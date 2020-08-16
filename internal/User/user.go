package User

import (
	"context"
	"fmt"
	"github.com/volatiletech/authboss/v3"
)

type User struct {
	Name     string
	Password string
	Role     string
}
type Store struct {
	Users  map[string]User
	Tokens map[string][]string
}

//implements user interface
func (u *User) PutPID(pid string) {
	u.Name = pid
}

//implements user interface
func (u User) GetPID() string {
	return u.Name
}

//implements user interface
func (u User) GetPassword() (password string) {
	return u.Password
}

//implements user interface
func (u *User) PutPassword(password string) {
	u.Password = password
}

//create store
func NewStore() *Store {
	fmt.Println("Store Make")
	s := Store{Users: map[string]User{
		"Admin": {Name: "Admin", Password: "$2a$10$XtW/BrS5HeYIuOCXYe8DFuInetDMdaarMUJEOg/VA/JAIDgw3l4aG", Role: "admin"},
		"User":  {Name: "User", Password: "$2a$10$XtW/BrS5HeYIuOCXYe8DFuInetDMdaarMUJEOg/VA/JAIDgw3l4aG", Role: "user"},
	},
		Tokens: make(map[string][]string, 1),
	}
	return &s
}

//...
func (s Store) Save(_ context.Context, user authboss.User) error {
	fmt.Println("user save")
	u := user.(*User)
	s.Users[u.Name] = *u
	return nil
}

//...
func (s *Store) Load(_ context.Context, key string) (user authboss.User, err error) {

	u, ok := s.Users[key]
	if !ok {
		fmt.Println("User not found", "KEY:", key)
		return nil, authboss.ErrUserNotFound
	}

	fmt.Println("Load user:", s.Users[key])
	return &u, nil
}

// New
func (s Store) New(_ context.Context) authboss.User {
	return &User{}
}
func (s Store) Create(_ context.Context, user authboss.User) error {
	u := user.(*User)

	if _, ok := s.Users[u.Name]; ok {
		return authboss.ErrUserFound
	}
	u.Role = "user"

	fmt.Println("New User", u)

	s.Users[u.Name] = *u
	return nil
}

func (s Store) AddRememberToken(_ context.Context, pid, token string) error {
	s.Tokens[pid] = append(s.Tokens[pid], token)
	fmt.Println("Token +")
	return nil
}

func (s Store) DelRememberTokens(_ context.Context, pid string) error {
	delete(s.Tokens, pid)
	fmt.Println("Token -")
	return nil
}

func (s Store) UseRememberToken(_ context.Context, pid, token string) error {
	tokens, ok := s.Tokens[pid]
	if !ok {
		fmt.Println("Token not found")
		return authboss.ErrTokenNotFound
	}

	for i, tok := range tokens {
		if tok == token {
			tokens[len(tokens)-1] = tokens[i]
			s.Tokens[pid] = tokens[:len(tokens)-1]
			fmt.Println("Token Found")
			return nil
		}
	}

	return authboss.ErrTokenNotFound
}
