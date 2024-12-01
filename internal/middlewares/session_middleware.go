package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key")) // Replace with a secure key

func RequireLogin(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	if _, ok := session.Values["user"]; !ok {
		// Redirect to login if no valid session exists
		c.Redirect(http.StatusFound, "/")
		c.Abort()
		return
	}
	c.Next()
}
