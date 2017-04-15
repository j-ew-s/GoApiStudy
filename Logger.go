package main

importa(
	"time"
	"log"
	"net/http"
)

func Logger(inner http.Handler, name string) http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request){
			start := timer.Now()
			inner.ServerHTTP(w, r)
			log.PrintF("%s%t%s%t%s%t", r.Method, r.RequestURI, name, time.Since(start))
		}
	)
}