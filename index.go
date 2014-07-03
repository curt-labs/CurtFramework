package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var (
	port = flag.String("port", "8000", "Port for the application to start on")
)

func main() {
	flag.Parse()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusFound, "Welcome home...")
	})

	s := &http.Server{
		Addr:           ":" + *port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Started on :%s\n", *port)
	s.ListenAndServe()

}
