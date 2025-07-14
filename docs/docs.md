# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [core/error_code.proto](#core_error_code-proto)
    - [ErrorCode](#core-ErrorCode)
  
- [core/error.proto](#core_error-proto)
    - [Error](#core-Error)
  
- [core/null.proto](#core_null-proto)
    - [Null](#core-Null)
  
- [core/resource.proto](#core_resource-proto)
    - [Resource](#core-Resource)
  
- [core/value.proto](#core_value-proto)
    - [Object](#core-Object)
    - [Object.FieldsEntry](#core-Object-FieldsEntry)
    - [Value](#core-Value)
    - [Values](#core-Values)
  
    - [ValueKind](#core-ValueKind)
  
- [Scalar Value Types](#scalar-value-types)



<a name="core_error_code-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## core/error_code.proto



<a name="core-ErrorCode"></a>

### ErrorCode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| http_status_code | [int32](#int32) |  |  |





 

 

 

 



<a name="core_error-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## core/error.proto



<a name="core-Error"></a>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [ErrorCode](#core-ErrorCode) |  |  |
| message | [string](#string) |  |  |
| details | [google.protobuf.ListValue](#google-protobuf-ListValue) |  |  |





 

 

 

 



<a name="core_null-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## core/null.proto



<a name="core-Null"></a>

### Null






 

 

 

 



<a name="core_resource-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## core/resource.proto



<a name="core-Resource"></a>

### Resource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |





 

 

 

 



<a name="core_value-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## core/value.proto



<a name="core-Object"></a>

### Object



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fields | [Object.FieldsEntry](#core-Object-FieldsEntry) | repeated |  |






<a name="core-Object-FieldsEntry"></a>

### Object.FieldsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Value](#core-Value) |  |  |






<a name="core-Value"></a>

### Value



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| null_value | [Null](#core-Null) |  |  |
| bool_value | [bool](#bool) |  |  |
| positive_value | [uint64](#uint64) |  |  |
| negative_value | [uint64](#uint64) |  |  |
| number_value | [double](#double) |  |  |
| string_value | [string](#string) |  |  |
| bytes_value | [bytes](#bytes) |  |  |
| object_value | [Object](#core-Object) |  |  |
| values_value | [Values](#core-Values) |  |  |






<a name="core-Values"></a>

### Values



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [Value](#core-Value) | repeated |  |





 


<a name="core-ValueKind"></a>

### ValueKind


| Name | Number | Description |
| ---- | ------ | ----------- |
| VALUE_KIND_UNSPECIFIED | 0 |  |
| VALUE_KIND_NULL | 1 |  |
| VALUE_KIND_BOOLEAN | 2 |  |
| VALUE_KIND_INTEGER | 3 |  |
| VALUE_KIND_NUMBER | 4 |  |
| VALUE_KIND_STRING | 5 |  |
| VALUE_KIND_BYTES | 6 |  |
| VALUE_KIND_ARRAY | 7 |  |
| VALUE_KIND_OBJECT | 8 |  |


 

 

 



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

