package actions

import (
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
)

// HealthcheckHandler default implementation.
func HealthcheckHandler(c buffalo.Context) error {
	status := "ok"
	code := http.StatusOK

	if strings.ToLower(c.Param("unhealthy")) == "true" {
		status = "unhealthy"
		code = http.StatusInternalServerError
	}

	res := map[string]string{"status": status}

	return c.Render(code, r.JSON(res))
}
