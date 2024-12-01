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

// ListOrders displays all orders
func ListOrders(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/orders")

	var orders []models.Order
	database.DB.Preload("Menu").Find(&orders) // Preload associated menu details

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/orders_list.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":     "Orders",
		"Username":  username,
		"Menus":     menus,
		"AllOrders": orders,
	})
}

// ShowCreateOrder displays the form to create a new order
func ShowCreateOrder(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/orders/create")

	var tables []models.Table
	var menusList []models.Menu

	database.DB.Find(&tables)    // Fetch all tables
	database.DB.Find(&menusList) // Fetch all menus

	fmt.Printf("Fetched menus: %+v\n", menusList)

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/orders_create.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "Add Order",
		"Username": username,
		"Menus":    menus,
		"Tables":   tables,
		"MenuList": menusList,
	})
}

// CreateOrder handles adding a new order
func CreateOrder(c *gin.Context) {
	var order models.Order

	// Parse form values
	nomorMeja, err := strconv.Atoi(c.PostForm("nomor_meja"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid table number format")
		return
	}

	menuID, err := strconv.Atoi(c.PostForm("menu_id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid menu ID format")
		return
	}

	jumlah, err := strconv.Atoi(c.PostForm("jumlah"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid quantity format")
		return
	}

	order.NomorMeja = nomorMeja
	order.MenuID = uint64(menuID)
	order.Jumlah = jumlah
	order.StatusPesanan = c.PostForm("status_pesanan")

	// Save to database
	database.DB.Create(&order)
	c.Redirect(http.StatusFound, "/orders")
}

// ShowEditOrder displays the form to edit an existing order
func ShowEditOrder(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	role, roleOk := session.Values["role"].(string)
	username, userOk := session.Values["user"].(string)

	if !roleOk || !userOk {
		c.Redirect(http.StatusFound, "/")
		return
	}

	menus := getMenus(role, "/orders/edit")

	var order models.Order
	var tables []models.Table
	var menusList []models.Menu

	id := c.Param("id")
	result := database.DB.First(&order, id)
	if result.Error != nil {
		c.String(http.StatusNotFound, "Order not found")
		return
	}

	database.DB.Find(&tables)    // Fetch all tables
	database.DB.Find(&menusList) // Fetch all menus

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/orders_edit.html"))
	tmpl.Execute(c.Writer, gin.H{
		"Title":    "Edit Order",
		"Username": username,
		"Menus":    menus,
		"Order":    order,
		"Tables":   tables,
		"MenuList": menusList,
		"OrderID":  order.OrderID,
	})
}

// UpdateOrder handles updating an existing order
func UpdateOrder(c *gin.Context) {
	var order models.Order

	// Find the order
	id := c.Param("id")
	result := database.DB.First(&order, id)
	if result.Error != nil {
		c.String(http.StatusNotFound, "Order not found")
		return
	}

	// Parse form values
	nomorMeja, err := strconv.Atoi(c.PostForm("nomor_meja"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid table number format")
		return
	}

	menuID, err := strconv.Atoi(c.PostForm("menu_id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid menu ID format")
		return
	}

	jumlah, err := strconv.Atoi(c.PostForm("jumlah"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid quantity format")
		return
	}

	order.NomorMeja = nomorMeja
	order.MenuID = uint64(menuID)
	order.Jumlah = jumlah
	order.StatusPesanan = c.PostForm("status_pesanan")

	// Save updates
	database.DB.Save(&order)
	c.Redirect(http.StatusFound, "/orders")
}

// DeleteOrder handles deleting an order
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Delete(&models.Order{}, id)
	if result.Error != nil {
		c.String(http.StatusNotFound, "Order not found")
		return
	}
	c.Redirect(http.StatusFound, "/orders")
}
