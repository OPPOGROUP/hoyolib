package request

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httputil"
	"net/url"
	"testing"
)

func TestHttpGet(t *testing.T) {
	u, _ := url.Parse("https://translate.google.com/")
	q := u.Query()
	q.Add("sl", "en")
	q.Add("tl", "zh-CN")
	q.Add("text", "start")
	q.Add("op", "translate")
	u.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	assert.Nil(t, err)
	dumpReq, err := httputil.DumpRequestOut(req, true)
	assert.Nil(t, err)
	t.Log(string(dumpReq))
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	dumpResp, err := httputil.DumpResponse(resp, false)
	assert.Nil(t, err)
	t.Log(string(dumpResp))
}
