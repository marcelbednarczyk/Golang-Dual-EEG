package main

import (
	"crypto/tls"
	"log/slog"
	"os"

	"github.com/marcelbednarczyk/Golang-Dual-EEG/cortex"
	"golang.org/x/net/websocket"
)

func main() {
	slog.Info("Hello, World!")

	origin := os.Getenv("ORIGIN")
	url := "wss://" + os.Getenv("WS_IP") + ":6868"
	config, err := websocket.NewConfig(url, origin)
	if err != nil {
		return
	}

	if os.Getenv("SKIP_TLS") == "true" {
		config.TlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	ws, err := websocket.DialConfig(config)
	if err != nil {
		slog.Error("Error dialing websocket", slog.String("error", err.Error()))
		return
	}

	token, headsetId, err := cortex.ConnectHeadset(ws, os.Getenv("HEADSET_NAME"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		slog.Error("Error connecting headset", slog.String("error", err.Error()))
		return
	}
	slog.Info("Connected to headset", slog.String("headset", headsetId))

	if err := cortex.Listen(ws, token, headsetId, os.Getenv("HEADSET_NAME")); err != nil {
		slog.Error("Error subscribing to headset", slog.String("error", err.Error()))
		return
	}
	slog.Info("Stopped")
}
