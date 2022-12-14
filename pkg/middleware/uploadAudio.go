package middleware

import (
	"context"
	dto "dumbsound/dto/result"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	// "github.com/golang-jwt/jwt/request"
)

func UploadAudio(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Upload file
		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		// r.ParseMultipartForm(10 * 1024 * 1024)
		file, _, err := r.FormFile("attache")

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Error Retrieving the File")
			return
		}
		defer file.Close()

		// setup file type filtering
		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		// filetype := http.DetectContentType(buff)

		filetype := "audio/mp3"
		fmt.Println(filetype)

		// if filetype != "audio/mp3" && filetype != "image/png" && filetype != "image/jpg" && filetype != "image/jpeg" {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "The provided file format is not allowed. Please upload a JPEG or PNG image"}
		// 	json.NewEncoder(w).Encode(response)
		// 	return
		// }

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		const MAX_UPLOAD_SIZE = 4000 << 20 // 10MB
		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "Max size in 1mb"}
			json.NewEncoder(w).Encode(response)
			return
		}

		fileTypeSplit := strings.Split(filetype, "/")
		tempFile, err := ioutil.TempFile("uploads", fileTypeSplit[0]+"-*."+fileTypeSplit[1])
		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload error")
			json.NewEncoder(w).Encode(err)
			return
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		// write this byte array to our temporary file
		tempFile.Write(fileBytes)

		data := tempFile.Name()
		filename := data[8:] // split uploads/
		filenameupdate := "http://localhost:5000/uploads/" + filename

		// add filename to ctx
		ctx := context.WithValue(r.Context(), "dataAudio", filenameupdate)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
