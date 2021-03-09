package graphics

import (
	log "github.com/sirupsen/logrus"
	"github.com/vehsamrak/game-roguelike-lonely/internal/logger"
)

func Render() {
	log.SetFormatter(&logger.TextFormatter{})
	log.Info("Client started\n")
	log.Info("Client finished\n")
}
