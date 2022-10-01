package websocket_l

import (
	"fmt"
	"net"
)

/*
- Connecting to a UDP socket
In contrast to a TCP socket, with a UDP socket, the client just sends a datagram
to the server. There's no Accept function, since the server doesn't need to accept
a connection and just waits for datagrams to arrive.
*/

func udp_socket_server(conn *net.UDPConn) {
	// send message
	buffer := make([]byte, 512)
	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		// handle error
		panic(fmt.Errorf("server side send error: %v", err))
	}

	// receive message
	buf := make([]byte, 512)
	n, err = conn.WriteToUDP(buf[:n], addr)
	if err != nil {
		// handle error
		panic(fmt.Errorf("server side receive error: %v", err))
	}

	fmt.Printf("server get message: %v\n", string(buf[:]))
	fmt.Printf("tcp socket server: %v\n", n)
}

func udp_socket_client() {
	udpAddr_client, err := net.ResolveUDPAddr("tcp", addr.String())
	if err != nil {
		// handle error
		panic(fmt.Errorf("client side resolveUdpAddr error: %v", err))
	}
	fmt.Println("client dialUDP")
	conn, err := net.DialUDP("udp", nil, udpAddr_client)
	if err != nil {
		// handle error
		panic(fmt.Errorf("DialUDP error: %v", err))
	}

	// send message
	buffer := make([]byte, 512)
	n, addr, err := conn.ReadFrom(buffer)
	if err != nil {
		// handle error
		panic(fmt.Errorf("client send message error: %v", err))
	}

	// receive message
	var buf [512]byte
	n, err = conn.WriteTo(buf[:n], addr)
	if err != nil {
		// handle error
		panic(fmt.Errorf("client receive message error: %v", err))
	}
	fmt.Printf("client get message: %v\n", string(buf[:]))
	fmt.Printf("udp socket client: %v\n", n)

}

func Execute_udp_socket_server() {
	udpAddr_server, err := net.ResolveUDPAddr("udp", addr.String())
	if err != nil {
		// handle error
		panic(fmt.Errorf("server side resolveTcpAddr error: %v", err))
	}

	listener, err := net.ListenUDP("udp", udpAddr_server)
	if err != nil {
		// handle error
		panic(fmt.Errorf("server side ListenUDP error: %v", err))
	}
	go udp_socket_server(listener)
}

func Execute_udp_socket_client() {
	go udp_socket_client()
}

func Execute_udp_socket_commnicate() {
	Execute_udp_socket_server()
	Execute_udp_socket_client()
}
