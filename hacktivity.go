package main

import  "net/url"
import "strconv"

const baseUrl string = "https://hackerone.com/hacktivity"

type hacktivity struct {
    url string
    filters url.Values
}

func (h *hacktivity) NextPage() *hacktivity {
    var currentPage string = h.filters.Get("page")

    nextPage, _ := strconv.Atoi(currentPage)
    h.filters.Set("page", strconv.Itoa(nextPage + 1))

    return h
}

func (h *hacktivity) String() string {
    url, err := url.Parse(h.url)
    if err != nil {
        panic("Failed to parse Hackitivy base url")
    }

    url.RawQuery = h.filters.Encode()

    return url.String()
}

// NewHacktivity : Prepare the results, including filters
func NewHacktivity(to string) *hacktivity {
    h := hacktivity{
        url: baseUrl, filters: url.Values{} }

        h.filters.Set("sort_type", "latest_disclosable_activity_at")
        h.filters.Set("filter", "type:public to:"+to)
        h.filters.Set("page", "1")
        h.filters.Set("range", "forever")
        h.filters.Set("format", "json")

        return &h
    }

