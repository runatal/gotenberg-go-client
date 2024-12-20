package gotenberg

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
	"github.com/starwalkn/gotenberg-go-client/v8/test"
)

func TestMarkdown(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	index, err := document.FromPath("index.html", test.MarkdownTestFilePath(t, "index.html"))
	require.NoError(t, err)
	markdown1, err := document.FromPath("paragraph1.md", test.MarkdownTestFilePath(t, "paragraph1.md"))
	require.NoError(t, err)
	markdown2, err := document.FromPath("paragraph2.md", test.MarkdownTestFilePath(t, "paragraph2.md"))
	require.NoError(t, err)
	markdown3, err := document.FromPath("paragraph3.md", test.MarkdownTestFilePath(t, "paragraph3.md"))
	require.NoError(t, err)
	req := NewMarkdownRequest(index, markdown1, markdown2, markdown3)
	req.Trace("testMarkdown")
	req.UseBasicAuth("foo", "bar")

	err = req.ExtraHTTPHeaders(map[string]string{
		"X-Header":        "Value",
		"X-Scoped-Header": `value;scope=https?:\\/\\/([a-zA-Z0-9-]+\\.)*domain\\.com\\/.*`,
	})
	require.NoError(t, err)

	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
}

func TestMarkdownComplete(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	index, err := document.FromPath("index.html", test.MarkdownTestFilePath(t, "index.html"))
	require.NoError(t, err)
	markdown1, err := document.FromPath("paragraph1.md", test.MarkdownTestFilePath(t, "paragraph1.md"))
	require.NoError(t, err)
	markdown2, err := document.FromPath("paragraph2.md", test.MarkdownTestFilePath(t, "paragraph2.md"))
	require.NoError(t, err)
	markdown3, err := document.FromPath("paragraph3.md", test.MarkdownTestFilePath(t, "paragraph3.md"))
	require.NoError(t, err)
	req := NewMarkdownRequest(index, markdown1, markdown2, markdown3)
	req.Trace("testMarkdownComplete")
	req.UseBasicAuth("foo", "bar")

	err = req.ExtraHTTPHeaders(map[string]string{
		"X-Header":        "Value",
		"X-Scoped-Header": `value;scope=https?:\\/\\/([a-zA-Z0-9-]+\\.)*domain\\.com\\/.*`,
	})
	require.NoError(t, err)

	header, err := document.FromPath("header.html", test.MarkdownTestFilePath(t, "header.html"))
	require.NoError(t, err)
	req.Header(header)
	footer, err := document.FromPath("footer.html", test.MarkdownTestFilePath(t, "footer.html"))
	require.NoError(t, err)
	req.Footer(footer)
	font, err := document.FromPath("font.woff", test.MarkdownTestFilePath(t, "font.woff"))
	require.NoError(t, err)
	img, err := document.FromPath("img.gif", test.MarkdownTestFilePath(t, "img.gif"))
	require.NoError(t, err)
	style, err := document.FromPath("style.css", test.MarkdownTestFilePath(t, "style.css"))
	require.NoError(t, err)
	req.Assets(font, img, style)
	req.OutputFilename("foo.pdf")
	req.WaitDelay(1 * time.Second)
	req.PaperSize(A4)
	req.Margins(NormalMargins)
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
}

func TestMarkdownPageRanges(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	index, err := document.FromPath("index.html", test.MarkdownTestFilePath(t, "index.html"))
	require.NoError(t, err)
	markdown1, err := document.FromPath("paragraph1.md", test.MarkdownTestFilePath(t, "paragraph1.md"))
	require.NoError(t, err)
	markdown2, err := document.FromPath("paragraph2.md", test.MarkdownTestFilePath(t, "paragraph2.md"))
	require.NoError(t, err)
	markdown3, err := document.FromPath("paragraph3.md", test.MarkdownTestFilePath(t, "paragraph3.md"))
	require.NoError(t, err)
	req := NewMarkdownRequest(index, markdown1, markdown2, markdown3)
	req.Trace("testMarkdownPageRanges")
	req.UseBasicAuth("foo", "bar")

	err = req.ExtraHTTPHeaders(map[string]string{
		"X-Header":        "Value",
		"X-Scoped-Header": `value;scope=https?:\\/\\/([a-zA-Z0-9-]+\\.)*domain\\.com\\/.*`,
	})
	require.NoError(t, err)

	req.NativePageRanges("1-1")
	resp, err := c.Send(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestMarkdownScreenshot(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	index, err := document.FromPath("index.html", test.MarkdownTestFilePath(t, "index.html"))
	require.NoError(t, err)
	markdown1, err := document.FromPath("paragraph1.md", test.MarkdownTestFilePath(t, "paragraph1.md"))
	require.NoError(t, err)
	markdown2, err := document.FromPath("paragraph2.md", test.MarkdownTestFilePath(t, "paragraph2.md"))
	require.NoError(t, err)
	markdown3, err := document.FromPath("paragraph3.md", test.MarkdownTestFilePath(t, "paragraph3.md"))
	require.NoError(t, err)
	req := NewMarkdownRequest(index, markdown1, markdown2, markdown3)
	req.Trace("testMarkdownScreenshot")
	req.UseBasicAuth("foo", "bar")

	err = req.ExtraHTTPHeaders(map[string]string{
		"X-Header":        "Value",
		"X-Scoped-Header": `value;scope=https?:\\/\\/([a-zA-Z0-9-]+\\.)*domain\\.com\\/.*`,
	})
	require.NoError(t, err)

	require.NoError(t, err)
	dirPath := t.TempDir()
	req.Format(JPEG)
	dest := fmt.Sprintf("%s/foo.jpeg", dirPath)
	err = c.StoreScreenshot(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
}
