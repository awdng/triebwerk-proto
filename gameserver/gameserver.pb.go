// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gameserver.proto

package gameserver

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A Player representation in the GameServerState.
type Player struct {
	// The name of the player.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The score of the player.
	Score int32 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	// The team of the player.
	Team                 int32    `protobuf:"varint,3,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a199f3261f0fcf, []int{0}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Player) GetTeam() int32 {
	if m != nil {
		return m.Team
	}
	return 0
}

// The state of the server sent as a Heartbeat
type ServerState struct {
	// The ID of the server
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The IP and port this server is running on.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// The elapsed time in the game in seconds.
	ElapsedTime int32 `protobuf:"varint,3,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	// The timestamp of the last heartbeat.
	UpdatedAt int32 `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// The list of players on this server.
	Players              []*Player `protobuf:"bytes,5,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ServerState) Reset()         { *m = ServerState{} }
func (m *ServerState) String() string { return proto.CompactTextString(m) }
func (*ServerState) ProtoMessage()    {}
func (*ServerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a199f3261f0fcf, []int{1}
}

func (m *ServerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerState.Unmarshal(m, b)
}
func (m *ServerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerState.Marshal(b, m, deterministic)
}
func (m *ServerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerState.Merge(m, src)
}
func (m *ServerState) XXX_Size() int {
	return xxx_messageInfo_ServerState.Size(m)
}
func (m *ServerState) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerState.DiscardUnknown(m)
}

var xxx_messageInfo_ServerState proto.InternalMessageInfo

func (m *ServerState) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServerState) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ServerState) GetElapsedTime() int32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func (m *ServerState) GetUpdatedAt() int32 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *ServerState) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

