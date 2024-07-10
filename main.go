package main

import (
	"fmt"
	"net/http"

	"github.com/Ccj-pro/admin_server/config"
	"github.com/Ccj-pro/admin_server/public/common"
	"github.com/Ccj-pro/admin_server/routes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 加载配置文件到全局配置结构体
	config.InitConfig()

	// 初始化日志
	common.InitLogger()

	// 初始化数据库(mysql)
	common.InitDB()

	// 注册所有路由
	r := routes.InitRoutes()

	host := "0.0.0.0"
	port := config.Conf.System.Port

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: r,
	}

	// go func() {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		common.Log.Fatalf("listen: %s\n", err)
	}
	// }()

	common.Log.Info(fmt.Sprintf("Server is running at http://%s:%d", host, port))
}
