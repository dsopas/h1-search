package main

import (
    "io/ioutil"
    "log"
    "html/template"
    "net/http"
    "encoding/json"
    "os"
    "time"
)

type TemplateData struct{
    Target string
    Reports []Report
}

type Report struct {
    Title string
    Url string
    Severity_rating string
}

type Response struct {
    Count int
    Pages int
    Reports []Report
}

func main() {
    var templateData TemplateData
    var hacktivity *hacktivity
    var response Response

    log.Println("::: Loading h1-search.go...")
    log.Println("::: by @dsopas and @pauloasilva_com")
    log.Println(":::::::::::::::::::::::::::::::::::.")

    templateData = TemplateData{Target: os.Args[1]}
    hacktivity = NewHacktivity(os.Args[1])

    res, err := http.Get(hacktivity.String())
    if err != nil {
        log.Fatal(err)
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    jsonErr := json.Unmarshal(body, &response)
    if jsonErr != nil {
        log.Fatal(jsonErr)
    }

    log.Println("::: Creating template file...")

    tmpl := template.Must(template.ParseFiles("./layout.html"))

    log.Println("::: Getting results from HackerOne...")

    for p := 1; p <= response.Pages; p++ {
        res, err := http.Get(hacktivity.NextPage().String())
        if err != nil {
            log.Fatal(err)
        }

        body, err := ioutil.ReadAll(res.Body)
        if err != nil {
            log.Fatal(err)
        }

        jsonErr := json.Unmarshal(body, &response)
        if jsonErr != nil {
            log.Fatal(jsonErr)
        }

        log.Printf("::: Getting page %d of %d...\n", p, response.Pages)

        for _, r := range response.Reports {
            report := Report {
                Title: r.Title,
                Severity_rating: r.Severity_rating,
                Url: r.Url}

            templateData.Reports = append(templateData.Reports, report)
        }
        time.Sleep(time.Millisecond * 500)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, templateData)
    })

    log.Println("::: File is ready...")
    log.Println("::: Click to view the results: http://localhost:3000/")

    http.ListenAndServe(":3000", nil)
}
