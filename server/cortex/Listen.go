package cortex

import (
	"context"
	"log/slog"

	"golang.org/x/net/websocket"
)

func Listen(ctx context.Context, ws *websocket.Conn, token, headsetId, name string, ch chan<- float64) error {
	if err := Send(ws, GetOpenSessionRequest(token, headsetId)); err != nil {
		return err
	}
	var openSession Response[SessionResponse]
	if err := Receive(ws, &openSession); err != nil {
		return err
	}

	if err := Send(ws, GetSubscribeRequest(token, openSession.Result.ID)); err != nil {
		return err
	}
	var subscribe Response[SubscribeResponse]
	if err := Receive(ws, &subscribe); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			var data SubscribeData
			if err := Receive(ws, &data); err != nil {
				return err
			}

			score, err := calculateScore(data.Pow)
			if err != nil {
				slog.Warn("Error calculating score", slog.String("error", err.Error()))
				continue
			}
			ch <- score
		}
	}
}
