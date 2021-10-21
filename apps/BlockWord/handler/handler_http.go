package handler

import (
	"encoding/json"
	"fmt"
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
	params.Str = service.BlockWorkTrie.Replace(params.Str, params.Sub)
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