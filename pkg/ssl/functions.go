package ssl

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"io/ioutil"
	"os"
	"strings"
)

func BuildTLS() (*tls.Config, error) {

	ssl := os.Getenv("SSL_DIR")

	caCert, err := ioutil.ReadFile(strings.Join([]string{ssl, "ca_certificate.pem"}, "/"))
	if err != nil {
		log.Fatal("Error leyendo CA certificate: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, err := tls.LoadX509KeyPair(strings.Join([]string{ssl, "client_ubuntu-us-southeast_certificate.pem"}, "/"), strings.Join([]string{ssl, "client_ubuntu-us-southeast_key.pem"}, "/"))
	if err != nil {
		log.Fatal("Error cargando client certificate/key pair: %v", err)
	}
	tlsConfig := &tls.Config{
		RootCAs:            caCertPool,
		Certificates:       []tls.Certificate{cert},
		ServerName:         "ubuntu-us-southeast",
		InsecureSkipVerify: false,
	}

	return tlsConfig, nil
}
