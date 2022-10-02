package websocket_l

import (
	"context"
	"fmt"
	"net"
	"time"
)

/*
- Connecting to a UDP socket
In contrast to a TCP socket, with a UDP socket, the client just sends a datagram
to the server. There's no Accept function, since the server doesn't need to accept
a connection and just waits for datagrams to arrive.
*/
var addr_udp = net.UDPAddr{
	IP:   net.IPv4(127, 0, 0, 1),
	Port: 9001,
}

func udp_socket_server(ctx context.Context, conn *net.UDPConn) {
	select {
	case <-ctx.Done():
		conn.Close()
		break
	default:
		buffer := make([]byte, 512)
		fmt.Println("server Read from UDP")
		// receive message
		_, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			// handle error
			panic(fmt.Errorf("server side read error: %v", err))
		}
		fmt.Printf("server get message: %v\n", string(buffer[:]))

		go func() {
			fmt.Println("server Write to UDP")
			// send message
			// buf := make([]byte, 512)
			_, err = conn.WriteToUDP([]byte("ping"), addr)
			if err != nil {
				// handle error
				panic(fmt.Errorf("server side send error: %v", err))
			}
		}()
	}
}

func udp_socket_client(ctx context.Context) {
	select {
	case <-ctx.Done():
		break
	default:
		udpAddr_client, err := net.ResolveUDPAddr("udp", addr_udp.String())
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
		defer conn.Close()

		fmt.Println("udp client receive data")

		fmt.Println("udp client send data")
		// send message
		_, err = conn.Write([]byte("pong"))
		if err != nil {
			// handle error
			panic(fmt.Errorf("client send message error: %v", err))
		}

		// receive message
		buffer := make([]byte, 512)
		_, _, err = conn.ReadFrom(buffer)
		if err != nil {
			// handle error
			panic(fmt.Errorf("client receive message error: %v", err))
		}

		fmt.Printf("client get message: %v\n", string(buffer[:]))

	}
	// fmt.Printf("udp socket client: %v\n", n)
}

func Execute_udp_socket_server(ctx context.Context) {
	udpAddr_server, err := net.ResolveUDPAddr("udp", addr_udp.String())
	if err != nil {
		// handle error
		panic(fmt.Errorf("server side resolveUdpAddr error: %v", err))
	}
	fmt.Println("server listen UDP")
	listener, err := net.ListenUDP("udp", udpAddr_server)
	if err != nil {
		// handle error
		panic(fmt.Errorf("server side ListenUDP error: %v", err))
	}
	go udp_socket_server(ctx, listener)
}

func Execute_udp_socket_client(ctx context.Context) {
	go udp_socket_client(ctx)
}

func Execute_udp_socket_commnicate() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	Execute_udp_socket_server(ctx)
	Execute_udp_socket_client(ctx)
	time.Sleep(time.Second * 3)
}
