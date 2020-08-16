package Route

import (
	"NPOtest/internal/User"
	"fmt"
	"net/http"
	"github.com/volatiletech/authboss/v3"
)

// Role middleware: admin/user
func RoleControl(ab *authboss.Authboss) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, e := ab.CurrentUser(r)
			if e != nil {
				fmt.Fprintln(w, e)
			}
			userRole := user.(*User.User).Role
			if userRole == "admin" {
				next.ServeHTTP(w, r)
			} else {
				fmt.Fprintln(w, "<p>Access is denied</p><br><a href ='/'>back</a>")
			}

		})
	}
}