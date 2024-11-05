package cortex

func GetDefaultInfoRequest() Request {
	return Request{
		ID:      0,
		JsonRPC: "2.0",
		Method:  "getCortexInfo",
		Params:  nil,
	}
}
