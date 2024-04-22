package core

import "net/http"

type HttpExchange func(w http.ResponseWriter, r *http.Request)
