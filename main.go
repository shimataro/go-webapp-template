package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-webapp-template/config"
	"go-webapp-template/controllers"
	"go-webapp-template/libs/cologger"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type options struct {
	fd *uint
}

const configDirName = "configFiles"

func main() {
	opt := parseOptions()

	err := config.Load(configDirName)
	if err != nil {
		panic(err)
	}

	listener, err := createListener(*opt.fd)
	if err != nil {
		panic(err)
	}
	server := createServer()
	runServer(server, listener)

	setupGracefulShutdown(server)
}

func parseOptions() options {
	opt := options{
		fd: flag.Uint("fd", 0, "file descriptor to listen and serve"),
	}
	flag.Parse()

	return opt
}

func createListener(fd uint) (net.Listener, error) {
	if fd != 0 {
		// ファイルディスクリプタをListen
		name := fmt.Sprintf("fd@%d", fd)
		file := os.NewFile(uintptr(fd), name)
		cologger.Infof("Listening on %s...", name)
		return net.FileListener(file)
	} else {
		// TCPポートをListen
		conf := config.Get()
		return net.Listen("tcp", conf.Settings.Address)
	}
}

func createServer() *http.Server {
	engine := gin.Default()
	controllers.Routing(engine)

	return &http.Server{
		Handler: engine,
	}
}

func runServer(server *http.Server, listener net.Listener) {
	// 先に setupGracefulShutdown() のシグナル待ちできるように、goroutineとして遅延実行
	go func() {
		err := server.Serve(listener)
		if err == nil {
			return
		}
		if err == http.ErrServerClosed {
			return
		}
		panic(err)
	}()
}

func setupGracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	cologger.Info("Shutting down server...")

	// 処理中のリクエストを終了するまでの猶予
	conf := config.Get()
	ctx, cancel := context.WithTimeout(context.Background(), conf.Settings.Timeout*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		cologger.Fatal("Server forced to shutdown:", err)
	}

	cologger.Info("🎉Shutting down completed.🎉")
}
