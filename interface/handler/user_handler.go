package handler

import (
	"context"
	"encoding/json"
	"github.com/yusuke-takatsu/fishing-api-server/errors"
	dto "github.com/yusuke-takatsu/fishing-api-server/interface/dto/input/user"
	"github.com/yusuke-takatsu/fishing-api-server/usecase/user"
	"net/http"
)

type SessionManager interface {
	RegenerateSession(ctx context.Context, w http.ResponseWriter, userID string) error
}

type UserHandler struct {
	loginUseCase   *user.LoginUseCase
	sessionManager SessionManager
}

func NewUserHandler(loginUseCase *user.LoginUseCase, sessionManager SessionManager) *UserHandler {
	return &UserHandler{
		loginUseCase:   loginUseCase,
		sessionManager: sessionManager,
	}
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,password"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.Handler(w, err)
		return
	}

	input := dto.LoginInputData{
		Email:    req.Email,
		Password: req.Password,
	}

	id, err := h.loginUseCase.Execute(r.Context(), input)
	if err != nil {
		errors.Handler(w, err)
		return
	}

	if err := h.sessionManager.RegenerateSession(r.Context(), w, id); err != nil {
		errors.Handler(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&LoginResponse{Message: "ログインしました。"}); err != nil {
		errors.Handler(w, err)
	}
}
