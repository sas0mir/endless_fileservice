package handler

import (
	"fmt"
	"net/http"
)

type File struct{}

func (f *File) Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Upload file")
}

func (f *File) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List files")
}

func (f *File) Download(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Download file")
}

func (f *File) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete file")
}
