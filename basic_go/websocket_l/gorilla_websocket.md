# Gorilla WebSocket

## [source](https://pkg.go.dev/github.com/gorilla/websocket#Conn)

## Overview

### get websocket

 A `server application` calls the `Upgrader.Upgrade` method from an HTTP request handler to get a `*Conn`:

 ```go
 // specifies paarameters for upgrading an HTTP connection to WebSocket connection.
 var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
    // get websocket connection
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    ... Use conn to send and receive messages.
}
 ```

### WriteMessage and ReadMessage

 Call the connection's `WriteMessage` and `ReadMessage` methods to `send` and `receive` messages as a slice of bytes.

 ```go
 for {
    messageType, p, err := conn.ReadMessage()
    if err != nil {
        log.Println(err)
        return
    }

    if err := conn.WriteMessage(messageType, p); err != nil {
        log.Println(err)
        return
    }
 }
 ```

 In above snippet of code, `p` is a `[]byte` and `messageType` is an `int` with value `websocket.BinaryMessage` or `websocket.TextMessage`.

### send or receive message with io.WriteCloser and io.Reader

 An application `can also send` and `receive messages` using the `io.WriteCloser` and `io.Reader` interfaces.

 To `send` a message, call the connection `NextWriter method to get an io.WriteCloser`, write the message to the writer and close the writer when done.

 To `receive` a message, call the connection `NextReader method to get an io.Reader` and read until io.EOF is returned.

 ```go
 for {
    messageType, r, err := conn.NextReader()
    if err != nil {
        return
    }
    w, err := conn.NextWriter(messageType)
    if err != nil {
        return err
    }
    if _, err := io.Copy(w, r); err != nil {
        return err
    }
    if err := w.Close(); err != nil {
        return err
    }
}
 ```

### Data Messages

The WebSocket protocol distinguishes between text and binary data messages.

`Text messages` are interpreted as `UTF-8` encoded text. The interpretation of `binary messages is left to the application`.

This package uses the TextMessage and BinaryMessage `integer constants to identify the two data message types`.

### Control Messages

The WebSocket protocol defines three types of control messages: `close`, `ping` and `pong`. Call the connection `WriteControl`, `WriteMessage` or `NextWriter` methods to `send a control message` to the peer.

`Connections handle received close messages` by calling the handler function set with the `SetCloseHandler` method and by returning a *CloseError from the `NextReader`, `ReadMessage` or the message `Read` method. The default `close` handler sends a close message to the peer.

`Connections handle received ping messages` by calling the handler function set with the `SetPingHandler` method. The `default ping handler sends a pong message to the peer`.

`Connections handle received pong messages` by calling the handler function set with the `SetPongHandler` method. The `default pong handler does nothing`. If an application sends ping messages, then the application should set a pong handler to receive the corresponding pong.

The `control message handler` functions are called from the `NextReader`, `ReadMessage` and message reader `Read` methods. The `default close and ping handlers` can block these methods for a short time when the handler writes to the connection.

The application must read the connection to process close, ping and pong messages sent from the peer. If the application is `not otherwise interested in messages from the peer, then the application should start a goroutine to read and discard messages from the peer`.

```go
func readLoop(c *websocket.Conn) {
    for {
        if _, _, err := c.NextReader(); err != nil {
            c.Close()
            break
        }
    }
}
```

## Concurrency

Connections support one concurrent reader and one concurrent writer.

Applications are responsible for ensuring that `no more than one goroutine calls the write methods` (NextWriter, SetWriteDeadline, WriteMessage, WriteJSON, EnableWriteCompression, SetCompressionLevel) concurrently and that `no more than one goroutine calls the read methods` (NextReader, SetReadDeadline, ReadMessage, ReadJSON, SetPongHandler, SetPingHandler) concurrently.

The `Close` and `WriteControl` methods can `be called` concurrently `with all other methods`.

## Origin Considerations

Web browsers allow Javascript applications to open a WebSocket connection to any host. `It's up to the server to enforce an origin policy using the Origin request header sent by the browser`.

The `Upgrader` calls the function specified in the `CheckOrigin field to check the origin`. If the CheckOrigin function `returns false`, then the Upgrade method `fails` the WebSocket `handshake with HTTP status 403`.

If the `CheckOrigin field` is `nil`, then the Upgrader uses a `safe default`: fail the handshake if the Origin request header is present and the Origin host is not equal to the Host request header.

## Buffer

Connections buffer network input and output to `reduce the number of system calls` when reading or writing messages.

Write buffers are also used for `constructing WebSocket frames`. See `RFC 6455, Section 5` for a discussion of message framing.

A WebSocket `frame header is written` to the network `each time a write buffer is flushed` to the network. `Decreasing the size of the write buffer` can `increase the amount of framing overhead` on the connection.

The `buffer sizes in bytes are specified` by the ReadBufferSize and WriteBufferSize fields `in the Dialer and Upgrader`.

The `Dialer` uses a `default size of 4096 when a buffer size field is set to zero`.

The `Upgrader` reuses buffers created by the HTTP server when a buffer size field is set to zero. The HTTP server buffers have a size of 4096 at the time of this writing.
