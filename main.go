package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shy0526/gin-embed-vue/common/logger"
	"github.com/shy0526/gin-embed-vue/common/router"
	"github.com/shy0526/gin-embed-vue/common/settings"
	"golang.org/x/sync/errgroup"
)

//go:embed static
var staticFS embed.FS

var (
	g errgroup.Group
)

func main() {
	if err := settings.Init(); err != nil {
		panic(fmt.Errorf("fatal error init settings: %w", err))
	}

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		panic(fmt.Errorf("fatal error init logger: %w", err))
	}

	//cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.Conf.Port),
		Handler:        router.Router(),
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
		//TLSConfig: &tls.Config{
		//	Certificates: []tls.Certificate{cert},
		//},
	}

	g.Go(func() error {
		return s.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		os.Exit(1)
	}

	fmt.Println("server running. press ctrl + c to exit.")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		fmt.Printf("server shutdown failed, err:%v", err)
	}

	fmt.Println("Server exiting")

}
