package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/Alv1ol/Todoapp/internal/core/logger"
	core_http_response "github.com/Alv1ol/Todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/Alv1ol/Todoapp/internal/core/transport/http/utils"
)

type GetTaskResponce TaskDTOResponce


// GetTask godoc
// @Summary Get task
// @Description Get task by id
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "task id"
// @Success 200 {object} GetTaskResponce
// @Failure 400 {object} core_http_response.ErrorResponse "bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "task not found"
// @Failure 500 {object} core_http_response.ErrorResponse "internal server error"
// @Router /tasks/{id} [get]
func (h *TasksHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responceHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responceHandler.ErrorResponse(
			err,
			"failde to get taskID path value",
		)
		return
	}

	taskDomain, err := h.tasksService.GetTask(ctx,taskID)
	if err != nil {
		responceHandler.ErrorResponse(
			err,
			"failed to get task",
		)
		return
	}

	responce := GetTaskResponce(taskDTOFromDomain(taskDomain))

	responceHandler.JSONResponce(responce, http.StatusOK)
}