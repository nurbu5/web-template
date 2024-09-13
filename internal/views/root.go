package views

func RootView() (*View, error) {
	return newView("views.RootView", "base.html", []string{"root.html"})
}
