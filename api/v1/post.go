package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samandar2605/post/api/models"
	"github.com/samandar2605/post/storage/repo"
)

// @Router /posts/{id} [get]
// @Summary Get post by id
// @Description Get post by id
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Post
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Post().Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Post{
		Id:          resp.Id,
		Title:       resp.Title,
		Description: resp.Description,
		ImageUrl:    resp.ImageUrl,
		UserId:      resp.UserId,
		CategoryId:  resp.CategoryId,
		UpdatedAt:   resp.UpdatedAt,
		ViewsCount:  resp.ViewsCount,
		CreatedAt:   resp.CreatedAt,
	})
}

// @Router /posts [post]
// @Summary Create a post
// @Description Create a post
// @Tags post
// @Accept json
// @Produce json
// @Param post body models.CreatePost true "post"
// @Success 201 {object} models.Post
// @Failure 500 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		req models.CreatePost
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Post().Create(&repo.Post{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserId:      req.UserId,
		CategoryId:  req.CategoryId,
		ViewsCount:  req.ViewsCount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Post{
		Id:          resp.Id,
		Title:       resp.Title,
		Description: resp.Description,
		ImageUrl:    resp.ImageUrl,
		UserId:      resp.UserId,
		CategoryId:  resp.CategoryId,
		ViewsCount:  resp.ViewsCount,
		UpdatedAt:   resp.UpdatedAt,
		CreatedAt:   resp.CreatedAt,
	})
}

// @Summary Get post
// @Description Get post
// @Tags post
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param search query string false "Search"
// @Success 200 {object} models.Post
// @Failure 500 {object} models.ErrorResponse
// @Router /posts [get]
func (h *handlerV1) GetPostAll(ctx *gin.Context) {
	queryParams, err := validateGetPostQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := h.storage.Post().GetAll(queryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func validateGetPostQuery(ctx *gin.Context) (repo.GetPostQuery, error) {
	var (
		limit int = 10
		page  int = 1
		err   error
	)
	if ctx.Query("limit") != "" {
		limit, err = strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			return repo.GetPostQuery{}, err
		}
	}

	if ctx.Query("page") != "" {
		page, err = strconv.Atoi(ctx.Query("page"))
		if err != nil {
			return repo.GetPostQuery{}, err
		}
	}

	return repo.GetPostQuery{
		Limit:  limit,
		Page:   page,
		Search: ctx.Query("search")}, nil
}

// @Summary Update a post
// @Description Update a post
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param user body models.CreatePost true "post"
// @Success 200 {object} models.Post
// @Failure 500 {object} models.ErrorResponse
// @Router /posts/{id} [put]
func (h *handlerV1) UpdatePost(ctx *gin.Context) {
	var b repo.Post

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
	post, err := h.storage.Post().Update(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create post",
		})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

// @Summary Delete a posts
// @Description Delete a posts
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Failure 500 {object} models.ErrorResponse
// @Router /posts/{id} [delete]
func (h *handlerV1) DeletePost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert",
		})
		return
	}

	err = h.storage.Post().Delete(id)
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
