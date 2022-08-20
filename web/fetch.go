package web

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lucas-clemente/quic-go/http3"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

func FetchJson[R any](mode Mode, req *http.Request, res *R, callback ...func(response *http.Response)) error {
	data, err := FetchData(mode, req, callback...)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data.Bytes(), res); err != nil {
		return err
	}
	return nil
}

func FetchData(mode Mode, req *http.Request, callback ...func(response *http.Response)) (*bytes.Buffer, error) {
	response, err := Fetch(req, mode)
	if err != nil {
		return nil, err
	}
	if callback != nil && len(callback) > 0 {
		callback[0](response)
	}
	body := &bytes.Buffer{}
	if _, err := io.Copy(body, response.Body); err != nil {
		return nil, err
	}
	return body, nil
}

func Fetch(req *http.Request, mode Mode) (*http.Response, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	switch mode {
	case ModeNormal:
	case ModeWeb3:
		client.Transport = &http3.RoundTripper{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	default:
		return nil, errors.New("unknown mode")
	}
	resp, reqErr := client.Do(req)
	if reqErr != nil {
		return nil, reqErr
	}
	return resp, nil
}

func UrlEncoded(v any) string {
	jsonStr, _ := json.Marshal(v)
	var data map[string]any
	_ = json.Unmarshal(jsonStr, &data)
	var build strings.Builder
	j := 0
	keys := make([]string, len(data))
	for d := range data {
		keys[j] = d
		j++
	}
	sort.Strings(keys)
	for _, s := range keys {
		val, ok := data[s]
		if !ok || val == nil {
			continue
		}
		valStr, flag := val.(string)
		if flag && len(valStr) == 0 {
			continue
		}
		if build.Len() > 0 {
			build.WriteString("&")
		}
		build.WriteString(s + "=")
		build.WriteString(fmt.Sprintf("%v", val))
	}
	return build.String()
}
