package tls

import (
	"crypto/rand"
	"crypto/tls"
	"errors"
	"net"
)

func Listen(network, laddr string) (net.Listener, error) {
	cert, err := tls.LoadX509KeyPair("./certs/server.pem", "./certs/server.key")
	if err != nil {
		panic(err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader

	lis, err := tls.Listen(network, laddr, &config)
	if err != nil {
		return nil, errors.New("wrpc tls listen: listen failed: " + err.Error())
	}

	return lis, err
}

func Dial(network, laddr string) (*tls.Conn, error) {
	cert, err := tls.LoadX509KeyPair("./certs/client.pem", "./certs/client.key")
	if err != nil {
		panic(err)
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	return tls.Dial(network, laddr, &config)
}
