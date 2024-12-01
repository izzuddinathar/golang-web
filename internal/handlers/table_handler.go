package handlers

import (
	"golang-web/internal/database"
	"golang-web/internal/models"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListTables displays all tables
func ListTables(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/tables")

	var allTables []models.Table
	database.DB.Find(&allTables)

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/tables_list.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":     "Tables",
		"Username":  username,
		"Menus":     menus,
		"AllTables": allTables,
	})
}

// ShowCreateTable displays the form to create a new table
func ShowCreateTable(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/tables/create")

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/tables_create.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "Add Table",
		"Username": username,
		"Menus":    menus,
	})
}

// CreateTable handles adding a new table
func CreateTable(c *gin.Context) {
	var table models.Table
	nomorMeja, err := strconv.Atoi(c.PostForm("nomor_meja"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid table number format")
		return
	}
	table.NomorMeja = nomorMeja
	kapasitas, err := strconv.Atoi(c.PostForm("kapasitas"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid capacity format")
		return
	}
	table.Kapasitas = kapasitas
	table.Status = c.PostForm("status")

	database.DB.Create(&table)
	c.Redirect(http.StatusFound, "/tables")
}

// ShowEditTable displays the form to edit a table
func ShowEditTable(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	var table models.Table
	id := c.Param("id")
	result := database.DB.First(&table, id)
	if result.Error != nil {
		c.String(http.StatusNotFound, "Table not found")
		return
	}

	menus := getMenus(role, "/tables/edit/"+id)

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/tables_edit.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "Edit Table",
		"Username": username,
		"Menus":    menus,
		"TableID":  table.TableID,
		"Table":    table,
	})
}

// UpdateTable handles updating a table
func UpdateTable(c *gin.Context) {
	var table models.Table
	id := c.Param("id")
	database.DB.First(&table, id)

	nomorMeja, err := strconv.Atoi(c.PostForm("nomor_meja"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid table number format")
		return
	}
	table.NomorMeja = nomorMeja
	kapasitas, err := strconv.Atoi(c.PostForm("kapasitas"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid capacity format")
		return
	}
	table.Kapasitas = kapasitas
	table.Status = c.PostForm("status")

	database.DB.Save(&table)
	c.Redirect(http.StatusFound, "/tables")
}

// DeleteTable handles deleting a table
func DeleteTable(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Table{}, id)
	c.Redirect(http.StatusFound, "/tables")
}
