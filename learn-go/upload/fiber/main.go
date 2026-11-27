package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	os.MkdirAll("./uploads", 0755)

	//Limit Total Request : Body Limit untuk membatasi ukuran body request, default 10MB jadi seluruh file totalnya tidak boleh lebih dari 10MB
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // 10 MB
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Type("html").SendString(`
<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Upload File</title>
</head>
<body>

    <h2>Upload Single File</h2>

     <form action="/upload" method="post" enctype="multipart/form-data">
        <input type="file" name="file" required>
        <button type="submit">Upload</button>
    </form>

    <h2>Upload Multiple File</h2>

    <form action="/upload-multiple" method="post" enctype="multipart/form-data">
        <input type="file" name="files" multiple required>
        <button type="submit">Upload</button>
    </form>

</body>
</html>
`)
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				SendString("File tidak ditemukan")
		}

		const maxSize = 5 * 1024 * 1024 // 5 MB

		if file.Size > maxSize {
			return c.Status(fiber.StatusBadRequest).
				SendString("Ukuran file maksimal 5 MB")
		}

		err = c.SaveFile(file, "./uploads/"+file.Filename)
		if err != nil {
			return err
		}

		return c.SendString("Upload berhasil")
	})

	app.Post("/upload-multiple", func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		files := form.File["files"]

		const maxSize = 5 * 1024 * 1024 // 5 MB

		if len(files) == 0 {
			return fiber.NewError(
				fiber.StatusBadRequest,
				"File wajib diisi",
			)
		}

		for _, file := range files {
			if file.Size > maxSize {
				return c.Status(fiber.StatusBadRequest).
					SendString(fmt.Sprintf(
						"File %s melebihi batas 5 MB",
						file.Filename,
					))
			}

			if err := c.SaveFile(file, "./uploads/"+file.Filename); err != nil {
				return err
			}
		}

		return c.SendString("Semua file berhasil diupload")
	})

	log.Fatal(app.Listen(":5500"))
}
