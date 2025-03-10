package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/playwright-community/playwright-go"
)

var args = []string{
	"--disable-blink-features=AutomationControlled",
	"--start-maximized",
	"--disable-infobars",
	"--no-sandbox",
	"--enable-gpu",
	"--use-gl=desktop",
	"--enable-webgl",
	"--enable-accelerated-2d-canvas",
	"--autoplay-policy=no-user-gesture-required",
	"--disable-dev-shm-usage",
	"--disable-extensions",
	"--remote-debugging-port=0",
	"--disable-web-security",
	"--enable-features=WebRTCPeerConnectionWithBlockIceAddresses",
	"--force-webrtc-ip-handling-policy=disable_non_proxied_udp",
}

type AmazonVisit struct {
	Link string
}

func (av *AmazonVisit) VisitAmazon() string {
	fmt.Printf("[DEBUG] Visiting Amazon link with http: %s\n", av.Link)
	client := &http.Client{}
	req, err := http.NewRequest("GET", av.Link, nil)
	if err != nil {
		log.Println("[DEBUG] Error creating request:", err)
		return av.VisitAmazonWithPlaywright()
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("DNT", "1")

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Println("[DEBUG] Visiting Amazon link with http failed (Trying with Playwright)...")
		return av.VisitAmazonWithPlaywright()
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("[DEBUG] Error reading response body:", err)
		return av.VisitAmazonWithPlaywright()
	}

	htmlContent, _ := doc.Html()
	return htmlContent
}

func (av *AmazonVisit) VisitAmazonWithPlaywright() string {
	fmt.Printf("[DEBUG] Now visiting Amazon link with Playwright: %s\n", av.Link)
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
		Args:     args,
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	if _, err := page.Goto(av.Link); err != nil {
		log.Fatalf("could not go to %s: %v", av.Link, err)
	}

	if _, err := page.WaitForLoadState("load"); err != nil {
		log.Fatalf("could not wait for load state: %v", err)
	}

	content, err := page.Content()
	if err != nil {
		log.Fatalf("could not get page content: %v", err)
	}

	if err := browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}

	if err := pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

	return content
}

type AmazonSERPExtractor struct {
	AmazonVisit
	Keyword       string
	ProductNumber int
	SearchLink    string
}

func NewAmazonSERPExtractor(keyword string, productNumber int) *AmazonSERPExtractor {
	searchLink := fmt.Sprintf("https://www.amazon.com/s?k=%s", strings.ReplaceAll(keyword, " ", "+"))
	return &AmazonSERPExtractor{
		AmazonVisit:   AmazonVisit{Link: searchLink},
		Keyword:       keyword,
		ProductNumber: productNumber,
		SearchLink:    searchLink,
	}
}

func (ase *AmazonSERPExtractor) SearchAmazon() []string {
	htmlContent := ase.VisitAmazon()
	return ase.ExtractLinks(htmlContent)
}

func (ase *AmazonSERPExtractor) GetProductNumber() int {
	if ase.ProductNumber > 0 {
		return ase.ProductNumber
	}
	return 10
}

func (ase *AmazonSERPExtractor) ExtractLinks(htmlContent string) []string {
	fmt.Println("[DEBUG] Extracting links...")
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatalf("Error parsing HTML: %v", err)
	}

	var links []string
	doc.Find("div[data-component-type='s-search-result']").Each(func(i int, s *goquery.Selection) {
		aTag := s.Find("a.a-link-normal")
		href, exists := aTag.Attr("href")
		if exists && strings.Contains(href, "/dp/") {
			link := fmt.Sprintf("https://www.amazon.com%s", strings.Split(href, "ref")[0])
			links = append(links, link)
			if len(links) >= ase.GetProductNumber() {
				return
			}
		}
	})

	return links
}

type ProductDataExtractor struct {
	AmazonVisit
	LinkList    []string
	ProductData []map[string]string
}

func NewProductDataExtractor(linkList []string) *ProductDataExtractor {
	return &ProductDataExtractor{
		LinkList:    linkList,
		ProductData: []map[string]string{},
	}
}

func (pde *ProductDataExtractor) ExtractProductData() []map[string]string {
	for _, link := range pde.LinkList {
		fmt.Printf("[DEBUG] Extracting data from product link: %s\n", link)
		pde.Link = link
		htmlContent := pde.VisitAmazon()
		pde.ProductData = append(pde.ProductData, pde.ParseProductPage(htmlContent))
	}
	return pde.ProductData
}

func (pde *ProductDataExtractor) ParseProductPage(htmlContent string) map[string]string {
	fmt.Println("[DEBUG] Parsing product page...")
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatalf("Error parsing HTML: %v", err)
	}

	productTitle := doc.Find("#productTitle").Text()
	if productTitle == "" {
		productTitle = "NO Title"
	}
	price := doc.Find("span.a-price-whole").Text()
	if price == "" {
		price = "Price not found"
	}
	description := doc.Find("div#feature-bullets").Text()
	technicalData := doc.Find("table.a-keyvalue.prodDetTable").Text()
	review := doc.Find("div[data-hook='review-collapsed']").Text()

	dataOK := true
	if productTitle == "NO Title" {
		dataOK = false
	}

	return map[string]string{
		"data_ok":        fmt.Sprintf("%t", dataOK),
		"title":          strings.TrimSpace(productTitle),
		"price":          strings.TrimSpace(price),
		"description":    strings.TrimSpace(description),
		"technical_data": strings.TrimSpace(technicalData),
		"review":         strings.TrimSpace(review),
	}
}

func main() {
	keyword := "laptop hp"
	extractor := NewAmazonSERPExtractor(keyword, 10)
	links := extractor.SearchAmazon()
	for _, link := range links {
		fmt.Println(link)
	}

	productExtractor := NewProductDataExtractor(links)
	productData := productExtractor.ExtractProductData()
	for _, data := range productData {
		fmt.Println(data)
	}
}
