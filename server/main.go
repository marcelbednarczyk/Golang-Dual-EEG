package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/marcelbednarczyk/Golang-Dual-EEG/cortex"
	"golang.org/x/net/websocket"
)

const (
	numberOfHeadsets = 2
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

	websockets := make([]*websocket.Conn, numberOfHeadsets)
	for i := 0; i < numberOfHeadsets; i++ {
		ws, err := websocket.DialConfig(config)
		if err != nil {
			slog.Error("Error dialing websocket", slog.String("error", err.Error()))
			return
		}
		websockets[i] = ws
	}

	scoreChans := make([]chan float64, numberOfHeadsets)
	for i := 0; i < numberOfHeadsets; i++ {
		scoreChans[i] = make(chan float64)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	clientId, clientSecret := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET")
	for i := 0; i < numberOfHeadsets; i++ {
		headsetName := os.Getenv("HEADSET_NAME_" + strconv.Itoa(i+1))
		go func(headsetName, clientId, clientSecret string) {
			token, headsetId, err := cortex.ConnectHeadset(websockets[i], headsetName, clientId, clientSecret)
			if err != nil {
				slog.Error("Error connecting headset", slog.String("headsetName", headsetName), slog.String("error", err.Error()))
				return
			}
			slog.Info("Connected to headset", slog.String("headsetName", headsetName), slog.String("headsetId", headsetId))

			if err := cortex.Listen(ctx, websockets[i], token, headsetId, headsetName, scoreChans[i]); err != nil {
				slog.Error("Error subscribing to headset", slog.String("headsetName", headsetName), slog.String("error", err.Error()))
				return
			}
			slog.Info("Stopped", slog.String("headsetName", headsetName), slog.String("headsetId", headsetId))
		}(headsetName, clientId, clientSecret)
	selectLoop:
		for {
			select {
			case score := <-scoreChans[i]:
				slog.Info("Score", slog.Float64("score", score))
				break selectLoop
			default:
				continue
			}
		}
	}

	value := 50 // %
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			score1, score2 := tryReceive(scoreChans[0]), tryReceive(scoreChans[1])
			if score1 == 0 && score2 == 0 {
				continue
			}
			if score1 > score2 {
				slog.Info("Score 1 is higher", slog.Float64("score1", score1), slog.Float64("score2", score2))
				value++
			} else {
				slog.Info("Score 2 is higher", slog.Float64("score1", score1), slog.Float64("score2", score2))
				value--
			}
		}
	}()

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		switch r.URL.Path {
		case "/api/dual-eeg/v1.0/score":
			switch r.Method {
			case http.MethodGet:
				score := ScoreResponse{Score: value}
				resp, err := json.Marshal(score)
				if err != nil {
					http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(resp)
			case http.MethodOptions:
				w.WriteHeader(http.StatusOK)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		default:
			http.NotFound(w, r)
		}
	}))
}

type ScoreResponse struct {
	Score int `json:"score"`
}

func tryReceive(ch <-chan float64) float64 {
	select {
	case score := <-ch:
		return score
	default:
		return 0
	}
}
