package websocket_l

import (
	"fmt"
	"net"
)

/*
	- Connecting to a TCP socket
	To establish a TCP connection, a Go client uses the "DialTCP" function in the "net" package which return TCPConn object.

	When a connection is established, the client and server begin exchanging data:
	the client sends a requrest to the server through a TCPConn object, the server parses
	the requrest and sends a response, and the TCPConn object receives the response from the server.

	This connection remains valid until the client or server closes it. The functions
	for creating a connection are as follows:

*/

var addr = net.TCPAddr{
	IP:   net.IPv4(127, 0, 0, 1),
	Port: 9001,
}

func tcp_socket_server(listener net.Listener) {
	conn, err := listener.Accept()
	if err != nil {
		// handle error
		panic(fmt.Errorf("server side Accept error: %v", err))
	}

	// send message
	if _, err := conn.Write([]byte("ping")); err != nil {
		// handle error
		panic(fmt.Errorf("server side send error: %v", err))
	}

	// receive message
	buf := make([]byte, 512)
	n, err := conn.Read(buf[0:])
	if err != nil {
		// handle error
		panic(fmt.Errorf("server side receive error: %v", err))
	}

	fmt.Printf("server get message: %v\n", string(buf[:]))
	fmt.Printf("tcp socket server: %v\n", n)
}

func tcp_socket_client() {
	tcpAddr_client, err := net.ResolveTCPAddr("tcp", addr.String())
	if err != nil {
		// handle error
		panic(fmt.Errorf("client side resolveTcpAddr error: %v", err))
	}
	fmt.Println("client dialTCP")
	conn, err := net.DialTCP("tcp", nil, tcpAddr_client)
	if err != nil {
		// handle error
		panic(fmt.Errorf("DialTCP error: %v", err))
	}

	// send message
	_, err = conn.Write([]byte("pong"))
	if err != nil {
		// handle error
		panic(fmt.Errorf("client send message error: %v", err))
	}

	// receive message
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		// handle error
		panic(fmt.Errorf("client receive message error: %v", err))
	}
	fmt.Printf("client get message: %v\n", string(buf[:]))
	fmt.Printf("tcp socket client: %v\n", n)

}

func Execute_tcp_socket_server() {
	tcpAddr_server, err := net.ResolveTCPAddr("tcp", addr.String())
	if err != nil {
		// handle error
		panic(fmt.Errorf("server side resolveTcpAddr error: %v", err))
	}

	listener, err := net.ListenTCP("tcp", tcpAddr_server)
	fmt.Println("server listening")
	if err != nil {
		// handle error
		panic(fmt.Errorf("server side ListenTCP error: %v", err))
	}
	go tcp_socket_server(listener)
}

func Execute_tcp_socket_client() {
	go tcp_socket_client()
}

func Execute_tcp_socket_commnicate() {
	Execute_tcp_socket_server()
	Execute_tcp_socket_client()
}
