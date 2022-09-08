package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func routerPOSTReq(method, path, contentType string, jsonData []byte) (*httptest.ResponseRecorder, *http.Request) {
	w := httptest.NewRecorder()                                        // <1>
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(jsonData)) // <2>
	req.Header.Set("Content-Type", contentType)                        // <3>
	return w, req                                                      // <4>
}

func TestMeanHandler(t *testing.T) { // <1>
	tests := map[string]struct { // <2>
		method      string  // <3>
		path        string  // <4>
		contentType string  // <5>
		jsonData    []byte  // <6>
		wantResult  float64 // <7>
		wantCode    int     // <8>
	}{
		"ok request":                  {method: "POST", path: "/stat/mean", contentType: "application/json", jsonData: []byte(`{"data":[1,1,2,3,5]}`), wantResult: 2.4, wantCode: 200}, // <9>
		"bad request: Invalid Json":   {method: "POST", path: "/stat/mean", contentType: "application/json", jsonData: []byte(`{"data":[1,1,2,3,5}`), wantResult: 2.4, wantCode: 400},  // <10>
		"bad request PUT: code":       {method: "PUT", path: "/stat/mean", contentType: "application/json", jsonData: []byte(`{"data":[1,1,2,3,5]}`), wantResult: 2.4, wantCode: 405},  // <11>
		"bad request 2: Content-Type": {method: "POST", path: "/stat/mean", contentType: "text/csv", jsonData: []byte(`{"data":[1,1,2,3,5]}`), wantResult: 2.4, wantCode: 400},         // <12>
	}
	for name, tc := range tests { // <13>
		got, req := routerPOSTReq(tc.method, tc.path, tc.contentType, tc.jsonData) // <14>
		meanhandler(got, req)
		if tc.wantCode != got.Code { // <15>
			t.Fatalf("%s: expectedCode: %v, got: %v expectedResult: %v, got: %v", name, tc.wantCode, got.Code, tc.wantResult, got.Body.String()) // <16>
		}
		if tc.wantCode == 200 { // <17>
			outputDataMean := &OutputDataMean{}                     // <18>
			err := json.NewDecoder(got.Body).Decode(outputDataMean) // <19>
			if err != nil {                                         // <20>
				t.Error(err) // <21>
			}
		}

	}
}

func TestMedianHandler(t *testing.T) {
	tests := map[string]struct {
		method      string
		path        string
		contentType string
		jsonData    []byte
		wantResult  float64
		wantCode    int
	}{
		"ok request":                 {method: "POST", path: "/stat/median", contentType: "text/csv", jsonData: []byte(`1,1,2,3,5`), wantResult: 2, wantCode: 200},
		"ok request 2":               {method: "POST", path: "/stat/median", contentType: "text/csv", jsonData: []byte(`1,1,2,3,5,6`), wantResult: 2, wantCode: 200},
		"ok request 3":               {method: "POST", path: "/stat/median", contentType: "text/csv", jsonData: []byte(` 1`), wantResult: 2, wantCode: 200},
		"bad request: code":          {method: "PUT", path: "/stat/median", contentType: "text/csv", jsonData: []byte(`1,1,2,3,5,6`), wantResult: 2, wantCode: 405},
		"bad request 2: code":        {method: "POST", path: "/stat/median", contentType: "application/json", jsonData: []byte(`1,1,2,3,5,6`), wantResult: 2, wantCode: 400},
		"bad request 3: Invalid CSV": {method: "POST", path: "/stat/median", contentType: "text/csv", jsonData: []byte(``), wantResult: 2, wantCode: 400},
		"bad request 4: Invalid CSV": {method: "POST", path: "/stat/median", contentType: "text/csv", jsonData: []byte(`"1`), wantResult: 2, wantCode: 400},
		"bad request 5: Invalid CSV": {method: "POST", path: "/stat/median", contentType: "text/csv", jsonData: []byte(`ff`), wantResult: 2, wantCode: 400},
		"bad request 6: Invalid CSV": {method: "POST", path: "/stat/median", contentType: "text/csv", jsonData: []byte(` `), wantResult: 2, wantCode: 400},
	}
	for name, tc := range tests {
		got, req := routerPOSTReq(tc.method, tc.path, tc.contentType, tc.jsonData)
		medianhandler(got, req)
		if tc.wantCode != got.Code {
			t.Fatalf("%s: expectedCode: %v, got: %v expectedResult: %v, got: %v", name, tc.wantCode, got.Code, tc.wantResult, got.Body.String())
		}
		if tc.wantCode == 200 {
			outputDataMean := &OutputDataMean{}
			err := json.NewDecoder(got.Body).Decode(outputDataMean)
			if err != nil {
				t.Error(err)
			}
		}

	}
}
func TestEchoHandler(t *testing.T) {
	tests := map[string]struct {
		method      string
		path        string
		contentType string
		jsonData    []byte
		wantResult  float64
		wantCode    int
	}{
		"ok request": {method: "GET", path: "/=", contentType: "application/json", jsonData: []byte(``), wantResult: 2, wantCode: 200},
	}
	for name, tc := range tests {
		got, req := routerPOSTReq(tc.method, tc.path, tc.contentType, tc.jsonData)
		echohandler(got, req)
		if tc.wantCode != got.Code {
			t.Fatalf("%s: expectedCode: %v, got: %v expectedResult: %v, got: %v", name, tc.wantCode, got.Code, tc.wantResult, got.Body.String())
		}

	}
}
