package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: embed_assets <base_templ_go_path> <css_file_path> <js_file_path>")
		os.Exit(1)
	}

	baseTemplGoPath := os.Args[1]
	cssFilePath := os.Args[2]
	jsFilePath := os.Args[3]

	baseTemplGoContent, err := ioutil.ReadFile(baseTemplGoPath)
	if err != nil {
		fmt.Printf("Error reading base_templ.go: %v\n", err)
		os.Exit(1)
	}

	cssContent, err := ioutil.ReadFile(cssFilePath)
	if err != nil {
		fmt.Printf("Error reading CSS file: %v\n", err)
		os.Exit(1)
	}

	jsContent, err := ioutil.ReadFile(jsFilePath)
	if err != nil {
		fmt.Printf("Error reading JS file: %v\n", err)
		os.Exit(1)
	}

	baseTemplGoStr := string(baseTemplGoContent)
	cssStr := string(cssContent)
	jsStr := string(jsContent)

	baseTemplGoStr = strings.Replace(
		baseTemplGoStr,
		`<style id="main-css"></style>`,
		fmt.Sprintf(`<style id="main-css">%s</style>`, cssStr),
		-1,
	)

	baseTemplGoStr = strings.Replace(
		baseTemplGoStr,
		`<script id="main-js"></script>`,
		fmt.Sprintf(`<script id="main-js">%s</script>`, jsStr),
		-1,
	)

	err = ioutil.WriteFile(baseTemplGoPath, []byte(baseTemplGoStr), 0644)
	if err != nil {
		fmt.Printf("Error writing to base_templ.go: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Assets embedded successfully.")
}
