package shufflings

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/neversitup-test/internal/core/shufflings"
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

	// unmarshal request to struct
	if err = json.Unmarshal(body, &req); err != nil {
		log.Print("ERROR Unmarshal: ", err)
		utility.WriteError(w, &respone)
		return
	}

	listString := shufflings.Process(prepareToCore(req.Body)) // เรียก core function เพื่อทำการคำนวณ

	respone = prepareToPesponse(listString)
}

func prepareToCore(req RequestBody) shufflings.BodyData {
	return shufflings.BodyData{
		Data: req.Text,
	}
}

func prepareToPesponse(resFromCore []string) Response {
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
