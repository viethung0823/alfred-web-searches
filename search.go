package main

import (
	"encoding/csv"
	"log"
	"os"
	"regexp"
	"strings"
)

type Link struct {
    Value    string
    Subtitle string
    Tags     string
}

func parseCSV(files []string) map[string]Link {
    links := make(map[string]Link)

    for _, file := range files {
        // Load file
        f, err := os.Open(file)
        if err != nil {
            panic(err)
        }

        r := csv.NewReader(f)

        records, err := r.ReadAll()
        if err != nil {
            log.Fatal(err)
        }
        f.Close()

        // Skip if the record does not contain a valid URL
        for _, record := range records {
            if strings.TrimSpace(record[1]) == "" {
                continue
            }
            links[record[0]] = Link{Value: record[1], Subtitle: record[2], Tags: record[3]}
        }
    }

    return links
}

// doSearch searches through the websites and returns results to Alfred.
func doSearch() error {
	showUpdateStatus()

	log.Printf("query=%s", query)

	files := []string{basePath + "websites.csv", basePath + "github.csv"}
  links := parseCSV(files)

	var re1 = regexp.MustCompile(`.: `)

	for key, link := range links {
    item := wf.NewItem(key).
    Valid(true).
    Var("URL", link.Value).
    Var("ARG", re1.ReplaceAllString(key, ``)).
    UID(key).
    Subtitle(strings.TrimSpace(link.Subtitle)).
    Match(key + " " + link.Tags)
    for _, iconInfo := range icons {
      for _, str := range iconInfo.Strings {
        if strings.Contains(strings.ToLower(key), strings.ToLower(str)) ||
        strings.Contains(strings.ToLower(link.Value), strings.ToLower(str)) ||
        strings.Contains(strings.ToLower(link.Tags), strings.ToLower(str)) {
            item.Icon(iconInfo.Icon)
            break
        }
      }
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
