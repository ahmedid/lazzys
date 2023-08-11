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


type headerArgs []string

func (h *headerArgs) Set(value string) error {
	*h = append(*h, value)
	return nil
}

func (h headerArgs) String() string {
	return strings.Join(h, ",")
}

var (
	headers headerArgs
	port = flag.Int("p", 3080,"Server port to listen.")
	pflag = flag.String("path", "", "URL Path.")
	cflag = flag.Int("c", 200, "HTTP status code.")
	bflag = flag.String("d", "", "Response data.")
)

func ParseFlags(){
	flag.Var(&headers, "H", "Header `\"Name: Value\"`. Multiple -H flags are accepted.")
	flag.Usage = usage
	flag.Parse()
}

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

	if len(headers) != 0 {
		for _, header := range headers {
			kv := strings.Split(header, ":")
			w.Header().Set(kv[0], kv[1])
		}
	}
	w.WriteHeader(*cflag)
	size, _ := w.Write([]byte(*bflag))

	fmt.Printf("[%s] %s %s - (%s) %s - %d %d\n",
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
	ParseFlags()
	if !strings.HasPrefix(*pflag, "/") {
		*pflag = "/" + *pflag
	}

	http.Handle(*pflag, &Server{})
	if err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil); err != nil {
		log.Fatal(err)
	}
}