// The list of servers
type ServerList struct {
	// The list of server states
	Servers              []*ServerState `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ServerList) Reset()         { *m = ServerList{} }
func (m *ServerList) String() string { return proto.CompactTextString(m) }
func (*ServerList) ProtoMessage()    {}
func (*ServerList) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a199f3261f0fcf, []int{2}
}

func (m *ServerList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerList.Unmarshal(m, b)
}
func (m *ServerList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerList.Marshal(b, m, deterministic)
}
func (m *ServerList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerList.Merge(m, src)
}
func (m *ServerList) XXX_Size() int {
	return xxx_messageInfo_ServerList.Size(m)
}
func (m *ServerList) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerList.DiscardUnknown(m)
}

var xxx_messageInfo_ServerList proto.InternalMessageInfo

func (m *ServerList) GetServers() []*ServerState {
	if m != nil {
		return m.Servers
	}
	return nil
}

// GetServer retrieves the state of a single server
type GetServerRequest struct {
	// The ID of the server
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServerRequest) Reset()         { *m = GetServerRequest{} }
func (m *GetServerRequest) String() string { return proto.CompactTextString(m) }
func (*GetServerRequest) ProtoMessage()    {}
func (*GetServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a199f3261f0fcf, []int{3}
}

func (m *GetServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServerRequest.Unmarshal(m, b)
}
func (m *GetServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServerRequest.Marshal(b, m, deterministic)
}
func (m *GetServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServerRequest.Merge(m, src)
}
func (m *GetServerRequest) XXX_Size() int {
	return xxx_messageInfo_GetServerRequest.Size(m)
}
func (m *GetServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServerRequest proto.InternalMessageInfo

func (m *GetServerRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// GetServer retrieves the state of a single server
type GetServerListRequest struct {
	// Which region the servers are located in
	Region               string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServerListRequest) Reset()         { *m = GetServerListRequest{} }
func (m *GetServerListRequest) String() string { return proto.CompactTextString(m) }
func (*GetServerListRequest) ProtoMessage()    {}
func (*GetServerListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a199f3261f0fcf, []int{4}
}

func (m *GetServerListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServerListRequest.Unmarshal(m, b)
}
func (m *GetServerListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServerListRequest.Marshal(b, m, deterministic)
}
func (m *GetServerListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServerListRequest.Merge(m, src)
}
func (m *GetServerListRequest) XXX_Size() int {
	return xxx_messageInfo_GetServerListRequest.Size(m)
}
func (m *GetServerListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServerListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServerListRequest proto.InternalMessageInfo

func (m *GetServerListRequest) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

type ServerStateRequest struct {
	// The state as sent from the game server
	State                *ServerState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ServerStateRequest) Reset()         { *m = ServerStateRequest{} }
func (m *ServerStateRequest) String() string { return proto.CompactTextString(m) }
func (*ServerStateRequest) ProtoMessage()    {}
func (*ServerStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a199f3261f0fcf, []int{5}
}

func (m *ServerStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStateRequest.Unmarshal(m, b)
}
func (m *ServerStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStateRequest.Marshal(b, m, deterministic)
}
func (m *ServerStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStateRequest.Merge(m, src)
}
func (m *ServerStateRequest) XXX_Size() int {
	return xxx_messageInfo_ServerStateRequest.Size(m)
}
func (m *ServerStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStateRequest proto.InternalMessageInfo

func (m *ServerStateRequest) GetState() *ServerState {
	if m != nil {
		return m.State
	}
	return nil
}

// Sent from Gameserver to register themselves
type ServerRegisterRequest struct {
	// The IP address and port combination of the server
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerRegisterRequest) Reset()         { *m = ServerRegisterRequest{} }
func (m *ServerRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*ServerRegisterRequest) ProtoMessage()    {}
func (*ServerRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a199f3261f0fcf, []int{6}
}

func (m *ServerRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerRegisterRequest.Unmarshal(m, b)
}
func (m *ServerRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerRegisterRequest.Marshal(b, m, deterministic)
}
func (m *ServerRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerRegisterRequest.Merge(m, src)
}
func (m *ServerRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_ServerRegisterRequest.Size(m)
}
func (m *ServerRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerRegisterRequest proto.InternalMessageInfo

func (m *ServerRegisterRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// The Gameserver register reponse
type ServerRegisterResponse struct {
	// The associated id of the server
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerRegisterResponse) Reset()         { *m = ServerRegisterResponse{} }
func (m *ServerRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*ServerRegisterResponse) ProtoMessage()    {}
func (*ServerRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a199f3261f0fcf, []int{7}
}

func (m *ServerRegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerRegisterResponse.Unmarshal(m, b)
}
func (m *ServerRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerRegisterResponse.Marshal(b, m, deterministic)
}
func (m *ServerRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerRegisterResponse.Merge(m, src)
}
func (m *ServerRegisterResponse) XXX_Size() int {
	return xxx_messageInfo_ServerRegisterResponse.Size(m)
}
func (m *ServerRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerRegisterResponse proto.InternalMessageInfo

func (m *ServerRegisterResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EndGameRequest struct {
	// The state as sent from the game server
	State                *ServerState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *EndGameRequest) Reset()         { *m = EndGameRequest{} }
func (m *EndGameRequest) String() string { return proto.CompactTextString(m) }
func (*EndGameRequest) ProtoMessage()    {}
func (*EndGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a199f3261f0fcf, []int{8}
}

func (m *EndGameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndGameRequest.Unmarshal(m, b)
}
func (m *EndGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndGameRequest.Marshal(b, m, deterministic)
}
func (m *EndGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndGameRequest.Merge(m, src)
}
func (m *EndGameRequest) XXX_Size() int {
	return xxx_messageInfo_EndGameRequest.Size(m)
}
func (m *EndGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndGameRequest proto.InternalMessageInfo

func (m *EndGameRequest) GetState() *ServerState {
	if m != nil {
		return m.State
	}
	return nil
}

type EndGameResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndGameResponse) Reset()         { *m = EndGameResponse{} }
func (m *EndGameResponse) String() string { return proto.CompactTextString(m) }
func (*EndGameResponse) ProtoMessage()    {}
func (*EndGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a199f3261f0fcf, []int{9}
}

func (m *EndGameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndGameResponse.Unmarshal(m, b)
}
func (m *EndGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndGameResponse.Marshal(b, m, deterministic)
}
func (m *EndGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndGameResponse.Merge(m, src)
}
func (m *EndGameResponse) XXX_Size() int {
	return xxx_messageInfo_EndGameResponse.Size(m)
}
func (m *EndGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EndGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EndGameResponse proto.InternalMessageInfo

type AuthorizePlayerRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizePlayerRequest) Reset()         { *m = AuthorizePlayerRequest{} }
func (m *AuthorizePlayerRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizePlayerRequest) ProtoMessage()    {}
func (*AuthorizePlayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a199f3261f0fcf, []int{10}
}

func (m *AuthorizePlayerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizePlayerRequest.Unmarshal(m, b)
}
func (m *AuthorizePlayerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizePlayerRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizePlayerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizePlayerRequest.Merge(m, src)
}
func (m *AuthorizePlayerRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizePlayerRequest.Size(m)
}
func (m *AuthorizePlayerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizePlayerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizePlayerRequest proto.InternalMessageInfo

func (m *AuthorizePlayerRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthorizePlayerRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AuthorizePlayerResponse struct {
	Authorized           bool     `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizePlayerResponse) Reset()         { *m = AuthorizePlayerResponse{} }
