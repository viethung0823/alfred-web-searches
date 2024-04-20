package main

import (
	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	// Icons
	updateAvailable = &aw.Icon{Value: "icons/update-available.png"}
	redditIcon      = &aw.Icon{Value: "icons/reddit.png"}
	githubIcon      = &aw.Icon{Value: "icons/github.png"}
	translateIcon   = &aw.Icon{Value: "icons/translate.png"}
	forumsIcon      = &aw.Icon{Value: "icons/forums.png"}
	stackIcon       = &aw.Icon{Value: "icons/stack.png"}
	docIcon         = &aw.Icon{Value: "icons/doc.png"}
	facebookIcon    = &aw.Icon{Value: "icons/facebook.png"}
	musicIcon    		= &aw.Icon{Value: "icons/music.png"}
	instagramIcon    		= &aw.Icon{Value: "icons/instagram.png"}
	youtubeIcon    		= &aw.Icon{Value: "icons/youtube.png"}

	query string

	repo = "viethung0823/alfred-web-searches"

	wf *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func run() {
	doSearch()
}

func main() {
	wf.Run(run)
}
