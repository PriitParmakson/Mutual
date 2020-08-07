package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	// Marsruudid
	// Go-s "/" käsitleb ka need teed, millele oma käsitlejat ei leidu.
	http.HandleFunc("/", healthCheck)

	// Loe kliendi CA sert
	caCert, err := ioutil.ReadFile("rootCA.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConf := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	s := &http.Server{
		Addr: ":8080",
		// http.DefaultServeMux if nil
		Handler:   nil,
		TLSConfig: tlsConf,
		// Päringu lugemise maks. aeg
		ReadTimeout: 10 * time.Second,
		// Vastuse kirjutamise maks. aeg
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("** POC-ClientAuth server käivitatud pordil %v **\n", "8080")
	err2 := s.ListenAndServeTLS("https.cert", "https.key")
	if err2 != nil {
		log.Fatal(err2)
	}

}

// healthCheck pakub elutukset (/health).
func healthCheck(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"name":"POC-ClientAuth", "status":"ok"}`)
}

/* Vt: https://github.com/jcbsmpsn/golang-https-example/blob/master/https_server.go
 */
