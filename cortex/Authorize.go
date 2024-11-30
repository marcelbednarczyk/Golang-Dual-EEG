package cortex

type AuthorizeResponse struct {
	CortexToken string `json:"cortexToken"`
}

func GetAuthorizeRequest(clientId, clientSecret string) Request {
	return Request{
		ID:      3,
		JsonRPC: "2.0",
		Method:  "authorize",
		Params: AuthParams{
			ClientID:     clientId,
			ClientSecret: clientSecret,
		},
	}
}
