package server

import(
	_ "github.com/vishvananda/netlink"
	"github.com/torukita/go-udpgen/util"
	"github.com/torukita/go-udpgen/server/resource"
	"github.com/torukita/go-udpgen/api"	
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"	
	"text/template"
	"net/http"
	"io"
	"strconv"
	"fmt"
)

var (
	logConfig = middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${uri} ${latency_human}\n",
	}
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
	return c.Render(http.StatusOK, "html", nil)
}

func JS(c echo.Context) error {
	return c.Render(http.StatusOK, "js", nil)
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

	e.Use(middleware.LoggerWithConfig(logConfig))
	//e.Use(middleware.Logger())
	// Define Rest API
	g := e.Group("/api")
	g.POST("/send", api.WebSend)
	g.POST("/config", api.WebSend)
	g.GET("/devices", getDevices) // Not used on this app
	g.GET("/device/:index/arp", getArpTable)
	g.GET("/device/:index/ipv4", getIPByIndex)
	
	e.GET("/", MainPage)
	e.GET("/js/utils.js", JS)
	e.Logger.Fatal(e.Start(addr))
}

func Test(c echo.Context) error {
	req := api.NewConfig()
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Printf("%+v", req)
	return c.JSON(http.StatusOK, nil)
}

func getDevices(c echo.Context) error {
	devices := util.DeviceList()
	return c.JSON(http.StatusOK, devices)
}

func getArpTable(c echo.Context) error {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	arps := util.ArpList(index)
	return c.JSON(http.StatusOK, arps)
}

func getIPByIndex(c echo.Context) error {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	ips := util.IPByIndex(index)
	return c.JSON(http.StatusOK, ips)
}
