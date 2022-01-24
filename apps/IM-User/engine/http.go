package engine

import (
	"github.com/mangenotwork/extras/apps/IM-User/handler"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
)

func StartHTTP(){
	go func() {
		logger.Info("StartHTTP")

		mux := httpser.NewEngine()

		mux.Router("/register", handler.Register)

		mux.Run()

	}()
}