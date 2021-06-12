package internal

import "net/http"

type Sessioner interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Invalidate()
	Save(r *http.Request, w http.ResponseWriter)
}
