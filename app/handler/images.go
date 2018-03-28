package handler

import (
	"dienlanhphongvan/app/entity"
	"dienlanhphongvan/app/view"
	"dienlanhphongvan/cdnmodel"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
)

type imageHandler struct {
	imageEntity entity.Image
}

func (h imageHandler) Upload(c *gin.Context) {
	file, _, err := c.Request.FormFile("files")
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}
	name, err := h.imageEntity.Upload(file)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	url := view.Url{Value: name.Name()}
	c.JSON(200, view.Urls{
		Data: []view.Url{url},
	})
}

func (h imageHandler) Move(c *gin.Context) {
	uploadName, err := model.ParseUploadFilename(c.Param("name"))
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}
	imageName, err := model.ParseImageFilename(c.Query("name"), "", "")
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}
	err = h.imageEntity.Move(*uploadName, *imageName)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}
}

func (h imageHandler) GetOriginal(c *gin.Context) {
	imageName, err := model.ParseImageFilename(c.Param("name"), c.Query("type"), c.Query("w"))
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}
	filepath, err := h.imageEntity.GetOriginal(*imageName)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}
	c.File(filepath)
}
