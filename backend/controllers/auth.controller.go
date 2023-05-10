package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hiroto0222/kintai-kanri-web-app/config"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/middlewares"
	"github.com/hiroto0222/kintai-kanri-web-app/token"
	"github.com/hiroto0222/kintai-kanri-web-app/utils"
)

// TODO: Refactor
type AuthController struct {
	config     config.Config
	TokenMaker token.Maker
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

// RegisterAccount godoc
// @Summary      従業員登録（管理者 Admin Only）
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param request body registerEmployeeRequest true "Request body"
// @Success      201  {object}  EmployeeResponse
// @Failure      400  {object}   utils.ErrorResponse
// @Failure      401  {object}   utils.ErrorResponse
// @Failure      500  {object}   utils.ErrorResponse
// @Router       /auth/register [post]
func (ac *AuthController) RegisterEmployee(ctx *gin.Context) {
	var req registerEmployeeRequest

	// application/jsonでレスポンスを返したいため、ShouldBindJSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorResponse(err))
		return
	}

	// 取得するユーザーがログインしているユーザー・管理者でない場合はエラーを返す
	authPayload := ctx.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if !authPayload.IsAdmin {
		err := errors.New("you do not have permission")
		ctx.JSON(http.StatusUnauthorized, utils.CreateErrorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorResponse(err))
		return
	}

	response := NewUserResponse(employee)
	ctx.JSON(http.StatusCreated, response)
}

type logInEmployeeRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type logInEmployeeResponse struct {
	SessionID             uuid.UUID        `json:"session_id"`
	AccessToken           string           `json:"access_token"`
	AccessTokenExpiresAt  time.Time        `json:"access_token_expires_at"`
	RefreshToken          string           `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time        `json:"refresh_token_expires_at"`
	User                  EmployeeResponse `json:"user"`
}

// LogInEmployee godoc
// @Summary      従業員認証
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param request body logInEmployeeRequest true "Request body"
// @Success      200  {object}  logInEmployeeResponse
// @Failure      400  {object}   utils.ErrorResponse
// @Failure      401  {object}   utils.ErrorResponse
// @Failure      404  {object}   utils.ErrorResponse
// @Failure      500  {object}   utils.ErrorResponse
// @Router       /auth/login [post]
func (ac *AuthController) LogInEmployee(ctx *gin.Context) {
	var req logInEmployeeRequest

	// application/jsonでレスポンスを返したいため、ShouldBindJSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorResponse(err))
		return
	}

	// ユーザーを取得
	employee, err := ac.store.GetEmployeeByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// ユーザーが存在しない場合404エラーを返す
			ctx.JSON(http.StatusNotFound, utils.CreateErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorResponse(err))
		return
	}

	err = utils.CheckPassword(req.Password, employee.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.CreateErrorResponse(err))
		return
	}

	// アクセストークンを作成
	accessToken, accessPayload, err := ac.TokenMaker.CreateToken(employee.ID.String(), employee.IsAdmin, ac.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorResponse(err))
		return
	}

	// リフレッシュトークンを作成
	refreshToken, refreshPayload, err := ac.TokenMaker.CreateToken(employee.ID.String(), employee.IsAdmin, ac.config.RefreshTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorResponse(err))
		return
	}

	session, err := ac.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Email:        employee.Email,
		EmployeeID:   employee.ID,
		RefreshToken: refreshToken,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorResponse(err))
		return
	}

	response := logInEmployeeResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  NewUserResponse(employee),
	}

	ctx.JSON(http.StatusOK, response)
}

type refreshAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type refreshAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

// RefreshAccessToken godoc
// @Summary      トークンのリフレッシュ
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param request body refreshAccessTokenRequest true "Request body"
// @Success      200  {object}  refreshAccessTokenResponse
// @Failure      400  {object}   utils.ErrorResponse
// @Failure      401  {object}   utils.ErrorResponse
// @Failure      404  {object}   utils.ErrorResponse
// @Failure      500  {object}   utils.ErrorResponse
// @Router       /auth/refresh [post]
func (ac *AuthController) RefreshAccessToken(ctx *gin.Context) {
	var req refreshAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CreateErrorResponse(err))
		return
	}

	// リフレッシュトークンを検証
	refreshPayload, err := ac.TokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.CreateErrorResponse(err))
		return
	}

	// セッションを取得
	session, err := ac.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.CreateErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorResponse(err))
		return
	}

	// セッションがブロックされているか確認
	if session.IsBlocked {
		err := fmt.Errorf("blocked session")
		ctx.JSON(http.StatusUnauthorized, utils.CreateErrorResponse(err))
		return
	}

	// セッションのユーザーとリフレッシュトークンのユーザーが一致するか確認
	if session.EmployeeID.String() != refreshPayload.EmployeeID {
		err := fmt.Errorf("incorrect session user")
		ctx.JSON(http.StatusUnauthorized, utils.CreateErrorResponse(err))
		return
	}

	// セッションのリフレッシュトークンとリクエストのリフレッシュトークンが一致するか確認
	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("incorrect session refresh token")
		ctx.JSON(http.StatusUnauthorized, utils.CreateErrorResponse(err))
		return
	}

	// リフレッシュトークンが有効期限切れか確認
	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("refresh token expired")
		ctx.JSON(http.StatusUnauthorized, utils.CreateErrorResponse(err))
		return
	}

	accessToken, accessPayload, err := ac.TokenMaker.CreateToken(refreshPayload.EmployeeID, refreshPayload.IsAdmin, ac.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.CreateErrorResponse(err))
		return
	}

	response := refreshAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}

	ctx.JSON(http.StatusOK, response)
}

type EmployeeResponse struct {
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

func NewUserResponse(employee db.Employee) EmployeeResponse {
	return EmployeeResponse{
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
