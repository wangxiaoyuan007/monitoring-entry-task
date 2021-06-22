package main

import (
	"log"
	"math/rand"
	"monitoring-entry-task/router"
	"net/http"
	"time"
)

func main()  {
	go startClient()
	router.InitRouter()
}

func startClient() {
	sleepTime := 1000

	go func() {
		ticker := time.NewTicker(2 * time.Minute)
		for {
			<-ticker.C
			sleepTime = 200
			<-time.After(30 * time.Second)
			sleepTime = 1000
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				n := rand.Intn(100)
				if n >= 95 {
					sendErrRequest()
				} else {
					sendRequest()
				}

				time.Sleep((time.Duration)(sleepTime) * time.Millisecond)
			}
		}()
	}
}

func sendRequest() {
	resp, err := http.Get("http://localhost:9094/hello?name=tom")
	if err != nil {
		log.Println(err)
		return
	}
	resp.Body.Close()
}

func sendErrRequest() {
	resp, err := http.Get("http://localhost:9094/hello")
	if err != nil {
		log.Println(err)
		return
	}
	resp.Body.Close()
}