package websocket_l

// source: https://yalantis.com/blog/how-to-build-websockets-in-go/

// Network sockets
/*
	A network socket, or simply a socket, serves as an internal endpoint for exchanging data between
	applications running on the same computer or on different computers on the same network.

	Sockets are a key part of Unix and Windows-based os, and they make it easier for developers
	to create network-enabled software. Instead of constructing network from scratch, app
	developers can include sockets in their programs. Since network sockets are used for several network
	protocols (HTTP, FTP, etc.), multiple sockets can be used simultaneously.

	Sockets are created and used with a set of function calls defined by a socket's application
	programming interface (API).

	There are serval types of network socket:
	- Datagram sockets (SOCK_DGRAM), also knwon as connectionless sockets, use the User
	Datagram Protocol (UDP). Datagram sockets supports a bidirectional flow of messages
	and preserve record boundaries.

	- Stream sockets (SOCK_STREAM), also known as connection-oriented sockets, use the
	Transmission Control Protocol (TCP), Stream Control Transmission Protocol (SCTP), or
	Datagram Congestion Control Protocol (DCCP). These sockets provide a bidirectional,
	reliable, sequenced, and unduplicated flow of data with no record boundaries.

	- Raw sockets: (or raw IP sockets) are typically available in routers and other networking
	equipment. These sockets are normally datagram-oriented, although their exact characteristics
	depend on the interface provided by the protocol. Raw sockets are not used by most applications.
	They're provided to support the development of new communication protocols and to provide
	access to more esoteric facilities of existing protocols.
*/

// Socket communication
/*
	Each network socket is identified by the address, which is a triad of a
	transport protocol, Ip address, and port number.

	There are two major protocols for communicating between hosts: TCP and UDP.

*/
