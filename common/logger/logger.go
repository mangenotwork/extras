package logger

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/*

	1. 单机 - 终端输出

	2. 单机 - 输出文件

	3. 网络 - 输出到日志服务

*/


var std = newStd()

type logger struct {
	appName string
	terminal bool
	outFile bool
	outFileWriter *os.File
	outService bool // 日志输出到服务
	outServiceIp string
	outServicePort int
	outServiceConn  *net.UDPConn
	outServiceLevel []int
}

func newStd() *logger {
	return &logger{
		terminal: true,
		outFile: false,
		outService: false,
		outServiceLevel: []int{3, 4, 5},
	}
}

func SetLogFile(name string) {
	std.outFile = true
	std.appName = name
	std.outFileWriter, _ = os.OpenFile( name+time.Now().Format("-20060102")+".log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}

func SetAppName(name string) {
	std.appName = name
}

func SetOutService(ip string, port int) {
	var err error
	std.outService = true
	std.outServiceIp = ip
	std.outServicePort = port
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: net.ParseIP(std.outServiceIp), Port: std.outServicePort}
	std.outServiceConn, err = net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		Error(err)
	}
}

func SetOutServiceWarn2Panic() {
	std.outServiceLevel = []int{3, 4, 5}
}

func SetOutServiceInfo2Panic() {
	std.outServiceLevel = []int{1, 2, 3, 4, 5}
}

func DisableTerminal() {
	std.terminal = false
}

type Level int

var LevelMap = map[Level]string {
	1 : "Info  ",
	2 : "Debug ",
	3 : "Warn  ",
	4 : "Error ",
	5 : "Panic ",
}

func (l *logger) Log(level Level, args string, times int) {
	var buffer bytes.Buffer
	buffer.WriteString(time.Now().Format("2006-01-02 15:04:05 |"))
	buffer.WriteString(LevelMap[level])
	_, file, line, _ := runtime.Caller(times)
	buffer.WriteString("|")
	buffer.WriteString(file)
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(line))
	buffer.WriteString(" : ")
	buffer.WriteString(args)
	buffer.WriteString("\n")
	out := buffer.Bytes()
	if l.terminal {
		_,_ = buffer.WriteTo(os.Stdout)
	}

	go func(out []byte) {
		if l.outFile {
			_,_ = l.outFileWriter.Write(out)
		}

		if l.outService {
			for _, v := range l.outServiceLevel {
				if Level(v) == level {
					out = append([]byte("1"+l.appName+"|"), out...)
					_,_ = l.outServiceConn.Write(out)
				}
			}
		}
	}(out)

}

func Info(args ...interface{}) {
	std.Log(1, fmt.Sprint(args...), 2)
}

func Infof(format string, args ...interface{}) {
	std.Log(1, fmt.Sprintf(format, args...), 2)
}

func InfoTimes(times int, args ...interface{}) {
	std.Log(1, fmt.Sprint(args...), times)
}

func Debug(args ...interface{}) {
	std.Log(2, fmt.Sprint(args...), 2)
}

func Debugf(format string, args ...interface{}) {
	std.Log(2, fmt.Sprintf(format, args...), 2)
}

func DebugTimes(times int, args ...interface{}) {
	std.Log(2, fmt.Sprint(args...), times)
}

func Warn(args ...interface{}) {
	std.Log(3, fmt.Sprint(args...), 2)
}

func Warnf(format string, args ...interface{}) {
	std.Log(3, fmt.Sprintf(format, args...), 2)
}

func WarnTimes(times int, args ...interface{}) {
	std.Log(3, fmt.Sprint(args...), times)
}

func Error(args ...interface{}) {
	std.Log(4, fmt.Sprint(args...), 2)
}

func Errorf(format string, args ...interface{}) {
	std.Log(4, fmt.Sprintf(format, args...), 2)
}

func ErrorTimes(times int, args ...interface{}) {
	std.Log(4, fmt.Sprint(args...), times)
}

func Panic(args ...interface{}){
	std.Log(5, fmt.Sprint(args...), 2)
	panic(args)
}

func (l *logger) Http(log string) {
	if l.outService {
		var out bytes.Buffer
		out.WriteString("2"+l.appName+"|")
		out.WriteString(log)
		_,_ = l.outServiceConn.Write(out.Bytes())
	}
}

func Http(log string, show bool) {
	if show {
		Info(strings.Replace(log, "#", " | ", -1) + " ms")
	}
	go func() {
		std.Http(log)
	}()

}