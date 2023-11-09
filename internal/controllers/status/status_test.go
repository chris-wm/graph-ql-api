package status

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/electivetechnology/utility-library-go/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

type TestStatusItem struct {
	Method         string
	Path           string
	ExpectedStatus int
	ExpectedBody   gin.H
}

func TestStatus(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"status": http.StatusOK,
	}

	s1 := TestStatusItem{
		Method:         "GET",
		Path:           "/v2/status",
		ExpectedStatus: http.StatusOK,
		ExpectedBody:   body,
	}

	s2 := TestStatusItem{
		Method:         "HEAD",
		Path:           "/v2/status",
		ExpectedStatus: http.StatusNoContent,
		ExpectedBody:   nil,
	}

	testItmes := []TestStatusItem{
		s1,
		s2,
	}

	// Grab our router
	router := router.NewRouter()

	// Perform a GET request with that handler.
	for _, item := range testItmes {
		response := performRequest(router.Engine, item.Method, item.Path)

		// Assert we encoded correctly,
		// the request gives a 200
		assert.Equal(t, item.ExpectedStatus, response.Code)

		// Convert the JSON response to a map
		if item.ExpectedBody != nil {

			responseBody := StatusResponse{}
			err := json.Unmarshal([]byte(response.Body.String()), &responseBody)

			// Grab the value & whether or not it exists
			value := responseBody.Status

			// Make some assertions on the correctness of the response.
			assert.Nil(t, err)
			assert.Equal(t, body["status"], value)
		}
	}
}
