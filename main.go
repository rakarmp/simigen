package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

type SitemapIndex struct {
	XMLName xml.Name `xml:"sitemapindex"`
	Xmlns   string   `xml:"xmlns,attr"`
	Sitemap []Sitemap
}

type Sitemap struct {
	Loc     string    `xml:"loc"`
	LastMod time.Time `xml:"lastmod"`
}

func main() {
	urls := []string{
		"https://example.com/page1",
		"https://example.com/page2",
		"https://example.com/page3",
	}

	sitemap := GenerateSitemap(urls)
	if err := SaveSitemap(sitemap, "sitemap.xml"); err != nil {
		fmt.Println("Failed to save sitemap:", err)
		return
	}

	fmt.Println("Sitemap generated successfully!")
}

func GenerateSitemap(urls []string) SitemapIndex {
	var sitemapIndex SitemapIndex
	now := time.Now()

	for _, url := range urls {
		sitemap := Sitemap{
			Loc:     url,
			LastMod: now,
		}
		sitemapIndex.Sitemap = append(sitemapIndex.Sitemap, sitemap)
	}

	return sitemapIndex
}

func SaveSitemap(sitemap SitemapIndex, filename string) error {
	sitemapIndexXml, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		return err
	}

	xmlStr := xml.Header + string(sitemapIndexXml)
	if err := os.WriteFile(filename, []byte(xmlStr), os.ModePerm); err != nil {
		return err
	}

	return nil
}
