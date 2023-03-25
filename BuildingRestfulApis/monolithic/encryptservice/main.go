package main

import (
	"encryptservice/helpers"
	transport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

func main() {
	svc := helpers.EncryptServiceInstance{}
	encryptHandler := transport.NewServer(helpers.MakeEncryptEndpoint(svc), helpers.DecodeEncryptRequest, helpers.EncodeResponse)
	decryptHandler := transport.NewServer(helpers.MakeDecryptEndpoint(svc), helpers.DecodeDecryptRequest, helpers.EncodeResponse)

	http.Handle("/encrypt", encryptHandler)
	http.Handle("/decrypt", decryptHandler)
	log.Fatal(http.ListenAndServe(":1205", nil))
}
