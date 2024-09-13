package views

func TestView() (*View, error) {
	return newView("view.TestView", "base.html", []string{"test.html"})
}
