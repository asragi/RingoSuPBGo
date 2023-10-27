// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: gateway/schema.proto

package gateway

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetStageActionDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StageId   string `protobuf:"bytes,2,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	ExploreId string `protobuf:"bytes,4,opt,name=explore_id,json=exploreId,proto3" json:"explore_id,omitempty"`
}

func (x *GetStageActionDetailRequest) Reset() {
	*x = GetStageActionDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStageActionDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStageActionDetailRequest) ProtoMessage() {}

func (x *GetStageActionDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStageActionDetailRequest.ProtoReflect.Descriptor instead.
func (*GetStageActionDetailRequest) Descriptor() ([]byte, []int) {
	return file_gateway_schema_proto_rawDescGZIP(), []int{0}
}

func (x *GetStageActionDetailRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetStageActionDetailRequest) GetStageId() string {
	if x != nil {
		return x.StageId
	}
	return ""
}

func (x *GetStageActionDetailRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetStageActionDetailRequest) GetExploreId() string {
	if x != nil {
		return x.ExploreId
	}
	return ""
}

type GetStageActionDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId            string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StageId           string           `protobuf:"bytes,2,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	DisplayName       string           `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	ActionDisplayName string           `protobuf:"bytes,4,opt,name=action_display_name,json=actionDisplayName,proto3" json:"action_display_name,omitempty"`
	RequiredPayment   int32            `protobuf:"varint,5,opt,name=required_payment,json=requiredPayment,proto3" json:"required_payment,omitempty"`
	RequiredStamina   int32            `protobuf:"varint,6,opt,name=required_stamina,json=requiredStamina,proto3" json:"required_stamina,omitempty"`
	RequiredItems     []*RequiredItem  `protobuf:"bytes,7,rep,name=required_items,json=requiredItems,proto3" json:"required_items,omitempty"`
	EarningItems      []*EarningItem   `protobuf:"bytes,8,rep,name=earning_items,json=earningItems,proto3" json:"earning_items,omitempty"`
	RequiredSkills    []*RequiredSkill `protobuf:"bytes,9,rep,name=required_skills,json=requiredSkills,proto3" json:"required_skills,omitempty"`
}

func (x *GetStageActionDetailResponse) Reset() {
	*x = GetStageActionDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStageActionDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStageActionDetailResponse) ProtoMessage() {}

func (x *GetStageActionDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStageActionDetailResponse.ProtoReflect.Descriptor instead.
func (*GetStageActionDetailResponse) Descriptor() ([]byte, []int) {
	return file_gateway_schema_proto_rawDescGZIP(), []int{1}
}

func (x *GetStageActionDetailResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetStageActionDetailResponse) GetStageId() string {
	if x != nil {
		return x.StageId
	}
	return ""
}

func (x *GetStageActionDetailResponse) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *GetStageActionDetailResponse) GetActionDisplayName() string {
	if x != nil {
		return x.ActionDisplayName
	}
	return ""
}

func (x *GetStageActionDetailResponse) GetRequiredPayment() int32 {
	if x != nil {
		return x.RequiredPayment
	}
	return 0
}

func (x *GetStageActionDetailResponse) GetRequiredStamina() int32 {
	if x != nil {
		return x.RequiredStamina
	}
	return 0
}

func (x *GetStageActionDetailResponse) GetRequiredItems() []*RequiredItem {
	if x != nil {
		return x.RequiredItems
	}
	return nil
}

func (x *GetStageActionDetailResponse) GetEarningItems() []*EarningItem {
	if x != nil {
		return x.EarningItems
	}
	return nil
}

func (x *GetStageActionDetailResponse) GetRequiredSkills() []*RequiredSkill {
	if x != nil {
		return x.RequiredSkills
	}
	return nil
}

type RequiredItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId        string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	RequiredCount int32  `protobuf:"varint,2,opt,name=required_count,json=requiredCount,proto3" json:"required_count,omitempty"`
	Stock         int32  `protobuf:"varint,3,opt,name=stock,proto3" json:"stock,omitempty"`
	IsKnown       bool   `protobuf:"varint,4,opt,name=is_known,json=isKnown,proto3" json:"is_known,omitempty"`
}

func (x *RequiredItem) Reset() {
	*x = RequiredItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequiredItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredItem) ProtoMessage() {}