func (m *AuthorizePlayerResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorizePlayerResponse) ProtoMessage()    {}
func (*AuthorizePlayerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a199f3261f0fcf, []int{11}
}

func (m *AuthorizePlayerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizePlayerResponse.Unmarshal(m, b)
}
func (m *AuthorizePlayerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizePlayerResponse.Marshal(b, m, deterministic)
}
func (m *AuthorizePlayerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizePlayerResponse.Merge(m, src)
}
func (m *AuthorizePlayerResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorizePlayerResponse.Size(m)
}
func (m *AuthorizePlayerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizePlayerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizePlayerResponse proto.InternalMessageInfo

func (m *AuthorizePlayerResponse) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

func (m *AuthorizePlayerResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Player)(nil), "gameserver.Player")
	proto.RegisterType((*ServerState)(nil), "gameserver.ServerState")
	proto.RegisterType((*ServerList)(nil), "gameserver.ServerList")
	proto.RegisterType((*GetServerRequest)(nil), "gameserver.GetServerRequest")
	proto.RegisterType((*GetServerListRequest)(nil), "gameserver.GetServerListRequest")
	proto.RegisterType((*ServerStateRequest)(nil), "gameserver.ServerStateRequest")
	proto.RegisterType((*ServerRegisterRequest)(nil), "gameserver.ServerRegisterRequest")
	proto.RegisterType((*ServerRegisterResponse)(nil), "gameserver.ServerRegisterResponse")
	proto.RegisterType((*EndGameRequest)(nil), "gameserver.EndGameRequest")
	proto.RegisterType((*EndGameResponse)(nil), "gameserver.EndGameResponse")
	proto.RegisterType((*AuthorizePlayerRequest)(nil), "gameserver.AuthorizePlayerRequest")
	proto.RegisterType((*AuthorizePlayerResponse)(nil), "gameserver.AuthorizePlayerResponse")
}

func init() { proto.RegisterFile("gameserver.proto", fileDescriptor_97a199f3261f0fcf) }

