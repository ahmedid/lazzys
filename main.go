package main

import (
	"os"
	"flag"
	"fmt"
	"log"
	"strings"
	"net/http"
	"time"
)

var (
	port = flag.Int("p", 3080,"Server port to listen")
	pflag = flag.String("path", "", "URL Path")
	cflag = flag.Int("c", 200, "HTTP status code")
	bflag = flag.String("d", "", "The response body")
)

func usage(){
	fmt.Fprintf(os.Stderr, "usage: lazzys [-p port] [-path] [-c code] [-d body]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func getOriginalIP(r *http.Request) string {
        if forwardedIP := r.Header.Get("X-Forwarded-For"); forwardedIP != "" {
                return forwardedIP
        }

        return r.RemoteAddr
}

type Server struct {}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	w.WriteHeader(*cflag)
	size, _ := w.Write([]byte(*bflag))

	fmt.Printf("[%s] %s %s - (%s) %s - %d %d",
                startTime.Format("2006-01-02 15:04:05"),
                r.Method,
                r.URL.Path,
                r.Header.Get("User-Agent"),
                getOriginalIP(r),
                *cflag,
                size,
        )
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if !strings.HasPrefix(*pflag, "/") {
		*pflag = "/" + *pflag
	}

	http.Handle(*pflag, &Server{})
	if err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil); err != nil {
		log.Fatal(err)
	}
}
