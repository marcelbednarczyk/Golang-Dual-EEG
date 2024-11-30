package cortex

type SessionResponse struct {
	ID        string          `json:"id"`
	Status    string          `json:"status"`
	Owner     string          `json:"owner"`
	License   string          `json:"license"`
	AppID     string          `json:"appId"`
	Started   string          `json:"started"`
	Stopped   string          `json:"stopped"`
	Streams   []string        `json:"streams"`
	RecordIDs []string        `json:"recordIds"`
	Recording bool            `json:"recording"`
	Headset   HeadsetResponse `json:"headset"`
}

func GetOpenSessionRequest(token, headsetID string) Request {
	return Request{
		ID:      6,
		JsonRPC: "2.0",
		Method:  "createSession",
		Params: SessionCreateParams{
			CortexToken: token,
			Status:      "open",
			Headset:     headsetID,
		},
	}
}

func GetCloseSessionRequest(token, sessionID string) Request {
	return Request{
		ID:      7,
		JsonRPC: "2.0",
		Method:  "updateSession",
		Params: SessionUpdateParams{
			CortexToken: token,
			Session:     sessionID,
			Status:      "close",
		},
	}
}
