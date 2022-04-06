package client

import (
	"core/infra/errors"
	"core/infra/logger"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"mime/multipart"
	"net/http"
)

func UploadSingleImage(file *multipart.FileHeader, oldImagePath string, uploadServiceUrl string) (string, *errors.RestErr) {
	fileReader, frErr := file.Open()
	if frErr != nil {
		return "", errors.NewInternalServerError(frErr)

	}

	client := resty.New()
	resp, err := client.R().
		SetFileReader("file", file.Filename, fileReader).
		SetFormData(map[string]string{
			"old_image_path": oldImagePath,
		}).
		Post(uploadServiceUrl)

	if err != nil {
		return "", errors.NewInternalServerError(err)
	}

	if resp.StatusCode() == http.StatusCreated {
		var respBody UploadResponse

		tempBody := resp.Body()
		decodeErr := json.Unmarshal(tempBody, &respBody)
		if decodeErr != nil {
			return "", errors.NewInternalServerError(decodeErr)
		}

		return respBody.Data.ImagePath, nil
	} else {
		var respBody errors.RestErr

		decodeErr := json.NewDecoder(resp.RawResponse.Body).Decode(&respBody)
		if decodeErr != nil {
			return "", errors.NewInternalServerError(decodeErr)
		}

		logger.ErrorAsJson(fmt.Sprintf("Error on file upload:: %s ", uploadServiceUrl), respBody)
		return "", errors.NewBadRequestError(fmt.Sprintf("%s: %d", respBody.Error(), respBody.Status), nil)
	}
}

func UploadMultipleImages(files []*multipart.FileHeader, oldImagePaths []string, uploadServiceUrl string) (map[string]string, *errors.RestErr) {
	client := resty.New()
	req := client.R()

	for _, file := range files {
		fileReader, frErr := file.Open()
		if frErr != nil {
			return nil, errors.NewInternalServerError(frErr)
		}
		req = req.SetFileReader("file", file.Filename, fileReader)
	}

	if oldImagePaths != nil {
		for _, oldImagePath := range oldImagePaths {
			req = req.SetFormData(map[string]string{
				"old_image_path": oldImagePath,
			})
		}
	}

	resp, restErr := req.Post(uploadServiceUrl)
	if restErr != nil {
		return nil, errors.NewInternalServerError(restErr)
	}

	if resp.StatusCode() == http.StatusCreated {
		var respBody UploadResponse

		tempBody := resp.Body()
		decodeErr := json.Unmarshal(tempBody, &respBody)
		if decodeErr != nil {
			return nil, errors.NewInternalServerError(decodeErr)
		}

		return respBody.Data.ImagePathMap, nil
	} else {
		var respBody errors.RestErr

		decodeErr := json.NewDecoder(resp.RawResponse.Body).Decode(&respBody)
		if decodeErr != nil {
			return nil, errors.NewInternalServerError(decodeErr)
		}

		logger.ErrorAsJson(fmt.Sprintf("Error on file upload:: %s ", uploadServiceUrl), respBody)
		return nil, errors.NewBadRequestError(fmt.Sprintf("%s: %d", respBody.Error(), respBody.Status), nil)
	}
}

//func DeleteFile(filepath string) *errors.RestErr {
//	if filepath == "" {
//		return nil
//	}
//
//	client := resty.New()
//	resp, err := client.R().
//		SetBody(map[string]interface{}{
//			"file_path": filepath,
//		}).
//		Delete(config.InternalService().FileRemoveUrl)
//
//	if err != nil {
//		return errors.NewInternalServerError(err)
//	}
//
//	if resp.StatusCode() != http.StatusOK {
//		return errors.NewInternalServerError(errors.NewError("file remove error"))
//	}
//
//	return nil
//}

type UploadResponse struct {
	Message string       `json:"message"`
	Data    UploadedData `json:"data"`
}

type UploadedData struct {
	ImagePath    string            `json:"image_path"`
	ImagePathMap map[string]string `json:"image_path_map"`
	DeletedCount int               `json:"deleted_count"`
}
