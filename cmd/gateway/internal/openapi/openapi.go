package openapi

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/adm87/onyx-server/pkg/proto/gen/grpc/auth"
	"github.com/adm87/onyx-server/pkg/proto/gen/grpc/user"
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

var tmpl = template.Must(template.New("swagger-ui").Parse(swaggerUITemplate))

func RegisterSwaggerUI(mux *http.ServeMux) {
	var specs []specEntry
	for name := range openAPISpecs {
		specs = append(specs, specEntry{Name: name, URL: "/openapi/" + name + ".json"})
	}

	specsJSON, err := json.Marshal(specs)
	if err != nil {
		panic(err) // specs is static at startup; this can only fail on a programming error
	}

	data := templateData{SpecsJSON: template.JS(specsJSON)}

	for name, spec := range openAPISpecs {
		spec := spec
		mux.HandleFunc("/openapi/"+name+".json", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write(spec)
		})
	}

	mux.HandleFunc("/docs/index.html", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "failed to render docs", http.StatusInternalServerError)
		}
	})
}
