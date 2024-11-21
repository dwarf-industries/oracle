package commands

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"oracle/controllers"
	"oracle/middlewhere"
)

type SyncCommand struct {
}

func (s *SyncCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "sync [port]",
		Short: "starts the oracle and connects to the network",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			port := args[0]
			s.Execute(&port)
		},
	}
}

func (s *SyncCommand) Execute(port *string) {
	router := gin.New()
	router.Use(middlewhere.Cors())
	v1 := router.Group("/v1")

	nodesController := controllers.NodesController{}
	nodesController.Init(v1)

	srv := &http.Server{
		Addr:    *port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	got := <-quit
	fmt.Println(got)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	connectionDone := <-ctx.Done()
	fmt.Println(connectionDone)
	log.Println("Server exiting")
}
