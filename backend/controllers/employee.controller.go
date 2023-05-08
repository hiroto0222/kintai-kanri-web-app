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

// GetEmployee: api/employees/:id 従業員情報取得 (権限: Admin, 自分)
func (c *EmployeeController) GetEmployee(ctx *gin.Context) {
	var req getEmployeeRequest
	// ShouldBindUri はリクエストのURIからパラメータを取得
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	// 文字列で渡された employee_id をUUIDに変換
	employee_id, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	// employee_id から従業員情報を取得
	employee, err := c.store.GetEmployeeById(ctx, employee_id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)

	// req ユーザーが Admin か確認する
	reqEmployeeId, err := uuid.Parse(authPayload.EmployeeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	reqEmployee, err := c.store.GetEmployeeById(ctx, reqEmployeeId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	// 取得するユーザーがログインしているユーザー・管理者でない場合はエラーを返す
	if employee.ID.String() != authPayload.EmployeeID && !reqEmployee.IsAdmin {
		err := errors.New("you do not have permission")
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	response := NewUserResponse(employee)

	ctx.JSON(http.StatusOK, response)
}

type listEmployeesRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// GetEmployees: api/employees 全従業員情報を取得 (権限: Admin)
func (c *EmployeeController) ListEmployees(ctx *gin.Context) {
	var req listEmployeesRequest
	// ShouldBindQuery はリクエストのクエリパラメータを取得
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)

	// req ユーザーが Admin か確認する
	reqEmployeeId, err := uuid.Parse(authPayload.EmployeeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	reqEmployee, err := c.store.GetEmployeeById(ctx, reqEmployeeId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	if !reqEmployee.IsAdmin {
		err := errors.New("you do not have permission")
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	arg := db.ListEmployeesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	employees, err := c.store.ListEmployees(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	var response []EmployeeResponse
	for _, employee := range employees {
		response = append(response, NewUserResponse(employee))
	}

	ctx.JSON(http.StatusOK, response)
}
