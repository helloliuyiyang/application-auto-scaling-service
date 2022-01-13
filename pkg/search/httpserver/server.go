package httpserver

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"

	"nanto.io/application-auto-scaling-service/pkg/common/config"
	"nanto.io/application-auto-scaling-service/pkg/common/utils/logutil"
)

const (
	defaultKeepAlivePeriod = 3 * time.Minute
)

var logger = logutil.GetLogger()

// HttpServer contains a http server
type HttpServer struct {
	httpServer *http.Server
	listener   net.Listener
}

// Start snapshot API for resource-collector
func StartHttpServer(conf *config.HttpServerConf, stopCh <-chan struct{}) error {
	Register()
	hs, err := NewHttpServer(conf)
	if err != nil {
		logger.Errorf("new http server err: %v", err)
		return err
	}
	logger.Infof("Start http server listen[%s]", hs.httpServer.Addr)
	return hs.BlockingRun(stopCh)
}

func getServerAddress(conf *config.HttpServerConf) string {
	hostIP := conf.HttpAddr
	if hostIP == "" {
		logger.Errorf("server IP address not configured.")
		os.Exit(1)
	}
	port := conf.HttpPort
	return fmt.Sprintf("%s:%d", hostIP, port)
}

// NewHttpServer construct a new http server with ssl certificate
func NewHttpServer(conf *config.HttpServerConf) (*HttpServer, error) {
	hs := &HttpServer{}

	httpAddr := getServerAddress(conf)
	l, err := net.Listen("tcp", httpAddr)
	if err != nil {
		logger.Errorf("failed to http(https) listen: %s err: %s", httpAddr, err.Error())
		return nil, err
	}
	hs.listener = l
	hs.httpServer = &http.Server{
		Addr:           l.Addr().String(),
		MaxHeaderBytes: 1 << 20,
	}

	return hs, nil
}

// BlockingRun make server running with blocking and shutdown with certain time timeout
func (hs *HttpServer) BlockingRun(stopCh <-chan struct{}) error {
	return RunServer(hs.httpServer, hs.listener, 60, stopCh)
}

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
//
// Copied from Go 1.7.2 net/http/server.go
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}
	err = tc.SetKeepAlive(true)
	if err != nil {
		logger.Infof("Meet error when setting KeepAlive as true: %s", err.Error())
	}
	err = tc.SetKeepAlivePeriod(defaultKeepAlivePeriod)
	if err != nil {
		logger.Infof("Meet error when setting KeepAlivePeriod (%s): %s", defaultKeepAlivePeriod, err.Error())
	}
	return tc, nil
}

// RunServer run server gracefully
func RunServer(
	server *http.Server,
	ln net.Listener,
	shutDownTimeout time.Duration,
	stopCh <-chan struct{}) error {
	if ln == nil {
		return errors.New("listener must not be nil")
	}

	// Shutdown server gracefully.
	go func() {
		<-stopCh
		logger.Infof("Shutdown Server gracefully...")
		ctx, cancel := context.WithTimeout(context.Background(), shutDownTimeout)
		err := server.Shutdown(ctx)
		if err != nil {
			logger.Errorf("Shutdown Server failed, err: %v", err)
		}
		cancel()
	}()

	var listener net.Listener
	listener = tcpKeepAliveListener{ln.(*net.TCPListener)}

	if server.TLSConfig != nil {
		listener = tls.NewListener(listener, server.TLSConfig)
	}

	err := server.Serve(listener)
	if err != nil {
		logger.Errorf("Server runs failed, err: %v", err)
	}

	msg := fmt.Sprintf("Stopped listening on %s", ln.Addr().String())
	select {
	case _, ok := <-stopCh:
		if !ok {
			return nil
		}
		fmt.Println("continue")
	default:
		errMsg := fmt.Sprintf("%s due to error: %v", msg, err)
		fmt.Println(errMsg)
		os.Exit(1)
	}

	return nil
}
