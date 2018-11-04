//
// A simple http server that either server ACME http-01 request
// or redirect to https
//
// ./acme-helper /path/to/.well-known
//
package main

import "net/http"
import "fmt"
import "os"
import "strings"

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s /path/to/.well-known\n", os.Args[0])
		os.Exit(1)
	}

	acme_prefix := "/.well-known/"
	http_root := os.Args[1]
	fmt.Printf("Serving ACME challenges from: %s\n", http_root)

	// Handle ACME challenges by serving the folder.
	// XXX: Disable directory listing.
	http.Handle(acme_prefix,
		http.StripPrefix(acme_prefix, http.FileServer(http.Dir(http_root))))
	// Everything else is redirected to https
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s: %s %s\n", r.Host, r.Method, r.URL)
		host := strings.Split(r.Host, ":")
		redirect_uri := fmt.Sprintf("https://%s%s", host[0], r.URL)
		http.Redirect(w, r, redirect_uri, 301)
	})

	fmt.Println(http.ListenAndServe(":80", nil))
}
