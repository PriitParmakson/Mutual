package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	cFilePtr := flag.String("conf", "config.json", "Seadistusfaili asukoht")
	flag.Parse()

	// Loe seadistus sisse
	conf = loadConf(*cFilePtr)

	// Loe kontrolli objektid sisse
	VObjects := loadVObjects(conf.VObjectsFile)

	// Valmista HTTPS klient

	// Loe kliendi HTTPS võti ja sert
	cert, err := tls.LoadX509KeyPair(
		conf.ClientCert,
		conf.ClientKey)
	if err != nil {
		log.Fatal(err)
	}

	// Loe CA sert
	caCert, err := ioutil.ReadFile(
		conf.RootCA,
	)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Sea HTTPS klient valmis
	// Vt: https://golang.org/pkg/net/http/
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	for _, vo := range VObjects.VObjects {
		fmt.Printf("--- Kontrollin masinat:\n    %s\n    %s\n", vo.Name, vo.URL)
		getObjectCert(client, vo.Name, vo.URL)
	}
}

func getObjectCert(client *http.Client, name, url string) {
	r, err := client.Get(url)
	if err == nil {
		defer r.Body.Close()
		for _, cert := range r.TLS.PeerCertificates {
			fmt.Printf("    Issuer:\n    O=%v, CN=%v\n", cert.Issuer.Organization,
				cert.Issuer.CommonName)
			fmt.Printf("    Subject:\n    O=%v, CN=%v\n", cert.Subject.Organization,
				cert.Subject.CommonName)
			fmt.Printf("    Subject Alternate Names:\n    %+v\n", cert.DNSNames)
		}
	} else {
		fmt.Printf("--- Kontrolliobjekt ei ole kättesaadav:\n    %s: %s\n", name, url)
		fmt.Printf("    %v\n", err)
	}
}

// https://stackoverflow.com/questions/12122159/how-to-do-a-https-request-with-bad-certificate
// https://serverfault.com/questions/661978/displaying-a-remote-ssl-certificate-details-using-cli-tools
