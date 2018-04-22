// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bases.proto

package zproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ListLoadMode
type EnumHelper_ListLoadMode int32

const (
	EnumHelper_ListLoadMode_NONE EnumHelper_ListLoadMode = 0
	EnumHelper_FOWRARD           EnumHelper_ListLoadMode = 1
	EnumHelper_BACKWARD          EnumHelper_ListLoadMode = 2
	EnumHelper_BOTH              EnumHelper_ListLoadMode = 3
)

var EnumHelper_ListLoadMode_name = map[int32]string{
	0: "ListLoadMode_NONE",
	1: "FOWRARD",
	2: "BACKWARD",
	3: "BOTH",
}
var EnumHelper_ListLoadMode_value = map[string]int32{
	"ListLoadMode_NONE": 0,
	"FOWRARD":           1,
	"BACKWARD":          2,
	"BOTH":              3,
}

func (x EnumHelper_ListLoadMode) String() string {
	return proto.EnumName(EnumHelper_ListLoadMode_name, int32(x))
}
func (EnumHelper_ListLoadMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

// MessageState
type EnumHelper_MessageState int32

const (
	EnumHelper_MessageState_NONE EnumHelper_MessageState = 0
	EnumHelper_SENT              EnumHelper_MessageState = 1
	EnumHelper_RECEIVED          EnumHelper_MessageState = 2
	EnumHelper_READ              EnumHelper_MessageState = 3
)

var EnumHelper_MessageState_name = map[int32]string{
	0: "MessageState_NONE",
	1: "SENT",
	2: "RECEIVED",
	3: "READ",
}
var EnumHelper_MessageState_value = map[string]int32{
	"MessageState_NONE": 0,
	"SENT":              1,
	"RECEIVED":          2,
	"READ":              3,
}

func (x EnumHelper_MessageState) String() string {
	return proto.EnumName(EnumHelper_MessageState_name, int32(x))
}
func (EnumHelper_MessageState) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 1} }

// 消息类型
type EnumHelper_MessageType int32

const (
	EnumHelper_MessageType_NONE EnumHelper_MessageType = 0
	EnumHelper_TEXT             EnumHelper_MessageType = 1
	EnumHelper_AUDIO            EnumHelper_MessageType = 2
	EnumHelper_VIDEO            EnumHelper_MessageType = 3
)

var EnumHelper_MessageType_name = map[int32]string{
	0: "MessageType_NONE",
	1: "TEXT",
	2: "AUDIO",
	3: "VIDEO",
}
var EnumHelper_MessageType_value = map[string]int32{
	"MessageType_NONE": 0,
	"TEXT":             1,
	"AUDIO":            2,
	"VIDEO":            3,
}

func (x EnumHelper_MessageType) String() string {
	return proto.EnumName(EnumHelper_MessageType_name, int32(x))
}
func (EnumHelper_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 2} }

type EnumHelper struct {
}

func (m *EnumHelper) Reset()                    { *m = EnumHelper{} }
func (m *EnumHelper) String() string            { return proto.CompactTextString(m) }
func (*EnumHelper) ProtoMessage()               {}
func (*EnumHelper) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type UserID struct {
	AppId  uint32 `protobuf:"varint,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *UserID) Reset()                    { *m = UserID{} }
func (m *UserID) String() string            { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()               {}
func (*UserID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *UserID) GetAppId() uint32 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *UserID) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type OnlineSessionStatusEntry struct {
	AppId  uint32 `protobuf:"varint,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *OnlineSessionStatusEntry) Reset()                    { *m = OnlineSessionStatusEntry{} }
func (m *OnlineSessionStatusEntry) String() string            { return proto.CompactTextString(m) }
func (*OnlineSessionStatusEntry) ProtoMessage()               {}
func (*OnlineSessionStatusEntry) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *OnlineSessionStatusEntry) GetAppId() uint32 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *OnlineSessionStatusEntry) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// 连接认证
type UserToeken struct {
	AppKey    string `protobuf:"bytes,1,opt,name=app_key,json=appKey" json:"app_key,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserToken string `protobuf:"bytes,3,opt,name=user_token,json=userToken" json:"user_token,omitempty"`
}

func (m *UserToeken) Reset()                    { *m = UserToeken{} }
func (m *UserToeken) String() string            { return proto.CompactTextString(m) }
func (*UserToeken) ProtoMessage()               {}
func (*UserToeken) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *UserToeken) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *UserToeken) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserToeken) GetUserToken() string {
	if m != nil {
		return m.UserToken
	}
	return ""
}

// 用户信息
type UserInfo struct {
	AppId    uint32 `protobuf:"varint,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	PushName string `protobuf:"bytes,3,opt,name=push_name,json=pushName" json:"push_name,omitempty"`
	Avatar   string `protobuf:"bytes,4,opt,name=avatar" json:"avatar,omitempty"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *UserInfo) GetAppId() uint32 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *UserInfo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserInfo) GetPushName() string {
	if m != nil {
		return m.PushName
	}
	return ""
}

func (m *UserInfo) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type OnlineUser struct {
	AppId    uint32 `protobuf:"varint,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ServerId uint32 `protobuf:"varint,3,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	ConnId   uint64 `protobuf:"varint,4,opt,name=conn_id,json=connId" json:"conn_id,omitempty"`
}

