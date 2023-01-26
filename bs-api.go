package bsapi

import (
	"bytes"
	"context"
	"net/http"
	jsoniter "github.com/json-iterator/go"
)

const apiUrl = `https://api.brawlstars.com/v1`

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type BsApi struct {
	token string
}

func (api BsApi) makeRequest(ctx context.Context, urlSuffix string) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", apiUrl+urlSuffix, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+api.token)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf := bytes.Buffer{}
	buf.ReadFrom(resp.Body)
	data := buf.Bytes()
	if resp.StatusCode == http.StatusOK {
		return data, nil
	}
	ce := ClientError{}
	err = json.Unmarshal(data, &ce)
	if err != nil {
		return nil, err
	}
	return nil, ce
}

func NewApi(token string) BsApi {
	bs := BsApi{}
	bs.token = token
	return bs
}
