package handlers

import (
	"fmt"
	"golang-web/internal/database"
	"golang-web/internal/models"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListPayments displays all payments
func ListPayments(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/payments")

	var payments []models.Payment
	database.DB.Preload("Menu").Find(&payments)

	// Debugging: Check the fetched payments and their menus
	fmt.Printf("Fetched payments: %+v\n", payments)

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/payments_list.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":       "Payments",
		"Username":    username,
		"Menus":       menus,
		"AllPayments": payments,
	})
}

// ShowCreatePayment displays the form to create a new payment
func ShowCreatePayment(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/payments/create")

	var tables []models.Table
	var menusList []models.Menu

	resultTables := database.DB.Find(&tables)
	resultMenus := database.DB.Find(&menusList)

	if resultTables.Error != nil || resultMenus.Error != nil {
		c.String(http.StatusInternalServerError, "Error fetching data for payments")
		return
	}

	// Debugging: Check fetched menus
	fmt.Printf("Fetched tables: %+v\n", tables)
	fmt.Printf("Fetched menus: %+v\n", menusList)

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/payments_create.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "Add Payment",
		"Username": username,
		"Menus":    menus,
		"Tables":   tables,
		"MenuList": menusList, // Ensure this matches the template key
	})
}

// CreatePayment handles adding a new payment
func CreatePayment(c *gin.Context) {

	var payment models.Payment

	nomorMeja, _ := strconv.Atoi(c.PostForm("nomor_meja"))
	menuID, _ := strconv.Atoi(c.PostForm("menu_id"))
	jumlah, _ := strconv.Atoi(c.PostForm("jumlah"))

	payment.NomorMeja = nomorMeja
	payment.MenuID = uint64(menuID)
	payment.Jumlah = jumlah
	payment.MetodePembayaran = c.PostForm("metode_pembayaran")
	payment.Status = "belum dibayar"

	database.DB.Create(&payment)
	c.Redirect(http.StatusFound, "/payments")
}

// ShowEditPayment displays the form to edit an existing payment
func ShowEditPayment(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/payments/edit")

	var payment models.Payment
	var tables []models.Table
	var menusList []models.Menu

	id := c.Param("id")
	database.DB.Preload("Menu").First(&payment, id) // Preload the associated Menu
	database.DB.Find(&tables)
	database.DB.Find(&menusList)

	// Debugging: Check the fetched payment and its associated menu
	fmt.Printf("Fetched payment: %+v\n", payment)

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/payments_edit.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "Edit Payment",
		"Username": username,
		"Menus":    menus,
		"Payment":  payment,
		"Tables":   tables,
		"MenuList": menusList,
	})
}

// UpdatePayment handles updating an existing payment
func UpdatePayment(c *gin.Context) {
	var payment models.Payment

	id := c.Param("id")
	database.DB.First(&payment, id)

	nomorMeja, _ := strconv.Atoi(c.PostForm("nomor_meja"))
	menuID, _ := strconv.Atoi(c.PostForm("menu_id"))
	jumlah, _ := strconv.Atoi(c.PostForm("jumlah"))

	payment.NomorMeja = nomorMeja
	payment.MenuID = uint64(menuID)
	payment.Jumlah = jumlah
	payment.MetodePembayaran = c.PostForm("metode_pembayaran")
	payment.Status = c.PostForm("status")

	database.DB.Save(&payment)
	c.Redirect(http.StatusFound, "/payments")
}

// DeletePayment handles deleting a payment
func DeletePayment(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Payment{}, id)
	c.Redirect(http.StatusFound, "/payments")
}
