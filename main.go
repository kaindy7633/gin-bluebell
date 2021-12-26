package main

import (
	"context"
	"fmt"
	"gin-bluebell/dao/mysql"
	"gin-bluebell/dao/redis"
	"gin-bluebell/logger"
	"gin-bluebell/pkg/snowflake"
	"gin-bluebell/routes"
	"gin-bluebell/settings"
	"gin-bluebell/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Go Web 开发通用脚手架模板

// @title gin-bluebell api docs
// @version 0.1.0
// @description gin-bluebell 基于Gin框架的Web通用脚手架模板
// @termsOfService http://swagger.io/terms/

// @contact.name Liuzhen
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/license-2.0.html

// @host 127.0.0.1:8088
// @BasePath /api/v1
func main() {
	// 通过执行参数指定配置文件
	// if len(os.Args) < 2 {
	// 	return
	// }

	// 1. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err: %v\n", err)
		return
	}

	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.AppConfig.Mode); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}
	defer zap.L().Sync()

	// 3. 初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err: %v\n", err)
		return
	}
	defer mysql.Close()

	// 4. 初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err: %v\n", err)
		return
	}
	defer redis.Close()

	// 初始化雪花算法 ID 生成函数
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err: %v\n", err)
		return
	}

	// 初始化翻译器
	if err := utils.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err: %v\n", err)
		return
	}

	// 5. 注册路由
	r := routes.Setup()

	// 6. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个 goroutine 启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	// 等待中断信号来地关闭服务器，为关闭服务器才做设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 Ctrl + c 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号， 但是不能被捕获，所以不需要添加它
	// signal.Notify 把接收到的 syscall.SIGINT 或 syscall.SIGTERM 信号转发给 quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