func (x *RequiredItem) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequiredItem.ProtoReflect.Descriptor instead.
func (*RequiredItem) Descriptor() ([]byte, []int) {
	return file_gateway_schema_proto_rawDescGZIP(), []int{2}
}

func (x *RequiredItem) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *RequiredItem) GetRequiredCount() int32 {
	if x != nil {
		return x.RequiredCount
	}
	return 0
}

func (x *RequiredItem) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *RequiredItem) GetIsKnown() bool {
	if x != nil {
		return x.IsKnown
	}
	return false
}

type EarningItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId  string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	IsKnown bool   `protobuf:"varint,2,opt,name=is_known,json=isKnown,proto3" json:"is_known,omitempty"`
}

func (x *EarningItem) Reset() {
	*x = EarningItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EarningItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EarningItem) ProtoMessage() {}

func (x *EarningItem) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EarningItem.ProtoReflect.Descriptor instead.
func (*EarningItem) Descriptor() ([]byte, []int) {
	return file_gateway_schema_proto_rawDescGZIP(), []int{3}
}

func (x *EarningItem) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *EarningItem) GetIsKnown() bool {
	if x != nil {
		return x.IsKnown
	}
	return false
}

type RequiredSkill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillId     string `protobuf:"bytes,1,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	SkillLv     int32  `protobuf:"varint,3,opt,name=skill_lv,json=skillLv,proto3" json:"skill_lv,omitempty"`
	RequiredLv  int32  `protobuf:"varint,4,opt,name=required_lv,json=requiredLv,proto3" json:"required_lv,omitempty"`
}

func (x *RequiredSkill) Reset() {
	*x = RequiredSkill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequiredSkill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredSkill) ProtoMessage() {}

func (x *RequiredSkill) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequiredSkill.ProtoReflect.Descriptor instead.
func (*RequiredSkill) Descriptor() ([]byte, []int) {
	return file_gateway_schema_proto_rawDescGZIP(), []int{4}
}

func (x *RequiredSkill) GetSkillId() string {
	if x != nil {
		return x.SkillId
	}
	return ""
}

func (x *RequiredSkill) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *RequiredSkill) GetSkillLv() int32 {
	if x != nil {
		return x.SkillLv
	}
	return 0
}

func (x *RequiredSkill) GetRequiredLv() int32 {
	if x != nil {
		return x.RequiredLv
	}
	return 0
}

type PostActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ExploreId string `protobuf:"bytes,3,opt,name=explore_id,json=exploreId,proto3" json:"explore_id,omitempty"`
	ExecCount int32  `protobuf:"varint,4,opt,name=exec_count,json=execCount,proto3" json:"exec_count,omitempty"`
}

func (x *PostActionRequest) Reset() {
	*x = PostActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostActionRequest) ProtoMessage() {}

func (x *PostActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostActionRequest.ProtoReflect.Descriptor instead.
func (*PostActionRequest) Descriptor() ([]byte, []int) {
	return file_gateway_schema_proto_rawDescGZIP(), []int{5}
}

func (x *PostActionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PostActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PostActionRequest) GetExploreId() string {
	if x != nil {
		return x.ExploreId
	}
	return ""
}

func (x *PostActionRequest) GetExecCount() int32 {
	if x != nil {
		return x.ExecCount
	}
	return 0
}

type EarnedItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId string `protobuf:"bytes,10,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Count  int32  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *EarnedItems) Reset() {
	*x = EarnedItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_schema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EarnedItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EarnedItems) ProtoMessage() {}

func (x *EarnedItems) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_schema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EarnedItems.ProtoReflect.Descriptor instead.
func (*EarnedItems) Descriptor() ([]byte, []int) {
	return file_gateway_schema_proto_rawDescGZIP(), []int{6}
}

