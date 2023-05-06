package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/utils"
)

type AuthController struct {
	store db.Store
}

func NewAuthController(store db.Store) *AuthController {
	return &AuthController{store}
}

type signUpEmployeeRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Address   string `json:"address" binding:"required"`
	RoleID    int32  `json:"role_id" binding:"required"`
	IsAdmin   bool   `json:"is_admin"`
	Password  string `json:"password" binding:"required,min=6"`
}

type signUpEmployeeResponse struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	RoleID    int32     `json:"role_id"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
}

// SignUpEmployee: api/auth/register ユーザー登録
func (ac *AuthController) SignUpEmployee(ctx *gin.Context) {
	var req signUpEmployeeRequest

	// application/jsonでレスポンスを返したいため、ShouldBindJSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	args := db.CreateEmployeeParams{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		Address:        req.Address,
		RoleID:         req.RoleID,
		IsAdmin:        req.IsAdmin,
		HashedPassword: hashedPassword,
	}

	employee, err := ac.store.CreateEmployee(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := newEmployeeResponse(employee)
	ctx.JSON(http.StatusCreated, response)
}

type signInEmployee struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// SignInEmployee: api/auth/login ユーザー認証
func (ac *AuthController) SignInEmployee(ctx *gin.Context) {
	var req signInEmployee

	// application/jsonでレスポンスを返したいため、ShouldBindJSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// ユーザーを取得
	employee, err := ac.store.GetEmployeeByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// ユーザーが存在しない場合404エラーを返す
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = utils.CheckPassword(req.Password, employee.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	// TODO: Access token

	response := newEmployeeResponse(employee)
	ctx.JSON(http.StatusOK, response)
}

func newEmployeeResponse(employee db.Employee) signUpEmployeeResponse {
	return signUpEmployeeResponse{
		ID:        employee.ID,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Email:     employee.Email,
		Phone:     employee.Phone,
		Address:   employee.Address,
		RoleID:    employee.RoleID,
		IsAdmin:   employee.IsAdmin,
		CreatedAt: employee.CreatedAt,
	}
}
