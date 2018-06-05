package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/narenaryan/encryptService/helpers"
)

func main() {
	svc := helpers.EncryptServiceInstance{}
	encryptHandler := httptransport.NewServer(helpers.MakeEncryptEndpoint(svc),
		helpers.DecodeEncryptRequest,
		helpers.EncodeResponse)

	decryptHandler := httptransport.NewServer(helpers.MakeDecryptEndpoint(svc),
		helpers.DecodeDecryptRequest,
		helpers.EncodeResponse)

	http.Handle("/encrypt", encryptHandler)
	http.Handle("/decrypt", decryptHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


// curl -XPOST -d'{"key":"111023043350789514532147", "text": "I am A Message"}' localhost:8080/encrypt
// {"message":"8/+JCfTb+ibIjzQtmCo=","error":""}

// curl -XPOST -d'{"key":"111023043350789514532147", "message":"8/+JCfTb+ibIjzQtmCo="}' localhost:8080/decrypt
//  {"text":"I am A Message","error":""}

