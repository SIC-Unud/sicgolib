package sicgolib

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

/*
ParseImageFile returns byte file, multipart file header, and an error.
This function will need *http.Request and an <form> <input> name as parameters.
This function is used to parse file into file bytes that needed in order to work with UploadToImagekit function.
*/
func ParseImageFile(r *http.Request, inputName string) ([]byte, *multipart.FileHeader, error) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile(inputName)

	if err != nil {
		log.Println("Error Retrieving File: ", err.Error())
		return nil, nil, err
	}

	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error converting file: ", err)
		return nil, nil, err
	}
	return fileBytes, handler, nil
}
