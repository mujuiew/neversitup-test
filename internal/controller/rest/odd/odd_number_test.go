package odd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	// mock data request
	request, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte(`{
		"rq_body" : {
			"number" : [1,2,3,4,5,6,1,2,3,4,5]
		}
	}`)))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	// ========================
	Handle(w, request)
	// ========================

	assert.Equal(t, 200, w.Result().StatusCode)
	// unmarshal data response for assert
	var respBody Response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &respBody))

	// assert len of response list
	assert.Equal(t, 1, len(respBody.Body.Data))

	// sort data of array before assert
	sort.SliceStable(respBody.Body.Data, func(i, j int) bool {
		return respBody.Body.Data[i].Value < respBody.Body.Data[j].Value
	})

	// {"rs_body":{"data":[{"value":6}]}}
	assert.Equal(t, 6, respBody.Body.Data[0].Value)

}

func TestHandle_case_zro(t *testing.T) {
	// mock data request
	request, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte(`{
		"rq_body" : {
			"number" : [1,2,3,4,5,6,1,2,3,4,5,6]
		}
	}`)))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	// ========================
	Handle(w, request)
	// ========================

	assert.Equal(t, 200, w.Result().StatusCode)
	// unmarshal data response for assert
	var respBody Response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &respBody))

	// assert len of response list
	assert.Equal(t, 1, len(respBody.Body.Data))

	// sort data of array before assert
	sort.SliceStable(respBody.Body.Data, func(i, j int) bool {
		return respBody.Body.Data[i].Value < respBody.Body.Data[j].Value
	})

	// {"rs_body":{"data":[{"value":6}]}}
	assert.Equal(t, 0, respBody.Body.Data[0].Value)

}
func TestHandle_case_error(t *testing.T) {
	// mock data request
	request, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte(`{
		"rq_body" : {
			"number" : [1,2,3,4,5,6,1,2,3,4,5
		}
	}`)))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	// ========================
	Handle(w, request)
	// ========================

	assert.Equal(t, 400, w.Result().StatusCode)

}
