package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samandar2605/post/api/models"
	"github.com/samandar2605/post/storage/repo"
)

// @Router /likes/{id} [get]
// @Summary Get Like by id
// @Description Get Like by id
// @Tags Like
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Like
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetLike(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Like().Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Like{
		Id:     resp.Id,
		PostId: resp.PostId,
		UserId: resp.UserId,
		Status: resp.Status,
	})
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
func (h *handlerV1) CreateLike(c *gin.Context) {
	var (
		req models.CreateLike
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Like().Create(&repo.Like{
		PostId: req.PostId,
		UserId: req.UserId,
		Status: req.Status})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Like{
		Id:     resp.Id,
		PostId: resp.PostId,
		UserId: resp.UserId,
		Status: resp.Status,
	})
}

// @Summary Get Likes
// @Description Get Likes
// @Tags Like
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param post_id query int false "post_id"
// @Param user_id query int false "user_id"
// @Success 200 {object} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /likes [get]
func (h *handlerV1) GetAllLike(ctx *gin.Context) {
	queryParams, err := validateGetLikeQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := h.storage.Like().GetAll(queryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func validateGetLikeQuery(ctx *gin.Context) (repo.GetLikesQuery, error) {
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
			return repo.GetLikesQuery{}, err
		}
	}

	if ctx.Query("page") != "" {
		page, err = strconv.Atoi(ctx.Query("page"))
		if err != nil {
			return repo.GetLikesQuery{}, err
		}
	}
	if ctx.Query("post_id") != "" {
		postId, err = strconv.Atoi(ctx.Query("post_id"))
		if err != nil {
			return repo.GetLikesQuery{}, err
		}
	}

	if ctx.Query("user_id") != "" {
		userId, err = strconv.Atoi(ctx.Query("user_id"))
		if err != nil {
			return repo.GetLikesQuery{}, err
		}
	}
	return repo.GetLikesQuery{
		Limit:  limit,
		Page:   page,
		PostId: postId,
		UserId: userId,
	}, nil
}

// @Summary Update a like
// @Description Update a likes
// @Tags Like
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param like body models.CreateLike true "like"
// @Success 200 {object} models.Like
// @Failure 500 {object} models.ErrorResponse
// @Router /likes/{id} [put]
func (h *handlerV1) UpdateLike(ctx *gin.Context) {
	var b models.Like

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
	like, err := h.storage.Like().Update(&repo.Like{
		PostId: b.PostId,
		UserId: b.UserId,
		Status: b.Status,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create like",
		})
		return
	}

	ctx.JSON(http.StatusOK, like)
}

// @Summary Delete a like
// @Description Delete a like
// @Tags Like
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Failure 500 {object} models.ErrorResponse
// @Router /likes/{id} [delete]
func (h *handlerV1) DeleteLike(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert",
		})
		return
	}

	err = h.storage.Like().Delete(id)
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

