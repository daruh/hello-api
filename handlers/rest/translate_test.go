package rest_test

import (
	"encoding/json"
	"hello-api/handlers/rest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTranslateAPI(t *testing.T) {
	tt := []struct { // <1>
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/hello",
			StatusCode:          200,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:            "/hello?language=german",
			StatusCode:          200,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
	}

	handler := http.HandlerFunc(rest.TranslateHandler)

	for _, test := range tt { // <3>
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)

		handler.ServeHTTP(rr, req)

		if rr.Code != test.StatusCode {
			t.Errorf(`expected status %d but received %d`,
				test.StatusCode, rr.Code)
		}

		var resp rest.Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.Language != test.ExpectedLanguage {
			t.Errorf(`expected language "%s" but received %s`,
				test.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != test.ExpectedTranslation {
			t.Errorf(`expected Translation "%s" but received "%s"`,
				test.ExpectedTranslation, resp.Translation)
		}

	}
}
