package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mirasildev/exam_project_2.0/api/models"
	"github.com/mirasildev/exam_project_2.0/storage/repo"
)
// @Security ApiKeyAuth


// @Router /hotels [post]
// @Summary Create a hotel
// @Description Create a hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Param hotel body models.CreateHotelRequest true "hotel"
// @Success 200 {object} models.CreateHotelResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateHotel(c *gin.Context) {
	var (
		req models.CreateHotelRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	images := []*repo.HotelImage{}
	for _, image := range req.Images {
		i := repo.HotelImage{
			ImageUrl: image.ImageUrl,
			SequenceNumber: image.SequenceNumber,
		}
		images = append(images, &i)
	}

	id, err := h.storage.Hotel().Create(&repo.Hotel{
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		ImageUrl:    req.ImageUrl,
		NumOfRooms:  req.NumOfRooms,
		UserID:      payload.UserID,
		Images:      images,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, id)
}

func parseHotelModel(hotel *repo.Hotel) models.Hotel {
	return models.Hotel{
		ID:          hotel.ID,
		Name:        hotel.Name,
		Description: hotel.Description,
		Address:     hotel.Address,
		ImageUrl:    hotel.ImageUrl,
		NumOfRooms:  hotel.NumOfRooms,
		UserID:      hotel.UserID,
	}
}

// @Router /hotels/{id} [get]
// @Summary Get hotel by id
// @Description Get hotel by id
// @Tags hotel
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Hotel
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetHotel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Hotel().Get(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	hotel := parseHotelModel(resp)

	c.JSON(http.StatusOK, hotel)
}

// @Router /hotels [get]
// @Summary Get all hotels
// @Description Get all hotels
// @Tags hotel
// @Accept json
// @Produce json
// @Param filter query models.GetAllHotelsParams false "Filter"
// @Success 200 {object} models.GetAllHotelsResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllHotels(c *gin.Context) {
	req, err := validateGetAllHotelsParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Hotel().GetAll(&repo.GetAllHotelsParams{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getHotelsResponse(result))
}

func validateGetAllHotelsParams(c *gin.Context) (*models.GetAllHotelsParams, error) {
	var (
		limit int = 10
		page  int = 1
		err   error
	)

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			return nil, err
		}
	}

	return &models.GetAllHotelsParams{
		Limit: int32(limit),
		Page:  int32(page),
	}, nil
}

func getHotelsResponse(data *repo.GetAllHotelsResult) *models.GetAllHotelsResponse {
	response := models.GetAllHotelsResponse{
		Hotels: make([]*models.Hotel, 0),
		Count:  data.Count,
	}

	for _, hotel := range data.Hotels {
		p := parseHotelModel(hotel)
		response.Hotels = append(response.Hotels, &p)
	}

	return &response
}

// @Router /hotels/{id} [put]
// @Summary Update a hotel
// @Description Update a hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param hotel body models.UpdateHotelRequest true "Hotel"
// @Success 200 {object} models.Hotel
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateHotel(c *gin.Context) {
	var (
		req models.UpdateHotelRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	images := []*repo.HotelImage{}
	for _, image := range req.Images {
		i := repo.HotelImage{
			ImageUrl: image.ImageUrl,
			SequenceNumber: image.SequenceNumber,
		}
		images = append(images, &i)
	}

	updated, err := h.storage.Hotel().Update(&repo.Hotel{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		ImageUrl:    req.ImageUrl,
		NumOfRooms:  req.NumOfRooms,
		Images:      images,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// @Router /hotels/{id} [delete]
// @Summary Delete a hotels
// @Description Delete a hotels
// @Tags hotel
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteHotel(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.storage.Hotel().Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted!",
	})
}