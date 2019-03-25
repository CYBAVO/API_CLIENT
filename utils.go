// Copyright (c) 2018-2019 The Cybavo developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of Cybavo and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to Cybavo
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from Cybavo.

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

func makeRequest(method string, api string, params []string, postBody *string) ([]byte, error) {
	if method == "" || api == "" {
		return nil, errors.New("invalid parameters")
	}

	client := &http.Client{}
	t := time.Now().Unix()
	url := fmt.Sprintf("%s%s?ac=%s&t=%d", baseURL, api, myAPICode, t)

	var req *http.Request
	var err error
	if postBody == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest("POST", url, strings.NewReader(*postBody))
		params = append(params, *postBody)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set(sigHeader, buildSign(params, mySecretKey, t))
	if postBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
