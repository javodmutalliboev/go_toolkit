package struct_package

import (
	"encoding/json"
	"nam_tourism_web_server/type_package"
	"net/http"
)

type Response struct {
	Status type_package.ResponseStatus `json:"status"`
	Code   int                         `json:"code"`
	Data   any                         `json:"data"`
}

func (res *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}
