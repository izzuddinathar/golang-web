package handlers

type Menu struct {
	Name   string
	URL    string
	Active bool
}

var roleMenus = map[string][]string{
	"owner":  {"Meja", "Pesanan", "User", "Menu", "Pembayaran"},
	"waiter": {"Meja", "Pesanan"},
	"admin":  {"User", "Menu", "Pembayaran"},
}

func getMenus(role string, currentPath string) []Menu {
	var menus []Menu
	menuURLs := map[string]string{
		"Meja":       "/meja",
		"Pesanan":    "/pesanan",
		"User":       "/users",
		"Menu":       "/menus",
		"Pembayaran": "/pembayaran",
		"Dashboard":  "/dashboard",
		"CRUD":       "/crud",
		"Profile":    "/profile",
	}

	if menuNames, ok := roleMenus[role]; ok {
		for _, name := range menuNames {
			menus = append(menus, Menu{
				Name:   name,
				URL:    menuURLs[name],
				Active: currentPath == menuURLs[name],
			})
		}
	}
	return menus
}