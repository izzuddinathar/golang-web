package handlers

import (
	"golang-web/internal/database"
	"golang-web/internal/models"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListMenus displays all menus
func ListMenus(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/menus")

	var allMenus []models.Menu
	result := database.DB.Find(&allMenus)
	if result.Error != nil {
		c.String(http.StatusInternalServerError, "Database error: %v", result.Error)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/menus_list.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "Menus",
		"Username": username,
		"Menus":    menus,
		"AllMenus": allMenus,
	})
}

// ShowCreateMenu displays the form to create a new menu
func ShowCreateMenu(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/menus/create")

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/menus_create.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "Add Menu",
		"Username": username,
		"Menus":    menus,
	})
}

// CreateMenu handles adding a new menu
func CreateMenu(c *gin.Context) {
	var menu models.Menu
	menu.NamaMenu = c.PostForm("nama_menu")
	menu.Deskripsi = c.PostForm("deskripsi")
	// Convert harga from string to float64
	harga, err := strconv.ParseFloat(c.PostForm("harga"), 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid price format")
		return
	}
	menu.Harga = harga
	menu.Kategori = c.PostForm("kategori")

	database.DB.Create(&menu)
	c.Redirect(http.StatusFound, "/menus")
}

// ShowEditMenu displays the form to edit a menu
func ShowEditMenu(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	var menu models.Menu
	id := c.Param("id")
	result := database.DB.First(&menu, id)
	if result.Error != nil {
		c.String(http.StatusNotFound, "Menu not found")
		return
	}

	menus := getMenus(role, "/menus/edit/"+id)

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/menus_edit.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "Edit Menu",
		"Username": username,
		"Menus":    menus,
		"MenuID":   menu.MenuID,
		"Menu":     menu,
	})
}

// UpdateMenu handles updating a menu
func UpdateMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	database.DB.First(&menu, id)

	menu.NamaMenu = c.PostForm("nama_menu")
	menu.Deskripsi = c.PostForm("deskripsi")
	// Convert harga from string to float64
	harga, err := strconv.ParseFloat(c.PostForm("harga"), 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid price format")
		return
	}
	menu.Harga = harga
	menu.Kategori = c.PostForm("kategori")

	database.DB.Save(&menu)
	c.Redirect(http.StatusFound, "/menus")
}

// DeleteMenu handles deleting a menu
func DeleteMenu(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Menu{}, id)
	c.Redirect(http.StatusFound, "/menus")
}
