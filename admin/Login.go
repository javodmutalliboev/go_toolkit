package admin

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/javodmutalliboev/go_toolkit/constants"
	"github.com/javodmutalliboev/go_toolkit/interface_package"
	"github.com/javodmutalliboev/go_toolkit/session"
	"github.com/javodmutalliboev/go_toolkit/struct_package"
)

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var admin interface_package.Admin = &struct_package.Admin{}
		var res interface_package.Response = &struct_package.Response{}

		err := json.NewDecoder(r.Body).Decode(admin)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err.Error())
			res = &struct_package.Response{
				Status: constants.Error,
				Code:   http.StatusBadRequest,
				Data:   err.Error(),
			}
			res.Send(w)
			return
		}

		err = admin.Login()
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				log.Printf("%s: %s", r.URL.Path, err.Error())
				res = &struct_package.Response{
					Status: constants.Error,
					Code:   http.StatusUnauthorized,
					Data:   "Invalid email",
				}
				res.Send(w)
				return
			}

			if err.Error() == "invalid password" {
				log.Printf("%s: %s", r.URL.Path, err.Error())
				res = &struct_package.Response{
					Status: constants.Error,
					Code:   http.StatusUnauthorized,
					Data:   "Invalid password",
				}
				res.Send(w)
				return
			}

			log.Printf("%s: %s", r.URL.Path, err.Error())
			res = &struct_package.Response{
				Status: constants.Error,
				Code:   http.StatusInternalServerError,
				Data:   "Internal server error",
			}
			res.Send(w)
			return
		}

		session, err := session.GetSession(r, "session")
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err.Error())
			res = &struct_package.Response{
				Status: constants.Error,
				Code:   http.StatusInternalServerError,
				Data:   "Internal server error",
			}
			res.Send(w)
			return
		}

		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   60 * 60 * 24,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteNoneMode,
		}
		session.Values["authenticated"] = true
		session.Save(r, w)

		res = &struct_package.Response{
			Status: constants.Success,
			Code:   http.StatusOK,
			Data:   "Logged in",
		}
		res.Send(w)
	}
}
