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
	ID string `uri:"id" binding:"required,min=1"`
}

// GetEmployee godoc
// @Summary      従業員情報を取得
// @Tags         employees
// @Accept       json
// @Produce      json
// @Security		 BearerAuth
// @Param id path string true "ID of the employee to retrieve"
// @Success      200  {object}   EmployeeResponse
// @Failure      400  {object}   utils.ErrorResponse
// @Failure      401  {object}   utils.ErrorResponse
// @Failure      404  {object}   utils.ErrorResponse
// @Failure      500  {object}   utils.ErrorResponse
// @Router       /employees/:id [get]
func (c *EmployeeController) GetEmployee(ctx *gin.Context) {
	var req getEmployeeRequest
	// ShouldBindUri はリクエストのURIからパラメータを取得
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorResponse(err))
		return
	}

	// 文字列で渡された employee_id をUUIDに変換
	employee_id, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorResponse(err))
		return
	}

	// 取得するユーザーがログインしているユーザー・管理者でない場合はエラーを返す
	authPayload := ctx.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if employee_id.String() != authPayload.EmployeeID && !authPayload.IsAdmin {
		err := errors.New("you do not have permission")
		ctx.JSON(http.StatusUnauthorized, utils.CreateErrorResponse(err))
		return
	}

	// employee_id から従業員情報を取得
	employee, err := c.store.GetEmployeeById(ctx, employee_id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.CreateErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorResponse(err))
		return
	}

	response := NewUserResponse(employee)

	ctx.JSON(http.StatusOK, response)
}

type listEmployeesRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5"`
}

func (c *EmployeeController) ListEmployees(ctx *gin.Context) {
	var req listEmployeesRequest
	// ShouldBindQuery はリクエストのクエリパラメータを取得
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorResponse(err))
		return
	}

	// Admin でない場合はエラーを返す
	authPayload := ctx.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if !authPayload.IsAdmin {
		err := errors.New("you do not have permission")
		ctx.JSON(http.StatusUnauthorized, utils.CreateErrorResponse(err))
		return
	}

	arg := db.ListEmployeesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	employees, err := c.store.ListEmployees(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorResponse(err))
		return
	}

	var response []EmployeeResponse
	for _, employee := range employees {
		response = append(response, NewUserResponse(employee))
	}

	ctx.JSON(http.StatusOK, response)
}
