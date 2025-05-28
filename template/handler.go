package template

import (
	"encoding/json"
	"fmt"
	"github.com/aymerick/raymond"
	"net/http"
)

func GenerateCV(w http.ResponseWriter, r *http.Request) {
	source := `<div class="entry">
  <h1>{{title}}</h1>
  <div class="body">
    {{body}}
  </div>
</div>
`

	ctxList := []map[string]string{
		{
			"title": "My New Post",
			"body":  "This is my first post!",
		},
		{
			"title": "Here is another post",
			"body":  "This is my second post!",
		},
	}

	// parse template
	tpl, err := raymond.Parse(source)
	if err != nil {
		panic(err)
	}

	result := ""

	for _, ctx := range ctxList {
		// render template
		item, err := tpl.Exec(ctx)
		if err != nil {
			panic(err)
		}

		result += item + "\n"
	}

	respondJSON(w, 200, BaseResponse{
		Message: "success",
		Data:    fmt.Sprintf(result),
	})
}

func respondJSON(w http.ResponseWriter, status int, payload BaseResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		return
	}
}
