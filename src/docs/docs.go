package docs

import (
	_ "embed"

	"html/template"
	"strings"

	"github.com/gofiber/fiber/v2"
)

//go:embed templates/docs.html
var docTemplate string

//go:embed templates/docs-detail.html
var detailDocTemplate string

func RenderMainDoc(c *fiber.Ctx, app *fiber.App) error {
	documentedRoutes := getDocumentedRoutes(app)

	return renderDocs(
		c,
		docTemplate,
		fiber.Map{
			"Routes": documentedRoutes,
		},
	)
}

func RenderDetailDoc(c *fiber.Ctx, data fiber.Map) error {
	return renderDocs(c, detailDocTemplate, data)
}

func renderDocs(c *fiber.Ctx, docTemplate string, data fiber.Map) error {
	index, err := template.New("docs.html").Parse(docTemplate)
	if err != nil {
		return err
	}

	c.Response().Header.SetContentType(fiber.MIMETextHTMLCharsetUTF8)

	return index.Execute(c, data)
}

func getDocumentedRoutes(app *fiber.App) map[string]string {
	documentedRoutes := make(map[string]string)

	for _, route := range app.GetRoutes() {
		if !strings.Contains(route.Path, "/docs") {
			continue
		}

		name, found := documentedRoutes[route.Path]
		if found && name != "" {
			continue
		}

		documentedRoutes[route.Path] = route.Name
	}

	return documentedRoutes
}
