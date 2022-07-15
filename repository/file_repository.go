package repository

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileRepository interface {
	Save(file multipart.File, fileName string) (string, error)
}

type fileRepository struct {
	path string
}

func (fr *fileRepository) Save(file multipart.File, fileName string) (string, error) {
	fileLocation := filepath.Join(fr.path, fileName)
	out, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	// defer out.Close()

	defer func(out *os.File) {
		err:= out.Close()
		if err != nil {
			panic(err)
		}
	}(out)
	
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}
	return fileLocation, nil
}

func NewFileRepository(path string) FileRepository {
	return &fileRepository{
		path: path,
	}
}
