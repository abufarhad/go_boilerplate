package client

import (
	"bytes"
	"core/infra/config"
	"core/infra/logger"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

type serviceClient struct {
	Client *http.Client
}

var sClient serviceClient

func initiateHttpClient() *serviceClient {
	timeout := config.InternalService().Timeout * time.Second
	var netTransport = &http.Transport{
		DialContext:         (&net.Dialer{Timeout: timeout, KeepAlive: time.Minute}).DialContext,
		TLSHandshakeTimeout: timeout,
		MaxIdleConnsPerHost: 10,
	}

	sClient = serviceClient{
		Client: &http.Client{
			Timeout:   timeout,
			Transport: netTransport,
		},
	}

	return &sClient
}

func prepareFileUploadRequest(bodyBuffer *bytes.Buffer, contentType string, reqURL *url.URL) (http.Request, error) {

	req := http.Request{
		Method: "POST",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type": {contentType},
		},
		Body: ioutil.NopCloser(bodyBuffer),
	}

	logger.InfoAsJson(fmt.Sprintf("prepared request for %v", reqURL))
	return req, nil
}

func prepareGetRequest(headers map[string][]string, method string, reqURL *url.URL) http.Request {

	if headers == nil {
		headers = make(map[string][]string)
	}
	headers["Content-Type"] = []string{"application/json"}

	req := http.Request{
		Method: method,
		URL:    reqURL,
		Header: headers,
	}

	logger.InfoAsJson(fmt.Sprintf("prepared request for %v", reqURL))
	return req
}

func prepareRequest(body []byte, method string, reqURL *url.URL) http.Request {

	req := http.Request{
		Method: method,
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body:          ioutil.NopCloser(bytes.NewReader(body)),
		ContentLength: int64(len(body)),
	}

	logger.InfoAsJson(fmt.Sprintf("prepared request for %v", reqURL))
	return req
}
