# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [fino2/options.proto](#fino2_options-proto)
    - [DBOptions](#fino2-DBOptions)
    - [MessagingOptions](#fino2-MessagingOptions)
    - [ModelOptions](#fino2-ModelOptions)
    - [ValidateOptions](#fino2-ValidateOptions)
  
    - [File-level Extensions](#fino2_options-proto-extensions)
    - [File-level Extensions](#fino2_options-proto-extensions)
    - [File-level Extensions](#fino2_options-proto-extensions)
    - [File-level Extensions](#fino2_options-proto-extensions)
  
- [Scalar Value Types](#scalar-value-types)



<a name="fino2_options-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fino2/options.proto



<a name="fino2-DBOptions"></a>

### DBOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| null | [bool](#bool) |  |  |
| len | [int64](#int64) |  |  |
| default | [string](#string) |  |  |
| comment | [string](#string) |  |  |
| index | [string](#string) |  |  |
| unique_index | [string](#string) |  |  |






<a name="fino2-MessagingOptions"></a>

### MessagingOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [bool](#bool) |  |  |






<a name="fino2-ModelOptions"></a>

### ModelOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| generate | [bool](#bool) |  |  |
| services | [string](#string) | repeated |  |






<a name="fino2-ValidateOptions"></a>

### ValidateOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| required | [bool](#bool) |  |  |
| omitempty | [bool](#bool) |  |  |
| dive | [bool](#bool) |  |  |
| len | [int64](#int64) |  |  |
| min | [int64](#int64) |  |  |
| max | [int64](#int64) |  |  |
| oneof | [string](#string) |  |  |
| email | [bool](#bool) |  |  |
| url | [bool](#bool) |  |  |
| uuid | [bool](#bool) |  |  |





 

 


<a name="fino2_options-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| db | DBOptions | .google.protobuf.FieldOptions | 51002 |  |
| validate | ValidateOptions | .google.protobuf.FieldOptions | 51003 |  |
| model | ModelOptions | .google.protobuf.MessageOptions | 51001 |  |
| messaging | MessagingOptions | .google.protobuf.MethodOptions | 51004 |  |

 

 



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

