package main

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/cloudfront/sign"
)

var (
	objKey string = "fish.jpeg"
)

var (
	cfDomain     = os.Getenv("CF_DOMAIN")
	cfAccessKey  = os.Getenv("CF_PUBLIC_KEY_ID")
	cfPrikeyPath = os.Getenv("CF_PRIKEY_PATH")
)

func main() {
	priKeyFile, err := os.Open(cfPrikeyPath)
	if err != nil {
		log.Fatalln((err))
	}

	var priKey *rsa.PrivateKey
	priKey, err = sign.LoadPEMPrivKey(priKeyFile)
	if err != nil {
		log.Fatalln((err))
	}

	var signedURL string
	signer := sign.NewURLSigner(cfAccessKey, priKey)

	rawURL := url.URL{
		Scheme: "https",
		Host:   cfDomain,
		Path:   objKey,
	}
	signedURL, err = signer.Sign(rawURL.String(), time.Now().Add(1*time.Hour))
	if err != nil {
		log.Fatalln((err))
	}
	fmt.Printf("Get signed URL %q\n", signedURL)
}
