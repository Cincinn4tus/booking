package core

import "github.com/gorilla/sessions"

// create a session with a given email
var sessionCookie = sessions.NewCookieStore([]byte("secret"))
