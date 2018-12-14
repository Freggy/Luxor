package rest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/valyala/fasthttp"
)

// This interface provides abstraction for different signature creating algorithms like HMAC.
// It should be implemented by every type that represents a signature for authentication.
type Signature interface {
	// Verifies the given signature string.
	Verify(signature string) (bool, error)

	// Creates a signature signed with the given secret key.
	Create(key string) (string, error)
}

// Creates a Signature for a response message with the given RequestCtx.
// The algorithm can be chosen by passing the specific supplier to the method.
//  sig := rest.CreateForResponse(NewHmacSignatureForResponse, ctx)
func CreateForResponse(supplier func(ctx *fasthttp.RequestCtx) Signature, ctx *fasthttp.RequestCtx) Signature {
	return supplier(ctx)
}

// Creates a Signature for a request with the given RequestCtx.
// The algorithm can be chosen by passing the specific supplier to the method.
//  sig := rest.CreateFromRequest(NewHmacSignatureFromRequest, ctx)
func CreateFromRequest(supplier func(ctx *fasthttp.RequestCtx) Signature, ctx *fasthttp.RequestCtx) Signature {
	return supplier(ctx)
}

// Represents the HMAC SHA256 signature struct.
// This struct holds all information that will be used for creating the signature.
// It implements the Signature interface by providing
//  func Verify(signature string) (bool, error)
// and
//  func Create(key string) (string, error)
type HmacSignature struct {
	httpVerb    string
	endpoint    string
	contentType string
	date        string
	body        []byte
	signature   string
}

// NewHmacSignatureForResponse creates a type that uses HMAC sha256 for creating a signature
// Internally this method uses ctx.Response to retrieve header and body data.
func NewHmacSignatureForResponse(ctx *fasthttp.RequestCtx) Signature {
	return &HmacSignature{
		string(ctx.Method()),
		string(ctx.Path()),
		string(ctx.Response.Header.ContentType()),
		string(ctx.Response.Header.Peek("Date")),
		ctx.Response.Body(),
		"",
	}
}

// Creates a type that uses HMAC sha256 for creating a signature
// Internally this method uses ctx.Request to retrieve header and body data.
func NewHmacSignatureFromRequest(ctx *fasthttp.RequestCtx) Signature {
	return &HmacSignature{
		string(ctx.Method()),
		string(ctx.Path()),
		string(ctx.Request.Header.ContentType()),
		string(ctx.Request.Header.Peek("Date")),
		ctx.Response.Body(),
		"",
	}
}

// String concatenates all signature details. This string will later be used for creating the HMAC signature.
func (info *HmacSignature) String() string {
	return info.httpVerb + info.endpoint + info.contentType /*+ info.date*/ + string(info.body)
}

// Create creates the HMAC SHA256 signature with the given secret key.
// The result will then encoded to a base64 string.
func (info *HmacSignature) Create(key string) (string, error) {
	mac := hmac.New(sha256.New, []byte(key))

	if _, err := mac.Write([]byte(info.String())); err != nil {
		return "", err
	}

	info.signature = base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return info.signature, nil
}

// Verify verifies the given HMAC SHA256 base64 encoded string.
// The method will return whether or not the given signature matches.
func (info *HmacSignature) Verify(signature string) (bool, error) {
	other, err := base64.StdEncoding.DecodeString(signature)

	if err != nil {
		return false, err
	}

	self, err := base64.StdEncoding.DecodeString(info.signature)

	if err != nil {
		return false, err
	}

	return hmac.Equal(other, self), nil
}
