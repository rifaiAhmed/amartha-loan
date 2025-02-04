package api

import (
	"amartha-loan/constants"
	"amartha-loan/helpers"
	"amartha-loan/internal/input"
	"amartha-loan/internal/interfaces"
	"amartha-loan/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoanAPI struct {
	LoanService interfaces.ILoanService
}

func (api *LoanAPI) Create(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	req := models.Loan{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if req.UserID == 0 {
		log.Error("user_id is empty")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	err := api.LoanService.Create(c.Request.Context(), &req)
	if err != nil {
		log.Error("failed to create loan: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, req)
}

func (api *LoanAPI) GetLoanByID(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	id := c.Query("id")
	ID, err := strconv.Atoi(id)
	if err != nil || ID == 0 {
		log.Error("failed to get id")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	resp, err := api.LoanService.GetLoanByID(c, ID)
	if err != nil {
		log.Error("failed to get loan: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)
}

func (api *LoanAPI) IsDelinquent(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	id := c.Query("id")
	ID, err := strconv.Atoi(id)
	if err != nil || ID == 0 {
		log.Error("failed to get id")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	resp, err := api.LoanService.IsDelinquent(c, ID)
	if err != nil {
		log.Error("failed to get loan: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)
}

func (api *LoanAPI) MakePayment(c *gin.Context) {
	var (
		log   = helpers.Logger
		input = input.DataURI{}
	)
	err := c.ShouldBindUri(&input)
	if err != nil {
		log.Error("failed to get id")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	err = api.LoanService.MakePayment(c, input.ID)
	if err != nil {
		log.Error("failed to get loan: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, nil)
}
