package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/OPPOGROUP/hoyolib/internal/errors"
	"github.com/OPPOGROUP/hoyolib/internal/log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type OptionFunc func(*Request) error

type Request struct {
	client   *http.Client
	url      *url.URL
	method   string
	params   map[string]string
	payloads map[string]interface{}
	cookies  map[string]string
	headers  map[string]string
}

func NewRequest(opts ...OptionFunc) (*Request, error) {
	r := &Request{
		client: &http.Client{
			Timeout: 2 * time.Second,
		},
	}
	err := r.Apply(opts...)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Request) Apply(opts ...OptionFunc) error {
	for _, o := range opts {
		err := o(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func WithUrl(u string) OptionFunc {
	return func(request *Request) error {
		var err error
		request.url, err = url.Parse(u)
		return err
	}
}

func WithMethod(method string) OptionFunc {
	return func(request *Request) error {
		if method != http.MethodPost && method != http.MethodGet {
			return errors.ErrHttpMethod
		}
		request.method = method
		return nil
	}
}

func WithParams(params map[string]string) OptionFunc {
	return func(request *Request) error {
		request.params = params
		return nil
	}
}

func WithPayloads(payloads map[string]interface{}) OptionFunc {
	return func(request *Request) error {
		request.payloads = payloads
		return nil
	}
}

func WithCookies(cookies map[string]string) OptionFunc {
	return func(request *Request) error {
		request.cookies = cookies
		request.client.Jar = request.transformCookies()
		return nil
	}
}

func WithHeaders(headers map[string]string) OptionFunc {
	return func(request *Request) error {
		request.headers = headers
		return nil
	}
}

func (r *Request) verify() bool {
	return true
}

func (r *Request) transformCookies() *cookiejar.Jar {
	cookies := make([]*http.Cookie, 0, len(r.cookies))
	for name, value := range r.cookies {
		cookies = append(cookies, &http.Cookie{
			Name:  name,
			Value: value,
		})
	}
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(r.url, cookies)
	return jar
}

func (r *Request) Do() (*http.Response, error) {
	var (
		pReader *bytes.Reader
	)
	if r.payloads != nil {
		payloads, _ := json.Marshal(r.payloads)
		pReader = bytes.NewReader(payloads)
	}
	req, err := http.NewRequest(r.method, r.url.String(), pReader)
	if err != nil {
		return nil, errors.ErrBuildRequest
	}
	if r.params != nil {
		q := req.URL.Query()
		for key, value := range r.params {
			q.Add(key, value)
		}
	}
	log.Debug().Msg(fmt.Sprintf("[http request] request = %v", req))
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, errors.ErrSendRequest
	}
	return resp, nil
}
