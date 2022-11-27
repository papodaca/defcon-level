package main

import (
	"bytes"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
)

func collectText(n *html.Node) string {
	buf := &bytes.Buffer{}
	collectTextInner(n, buf)
	buf.WriteString(";")
	str, err := buf.ReadString(';')
	if err != nil {
		panic(err)
	}
	return str
}

func collectTextInner(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectTextInner(c, buf)
	}
}

func get_level() string {
	resp, err := http.Get("https://www.defconlevel.com/current-level.php")
	if err != nil || resp.StatusCode != 200 {
		panic(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	} // /Defcon Level:\s+(\d)/.exec($('.header-defcon-level').textContent)[1]

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	title := doc.Find(".header-defcon-level").Nodes[0]
	text_content := collectText(title)
	r := regexp.MustCompile(`Defcon Level:\s(\d+)`)
	matched := r.FindStringSubmatch(text_content)

	if len(matched) >= 2 {
		return matched[1]
	} else {
		return ""
	}
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, get_level())
	})
	r.Run()
}
