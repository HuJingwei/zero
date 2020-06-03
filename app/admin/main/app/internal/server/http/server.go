package http

import (
	"net/http"

	pb "github.com/HuJingwei/go-zero/app/admin/main/app/api"
	"github.com/HuJingwei/go-zero/app/admin/main/app/internal/model"
	"github.com/HuJingwei/go-zero/pkg/conf/paladin"
	"github.com/HuJingwei/go-zero/pkg/log"
	bm "github.com/HuJingwei/go-zero/pkg/net/http/blademaster"
)

var svc pb.DemoServer

// New new a bm server.
func New(s pb.DemoServer) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	engine = bm.DefaultServer(&cfg)
	pb.RegisterDemoBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/app")
	{
		g.GET("/start", howToStart)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}