package users_transport_http

import (
	"net/http"

	core_logger "github.com/Alv1ol/Todoapp/internal/core/logger"
	core_http_response "github.com/Alv1ol/Todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/Alv1ol/Todoapp/internal/core/transport/http/utils"
)

type GetUserResponse UserDTOResponce

// GetUser godoc
// @Summary Get user
// @Description Get user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} core_http_response.ErrorResponse "bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "user not found"
// @Failure 500 {object} core_http_response.ErrorResponse "internal server error"
// @Router /users/{id} [get]
func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responceHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil { 
		responceHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)

		return
	}

	user, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responceHandler.ErrorResponse(
			err,
			"failed to get user",
		)

		return
	}

	responce := GetUserResponse(userDTOFromDomain(user))

	responceHandler.JSONResponce(responce, http.StatusOK)
}