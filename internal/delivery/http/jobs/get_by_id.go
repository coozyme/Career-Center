package jobs

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *JobsHandler) GetJobById(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = context.TODO()
		vars  = mux.Vars(r)
		jobId = vars["job_id"]
		log   = logger.NewLogger(fmt.Sprintf("/v1/job-detail/%s", jobId))
	)

	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		log.General("", errToken)
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	jobs, err := h.UCJobs.GetJobById(ctx, user.Email, jobId)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.General("", err)
		return
	}

	jobsResponse := response.GetDetailJobResponse(jobs)
	helper.ResponseInterface(w, "success get detail job", jobsResponse, http.StatusInternalServerError)
	log.General("success Get detail job", jobsResponse)
	return
}
