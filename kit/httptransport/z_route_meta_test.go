package httptransport_test

import (
	"fmt"
	"os"
	"sort"

	"github.com/fatih/color"

	. "github.com/saitofun/qkit/kit/httptransport"
	"github.com/saitofun/qkit/testutil/httptransporttestutil/server/cmd/app/routes"
)

func ExampleGroup() {
	g := Group("/test")
	fmt.Println(g.Path())
	// Output:
	// /test
}

func ExampleHttpRouteMeta() {
	os.Setenv(EnvProjectName, "service-example")
	os.Setenv(EnvProjectVersion, "1.0.0")

	color.NoColor = true

	routes := routes.RootRouter.Routes()

	sort.Slice(routes, func(i, j int) bool {
		return NewHttpRouteMeta(routes[i]).Key() <
			NewHttpRouteMeta(routes[j]).Key()
	})

	for i := range routes {
		rm := NewHttpRouteMeta(routes[i])
		fmt.Println(rm.String())
	}
	// Output:
	// GET /demo swagger.OpenAPI
	// GET /demo/binary/files routes.DownloadFile
	// GET /demo/binary/images routes.ShowImage
	// POS /demo/cookie routes.Cookie
	// POS /demo/forms/multipart routes.FormMultipartWithFile
	// POS /demo/forms/multipart-with-files routes.FormMultipartWithFiles
	// POS /demo/forms/urlencoded routes.FormURLEncoded
	// GET /demo/proxy routes.Proxy
	// GET /demo/redirect routes.Redirect
	// POS /demo/redirect routes.RedirectWhenError
	// POS /demo/restful routes.Create
	// HEA /demo/restful routes.HealthCheck
	// GET /demo/restful/{id} routes.DataProvider routes.GetByID
	// DEL /demo/restful/{id} routes.DataProvider routes.RemoveByID
	// PUT /demo/restful/{id} routes.DataProvider routes.UpdateByID
	// GET /demo/v2/proxy routes.ProxyV2
}
