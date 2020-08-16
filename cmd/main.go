package main

import (
	"NPOtest/internal/Server"
	"NPOtest/internal/User"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	_"github.com/volatiletech/authboss/v3/logout"
	abclientstate "github.com/volatiletech/authboss-clientstate"
	abrenderer "github.com/volatiletech/authboss-renderer"
	"github.com/volatiletech/authboss/v3"
	_ "github.com/volatiletech/authboss/v3/auth"
	"github.com/volatiletech/authboss/v3/defaults"
	remember2 "github.com/volatiletech/authboss/v3/remember"
	"net/http"
)

func main() {

	ab := authboss.New()
	//Create store
	cookieStoreKey, _ := base64.StdEncoding.DecodeString(`NpEPi8pEjKVjLGJ6kYCS+VTCzi6BUuDzU0wrwXyf5uDPArtlofn2AG6aTMiPmN3C909rsEWMNqJqhIVPGP3Exg==`)
	sessionStoreKey, _ := base64.StdEncoding.DecodeString(`AbfYwmmt8UCwUuhd9qvfNA9UCuN1cVcKJN1ofbiky6xCyyBj20whe40rJa3Su0WOWLWcPpO1taqJdsEI/65+JA==`)
	cookieStore := abclientstate.NewCookieStorer(cookieStoreKey, nil)
	sessionStore := abclientstate.NewSessionStorer("NPO", sessionStoreKey)
	cookieStore.HTTPOnly = false
	cookieStore.Secure = false
	UserStore := User.NewStore()

	//setting ab
	ab.Storage.Server = UserStore
	ab.Storage.CookieState = cookieStore
	ab.Storage.SessionState = sessionStore
	ab.Paths.Mount = "/auth"
	ab.Modules.LogoutMethod = "GET"
	ab.Paths.RootURL = "http://localhost:8080/"
	ab.Paths.AuthLoginOK = "http://localhost:8080/"
	ab.Config.Core.ViewRenderer = abrenderer.NewHTML("/login", "tmp")
	defaults.SetCore(&ab.Config, false, false)
	if err := ab.Init(); err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	//authboss route
	router.Use(ab.LoadClientStateMiddleware, remember2.Middleware(ab), authboss.ModuleListMiddleware(ab))
	router.HandleFunc("/", index())
	router.PathPrefix("/auth").Handler(http.StripPrefix("/auth", ab.Config.Core.Router))
	router.PathPrefix("/login").Handler(http.StripPrefix("/login", ab.Config.Core.Router))
	router.PathPrefix("/logout").Handler(http.StripPrefix("/logout", ab.Config.Core.Router))
	//route
	u := router.Host("localhost").Subrouter()
	u.Use(authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondUnauthorized))
	u.HandleFunc("/foo", authFoo())
	u.HandleFunc("/bar", authBar())

	//admin route
	a := router.Host("localhost").Subrouter()
	a.Use(authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondUnauthorized), RoleControl(ab))
	a.HandleFunc("/sigma", authSigma())

	serv := Server.NewServer(router)
	serv.ListenAndServe()
}
//index page
func index() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1> Test Page</h1>" +
							 "<a href ='/foo'>Foo</a><br>"+
		                     "<a href ='/bar'>Bar</a><br>"+
			                 "<a href ='/sigma'>Sigma</a><br>"+
			                 "<a href ='/auth/login'>Login</a><br>"+
			                 "<a href ='/auth/logout'>Logout</a>")
	}}
//foo
func authFoo() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<br><p>Foo</p><br><a href ='/'>back</a>")

	}
}
//bar
func authBar() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<br><p>Bar</p><br><a href ='/'>back</a>")

	}
}
//sigma
func authSigma() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<br><p>Sigma</p><br><a href ='/'>back</a>")

	}
}



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
