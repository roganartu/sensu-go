package middlewares

import (
	"net/http"

	types "github.com/sensu/sensu-go/api/core/v2"
)

// Edition is an HTTP middleware that provides the Sensu Edition through a header
type Edition struct {
	Name string
}

// Then middleware
func (e Edition) Then(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(types.EditionHeader, e.Name)
		next.ServeHTTP(w, r)
	})
}
