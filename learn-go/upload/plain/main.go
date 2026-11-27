package main

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	MaxRequestSize = 10 * 1024 * 1024 // 10 MB
	MaxFileSize    = 5 * 1024 * 1024  // 5 MB
)

func main() {

	if err := os.MkdirAll("./uploads", 0755); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/upload", singleUploadHandler)
	http.HandleFunc("/upload-multiple", multipleUploadHandler)

	log.Println("Server running :5500")
	log.Fatal(http.ListenAndServe(":5500", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	html := `
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

	<hr>

	<h2>Upload Multiple File</h2>

	<form action="/upload-multiple" method="post" enctype="multipart/form-data">
		<input type="file" name="files" multiple required>
		<button type="submit">Upload</button>
	</form>

</body>
</html>
`

	w.Header().Set("Content-Type", "text/html")
	_, _ = w.Write([]byte(html))
}

func singleUploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MaxRequestSize)

	if err := r.ParseMultipartForm(MaxRequestSize); err != nil {
		http.Error(w, "Request terlalu besar", http.StatusBadRequest)
		return
	}

	fileHeader, err := getSingleFile(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if fileHeader.Size > MaxFileSize {
		http.Error(w, "Ukuran file maksimal 5 MB", http.StatusBadRequest)
		return
	}

	if err := saveFile(fileHeader); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Upload berhasil")
}

func multipleUploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MaxRequestSize)

	if err := r.ParseMultipartForm(MaxRequestSize); err != nil {
		http.Error(w, "Request terlalu besar", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]

	if len(files) == 0 {
		http.Error(w, "File wajib diisi", http.StatusBadRequest)
		return
	}

	for _, file := range files {

		if file.Size > MaxFileSize {
			http.Error(
				w,
				fmt.Sprintf("File %s melebihi 5 MB", file.Filename),
				http.StatusBadRequest,
			)
			return
		}

		if err := saveFile(file); err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}
	}

	fmt.Fprintln(w, "Semua file berhasil diupload")
}

func getSingleFile(r *http.Request) (*multipart.FileHeader, error) {

	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("file tidak ditemukan")
	}

	file.Close()

	return header, nil
}

func saveFile(fileHeader *multipart.FileHeader) error {

	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(
		filepath.Join("./uploads", fileHeader.Filename),
	)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)

	return err
}
