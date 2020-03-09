package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// in order for us to use the http handler, we need to create a struct that uses the interface
//handler is just an http interface
//handleFunc is a convenience method that takes the function and creats an http handler from it and adds it to the default serveMux.
//hello is a simple handler
type Hello struct {
	//dependencies injection here later allows more freedom and control over things when these things are connected to the database. this allows us to do good unit testing

	l *log.Logger
}

//returns Hello handler as reference
//NewHello
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) { //responseWriter is an interface while request is a struct
	h.l.Println("hello world") //l is log.logger
	//http request takes r.body interface that uses io reader which returns data and error if there is any
	//read the body
	b, err := ioutil.ReadAll(r.Body)
	//dealing with the error
	if err != nil {
		http.Error(rw, "OOppss, sth went wrong", http.StatusBadRequest)
		return
	}
	//write the response
	fmt.Fprintf(rw, "data %s\n", b) // rw responseWriter implements io interface, we can use standard string which will write out to serve HTTP
}
