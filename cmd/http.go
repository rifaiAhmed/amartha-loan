package cmd

import (
	"amartha-loan/external"
	"amartha-loan/helpers"
	"amartha-loan/internal/api"
	"amartha-loan/internal/interfaces"
	"amartha-loan/internal/repository"
	"amartha-loan/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()

	r := gin.Default()

	r.GET("/health", d.HealthcheckAPI.HealthcheckHandlerHTTP)

	loanV1 := r.Group("/loan/v1")
	loanV1.POST("/", d.MiddlewareValidateToken, d.LoanAPI.Create)
	loanV1.GET("/get", d.MiddlewareValidateToken, d.LoanAPI.GetLoanByID)
	loanV1.GET("/is-delinquent", d.MiddlewareValidateToken, d.LoanAPI.IsDelinquent)
	loanV1.PATCH("/make-payment/:id", d.MiddlewareValidateToken, d.LoanAPI.MakePayment)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	HealthcheckAPI interfaces.IHealthcheckAPI
	LoanAPI        interfaces.ILoanAPI
	External       interfaces.IExternal
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	loanRepo := &repository.LoanRepo{
		DB: helpers.DB,
	}

	loanSvc := &services.LoanService{
		LoanRepo: loanRepo,
	}
	LoanAPI := &api.LoanAPI{
		LoanService: loanSvc,
	}

	external := &external.External{}

	return Dependency{
		HealthcheckAPI: healthcheckAPI,
		LoanAPI:        LoanAPI,
		External:       external,
	}
}