var fileDescriptor_97a199f3261f0fcf = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xae, 0xd3, 0x26, 0xf9, 0x65, 0xf2, 0x4b, 0x1a, 0x86, 0x90, 0x5a, 0xa6, 0x54, 0xe9, 0x72,
	0xc9, 0x01, 0x12, 0xa5, 0xdc, 0xa9, 0x2a, 0xfe, 0x14, 0x10, 0x48, 0xc8, 0xe1, 0x82, 0x84, 0xa8,
	0xb6, 0x78, 0x14, 0x2c, 0x62, 0xaf, 0xf1, 0x6e, 0x2a, 0x01, 0xe2, 0xc2, 0x2b, 0xf0, 0x02, 0xdc,
	0x78, 0x20, 0x5e, 0x81, 0x07, 0x41, 0xde, 0x5d, 0xa7, 0x8e, 0x93, 0x06, 0x89, 0x9b, 0x67, 0xf6,
	0x9b, 0xf9, 0xbe, 0x99, 0xf9, 0x64, 0xe8, 0x4c, 0x79, 0x44, 0x92, 0xd2, 0x0b, 0x4a, 0x87, 0x49,
	0x2a, 0x94, 0x40, 0xb8, 0xcc, 0x78, 0xfb, 0x53, 0x21, 0xa6, 0x33, 0x1a, 0xf1, 0x24, 0x1c, 0xf1,
	0x38, 0x16, 0x8a, 0xab, 0x50, 0xc4, 0xd2, 0x20, 0xd9, 0x63, 0xa8, 0xbd, 0x9c, 0xf1, 0x4f, 0x94,
	0x22, 0xc2, 0x4e, 0xcc, 0x23, 0x72, 0x9d, 0xbe, 0x33, 0x68, 0xf8, 0xfa, 0x1b, 0xbb, 0x50, 0x95,
	0xef, 0x44, 0x4a, 0x6e, 0xa5, 0xef, 0x0c, 0xaa, 0xbe, 0x09, 0x32, 0xa4, 0x22, 0x1e, 0xb9, 0xdb,
	0x3a, 0xa9, 0xbf, 0xd9, 0x4f, 0x07, 0x9a, 0x13, 0x4d, 0x38, 0x51, 0x5c, 0x11, 0xb6, 0xa1, 0x12,
	0x06, 0xb6, 0x57, 0x25, 0x0c, 0xd0, 0x85, 0x3a, 0x0f, 0x82, 0x94, 0xa4, 0xd4, 0xbd, 0x1a, 0x7e,
	0x1e, 0xe2, 0x21, 0xfc, 0x4f, 0x33, 0x9e, 0x48, 0x0a, 0xce, 0x54, 0x18, 0x91, 0xed, 0xda, 0xb4,
	0xb9, 0x57, 0x61, 0x44, 0x78, 0x0b, 0x60, 0x9e, 0x04, 0x5c, 0x51, 0x70, 0xc6, 0x95, 0xbb, 0xa3,
	0x01, 0x0d, 0x9b, 0x39, 0x51, 0x78, 0x07, 0xea, 0x89, 0x9e, 0x41, 0xba, 0xd5, 0xfe, 0xf6, 0xa0,
	0x79, 0x84, 0xc3, 0xc2, 0x46, 0xcc, 0x78, 0x7e, 0x0e, 0x61, 0xc7, 0x00, 0x46, 0xe8, 0xf3, 0x50,
	0x2a, 0x1c, 0x43, 0xdd, 0xe0, 0xa4, 0xeb, 0xe8, 0xda, 0xbd, 0x62, 0x6d, 0x61, 0x22, 0x3f, 0xc7,
	0x31, 0x06, 0x9d, 0x53, 0x52, 0xe6, 0xc9, 0xa7, 0x8f, 0x73, 0x92, 0xaa, 0x3c, 0x2e, 0x1b, 0x42,
	0x77, 0x81, 0xc9, 0x78, 0x72, 0x5c, 0x0f, 0x6a, 0x29, 0x4d, 0x43, 0x11, 0x5b, 0xac, 0x8d, 0xd8,
	0x03, 0xc0, 0x22, 0x97, 0x45, 0xdf, 0x85, 0xaa, 0xcc, 0x62, 0x0d, 0xde, 0x20, 0xcd, 0xa0, 0xd8,
	0x18, 0x6e, 0xe4, 0xaa, 0xa6, 0xa1, 0x54, 0x97, 0xea, 0x0a, 0xcb, 0x77, 0x96, 0x96, 0xcf, 0x06,
	0xd0, 0x2b, 0x97, 0xc8, 0x44, 0xc4, 0x72, 0xe5, 0x80, 0xec, 0x18, 0xda, 0x8f, 0xe2, 0xe0, 0x94,
	0x47, 0xff, 0xaa, 0xee, 0x1a, 0xec, 0x2e, 0x1a, 0x18, 0x0e, 0x76, 0x1f, 0x7a, 0x27, 0x73, 0xf5,
	0x5e, 0xa4, 0xe1, 0x67, 0xb2, 0x67, 0xb2, 0xbd, 0xbb, 0x50, 0x55, 0xe2, 0x03, 0xe5, 0x6b, 0x32,
	0x81, 0xd5, 0x54, 0x59, 0x68, 0x7a, 0x0a, 0x7b, 0x2b, 0xf5, 0x56, 0xfe, 0x01, 0x00, 0xcf, 0x9f,
	0xcc, 0x18, 0xff, 0xf9, 0x85, 0x4c, 0xb9, 0xd5, 0xd1, 0x8f, 0x1d, 0xe8, 0x64, 0xda, 0x8c, 0xf0,
	0x17, 0x3c, 0xdb, 0x05, 0x72, 0x68, 0x2f, 0xae, 0x68, 0x6c, 0xbd, 0x5f, 0x1c, 0xb2, 0xec, 0x02,
	0xef, 0xaa, 0x15, 0x30, 0xf7, 0xdb, 0xaf, 0xdf, 0xdf, 0x2b, 0x88, 0x9d, 0xd1, 0xc5, 0x78, 0x64,
	0x7d, 0x34, 0xfa, 0x12, 0x06, 0x5f, 0xf1, 0x2d, 0xb4, 0x96, 0x8c, 0x82, 0xfd, 0xb5, 0x0c, 0x05,
	0x0f, 0x79, 0xbd, 0x55, 0x96, 0xec, 0x99, 0x5d, 0xd7, 0x24, 0x2d, 0x6c, 0x16, 0x48, 0xf0, 0x35,
	0xb4, 0xf3, 0xd3, 0x1a, 0x28, 0x1e, 0xae, 0x96, 0x97, 0xfc, 0xe2, 0xb1, 0x4d, 0x10, 0x7b, 0xbb,
	0x2d, 0x7c, 0x06, 0xad, 0x09, 0xc5, 0xc1, 0x13, 0xe2, 0xa9, 0x3a, 0x27, 0xae, 0xf0, 0xe0, 0x2a,
	0x07, 0xfc, 0x6d, 0x3d, 0x5b, 0xf8, 0x10, 0xea, 0xd6, 0x1c, 0xe8, 0x15, 0x51, 0xcb, 0x96, 0xf3,
	0x6e, 0xae, 0x7d, 0x5b, 0x28, 0x7a, 0x03, 0xbb, 0x25, 0x3f, 0xe0, 0xd2, 0x28, 0xeb, 0xcd, 0xe6,
	0xdd, 0xde, 0x88, 0xc9, 0xbb, 0x9f, 0xd7, 0xf4, 0x1f, 0xf3, 0xde, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x7a, 0x5f, 0x0f, 0x50, 0x6f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameServerMasterClient is the client API for GameServerMaster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServerMasterClient interface {
	GetServerState(ctx context.Context, in *GetServerRequest, opts ...grpc.CallOption) (*ServerState, error)
	GetServerList(ctx context.Context, in *GetServerListRequest, opts ...grpc.CallOption) (*ServerList, error)
	RegisterServer(ctx context.Context, in *ServerRegisterRequest, opts ...grpc.CallOption) (*ServerRegisterResponse, error)
	SendHeartbeat(ctx context.Context, in *ServerStateRequest, opts ...grpc.CallOption) (*ServerState, error)
	EndGame(ctx context.Context, in *EndGameRequest, opts ...grpc.CallOption) (*EndGameResponse, error)
	AuthorizePlayer(ctx context.Context, in *AuthorizePlayerRequest, opts ...grpc.CallOption) (*AuthorizePlayerResponse, error)
}

