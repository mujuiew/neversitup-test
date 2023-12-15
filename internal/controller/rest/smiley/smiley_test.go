package smiley

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {

	// mock data request
	request, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte(`{
		"rq_body" : {
			"text" : [":)", ";(", ";}", ":-D"]
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

	// {"rs_body":{"data":[{"value":2}]}}
	assert.Equal(t, 2, respBody.Body.Value)

}

func TestHandle_case_null(t *testing.T) {

	// mock data request
	request, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte(`{
		"rq_body" : {
			"text" : [":1)",":-P"]
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

	// {"rs_body":{"data":[{"value":2}]}}
	assert.Equal(t, 0, respBody.Body.Value)

}

func TestHandle_error(t *testing.T) {

	// mock data request
	request, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte(`{
		"rq_body" : {
			"text" : [":)", ";(", ";}", ":-D]
		}
	}`)))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	// ========================
	Handle(w, request)
	// ========================

	assert.Equal(t, 400, w.Result().StatusCode)

}
