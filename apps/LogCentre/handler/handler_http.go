package handler

import (
	"github.com/mangenotwork/extras/common/boltdb"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
	"net/http"
)


var BoltdbFileName = "data.db"

func Hello(w http.ResponseWriter, r *http.Request) {
	_,_= w.Write([]byte("Hello l'm block word.\n"+utils.Logo))
}

func GetLogTable(w http.ResponseWriter, r *http.Request) {
	bo, err := boltdb.NewBoltDB(BoltdbFileName)
	defer bo.Close()
	if err != nil {
		logger.Error(err)
	}
	data, err := bo.GetTable()
	if err != nil {
		httpser.OutErrBody(w, 1, err)
		return
	}
	httpser.OutSucceedBody(w, data)
	return
}

func HttpReqLog(w http.ResponseWriter, r *http.Request) {
	start := httpser.GetUrlArg(r, "start")
	end := httpser.GetUrlArg(r, "end")
	table := httpser.GetUrlArg(r, "table")

	startKey := utils.Str2TimestampStr(start)
	endKey := utils.Str2TimestampStr(end)
	logger.Error(startKey, endKey)

	bo, err := boltdb.NewBoltDB(BoltdbFileName)
	if err != nil {
		logger.Error(err)
	}
	defer bo.Close()


	//data := bo.SelectFront(table, 100)
	data, err := bo.SelectInterval(table, startKey, endKey)
	if err != nil {
		httpser.OutErrBody(w, 1, err)
		return
	}
	httpser.OutSucceedBody(w, data)
	return
}

