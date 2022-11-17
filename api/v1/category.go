package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samandar2605/post/api/models"
	"github.com/samandar2605/post/storage/repo"
)

// @Router /categories/{id} [get]
// @Summary Get category by id
// @Description Get category by id
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Category().Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Category{
		Id:        resp.Id,
		Title:     resp.Title,
		CreatedAt: resp.CreatedAt,
	})
}

// @Router /categories [post]
// @Summary Create a category
// @Description Create a category
// @Tags category
// @Accept json
// @Produce json
// @Param category body models.CreateCategory true "Category"
// @Success 201 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var (
		req models.CreateCategory
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Category().Create(&repo.Category{
		Title: req.Title,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Category{
		Id:        resp.Id,
		Title:     resp.Title,
		CreatedAt: resp.CreatedAt,
	})
}

// @Summary Get Category
// @Description Get Category
// @Tags category
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param search query string false "Search"
// @Success 200 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
// @Router /categories [get]
func (h *handlerV1) GetCategoryAll(ctx *gin.Context) {
	queryParams, err := validateGetCategoryQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := h.storage.Category().GetAll(queryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func validateGetCategoryQuery(ctx *gin.Context) (repo.GetCategoryQuery, error) {
	var (
		limit int = 10
		page  int = 1
		err   error
	)
	if ctx.Query("limit") != "" {
		limit, err = strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			return repo.GetCategoryQuery{}, err
		}
	}

	if ctx.Query("page") != "" {
		page, err = strconv.Atoi(ctx.Query("page"))
		if err != nil {
			return repo.GetCategoryQuery{}, err
		}
	}

	return repo.GetCategoryQuery{
		Limit:  limit,
		Page:   page,
		Search: ctx.Query("search")}, nil
}

// @Summary Update a Category
// @Description Update a Category
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param user body models.CreateCategory true "Category"
// @Success 200 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
// @Router /categories/{id} [put]
func (h *handlerV1) UpdateCategory(ctx *gin.Context) {
	var b models.Category

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
	category, err := h.storage.Category().Update(&repo.Category{
		Title: b.Title,
	})
	fmt.Println(category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create category",
		})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// @Summary Delete a categories
// @Description Delete a categories
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Failure 500 {object} models.ErrorResponse
// @Router /categories/{id} [delete]
func (h *handlerV1) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert",
		})
		return
	}

	err = h.storage.Category().Delete(id)
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
