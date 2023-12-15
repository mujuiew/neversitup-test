package smiley

type Request struct {
	Body RequestBody `json:"rq_body"`
}

// RequestBody : body
type RequestBody struct {
	Text []string `json:"text" validate:"require"`
}

type Response struct {
	Body ResponseBody `json:"rs_body"`
}

// ResponseBody : body for response
type ResponseBody struct {
	Value int `json:"value"`
}
