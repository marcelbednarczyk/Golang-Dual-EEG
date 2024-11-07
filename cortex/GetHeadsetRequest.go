package cortex

func GetHeadsetRequest() Request {
	return Request{
		ID:      2,
		JsonRPC: "2.0",
		Method:  "queryHeadsets",
		Params:  nil,
	}
}
