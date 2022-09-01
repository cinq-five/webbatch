package webbach

import (
	"context"
	"net/http"
)

type Step func(writer http.ResponseWriter, request *http.Request, ctx *context.Context) bool
