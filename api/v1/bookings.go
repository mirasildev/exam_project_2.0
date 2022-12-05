package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mirasildev/exam_project_2.0/api/models"
	"github.com/mirasildev/exam_project_2.0/storage/repo"
)

// @Security ApiKeyAuth

// @Router /bookings [post]
// @Summary Create a booking
// @Description Create a booking
// @Tags booking
// @Accept json
// @Produce json
// @Param booking body models.CreateBookingRequest true "booking"
// @Success 200 {object} models.CreateBookingResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateBooking(c *gin.Context) {
	var (
		req models.CreateBookingRequest
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

	id, err := h.storage.Booking().Create(&repo.Booking{
		Arrival:    req.Arrival,
		Checkout:   req.Checkout,
		RoomID:     req.RoomID,
		RoomNumber: req.RoomNumber,
		UserID:     payload.UserID,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, id)
}

func parseBookingModel(booking *repo.Booking) models.Booking {
	return models.Booking{
		ID:         booking.ID,
		// Arrival:    booking.Arrival,
		// Checkout:   booking.Checkout,
		RoomID:     booking.RoomID,
		RoomNumber: booking.RoomNumber,
		UserID:     booking.UserID,
	}
}

// @Router /bookings/{id} [get]
// @Summary Get bookings by id
// @Description Get bookings by id
// @Tags booking
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Booking
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetBooking(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Booking().Get(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	booking := parseBookingModel(resp)

	c.JSON(http.StatusOK, booking)
}

// @Router /bookings [get]
// @Summary Get all bookings
// @Description Get all bookings
// @Tags booking
// @Accept json
// @Produce json
// @Param filter query models.GetAllBookingsParams false "Filter"
// @Success 200 {object} models.GetAllBookingsResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllBookings(c *gin.Context) {
	req, err := validateGetAllBookingsParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Booking().GetAll(&repo.GetAllBookingsParams{
		Page:  req.Page,
		Limit: req.Limit,
		RoomID: req.RoomID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getBookingsResponse(result))
}

func validateGetAllBookingsParams(c *gin.Context) (*models.GetAllBookingsParams, error) {
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

	return &models.GetAllBookingsParams{
		Limit: int32(limit),
		Page:  int32(page),
	}, nil
}

func getBookingsResponse(data *repo.GetAllBookingsResult) *models.GetAllBookingsResponse {
	response := models.GetAllBookingsResponse{
		Bookings: make([]*models.Booking, 0),
		Count:    data.Count,
	}

	for _, booking := range data.Bookings {
		p := parseBookingModel(booking)
		response.Bookings = append(response.Bookings, &p)
	}

	return &response
}

// @Router /bookings/{id} [put]
// @Summary Update a booking
// @Description Update a booking
// @Tags booking
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param booking body models.UpdateBookingRequest true "Booking"
// @Success 200 {object} models.Booking
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateBooking(c *gin.Context) {
	var (
		req models.UpdateBookingRequest
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

	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	updated, err := h.storage.Booking().Update(&repo.Booking{
		ID: id,
		Arrival:    req.Arrival,
		Checkout:   req.Checkout,
		RoomID:     req.RoomID,
		RoomNumber: req.RoomNumber,
		UserID:     payload.UserID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// @Router /bookings/{id} [delete]
// @Summary Delete a booking
// @Description Delete a booking
// @Tags booking
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteBooking(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.storage.Booking().Delete(id)
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
