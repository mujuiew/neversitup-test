package odd

type Request struct {
	Body RequestBody `json:"rq_body"`
}

// RequestBody : body
type RequestBody struct {
	Number []int `json:"number" validate:"require"`
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
	Value int `json:"value"`
}
