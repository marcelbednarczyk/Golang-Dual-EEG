package cortex

type AuthParams struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	License      string `json:"license,omitempty"`
	Debit        int    `json:"debit,omitempty"`
}

type AuthResult struct {
	CortexToken string      `json:"cortexToken"`
	Warning     interface{} `json:"warning"`
}

type DataSample struct {
	Com  []interface{} `json:"com"`
	SID  string        `json:"sid"`
	Time float32       `json:"time"`
}

type GetProfileParams struct {
	CortexToken string `json:"cortexToken"`
	Headset     string `json:"headset"`
}

type Request struct {
	ID      int         `json:"id"`
	JsonRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type Response[T any] struct {
	ID      int         `json:"id"`
	JsonRPC string      `json:"jsonrpc"`
	Result  T           `json:"result"`
	Error   interface{} `json:"error"`
}

type SetupProfileParams struct {
	CortexToken string `json:"cortexToken"`
	Headset     string `json:"headset"`
	Profile     string `json:"profile"`
	Status      string `json:"status"`
}

type SessionCreateParams struct {
	CortexToken string `json:"cortexToken"`
	Status      string `json:"status"`
	Headset     string `json:"headset"`
}

type SessionUpdateParams struct {
	CortexToken string `json:"cortexToken"`
	Session     string `json:"session"`
	Status      string `json:"status"`
}

type SubscribeParams struct {
	CortexToken string   `json:"cortexToken"`
	Session     string   `json:"session"`
	Streams     []string `json:"streams"`
}

type TrainingParams struct {
	CortexToken string `json:"cortexToken"`
	Session     string `json:"session"`
	Detection   string `json:"detection"`
	Status      string `json:"status"`
	Action      string `json:"action"`
}
