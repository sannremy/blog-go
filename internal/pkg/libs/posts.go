package libs

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gosimple/slug"
)

// PostTitles contains post titles
var PostTitles = make(map[string]string)

// PostDates contains post dates
var PostDates = make(map[string]time.Time)

// PostHTMLs contains post HTML contents
var PostHTMLs = make(map[string]string)

// PostSlugs contains post slugs
var PostSlugs []string

// InitPosts Initialize all posts
func InitPosts() {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	postsDir := "web/posts/"

	// Find all posts
	err := filepath.Walk(postsDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".md") {
			mdText, err := ioutil.ReadFile(path)

			if err != nil {
				log.Fatal(err)
			}

			md := []byte(mdText)

			// Post HTML
			html := string(markdown.ToHTML(md, nil, renderer))

			// Generate slug (first <h1> title from md)
			startIndexTitle := strings.Index(html, "<h1>") + 4
			endIndexTitle := strings.Index(html, "</h1>")
			title := html[startIndexTitle:endIndexTitle]

			// Post slug
			postSlug := slug.Make(title)
			PostSlugs = append(PostSlugs, postSlug)

			// Date (YYYY-MM-DD)
			postTime, errTime := time.Parse(time.RFC3339, path[len(postsDir):len(postsDir)+10]+"T00:00:01Z")
			PostDates[postSlug] = postTime

			if errTime != nil {
				log.Fatal(errTime)
			}

			// Post title
			PostTitles[postSlug] = title

			// Post HTML
			PostHTMLs[postSlug] = html
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// Reverse slugs
	for i, j := 0, len(PostSlugs)-1; i < j; i, j = i+1, j-1 {
		PostSlugs[i], PostSlugs[j] = PostSlugs[j], PostSlugs[i]
	}
}