func (m *OnlineUser) Reset()                    { *m = OnlineUser{} }
func (m *OnlineUser) String() string            { return proto.CompactTextString(m) }
func (*OnlineUser) ProtoMessage()               {}
func (*OnlineUser) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *OnlineUser) GetAppId() uint32 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *OnlineUser) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *OnlineUser) GetServerId() uint32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *OnlineUser) GetConnId() uint64 {
	if m != nil {
		return m.ConnId
	}
	return 0
}

// 聊天消息内容
type MessageContainer struct {
	MessageId       uint64                  `protobuf:"varint,1,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	SenderUserId    string                  `protobuf:"bytes,2,opt,name=sender_user_id,json=senderUserId" json:"sender_user_id,omitempty"`
	Peer            *Peer                   `protobuf:"bytes,3,opt,name=peer" json:"peer,omitempty"`
	ClientMessageId uint64                  `protobuf:"varint,4,opt,name=client_message_id,json=clientMessageId" json:"client_message_id,omitempty"`
	MessageSeq      uint64                  `protobuf:"varint,5,opt,name=message_seq,json=messageSeq" json:"message_seq,omitempty"`
	State           EnumHelper_MessageState `protobuf:"varint,6,opt,name=state,enum=zproto.EnumHelper_MessageState" json:"state,omitempty"`
	MessageType     EnumHelper_MessageType  `protobuf:"varint,7,opt,name=message_type,json=messageType,enum=zproto.EnumHelper_MessageType" json:"message_type,omitempty"`
	MessageContent  []byte                  `protobuf:"bytes,8,opt,name=message_content,json=messageContent,proto3" json:"message_content,omitempty"`
	PassthroughData []byte                  `protobuf:"bytes,9,opt,name=passthrough_data,json=passthroughData,proto3" json:"passthrough_data,omitempty"`
	UpdatedAt       uint64                  `protobuf:"varint,10,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *MessageContainer) Reset()                    { *m = MessageContainer{} }
func (m *MessageContainer) String() string            { return proto.CompactTextString(m) }
func (*MessageContainer) ProtoMessage()               {}
func (*MessageContainer) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *MessageContainer) GetMessageId() uint64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *MessageContainer) GetSenderUserId() string {
	if m != nil {
		return m.SenderUserId
	}
	return ""
}

func (m *MessageContainer) GetPeer() *Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *MessageContainer) GetClientMessageId() uint64 {
	if m != nil {
		return m.ClientMessageId
	}
	return 0
}

func (m *MessageContainer) GetMessageSeq() uint64 {
	if m != nil {
		return m.MessageSeq
	}
	return 0
}

func (m *MessageContainer) GetState() EnumHelper_MessageState {
	if m != nil {
		return m.State
	}
	return EnumHelper_MessageState_NONE
}

func (m *MessageContainer) GetMessageType() EnumHelper_MessageType {
	if m != nil {
		return m.MessageType
	}
	return EnumHelper_MessageType_NONE
}

func (m *MessageContainer) GetMessageContent() []byte {
	if m != nil {
		return m.MessageContent
	}
	return nil
}

func (m *MessageContainer) GetPassthroughData() []byte {
	if m != nil {
		return m.PassthroughData
	}
	return nil
}

func (m *MessageContainer) GetUpdatedAt() uint64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

