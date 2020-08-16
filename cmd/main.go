package main

import (
	"NPOtest/internal/Route"
	"NPOtest/internal/Server"
	"NPOtest/internal/User"
	"encoding/base64"
	"github.com/gorilla/mux"
	abclientstate "github.com/volatiletech/authboss-clientstate"
	abrenderer "github.com/volatiletech/authboss-renderer"
	"github.com/volatiletech/authboss/v3"
	_ "github.com/volatiletech/authboss/v3/auth"
	"github.com/volatiletech/authboss/v3/defaults"
	_ "github.com/volatiletech/authboss/v3/logout"
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
	router.HandleFunc("/", Route.Index())
	router.PathPrefix("/auth").Handler(http.StripPrefix("/auth", ab.Config.Core.Router))
	router.PathPrefix("/login").Handler(http.StripPrefix("/login", ab.Config.Core.Router))
	router.PathPrefix("/logout").Handler(http.StripPrefix("/logout", ab.Config.Core.Router))
	//route
	u := router.Host("localhost").Subrouter()
	u.Use(authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondUnauthorized))
	u.HandleFunc("/foo", Route.AuthFoo())
	u.HandleFunc("/bar", Route.AuthBar())

	//admin route
	a := router.Host("localhost").Subrouter()
	a.Use(authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondUnauthorized), Route.RoleControl(ab))
	a.HandleFunc("/sigma", Route.AuthSigma())

	serv := Server.NewServer(router)
	serv.ListenAndServe()
}
