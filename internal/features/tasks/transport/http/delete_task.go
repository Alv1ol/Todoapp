package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/Alv1ol/Todoapp/internal/core/logger"
	core_http_response "github.com/Alv1ol/Todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/Alv1ol/Todoapp/internal/core/transport/http/utils"
)

// DeleteTask godoc
// @Summary Delete task
// @Description Delete task by id
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "task id"
// @Success 204 "no content"
// @Failure 400 {object} core_http_response.ErrorResponse "bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "task not found"
// @Failure 500 {object} core_http_response.ErrorResponse "internal server error"
// @Router /tasks/{id} [delete]
func (h *TasksHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responcehandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil { 
		responcehandler.ErrorResponse(
			err,
			"failed to get taskID path value",
		)
		return
	}

	if err := h.tasksService.DeleteTask(ctx, taskID); err != nil {
		responcehandler.ErrorResponse(
			err,
			"failed to delete task",
		)
		return
	}

	responcehandler.NoContentresponce()
}