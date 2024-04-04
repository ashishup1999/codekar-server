package handlers

import (
	"bytes"
	"codekar/app/services"
	"encoding/base64"
	"image"
	"image/png"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

func UploadProfileImgHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("image")
	userName := c.PostForm("userName")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Resize the image to width 300 and keep the aspect ratio
	resizedImg := imaging.Resize(img, 300, 0, imaging.Lanczos)

	// Creating file buffer
	fileBytes := bytes.NewBuffer(nil)
	err = png.Encode(fileBytes, resizedImg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Encode the file content as base64
	base64String := base64.StdEncoding.EncodeToString(fileBytes.Bytes())

	//uploading base64String to db
	resp := services.UpdateProfilePictureService(userName, base64String)

	c.JSON(http.StatusOK, gin.H{"status": resp.Status, "message": resp.Message, "profileImg": base64String})
}
