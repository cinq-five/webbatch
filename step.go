package webbach

import "net/http"

type Step func(writer http.ResponseWriter, request *http.Request) bool
