package sicgolib

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/codedius/imagekit-go"
)

/*
CreateNewImagekitClient return an imagekit client and an error.
The function will need public key and private key from desired imagekit in order to use the function.
*/
func CreateNewImagekitClient(publicKey string, privateKey string) (*imagekit.Client, error) {
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
func UploadToImagekit(ctx context.Context, imageKitClient *imagekit.Client, publicKey string, privateKey string, file []byte, fileName string, folderName string) (*imagekit.UploadResponse, error) {
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

	upr, err := imageKitClient.Upload.ServerUpload(ctx, &ur)
	if err != nil {
		log.Println(upr)
	}
	return upr, nil
}

/*
ParseImageFile returns byte file, multipart file header, and an error.
This function will need *http.Request and an <form> <input> name from the Front-End as specifier.
Example: <input type="file" name="{inputName}" accept="image/*" />.
This function is used to parse file into file bytes that needed in order to work with UploadToImagekit function.
Defining maxFileSize Example: (10 << 20) it will be around 10mb because (10 << 20) will be the same as (10 * (2^20)) = (10 * 1,048,576) = 10,048,576 around 10mb
*/
func ParseImageFile(r *http.Request, inputName string, maxFileSize int64) ([]byte, *multipart.FileHeader, error) {
	err := validateFileSize(maxFileSize)
	if err != nil {
		return nil, nil, err
	}
	r.ParseMultipartForm(maxFileSize)

	file, handler, err := r.FormFile(inputName)

	if err != nil {
		log.Println("Error Retrieving File: ", err.Error())
		return nil, nil, err
	}

	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error Converting File: ", err)
		return nil, nil, err
	}
	return fileBytes, handler, nil
}

func validateFileSize(maxFileSize int64) error {
	if maxFileSize < 1 {
		err := errors.New("INVALID MEMORY SIZE")
		return err
	}
	return nil
}
