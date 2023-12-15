package shufflings

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
			"text" : "abc"
		}
	}`)))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	// ========================
	Handle(w, request)
	// ========================

	// unmarshal data response for assert
	var respBody Response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &respBody))

	// assert len of response list
	assert.Equal(t, 6, len(respBody.Body.Data))

	// sort data of array before assert
	sort.SliceStable(respBody.Body.Data, func(i, j int) bool {
		return respBody.Body.Data[i].Value < respBody.Body.Data[j].Value
	})

	// {"rs_body":{"data":[{"value":"abc"},{"value":"acb"},{"value":"bac"},{"value":"bca"},{"value":"cba"},{"value":"cab"}]}}
	assert.Equal(t, "abc", respBody.Body.Data[0].Value)
	assert.Equal(t, "acb", respBody.Body.Data[1].Value)
	assert.Equal(t, "bac", respBody.Body.Data[2].Value)
	assert.Equal(t, "bca", respBody.Body.Data[3].Value)
	assert.Equal(t, "cab", respBody.Body.Data[4].Value)
	assert.Equal(t, "cba", respBody.Body.Data[5].Value)

}
