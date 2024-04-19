package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// TODO:
func updateList() {
	// Get websites list from GitHub
	res, err := http.Get("https://raw.githubusercontent.com/viethung0823/alfred-web-searches/master/workflow/websites.csv")
	if err != nil {
		return
	}

	defer res.Body.Close()
	// body, err := ioutil.ReadAll(res.Body)
	// TODO: Cache it
}

// parseCSV parses CSV for links and arguments.
type Link struct {
    Value    string
    Subtitle string
    Tags     string
}

func parseCSV() map[string]Link {
    var err error

    // Load file
    f, err := os.Open("/Users/viethung/Library/Mobile Documents/iCloud~md~obsidian/Documents/Vault/Attachment/websites.csv")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    r := csv.NewReader(f)

    records, err := r.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    // Holds user's search arguments and an appropriate search URL
    links := make(map[string]Link)
    // Skip if the record does not contain a valid URL
    for _, record := range records {
        if strings.TrimSpace(record[1]) == "" {
            continue
        }
        links[record[0]] = Link{Value: record[1], Subtitle: record[2], Tags: record[3]}
    }

    return links
}

// doSearch searches through the websites and returns results to Alfred.
func doSearch() error {
	showUpdateStatus()

	log.Printf("query=%s", query)

	links := parseCSV()

	var re1 = regexp.MustCompile(`.: `)

	for key, link := range links {
    if strings.Contains(key, "r: ") {
        wf.NewItem(key).Valid(true).Var("URL", link.Value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Subtitle(link.Subtitle).Match(key + " " + link.Tags).Icon(redditIcon)
    } else if strings.Contains(key, "d: ") {
        wf.NewItem(key).Valid(true).Var("URL", link.Value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Subtitle(strings.TrimSpace(link.Subtitle)).Match(key + " " + link.Tags).Icon(docIcon)
    } else if strings.Contains(key, "g: ") {
        wf.NewItem(key).Valid(true).Var("URL", link.Value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Subtitle(strings.TrimSpace(link.Subtitle)).Match(key + " " + link.Tags).Icon(githubIcon)
    } else if strings.Contains(key, "s: ") {
        wf.NewItem(key).Valid(true).Var("URL", link.Value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Subtitle(strings.TrimSpace(link.Subtitle)).Match(key + " " + link.Tags).Icon(stackIcon)
    } else if strings.Contains(key, "f: ") {
        wf.NewItem(key).Valid(true).Var("URL", link.Value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Subtitle(strings.TrimSpace(link.Subtitle)).Match(key + " " + link.Tags).Icon(forumsIcon)
    } else if strings.Contains(key, "t: ") {
        wf.NewItem(key).Valid(true).Var("URL", link.Value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Subtitle(strings.TrimSpace(link.Subtitle)).Match(key + " " + link.Tags).Icon(translateIcon)
    } else {
        wf.NewItem(key).Valid(true).Var("URL", link.Value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Subtitle(strings.TrimSpace(link.Subtitle)).Match(key + " " + link.Tags)
    }
}

	query = os.Args[1]

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()
	return nil
}
