# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [gateway/schema.proto](#gateway_schema-proto)
    - [ConsumedItems](#ringosu-ConsumedItems)
    - [EarnedItems](#ringosu-EarnedItems)
    - [EarningItem](#ringosu-EarningItem)
    - [Error](#ringosu-Error)
    - [GetItemDetailRequest](#ringosu-GetItemDetailRequest)
    - [GetItemDetailResponse](#ringosu-GetItemDetailResponse)
    - [GetItemListRequest](#ringosu-GetItemListRequest)
    - [GetItemListResponse](#ringosu-GetItemListResponse)
    - [GetItemListResponseRow](#ringosu-GetItemListResponseRow)
    - [GetStageActionDetailRequest](#ringosu-GetStageActionDetailRequest)
    - [GetStageActionDetailResponse](#ringosu-GetStageActionDetailResponse)
    - [GetStageListRequest](#ringosu-GetStageListRequest)
    - [GetStageListResponse](#ringosu-GetStageListResponse)
    - [PostActionRequest](#ringosu-PostActionRequest)
    - [PostActionResponse](#ringosu-PostActionResponse)
    - [RequiredItem](#ringosu-RequiredItem)
    - [RequiredSkill](#ringosu-RequiredSkill)
    - [SkillGrowthResult](#ringosu-SkillGrowthResult)
    - [StageInformation](#ringosu-StageInformation)
    - [UserExplore](#ringosu-UserExplore)
  
- [Scalar Value Types](#scalar-value-types)



<a name="gateway_schema-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gateway/schema.proto



<a name="ringosu-ConsumedItems"></a>

### ConsumedItems



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_id | [string](#string) |  |  |
| count | [int32](#int32) |  |  |






<a name="ringosu-EarnedItems"></a>

### EarnedItems



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_id | [string](#string) |  |  |
| count | [int32](#int32) |  |  |






<a name="ringosu-EarningItem"></a>

### EarningItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_id | [string](#string) |  |  |
| is_known | [bool](#bool) |  |  |






<a name="ringosu-Error"></a>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_occured | [bool](#bool) |  |  |
| display_message | [string](#string) |  |  |






<a name="ringosu-GetItemDetailRequest"></a>

### GetItemDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| item_id | [string](#string) |  |  |
| token | [string](#string) |  |  |






<a name="ringosu-GetItemDetailResponse"></a>

### GetItemDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| item_id | [string](#string) |  |  |
| price | [int32](#int32) |  |  |
| display_name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| max_stock | [int32](#int32) |  |  |
| stock | [int32](#int32) |  |  |
| user_explore | [UserExplore](#ringosu-UserExplore) | repeated |  |






<a name="ringosu-GetItemListRequest"></a>

### GetItemListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| token | [string](#string) |  |  |






<a name="ringosu-GetItemListResponse"></a>

### GetItemListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_list | [GetItemListResponseRow](#ringosu-GetItemListResponseRow) | repeated |  |






<a name="ringosu-GetItemListResponseRow"></a>

### GetItemListResponseRow



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_id | [string](#string) |  |  |
| display_name | [string](#string) |  |  |
| price | [int32](#int32) |  |  |
| stock | [int32](#int32) |  |  |
| max_stock | [int32](#int32) |  |  |






<a name="ringosu-GetStageActionDetailRequest"></a>

### GetStageActionDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| stage_id | [string](#string) |  |  |
| token | [string](#string) |  |  |
| explore_id | [string](#string) |  |  |






<a name="ringosu-GetStageActionDetailResponse"></a>

### GetStageActionDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| stage_id | [string](#string) |  |  |
| display_name | [string](#string) |  |  |
| action_display_name | [string](#string) |  |  |
| required_payment | [int32](#int32) |  |  |
| required_stamina | [int32](#int32) |  |  |
| required_items | [RequiredItem](#ringosu-RequiredItem) | repeated |  |
| earning_items | [EarningItem](#ringosu-EarningItem) | repeated |  |
| required_skills | [RequiredSkill](#ringosu-RequiredSkill) | repeated |  |






<a name="ringosu-GetStageListRequest"></a>

### GetStageListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| token | [string](#string) |  |  |






<a name="ringosu-GetStageListResponse"></a>

### GetStageListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stage_information | [StageInformation](#ringosu-StageInformation) | repeated |  |






<a name="ringosu-PostActionRequest"></a>

### PostActionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| token | [string](#string) |  |  |
| explore_id | [string](#string) |  |  |
| exec_count | [int32](#int32) |  |  |






<a name="ringosu-PostActionResponse"></a>

### PostActionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [Error](#ringosu-Error) |  |  |
| earned_items | [EarnedItems](#ringosu-EarnedItems) | repeated |  |
| consumed_items | [ConsumedItems](#ringosu-ConsumedItems) | repeated |  |
| skill_growth_result | [SkillGrowthResult](#ringosu-SkillGrowthResult) | repeated |  |






<a name="ringosu-RequiredItem"></a>

### RequiredItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_id | [string](#string) |  |  |
| required_count | [int32](#int32) |  |  |
| stock | [int32](#int32) |  |  |
| is_known | [bool](#bool) |  |  |






<a name="ringosu-RequiredSkill"></a>

### RequiredSkill



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| skill_id | [string](#string) |  |  |
| display_name | [string](#string) |  |  |
| skill_lv | [int32](#int32) |  |  |
| required_lv | [int32](#int32) |  |  |






<a name="ringosu-SkillGrowthResult"></a>

### SkillGrowthResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| skill_id | [string](#string) |  |  |
| before_exp | [int32](#int32) |  |  |
| before_lv | [int32](#int32) |  |  |
| after_exp | [int32](#int32) |  |  |
| after_lv | [int32](#int32) |  |  |
| display_name | [string](#string) |  |  |






<a name="ringosu-StageInformation"></a>

### StageInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stage_id | [string](#string) |  |  |
| display_name | [string](#string) |  |  |
| is_known | [bool](#bool) |  |  |
| description | [string](#string) |  |  |
| user_explore | [UserExplore](#ringosu-UserExplore) | repeated |  |






<a name="ringosu-UserExplore"></a>

### UserExplore



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| explore_id | [string](#string) |  |  |
| display_name | [string](#string) |  |  |
| is_known | [bool](#bool) |  |  |
| is_possible | [bool](#bool) |  |  |





 

 

 

 



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

