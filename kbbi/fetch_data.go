package kbbi

import (
	"encoding/json"
	"errors"
	"github.com/dimassfeb-09/kbbi.git/entity"
	"io"
	"net/http"
)

func Search(keyword string) (*entity.ResponseKBBI, error) {

	var responseKBBI *entity.ResponseKBBI

	response, err := http.Get("https://kbbi-api-zhirrr.vercel.app/api/kbbi?text=" + keyword)
	if err != nil {
		return nil, err
	} else {
		if response.StatusCode == 200 {
			bytesResponse, err := io.ReadAll(response.Body)
			if err != nil {
				return nil, err
			} else {
				err := json.Unmarshal(bytesResponse, &responseKBBI)
				if err != nil {
					return nil, err
				}
			}
		} else if response.StatusCode == http.StatusGatewayTimeout {
			return nil, errors.New(keyword + " not found.")
		}
	}

	return responseKBBI, nil
}
