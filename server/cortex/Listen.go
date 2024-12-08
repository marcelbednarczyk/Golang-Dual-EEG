package cortex

import (
	"log/slog"
	"time"

	"golang.org/x/net/websocket"
)

func Listen(ws *websocket.Conn, token, headsetId, name string) error {
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

	// f, err := os.Create(name + ".csv")
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()
	// if _, err := f.WriteString(`"P1 alpha";"P1 low beta";"P1 high beta";"P1 gamma";"P1 theta";"P2 alpha";"P2 low beta";"P2 high beta";"P2 gamma";"P2 theta";"P3 alpha";"P3 low beta";"P3 high beta";"P3 gamma";"P3 theta";"P4 alpha";"P4 low beta";"P4 high beta";"P4 gamma";"P4 theta";"P5 alpha";"P5 low beta";"P5 high beta";"P5 gamma";"P5 theta"`); err != nil {
	// 	return err
	// }
	// if _, err := f.WriteString("\n"); err != nil {
	// 	return err
	// }

	sum, count := 0.0, 0
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			slog.Info("Average score", slog.Float64("score", sum/float64(count)))
			return nil
		default:
			var data SubscribeData
			if err := Receive(ws, &data); err != nil {
				return err
			}
			// for _, p := range data.Pow {
			// 	if _, err := f.WriteString(fmt.Sprintf("%f;", p)); err != nil {
			// 		return err
			// 	}
			// }
			// if _, err := f.WriteString("\n"); err != nil {
			// 	return err
			// }

			score, err := calculateScore(data.Pow)
			if err != nil {
				slog.Warn("Error calculating score", slog.String("error", err.Error()))
				continue
			}
			sum += score
			count++
		}
	}
}

// var x map[string]interface{}
// if err := Receive(ws, &x); err != nil {
// 	return err
// }
// jsonString, _ := json.Marshal(x)
// fmt.Println(string(jsonString))