func (x *EarnedItems) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *EarnedItems) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ConsumedItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId string `protobuf:"bytes,10,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Count  int32  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ConsumedItems) Reset() {
	*x = ConsumedItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_schema_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumedItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumedItems) ProtoMessage() {}

func (x *ConsumedItems) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_schema_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumedItems.ProtoReflect.Descriptor instead.
func (*ConsumedItems) Descriptor() ([]byte, []int) {
	return file_gateway_schema_proto_rawDescGZIP(), []int{7}
}

func (x *ConsumedItems) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *ConsumedItems) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type SkillGrowthResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillId     string `protobuf:"bytes,10,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	BeforeExp   int32  `protobuf:"varint,1,opt,name=before_exp,json=beforeExp,proto3" json:"before_exp,omitempty"`
	BeforeLv    int32  `protobuf:"varint,2,opt,name=before_lv,json=beforeLv,proto3" json:"before_lv,omitempty"`
	AfterExp    int32  `protobuf:"varint,3,opt,name=after_exp,json=afterExp,proto3" json:"after_exp,omitempty"`
	AfterLv     int32  `protobuf:"varint,4,opt,name=after_lv,json=afterLv,proto3" json:"after_lv,omitempty"`
	DisplayName string `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *SkillGrowthResult) Reset() {
	*x = SkillGrowthResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_schema_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkillGrowthResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillGrowthResult) ProtoMessage() {}

func (x *SkillGrowthResult) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_schema_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillGrowthResult.ProtoReflect.Descriptor instead.
func (*SkillGrowthResult) Descriptor() ([]byte, []int) {
	return file_gateway_schema_proto_rawDescGZIP(), []int{8}
}

func (x *SkillGrowthResult) GetSkillId() string {
	if x != nil {
		return x.SkillId
	}
	return ""
}

func (x *SkillGrowthResult) GetBeforeExp() int32 {
	if x != nil {
		return x.BeforeExp
	}
	return 0
}

func (x *SkillGrowthResult) GetBeforeLv() int32 {
	if x != nil {
		return x.BeforeLv
	}
	return 0
}

func (x *SkillGrowthResult) GetAfterExp() int32 {
	if x != nil {
		return x.AfterExp
	}
	return 0
}

func (x *SkillGrowthResult) GetAfterLv() int32 {
	if x != nil {
		return x.AfterLv
	}
	return 0
}

func (x *SkillGrowthResult) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type PostActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error             *Error               `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	EarnedItems       []*EarnedItems       `protobuf:"bytes,2,rep,name=earned_items,json=earnedItems,proto3" json:"earned_items,omitempty"`
	ConsumedItems     []*ConsumedItems     `protobuf:"bytes,3,rep,name=consumed_items,json=consumedItems,proto3" json:"consumed_items,omitempty"`
	SkillGrowthResult []*SkillGrowthResult `protobuf:"bytes,4,rep,name=skill_growth_result,json=skillGrowthResult,proto3" json:"skill_growth_result,omitempty"`
}

func (x *PostActionResponse) Reset() {
	*x = PostActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_schema_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostActionResponse) ProtoMessage() {}

