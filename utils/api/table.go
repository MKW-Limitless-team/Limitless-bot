package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/MKW-Limitless-team/limitless-types/table"
)

func GetTable(data string) (*table.Table, error) {
	url := fmt.Sprintf("http://localhost:8080/table?data=%s", url.QueryEscape(data))
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	table := &table.Table{}
	err = json.NewDecoder(response.Body).Decode(table)
	if err != nil {
		return nil, err
	}

	return table, err
}
