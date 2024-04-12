package middleware

import (
	"log"
	"net/http"

	"github.com/javodmutalliboev/go_toolkit/constants"
	"github.com/javodmutalliboev/go_toolkit/interface_package"
	"github.com/javodmutalliboev/go_toolkit/session"
	"github.com/javodmutalliboev/go_toolkit/struct_package"
	"github.com/javodmutalliboev/go_toolkit/type_package"
)

func AdminAuth() type_package.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var res interface_package.Response = &struct_package.Response{}
			session, err := session.GetSession(r, "session")
			if err != nil {
				log.Printf("%s: %s", r.URL.Path, err.Error())
				res = &struct_package.Response{
					Status: constants.Error,
					Code:   http.StatusUnauthorized,
					Data:   "Unauthorized",
				}

				res.Send(w)
				return
			}

			// Check if user is authenticated
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				log.Printf("%s: %s", r.URL.Path, "Unauthorized")
				res = &struct_package.Response{
					Status: constants.Error,
					Code:   http.StatusUnauthorized,
					Data:   "Unauthorized",
				}

				res.Send(w)
				return
			}

			session.Options.MaxAge = 24 * 60 * 60
			session.Save(r, w)

			f(w, r)
		}
	}
}
