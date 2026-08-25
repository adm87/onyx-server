package openapi

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/adm87/onyx-server/pkg/proto/gen/grpc/auth"
	"github.com/adm87/onyx-server/pkg/proto/gen/grpc/user"
	"go.uber.org/zap"
)

var openAPISpecs = map[string][]byte{
	"auth.v1": auth.AuthV1OpenAPISpec,
	"user.v1": user.UserV1OpenAPISpec,
}

const swaggerUITemplate = `<!DOCTYPE html>
<html>
<head>
  <title>Onyx API Docs</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
  <style>
    body { margin: 0; }
  </style>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-standalone-preset.js"></script>
  <script>
    window.onload = () => {
      SwaggerUIBundle({
        urls: {{ .SpecsJSON }},
        dom_id: "#swagger-ui",
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        plugins: [
          SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout"
      });
    };
  </script>
</body>
</html>`

type specEntry struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type templateData struct {
	SpecsJSON template.JS
}

func RegisterSwaggerUI(mux *http.ServeMux, log *zap.Logger) http.HandlerFunc {
	const openAPISpecsPath = "/openapi/%s.json"

	var specs []specEntry
	for name := range openAPISpecs {
		specs = append(specs, specEntry{Name: name, URL: fmt.Sprintf(openAPISpecsPath, name)})
	}

	specsJSON, err := json.Marshal(specs)
	if err != nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "failed to marshal OpenAPI specs", http.StatusInternalServerError)
		})
	}

	for name, spec := range openAPISpecs {
		mux.HandleFunc(fmt.Sprintf(openAPISpecsPath, name), func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write(spec)
		})
	}

	tmpl := template.Must(template.New("swagger-ui").Parse(swaggerUITemplate))
	data := templateData{SpecsJSON: template.JS(specsJSON)}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "failed to render docs", http.StatusInternalServerError)
		}
	})
}
