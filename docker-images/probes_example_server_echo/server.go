package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var LivenessProbeOK = true
var ReadinessProbeOK = true

func server() error {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	type baseResponse struct {
		Status   int    `json:"status"`
		Hostname string `json:"hostname"`
	}

	type response struct {
		Status    int    `json:"status"`
		Hostname  string `json:"hostname"`
		Liveness  bool   `json:"liveness"`
		Readiness bool   `json:"readiness"`
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response{
			Status:    http.StatusOK,
			Hostname:  hostname,
			Liveness:  LivenessProbeOK,
			Readiness: ReadinessProbeOK,
		})
	})

	e.GET("/livez", func(c echo.Context) error {
		if LivenessProbeOK {
			return c.JSON(http.StatusOK, baseResponse{
				Status:   http.StatusOK,
				Hostname: hostname,
			})
		}
		return c.JSON(http.StatusInternalServerError, baseResponse{
			Status:   http.StatusInternalServerError,
			Hostname: hostname,
		})
	})

	e.GET("/readyz", func(c echo.Context) error {
		if ReadinessProbeOK {
			return c.JSON(http.StatusOK, baseResponse{
				Status:   http.StatusOK,
				Hostname: hostname,
			})
		}
		return c.JSON(http.StatusInternalServerError, baseResponse{
			Status:   http.StatusInternalServerError,
			Hostname: hostname,
		})
	})

	e.GET("/api/fail-liveness", func(c echo.Context) error {
		LivenessProbeOK = false
		log.Printf("[%s] fail liveness\n", hostname)
		return c.JSON(http.StatusOK, response{
			Status:    http.StatusOK,
			Hostname:  hostname,
			Liveness:  LivenessProbeOK,
			Readiness: ReadinessProbeOK,
		})
	})

	e.GET("/api/fail-readiness", func(c echo.Context) error {
		ReadinessProbeOK = false
		log.Printf("[%s] fail readiness\n", hostname)
		return c.JSON(http.StatusOK, response{
			Status:    http.StatusOK,
			Hostname:  hostname,
			Liveness:  LivenessProbeOK,
			Readiness: ReadinessProbeOK,
		})
	})

	e.RouteNotFound("/*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, baseResponse{
			Status:   http.StatusNotFound,
			Hostname: hostname,
		})
	})

	e.GET("/api/fix-readiness", func(c echo.Context) error {
		ReadinessProbeOK = true
		log.Printf("[%s] fix readiness\n", hostname)
		return c.JSON(http.StatusOK, response{
			Status:    http.StatusOK,
			Hostname:  hostname,
			Liveness:  LivenessProbeOK,
			Readiness: ReadinessProbeOK,
		})
	})

	log.Printf("[%s] starting probes-example-server\n", hostname)
	log.Printf("[%s] Listening on 0.0.0.0:8000, see http://127.0.0.1:8000\n", hostname)
	return e.Start(":8000")
}
