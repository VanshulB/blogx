package server

import (
	"fmt"
	"net/http"

	"github.com/SohamRatnaparkhi/blogx-backend-go/blog/pkg/utils"
	"github.com/SohamRatnaparkhi/blogx-backend-go/db"
)

func HealthCheck(res http.ResponseWriter, _ *http.Request) {
	databaseObject := db.DbClient
	if databaseObject == nil {
		utils.ErrorResponse(res, http.StatusInternalServerError, fmt.Errorf("database error"))
		return
	}
	type resp struct {
		Status string `json:"status"`
	}
	utils.ResponseJson(res, 200, resp{
		Status: "ok",
	})
}
