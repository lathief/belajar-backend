package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	os.MkdirAll("./uploads", 0755)

	router := gin.Default()

	// Limit total request body 10 MB
	router.MaxMultipartMemory = 10 << 20 // 10 MB

	router.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
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
		`))
	})

	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Single Upload
	router.POST("/upload", func(c *gin.Context) {

		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "File tidak ditemukan")
			return
		}

		const maxSize = 5 * 1024 * 1024 // 5 MB

		if file.Size > maxSize {
			c.String(http.StatusBadRequest, "Ukuran file maksimal 5 MB")
			return
		}

		if err := c.SaveUploadedFile(
			file,
			"./uploads/"+file.Filename,
		); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.String(http.StatusOK, "Upload berhasil")
	})

	// Multiple Upload
	router.POST("/upload-multiple", func(c *gin.Context) {

		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		files := form.File["files"]

		if len(files) == 0 {
			c.String(http.StatusBadRequest, "File wajib diisi")
			return
		}

		const maxSize = 5 * 1024 * 1024 // 5 MB

		for _, file := range files {

			if file.Size > maxSize {
				c.String(
					http.StatusBadRequest,
					fmt.Sprintf(
						"File %s melebihi batas 5 MB",
						file.Filename,
					),
				)
				return
			}

			if err := c.SaveUploadedFile(
				file,
				"./uploads/"+file.Filename,
			); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		}

		c.String(http.StatusOK, "Semua file berhasil diupload")
	})

	log.Fatal(router.Run(":5500"))
}
