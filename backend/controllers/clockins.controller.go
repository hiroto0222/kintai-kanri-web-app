package controllers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hiroto0222/kintai-kanri-web-app/config"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/middlewares"
	"github.com/hiroto0222/kintai-kanri-web-app/token"
	"github.com/hiroto0222/kintai-kanri-web-app/utils"
)

type ClockInController struct {
	config     config.Config
	TokenMaker token.Maker
	store      db.Store
}

func NewClockInController(config config.Config, store db.Store, tokenMaker token.Maker) *ClockInController {
	return &ClockInController{
		config:     config,
		store:      store,
		TokenMaker: tokenMaker,
	}
}

type reqCreateClockIn struct {
	EmployeeID string `json:"employee_id" binding:"required"`
}

// POST: api/clockins 出勤打刻
func (c *ClockInController) CreateClockIn(ctx *gin.Context) {
	var req reqCreateClockIn

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if authPayload.EmployeeID != req.EmployeeID {
		err := errors.New("you do not have permission")
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	reqEmployeeID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	prevClockIn, err := c.store.GetMostRecentClockIn(ctx, reqEmployeeID)
	if err != nil && err != sql.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	// 退出打刻しないまま出勤打刻している場合はエラー
	if prevClockIn != (db.ClockIn{}) && !prevClockIn.ClockedOut {
		err := errors.New("you have not clocked out yet")
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	clockIn, err := c.store.CreateClockIn(ctx, reqEmployeeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, clockIn)
}
