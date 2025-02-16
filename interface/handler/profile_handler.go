package handler

import (
	"encoding/json"
	"errors"
	custumError "github.com/yusuke-takatsu/fishing-api-server/errors"
	dto "github.com/yusuke-takatsu/fishing-api-server/interface/dto/input/profile"
	"github.com/yusuke-takatsu/fishing-api-server/usecase/profile"
	"mime/multipart"
	"net/http"
)

type ProfileHandler struct {
	registerUseCase *profile.RegisterUseCase
}

func NewProfileHandler(registerUseCase *profile.RegisterUseCase) *ProfileHandler {
	return &ProfileHandler{registerUseCase: registerUseCase}
}

type RegisterRequest struct {
	NickName           string                `form:"nick_name" validate:"required,max=255"`
	DateOfBirth        string                `form:"date_of_birth" validate:"required,datetime=2006-01-02"`
	FishingStartedDate string                `form:"fishing_started_date" validate:"required,datetime=2006-01-02"`
	Image              *multipart.FileHeader `form:"image" validate:"omitempty"`
}

func (h *ProfileHandler) Register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		custumError.Handler(w, err)
		return
	}

	req := RegisterRequest{
		NickName:           r.FormValue("nick_name"),
		DateOfBirth:        r.FormValue("date_of_birth"),
		FishingStartedDate: r.FormValue("fishing_started_date"),
	}

	var fileHeader *multipart.FileHeader
	file, header, err := r.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		custumError.Handler(w, err)
		return
	}

	defer file.Close()

	fileHeader = header
	req.Image = fileHeader

	err = h.registerUseCase.Execute(r.Context(), dto.RegisterInputData{
		NickName:           req.NickName,
		DateOfBirth:        req.DateOfBirth,
		FishingStartedDate: req.FishingStartedDate,
		Image:              req.Image,
	})

	if err != nil {
		custumError.Handler(w, err)
		return
	}

	response := map[string]string{
		"message": "プロフィールが更新されました。",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		custumError.Handler(w, err)
	}
}
