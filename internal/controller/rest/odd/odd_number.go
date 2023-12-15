package odd

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/neversitup-test/internal/core/odd"
	"github.com/neversitup-test/internal/core/utility"
)

func Handle(w http.ResponseWriter, r *http.Request) {

	var req Request
	var respone Response

	defer utility.WriteBody(w, &respone) // เขียน response

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("ERROR ReadAll: ", err)
		utility.WriteError(w, &respone)
		return
	}
	if err = r.Body.Close(); err != nil {
		log.Print("ERROR: ", err)
		utility.WriteError(w, &respone)
		return
	}

	err = json.Unmarshal(body, &req) // unmarshal request to struct
	if err != nil {
		log.Print("ERROR Unmarshal: ", err)
		utility.WriteError(w, &respone)
		return
	}

	listString := odd.Process(req.Body.Number) // เรียก core function เพื่อทำการคำนวณ

	respone = prepareToPesponse(listString)
}

func prepareToPesponse(resFromCore []int) Response {
	var dataResList []AnswerList
	for _, res := range resFromCore {
		dataResList = append(dataResList, AnswerList{
			Value: res,
		})
	}

	return Response{
		Body: ResponseBody{
			Data: dataResList,
		},
	}
}
