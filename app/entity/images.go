package entity

import (
	"crypto/md5"
	"dienlanhphongvan/cdnmodel"
	"dienlanhphongvan/utilities/file"
	"dienlanhphongvan/utilities/uer"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type imageEntity struct {
	UploadDir   string
	OriginalDir string
	CachedDir   string
	debug       bool
}

type Image interface {
	Upload(f io.Reader) (*model.UploadFilename, error)
	GetOriginal(name model.Filename) (string, error)
	Move(fromName model.UploadFilename, toName model.Filename) error
	MoveImagesOfProduct(images []string) ([]string, error)
}

func NewImage(uploadDir, originalDir, cachedDir string, debug bool) *imageEntity {
	return &imageEntity{
		UploadDir:   uploadDir,
		OriginalDir: originalDir,
		CachedDir:   cachedDir,
		debug:       debug,
	}
}

func (i imageEntity) Upload(f io.Reader) (*model.UploadFilename, error) {
	uuid := uniqueStr(time.Now().UnixNano())
	name := model.NewUploadFilename(uuid, time.Now())
	filepath := path.Join("/home/justin/workspace/src/dienlanhphongvan/images/products/tmp/", name.Path())
	if err := file.WriteFile(filepath, f); err != nil {
		return nil, err
	}
	return &name, nil
}

func (i imageEntity) GetOriginal(name model.Filename) (string, error) {
	filepath := path.Join("/home/justin/workspace/src/dienlanhphongvan/images/products/original/", name.Path())
	if file.ExistFile(filepath) {
		return filepath, nil
	}

	return "", uer.NotFoundError(errors.New("image not found"))
}

func (i imageEntity) MoveImagesOfProduct(images []string) (oimages []string, err error) {
	for _, image := range images {
		if len(image) == 0 {
			continue
		}
		fromName, err := model.ParseUploadFilename(image)
		if err != nil {
			return oimages, uer.InternalError(err)
		}
		fromPath := path.Join("/home/justin/workspace/src/dienlanhphongvan/images/products/tmp/", fromName.Path())
		toName := model.NewImageFilename(image, time.Now())
		toPath := path.Join("/home/justin/workspace/src/dienlanhphongvan/images/products/original/", toName.Path())

		if !file.ExistFile(fromPath) {
			fmt.Println(fromPath)
			return oimages, uer.NotFoundError(errors.New("image not found"))
		}
		f, err := os.Open(fromPath)
		if err != nil {
			err = uer.InternalError(err)
			return images, err
		}
		if err := file.WriteFile(toPath, f); err != nil {
			err = uer.InternalError(err)
			return images, err
		}

		oimages = append(oimages, toName.Name())
	}

	return
}

func (i imageEntity) Move(fromName model.UploadFilename, toName model.Filename) error {
	fromPath := path.Join("/home/justin/workspace/src/dienlanhphongvan/images/products/tmp/", fromName.Path())
	toPath := path.Join("/home/justin/workspace/src/dienlanhphongvan/images/products/original/", toName.Path())

	if !file.ExistFile(fromPath) {
		return uer.NotFoundError(errors.New("image not found"))
	}
	f, err := os.Open(fromPath)
	if err != nil {
		return uer.InternalError(err)
	}
	if err := file.WriteFile(toPath, f); err != nil {
		return uer.InternalError(err)
	}
	return nil
}

func uniqueStr(sec int64) string {
	hash := md5.New()
	v4, _ := uuid.NewV4()
	uid := strings.Replace(v4.String(), "-", "", -1)
	io.WriteString(hash, fmt.Sprintf("%s%v", uid, sec))
	return hex.EncodeToString(hash.Sum(nil))
}
