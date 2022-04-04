package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./scanner <ip>")
		return
	}
	var openPort []int
	wg := sync.WaitGroup{}
	for port := 0; port < 65536; port++ {
		wg.Add(1)
		go func(p int) {
			if open(p) {
				openPort = append(openPort, p)
			}
			defer wg.Done()
		}(port)
	}
	wg.Wait()
	for _, v := range openPort {
		fmt.Printf("Port %d: success\n", v)
	}
}

func open(port int) (flag bool) {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", os.Args[1], port), time.Millisecond*200)
	if err == nil {
		flag = true
	}
	return
}
