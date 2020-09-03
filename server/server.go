package server

import (
	"crypto/tls"
	"fmt"
	"github.com/ovadiaK/bingo/router"
	"golang.org/x/crypto/acme/autocert"
	"os"
	"strings"

	"log"
	"net/http"
	"time"
)

var (
	domain                 = os.Getenv("DOMAIN")
	httpPort               = os.Getenv("HTTP_PORT")
	production             = os.Getenv("PRODUCTION")
	redirect               = os.Getenv("REDIRECT")
	dataDir                = os.Getenv("TLS_DIR")
	flgProduction          = false
	flgRedirectHTTPToHTTPS = false
)

func init() {
	if strings.ToLower(production) == "yes" {
		flgProduction = true
		if domain == "" {
			panic("no domain provided in environment, required")
		}
	}
	if strings.ToLower(redirect) == "yes" {
		flgRedirectHTTPToHTTPS = true
	}
}

func makeServerFromMux(mux http.Handler) *http.Server {
	// collection timeouts so that a slow or malicious client doesn't
	// hold resources forever
	return &http.Server{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  360 * time.Second,
		Handler:      mux,
	}
}

func makeHTTPServer() *http.Server {

	mux := router.New()
	return makeServerFromMux(mux)
}

func makeHTTPToHTTPSRedirectServer() *http.Server {
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		newURI := "https://" + r.Host + r.URL.String()
		http.Redirect(w, r, newURI, http.StatusFound)
	}
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handleRedirect)
	return makeServerFromMux(mux)
}

func Start() {
	var m *autocert.Manager

	var httpsSrv *http.Server
	if flgProduction {
		fmt.Println("production flag collection")
		m = &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(domain, "www."+domain),
			Cache:      autocert.DirCache(dataDir),
		}

		httpsSrv = makeHTTPServer()
		httpsSrv.Addr = domain + ":443"
		httpsSrv.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}

		go func() {
			fmt.Printf("Starting HTTPS server on %s\n", httpsSrv.Addr)
			err := httpsSrv.ListenAndServeTLS("", "")
			if err != nil {
				log.Fatalf("httpsSrv.ListendAndServeTLS() failed with %s", err)
			}
		}()
	}

	var httpSrv *http.Server
	if flgRedirectHTTPToHTTPS {
		fmt.Println("redirect flag collection")
		httpSrv = makeHTTPToHTTPSRedirectServer()
	} else {
		httpSrv = makeHTTPServer()
	}
	// allow autocert handle Let's Encrypt callbacks over http
	if m != nil {
		httpSrv.Handler = m.HTTPHandler(httpSrv.Handler)
	}

	httpSrv.Addr = httpPort
	fmt.Printf("Starting HTTP server on %s\n", httpPort)
	err := httpSrv.ListenAndServe()
	if err != nil {
		log.Fatalf("httpSrv.ListenAndServe() failed with %s", err)
	}
}
