package cortex

func GetAccessRequest(clientId, clientSecret string) Request {
	return Request{
		ID:      1,
		JsonRPC: "2.0",
		Method:  "requestAccess",
		Params: AuthParams{
			ClientID:     clientId,
			ClientSecret: clientSecret,
		},
	}
}
