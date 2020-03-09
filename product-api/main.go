package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"workspace/working/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l) //hello handler
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux() // serveMux itself is a handler

	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh) // here goodbye is the example of endpoint
	//below is the manual server
	s := &http.Server{ // here s is the server
		Addr:         ":8881",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//listenAndServe is automatic server
	//http.ListenAndServe(":8881", sm) //listenAndServe is using the default serveMux as routeHandler to determine which code block gets executed when a request is made

	//graceful shutdown to allow the clients finish their task during our upgrading of the server
	//handling listen and serve in go func so its not gonna block
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal) //channel is of type os.Signal
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	//this is a block because reading from channel needs the message to be availible to be displayed
	l.Println("recived terminate, graceful shutdown", sig)
	//context timeout returns a function as well
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
