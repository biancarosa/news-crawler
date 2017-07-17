package requests

import (
	"io/ioutil"
	"net/http"
)

func GetBody(url string) (*string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(body)
	return &stringBody, nil
}
