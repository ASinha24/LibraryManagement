package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ASinha24/LibraryManagementSystem"
	"github.com/ASinha24/LibraryManagementSystem/bookStore"
	phttp "github.com/ASinha24/LibraryManagementSystem/http"
	"github.com/gorilla/mux"
)

var port = flag.String("port", "8080", "port to listen")

func main() {
	flag.Parse()
	router := mux.NewRouter().StrictSlash(true)
	bookstore := bookStore.NewBookStore()
	bookService := LibraryManagementSystem.NewBookService(bookstore)
	bookHandler := phttp.NewbookHandler(bookService, bookstore)
	bookHandler.MuxInstaller(router)

	//server starting stopping gracefully
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("err in listening port: %s\n", err)
		}
	}()
	log.Print("Server Started")
	<-done
	log.Print("server stopper")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")

}
