package ssl

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

func TLS(serverName string, caCertificate string, clientCertificate string, clientKey string) (*tls.Config, error) {
	assert.NotEmpty(serverName, "ssl - error setting up tls: serverName is empty")
	assert.NotEmpty(caCertificate, "ssl - error setting up tls: caCertificate is empty")
	assert.NotEmpty(clientCertificate, "ssl - error setting up tls: clientCertificate is empty")
	assert.NotEmpty(clientKey, "ssl - error setting up tls: clientKey is empty")

	var err error

	var caCert []byte
	if caCert, err = os.ReadFile(caCertificate); err != nil {
		return nil, fmt.Errorf("ca certificate: %s", err.Error())
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	var cert tls.Certificate
	if cert, err = tls.LoadX509KeyPair(clientCertificate, clientKey); err != nil {
		return nil, fmt.Errorf("client certificate: %s", err.Error())
	}

	tlsConfig := &tls.Config{
		RootCAs:            caCertPool,
		Certificates:       []tls.Certificate{cert},
		ServerName:         serverName,
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS12,
	}

	return tlsConfig, nil
}
