package handlers

import (
	"golang-web/internal/database"
	"golang-web/internal/models"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("secret-key")) // Replace with a secure key

func ShowLoginPage(c *gin.Context) {
	// Retrieve session
	session, _ := store.Get(c.Request, "session")

	// Check if the user is already logged in
	if _, ok := session.Values["user"]; ok {
		// Redirect to /dashboard if a session exists
		c.Redirect(http.StatusFound, "/dashboard")
		return
	}

	// Render the login page if no session exists
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(c.Writer, gin.H{"Error": ""})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(c.Writer, gin.H{"Error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(c.Writer, gin.H{"Error": "Invalid username or password"})
		return
	}

	// Create session
	session, _ := store.Get(c.Request, "session")
	session.Values["user"] = username
	session.Values["role"] = user.Role
	session.Save(c.Request, c.Writer)

	// Redirect to dashboard
	c.Redirect(http.StatusFound, "/dashboard")
}

func Logout(c *gin.Context) {
	// Retrieve session
	session, _ := store.Get(c.Request, "session")

	// Invalidate the session
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)

	// Redirect to login page
	c.Redirect(http.StatusFound, "/")
}
