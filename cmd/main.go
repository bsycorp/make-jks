package main

import (
	"crypto/tls"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"flag"
	"github.com/pavel-v-chernykh/keystore-go"
	"time"
	"log"
	"os"
)

func main() {
	pemPath := flag.String("input", "/etc/ssl/pemCertificates", "The path to find PEM encoded certificates to convert")
	jksPath := flag.String("output", "./cacerts.jks", "The path to output the Java keystore to")
	flag.Parse()

	pemCertificates, err := LoadCertificateDirectory(string(*pemPath))
	if err != nil {
		fmt.Errorf("Error loading certificates %s", err)
	}

	fmt.Printf("Building keystore..\n")
	keyStore := make(map[string]interface{})
	certCount := 1
	for pemIndex := 0; pemIndex < len(pemCertificates); pemIndex += 1 {
		pemCertificate := pemCertificates[pemIndex].Certificate
		for certIndex := 0; certIndex < len(pemCertificate); certIndex += 1 {
			keyStore[fmt.Sprintf("cacerts-%d", certCount)] = &keystore.TrustedCertificateEntry{
				Entry: keystore.Entry{
					CreationDate: time.Now(),
				},
				Certificate: keystore.Certificate{
					Content: pemCertificate[certIndex],
					Type:    "X509",
				},
			}
			certCount++
		}
	}

	password := []byte{'c', 'h', 'a', 'n', 'g', 'e', 'i', 't'}

	fmt.Printf("Writing keystore to %s\n", string(*jksPath))
	writeKeyStore(keyStore, string(*jksPath), password)
}

func writeKeyStore(keyStore keystore.KeyStore, filename string, password []byte) {
	o, err := os.Create(filename)
	defer o.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = keystore.Encode(o, keyStore, password)
	if err != nil {
		log.Fatal(err)
	}
}

// LoadCertficatesFromFile reads file, divides into key and certificates
func LoadCertficatesFromFile(path string) (tls.Certificate, error) {
	var cert tls.Certificate
	fmt.Printf("Loading certificates from file: %s\n", path)
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return cert, err
	}

	for {
		block, rest := pem.Decode(raw)
		if block == nil {
			break
		}
		if block.Type == "CERTIFICATE" {
			cert.Certificate = append(cert.Certificate, block.Bytes)
		}
		raw = rest
	}

	if len(cert.Certificate) == 0 {
		return cert, fmt.Errorf("No certificate found in \"%s\"\n", path)
	}

	return cert, nil
}

// LoadCertificateDirectory globs all .pem files in given directory, parses them
// for certs and returns them
func LoadCertificateDirectory(dir string) ([]tls.Certificate, error) {
	fmt.Printf("Checking %s for certificates..\n", dir)
	// read certificate files
	certficateFiles, err := filepath.Glob(filepath.Join(dir, "*.pem"))
	if err != nil {
		return nil, fmt.Errorf("Failed to scan certificate dir \"%s\": %s\n", dir, err)
	}
	fmt.Printf("Found %d files to check\n", len(certficateFiles))
	sort.Strings(certficateFiles)
	certs := make([]tls.Certificate, 0)
	for _, file := range certficateFiles {
		cert, err := LoadCertficatesFromFile(file)
		if err != nil {
			fmt.Printf("%s", err)
		} else {
			certs = append(certs, cert)
		}
	}
	fmt.Printf("Loaded %d certificate(s)\n", len(certs))
	return certs, nil
}
