package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	ip, max := parse()
	if ip == "" {
		return
	}

	var openPort []int
	wg := sync.WaitGroup{}
	//scan port
	fmt.Printf("Scanning %s from 0 to %d\n", ip, max)
	for port := 0; port < max; port++ {
		wg.Add(1)
		go func(p int) {
			if open(ip, p) {
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

func parse() (ip string, max int) {
	// Parse parameters
	if len(os.Args) != 2 && len(os.Args) != 4 {
		fmt.Println("Usage: ./scanner <ip> [option]\n\t-m\tMaximum port number")
		return
	}

	ip = os.Args[1]
	max = 65536
	var err error

	if len(os.Args) == 4 {
		if os.Args[2] == "-m" {
			max, err = strconv.Atoi(os.Args[3])
			if err != nil || max <= 0 {
				fmt.Println("port number must be a positive number which is not greater than 65536")
				ip = ""
				return
			}

		} else {
			fmt.Println("Usage: ./scanner <ip> [option]\n\t-m\tMaximum port number")
			ip = ""
			return
		}
	}
	return
}

func open(ip string, port int) (flag bool) {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Millisecond*200)
	if err == nil {
		flag = true
	}
	return
}
