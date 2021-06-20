package handlers

import (
	"api-storage-github/models"
	"api-storage-github/repositories"
	"api-storage-github/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SaveCode(c *fiber.Ctx) error {
	body := models.CodeBody{}
	c.BodyParser(&body)
	codeInRepository := repositories.SendCodeInRepository(body)
	return c.Status(201).JSON(codeInRepository)
}

func ExecuteCodefunc(c *fiber.Ctx) error {
	scriptName := c.Params("scriptName")
	code := repositories.GetCodeInRepositoryByScriptName(scriptName)
	command := fmt.Sprintf("docker run node:12-alpine3.10 echo '%s' > index.js && node index.js", code)
	return c.SendString(utils.Command(command))
}
