package handlers

import (
	"golang-web/internal/database"
	"golang-web/internal/models"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ListUsers displays a list of all users
func ListUsers(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/users")

	var users []models.User
	database.DB.Find(&users)

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/users_list.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "Users",
		"Username": username,
		"Menus":    menus,
		"Users":    users,
	})
}

// ShowCreateUser displays the form to create a new user
func ShowCreateUser(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/users/create")

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/users_create.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "Add User",
		"Username": username,
		"Menus":    menus,
	})
}

// CreateUser handles the submission of a new user
func CreateUser(c *gin.Context) {
	var user models.User
	user.Nama = c.PostForm("nama")
	user.Email = c.PostForm("email")
	user.Username = c.PostForm("username")
	password := c.PostForm("password")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	user.Role = c.PostForm("role")

	database.DB.Create(&user)
	c.Redirect(http.StatusFound, "/users")
}

func ShowEditUser(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	var user models.User
	id := c.Param("id") // Extract the ID from the URL
	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.String(http.StatusNotFound, "User not found")
		return
	}

	menus := getMenus(role, "/users/edit/"+id)

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/users_edit.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "Edit User",
		"Username": username,
		"Menus":    menus,
		"UserID":   user.UserID, // Pass the UserID to the template
		"User":     user,
	})
}

// UpdateUser handles the submission of updated user data
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	database.DB.First(&user, id)

	user.Nama = c.PostForm("nama")
	user.NoTelp = c.PostForm("no_telp")
	user.Email = c.PostForm("email")
	user.Username = c.PostForm("username")
	user.Role = c.PostForm("role")

	database.DB.Save(&user)
	c.Redirect(http.StatusFound, "/users")
}

// DeleteUser handles the deletion of a user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.User{}, id)
	c.Redirect(http.StatusFound, "/users")
}
