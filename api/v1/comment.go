package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samandar2605/post/api/models"
	"github.com/samandar2605/post/storage/repo"
)

// @Router /comments/{id} [get]
// @Summary Get comment by id
// @Description Get comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Comment
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Comment().Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Comment{
		Id:          resp.Id,
		PostId:      resp.PostId,
		UserId:      resp.UserId,
		Description: resp.Description,
		CreatedAt:   resp.CreatedAt,
		UpdatedAt:   resp.UpdatedAt,
	})
}

// @Router /comments [post]
// @Summary Create a comment
// @Description Create a comment
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body models.CreateComment true "comment"
// @Success 201 {object} models.Comment
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateComment(c *gin.Context) {
	var (
		req models.CreateComment
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Comment().Create(&repo.Comment{
		PostId:      req.PostId,
		UserId:      req.UserId,
		Description: req.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Comment{
		Id:          resp.Id,
		PostId:      resp.PostId,
		UserId:      resp.UserId,
		Description: resp.Description,
		CreatedAt:   resp.CreatedAt,
		UpdatedAt:   resp.UpdatedAt,
	})
}

// @Summary Get Likes
// @Description Get Likes
// @Tags comments
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param post_id query int false "post_id"
// @Param user_id query int false "user_id"
// @Success 200 {object} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /comments [get]
func (h *handlerV1) GetAllComment(ctx *gin.Context) {
	queryParams, err := validateGetCommentQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := h.storage.Comment().GetAll(queryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func validateGetCommentQuery(ctx *gin.Context) (repo.GetCommentQuery, error) {
	var (
		limit  int = 10
		page   int = 1
		postId int
		userId int
		err    error
	)
	if ctx.Query("limit") != "" {
		limit, err = strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			return repo.GetCommentQuery{}, err
		}
	}

	if ctx.Query("page") != "" {
		page, err = strconv.Atoi(ctx.Query("page"))
		if err != nil {
			return repo.GetCommentQuery{}, err
		}
	}
	if ctx.Query("post_id") != "" {
		postId, err = strconv.Atoi(ctx.Query("post_id"))
		if err != nil {
			return repo.GetCommentQuery{}, err
		}
	}

	if ctx.Query("user_id") != "" {
		userId, err = strconv.Atoi(ctx.Query("user_id"))
		if err != nil {
			return repo.GetCommentQuery{}, err
		}
	}
	return repo.GetCommentQuery{
		Limit:  limit,
		Page:   page,
		PostId: postId,
		UserId: userId,
	}, nil
}

// @Summary Update a comment
// @Description Update a comments
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param comment body models.CreateComment true "comment"
// @Success 200 {object} models.Comment
// @Failure 500 {object} models.ErrorResponse
// @Router /comments/{id} [put]
func (h *handlerV1) UpdateComment(ctx *gin.Context) {
	var b models.Comment

	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	b.Id = id
	comment, err := h.storage.Comment().Update(&repo.Comment{
		PostId:      b.PostId,
		UserId:      b.UserId,
		Description: b.Description,
	})
	fmt.Println(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create comment",
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// @Summary Delete a comment
// @Description Delete a comment
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Failure 500 {object} models.ErrorResponse
// @Router /comments/{id} [delete]
func (h *handlerV1) DeleteComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert",
		})
		return
	}

	err = h.storage.Comment().Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to Delete method",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "successful delete method",
	})
}
