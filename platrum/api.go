package platrum

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func buildUrl(method, query string) string {
	u := url.URL{
		Scheme:   "http",
		Host:     fmt.Sprintf("%s:%s", coreHost(), corePort()),
		Path:     "/core-api/" + method,
		RawQuery: query,
	}

	return u.String()
}

func requestBool(method string, query url.Values) (bool, error) {
	response, err := sendPostRequest(method, query.Encode(), []interface{}{})
	if err != nil {
		return false, err
	}

	var res bool
	if err = json.Unmarshal(response, &res); err != nil {
		return false, err
	}

	return res, nil
}

func sendPostRequest(method, query string, body interface{}) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	u := buildUrl(method, query)

	resp, err := http.Post(u, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	isSuccess, ok := res["success"]
	if !ok || !isSuccess.(bool) {
		return nil, fmt.Errorf(res["error"].(string))
	}

	return json.Marshal(res["result"])
}
