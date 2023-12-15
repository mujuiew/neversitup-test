package shufflings

type Request struct {
	Body RequestBody `json:"rq_body"`
}

// RequestBody : body for shufflings
type RequestBody struct {
	Text string `json:"text" validate:"require"`
}

type Response struct {
	Body ResponseBody `json:"rs_body"`
}

// ResponseBody : body for response
type ResponseBody struct {
	Data []AnswerList `json:"data"`
}

// AnswerList : body for response
type AnswerList struct {
	Value string `json:"value"`
}
