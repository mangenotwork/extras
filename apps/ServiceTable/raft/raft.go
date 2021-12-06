package raft

import (
	"bytes"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/utils"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

// 全局变量
var (
	rafter *Raft
	MyAddr string
	cluster []string
	IsCluster bool // 是否集群
	once sync.Once
	ClusterTable map[string]*net.UDPAddr // 集群表
	ClusterTableLock sync.Mutex  // 集群锁

	electionTime = 5 // follower -> candidate 的最大时间, 单位秒
	heartBeatTime = 1 // leader 发送心跳周期
)

func InitRaft(){
	MyAddr = conf.Arg.Cluster.MyAddr
	cluster = strings.Split(conf.Arg.Cluster.InitCluster, ";")
	IsCluster = conf.Arg.Cluster.Open
}

type Raft struct {
	Conn *net.UDPConn

	myAddr string

	// 锁
	mu sync.Mutex

	// 节点编号
	me int

	// 当前任期
	leader string

	// 为哪个节点投票
	votedFor string

	// 当前节点状态
	// 0 follower  1 candidate  2 leader
	state int

	// 发送最后一条消息的时间
	lastMessageTime int64

	// 当前节点的领导
	currentLeader int

	// 当前选票
	vote int

	// 心跳信号
	heartBeat chan bool

	// 成功当选信号
	beElection chan bool

	// 投票信号
	electCh chan bool

	// 随机
	randEr *rand.Rand
}

// 启动集群, 参照raft算法 - leader 选举实现
func StartCluster() {

	if !IsCluster {
		return
	}

	once.Do(
		func() {
			rafter = &Raft{
				state:      0, // 初始化 follower状态
				myAddr:     MyAddr,
				heartBeat:  make(chan bool),
				beElection: make(chan bool),
				electCh:    make(chan bool),
				votedFor:   "",
				vote:       0,
				randEr:     rand.New(rand.NewSource(time.Now().UnixNano())),
				leader:     "",
			}
		},
	)

	listener, err := net.ListenUDP("udp", getUdpAddr(rafter.myAddr))
	if err != nil {
		panic(err)
	}

	rafter.Conn =  listener

	go join()
	go read()
	go election()

}


func getUdpAddr(str string) *net.UDPAddr {
	strList := strings.Split(str, ":")
	return &net.UDPAddr{IP: net.ParseIP(strList[0]), Port: utils.Str2Int(strList[1])}
}

func join() {
	ClusterTable = make(map[string]*net.UDPAddr, len(cluster))
	for _,v := range cluster{
		ClusterTableLock.Lock()
		ClusterTable[v] = getUdpAddr(v)
		ClusterTableLock.Unlock()
	}
}

func send(data []byte, to *net.UDPAddr) error {
	_, err := rafter.Conn.WriteToUDP(data, to)
	return err
}

func read() {

	for {
		data := make([]byte, 1024)
		n, remoteAddr, err := rafter.Conn.ReadFromUDP(data)
		if err != nil {
			log.Printf("error during read: %s", err)
			continue
		}
		log.Printf("receive %s from <%s>  leader = %s \n", data[:n], remoteAddr, rafter.leader)
		cmd := string(data[:n])

		// 收到拉票的消息
		if cmd == "拉票" {
			log.Println("remoteAddr = ", remoteAddr, "要你给他投票", "我的状态 = ", rafter.state, rafter.votedFor )

			// 已经我是leader, 通知你结束竞选
			if rafter.state == 3 {
				_= send([]byte("结束竞选"), remoteAddr)
			}

			// 一个节点某一任期内最多只能投一票
			// follower 才能投票
			if rafter.votedFor == "" && rafter.state == 0 {
				log.Println("是否有投票 = ", rafter.votedFor)
				rafter.votedFor = remoteAddr.String()

				_= send([]byte("给你投票"), remoteAddr)
			}

		}

		// 收到竞选结束的消息
		if cmd == "结束竞选" {
			rafter.beElection <- false
		}

		// 收到选票
		if cmd == "给你投票" {
			rafter.electCh <- true
		}

		// 公布新的leader
		if strings.Index(cmd, "leader") != -1{
			log.Println("有新的leader = ", cmd)
			rafter.state = 0 // 有新的leader,当前节点就是follower状态
		}

		// 来自leader的心跳
		if cmd == "心跳" {
			rafter.leader = remoteAddr.String() // 来自 leader 的心跳
			rafter.heartBeat <- true
		}
	}

}

func sendHeartBeat() {
	log.Println("广播心跳")
	for _, v := range ClusterTable {
		_=send([]byte("心跳"), v)
	}
}

// follower -> candidate
func election(){
	// 随机秒后将变为竞选者, 必须大于心跳时间, 如当前心跳为1, 则+2
	for {
		t := rafter.randEr.Intn(electionTime)+heartBeatTime+1
		tTime := time.Duration(t)
		log.Println("随机值t = ", tTime)
		timer := time.NewTimer(tTime * time.Second)
		select {
		case t := <-timer.C:
			// 2秒内收不到来自 leader 的心跳
			rafter.votedFor = ""
			rafter.state = 1
			rafter.leader = ""
			rafter.currentLeader = 0
			rafter.vote = 0

			log.Println("开始拉票 : ", t)
			canvass()

		case  <- rafter.heartBeat:
			// 重置
			log.Println("重置 = ", tTime)
			timer.Reset(tTime*time.Second)

		case <- rafter.electCh:
			log.Println("获得选票")
			rafter.vote++
			log.Println(rafter.vote+1, ",", len(ClusterTable)/2)
			// 这里做了一下改动:  +1 的目的是投自己一票, >=一半服务投票就竞选成功
			if rafter.vote+1 >= len(ClusterTable)/2 {
				// 结束竞选
				timer.Stop()
				electionEnd(true)
			}

		case be := <- rafter.beElection:
			// 结束竞选
			timer.Stop()
			electionEnd(be)

		}

		// 等待下次竞选
		rafter.state = 0
	}
}

func (rf *Raft) depiao(){
	rf.electCh <- true
}

func electionEnd(be bool){
	// 没竞选上的自行切换到follower
	if !be {
		log.Println("竞选失败, 别人已当选")
		rafter.state = 0
		rafter.vote = 0 // 重置票数
		return
	}

	rafter.state = 3
	// 给所有节点宣布
	for _, v := range ClusterTable {
		// TODO 日志复制
		_=send([]byte("leader:"+rafter.myAddr), v)
	}

	// 发送心跳
	sendHeartBeat()
	for {
		// 心跳时间要小于变为竞选者的时间周期
		timer := time.NewTimer(time.Duration(heartBeatTime) * time.Second)
		select {
		case <-timer.C:
			sendHeartBeat()
		}
	}

}

// 选举拉票,竞选 leader
// candidate -> leader
func canvass(){
	succeed := true
	// 给其他节点发送投票请求
	for k, v := range ClusterTable {
		if rafter.myAddr == k {
			continue
		}
		err := send([]byte("拉票"), v)
		log.Println("向",k,"进行拉票, err = ", err)
		if err != nil {
			succeed = false
			break
		}
	}
	if !succeed {
		log.Println("拉票失败,重新拉票")
		canvass()
	}
}


/* 日志

每个节点存储自己的日志副本(log[])，每条日志记录包含：

索引：该记录在日志中的位置
任期号：该记录首次被创建时的任期号
命令

*/

type LogData struct {
	Index int64
	Term string // 任期号
	Command string
}

func (data LogData) ToStr() string {
	var buffer bytes.Buffer
	buffer.WriteString(utils.Any2String(data.Index))
	buffer.WriteString("&")
	buffer.WriteString(data.Term)
	buffer.WriteString("&")
	buffer.WriteString(data.Command)
	buffer.WriteString("\n")
	return buffer.String()
}

func (data LogData) ToObj(str string){
	strList := strings.Split(str, "&")

	if len(strList) == 3 {
		data.Index = utils.Str2Int64(strList[0])
		data.Term = strList[1]
		data.Command = strList[2]
	}
}

// 追加写入日志
func (data LogData) Wait(){
	fileName := "log.data"

	var f *os.File
	var err error

	if checkFileExist(fileName) {  //文件存在
		f, err = os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, 0666) //打开文件
		if err != nil{
			log.Println("file open fail", err)
			return
		}
	}else {  //文件不存在
		f, err = os.Create(fileName) //创建文件
		if err != nil {
			log.Println("file create fail")
			return
		}
	}

	strTest := data.ToStr()

	//将文件写进去
	n, err1 := io.WriteString(f, strTest)
	if err1 != nil {
		log.Println("write error", err1)
		return
	}
	log.Println("写入的字节数是：", n)

	_=f.Close()
}

func checkFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
