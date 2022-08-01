package swagger

import (
	"bytes"
	"context"
	"io/ioutil"

	"github.com/saitofun/qkit/kit/httptransport/httpx"
	"github.com/saitofun/qkit/kit/kit"
)

var raw = bytes.NewBuffer(nil)

func init() {
	data, err := ioutil.ReadFile("./swagger.json")
	if err == nil {
		raw.Write(data)
	} else {
		raw.Write([]byte("{}"))
	}
}

var Router = kit.NewRouter(OpenAPI{})

type OpenAPI struct {
	httpx.MethodGet
}

func (s OpenAPI) Output(ctx context.Context) (interface{}, error) {
	return httpx.WrapContentType(httpx.MIME_JSON)(bytes.NewBuffer(raw.Bytes())), nil
}
