package main

import (
	"github.com/WiFeng/go-sky"
	"github.com/xiwujie/article/pkg/config"
	"github.com/xiwujie/article/pkg/endpoint"
	"github.com/xiwujie/article/pkg/service"
	"github.com/xiwujie/article/pkg/task"
	"github.com/xiwujie/article/pkg/transport/http"
)

func main() {

	var (
		service     = service.New()
		endpoints   = endpoint.New(service)
		httpHandler = http.NewHandler(endpoints)
	)

	sky.LoadAppConfig(&config.GlobalAppConfig)
	// 启动后定时执行的任务
	sky.RegisterTask(task.Start, nil, true)
	sky.Run(httpHandler)
}
