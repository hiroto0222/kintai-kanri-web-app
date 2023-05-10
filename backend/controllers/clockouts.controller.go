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

type ClockOutController struct {
	config     config.Config
	TokenMaker token.Maker
	store      db.Store
}

func NewClockOutController(config config.Config, store db.Store, tokenMaker token.Maker) *ClockOutController {
	return &ClockOutController{
		config:     config,
		store:      store,
		TokenMaker: tokenMaker,
	}
}

type reqCreateClockOut struct {
	EmployeeID string `json:"employee_id" binding:"required"`
}

// CreateClockOut godoc
// @Summary      退出打刻
// @Tags         clockouts
// @Accept       json
// @Produce      json
// @Security		 BearerAuth
// @Param request body reqCreateClockOut true "Request body"
// @Success      201  {object}   db.ClockOutTxResult
// @Failure      400  {object}   utils.ErrorResponse
// @Failure      401  {object}   utils.ErrorResponse
// @Failure      500  {object}   utils.ErrorResponse
// @Router       /clockouts [post]
func (c *ClockOutController) CreateClockOut(ctx *gin.Context) {
	var req reqCreateClockOut

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if authPayload.EmployeeID != req.EmployeeID {
		err := errors.New("you do not have permission")
		ctx.JSON(http.StatusUnauthorized, utils.CreateErrorResponse(err))
		return
	}

	reqEmployeeID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorResponse(err))
		return
	}

	prevClockIn, err := c.store.GetMostRecentClockIn(ctx, reqEmployeeID)
	if err != nil && err != sql.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorResponse(err))
		return
	}

	// 出勤打刻していない場合はエラー
	if prevClockIn == (db.ClockIn{}) || prevClockIn.ClockedOut {
		err := errors.New("you have not clocked in yet")
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorResponse(err))
		return
	}

	arg := db.ClockOutTxParams{
		EmployeeID: reqEmployeeID,
		ClockInID:  prevClockIn.ID,
	}

	// clockout transaction
	result, err := c.store.ClockOutTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
