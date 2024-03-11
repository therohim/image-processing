package main

import (
	"Ubersnap-middle-backend-programmer-test/exception"
	"Ubersnap-middle-backend-programmer-test/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/google/uuid"
)

func main() {
	// Setup Configuration
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})
	app.Use(recover.New())

	// Setup Routing
	app.Get("/health-check", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(map[string]interface{}{
			"code":   200,
			"status": "success",
			"IP":     ctx.IP(),
			"IPs":    ctx.IPs(),
		})
	})

	app.Get("/upload", func(c *fiber.Ctx) error {
		return c.Render("upload", fiber.Map{
			"Title": "Upload File",
		})
	})
	app.Post("/process", func(c *fiber.Ctx) error {
		return handleFileupload(c)
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Start App
	go func() {
		if err := app.Listen(":3000"); err != nil {
			exception.PanicIfNeeded(err)
		}
	}()

	// graceful shutdown
	<-stop
	log.Println("Stopping server...")
}

func handleFileupload(c *fiber.Ctx) error {
	// parse incomming image file
	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": http.StatusInternalServerError, "message": "Server error", "data": nil})

	}

	// generate new uuid for image name
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename
	fileArr := strings.Split(file.Filename, ".")
	fileExt := fileArr[len(fileArr)-1]

	if fileExt != "png" {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": http.StatusBadRequest, "message": "File uploaded must be png format", "data": nil})
	}

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	path := fmt.Sprintf("./images/%s", image)
	// save image to ./images dir
	err = c.SaveFile(file, path)

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": http.StatusInternalServerError, "message": "Server error", "data": nil})
	}

	var channel = make(chan *string)
	go func() {
		data := utils.ResizeImage(path, "images")
		channel <- data
	}()

	result := <-channel
	if result == nil {
		return c.JSON(fiber.Map{"status": http.StatusInternalServerError, "message": "Resize image fail", "data": nil})
	}

	err = os.Remove(path) //remove the file using built-in functions
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("Error removing file: %s", err.Error())) //print error if file is not removed
		return nil
	}
	// generate image url to serve to client using CDN

	imageUrl := fmt.Sprintf("http://localhost:3000/%s", *result)

	// create meta data and send to client

	data := map[string]interface{}{
		"imageName": result,
		"imageUrl":  imageUrl,
	}

	return c.JSON(fiber.Map{"status": http.StatusOK, "message": "Image uploaded successfully", "data": data})
}
