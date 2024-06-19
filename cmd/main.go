package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.gohtml")),
	}
}

type Pollinator struct {
	ID int
	Name string
	Description string
	Count int
}

type Pollinators = []Pollinator

func newPollinators() Pollinators {
	return Pollinators {
		{
			ID: 0,
			Name: "Small Bee",
			Description: "Small bees are important pollinators. They are responsible for pollinating a variety of plants, including some that are not pollinated by honey bees.",
			Count: 0,
		},
		{
			ID: 1,
			Name: "Fly",
			Description: "Flies are also important pollinators. They are responsible for pollinating a variety of plants, including some that are not pollinated by bees.",
			Count: 0,
		},
		{
			ID: 2,
			Name: "Wasp",
			Description: "Wasps are important pollinators. They are responsible for pollinating a variety of plants, including some that are not pollinated by bees.",
			Count: 0,
		},
		{
			ID: 3,
			Name: "Carpenter Bee",
			Description: "Carpenter bees are important pollinators. They are responsible for pollinating a variety of plants, including some that are not pollinated by honey bees.",
			Count: 0,
		},
		{
			ID: 4,
			Name: "Bumblebee",
			Description: "Bumblebees are important pollinators. They are responsible for pollinating a variety of plants, including some that are not pollinated by honey bees.",
			Count: 0,
		},
		{
			ID: 5,
			Name: "Butterflies & Moths",
			Description: "Butterflies and moths are important pollinators. They are responsible for pollinating a variety of plants, including some that are not pollinated by bees.",
			Count: 0,
		},
		{
			ID: 6,
			Name: "Honey Bee",
			Description: "Honey bees are important pollinators. They are responsible for pollinating a variety of plants, including some that are not pollinated by bees.",
			Count: 0,
		},
		{
			ID: 7,
			Name: "Other Insects",
			Description: "Other insects are important pollinators. They are responsible for pollinating a variety of plants, including some that are not pollinated by bees.",
			Count: 0,
		},
	}
}

func addPollinator(p Pollinators, id int) Pollinators {
	for i, pollinator := range p {
		if pollinator.ID == id {
			p[i].Count++
		}
	}
	return p
}

type User struct {
	Name string
	Email string
	Phone string
}

type PageData struct {
	Pollinators Pollinators
	Time int
	User User
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplates()

	e.Static("/images", "images");
	e.Static("/fonts", "fonts");
	e.Static("/css", "css");

	e.GET("/", func(c echo.Context) error {
		// render index with an undefined object
		return c.Render(200, "index", nil)
	})

	e.GET("/begin", func(c echo.Context) error {
		pageData := PageData { Pollinators: newPollinators(), Time: 61, User: User { Name: c.QueryParam("name"), Email: c.QueryParam("email"), Phone: c.QueryParam("phone") }}

        if c.Request().Header.Get("HX-Request") == "true" {
			return c.Render(200, "button-grid", pageData)
        } else {
            return c.Render(200, "index", pageData)
		}
	})

	e.Logger.Fatal(e.Start(":1323"))
}