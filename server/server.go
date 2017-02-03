package server

import(
	_ "github.com/vishvananda/netlink"
	"github.com/torukita/go-udpgen/util"
	"github.com/torukita/go-udpgen/server/resource"
	"github.com/torukita/go-udpgen/api"	
	"github.com/labstack/echo"
	"text/template"
	"net/http"
	"io"
	_ "fmt"
)

type Template struct {
	templates []*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	for _, v := range(t.templates) {
		if template := v.Lookup(name); template != nil {
			return template.ExecuteTemplate(w, name, data)
		}
	}
	return nil
}

func MainPage(c echo.Context) error {
	devices := util.DeviceList()
	/* test for api 
	arps := util.ArpList(2)
	for _, v := range(arps) {
		fmt.Println(v)
	}
    */
	return c.Render(http.StatusOK, "html", devices)
}

func JS(c echo.Context) error {
	return c.Render(http.StatusOK, "js", "")
}

func Run(addr string, debug bool) {
	// FIXME: should have better template handling
	jstemplate := template.Must(template.New("js").Parse(resource.JSText))
	htmltemplate := template.Must(template.New("html").Parse(resource.TemplateText))
	t := &Template{
		templates: []*template.Template{htmltemplate, jstemplate},
	}
	
	e := echo.New()
	e.Debug = debug
	e.Renderer = t

	// Define Rest API
	g := e.Group("/api")
	g.POST("/send", api.WebSend)
	
	e.GET("/", MainPage)
	e.GET("/js/utils.js", JS)
	e.Logger.Fatal(e.Start(addr))
}

