package requests

import (
	"encoding/xml"
	"net/http"

	"kit.yo-star.com/go-kit/gf"
	"kit.yo-star.com/go-kit/json"
)

// Response request请求的返回结果
type Response struct {
	StatusCode  int
	Body        ResponseBody
	Header      http.Header
	RawResponse *http.Response
}

type ResponseBody []byte

func (rb ResponseBody) JsonBind(obj any) error {
	return json.Unmarshal(rb, obj)
}

func (rb ResponseBody) XMLBind(obj any) error {
	return xml.Unmarshal(rb, obj)
}

func (rb ResponseBody) String() string {
	return gf.BytesToString(rb)
}
