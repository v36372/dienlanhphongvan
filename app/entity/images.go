package entity

import (
	"crypto/md5"
	"dienlanhphongvan-cdn/client"
	"dienlanhphongvan-cdn/model"
	"dienlanhphongvan-cdn/util"
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
	Imgx        *client.Client
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
	GetCached(name model.Filename) (string, error)
	Delete(images []string) error
}

const (
	cachedImagePrefix = "/images/cached/"
)

func NewImage(imgx *client.Client, uploadDir, originalDir, cachedDir string, debug bool) *imageEntity {
	return &imageEntity{
		Imgx:        imgx,
		UploadDir:   uploadDir,
		OriginalDir: originalDir,
		CachedDir:   cachedDir,
		debug:       debug,
	}

}

func (i imageEntity) Upload(f io.Reader) (*model.UploadFilename, error) {
	uuid := uniqueStr(time.Now().UnixNano())
	name := model.NewUploadFilename(uuid, time.Now())
	filepath := path.Join(i.UploadDir, name.Path())
	if err := file.WriteFile(filepath, f); err != nil {
		return nil, err
	}
	return &name, nil
}

func (i imageEntity) GetOriginal(name model.Filename) (string, error) {
	filepath := path.Join(i.OriginalDir, name.Path())
	if file.ExistFile(filepath) {
		return filepath, nil
	}

	return "", uer.NotFoundError(errors.New("image not found"))
}

func (i imageEntity) GetCached(name model.Filename) (string, error) {
	cachedPath := path.Join(i.CachedDir, name.Path())
	if util.ExistFile(cachedPath) {
		return cachedPath, nil

	}
	// crop or resize original file
	originalPath, err := i.GetOriginal(name)
	if err != nil {
		return "", err

	}
	var f io.Reader
	switch name.Shape() {
	case "o":
		f, err = i.Imgx.Image.Resize(originalPath, client.ResizeOption{
			Width: name.Width(),
		})
	case "s":
		f, err = i.Imgx.Image.Crop(originalPath, client.CropOption{
			Width: name.Width(),
		})
	default:
		f, err = i.Imgx.Image.Resize(originalPath, client.ResizeOption{
			Width: name.Width(),
		})

	}
	if err != nil {
		return "", uer.InternalError(err)

	}
	if err := util.WriteFile(cachedPath, f); err != nil {
		return "", uer.InternalError(err)

	}
	return cachedPath, nil

}

func isOldImage(image string) bool {
	return strings.HasPrefix(image, cachedImagePrefix)
}

func (i imageEntity) Delete(images []string) error {
	for _, image := range images {
		if len(image) == 0 {
			continue
		}
		fromName, err := model.ParseUploadFilename(image)
		if err != nil {
			return uer.InternalError(err)
		}
		originImage := path.Join(i.OriginalDir, fromName.Path())
		cachedImage := path.Join(i.CachedDir, fromName.Path())

		if !file.ExistFile(originImage) {
			continue
		}
		if err := file.RemoveFile(originImage); err != nil {
			return uer.InternalError(err)
		}
		if !file.ExistFile(cachedImage) {
			continue
		}
		if err := file.RemoveFile(cachedImage); err != nil {
			return uer.InternalError(err)
		}
	}

	return nil
}

func (i imageEntity) MoveImagesOfProduct(images []string) (oimages []string, err error) {
	for _, image := range images {
		if len(image) == 0 {
			oimages = append(oimages, image)
			continue
		}
		if isOldImage(image) {
			if index := strings.Index(image, "?"); index > 0 {
				image = image[:index]
			}

			oimages = append(oimages, image[len(cachedImagePrefix):])
			continue
		}
		fromName, err := model.ParseUploadFilename(image)
		if err != nil {
			return oimages, uer.InternalError(err)
		}
		fromPath := path.Join(i.UploadDir, fromName.Path())
		toName := model.NewImageFilename(image, time.Now())
		toPath := path.Join(i.OriginalDir, toName.Path())

		if !file.ExistFile(fromPath) {
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
	fromPath := path.Join(i.UploadDir, fromName.Path())
	toPath := path.Join(i.OriginalDir, toName.Path())

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
