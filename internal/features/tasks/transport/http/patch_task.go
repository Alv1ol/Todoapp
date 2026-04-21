package tasks_transport_http

import (
	"fmt"
	"net/http"

	"github.com/Alv1ol/Todoapp/internal/core/domain"
	core_logger "github.com/Alv1ol/Todoapp/internal/core/logger"
	core_http_request "github.com/Alv1ol/Todoapp/internal/core/transport/http/request"
	core_http_response "github.com/Alv1ol/Todoapp/internal/core/transport/http/response"
	core_http_types "github.com/Alv1ol/Todoapp/internal/core/transport/http/types"
	core_http_utils "github.com/Alv1ol/Todoapp/internal/core/transport/http/utils"
)

type PatchTaskrequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed   core_http_types.Nullable[bool]   `json:"completed"`
}

func (r *PatchTaskrequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("`Title` can't be null")
		}

		titleLen := len([]rune(*r.Title.Value))
		if titleLen < 1 || titleLen > 100 {
			return fmt.Errorf("`Title` must be between 1 and 100 symbols")
		}
	}

	if r.Description.Set {
		if r.Description.Value != nil {
			descriptionLen := len([]rune(*r.Description.Value))
			if descriptionLen < 1 || descriptionLen > 1000 {
				return fmt.Errorf("`Description` must be between 1 and 1000 symbols")
			}
		}
	}

	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("`Completed` can't be null")
		}
	}

	return nil
}

type PatchTaskResponce TaskDTOResponce

// PatchTask godoc
// @Summary Patch task
// @Description Patch task by id. All fields are optional, but at least one must be provided. Title must be between 1 and 100 symbols. Description must be between 1 and 1000 symbols.
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "task id"
// @Param request body PatchTaskrequest true "patch task request"
// @Success 200 {object} PatchTaskResponce
// @Failure 400 {object} core_http_response.ErrorResponse "bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "task not found"
// @Failure 500 {object} core_http_response.ErrorResponse "internal server error"
// @Router /tasks/{id} [patch]
func (h *TasksHTTPHandler) PatchTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responceHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_utils.GetIntPathValue(
		r,
		"id",
	)
	if err != nil {
		responceHandler.ErrorResponse(
			err,
			"failed to get taskID path value",
		)
		return
	}

	var request PatchTaskrequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responceHandler.ErrorResponse(
			err,
			"failed to decode http request",
		)
		return
	}

	taskPatch := taskPatchFromRequest(request)

	taskDomain, err := h.tasksService.PatchTask(ctx, taskID, taskPatch)
	if err != nil {
		responceHandler.ErrorResponse(
			err,
			"failed to patch task",
		)
		return
	}

	responce := PatchTaskResponce(taskDTOFromDomain(taskDomain))
	responceHandler.JSONResponce(responce, http.StatusOK)
}

func taskPatchFromRequest(request PatchTaskrequest) domain.TaskPatch {
	return domain.NewTaskPatch(
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Completed.ToDomain(),
	)
}
