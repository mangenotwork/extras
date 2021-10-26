package handler

import (
	"encoding/json"
	"fmt"
	"github.com/mangenotwork/extras/apps/BlockWord/model"
	"net/http"

	"github.com/mangenotwork/extras/apps/BlockWord/service"
	"github.com/mangenotwork/extras/common/utils"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_,_= w.Write([]byte("Hello l'm block word.\n"+utils.Logo))
}

type BlockPostParam struct {
	Str string `json:"str"`
	Sub string `json:"sub"` // 替换符号
}

func Do(w http.ResponseWriter, r *http.Request) {
	decoder:=json.NewDecoder(r.Body)
	params := &BlockPostParam{}
	decoder.Decode(&params)
	params.Str = service.BlockWorkTrie.BlockWord(params.Str, params.Sub)
	fmt.Println("POST json req: ",params.Str)
	utils.OutSucceedBodyJsonP(w, params)
	return
}

func Add(w http.ResponseWriter, r *http.Request) {
	word := utils.GetUrlArg(r, "word")
	if err := service.AddWord(word); err != nil {
		utils.OutErrBody(w, 201, err)
		return
	}
	utils.OutSucceedBodyJsonP(w,"")
	return
}

func Del(w http.ResponseWriter, r *http.Request) {
	word := utils.GetUrlArg(r, "word")
	if err := service.DelWord(word); err != nil {
		utils.OutErrBody(w, 201, err)
		return
	}
	utils.OutSucceedBodyJsonP(w,"")
	return
}

func List(w http.ResponseWriter, r *http.Request) {
	utils.OutSucceedBodyJsonP(w, service.GetWord())
	return
}

func WhiteAdd(w http.ResponseWriter, r *http.Request) {
	word := utils.GetUrlArg(r, "word")
	model.WhiteWord.Insert(word)
	utils.OutSucceedBodyJsonP(w,"")
	return
}

func WhiteDel(w http.ResponseWriter, r *http.Request) {
	word := utils.GetUrlArg(r, "word")
	model.WhiteWord.Remove(word)
	utils.OutSucceedBodyJsonP(w,"")
	return
}

func WhiteList(w http.ResponseWriter, r *http.Request) {
	for _, v := range []string{"路口", "路上", "口交", "赶路", "交通","费出口"} {
		model.WhiteWord.Insert(v)
	}
	model.WhiteWord.WordList()
	utils.OutSucceedBodyJsonP(w,"")
	return
}