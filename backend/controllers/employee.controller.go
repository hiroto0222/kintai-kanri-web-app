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

type EmployeeController struct {
	config     config.Config
	store      db.Store
	TokenMaker token.Maker
}

func NewEmployeeController(config config.Config, store db.Store, tokenMaker token.Maker) *EmployeeController {
	return &EmployeeController{
		config:     config,
		store:      store,
		TokenMaker: tokenMaker,
	}
}

type getEmployeeRequest struct {
	ID uuid.UUID `uri:"id" binding:"required,min=1"`
}

// GetEmployee: api/employees/:id 従業員情報取得
func (c *EmployeeController) GetEmployee(ctx *gin.Context) {
	var req getEmployeeRequest
	// ShouldBindUri はリクエストのURIからパラメータを取得
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	employee, err := c.store.GetEmployeeById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if employee.Email != authPayload.Email {
		err := errors.New("you do not have permission")
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, employee)
}
