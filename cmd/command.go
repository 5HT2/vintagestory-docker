package main

import (
	"C"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/sys/unix"
)

func main() {
	//
	// Setup terminal
	//

	saved, err := tcget(os.Stdin.Fd())
	if err != nil {
		panic(err)
	}
	defer func() {
		tcset(os.Stdin.Fd(), saved)
	}()

	raw := makeraw(*saved)
	tcset(os.Stdin.Fd(), &raw)

	//
	// Setup args
	//

	// Remove first two args
	a := os.Args
	a = append(a[:0], a[2:]...)

	// Make sure message has a forward slash
	cmd := "/" + strings.TrimPrefix(strings.Join(a, " "), "/")

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port.")
		return
	}

	//
	// Setup TCP
	//

	c, err := net.Dial("tcp", "localhost:"+arguments[1])
	if err != nil || c == nil {
		fmt.Println(err)
		return
	}

	//
	// Connect to Docker, run command and get output, before exiting
	//

	last := time.Now().UnixMilli()
	var wg sync.WaitGroup
	wg.Add(3)

	// Read stdout from Docker
	go func() {
		scanner := bufio.NewScanner(c)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			last = time.Now().UnixMilli()
			io.Copy(os.Stdout, bytes.NewReader([]byte(scanner.Text()+"\r\n")))
		}
	}()

	// Check timeout of stdout lines
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			now := time.Now().UnixMilli()

			// No message received for 500ms, exit
			if now-last > 150 {
				wg.Done()
			}
		}
	}()

	// Run command
	go func() {
		_, _ = fmt.Fprintln(c, cmd)
		wg.Done()
	}()

	wg.Wait()
}

//
// Set current terminal interface; Docs: https://man7.org/linux/man-pages/man3/tcflow.3.html
//

func tcget(fd uintptr) (*unix.Termios, error) {
	termios, err := unix.IoctlGetTermios(int(fd), unix.TCGETS)
	if err != nil {
		return nil, err
	}
	return termios, nil
}

func tcset(fd uintptr, p *unix.Termios) error {
	return unix.IoctlSetTermios(int(fd), unix.TCSETS, p)
}

func makeraw(t unix.Termios) unix.Termios {
	t.Iflag &^= unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON
	t.Oflag &^= unix.OPOST
	t.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN
	t.Cflag &^= unix.CSIZE | unix.PARENB
	t.Cflag &^= unix.CS8
	t.Cc[unix.VMIN] = 1
	t.Cc[unix.VTIME] = 0
	return t
}
