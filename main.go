package main

import (
	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	query string
	basePath = "/Users/viethung/Library/Mobile Documents/iCloud~md~obsidian/Documents/Vault/Attachment/csv/"
	repo = "viethung0823/alfred-web-searches"
	wf *aw.Workflow
 icons = getIcons()
)

type IconInfo struct {
    Icon    *aw.Icon
    Strings []string
}

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func run() {
	doSearch()
}

func getIcons() []IconInfo  {
    return []IconInfo {
      {&aw.Icon{Value: "icons/update-available.png"}, nil},
			{&aw.Icon{Value: "icons/reddit.png"}, []string{"r: ", "reddit"}},
			{&aw.Icon{Value: "icons/github.png"}, []string{"g: ", "github"}},
			{&aw.Icon{Value: "icons/translate.png"}, []string{"language"}},
			{&aw.Icon{Value: "icons/forums.png"}, []string{"forum"}},
			{&aw.Icon{Value: "icons/stack.png"}, []string{"s: "}},
			{&aw.Icon{Value: "icons/doc.png"}, []string{"d: "}},
			{&aw.Icon{Value: "icons/facebook.png"}, []string{"f: ", "facebook"}},
			{&aw.Icon{Value: "icons/music.png"}, []string{"music"}},
			{&aw.Icon{Value: "icons/instagram.png"}, []string{"instagram"}},
			{&aw.Icon{Value: "icons/youtube.png"}, []string{"youtube"}},
			{&aw.Icon{Value: "icons/linkedin.png"}, []string{"linkedin"}},
			{&aw.Icon{Value: "icons/utils.png"}, []string{"utils"}},
			{&aw.Icon{Value: "icons/twitter.png"}, []string{"twitter"}},
			{&aw.Icon{Value: "icons/news.png"}, []string{"news"}},
			{&aw.Icon{Value: "icons/stackoverflow.png"}, []string{"stackoverflow"}},
			{&aw.Icon{Value: "icons/tiktok.png"}, []string{"tiktok"}},
    }
}

func main() {
	wf.Run(run)
}
