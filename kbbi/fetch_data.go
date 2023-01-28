package kbbi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ResponseKBBI struct {
	Lema string   `json:"lema"`
	Arti []string `json:"arti"`
}

func Search(keyword string) (*ResponseKBBI, error) {

	var responseKBBI *ResponseKBBI

	response, err := http.Get("https://kbbi-api-zhirrr.vercel.app/api/kbbi?text=" + keyword)
	if err != nil {
		log.Print(err)
	} else {
		bytesResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		} else {
			err := json.Unmarshal(bytesResponse, &responseKBBI)
			if err != nil {
				return nil, err
			}
		}
	}
	return responseKBBI, nil
}
