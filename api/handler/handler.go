package handler

import (
	"main/config"
	"main/pkg/logger"
	"main/storage"
)

type Handler struct {
	strg storage.StoregeI
	cfg  config.Config
	log  logger.LoggerI
}

func NewHandler(strg storage.StoregeI, conf config.Config, loger logger.LoggerI) *Handler {
	return &Handler{strg: strg, cfg: conf, log: loger}
}
