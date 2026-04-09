package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Alv1ol/Todoapp/internal/core/domain"
	core_logger "github.com/Alv1ol/Todoapp/internal/core/logger"
	core_http_response "github.com/Alv1ol/Todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/Alv1ol/Todoapp/internal/core/transport/http/utils"
)

type GetStatisticsResponce struct {
	TaskCreated               int
	TaskCompleted             int
	TaskCompletedRate         *float64
	TaskAverageComplitionTime *string
}

func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responcehandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, from, to, err := getUserIDFromtoQueryParams(r)
	if err != nil {
		responcehandler.ErrorResponse(
			err,
			"failed to get id/from/to from query params",
		)
		return
	}

	statistics, err := h.statisticsService.GetStatistics(ctx, userID, from, to)
	if err != nil {
		responcehandler.ErrorResponse(err, "failed to get statistics")
		return
	}

	responce := toDTOFromDomain(statistics)

	responcehandler.JSONResponce(responce, http.StatusOK)
}

func toDTOFromDomain(statistics domain.Statistics) GetStatisticsResponce {
	var avgTime *string
	if statistics.TaskAverageComplitionTime != nil {
		duration := statistics.TaskAverageComplitionTime.String()
		avgTime = &duration
	}
	return GetStatisticsResponce{
		TaskCreated:               statistics.TaskCreated,
		TaskCompleted:             statistics.TaskCompleted,
		TaskCompletedRate:         statistics.TaskCompletedRate,
		TaskAverageComplitionTime: avgTime,
	}
}

func getUserIDFromtoQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		userIDQueryParamkey = "user_id"
		fromQueryParamkey   = "from"
		toQueryParamkey     = "to"
	)

	userID, err := core_http_utils.GetIntQueryParam(r, userIDQueryParamkey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'userID' query param: %w", err)
	}

	from, err := core_http_utils.GetDateQueryParam(r, fromQueryParamkey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Get 'from' query param: %w", err)
	}

	to, err := core_http_utils.GetDateQueryParam(r, toQueryParamkey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Get 'to' query param: %w", err)
	}

	return userID, from, to, nil
}
