package fiber

import "C"
import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

type Context struct {
	ctx *fiber.Ctx
}

func IContext(ctx *fiber.Ctx) Context {
	return Context{
		ctx: ctx,
	}
}

func (c *Context) Render(name string, bind interface{}, layouts ...string) error {
	return c.ctx.Render(name, bind, layouts...)
}

func (c *Context) Params(key string, defaultValue ...string) string {
	return c.ctx.Params(key, defaultValue...)
}

func (c *Context) Get(key string, defaultValue ...string) string {
	return c.ctx.Get(key, defaultValue...)
}

func (c *Context) Set(key, value string) {
	c.ctx.Set(key, value)
}

func (c *Context) SendStatus(code int) error {
	return c.ctx.SendStatus(code)
}

func (c *Context) Cookies(key string) string {
	return c.ctx.Cookies(key)
}

func (c *Context) SetCookie(cookie *http.Cookie) {
	fc := &fiber.Cookie{
		Name:     cookie.Name,
		Value:    cookie.Value,
		Path:     cookie.Path,
		Domain:   cookie.Domain,
		MaxAge:   cookie.MaxAge,
		Expires:  cookie.Expires,
		Secure:   cookie.Secure,
		HTTPOnly: cookie.HttpOnly,
		SameSite: strconv.Itoa(int(cookie.SameSite)),
	}
	c.ctx.Cookie(fc)
}

func (c *Context) ClearCookie(key string) {
	c.ctx.ClearCookie(key)
}

func (c *Context) Redirect(location string, status int) error {
	return c.ctx.Redirect(location, status)
}

func (c *Context) Path() string {
	return c.ctx.Path()
}

func (c *Context) SendString(code int, s string) error {
	return c.ctx.Status(code).SendString(s)
}

func (c *Context) Send(code int, contentType string, b []byte) error {
	c.ctx.Set(fiber.HeaderContentType, contentType)
	return c.ctx.Status(code).Send(b)
}

func (c *Context) SendFile(file string) error {
	return c.ctx.SendFile(file)
}

func (c *Context) SaveFile(fileHeader *multipart.FileHeader, path string) error {
	return c.ctx.SaveFile(fileHeader, path)
}

func (c *Context) SendStream(code int, contentType string, stream io.Reader) error {
	c.ctx.Set(fiber.HeaderContentType, contentType)
	return c.ctx.Status(code).SendStream(stream)
}

func (c *Context) JSON(code int, data interface{}) error {
	return c.ctx.Status(code).JSON(data)
}

func (c *Context) Body() []byte {
	return c.ctx.Body()
}

func (c *Context) BodyParser(out interface{}) error {
	return c.ctx.BodyParser(out)
}

func (c *Context) QueryParam(name string, defaultValue ...string) string {
	return c.ctx.Query(name, defaultValue...)
}

func (c *Context) QueryValues() url.Values {
	values := url.Values{}
	c.ctx.Request().URI().QueryArgs().VisitAll(func(key, value []byte) {
		values.Set(string(key), string(value))
	})
	return values
}

func (c *Context) QueryParams(h func(key, value string)) {
	c.ctx.Request().URI().QueryArgs().VisitAll(func(key, value []byte) {
		h(string(key), string(value))
	})
}

func (c *Context) Hostname() string {
	return c.ctx.Hostname()
}

func (c *Context) FormValue(name string) string {
	return c.ctx.FormValue(name)
}

func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	c.ctx.Port()
	return c.ctx.FormFile(name)
}

func (c *Context) Scheme() string {
	return c.ctx.Protocol()
}

func (c *Context) MultipartForm() (*multipart.Form, error) {
	return c.ctx.MultipartForm()
}

func (c *Context) IP() string {
	return c.ctx.IP()
}

func (c *Context) Context() context.Context {
	return c.ctx.Context()
}

func (c *Context) Ctx() interface{} {
	return c.ctx
}

func (c *Context) Method() string {
	return c.ctx.Method()
}

func (c *Context) HttpRequest() *http.Request {
	return nil
}

func (c *Context) Request() interface{} {
	return c.ctx.Request()
}
