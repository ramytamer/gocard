package actions

import "net/http"

func (as *ActionSuite) Test_Healthcheck_Handler() {
	res := as.JSON("/status").Get()

	as.Equal(http.StatusOK, res.Code)
	as.JSONEq(`{"status": "ok"}`, res.Body.String())
}

func (as *ActionSuite) Test_Healthcheck_Unhealthy() {
	res := as.JSON("/status?unhealthy=false").Get()

	as.Equal(http.StatusOK, res.Code)
	as.JSONEq(`{"status": "ok"}`, res.Body.String())

	res = as.JSON("/status?unhealthy=true").Get()

	as.Equal(http.StatusInternalServerError, res.Code)
	as.JSONEq(`{"status": "unhealthy"}`, res.Body.String())
}
