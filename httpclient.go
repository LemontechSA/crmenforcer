package crmenforcer

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"time"
)

var client *http.Client
var url string

func init() {
	client = &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   2 * time.Second,
				KeepAlive: 2 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 2 * time.Second,
		},
	}
	if os.Getenv("ENFORCER_URL") != "" {
		url = os.Getenv("ENFORCER_URL")
	}
}

func exec(r rbac, s session) (status int, err error) {
	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(r)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Consumer-Id", s.ID)
	req.Header.Set("X-Consumer-Username", s.Username)
	req.Header.Set("token", s.Token)

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	status = res.StatusCode
	return
}
