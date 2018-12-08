Rest Provider
=============

This is a library wrapping [`fasthttp`](https://github.com/valyala/fasthttp) and [`fasthttprouter`](https://github.com/buaazp/fasthttprouter).
It is used in all Luxor projects exposing a REST API. The library performs HMAC authentication and retrieves client information based on that.

HMAC Authentication
-------------------

Every client has unique API credentials, those consist of the following:
* public **`CLIENT_ID`**
* private **`CLIENT_SECRET`**

To create a valid API request the Authorization Header must be set:
```HTTP
Authentication: LUXOR CLIENT_ID:SIGNATURE
```

Here is an example
```HTTP
Authentication: LUXOR ff7b185f-4557-48db-8783-ca8b9f9435b0:Q0V5c21RRWRaUGRhTjZGbEF6ZUF1d0JXSnVkNUNHNWVsTDhkQUtXclU=
```

The **`SIGNATURE`** will be calculated as follows
```
BASE64
(
    HMAC_SHA256
    (
        CLIENT_SECRET, // Key
        (              // Attributes to sign
            HTTP-Verb,
            ENDPOINT,
            Content-Type,
            Date,
            Body 
        )
)
```

If the server could not verify the signature an error response will be sent back
```JSON
{"code": 400, "message": "Error while verifying signature", "type": "SignatureVerifyException"}
```

Error Response
--------------

If the request wasn't valid you will receive an error response in JSON format.

| Field   | Type    | Description                           |
|---------|---------|---------------------------------------|
| code    | Integer |  HTTP error code                      |
| message | String  | Description of the error              |
| type    | String  | Type of the error i.e `AuthException` |

### Example

```JSON
{"code": 400, "message": "My Error Message", "type": "MyErrorType"}
```

Usage
-----

**TODO**
