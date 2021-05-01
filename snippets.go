package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

//go:embed templates/*.md
var templates embed.FS

type Snippet struct {
	Key         string
	Name        string
	Prefix      string   `json:"prefix"`
	Body        []string `json:"body"`
	Description string   `json:"description"`
}

func readTemplate(name string) string {
	f, _ := templates.ReadFile(name)
	return string(f)
}

func splitLines(path string) []string {
	return strings.Split(readTemplate(path), "\n")
}

func combineSnippets(snippets []Snippet) []string {
	var output []string
	for _, snippet := range snippets {
		output = append(output, snippet.Body...)
	}
	return output
}

func createSnippet(name, prefix, description string) Snippet {
	n := strings.ToLower(name)
	filename := fmt.Sprintf("templates/%s.md", n)
	return Snippet{
		Key:         n,
		Name:        name,
		Prefix:      prefix,
		Body:        splitLines(filename),
		Description: description,
	}
}

func main() {
	snippets := []Snippet{
		createSnippet("Title", "title", "Title and Description"),
		createSnippet("Acknowledgements", "ack", "Acknowledgements"),
		createSnippet("API Reference", "api", "API Reference"),
		createSnippet("Appendix", "apex", "Appendix"),
		createSnippet("Authors", "auth", "Authors"),
		createSnippet("Badges", "bad", "Badges"),
		createSnippet("Contributing", "con", "Contributing"),
		createSnippet("Demo", "demo", "Demo"),
		createSnippet("Used By", "used", "Used By"),
		createSnippet("Usage/Examples", "usa", "Usage/Examples"),
		createSnippet("Running Tests", "test", "Running Tests"),
		createSnippet("Tech Stack", "tech", "Tech Stack"),
		createSnippet("Support", "sup", "Support"),
		createSnippet("Screenshots", "screen", "Screenshots"),
		createSnippet("Run Local", "local", "Running Locally"),
		createSnippet("Roadmap", "road", "Roadmap"),
		createSnippet("Related", "rel", "Related"),
		createSnippet("Optimizations", "opt", "Optimizations"),
		createSnippet("Logo", "logo", "Logo"),
		createSnippet("License", "lic", "License"),
		createSnippet("Lessons", "les", "Lessons"),
		createSnippet("Installation", "inst", "Installation"),
		createSnippet("Feedback", "feed", "Feedback"),
		createSnippet("Features", "feat", "Features"),
		createSnippet("FAQ", "faq", "FAQ"),
		createSnippet("Environment Variables", "env", "Environment Variables"),
		createSnippet("Documentation", "doc", "Documentation"),
		createSnippet("Deployment", "deploy", "Deployment"),
	}

	data := make(map[string]Snippet)
	for _, snippet := range snippets {
		key := snippet.Key
		data[key] = snippet
	}
	// combine them all together and create a `readme` snippet
	data["readme"] = Snippet{
		Key:         "readme",
		Prefix:      "readme",
		Body:        combineSnippets(snippets),
		Description: "Example README with all the sections",
	}
	fmt.Println("Writing out the following sections:")
	for _, s := range snippets {
		fmt.Println(s.Name)
	}
	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("snippets.json", b, 0644)
	if err != nil {
		panic(err)
	}
}
