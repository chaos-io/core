# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [chaos-io/core/error_code.proto](#chaos-io_core_error_code-proto)
    - [ErrorCode](#chaos_io-core-ErrorCode)
  
- [chaos-io/core/null.proto](#chaos-io_core_null-proto)
    - [Null](#chaos_io-core-Null)
  
- [chaos-io/core/value.proto](#chaos-io_core_value-proto)
    - [Object](#chaos_io-core-Object)
    - [Object.FieldsEntry](#chaos_io-core-Object-FieldsEntry)
    - [Value](#chaos_io-core-Value)
    - [Values](#chaos_io-core-Values)
  
    - [ValueKind](#chaos_io-core-ValueKind)
  
- [chaos-io/core/error.proto](#chaos-io_core_error-proto)
    - [Error](#chaos_io-core-Error)
  
- [chaos-io/core/resource.proto](#chaos-io_core_resource-proto)
    - [Resource](#chaos_io-core-Resource)
  
- [Scalar Value Types](#scalar-value-types)



<a name="chaos-io_core_error_code-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaos-io/core/error_code.proto



<a name="chaos_io-core-ErrorCode"></a>

### ErrorCode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| http_status_code | [int32](#int32) |  |  |





 

 

 

 



<a name="chaos-io_core_null-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaos-io/core/null.proto



<a name="chaos_io-core-Null"></a>

### Null






 

 

 

 



<a name="chaos-io_core_value-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaos-io/core/value.proto



<a name="chaos_io-core-Object"></a>

### Object



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fields | [Object.FieldsEntry](#chaos_io-core-Object-FieldsEntry) | repeated |  |






<a name="chaos_io-core-Object-FieldsEntry"></a>

### Object.FieldsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Value](#chaos_io-core-Value) |  |  |






<a name="chaos_io-core-Value"></a>

### Value



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| null_value | [Null](#chaos_io-core-Null) |  |  |
| bool_value | [bool](#bool) |  |  |
| positive_value | [uint64](#uint64) |  |  |
| negative_value | [uint64](#uint64) |  |  |
| number_value | [double](#double) |  |  |
| string_value | [string](#string) |  |  |
| bytes_value | [bytes](#bytes) |  |  |
| object_value | [Object](#chaos_io-core-Object) |  |  |
| values_value | [Values](#chaos_io-core-Values) |  |  |






<a name="chaos_io-core-Values"></a>

### Values



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [Value](#chaos_io-core-Value) | repeated |  |





 


<a name="chaos_io-core-ValueKind"></a>

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


 

 

 



<a name="chaos-io_core_error-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaos-io/core/error.proto



<a name="chaos_io-core-Error"></a>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [ErrorCode](#chaos_io-core-ErrorCode) |  |  |
| message | [string](#string) |  |  |
| details | [Value](#chaos_io-core-Value) | repeated |  |





 

 

 

 



<a name="chaos-io_core_resource-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaos-io/core/resource.proto



<a name="chaos_io-core-Resource"></a>

### Resource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |





 

 

 

 



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

