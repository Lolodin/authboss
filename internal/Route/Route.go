package Route

import (
	"fmt"
	"net/http"
)
//index localhost:8080/
func Index() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1> Test Page</h1>" +
			"<a href ='/foo'>Foo</a><br>"+
			"<a href ='/bar'>Bar</a><br>"+
			"<a href ='/sigma'>Sigma</a><br>"+
			"<a href ='/auth/login'>Login</a><br>"+
			"<a href ='/auth/logout'>Logout</a>")
	}}
//foo
func AuthFoo() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<br><p>Foo</p><br><a href ='/'>back</a>")

	}
}
//bar
func AuthBar() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<br><p>Bar</p><br><a href ='/'>back</a>")

	}
}
//sigma
func AuthSigma() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<br><p>Sigma</p><br><a href ='/'>back</a>")

	}
}