func (x *PostActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_schema_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostActionResponse.ProtoReflect.Descriptor instead.
func (*PostActionResponse) Descriptor() ([]byte, []int) {
	return file_gateway_schema_proto_rawDescGZIP(), []int{9}
}

func (x *PostActionResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PostActionResponse) GetEarnedItems() []*EarnedItems {
	if x != nil {
		return x.EarnedItems
	}
	return nil
}

func (x *PostActionResponse) GetConsumedItems() []*ConsumedItems {
	if x != nil {
		return x.ConsumedItems
	}
	return nil
}

func (x *PostActionResponse) GetSkillGrowthResult() []*SkillGrowthResult {
	if x != nil {
		return x.SkillGrowthResult
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorOccured   bool   `protobuf:"varint,1,opt,name=error_occured,json=errorOccured,proto3" json:"error_occured,omitempty"`
	DisplayMessage string `protobuf:"bytes,2,opt,name=display_message,json=displayMessage,proto3" json:"display_message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_schema_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_schema_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_gateway_schema_proto_rawDescGZIP(), []int{10}
}

func (x *Error) GetErrorOccured() bool {
	if x != nil {
		return x.ErrorOccured
	}
	return false
}

func (x *Error) GetDisplayMessage() string {
	if x != nil {
		return x.DisplayMessage
	}
	return ""
}

var File_gateway_schema_proto protoreflect.FileDescriptor

var file_gateway_schema_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x69, 0x6e, 0x67, 0x6f, 0x73, 0x75, 0x22,
	0x86, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x6c, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22, 0xb5, 0x03, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x13, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53,
	0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x12, 0x3c, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x72, 0x69, 0x6e, 0x67, 0x6f, 0x73, 0x75, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x39, 0x0a, 0x0d, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x69,
	0x6e, 0x67, 0x6f, 0x73, 0x75, 0x2e, 0x45, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x0c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x3f, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x6f,
	0x73, 0x75, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x22, 0x7f, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4b, 0x6e, 0x6f, 0x77,
	0x6e, 0x22, 0x41, 0x0a, 0x0b, 0x45, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4b,
	0x6e, 0x6f, 0x77, 0x6e, 0x22, 0x89, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x76,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x76, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6c, 0x76, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x4c, 0x76,
	0x22, 0x80, 0x01, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x78, 0x65, 0x63, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x0b, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x3e, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0xc5, 0x01, 0x0a, 0x11, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x77, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x65, 0x78, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x45, 0x78,
	0x70, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x6c, 0x76, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x4c, 0x76, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x61, 0x66, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x4c, 0x76, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xfe, 0x01, 0x0a, 0x12, 0x50, 0x6f,
	0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x24, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x6f, 0x73, 0x75, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x0c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72,
	0x69, 0x6e, 0x67, 0x6f, 0x73, 0x75, 0x2e, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x0b, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x3d, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x6f, 0x73,
	0x75, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x4a,
	0x0a, 0x13, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x69,
	0x6e, 0x67, 0x6f, 0x73, 0x75, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x77, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x11, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72,
	0x6f, 0x77, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x55, 0x0a, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6f, 0x63, 0x63,
	0x75, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_schema_proto_rawDescOnce sync.Once
	file_gateway_schema_proto_rawDescData = file_gateway_schema_proto_rawDesc
)

func file_gateway_schema_proto_rawDescGZIP() []byte {
	file_gateway_schema_proto_rawDescOnce.Do(func() {
		file_gateway_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_schema_proto_rawDescData)
	})
	return file_gateway_schema_proto_rawDescData
}

var file_gateway_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_gateway_schema_proto_goTypes = []interface{}{
	(*GetStageActionDetailRequest)(nil),  // 0: ringosu.GetStageActionDetailRequest
	(*GetStageActionDetailResponse)(nil), // 1: ringosu.GetStageActionDetailResponse
	(*RequiredItem)(nil),                 // 2: ringosu.RequiredItem
	(*EarningItem)(nil),                  // 3: ringosu.EarningItem
	(*RequiredSkill)(nil),                // 4: ringosu.RequiredSkill
	(*PostActionRequest)(nil),            // 5: ringosu.PostActionRequest
	(*EarnedItems)(nil),                  // 6: ringosu.EarnedItems
	(*ConsumedItems)(nil),                // 7: ringosu.ConsumedItems
	(*SkillGrowthResult)(nil),            // 8: ringosu.SkillGrowthResult
	(*PostActionResponse)(nil),           // 9: ringosu.PostActionResponse
	(*Error)(nil),                        // 10: ringosu.Error
}
var file_gateway_schema_proto_depIdxs = []int32{
	2,  // 0: ringosu.GetStageActionDetailResponse.required_items:type_name -> ringosu.RequiredItem
	3,  // 1: ringosu.GetStageActionDetailResponse.earning_items:type_name -> ringosu.EarningItem
	4,  // 2: ringosu.GetStageActionDetailResponse.required_skills:type_name -> ringosu.RequiredSkill
	10, // 3: ringosu.PostActionResponse.error:type_name -> ringosu.Error
	6,  // 4: ringosu.PostActionResponse.earned_items:type_name -> ringosu.EarnedItems
	7,  // 5: ringosu.PostActionResponse.consumed_items:type_name -> ringosu.ConsumedItems
	8,  // 6: ringosu.PostActionResponse.skill_growth_result:type_name -> ringosu.SkillGrowthResult
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_gateway_schema_proto_init() }
func file_gateway_schema_proto_init() {
	if File_gateway_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStageActionDetailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStageActionDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequiredItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EarningItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequiredSkill); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_schema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EarnedItems); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_schema_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumedItems); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_schema_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkillGrowthResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_schema_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostActionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_schema_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_schema_proto_goTypes,
		DependencyIndexes: file_gateway_schema_proto_depIdxs,
		MessageInfos:      file_gateway_schema_proto_msgTypes,
	}.Build()
	File_gateway_schema_proto = out.File
	file_gateway_schema_proto_rawDesc = nil
	file_gateway_schema_proto_goTypes = nil
	file_gateway_schema_proto_depIdxs = nil
}
