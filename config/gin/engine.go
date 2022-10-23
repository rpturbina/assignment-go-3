package gin

import (
	"log"

	"github.com/gin-gonic/gin"
)

type HttpServer interface {
	GetGin() *gin.Engine
	Serve()
}

type Config struct {
	Port string
}

type GinImpl struct {
	engine *gin.Engine
	config Config
}

func (g *GinImpl) GetGin() *gin.Engine {
	return g.engine
}

func (g *GinImpl) Serve() {
	if err := g.engine.Run(g.config.Port); err != nil {
		log.Panic(err)
	}
}

func NewGinHttp(conf Config) HttpServer {
	ginEngine := gin.Default()
	return &GinImpl{engine: ginEngine, config: conf}
}
