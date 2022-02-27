package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	// Remove first two args
	a := os.Args
	a = append(a[:0], a[2:]...)

	// Make sure message has a forward slash
	cmd := "/" + strings.TrimPrefix(strings.Join(a, " "), "/")

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	c, err := net.Dial("tcp", arguments[1])
	if err != nil || c == nil {
		fmt.Println(err)
		return
	}

	last := time.Now().UnixMilli()

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		scanner := bufio.NewScanner(c)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			last = time.Now().UnixMilli()
			fmt.Println(scanner.Text())
		}
	}()

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			now := time.Now().UnixMilli()

			// No message received for 500ms, exit
			if now-last > 500 {
				os.Exit(0)
			}
		}
	}()

	go func() {
		_, _ = fmt.Fprintln(c, cmd)
	}()

	wg.Wait()
}
