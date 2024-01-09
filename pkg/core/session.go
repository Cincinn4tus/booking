package core

import "github.com/gorilla/sessions"

// create a session with a given email
var SessionCookie = sessions.NewCookieStore([]byte("secret"))
