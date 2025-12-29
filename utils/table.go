package utils

import (
	"net/url"
	"strings"
)

func GetURLParams(link string) (map[string]string, error) {
	paramsMap := make(map[string]string)

	link, err := url.QueryUnescape(link)

	if err != nil {
		return nil, err
	}

	paramString := strings.Split(link, "?")[1]
	params := strings.SplitSeq(paramString, "&")
	for param := range params {
		keyValPair := strings.Split(param, "=")
		paramsMap[keyValPair[0]] = keyValPair[1]
	}

	return paramsMap, nil
}
