package handler

import (
	"encoding/json"
	"github.com/yusuke-takatsu/fishing-api-server/errors"
	dto "github.com/yusuke-takatsu/fishing-api-server/interface/dto/input/user"
	"github.com/yusuke-takatsu/fishing-api-server/usecase/user"
	"net/http"
)

type UserHandler struct {
	loginUseCase *user.LoginUseCase
}

func NewUserHandler(loginUseCase *user.LoginUseCase) *UserHandler {
	return &UserHandler{loginUseCase: loginUseCase}
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

	if err := h.loginUseCase.Execute(r.Context(), input); err != nil {
		errors.Handler(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&LoginResponse{Message: "ログインしました。"}); err != nil {
		errors.Handler(w, err)
	}
}
