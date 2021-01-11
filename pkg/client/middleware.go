package client

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type EchoMiddleware struct {
	Client  Client
	Service string
}

func (m *EchoMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	piped := pipeMiddlewares(
		m.ExtractWorkspace,
		m.ExtractUserToken,
		m.Authenticate,
		m.Authorize,
	)
	return piped(next)
}

func (m *EchoMiddleware) ExtractWorkspace(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if wsid := c.Param("workspace"); wsid != "" {
			c.Set("workspaceuuid", wsid)
		}
		return next(c)
	}
}

func (m *EchoMiddleware) ExtractUserToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if userToken := bearerToken(c.Request()); userToken != "" {
			c.Set("usertoken", userToken)
		}
		return next(c)
	}
}

func (m *EchoMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken, ok := c.Get("usertoken").(string)
		if userToken == "" || !ok {
			return c.NoContent(http.StatusBadRequest)
		}

		user, err := m.Client.Identify(userToken, "")
		if err != nil {
			return c.NoContent(http.StatusForbidden)
		}

		c.Set("useruuid", user.UUID)

		return next(c)
	}
}

func (m *EchoMiddleware) Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		workspaceUUID, ok := c.Get("workspaceuuid").(string)
		if workspaceUUID == "" || !ok {
			return c.NoContent(http.StatusForbidden)
		}

		identity, ok := c.Get("useruuid").(string)
		if identity == "" || !ok {
			return c.NoContent(http.StatusForbidden)
		}
		method := c.Request().Method
		object := rri(workspaceUUID, m.Service, c.Request().URL.Path)

		if err := m.Client.Authorize(identity, method, object); err != nil {
			return c.NoContent(http.StatusForbidden)
		}

		return next(c)
	}
}

func pipeMiddlewares(middlewares ...echo.MiddlewareFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		iterHandler := next
		for i := len(middlewares) - 1; i >= 0; i-- {
			iterHandler = middlewares[i](iterHandler)
		}
		return iterHandler
	}
}

func rri(workspace, service, resource string) string {
	return fmt.Sprintf("rri:v1:cafebazaar.cloud:%s:%s:%s", workspace, service, resource)
}

func bearerToken(r *http.Request) string {
	value := r.Header.Get("Authorization")
	splitted := strings.Split(value, " ")
	if len(splitted) < 2 {
		return ""
	}
	if splitted[0] != "Bearer" {
		return ""
	}
	return splitted[1]
}
