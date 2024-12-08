package cortex

import (
	"errors"

	"golang.org/x/net/websocket"
)

func ConnectHeadset(ws *websocket.Conn, headsetName, clientId, clientSecret string) (string, string, error) {
	if err := Send(ws, GetDefaultInfoRequest()); err != nil {
		return "", "", err
	}
	var defaultInfo Response[DefaultInfoResponse]
	if err := Receive(ws, &defaultInfo); err != nil {
		return "", "", err
	}

	if err := Send(ws, GetAccessRequest(clientId, clientSecret)); err != nil {
		return "", "", err
	}
	var access Response[AccessResponse]
	if err := Receive(ws, &access); err != nil {
		return "", "", err
	}

	if err := Send(ws, GetHeadsetsRequest()); err != nil {
		return "", "", err
	}
	var headsets Response[HeadsetsResponse]
	if err := Receive(ws, &headsets); err != nil {
		return "", "", err
	}

	if err := Send(ws, GetAuthorizeRequest(clientId, clientSecret)); err != nil {
		return "", "", err
	}
	var authorize Response[AuthorizeResponse]
	if err := Receive(ws, &authorize); err != nil {
		return "", "", err
	}
	token := authorize.Result.CortexToken

	for _, headset := range headsets.Result {
		if headset.CustomName != headsetName {
			continue
		}
		if err := Send(ws, GetConnectHeadsetRequest(headset.ID)); err != nil {
			return "", "", err
		}
		var connectHeadset Response[ConnectHeadsetResponse]
		if err := Receive(ws, &connectHeadset); err != nil {
			return "", "", err
		}
		if isConnected(connectHeadset.Result) {
			return token, headset.ID, nil
		}
	}
	return token, "", errors.New(noHeadsetConnectedError)
}
