package response

import (
	"github.com/elton/project-layout/app/myapp/pkg/e"
	"github.com/gofiber/fiber/v2"
)

// ResultJSON show the result of responses using JSON.
func ResultJSON(c *fiber.Ctx, errCode int, data interface{}) {
	c.JSON(fiber.Map{"code": errCode, "msg": e.GetMsg(errCode), "data": data})
}
