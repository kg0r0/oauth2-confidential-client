package main

import (
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kg0r0/oauth2-confidential-client/oauth2client"
	"golang.org/x/oauth2"
)

var tpl *template.Template
var conf *oauth2.Config
var ctx context.Context

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	ctx = context.Background()
	config, err := oauth2client.NewConfig("conf/config.json")
	if err != nil {
		log.Fatal(err)
	}

	endpoint := oauth2.Endpoint{
		AuthURL:  config.ClientConfig.Endpoint.AuthURL,
		TokenURL: config.ClientConfig.Endpoint.TokenURL,
	}

	conf = &oauth2.Config{
		ClientID:     config.ClientConfig.ClientID,
		ClientSecret: config.ClientConfig.ClientSecret,
		Scopes:       config.ClientConfig.Scopes,
		RedirectURL:  config.ClientConfig.RedirectURL,
		Endpoint:     endpoint,
	}

}

func AuthCodeHandler(w http.ResponseWriter, r *http.Request) {
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOnline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func AuthCodeCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	client := conf.Client(ctx, tok)
	resp, err := client.Get("...")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(w, "index.html", body)

}

func main() {
	http.HandleFunc("/", AuthCodeHandler)
	http.HandleFunc("/callback", AuthCodeCallbackHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
