package websocket_l

// source: https://yalantis.com/blog/how-to-build-websockets-in-go/

/*
	What is meant by record or data boundaries in the sense of TCP & UDP protocol?
	https://stackoverflow.com/questions/51661519/what-is-meant-by-record-or-data-boundaries-in-the-sense-of-tcp-udp-protocol
*/

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

// What WebSockets are
/*
	The WebSocket communication package provides a full-duplex communication channel
	over a single TCP connection. That means that both the client and the server can
	simultaneously send data whenever they need without any requrest.

	WebSockets are a good solution for services that require continuous data exchange -
	for instance, instant messages, online games, and real-time trading systems. you see
	IETF RFC 6455 specification(https://www.rfc-editor.org/rfc/rfc6455).

	WebSocket connections are requested by browsers and are responded to by servers,
	after which a connection is established. This process is often called a "handshake".

	The special kind of header in WebSockets requires only one handshake(https://en.wikipedia.org/wiki/Handshaking)
	between a browser andserver for establishing a connection that will
	remain active throughout its lifetime.

	The Websocket protocol uses port 80 for an unsecure connection and port 443 for a
	secure connection.

	The WebSocket specification determines which uniform resource identifier schemes
	are required for the wd (WebSocket) and wss (WebSocket Secure) protocols.

	WebSockets have serveral benefits over traditional HTTP:
	- The lightweight header reduces data transmission overhead.
	- Only one TCP connection is required for a single web client.
	- WebSocket servers can push data to web clients.

	The WebSocket protocol uses HTTP protocol for the initial handshake.
	After a successful handshake, a connection is established and the
	WebSocket essentially uses raw TCP to read/write data.
*/
