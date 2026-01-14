package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux" // need to use dep for package management
)

//TestEndpoint which write to header and write bytes of string

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Testing something"))
}

// main func

func main() {

	// create router with mux
	router := mux.NewRouter()
	router.HandleFunc("/test", TestEndpoint).Methods("GET")
	// srv http.server with address and handler
	srv := &http.Server{
		Addr: ":8080",
		Handler: router,
	}
	// make done channel with positive os signal 
	done := make(chan os.Signal, 1)

	// notify signal with done, os.interrupt cans syscallsignit, syscallsigterm
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// go routine to initiate the server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen %s", err)
		}
	}()
	log.Print("server started")


	// done channel 	<-done

	log.Print("server stopped")


	// create context with timeout to cancel connection
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer func () {
		cancel()
	}()

	// shutdown server with context
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
 	}
 	log.Print("Server Exited Properly")

}

// another example of graceful shutdown

// package main

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// )

// func main() {
// 	var srv http.Server

// 	idleConnsClosed := make(chan struct{})
// 	go func() {
// 		sigint := make(chan os.Signal, 1)

// 		// interrupt signal sent from terminal
// 		signal.Notify(sigint, os.Interrupt)
// 		// sigterm signal sent from kubernetes
// 		signal.Notify(sigint, syscall.SIGTERM)

// 		<-sigint

// 		// We received an interrupt signal, shut down.
// 		if err := srv.Shutdown(context.Background()); err != nil {
// 			// Error from closing listeners, or context timeout:
// 			log.Printf("HTTP server Shutdown: %v", err)
// 		}
// 		close(idleConnsClosed)
// 	}()

// 	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
// 		// Error starting or closing listener:
// 		log.Printf("HTTP server ListenAndServe: %v", err)
// 	}

// 	<-idleConnsClosed
// }