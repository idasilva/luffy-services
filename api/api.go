package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type API struct {
	mux     *gin.Engine
	context *Context
	logger  *zap.Logger
}

func (a *API) Server() error {
	a.logger.Info("Initializing server")
	a.context.handlers(a)
	fmt.Println(a.String())
	a.logger.Info("Starting HTTP server", zap.String("address", ":8080"))
	err := a.mux.Run()
	if err != nil {
		a.logger.Error("Failed to start server", zap.Error(err))
	}
	return err
}

func (a *API) Handlers(loadPaths func() []Routes) {
	a.context.Add(loadPaths)
}

func (a *API) String() string {
	return `
  _   _   _ ___ _____   __  ___ ___ _____   _____ ___ ___ ___ 
 | | | | | | __| __\ \ / / / __| __| _ \ \ / /_ _/ __| __/ __|
 | |_| |_| | _|| _| \ V /  \__ \ _||   /\ V / | | (__| _|\__ \
 |____\___/|_| |_|   |_|   |___/___|_|_\ \_/ |___\___|___|___/

`

}

func New() *API {
	logger, _ := zap.NewProduction()
	logger.Info("Creating new API instance")
	return &API{
		mux:     gin.Default(),
		context: NewContext(),
		logger:  logger,
	}

}
