package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mangenotwork/extras/apps/BlockWord/service"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/utils"
)


func Hello(w http.ResponseWriter, r *http.Request) {
	_,_= w.Write([]byte("Hello l'm block word.\n"+utils.Logo))
}

type BlockPostParam struct {
	Str string `json:"str"`
	Sub string `json:"sub"` // 替换符号
	Time string `json:"time"`
}

func Do(w http.ResponseWriter, r *http.Request) {
	params := &BlockPostParam{}
	httpser.GetJsonParam(r, params)
	params.Str, params.Time  = service.BlockWorkTrie.BlockWord(params.Str, params.Sub)
	httpser.OutSucceedBodyJsonP(w, params)
	return
}

func Add(w http.ResponseWriter, r *http.Request) {
	word := httpser.GetUrlArg(r, "word")
	service.AddWord(word)
	httpser.OutSucceedBodyJsonP(w,"")
	return
}

func Del(w http.ResponseWriter, r *http.Request) {
	word := httpser.GetUrlArg(r, "word")
	service.DelWord(word)
	httpser.OutSucceedBodyJsonP(w,"")
	return
}

func List(w http.ResponseWriter, r *http.Request) {
	httpser.OutSucceedBodyJsonP(w, service.GetWord())
	return
}

func WhiteAdd(w http.ResponseWriter, r *http.Request) {
	word := httpser.GetUrlArg(r, "word")
	service.WhiteAddWord(word)
	httpser.OutSucceedBodyJsonP(w,"")
	return
}

func WhiteDel(w http.ResponseWriter, r *http.Request) {
	word := httpser.GetUrlArg(r, "word")
	service.WhiteDelWord(word)
	httpser.OutSucceedBodyJsonP(w,"")
	return
}

func WhiteList(w http.ResponseWriter, r *http.Request) {
	httpser.OutSucceedBodyJsonP(w, service.WhiteGetWord())
	return
}

func IsHave(w http.ResponseWriter, r *http.Request){
	word := ""
	if r.Method == "POST" {
		decoder:=json.NewDecoder(r.Body)
		params := &BlockPostParam{}
		_=decoder.Decode(&params)
		word = params.Str
	} else {
		word = httpser.GetUrlArg(r, "str")
	}
	httpser.OutSucceedBodyJsonP(w, service.BlockWorkTrie.IsHave(word))
	return
}

func IsHaveList(w http.ResponseWriter, r *http.Request){
	word := ""
	if r.Method == "POST" {
		decoder:=json.NewDecoder(r.Body)
		params := &BlockPostParam{}
		_=decoder.Decode(&params)
		word = params.Str
	} else {
		word = httpser.GetUrlArg(r, "str")
	}
	httpser.OutSucceedBodyJsonP(w,service.BlockWorkTrie.BlockHaveList(word))
	return
}