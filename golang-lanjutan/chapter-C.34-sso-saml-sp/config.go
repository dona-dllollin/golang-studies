package main

import "fmt"

var (
	samlCertificatePath = "./myservice.cert"
	samlPrivateKeyPath  = "./myservice.key"
	samlIDPMetadata     = "https://idp.ssocircle.com/meta-idp.xml"
	webserverPort       = 9000
	webserverRootURL    = fmt.Sprintf("http://localhost:%d", webserverPort)
)
