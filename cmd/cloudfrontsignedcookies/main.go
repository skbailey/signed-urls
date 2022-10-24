package main

import (
	"crypto/rsa"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
	"github.com/labstack/echo/v4"
)

var (
	cfDomain     = os.Getenv("CF_DOMAIN")
	cfPrikeyPath = os.Getenv("CF_PRIKEY_PATH")
	cfAccessKey  = os.Getenv("CF_PUBLIC_KEY_ID")
)

func main() {
	e := echo.New()
	e.GET("/cloudfront", func(c echo.Context) error {
		priKeyFile, err := os.Open(cfPrikeyPath)
		if err != nil {
			log.Fatalln((err))
		}

		var priKey *rsa.PrivateKey
		priKey, err = sign.LoadPEMPrivKey(priKeyFile)
		if err != nil {
			log.Fatalln((err))
		}

		// generate signed cookies
		expireAt := time.Now().Add(24 * time.Hour)
		cookieSigner := sign.NewCookieSigner(cfAccessKey, priKey)
		signedCookies, err := cookieSigner.Sign("https://", expireAt, func(o *sign.CookieOptions) {
			o.Path = "/"
			o.Domain = cfDomain
		})
		if err != nil {
			log.Fatal("generate signed cookie failed:", err)
		}

		for _, cookie := range signedCookies {
			c.SetCookie(cookie)
		}

		return c.String(http.StatusOK, "Hello, World!")
	})

	if err := e.Start(":8081"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
