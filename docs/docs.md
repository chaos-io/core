# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [chaos/http/cookie.proto](#chaos_http_cookie-proto)
    - [Cookie](#chaos-http-Cookie)
  
    - [SameSite](#chaos-http-SameSite)
  
- [chaos/http/header.proto](#chaos_http_header-proto)
    - [Header](#chaos-http-Header)
    - [Header.ValsEntry](#chaos-http-Header-ValsEntry)
  
- [chaos/http/method.proto](#chaos_http_method-proto)
    - [Method](#chaos-http-Method)
  
- [chaos/http/version.proto](#chaos_http_version-proto)
    - [Version](#chaos-http-Version)
  
- [chaos/http/request.proto](#chaos_http_request-proto)
    - [Request](#chaos-http-Request)
  
- [chaos/http/status.proto](#chaos_http_status-proto)
    - [Status](#chaos-http-Status)
  
- [Scalar Value Types](#scalar-value-types)



<a name="chaos_http_cookie-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaos/http/cookie.proto



<a name="chaos-http-Cookie"></a>

### Cookie
表示一个 HTTP Cookie（RFC 6265）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | cookie-name |
| value | [string](#string) |  | cookie-value |
| expires | [chaos.core.Timestamp](#chaos-core-Timestamp) |  | Expires attribute (HTTP-date) 例如：Wed, 21 Oct 2015 07:28:00 GMT |
| max_age | [int32](#int32) |  | Max-Age attribute (seconds) |
| domain | [string](#string) |  | Domain attribute |
| path | [string](#string) |  | Path attribute |
| secure | [bool](#bool) |  | Secure attribute |
| http_only | [bool](#bool) |  | HttpOnly attribute |
| same_site | [SameSite](#chaos-http-SameSite) |  | SameSite attribute |





 


<a name="chaos-http-SameSite"></a>

### SameSite
SameSite 枚举（RFC 6265bis / 实际浏览器行为）

| Name | Number | Description |
| ---- | ------ | ----------- |
| SAME_SITE_UNSPECIFIED | 0 |  |
| SAME_SITE_LAX | 1 |  |
| SAME_SITE_STRICT | 2 |  |
| SAME_SITE_NONE | 3 |  |


 

 

 



<a name="chaos_http_header-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaos/http/header.proto



<a name="chaos-http-Header"></a>

### Header



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Vals | [Header.ValsEntry](#chaos-http-Header-ValsEntry) | repeated |  |






<a name="chaos-http-Header-ValsEntry"></a>

### Header.ValsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [chaos.core.StringValues](#chaos-core-StringValues) |  |  |





 

 

 

 



<a name="chaos_http_method-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaos/http/method.proto


 


<a name="chaos-http-Method"></a>

### Method


| Name | Number | Description |
| ---- | ------ | ----------- |
| METHOD_UNSPECIFIED | 0 |  |
| METHOD_GET | 1 |  |
| METHOD_HEAD | 2 |  |
| METHOD_POST | 3 |  |
| METHOD_PUT | 4 |  |
| METHOD_DELETE | 5 |  |
| METHOD_CONNECT | 6 |  |
| METHOD_OPTIONS | 7 |  |
| METHOD_TRACE | 8 |  |


 

 

 



<a name="chaos_http_version-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaos/http/version.proto



<a name="chaos-http-Version"></a>

### Version



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| major | [int32](#int32) |  |  |
| minor | [int32](#int32) |  |  |





 

 

 

 



<a name="chaos_http_request-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaos/http/request.proto



<a name="chaos-http-Request"></a>

### Request



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| method | [Method](#chaos-http-Method) |  |  |
| url | [chaos.core.Url](#chaos-core-Url) |  |  |
| version | [Version](#chaos-http-Version) |  |  |
| header | [Header](#chaos-http-Header) |  |  |
| body | [chaos.core.Value](#chaos-core-Value) |  |  |





 

 

 

 



<a name="chaos_http_status-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaos/http/status.proto



<a name="chaos-http-Status"></a>

### Status



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int32](#int32) |  |  |
| reason | [string](#string) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

