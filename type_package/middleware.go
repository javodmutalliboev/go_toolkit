package type_package

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc
