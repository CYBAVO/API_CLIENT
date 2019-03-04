package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

func sortParams(params []string) string {
	if len(params) > 0 {
		sort.Strings(params)
		return strings.Join(params, "&")
	}
	return ""
}

func buildSign(params []string, secretKey string, time int64) string {
	params = append(params, fmt.Sprintf("t=%d", time))
	sorted := sortParams(params)
	fullParams := fmt.Sprintf("%s&secret_key=%s", sorted, secretKey)
	sha := sha256.Sum256([]byte(fullParams))
	return fmt.Sprintf("%x", sha)
}

func makeRequest(method string, api string, params []string) (string, error) {
	if method == "" || api == "" {
		return "", errors.New("invalid parameters")
	}

	client := &http.Client{}
	t := time.Now().Unix()
	url := fmt.Sprintf("%s%s?t=%d", baseURL, api, t)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set(apiHeader, myAPICode)
	req.Header.Set(sigHeader, buildSign(params, mySecretKey, t))

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
