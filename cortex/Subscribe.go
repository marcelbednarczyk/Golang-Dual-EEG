package cortex

type SubscribeResponse struct {
	Success []SubscribeSuccess `json:"success"`
}

type SubscribeData struct {
	Pow       []float64 `json:"pow"`
	SessionID string    `json:"sid"`
	Time      float64   `json:"time"`
}

type SubscribeSuccess struct {
	StreamName string   `json:"streamName"`
	Cols       []string `json:"cols"`
	SessionID  string   `json:"sid"`
}

func GetSubscribeRequest(token, sessionId string) Request {
	return Request{
		ID:      1,
		JsonRPC: "2.0",
		Method:  "subscribe",
		Params: SubscribeParams{
			CortexToken: token,
			Session:     sessionId,
			Streams:     []string{"pow"},
		},
	}
}
