package cortex

import "golang.org/x/net/websocket"

func Send(ws *websocket.Conn, request Request) error {
	return websocket.JSON.Send(ws, request)
}
