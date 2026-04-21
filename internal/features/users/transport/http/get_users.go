package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Alv1ol/Todoapp/internal/core/logger"
	core_http_response "github.com/Alv1ol/Todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/Alv1ol/Todoapp/internal/core/transport/http/utils"
)

type GetuserResponse []UserDTOResponce

// GetUsers godoc
// @Summary Get users
// @Description Get users with pagination support using 'limit' and 'offset' query parameters
// @Tags users
// @Accept json
// @Produce json
// @Param limit query int false "number of users to return"
// @Param offset query int false "number of users to skip"
// @Success 200 {object} GetuserResponse
// @Failure 400 {object} core_http_response.ErrorResponse "bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "internal server error"
// @Router /users [get]
func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'limit'/'offset' query param",
		)
		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get users",
		)
		return
	}

	responce := GetuserResponse(usersDTOFromDomains(userDomains))

	responseHandler.JSONResponce(responce, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return limit, offset, nil
}
