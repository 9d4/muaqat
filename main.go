package main

import (
	"log"
	"net"
	"sync"
	"time"
)

var (
	host string = "ladoosingh.herokuapp.com:80"
	path string = "/l/kVFQaOMXYpm"
	method string = "GET"
	routines int = 100
)

var wg = sync.WaitGroup{}

func main() {
	sendRequest := func(path string, method string, start chan int8) {
		wg.Add(1)
		num := <-start

		fn := func ()  {
			// open tcp connection to url
			conn, err := net.Dial("tcp", host)
			if err != nil {
				log.Fatal(err)
			}

			// send request
			conn.Write([]byte(method + " " + path + " HTTP/1.1\r\n"))
			conn.Write([]byte("Host: " + host + "\r\n"))
			conn.Write([]byte("Connection: close\r\n"))
			conn.Write([]byte("\r\n"))
			log.Println("sent. Routine:", num)
		}

		// make a ticker
		ticker := time.NewTicker(time.Millisecond * 100)
		
		for {
			<-ticker.C
			fn()
		}
	}

	var rs []chan int8

	// make routines
	for i := 0; i < routines; i++ {
		start := make(chan int8)
		rs = append(rs, start)

		go (func() {
			// send request
			sendRequest(path, method, start)
		})()
	}

	log.Println("Routines Created!")
	// set delay 5s
	time.Sleep(time.Second * 3)
	log.Println("Starting Routines!")

	// start all routines
	for i, r := range rs {
		r <- int8(i) 
	}
	wg.Wait()
}
