package handlers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowDashboard(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/dashboard")

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/dashboard.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "Dashboard",
		"Username": username,
		"Menus":    menus,
	})
}
