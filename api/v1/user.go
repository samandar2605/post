package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samandar2605/post/api/models"
	"github.com/samandar2605/post/storage/repo"
)

// @Router /users/{id} [get]
// @Summary Get user by id
// @Description Get user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.User
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.User().Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.User{
		Id:              resp.Id,
		FirstName:       resp.FirstName,
		LastName:        resp.LastName,
		PhoneNumber:     resp.PhoneNumber,
		Email:           resp.Email,
		Gender:          resp.Gender,
		Password:        resp.Password,
		Username:        resp.UserName,
		ProfileImageUrl: resp.ProfileImageUrl,
		Type:            resp.Type,
		CreatedAt:       resp.CreatedAt,
	})
}

// @Router /users [post]
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "user"
// @Success 201 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		req models.CreateUser
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.User().Create(&repo.User{
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		PhoneNumber:     req.PhoneNumber,
		Email:           req.Email,
		Gender:          req.Gender,
		UserName:        req.UserName,
		Password:        req.Password,
		ProfileImageUrl: req.ProfileImageUrl,
		Type:            req.Type,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.User{
		Id:              resp.Id,
		FirstName:       resp.FirstName,
		LastName:        resp.LastName,
		PhoneNumber:     resp.PhoneNumber,
		Email:           resp.Email,
		Gender:          resp.Gender,
		Password:        resp.Password,
		Username:        resp.UserName,
		ProfileImageUrl: resp.ProfileImageUrl,
		Type:            resp.Type,
	})
}

// @Summary Get users
// @Description Get users
// @Tags users
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param search query string false "Search"
// @Success 200 {object} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /users [get]
func (h *handlerV1) GetUserAll(ctx *gin.Context) {
	queryParams, err := validateGetUsersQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := h.storage.User().GetAll(repo.GetUserQuery{
		Page:   queryParams.Page,
		Limit:  queryParams.Limit,
		Search: queryParams.Search,
	} )
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func validateGetUsersQuery(ctx *gin.Context) (repo.GetUserQuery, error) {
	var (
		limit int = 10
		page  int = 1
		err   error
	)
	if ctx.Query("limit") != "" {
		limit, err = strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			return repo.GetUserQuery{}, err
		}
	}

	if ctx.Query("page") != "" {
		page, err = strconv.Atoi(ctx.Query("page"))
		if err != nil {
			return repo.GetUserQuery{}, err
		}
	}

	return repo.GetUserQuery{
		Limit:  limit,
		Page:   page,
		Search: ctx.Query("search"),
	}, nil
}

// @Summary Update a user
// @Description Update a userss
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param user body models.CreateUser true "User"
// @Success 200 {object} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [put]
func (h *handlerV1) UpdateUser(ctx *gin.Context) {
	var (
		req models.User
	)

	err := ctx.ShouldBindJSON(&req)
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

	req.Id = id
	user, err := h.storage.User().Update(&repo.User{
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		PhoneNumber:     req.PhoneNumber,
		Email:           req.Email,
		Gender:          req.Gender,
		UserName:        req.Username,
		Password:        req.Password,
		ProfileImageUrl: req.ProfileImageUrl,
		Type:            req.Type,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary Delete a User
// @Description Delete a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [delete]
func (h *handlerV1) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert",
		})
		return
	}

	err = h.storage.User().Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "successful delete method",
	})
}
