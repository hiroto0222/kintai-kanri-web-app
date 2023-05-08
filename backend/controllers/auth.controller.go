package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hiroto0222/kintai-kanri-web-app/config"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/token"
	"github.com/hiroto0222/kintai-kanri-web-app/utils"
)

// TODO: Refactor
type AuthController struct {
	config     config.Config
	tokenMaker token.Maker
	store      db.Store
}

func NewAuthController(config config.Config, store db.Store, tokenMaker token.Maker) *AuthController {
	return &AuthController{config, tokenMaker, store}
}

type registerEmployeeRequest struct {
	FirstName string        `json:"first_name" binding:"required"`
	LastName  string        `json:"last_name" binding:"required"`
	Email     string        `json:"email" binding:"required,email"`
	Phone     string        `json:"phone" binding:"required"`
	Address   string        `json:"address" binding:"required"`
	RoleID    sql.NullInt32 `json:"role_id"`
	IsAdmin   bool          `json:"is_admin"`
	Password  string        `json:"password" binding:"required,min=6"`
}

// RegisterEmployee: api/auth/register ユーザー登録
func (ac *AuthController) RegisterEmployee(ctx *gin.Context) {
	var req registerEmployeeRequest

	// application/jsonでレスポンスを返したいため、ShouldBindJSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	response := newUserResponse(employee)
	ctx.JSON(http.StatusCreated, response)
}

type logInEmployeeRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type logInEmployeeResponse struct {
	AccessToken string           `json:"access_token"`
	User        employeeResponse `json:"user"`
}

// LogInEmployee: api/auth/login ユーザー認証
func (ac *AuthController) LogInEmployee(ctx *gin.Context) {
	var req logInEmployeeRequest

	// application/jsonでレスポンスを返したいため、ShouldBindJSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	// ユーザーを取得
	employee, err := ac.store.GetEmployeeByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// ユーザーが存在しない場合404エラーを返す
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	err = utils.CheckPassword(req.Password, employee.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	// アクセストークンを作成
	accessToken, err := ac.tokenMaker.CreateToken(employee.Email, ac.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	response := logInEmployeeResponse{
		AccessToken: accessToken,
		User:        newUserResponse(employee),
	}
	ctx.JSON(http.StatusOK, response)
}

type employeeResponse struct {
	ID        uuid.UUID     `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Email     string        `json:"email"`
	Phone     string        `json:"phone"`
	Address   string        `json:"address"`
	RoleID    sql.NullInt32 `json:"role_id"`
	IsAdmin   bool          `json:"is_admin"`
	CreatedAt time.Time     `json:"created_at"`
}

func newUserResponse(employee db.Employee) employeeResponse {
	return employeeResponse{
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
