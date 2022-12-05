package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mirasildev/exam_project_2.0/api/models"
	"github.com/mirasildev/exam_project_2.0/storage/repo"
)

// @Security ApiKeyAuth

// @Router /rooms [post]
// @Summary Create a room
// @Description Create a room
// @Tags room
// @Accept json
// @Produce json
// @Param room body models.CreateRoomRequest true "room"
// @Success 200 {object} models.CreateRoomResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateRoom(c *gin.Context) {
	var (
		req models.CreateRoomRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	id, err := h.storage.Room().Create(&repo.Room{
		RoomNum:       req.RoomNum,
		Type:          req.Type,
		Description:   req.Description,
		HotelID:       req.HotelID,
		PricePerNight: req.PricePerNight,
		Status:        req.Status,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, id)
}

func parseRoomModel(room *repo.Room) models.Room {
	return models.Room{
		ID:            room.ID,
		RoomNum:       room.RoomNum,
		Type:          room.Type,
		Description:   room.Description,
		HotelID:       room.HotelID,
		PricePerNight: room.PricePerNight,
		Status:        room.Status,
	}
}

// @Router /rooms/{id} [get]
// @Summary Get rooms by id
// @Description Get rooms by id
// @Tags room
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Room
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetRoom(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Room().Get(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	room := parseRoomModel(resp)

	c.JSON(http.StatusOK, room)
}

// @Router /rooms [get]
// @Summary Get all rooms
// @Description Get all rooms
// @Tags room
// @Accept json
// @Produce json
// @Param filter query models.GetAllRoomsParams false "Filter"
// @Success 200 {object} models.GetAllRoomsResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllRooms(c *gin.Context) {
	req, err := validateGetAllRoomsParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Room().GetAll(&repo.GetAllRoomsParams{
		Page:    req.Page,
		Limit:   req.Limit,
		HotelID: req.HotelID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getRoomsResponse(result))
}

func validateGetAllRoomsParams(c *gin.Context) (*models.GetAllRoomsParams, error) {
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

	return &models.GetAllRoomsParams{
		Limit: int32(limit),
		Page:  int32(page),
	}, nil
}

func getRoomsResponse(data *repo.GetAllRoomsResult) *models.GetAllRoomsResponse {
	response := models.GetAllRoomsResponse{
		Rooms: make([]*models.Room, 0),
		Count: data.Count,
	}

	for _, room := range data.Rooms {
		p := parseRoomModel(room)
		response.Rooms = append(response.Rooms, &p)
	}

	return &response
}

// @Router /rooms/{id} [put]
// @Summary Update a room
// @Description Update a room
// @Tags room
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param room body models.UpdateRoomRequest true "Room"
// @Success 200 {object} models.Room
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateRoom(c *gin.Context) {
	var (
		req models.UpdateRoomRequest
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

	updated, err := h.storage.Room().Update(&repo.Room{
		ID:            id,
		RoomNum:       req.RoomNum,
		Type:          req.Type,
		Description:   req.Description,
		HotelID:       req.HotelID,
		PricePerNight: req.PricePerNight,
		Status:        req.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// @Router /rooms/{id} [delete]
// @Summary Delete a rooms
// @Description Delete a rooms
// @Tags room
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteRoom(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.storage.Room().Delete(id)
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
