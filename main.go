package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"

  	"github.com/google/uuid"
)

// UUIDパターンにマッチする正規表現
var uuidRegex = regexp.MustCompile(`^/([a-f\d]{8}-[a-f\d]{4}-[a-f\d]{4}-[a-f\d]{4}-[a-f\d]{12})$`)

// HTMLテンプレート
const htmlTemplateStr = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UUID Grid</title>
    <style>
        .grid-container {
            display: grid;
            grid-template-columns: repeat(5, 140px);
            gap: 10px;
        }
        .grid-item {
            width: 140px;
            height: 140px;
            border: 1px solid black;
            text-align: center;
            display: flex;
            justify-content: center;
            align-items: center;
        }
    </style>
</head>
<body>
    <div class="grid-container">
        {{range .}}
            <div class="grid-item">{{.}}</div>
        {{end}}
    </div>
</body>
</html>
`

func main() {
	http.Handle("/", http.HandlerFunc(handleRequest))

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	matches := uuidRegex.FindStringSubmatch(r.URL.Path)
	if len(matches) != 2 {
		http.NotFound(w, r)
		return
	}
	uuid := matches[1]

	uuids, err := generateUUIDs(uuid)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	renderTemplate(w, uuids)
}

// generateUUIDs は、与えられたUUIDを元に24個のUUIDを生成します。
func generateUUIDs(inputUUID string) ([]string, error) {
 	u, err := uuid.Parse(inputUUID)
	if err != nil {
		return nil, err
	}

        var uuids []string
        for i := 1; i <= 25; i++ {
                newUUID := manipulateUUID(u, i)
                uuids = append(uuids, newUUID)
        }

	// 中央のマス（13番目）を"FREE"に設定
	uuids[12] = "FREE"
	return uuids, nil
}

func manipulateUUID(u uuid.UUID, seed int) string {
	bytes := u[:]
	for i := range bytes {
		bytes[i] = bytes[i] + byte(seed)
	}

	newUUID := uuid.NewSHA1(u, bytes)
	return newUUID.String()
}

// renderTemplate は与えられたUUIDのリストを使ってHTMLをレンダリングし、レスポンスとして送ります。
func renderTemplate(w http.ResponseWriter, uuids []string) {
	tmpl, err := template.New("grid").Parse(htmlTemplateStr)
	if err != nil {
		http.Error(w, "Error creating template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, uuids)
}

