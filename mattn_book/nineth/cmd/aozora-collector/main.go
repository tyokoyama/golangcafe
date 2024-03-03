package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Entry struct {
	AuthorID string
	Author   string
	TitleID  string
	Title    string
	SiteURL  string
	ZipURL   string
}

func findEntries(siteURL string) ([]Entry, error) {
	doc, err := goquery.NewDocument(siteURL)
	if err != nil {
		return nil, err
	}
	// 処理
	entries := []Entry{}
	pat := regexp.MustCompile(`.*/cards/([0-9]+)/card([0-9]+).html$`)
	doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
		token := pat.FindStringSubmatch(elem.AttrOr("href", ""))
		if len(token) != 3 {
			return
		}

		title := elem.Text()
		pageURL := fmt.Sprintf("https://www.aozora.gr.jp/cards/%s/card%s.html", token[1], token[2])
		author, zipURL := findAuthorAndZIP(pageURL)
		if zipURL != "" {
			entries = append(entries, Entry{
				AuthorID: token[1],
				Author:   author,
				TitleID:  token[2],
				Title:    title,
				SiteURL:  siteURL,
				ZipURL:   zipURL,
			})
		}
		println(zipURL)
	})
	return entries, nil
}

func findAuthorAndZIP(siteURL string) (string, string) {
	doc, err := goquery.NewDocument(siteURL)
	if err != nil {
		return "", ""
	}

	author := doc.Find("table[summary=作家データ] tr:nth-child(1) td:nth-child(2)").Text()

	zipURL := ""
	doc.Find("table.download a").Each(func(n int, elem *goquery.Selection) {
		href := elem.AttrOr("href", "")
		if strings.HasSuffix(href, ".zip") {
			zipURL = href
		}
	})

	// 処理
	if zipURL == "" {
		return author, ""
	}
	if strings.HasPrefix(zipURL, "http://") || strings.HasPrefix(zipURL, "https://") {
		return author, zipURL
	}

	u, err := url.Parse(siteURL)
	if err != nil {
		return author, ""
	}

	u.Path = path.Join(path.Dir(u.Path), zipURL)
	return author, u.String()
}

func extractText(zipURL string) (string, error) {
	// 処理
	resp, err := http.Get(zipURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return "", err
	}

	for _, file := range r.File {
		if path.Ext(file.Name) == ".txt" {
			f, err := file.Open()
			if err != nil {
				return "", err
			}

			b, err := io.ReadAll(f)
			f.Close()
			if err != nil {
				return "", err
			}

			return string(b), nil
		}
	}
	return "", errors.New("contents not found")
}

func main() {
	listURL := "https://www.aozora.gr.jp/index_pages/person879.html"

	entries, err := findEntries(listURL)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		content, err := extractText(entry.ZipURL)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(entry.SiteURL)
		fmt.Println(content)
	}
}
