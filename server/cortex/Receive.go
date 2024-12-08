package cortex

import "golang.org/x/net/websocket"

func Receive[T any](ws *websocket.Conn, resp T) error {
	if err := websocket.JSON.Receive(ws, resp); err != nil {
		return err
	}
	return nil
}
