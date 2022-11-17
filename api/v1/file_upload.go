package v1

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samandar2605/post/api/models"
)

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// @Router /likes [post]
// @Summary Create a likes
// @Description Create a Likes
// @Tags Like
// @Accept json
// @Produce json
// @Param like body models.CreateLike true "like"
// @Success 201 {object} models.Like
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UploadFile(c *gin.Context) {
	var file File

	err := c.ShouldBindJSON(&file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	id := uuid.New()
	fileName := id.String() + filepath.Ext(file.File.Filename)

	dst, _ := os.Getwd()

	if _,err:=os.Stat(dst+"/media");os.IsNotExist(err){
		os.Mkdir(dst+"/media",os.ModePerm)
	}

	filePath:="/media/"+fileName
	if err=c.SaveUploadedFile(file.File,dst+filePath); err!=nil{
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}


	c.JSON(http.StatusCreated, gin.H{
		"filename": filePath,
	})
	
}
