package users

import (
	"net/http"

	"github.com/SohamRatnaparkhi/blogx-backend-go/authentication/db/database"

	"github.com/SohamRatnaparkhi/blogx-backend-go/authentication/pkg/utils"
)

func HandleUserLogout(w http.ResponseWriter, req *http.Request, user database.User) {
	//clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "auth_token",
		Value: "",
		Path:  "/",
	})
	utils.ResponseJson(w, http.StatusAccepted, user)
}
