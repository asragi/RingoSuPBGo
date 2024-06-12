# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [gateway/schema.proto](#gateway_schema-proto)
    - [AdminLoginRequest](#ringosu-AdminLoginRequest)
    - [AdminLoginResponse](#ringosu-AdminLoginResponse)
    - [ChangePeriodRequest](#ringosu-ChangePeriodRequest)
    - [ChangePeriodResponse](#ringosu-ChangePeriodResponse)
    - [ChangeTimeRequest](#ringosu-ChangeTimeRequest)
    - [ChangeTimeResponse](#ringosu-ChangeTimeResponse)
    - [ConsumedItems](#ringosu-ConsumedItems)
    - [EarnedItems](#ringosu-EarnedItems)
    - [EarningItem](#ringosu-EarningItem)
    - [GetDailyRankingRequest](#ringosu-GetDailyRankingRequest)
    - [GetDailyRankingResponse](#ringosu-GetDailyRankingResponse)
    - [GetItemActionDetailRequest](#ringosu-GetItemActionDetailRequest)
    - [GetItemActionDetailResponse](#ringosu-GetItemActionDetailResponse)
    - [GetItemDetailRequest](#ringosu-GetItemDetailRequest)
    - [GetItemDetailResponse](#ringosu-GetItemDetailResponse)
    - [GetItemListRequest](#ringosu-GetItemListRequest)
    - [GetItemListResponse](#ringosu-GetItemListResponse)
    - [GetItemListResponseRow](#ringosu-GetItemListResponseRow)
    - [GetMyShelfRequest](#ringosu-GetMyShelfRequest)
    - [GetMyShelfResponse](#ringosu-GetMyShelfResponse)
    - [GetResourceRequest](#ringosu-GetResourceRequest)
    - [GetResourceResponse](#ringosu-GetResourceResponse)
    - [GetShopsRequest](#ringosu-GetShopsRequest)
    - [GetShopsResponse](#ringosu-GetShopsResponse)
    - [GetStageActionDetailRequest](#ringosu-GetStageActionDetailRequest)
    - [GetStageActionDetailResponse](#ringosu-GetStageActionDetailResponse)
    - [GetStageListRequest](#ringosu-GetStageListRequest)
    - [GetStageListResponse](#ringosu-GetStageListResponse)
    - [InvokeAutoUpdateRequest](#ringosu-InvokeAutoUpdateRequest)
    - [InvokeAutoUpdateResponse](#ringosu-InvokeAutoUpdateResponse)
    - [LoginRequest](#ringosu-LoginRequest)
    - [LoginResponse](#ringosu-LoginResponse)
    - [PostActionRequest](#ringosu-PostActionRequest)
    - [PostActionResponse](#ringosu-PostActionResponse)
    - [RankingRow](#ringosu-RankingRow)
    - [RequiredItem](#ringosu-RequiredItem)
    - [RequiredSkill](#ringosu-RequiredSkill)
    - [Reservation](#ringosu-Reservation)
    - [Shelf](#ringosu-Shelf)
    - [Shop](#ringosu-Shop)
    - [SignUpRequest](#ringosu-SignUpRequest)
    - [SignUpResponse](#ringosu-SignUpResponse)
    - [SkillGrowthResult](#ringosu-SkillGrowthResult)
    - [StageInformation](#ringosu-StageInformation)
    - [UpdateShelfContentRequest](#ringosu-UpdateShelfContentRequest)
    - [UpdateShelfContentResponse](#ringosu-UpdateShelfContentResponse)
    - [UpdateShelfSizeRequest](#ringosu-UpdateShelfSizeRequest)
    - [UpdateShelfSizeResponse](#ringosu-UpdateShelfSizeResponse)
    - [UpdateShopNameRequest](#ringosu-UpdateShopNameRequest)
    - [UpdateShopNameResponse](#ringosu-UpdateShopNameResponse)
    - [UpdateUserNameRequest](#ringosu-UpdateUserNameRequest)
    - [UpdateUserNameResponse](#ringosu-UpdateUserNameResponse)
    - [UserExplore](#ringosu-UserExplore)
  
    - [ChangePeriod](#ringosu-ChangePeriod)
    - [DebugTime](#ringosu-DebugTime)
    - [InvokeUpdate](#ringosu-InvokeUpdate)
    - [Ringo](#ringosu-Ringo)
  
- [Scalar Value Types](#scalar-value-types)



<a name="gateway_schema-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gateway/schema.proto



<a name="ringosu-AdminLoginRequest"></a>

### AdminLoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| row_password | [string](#string) |  |  |






<a name="ringosu-AdminLoginResponse"></a>

### AdminLoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="ringosu-ChangePeriodRequest"></a>

### ChangePeriodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="ringosu-ChangePeriodResponse"></a>

### ChangePeriodResponse







<a name="ringosu-ChangeTimeRequest"></a>

### ChangeTimeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| recover_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="ringosu-ChangeTimeResponse"></a>

### ChangeTimeResponse







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






<a name="ringosu-GetDailyRankingRequest"></a>

### GetDailyRankingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






<a name="ringosu-GetDailyRankingResponse"></a>

### GetDailyRankingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ranking | [RankingRow](#ringosu-RankingRow) | repeated |  |






<a name="ringosu-GetItemActionDetailRequest"></a>

### GetItemActionDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_id | [string](#string) |  |  |
| explore_id | [string](#string) |  |  |
| access_token | [string](#string) |  |  |






<a name="ringosu-GetItemActionDetailResponse"></a>

### GetItemActionDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| item_id | [string](#string) |  |  |
| display_name | [string](#string) |  |  |
| action_display_name | [string](#string) |  |  |
| required_payment | [int32](#int32) |  |  |
| required_stamina | [int32](#int32) |  |  |
| required_items | [RequiredItem](#ringosu-RequiredItem) | repeated |  |
| earning_items | [EarningItem](#ringosu-EarningItem) | repeated |  |
| required_skills | [RequiredSkill](#ringosu-RequiredSkill) | repeated |  |






<a name="ringosu-GetItemDetailRequest"></a>

### GetItemDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
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






<a name="ringosu-GetMyShelfRequest"></a>

### GetMyShelfRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="ringosu-GetMyShelfResponse"></a>

### GetMyShelfResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shelves | [Shelf](#ringosu-Shelf) | repeated |  |






<a name="ringosu-GetResourceRequest"></a>

### GetResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="ringosu-GetResourceResponse"></a>

### GetResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| max_stamina | [int32](#int32) |  |  |
| fund | [int32](#int32) |  |  |
| recover_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="ringosu-GetShopsRequest"></a>

### GetShopsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| page | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |






<a name="ringosu-GetShopsResponse"></a>

### GetShopsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shop | [Shop](#ringosu-Shop) | repeated |  |






<a name="ringosu-GetStageActionDetailRequest"></a>

### GetStageActionDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
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
| token | [string](#string) |  |  |






<a name="ringosu-GetStageListResponse"></a>

### GetStageListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stage_information | [StageInformation](#ringosu-StageInformation) | repeated |  |






<a name="ringosu-InvokeAutoUpdateRequest"></a>

### InvokeAutoUpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="ringosu-InvokeAutoUpdateResponse"></a>

### InvokeAutoUpdateResponse







<a name="ringosu-LoginRequest"></a>

### LoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| row_password | [string](#string) |  |  |






<a name="ringosu-LoginResponse"></a>

### LoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |






<a name="ringosu-PostActionRequest"></a>

### PostActionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| explore_id | [string](#string) |  |  |
| exec_count | [int32](#int32) |  |  |






<a name="ringosu-PostActionResponse"></a>

### PostActionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| earned_items | [EarnedItems](#ringosu-EarnedItems) | repeated |  |
| consumed_items | [ConsumedItems](#ringosu-ConsumedItems) | repeated |  |
| skill_growth_result | [SkillGrowthResult](#ringosu-SkillGrowthResult) | repeated |  |






<a name="ringosu-RankingRow"></a>

### RankingRow



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| user_name | [string](#string) |  |  |
| rank | [int32](#int32) |  |  |
| total_score | [int32](#int32) |  |  |
| shelves | [Shelf](#ringosu-Shelf) | repeated |  |






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






<a name="ringosu-Reservation"></a>

### Reservation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reservation_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| index | [int32](#int32) |  |  |
| purchase_num | [int32](#int32) |  |  |
| scheduled_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="ringosu-Shelf"></a>

### Shelf



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [int32](#int32) |  |  |
| set_price | [int32](#int32) |  |  |
| item_id | [string](#string) |  |  |
| display_name | [string](#string) |  |  |
| stock | [int32](#int32) |  |  |
| max_stock | [int32](#int32) |  |  |
| user_id | [string](#string) |  |  |
| shelf_id | [string](#string) |  |  |






<a name="ringosu-Shop"></a>

### Shop



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| user_name | [string](#string) |  |  |
| rank | [int32](#int32) |  |  |
| shelves | [Shelf](#ringosu-Shelf) | repeated |  |






<a name="ringosu-SignUpRequest"></a>

### SignUpRequest







<a name="ringosu-SignUpResponse"></a>

### SignUpResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| row_password | [string](#string) |  |  |






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






<a name="ringosu-UpdateShelfContentRequest"></a>

### UpdateShelfContentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| index | [int32](#int32) |  |  |
| set_price | [int32](#int32) |  |  |
| item_id | [string](#string) |  |  |






<a name="ringosu-UpdateShelfContentResponse"></a>

### UpdateShelfContentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [int32](#int32) |  |  |
| set_price | [int32](#int32) |  |  |
| item_id | [string](#string) |  |  |
| reservations | [Reservation](#ringosu-Reservation) | repeated |  |






<a name="ringosu-UpdateShelfSizeRequest"></a>

### UpdateShelfSizeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| size | [int32](#int32) |  |  |






<a name="ringosu-UpdateShelfSizeResponse"></a>

### UpdateShelfSizeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| size | [int32](#int32) |  |  |






<a name="ringosu-UpdateShopNameRequest"></a>

### UpdateShopNameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| shop_name | [string](#string) |  |  |






<a name="ringosu-UpdateShopNameResponse"></a>

### UpdateShopNameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shop_name | [string](#string) |  |  |






<a name="ringosu-UpdateUserNameRequest"></a>

### UpdateUserNameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| user_name | [string](#string) |  |  |






<a name="ringosu-UpdateUserNameResponse"></a>

### UpdateUserNameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_name | [string](#string) |  |  |






<a name="ringosu-UserExplore"></a>

### UserExplore



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| explore_id | [string](#string) |  |  |
| display_name | [string](#string) |  |  |
| is_known | [bool](#bool) |  |  |
| is_possible | [bool](#bool) |  |  |





 

 

 


<a name="ringosu-ChangePeriod"></a>

### ChangePeriod


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Login | [AdminLoginRequest](#ringosu-AdminLoginRequest) | [AdminLoginResponse](#ringosu-AdminLoginResponse) |  |
| Do | [ChangePeriodRequest](#ringosu-ChangePeriodRequest) | [ChangePeriodResponse](#ringosu-ChangePeriodResponse) |  |


<a name="ringosu-DebugTime"></a>

### DebugTime


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Login | [AdminLoginRequest](#ringosu-AdminLoginRequest) | [AdminLoginResponse](#ringosu-AdminLoginResponse) |  |
| ChangeTime | [ChangePeriodRequest](#ringosu-ChangePeriodRequest) | [ChangePeriodResponse](#ringosu-ChangePeriodResponse) |  |


<a name="ringosu-InvokeUpdate"></a>

### InvokeUpdate


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Login | [AdminLoginRequest](#ringosu-AdminLoginRequest) | [AdminLoginResponse](#ringosu-AdminLoginResponse) |  |
| Do | [InvokeAutoUpdateRequest](#ringosu-InvokeAutoUpdateRequest) | [InvokeAutoUpdateResponse](#ringosu-InvokeAutoUpdateResponse) |  |


<a name="ringosu-Ringo"></a>

### Ringo


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SignUp | [SignUpRequest](#ringosu-SignUpRequest) | [SignUpResponse](#ringosu-SignUpResponse) |  |
| Login | [LoginRequest](#ringosu-LoginRequest) | [LoginResponse](#ringosu-LoginResponse) |  |
| GetResource | [GetResourceRequest](#ringosu-GetResourceRequest) | [GetResourceResponse](#ringosu-GetResourceResponse) |  |
| UpdateUserName | [UpdateUserNameRequest](#ringosu-UpdateUserNameRequest) | [UpdateUserNameResponse](#ringosu-UpdateUserNameResponse) |  |
| UpdateShopName | [UpdateShopNameRequest](#ringosu-UpdateShopNameRequest) | [UpdateShopNameResponse](#ringosu-UpdateShopNameResponse) |  |
| GetMyShelf | [GetMyShelfRequest](#ringosu-GetMyShelfRequest) | [GetMyShelfResponse](#ringosu-GetMyShelfResponse) |  |
| GetStageList | [GetStageListRequest](#ringosu-GetStageListRequest) | [GetStageListResponse](#ringosu-GetStageListResponse) |  |
| GetStageActionDetail | [GetStageActionDetailRequest](#ringosu-GetStageActionDetailRequest) | [GetStageActionDetailResponse](#ringosu-GetStageActionDetailResponse) |  |
| PostAction | [PostActionRequest](#ringosu-PostActionRequest) | [PostActionResponse](#ringosu-PostActionResponse) |  |
| GetItemList | [GetItemListRequest](#ringosu-GetItemListRequest) | [GetItemListResponse](#ringosu-GetItemListResponse) |  |
| GetItemDetail | [GetItemDetailRequest](#ringosu-GetItemDetailRequest) | [GetItemDetailResponse](#ringosu-GetItemDetailResponse) |  |
| GetItemActionDetail | [GetItemActionDetailRequest](#ringosu-GetItemActionDetailRequest) | [GetItemActionDetailResponse](#ringosu-GetItemActionDetailResponse) |  |
| UpdateShelfContent | [UpdateShelfContentRequest](#ringosu-UpdateShelfContentRequest) | [UpdateShelfContentResponse](#ringosu-UpdateShelfContentResponse) |  |
| UpdateShelfSize | [UpdateShelfSizeRequest](#ringosu-UpdateShelfSizeRequest) | [UpdateShelfSizeResponse](#ringosu-UpdateShelfSizeResponse) |  |
| GetShops | [GetShopsRequest](#ringosu-GetShopsRequest) | [GetShopsResponse](#ringosu-GetShopsResponse) |  |
| GetDailyRanking | [GetDailyRankingRequest](#ringosu-GetDailyRankingRequest) | [GetDailyRankingResponse](#ringosu-GetDailyRankingResponse) |  |

 



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

