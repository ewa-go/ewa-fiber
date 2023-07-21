package ef

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestIContext(t *testing.T) {
	ctx := IContext(&fiber.Ctx{})
	fmt.Println(ctx.Method())
	fmt.Println(ctx.Path())
	fmt.Println(ctx.Scheme())
	fmt.Println(string(ctx.Body()))
	fmt.Println(ctx.Hostname())
	fmt.Println(ctx.IP())
}

func TestHttpRequest(t *testing.T) {

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		req := IContext(ctx).HttpRequest()
		fmt.Printf("%s\n", req.URL)
		return ctx.SendString("OK")
	})

	err := app.Listen(":8877")
	if err != nil {
		t.Fatal(err)
	}
}