// Conversation from history
type Dialog struct {
	Peer *Peer `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
}

func (m *Dialog) Reset()                    { *m = Dialog{} }
func (m *Dialog) String() string            { return proto.CompactTextString(m) }
func (*Dialog) ProtoMessage()               {}
func (*Dialog) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *Dialog) GetPeer() *Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

func init() {
	proto.RegisterType((*EnumHelper)(nil), "zproto.EnumHelper")
	proto.RegisterType((*UserID)(nil), "zproto.UserID")
	proto.RegisterType((*OnlineSessionStatusEntry)(nil), "zproto.OnlineSessionStatusEntry")
	proto.RegisterType((*UserToeken)(nil), "zproto.UserToeken")
	proto.RegisterType((*UserInfo)(nil), "zproto.UserInfo")
	proto.RegisterType((*OnlineUser)(nil), "zproto.OnlineUser")
	proto.RegisterType((*MessageContainer)(nil), "zproto.MessageContainer")
	proto.RegisterType((*Dialog)(nil), "zproto.Dialog")
	proto.RegisterEnum("zproto.EnumHelper_ListLoadMode", EnumHelper_ListLoadMode_name, EnumHelper_ListLoadMode_value)
	proto.RegisterEnum("zproto.EnumHelper_MessageState", EnumHelper_MessageState_name, EnumHelper_MessageState_value)
	proto.RegisterEnum("zproto.EnumHelper_MessageType", EnumHelper_MessageType_name, EnumHelper_MessageType_value)
}

func init() { proto.RegisterFile("bases.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x4f, 0xdb, 0x4a,
	0x10, 0xbd, 0xce, 0x87, 0x49, 0x26, 0xb9, 0x60, 0x56, 0x97, 0x8b, 0xd5, 0xaa, 0x25, 0x72, 0x2b,
	0x35, 0xe5, 0x21, 0x0f, 0x54, 0x95, 0xfa, 0x9a, 0x0f, 0xb7, 0x98, 0x8f, 0x04, 0x39, 0x01, 0xaa,
	0x4a, 0x95, 0xb5, 0xc4, 0xd3, 0xc4, 0x4a, 0xb2, 0x6b, 0xbc, 0x1b, 0xa4, 0xf0, 0xb7, 0xfa, 0x9f,
	0xfa, 0x3b, 0xaa, 0x5d, 0xdb, 0xe0, 0x3e, 0xf0, 0x40, 0x9f, 0x9c, 0x39, 0x33, 0xe7, 0xcc, 0xc9,
	0xd8, 0x07, 0x1a, 0x37, 0x54, 0xa0, 0xe8, 0xc4, 0x09, 0x97, 0x9c, 0x98, 0xf7, 0xfa, 0xf9, 0xa2,
	0x11, 0x23, 0x26, 0x19, 0xe8, 0xfc, 0x32, 0x00, 0x5c, 0xb6, 0x5e, 0x1d, 0xe3, 0x32, 0xc6, 0xc4,
	0x39, 0x81, 0xe6, 0x59, 0x24, 0xe4, 0x19, 0xa7, 0xe1, 0x39, 0x0f, 0x91, 0xec, 0xc1, 0x6e, 0xb1,
	0x0e, 0x86, 0xa3, 0xa1, 0x6b, 0xfd, 0x43, 0x1a, 0xb0, 0xf5, 0x79, 0x74, 0xed, 0x77, 0xfd, 0x81,
	0x65, 0x90, 0x26, 0xd4, 0x7a, 0xdd, 0xfe, 0xe9, 0xb5, 0xaa, 0x4a, 0xa4, 0x06, 0x95, 0xde, 0x68,
	0x72, 0x6c, 0x95, 0x9d, 0x2f, 0xd0, 0x3c, 0x47, 0x21, 0xe8, 0x0c, 0xc7, 0x92, 0x4a, 0xad, 0x55,
	0xac, 0x73, 0xad, 0x1a, 0x54, 0xc6, 0xee, 0x70, 0x92, 0x0a, 0xf9, 0x6e, 0xdf, 0xf5, 0xae, 0xdc,
	0x4c, 0xc8, 0x77, 0xbb, 0x03, 0xab, 0xec, 0xf4, 0xa1, 0x91, 0x11, 0x27, 0x9b, 0x18, 0xc9, 0x7f,
	0x60, 0x15, 0xca, 0x82, 0xcc, 0xc4, 0xfd, 0xaa, 0x64, 0xea, 0x50, 0xed, 0x5e, 0x0e, 0xbc, 0x91,
	0x55, 0x52, 0x3f, 0xaf, 0xbc, 0x81, 0x3b, 0xb2, 0xca, 0xce, 0x27, 0x30, 0x2f, 0x05, 0x26, 0xde,
	0x80, 0xec, 0x81, 0x49, 0xe3, 0x38, 0x88, 0x42, 0xdb, 0x68, 0x19, 0xed, 0x7f, 0xfd, 0x2a, 0x8d,
	0x63, 0x2f, 0x24, 0xfb, 0xb0, 0xb5, 0x16, 0x98, 0x28, 0xbc, 0xd4, 0x32, 0xda, 0x75, 0xdf, 0x54,
	0xa5, 0x17, 0x3a, 0x27, 0x60, 0x8f, 0xd8, 0x32, 0x62, 0x38, 0x46, 0x21, 0x22, 0xce, 0x94, 0xfb,
	0xb5, 0x70, 0x99, 0x4c, 0x36, 0xcf, 0xd6, 0xfa, 0x0e, 0xa0, 0x5c, 0x4c, 0x38, 0x2e, 0x90, 0xa9,
	0x31, 0xc5, 0x5e, 0xe0, 0x46, 0xd3, 0xeb, 0xbe, 0x12, 0x3b, 0xc5, 0xcd, 0x93, 0x7c, 0xf2, 0x0a,
	0x40, 0x37, 0x24, 0x5f, 0x20, 0xb3, 0xcb, 0xba, 0x57, 0x5f, 0x6b, 0xc5, 0x05, 0x32, 0x87, 0x43,
	0x4d, 0xff, 0x49, 0xf6, 0x83, 0x3f, 0xd7, 0x1a, 0x79, 0x09, 0xf5, 0x78, 0x2d, 0xe6, 0x01, 0xa3,
	0x2b, 0xcc, 0x94, 0x6b, 0x0a, 0x18, 0xd2, 0x15, 0x92, 0xff, 0xc1, 0xa4, 0x77, 0x54, 0xd2, 0xc4,
	0xae, 0x64, 0x46, 0x75, 0xe5, 0x24, 0x00, 0xe9, 0x6d, 0xd4, 0xda, 0xbf, 0x59, 0x29, 0x30, 0xb9,
	0x4b, 0x5b, 0x65, 0x4d, 0xa9, 0xa5, 0x40, 0xca, 0x9a, 0x72, 0xc6, 0x54, 0x4b, 0xed, 0xac, 0xf8,
	0xa6, 0x2a, 0xbd, 0xd0, 0xf9, 0x59, 0x7e, 0xf8, 0x00, 0xfa, 0x9c, 0x49, 0x1a, 0x31, 0x4c, 0xd4,
	0x61, 0x56, 0x29, 0x96, 0xaf, 0xaf, 0xf8, 0xf5, 0x0c, 0xf1, 0x42, 0xf2, 0x16, 0xb6, 0x05, 0xb2,
	0x10, 0x93, 0xe0, 0x4f, 0x27, 0xcd, 0x14, 0xbd, 0x4c, 0xfd, 0xb4, 0xa0, 0xa2, 0xb2, 0xa1, 0xad,
	0x34, 0x8e, 0x9a, 0x9d, 0x34, 0x30, 0x9d, 0x0b, 0xc4, 0xc4, 0xd7, 0x1d, 0x72, 0x08, 0xbb, 0xd3,
	0x65, 0x84, 0x4c, 0x06, 0x85, 0x6d, 0xa9, 0xbd, 0x9d, 0xb4, 0x71, 0xfe, 0xb0, 0xf3, 0x00, 0x1a,
	0xf9, 0x90, 0xc0, 0x5b, 0xbb, 0xaa, 0xa7, 0x72, 0x97, 0x63, 0xbc, 0x25, 0x1f, 0xa1, 0x2a, 0x54,
	0x12, 0x6c, 0xb3, 0x65, 0xb4, 0xb7, 0x8f, 0x0e, 0xf2, 0x7d, 0x8f, 0x79, 0xec, 0x14, 0x03, 0xe3,
	0xa7, 0xd3, 0xa4, 0x0b, 0xcd, 0x5c, 0x57, 0x6e, 0x62, 0xb4, 0xb7, 0x34, 0xfb, 0xf5, 0xd3, 0x6c,
	0x15, 0x13, 0x3f, 0xf7, 0xa2, 0x23, 0xf4, 0x0e, 0x76, 0x72, 0x89, 0x29, 0x67, 0x12, 0x99, 0xb4,
	0x6b, 0x2d, 0xa3, 0xdd, 0xf4, 0xb7, 0x57, 0x8f, 0x87, 0x45, 0x26, 0xc9, 0x7b, 0xb0, 0x62, 0x2a,
	0x84, 0x9c, 0x27, 0x7c, 0x3d, 0x9b, 0x07, 0x21, 0x95, 0xd4, 0xae, 0xeb, 0xc9, 0x9d, 0x02, 0x3e,
	0xa0, 0x92, 0xea, 0x4f, 0x33, 0x0e, 0xa9, 0xc4, 0x30, 0xa0, 0xd2, 0x86, 0xf4, 0x0d, 0x64, 0x48,
	0x57, 0x3a, 0x87, 0x60, 0x0e, 0x22, 0xba, 0xe4, 0xb3, 0x87, 0x2b, 0x1b, 0x4f, 0x5d, 0xb9, 0xf7,
	0x06, 0xf6, 0xa7, 0x7c, 0xd5, 0xb9, 0x9f, 0xce, 0xa9, 0xec, 0x20, 0x9b, 0x45, 0x0c, 0xb3, 0xa9,
	0x9e, 0xf9, 0xed, 0x42, 0x3d, 0x8f, 0x4b, 0x37, 0xa6, 0x06, 0x3e, 0xfc, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xb2, 0x5e, 0xfb, 0x4c, 0xe4, 0x04, 0x00, 0x00,
}