package main

import (
	"confluence-checkout/container"
	"confluence-checkout/internal/infrastructure/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	err := config.LoadGlobalEnv(".")
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	g := gin.Default()

	initializer := container.NewInitializer(config.GlobalEnv)

	container.Router(g, initializer)

	g.Run(":" + config.GlobalEnv.Port)

}

// protoc --plugin=protoc-gen-ts_proto=".\\node_modules\\.bin\\protoc-gen-ts_proto.cmd" --ts_proto_out=./src ./proto/invoicer.proto
