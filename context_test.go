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
	ctx := IContext(&fiber.Ctx{})
	req, err := ctx.HttpRequest()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", req)
}
