package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ts httptest.Server

// Setup test http server
// and shut down the server and block after all request
// has gone through.
func init() {
	ts = *httptest.NewServer(router)
}

// Test method for the Look up user method
func TestLookUPUser(t *testing.T) {
	// Make a request to our server with the {base url}/ping
	resp, err := http.Get(fmt.Sprintf("%s/v1/user", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	//Test for 401 without bearerToken
	if resp.StatusCode != 401 {
		t.Fatalf("Expected status code 401, got %v", resp.StatusCode)
	}

	//Test for 401 with wrong bearerToken
	req := newRequest(ts.URL+"/v1/user", "badtoken")
	resp, err = http.DefaultClient.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 401 {
		t.Fatalf("Expected status code 401, got %v", resp.StatusCode)
	}

	//Test for 200 with user info with correct bearerToken
	req = newRequest(ts.URL+"/v1/user", "5230-SF20b-&21c1-%8vs1x41sd")
	resp, _ = http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}

// Helper method for making request with header
func newRequest(endpoint string, token string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	req.Header.Add("bearerToken", token)
	return req
}

//Test GetRecipe Method
func TestGetRecipe(t *testing.T) {
	// Test NO query
	resp, err := http.Get(fmt.Sprintf("%s/v1/user/recipe", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 404 {
		t.Fatalf("Expected status code 404, got %v", resp.StatusCode)
	}

	// test with ID in query param
	resp, err = http.Get(fmt.Sprintf("%s/v1/user/recipe?id=1000010", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	//test status code
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	staticBody := `{
		"title":"Boiled Goose",
		"uid":"100001-10",
		"date":"22-12-2022",
		"author":"Anders Wiggers",
		"basicInfo": {
			"prebTime":"10 mins",
			"cook":"25 mins",
			"difficulty":"Easy",
		},
		"coverImage":"http://linktoimg.com/10001.png",
		"ingredients":[
			{
				"Goose": {
					"amount":"2000kg",
					"purchasable":true
				}
			},
			{
				"Water": {
					"amount":"200ml",
					"purchasable":false
				}
			}
		],
		"Nutrition":{
			"kcal":"100",
			"fat":"12g",
			"saturates":"4g",
			"carbs":"100g",
			"sugars":"0g",
			"fibre":"3g",
			"protein":"20g",
			"salt":"0.5g"
		} ,
		"Method":[
			"warm water",
			"put goose in water",
			"take goose out of water"
		]
	}`
	//Test output of body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bodyStr := buf.String()
	bodyStr = bodyStr[1 : len(bodyStr)-1]
	fmt.Println(bodyStr)
	if bodyStr != staticBody {
		t.Fatalf("Got wrong body")
	}
}

func TestDeleteRecipe(t testing.T) {

}