type gameServerMasterClient struct {
	cc *grpc.ClientConn
}

func NewGameServerMasterClient(cc *grpc.ClientConn) GameServerMasterClient {
	return &gameServerMasterClient{cc}
}

func (c *gameServerMasterClient) GetServerState(ctx context.Context, in *GetServerRequest, opts ...grpc.CallOption) (*ServerState, error) {
	out := new(ServerState)
	err := c.cc.Invoke(ctx, "/gameserver.GameServerMaster/GetServerState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServerMasterClient) GetServerList(ctx context.Context, in *GetServerListRequest, opts ...grpc.CallOption) (*ServerList, error) {
	out := new(ServerList)
	err := c.cc.Invoke(ctx, "/gameserver.GameServerMaster/GetServerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServerMasterClient) RegisterServer(ctx context.Context, in *ServerRegisterRequest, opts ...grpc.CallOption) (*ServerRegisterResponse, error) {
	out := new(ServerRegisterResponse)
	err := c.cc.Invoke(ctx, "/gameserver.GameServerMaster/RegisterServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServerMasterClient) SendHeartbeat(ctx context.Context, in *ServerStateRequest, opts ...grpc.CallOption) (*ServerState, error) {
	out := new(ServerState)
	err := c.cc.Invoke(ctx, "/gameserver.GameServerMaster/SendHeartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServerMasterClient) EndGame(ctx context.Context, in *EndGameRequest, opts ...grpc.CallOption) (*EndGameResponse, error) {
	out := new(EndGameResponse)
	err := c.cc.Invoke(ctx, "/gameserver.GameServerMaster/EndGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServerMasterClient) AuthorizePlayer(ctx context.Context, in *AuthorizePlayerRequest, opts ...grpc.CallOption) (*AuthorizePlayerResponse, error) {
	out := new(AuthorizePlayerResponse)
	err := c.cc.Invoke(ctx, "/gameserver.GameServerMaster/AuthorizePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServerMasterServer is the server API for GameServerMaster service.
type GameServerMasterServer interface {
	GetServerState(context.Context, *GetServerRequest) (*ServerState, error)
	GetServerList(context.Context, *GetServerListRequest) (*ServerList, error)
	RegisterServer(context.Context, *ServerRegisterRequest) (*ServerRegisterResponse, error)
	SendHeartbeat(context.Context, *ServerStateRequest) (*ServerState, error)
	EndGame(context.Context, *EndGameRequest) (*EndGameResponse, error)
	AuthorizePlayer(context.Context, *AuthorizePlayerRequest) (*AuthorizePlayerResponse, error)
}

// UnimplementedGameServerMasterServer can be embedded to have forward compatible implementations.
type UnimplementedGameServerMasterServer struct {
}

func (*UnimplementedGameServerMasterServer) GetServerState(ctx context.Context, req *GetServerRequest) (*ServerState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerState not implemented")
}
func (*UnimplementedGameServerMasterServer) GetServerList(ctx context.Context, req *GetServerListRequest) (*ServerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerList not implemented")
}
func (*UnimplementedGameServerMasterServer) RegisterServer(ctx context.Context, req *ServerRegisterRequest) (*ServerRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterServer not implemented")
}
func (*UnimplementedGameServerMasterServer) SendHeartbeat(ctx context.Context, req *ServerStateRequest) (*ServerState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHeartbeat not implemented")
}
func (*UnimplementedGameServerMasterServer) EndGame(ctx context.Context, req *EndGameRequest) (*EndGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndGame not implemented")
}
func (*UnimplementedGameServerMasterServer) AuthorizePlayer(ctx context.Context, req *AuthorizePlayerRequest) (*AuthorizePlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizePlayer not implemented")
}

func RegisterGameServerMasterServer(s *grpc.Server, srv GameServerMasterServer) {
	s.RegisterService(&_GameServerMaster_serviceDesc, srv)
}

func _GameServerMaster_GetServerState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerMasterServer).GetServerState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameserver.GameServerMaster/GetServerState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerMasterServer).GetServerState(ctx, req.(*GetServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameServerMaster_GetServerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerMasterServer).GetServerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameserver.GameServerMaster/GetServerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerMasterServer).GetServerList(ctx, req.(*GetServerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameServerMaster_RegisterServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerMasterServer).RegisterServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameserver.GameServerMaster/RegisterServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerMasterServer).RegisterServer(ctx, req.(*ServerRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameServerMaster_SendHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerMasterServer).SendHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameserver.GameServerMaster/SendHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerMasterServer).SendHeartbeat(ctx, req.(*ServerStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameServerMaster_EndGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerMasterServer).EndGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameserver.GameServerMaster/EndGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerMasterServer).EndGame(ctx, req.(*EndGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameServerMaster_AuthorizePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerMasterServer).AuthorizePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameserver.GameServerMaster/AuthorizePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerMasterServer).AuthorizePlayer(ctx, req.(*AuthorizePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameServerMaster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gameserver.GameServerMaster",
	HandlerType: (*GameServerMasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerState",
			Handler:    _GameServerMaster_GetServerState_Handler,
		},
		{
			MethodName: "GetServerList",
			Handler:    _GameServerMaster_GetServerList_Handler,
		},
		{
			MethodName: "RegisterServer",
			Handler:    _GameServerMaster_RegisterServer_Handler,
		},
		{
			MethodName: "SendHeartbeat",
			Handler:    _GameServerMaster_SendHeartbeat_Handler,
		},
		{
			MethodName: "EndGame",
			Handler:    _GameServerMaster_EndGame_Handler,
		},
		{
			MethodName: "AuthorizePlayer",
			Handler:    _GameServerMaster_AuthorizePlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gameserver.proto",
}
