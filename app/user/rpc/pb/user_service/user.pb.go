// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: user.proto

package user_service

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

type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                         // 用户id
	NickName       string `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`              // 昵称
	AvatarUrl      string `protobuf:"bytes,3,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"`            // 头像
	Slogan         string `protobuf:"bytes,4,opt,name=slogan,proto3" json:"slogan,omitempty"`                  // 个性签名
	Gender         int32  `protobuf:"varint,5,opt,name=gender,proto3" json:"gender,omitempty"`                 // 性别
	Age            int32  `protobuf:"varint,6,opt,name=age,proto3" json:"age,omitempty"`                       // 年龄
	Mobile         string `protobuf:"bytes,7,opt,name=mobile,proto3" json:"mobile,omitempty"`                  // 手机号
	FollowingCount int64  `protobuf:"varint,8,opt,name=followingCount,proto3" json:"followingCount,omitempty"` // 关注数
	FollowerCount  int64  `protobuf:"varint,9,opt,name=followerCount,proto3" json:"followerCount,omitempty"`   // 粉丝数
	LikeCount      int64  `protobuf:"varint,10,opt,name=likeCount,proto3" json:"likeCount,omitempty"`          // 获得的点赞数
	CollectCount   int64  `protobuf:"varint,11,opt,name=collectCount,proto3" json:"collectCount,omitempty"`    // 收藏的视频数
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfoResponse) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserInfoResponse) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *UserInfoResponse) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *UserInfoResponse) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserInfoResponse) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserInfoResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserInfoResponse) GetFollowingCount() int64 {
	if x != nil {
		return x.FollowingCount
	}
	return 0
}

func (x *UserInfoResponse) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *UserInfoResponse) GetLikeCount() int64 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *UserInfoResponse) GetCollectCount() int64 {
	if x != nil {
		return x.CollectCount
	}
	return 0
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Mobile   string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Avatar   string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegisterRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ListByIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []int64 `protobuf:"varint,1,rep,packed,name=userIds,proto3" json:"userIds,omitempty"`
}

func (x *ListByIdsRequest) Reset() {
	*x = ListByIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListByIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByIdsRequest) ProtoMessage() {}

func (x *ListByIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListByIdsRequest.ProtoReflect.Descriptor instead.
func (*ListByIdsRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *ListByIdsRequest) GetUserIds() []int64 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type ListByIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32               `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*UserInfoResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListByIdsResponse) Reset() {
	*x = ListByIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListByIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByIdsResponse) ProtoMessage() {}

func (x *ListByIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListByIdsResponse.ProtoReflect.Descriptor instead.
func (*ListByIdsResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *ListByIdsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListByIdsResponse) GetData() []*UserInfoResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type FindByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *FindByIdRequest) Reset() {
	*x = FindByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIdRequest) ProtoMessage() {}

func (x *FindByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIdRequest.ProtoReflect.Descriptor instead.
func (*FindByIdRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *FindByIdRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FindByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Mobile   string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Avatar   string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *FindByIdResponse) Reset() {
	*x = FindByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIdResponse) ProtoMessage() {}

func (x *FindByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIdResponse.ProtoReflect.Descriptor instead.
func (*FindByIdResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *FindByIdResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FindByIdResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *FindByIdResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *FindByIdResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type FindByMobileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *FindByMobileRequest) Reset() {
	*x = FindByMobileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByMobileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByMobileRequest) ProtoMessage() {}

func (x *FindByMobileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByMobileRequest.ProtoReflect.Descriptor instead.
func (*FindByMobileRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *FindByMobileRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type FindByMobileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Mobile   string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Avatar   string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *FindByMobileResponse) Reset() {
	*x = FindByMobileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByMobileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByMobileResponse) ProtoMessage() {}

func (x *FindByMobileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByMobileResponse.ProtoReflect.Descriptor instead.
func (*FindByMobileResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *FindByMobileResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FindByMobileResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *FindByMobileResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *FindByMobileResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type SendSmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Mobile string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *SendSmsRequest) Reset() {
	*x = SendSmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsRequest) ProtoMessage() {}

func (x *SendSmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsRequest.ProtoReflect.Descriptor instead.
func (*SendSmsRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *SendSmsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SendSmsRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type SendSmsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendSmsResponse) Reset() {
	*x = SendSmsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsResponse) ProtoMessage() {}

func (x *SendSmsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsResponse.ProtoReflect.Descriptor instead.
func (*SendSmsResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xc6, 0x02, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x79,
	0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x73, 0x22, 0x58, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2d,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x29, 0x0a,
	0x0f, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x76, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x22, 0x2d, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22,
	0x7a, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x40, 0x0a, 0x0e, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x11, 0x0a,
	0x0f, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xd7, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4b, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07,
	0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_user_proto_goTypes = []interface{}{
	(*UserInfoResponse)(nil),     // 0: service.UserInfoResponse
	(*RegisterRequest)(nil),      // 1: service.RegisterRequest
	(*RegisterResponse)(nil),     // 2: service.RegisterResponse
	(*ListByIdsRequest)(nil),     // 3: service.ListByIdsRequest
	(*ListByIdsResponse)(nil),    // 4: service.ListByIdsResponse
	(*FindByIdRequest)(nil),      // 5: service.FindByIdRequest
	(*FindByIdResponse)(nil),     // 6: service.FindByIdResponse
	(*FindByMobileRequest)(nil),  // 7: service.FindByMobileRequest
	(*FindByMobileResponse)(nil), // 8: service.FindByMobileResponse
	(*SendSmsRequest)(nil),       // 9: service.SendSmsRequest
	(*SendSmsResponse)(nil),      // 10: service.SendSmsResponse
}
var file_user_proto_depIdxs = []int32{
	0,  // 0: service.ListByIdsResponse.data:type_name -> service.UserInfoResponse
	1,  // 1: service.User.Register:input_type -> service.RegisterRequest
	5,  // 2: service.User.FindById:input_type -> service.FindByIdRequest
	3,  // 3: service.User.ListByIds:input_type -> service.ListByIdsRequest
	7,  // 4: service.User.FindByMobile:input_type -> service.FindByMobileRequest
	9,  // 5: service.User.SendSms:input_type -> service.SendSmsRequest
	2,  // 6: service.User.Register:output_type -> service.RegisterResponse
	6,  // 7: service.User.FindById:output_type -> service.FindByIdResponse
	4,  // 8: service.User.ListByIds:output_type -> service.ListByIdsResponse
	8,  // 9: service.User.FindByMobile:output_type -> service.FindByMobileResponse
	10, // 10: service.User.SendSms:output_type -> service.SendSmsResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResponse); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListByIdsRequest); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListByIdsResponse); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByIdRequest); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByIdResponse); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByMobileRequest); i {
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
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByMobileResponse); i {
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
		file_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsRequest); i {
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
		file_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsResponse); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
