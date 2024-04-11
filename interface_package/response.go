package interface_package

import "net/http"

type Response interface {
	Send(w http.ResponseWriter)
}
