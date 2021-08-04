package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"

	"github.com/sirupsen/logrus"
)

func ListenAndServe(listener net.Listener, handler Handler, closeChan <-chan struct{}) {
	closeFunc := func(msg string) {
		err := listener.Close()
		if err != nil {
			logrus.Errorf("accept error: %v", err)
		}
		err = handler.Close()
		if err != nil {
			logrus.Errorf("accept error: %v", err)
		}
	}
	defer closeFunc("abnormal shutting down...")
	go func() {
		<-closeChan
		closeFunc("shutting down...")
	}()
	ctx := context.Background()
	var wg sync.WaitGroup
	for {
		conn, err := listener.Accept()
		if err != nil {
			logrus.Fatalf("accept error: %v", err)
			break
		}
		wg.Add(1)
		go func() {
			handler.Handle(ctx, conn)
			wg.Done()
		}()
	}
	wg.Wait()
}

func ListenAndServeWithSignal(addr string, handler Handler) error {
	closeCh := make(chan struct{})
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	go func() {
		sig := <-sigCh
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP:
			closeCh <- struct{}{}
		}
	}()
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	logrus.Info(fmt.Sprintf("bind: %s, start listening...", addr))
	ListenAndServe(listener, handler, closeCh)
	return nil
}

type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Close() error
}

type EchoHandler struct {
	activeConn sync.Map
	closing    int32
}

type ClientConn struct {
	net.Conn
}

func (h *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	if atomic.LoadInt32(&h.closing) != 0 {
		err := conn.Close()
		if err != nil {
			logrus.Errorln(err)
		}
		return
	}
	clientConn := ClientConn{conn}
	h.activeConn.Store(clientConn, struct{}{})
	reader := bufio.NewReader(clientConn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				logrus.Infoln("connection close")
				h.activeConn.Delete(clientConn)
			} else {
				logrus.Warnln(err)
			}
			return
		}
		conn.Write([]byte("receive: " + msg))
	}
}

func main() {
	// ListenAndServe(":4567")
}
