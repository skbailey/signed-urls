## S3 Signed URLs

There is an example application for generating S3 signed URLs.
It also has an example for chaining credentials providers

```bash
AWS_ACCESS_KEY_ID="<aws-access-id>" \
AWS_SECRET_ACCESS_KEY="<aws-sercret-access-key>" \
go run cmd/s3signedurls/main.go
```

## Cloudfront Signed URLs

There is an example for generating Cloudfront signed URLS.

```bash
CF_DOMAIN="<cloudfront domain>" \
CF_PUBLIC_KEY_ID="<cloudfront public key id>" \
CF_PRIKEY_PATH="<path to private key>" \
go run cmd/cloudfrontsignedurls/main.go
```

## Cloudfront Signed Cookies

There is an example for generating Cloudfront signed URLS.

```bash
CF_DOMAIN="<cloudfront domain>" \
CF_PUBLIC_KEY_ID="<cloudfront public key id>" \
CF_PRIKEY_PATH="<path to private key>" \
go run cmd/cloudfrontsignedcookies/main.go
```
