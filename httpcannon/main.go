package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for {
			fmt.Println("http://127.0.0.1:4000")
			resp, err := http.Get("http://127.0.0.1:4000")
			if err != nil {
				panic(err)
			}
			_ = resp.Body.Close()
			time.Sleep(50 * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		for {
			fmt.Println("http://127.0.0.1:4000/about")
			resp, err := http.Get("http://127.0.0.1:4000/about")
			if err != nil {
				panic(err)
			}
			_ = resp.Body.Close()
			time.Sleep(50 * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		for {
			fmt.Println("http://127.0.0.1:4000/contact")
			resp, err := http.Get("http://127.0.0.1:4000/contact")
			if err != nil {
				panic(err)
			}
			_ = resp.Body.Close()
			time.Sleep(50 * time.Millisecond)
		}
	}()

	wg.Wait()

}
