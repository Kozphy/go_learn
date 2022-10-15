package main

import (
	"time"

	pingpong "github.com/zixas/go_learn/websocket_l/ping_pong"
)

// import "github.com/zixas/go_learn/concurrency"

// generic "github.com/zixas/go_learn/Generics"
// "github.com/zixas/go_learn/crawler"

// _interface "github.com/zixas/go_learn/interface"
// json "github.com/zixas/go_learn/json_and_go"

// "github.com/zixas/go_learn/concurrency/server"

// "github.com/zixas/go_learn/methods"
// "github.com/zixas/go_learn/pointer"

func main() {
	// reflection.Law_reflect()
	// concurrency.Final_state()
	// server.Execute_server()
	// concurrency.Execute_timeout_digitoc()
	// methods.Execute_interface_digitoc()
	// methods.Execute_Pointer_receiver()
	// methods.Execute_pointer_receivers_interface()
	// pointer.Execute_define_using_pointer()
	// pointer.Execute_function_pointer_receivers()
	// pointer.Execute_nil_pointers()
	// pointer.Execute_nil_pointer_improve()
	// pointer.Execute_method_pointer_receivers()
	// reflection.Execute_reflect_with_name_return()
	// websocket_l.Execute_tcp_socket_commnicate()
	// websocket_l.Execute_udp_socket_commnicate()
	// websocket_l.Execute_simeple_http_endpoint()
	// chat_tutorial.Execute_chat_websocket()
	go pingpong.Execute_ping_pong_server()
	time.Sleep(time.Second)
	pingpong.Execute_ping_pong_client()
}
