package socket

import (
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func CreateSocket() *socketio_client.Client {
	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	uri := "http://127.0.0.1:3000/"
	client, _ := socketio_client.NewClient(uri, opts)
	return client
}
