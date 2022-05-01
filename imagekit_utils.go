package sicgolib

import (
	"context"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/codedius/imagekit-go"
)

/*
createNewImagekitClient return an imagekit client and an error.
The function will need public key and private key from desired imagekit in order to use the function.
*/
func createNewImagekitClient(publicKey string, privateKey string) (*imagekit.Client, error) {
	opts := imagekit.Options{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}

	imageKit, err := imagekit.NewClient(&opts)
	if err != nil {
		log.Println("ERROR (ImageKit): Error while creating imagekit client options")
		return imageKit, err
	}
	return imageKit, nil
}

/*
UploadToImagekit returns an imagekit Upload Response and an error that you can store if you need to know the upload responses and the error.
The function will need some parameters: context, public key, private key, byte image File, file name, destination folder name on imagekit.
You can only upload 1 image with this function according to imagekit rules. You need to use looping if you want to upload more than 1 file.
*/
func UploadToImagekit(ctx context.Context, publicKey string, privateKey string, file []byte, fileName string, folderName string) (*imagekit.UploadResponse, error) {
	ur := imagekit.UploadRequest{
		File:              file,
		FileName:          fileName,
		UseUniqueFileName: false,
		Tags:              []string{},
		Folder:            folderName,
		IsPrivateFile:     false,
		CustomCoordinates: "",
		ResponseFields:    nil,
	}
	client, err := createNewImagekitClient(publicKey, privateKey)
	if err != nil {
		return nil, err
	}

	upr, err := client.Upload.ServerUpload(ctx, &ur)
	if err != nil {
		log.Println(upr)
	}
	return upr, nil
}

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
