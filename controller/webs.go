package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-web-project/core"
	"go-web-project/service"
	"golang.org/x/net/http2"
	"io"
	"net/http"
	"os"
	"time"
)

type CuteWebCRUDCode int

const (
	GetCode CuteWebCRUDCode = iota
	PostCode
	UpdateCode
	DeleteCode
)

type CuteWebServerConfig struct {
	Port                 uint16
	TLSPort              uint16
	TLSCrtPath           string
	TLSKeyPath           string
	MaxConcurrentStreams uint32
	MaxReadFrameSize     uint32
	IdleTimeout          time.Duration
}

type CuteWebServer struct {
	servMap service.CuteServiceMap
	serv    *echo.Echo
	config  CuteWebServerConfig
	ctx     core.CuteContext
}

func CreateWebServer(config CuteWebServerConfig,
	ctx core.CuteContext) *CuteWebServer {

	serv := echo.New()
	serv.Logger.SetOutput(io.Discard)

	return &CuteWebServer{serv: serv, config: config, ctx: ctx}
}

func (ptr *CuteWebServer) AddServiceMap(servMap service.CuteServiceMap) {
	ptr.servMap = servMap
}

func (ptr *CuteWebServer) AddStaticPage(prefix string, root string) *core.CuteError {
	if _, err := os.Stat(root); err != nil {
		innerErr := &core.CuteError{Code: core.InternalError}
		innerErr.SetError(err)
		return innerErr
	}
	ptr.serv.Static(prefix, root)
	return nil
}

func (ptr *CuteWebServer) AddStaticFile(prefix string, root string) *core.CuteError {
	if _, err := os.Stat(root); err != nil {
		innerErr := &core.CuteError{Code: core.InternalError}
		innerErr.SetError(err)
		return innerErr
	}
	ptr.serv.Static(prefix, root)
	return nil
}

func (ptr *CuteWebServer) AddServiceHandler(handlerName string, code CuteWebCRUDCode, input core.CuteData) *core.CuteError {
	if len(handlerName) < 1 {
		return &core.CuteError{Code: core.InternalError, Message: "The length of handlerName must be greater than 1."}
	}
	innerFunc := func(c echo.Context) error {
		if input != nil {
			if err := c.Bind(input); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON input"})
			}
		}
		key := c.QueryParam("key")
		if key == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid key"})
		}
		task, err := ptr.servMap.GetService(key, input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.GetMessage()})
		} else {
			result, execErr := task.Execute(ptr.ctx)
			if execErr != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": execErr.GetMessage()})
			}
			task.Destroy(ptr.ctx)
			serialize, err2 := result.JsonSerialize()
			if err2 != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err2.GetMessage()})
			} else {
				response := serialize
				return c.String(http.StatusOK, response)
			}

		}
	}
	switch code {
	case GetCode:
		ptr.serv.GET(handlerName, innerFunc)
		break
	case PostCode:
		ptr.serv.POST(handlerName, innerFunc)
		break
	case UpdateCode:
		ptr.serv.PUT(handlerName, innerFunc)
		break
	case DeleteCode:
		ptr.serv.DELETE(handlerName, innerFunc)
		break
	}

	return nil
}

func (ptr *CuteWebServer) AddMiddleWare() {
	ptr.serv.Use(middleware.CORS())
	ptr.serv.Use(middleware.Logger())
}

func (ptr *CuteWebServer) Start() error {
	addr := fmt.Sprintf(":%d", ptr.config.Port)
	return ptr.serv.Start(addr)
}
func (ptr *CuteWebServer) AutoTLSStart() error {
	addr := fmt.Sprintf(":%d", ptr.config.Port)
	return ptr.serv.StartAutoTLS(addr)
}
func (ptr *CuteWebServer) TLSStart() error {
	addr := fmt.Sprintf(":%d", ptr.config.Port)
	return ptr.serv.StartTLS(addr, ptr.config.TLSCrtPath, ptr.config.TLSKeyPath)
}
func (ptr *CuteWebServer) H2CServerStart() error {

	addr := fmt.Sprintf(":%d", ptr.config.Port)

	h2Serv := &http2.Server{
		MaxConcurrentStreams: ptr.config.MaxConcurrentStreams,
		MaxReadFrameSize:     ptr.config.MaxReadFrameSize,
		IdleTimeout:          ptr.config.IdleTimeout,
	}
	return ptr.serv.StartH2CServer(addr, h2Serv)
}
