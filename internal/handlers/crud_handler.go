package handlers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowCRUDPage(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/crud")

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/crud.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "CRUD",
		"Username": username,
		"Menus":    menus,
	})
}
