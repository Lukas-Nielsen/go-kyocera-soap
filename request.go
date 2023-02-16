package kyocera_soap

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func (c *Config) reuqest(path string, body string, result any) bool {
	var protocol string
	if c.port == 9091 {
		protocol = "https"
	} else {
		protocol = "http"
	}

	client := resty.New().
		SetBaseURL(protocol + "://" + c.host + ":" + fmt.Sprint(c.port)).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetDebug(false)

	resp, err := client.R().
		SetHeader("Content-Type", "application/soap+xml; charset=utf-8").
		SetBody(body).
		SetResult(&result).
		Post(path)
	if err != nil {
		log.Println(err)
		return false
	}
	if resp.IsError() {
		log.Println(resp.Error())
		return false
	}
	if c.debug {
		log.Printf("%+v", result)
	}
	return true
}
