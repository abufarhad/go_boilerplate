package fileutil

import (
	"core/infra/errors"
	"core/utils/consts"
	"core/utils/hash"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func UploadFile(fileHeader *multipart.FileHeader, uploadPath string) (string, *errors.RestErr) {
	uploadPath = consts.AssetFolder + "/" + uploadPath
	os.MkdirAll(uploadPath, os.ModePerm)

	hashedFileName := GenerateHashName(fileHeader.Filename)
	filePath := filepath.Join(uploadPath, hashedFileName)

	file, err := fileHeader.Open()
	if err != nil {
		return "", errors.NewInternalServerError(err)
	}

	defer file.Close()

	createdFile, err := os.Create(filePath)
	if err != nil {
		return "", errors.NewInternalServerError(err)
	}

	defer createdFile.Close()
	_, err = io.Copy(createdFile, file)
	if err != nil {
		return "", errors.NewInternalServerError(err)
	}

	return filePath, nil
}

func GenerateHashName(filename string) string {
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	hashedName := hash.GetSha1Hash(name) + ext
	return hashedName
}

func RemoveFile(filepath string) *errors.RestErr {

	// Checking File exists or not
	if _, fErr := os.Stat(filepath); fErr == nil {
		err := os.Remove(filepath)

		if err != nil {
			return errors.NewInternalServerError(err)
		}
	}
	return nil
}
