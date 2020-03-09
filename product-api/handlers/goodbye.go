package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byee"))
}

/*
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) { //responseWriter is an interface while request is a struct
	g.l.Println("good bye world") //l is log.logger
	//http request takes r.body interface that uses io reader which returns data and errorif there is any
	d, err := ioutil.ReadAll(r.Body)
	//dealing with the error
	if err != nil {
		http.Error(rw, "OOppss, sth went wrong with the bye world", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "data %s\n", d) // rw responseWriter implements io interface, we can use standard string which will write out to serve HTTP
}
*/
