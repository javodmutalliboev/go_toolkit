package session

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/javodmutalliboev/go_toolkit/environment"
)

var key []byte
var store *sessions.CookieStore

func init() {
	environment.Load()

	key = []byte(os.Getenv("SESSION_KEY"))
	store = sessions.NewCookieStore(key)
}

func GetSession(r *http.Request, name string) (*sessions.Session, error) {
	session, err := store.Get(r, name)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func GetSession2() {
	log.Println(key)
}
