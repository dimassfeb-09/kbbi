package kbbi

import (
	"encoding/json"
	"github.com/dimassfeb-09/kbbi.git/entity"
	"io"
	"log"
	"net/http"
)

func Search(keyword string) (*entity.ResponseKBBI, error) {

	var responseKBBI *entity.ResponseKBBI

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
