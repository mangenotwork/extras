package handler

import (
	"github.com/mangenotwork/extras/common/utils"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_,_= w.Write([]byte("Hello l'm block word.\n"+utils.Logo))
}