package admin

import (
	"log"
	"net/http"

	"github.com/javodmutalliboev/go_toolkit/constants"
	"github.com/javodmutalliboev/go_toolkit/interface_package"
	"github.com/javodmutalliboev/go_toolkit/session"
	"github.com/javodmutalliboev/go_toolkit/struct_package"
)

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res interface_package.Response = &struct_package.Response{}
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

		session.Values["authenticated"] = false
		session.Save(r, w)

		res = &struct_package.Response{
			Status: constants.Success,
			Code:   http.StatusOK,
			Data:   "Logged out",
		}

		res.Send(w)
	}
}
