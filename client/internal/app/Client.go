package app

import (
	log "github.com/sirupsen/logrus"
	"github.com/vehsamrak/game-roguelike-lonely/client/internal/logger"
)

type Client struct {
}

func (client Client) Start() {
	log.SetFormatter(&logger.TextFormatter{})
	log.Info("Game client started")

	runRaycaster()

	log.Info("Game client finished")
}
