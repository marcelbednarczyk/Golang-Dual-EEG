package cortex

import (
	"golang.org/x/net/websocket"
)

func ConnectHeadset(ws *websocket.Conn, clientId, clientSecret string) error {
	if err := Send(ws, GetDefaultInfoRequest()); err != nil {
		return err
	}
	var defaultInfo Response[DefaultInfoResponse]
	if err := websocket.JSON.Receive(ws, &defaultInfo); err != nil {
		return err
	}
	_ = defaultInfo

	if err := Send(ws, GetAccessRequest(clientId, clientSecret)); err != nil {
		return err
	}
	var access Response[AccessResponse]
	if err := websocket.JSON.Receive(ws, &access); err != nil {
		return err
	}
	_ = access

	if err := Send(ws, GetHeadsetsRequest()); err != nil {
		return err
	}
	var headsets Response[HeadsetsResponse]
	if err := websocket.JSON.Receive(ws, &headsets); err != nil {
		return err
	}

	for _, headset := range headsets.Result {
		_ = headset
	}

	return nil
}

// var x map[string]interface{}
// if err := websocket.JSON.Receive(ws, &x); err != nil {
// 	return err
// }
// jsonString, _ := json.Marshal(x)
// fmt.Println(string(jsonString))
