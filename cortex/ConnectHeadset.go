package cortex

import "golang.org/x/net/websocket"

func ConnectHeadset(ws *websocket.Conn, clientId, clientSecret string) error {
	if err := Send(ws, GetDefaultInfoRequest()); err != nil {
		return err
	}

	if err := Send(ws, GetAccessRequest(clientId, clientSecret)); err != nil {
		return err
	}

	if err := Send(ws, GetHeadsetRequest()); err != nil {
		return err
	}

	var headsetResp ResponseSlice
	if err := websocket.JSON.Receive(ws, &headsetResp); err != nil {
		return err
	}

	headsets := make([]string, 0, len(headsetResp.Result))
	for _, headset := range headsetResp.Result {
		// TODO: change ResponseSlice to Response[T] where T is the type of Result in that case Headset
		_ = headset
	}
	_ = headsets

	return nil
}
