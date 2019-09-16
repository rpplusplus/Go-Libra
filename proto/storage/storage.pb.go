// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package storage

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SaveTransactionsRequest struct {
	// Transactions to persist.
	TxnsToCommit []*TransactionToCommit `protobuf:"bytes,1,rep,name=txns_to_commit,json=txnsToCommit,proto3" json:"txns_to_commit,omitempty"`
	// The version of the first transaction in `txns_to_commit`.
	FirstVersion uint64 `protobuf:"varint,2,opt,name=first_version,json=firstVersion,proto3" json:"first_version,omitempty"`
	// If this is set, Storage will check its state after applying the above
	// transactions matches info in this LedgerInfo before committing otherwise
	// it denies the request.
	LedgerInfoWithSignatures *LedgerInfoWithSignatures `protobuf:"bytes,3,opt,name=ledger_info_with_signatures,json=ledgerInfoWithSignatures,proto3" json:"ledger_info_with_signatures,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *SaveTransactionsRequest) Reset()         { *m = SaveTransactionsRequest{} }
func (m *SaveTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*SaveTransactionsRequest) ProtoMessage()    {}
func (*SaveTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}

func (m *SaveTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveTransactionsRequest.Unmarshal(m, b)
}
func (m *SaveTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveTransactionsRequest.Marshal(b, m, deterministic)
}
func (m *SaveTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveTransactionsRequest.Merge(m, src)
}
func (m *SaveTransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_SaveTransactionsRequest.Size(m)
}
func (m *SaveTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveTransactionsRequest proto.InternalMessageInfo

func (m *SaveTransactionsRequest) GetTxnsToCommit() []*TransactionToCommit {
	if m != nil {
		return m.TxnsToCommit
	}
	return nil
}

func (m *SaveTransactionsRequest) GetFirstVersion() uint64 {
	if m != nil {
		return m.FirstVersion
	}
	return 0
}

func (m *SaveTransactionsRequest) GetLedgerInfoWithSignatures() *LedgerInfoWithSignatures {
	if m != nil {
		return m.LedgerInfoWithSignatures
	}
	return nil
}

type SaveTransactionsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveTransactionsResponse) Reset()         { *m = SaveTransactionsResponse{} }
func (m *SaveTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*SaveTransactionsResponse) ProtoMessage()    {}
func (*SaveTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{1}
}

func (m *SaveTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveTransactionsResponse.Unmarshal(m, b)
}
func (m *SaveTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveTransactionsResponse.Marshal(b, m, deterministic)
}
func (m *SaveTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveTransactionsResponse.Merge(m, src)
}
func (m *SaveTransactionsResponse) XXX_Size() int {
	return xxx_messageInfo_SaveTransactionsResponse.Size(m)
}
func (m *SaveTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveTransactionsResponse proto.InternalMessageInfo

type GetTransactionsRequest struct {
	// The version to start with.
	StartVersion uint64 `protobuf:"varint,1,opt,name=start_version,json=startVersion,proto3" json:"start_version,omitempty"`
	// The size of the transaction batch.
	BatchSize uint64 `protobuf:"varint,2,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	// All the proofs returned in the response should be relative to this
	// given version.
	LedgerVersion uint64 `protobuf:"varint,3,opt,name=ledger_version,json=ledgerVersion,proto3" json:"ledger_version,omitempty"`
	// Used to return the events associated with each transaction
	FetchEvents          bool     `protobuf:"varint,4,opt,name=fetch_events,json=fetchEvents,proto3" json:"fetch_events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionsRequest) Reset()         { *m = GetTransactionsRequest{} }
func (m *GetTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionsRequest) ProtoMessage()    {}
func (*GetTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{2}
}

func (m *GetTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionsRequest.Unmarshal(m, b)
}
func (m *GetTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionsRequest.Marshal(b, m, deterministic)
}
func (m *GetTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionsRequest.Merge(m, src)
}
func (m *GetTransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTransactionsRequest.Size(m)
}
func (m *GetTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionsRequest proto.InternalMessageInfo

func (m *GetTransactionsRequest) GetStartVersion() uint64 {
	if m != nil {
		return m.StartVersion
	}
	return 0
}

func (m *GetTransactionsRequest) GetBatchSize() uint64 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

func (m *GetTransactionsRequest) GetLedgerVersion() uint64 {
	if m != nil {
		return m.LedgerVersion
	}
	return 0
}

func (m *GetTransactionsRequest) GetFetchEvents() bool {
	if m != nil {
		return m.FetchEvents
	}
	return false
}

type GetTransactionsResponse struct {
	TxnListWithProof     *TransactionListWithProof `protobuf:"bytes,1,opt,name=txn_list_with_proof,json=txnListWithProof,proto3" json:"txn_list_with_proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetTransactionsResponse) Reset()         { *m = GetTransactionsResponse{} }
func (m *GetTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetTransactionsResponse) ProtoMessage()    {}
func (*GetTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{3}
}

func (m *GetTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionsResponse.Unmarshal(m, b)
}
func (m *GetTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionsResponse.Marshal(b, m, deterministic)
}
func (m *GetTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionsResponse.Merge(m, src)
}
func (m *GetTransactionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetTransactionsResponse.Size(m)
}
func (m *GetTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionsResponse proto.InternalMessageInfo

func (m *GetTransactionsResponse) GetTxnListWithProof() *TransactionListWithProof {
	if m != nil {
		return m.TxnListWithProof
	}
	return nil
}

type GetAccountStateWithProofByVersionRequest struct {
	/// The account address to query with.
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	/// The version the query is based on.
	Version              uint64   `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountStateWithProofByVersionRequest) Reset() {
	*m = GetAccountStateWithProofByVersionRequest{}
}
func (m *GetAccountStateWithProofByVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountStateWithProofByVersionRequest) ProtoMessage()    {}
func (*GetAccountStateWithProofByVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{4}
}

func (m *GetAccountStateWithProofByVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountStateWithProofByVersionRequest.Unmarshal(m, b)
}
func (m *GetAccountStateWithProofByVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountStateWithProofByVersionRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountStateWithProofByVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountStateWithProofByVersionRequest.Merge(m, src)
}
func (m *GetAccountStateWithProofByVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountStateWithProofByVersionRequest.Size(m)
}
func (m *GetAccountStateWithProofByVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountStateWithProofByVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountStateWithProofByVersionRequest proto.InternalMessageInfo

func (m *GetAccountStateWithProofByVersionRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *GetAccountStateWithProofByVersionRequest) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type GetAccountStateWithProofByVersionResponse struct {
	/// The optional blob of account state blob.
	AccountStateBlob *AccountStateBlob `protobuf:"bytes,1,opt,name=account_state_blob,json=accountStateBlob,proto3" json:"account_state_blob,omitempty"`
	/// The state root hash the query is based on.
	SparseMerkleProof    *SparseMerkleProof `protobuf:"bytes,2,opt,name=sparse_merkle_proof,json=sparseMerkleProof,proto3" json:"sparse_merkle_proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetAccountStateWithProofByVersionResponse) Reset() {
	*m = GetAccountStateWithProofByVersionResponse{}
}
func (m *GetAccountStateWithProofByVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccountStateWithProofByVersionResponse) ProtoMessage()    {}
func (*GetAccountStateWithProofByVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{5}
}

func (m *GetAccountStateWithProofByVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountStateWithProofByVersionResponse.Unmarshal(m, b)
}
func (m *GetAccountStateWithProofByVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountStateWithProofByVersionResponse.Marshal(b, m, deterministic)
}
func (m *GetAccountStateWithProofByVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountStateWithProofByVersionResponse.Merge(m, src)
}
func (m *GetAccountStateWithProofByVersionResponse) XXX_Size() int {
	return xxx_messageInfo_GetAccountStateWithProofByVersionResponse.Size(m)
}
func (m *GetAccountStateWithProofByVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountStateWithProofByVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountStateWithProofByVersionResponse proto.InternalMessageInfo

func (m *GetAccountStateWithProofByVersionResponse) GetAccountStateBlob() *AccountStateBlob {
	if m != nil {
		return m.AccountStateBlob
	}
	return nil
}

func (m *GetAccountStateWithProofByVersionResponse) GetSparseMerkleProof() *SparseMerkleProof {
	if m != nil {
		return m.SparseMerkleProof
	}
	return nil
}

type GetStartupInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStartupInfoRequest) Reset()         { *m = GetStartupInfoRequest{} }
func (m *GetStartupInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetStartupInfoRequest) ProtoMessage()    {}
func (*GetStartupInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{6}
}

func (m *GetStartupInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStartupInfoRequest.Unmarshal(m, b)
}
func (m *GetStartupInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStartupInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetStartupInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStartupInfoRequest.Merge(m, src)
}
func (m *GetStartupInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetStartupInfoRequest.Size(m)
}
func (m *GetStartupInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStartupInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStartupInfoRequest proto.InternalMessageInfo

type GetStartupInfoResponse struct {
	// When this is empty, Storage needs to be bootstrapped via the bootstrap API
	Info                 *StartupInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetStartupInfoResponse) Reset()         { *m = GetStartupInfoResponse{} }
func (m *GetStartupInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetStartupInfoResponse) ProtoMessage()    {}
func (*GetStartupInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{7}
}

func (m *GetStartupInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStartupInfoResponse.Unmarshal(m, b)
}
func (m *GetStartupInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStartupInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetStartupInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStartupInfoResponse.Merge(m, src)
}
func (m *GetStartupInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetStartupInfoResponse.Size(m)
}
func (m *GetStartupInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStartupInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStartupInfoResponse proto.InternalMessageInfo

func (m *GetStartupInfoResponse) GetInfo() *StartupInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type StartupInfo struct {
	// The latest LedgerInfo. Note that at start up storage can have more
	// transactions than the latest LedgerInfo indicates due to an incomplete
	// start up sync.
	LedgerInfo *LedgerInfo `protobuf:"bytes,1,opt,name=ledger_info,json=ledgerInfo,proto3" json:"ledger_info,omitempty"`
	// The latest version. All fields below are based on this version.
	LatestVersion uint64 `protobuf:"varint,2,opt,name=latest_version,json=latestVersion,proto3" json:"latest_version,omitempty"`
	// The latest account state root hash.
	AccountStateRootHash []byte `protobuf:"bytes,3,opt,name=account_state_root_hash,json=accountStateRootHash,proto3" json:"account_state_root_hash,omitempty"`
	// From left to right, root hashes of all frozen subtrees.
	LedgerFrozenSubtreeHashes [][]byte `protobuf:"bytes,4,rep,name=ledger_frozen_subtree_hashes,json=ledgerFrozenSubtreeHashes,proto3" json:"ledger_frozen_subtree_hashes,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *StartupInfo) Reset()         { *m = StartupInfo{} }
func (m *StartupInfo) String() string { return proto.CompactTextString(m) }
func (*StartupInfo) ProtoMessage()    {}
func (*StartupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{8}
}

func (m *StartupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartupInfo.Unmarshal(m, b)
}
func (m *StartupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartupInfo.Marshal(b, m, deterministic)
}
func (m *StartupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartupInfo.Merge(m, src)
}
func (m *StartupInfo) XXX_Size() int {
	return xxx_messageInfo_StartupInfo.Size(m)
}
func (m *StartupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StartupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StartupInfo proto.InternalMessageInfo

func (m *StartupInfo) GetLedgerInfo() *LedgerInfo {
	if m != nil {
		return m.LedgerInfo
	}
	return nil
}

func (m *StartupInfo) GetLatestVersion() uint64 {
	if m != nil {
		return m.LatestVersion
	}
	return 0
}

func (m *StartupInfo) GetAccountStateRootHash() []byte {
	if m != nil {
		return m.AccountStateRootHash
	}
	return nil
}

func (m *StartupInfo) GetLedgerFrozenSubtreeHashes() [][]byte {
	if m != nil {
		return m.LedgerFrozenSubtreeHashes
	}
	return nil
}

func init() {
	proto.RegisterType((*SaveTransactionsRequest)(nil), "storage.SaveTransactionsRequest")
	proto.RegisterType((*SaveTransactionsResponse)(nil), "storage.SaveTransactionsResponse")
	proto.RegisterType((*GetTransactionsRequest)(nil), "storage.GetTransactionsRequest")
	proto.RegisterType((*GetTransactionsResponse)(nil), "storage.GetTransactionsResponse")
	proto.RegisterType((*GetAccountStateWithProofByVersionRequest)(nil), "storage.GetAccountStateWithProofByVersionRequest")
	proto.RegisterType((*GetAccountStateWithProofByVersionResponse)(nil), "storage.GetAccountStateWithProofByVersionResponse")
	proto.RegisterType((*GetStartupInfoRequest)(nil), "storage.GetStartupInfoRequest")
	proto.RegisterType((*GetStartupInfoResponse)(nil), "storage.GetStartupInfoResponse")
	proto.RegisterType((*StartupInfo)(nil), "storage.StartupInfo")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb) }

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 714 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x96, 0x9b, 0xea, 0xf4, 0x9c, 0x49, 0xd2, 0xd3, 0x6e, 0x73, 0x4e, 0x8c, 0xf9, 0x69, 0xea,
	0x0a, 0x29, 0xdc, 0x44, 0x22, 0x88, 0x6b, 0xa0, 0xa8, 0xb4, 0x48, 0x05, 0x81, 0x1d, 0xe8, 0x5d,
	0xad, 0x4d, 0x32, 0x49, 0x2c, 0x52, 0xaf, 0xd9, 0x9d, 0x94, 0xb6, 0x8f, 0xc0, 0x83, 0xf0, 0x14,
	0x3c, 0x0c, 0xaf, 0xc1, 0x1d, 0xf2, 0xee, 0xba, 0x71, 0x7e, 0x5a, 0xca, 0xe5, 0x7e, 0xf3, 0xed,
	0xb7, 0x33, 0x5f, 0xbe, 0x71, 0xa0, 0xaa, 0x48, 0x48, 0x3e, 0xc4, 0x56, 0x2a, 0x05, 0x09, 0xb6,
	0x66, 0x8f, 0x5e, 0x6d, 0x88, 0x14, 0x7d, 0x89, 0x69, 0x14, 0xa5, 0x52, 0x88, 0x81, 0x29, 0x7b,
	0x9b, 0x63, 0xec, 0x0f, 0x51, 0x46, 0x71, 0x32, 0x10, 0x39, 0x44, 0x92, 0x27, 0x8a, 0xf7, 0x28,
	0x16, 0x89, 0x85, 0x5c, 0xde, 0xeb, 0x89, 0x49, 0x42, 0x91, 0x22, 0x4e, 0x18, 0x75, 0xc7, 0xa2,
	0x6b, 0x2b, 0xe5, 0x82, 0x98, 0xff, 0xc3, 0x81, 0x7a, 0xc8, 0xcf, 0xb0, 0x33, 0x15, 0x50, 0x01,
	0x7e, 0x9e, 0xa0, 0x22, 0xf6, 0x1c, 0xd6, 0xe9, 0x3c, 0x51, 0x11, 0x89, 0xa8, 0x27, 0x4e, 0x4f,
	0x63, 0x72, 0x9d, 0x46, 0xa9, 0x59, 0x6e, 0x7b, 0x2d, 0xba, 0x48, 0x51, 0xb5, 0x0a, 0x77, 0x3a,
	0xe2, 0xa5, 0x66, 0x04, 0x95, 0xec, 0x46, 0x7e, 0x62, 0xbb, 0x50, 0x1d, 0xc4, 0x52, 0x51, 0x74,
	0x86, 0x52, 0xc5, 0x22, 0x71, 0x57, 0x1a, 0x4e, 0x73, 0x35, 0xa8, 0x68, 0xf0, 0xa3, 0xc1, 0xd8,
	0x09, 0xdc, 0x2d, 0x4c, 0x64, 0xe6, 0x55, 0xf1, 0x30, 0xe1, 0x34, 0x91, 0xa8, 0xdc, 0x52, 0xc3,
	0x69, 0x96, 0xdb, 0xdb, 0xf6, 0xcd, 0x23, 0xcd, 0x7c, 0x9d, 0x0c, 0xc4, 0x71, 0x4c, 0xa3, 0xf0,
	0x8a, 0x16, 0xb8, 0xe3, 0x6b, 0x2a, 0xbe, 0x07, 0xee, 0xe2, 0x84, 0x2a, 0x15, 0x89, 0x42, 0xff,
	0x9b, 0x03, 0xff, 0x1f, 0x20, 0x2d, 0x9b, 0x7e, 0x37, 0xfb, 0x59, 0xb8, 0x9c, 0xf6, 0xee, 0x98,
	0xde, 0x35, 0x98, 0xf7, 0x7e, 0x1f, 0xa0, 0xcb, 0xa9, 0x97, 0x35, 0x7c, 0x89, 0x76, 0xba, 0x7f,
	0x34, 0x12, 0xc6, 0x97, 0xc8, 0x1e, 0xc2, 0xba, 0x1d, 0x2d, 0x17, 0x29, 0x69, 0x4a, 0xd5, 0xa0,
	0xb9, 0xca, 0x0e, 0x54, 0x06, 0x98, 0xa9, 0xe0, 0x19, 0x26, 0xa4, 0xdc, 0xd5, 0x86, 0xd3, 0xfc,
	0x3b, 0x28, 0x6b, 0x6c, 0x5f, 0x43, 0x7e, 0x0c, 0xf5, 0x85, 0x3e, 0xcd, 0x0c, 0xec, 0x2d, 0x6c,
	0xd1, 0x79, 0x12, 0x8d, 0x63, 0x55, 0x0c, 0x8b, 0x6e, 0x77, 0xea, 0x5b, 0xe1, 0xe6, 0x51, 0xac,
	0x28, 0xb3, 0xe8, 0x5d, 0x46, 0x0b, 0x36, 0xe8, 0x7c, 0x16, 0xf1, 0x4f, 0xa0, 0x79, 0x80, 0xf4,
	0xc2, 0xc4, 0x27, 0xcc, 0xd2, 0x73, 0x55, 0xdb, 0xbb, 0xb0, 0x2d, 0xe7, 0x26, 0xb9, 0xb0, 0xc6,
	0xfb, 0x7d, 0x89, 0x4a, 0xe9, 0xf7, 0x2a, 0x41, 0x7e, 0xcc, 0x2a, 0xb3, 0x3f, 0x7a, 0x7e, 0xf4,
	0xbf, 0x3b, 0xf0, 0xe8, 0x16, 0x0f, 0xd8, 0xe9, 0xf6, 0x81, 0x2d, 0x26, 0xd9, 0x0e, 0x57, 0xb7,
	0xc3, 0x15, 0xa5, 0xf6, 0xc6, 0xa2, 0x1b, 0x6c, 0xf0, 0x39, 0x84, 0x1d, 0xc2, 0x96, 0x4a, 0xb9,
	0x54, 0x18, 0x9d, 0xa2, 0xfc, 0x34, 0x46, 0x6b, 0xd2, 0x8a, 0xd6, 0x71, 0xad, 0x4e, 0xa8, 0x19,
	0x6f, 0x34, 0xc1, 0xb8, 0xb3, 0xa9, 0xe6, 0x21, 0xbf, 0x0e, 0xff, 0x1d, 0x60, 0xa6, 0x2c, 0x69,
	0x92, 0x66, 0x71, 0xb3, 0x5e, 0xf8, 0x7b, 0x3a, 0x4a, 0x33, 0x05, 0x3b, 0x43, 0x13, 0x56, 0xb3,
	0x68, 0xdb, 0xae, 0x6b, 0xad, 0x7c, 0xdd, 0x8b, 0x5c, 0xcd, 0xc8, 0xd6, 0xb1, 0x5c, 0x40, 0x59,
	0x1b, 0xca, 0x85, 0xdd, 0xb0, 0x02, 0x9b, 0x0b, 0xbb, 0x10, 0xc0, 0x34, 0xfd, 0x3a, 0x74, 0x9c,
	0x70, 0x61, 0xeb, 0xaa, 0x06, 0xcd, 0x43, 0xf7, 0x14, 0xea, 0xb3, 0xc6, 0x4a, 0x21, 0x28, 0x1a,
	0x71, 0x35, 0xd2, 0x21, 0xad, 0x04, 0xb5, 0xa2, 0x89, 0x81, 0x10, 0x74, 0xc8, 0xd5, 0x88, 0x3d,
	0x83, 0x7b, 0xb6, 0xa3, 0x81, 0x14, 0x97, 0x98, 0x44, 0x6a, 0xd2, 0x25, 0x89, 0xa8, 0x6f, 0x62,
	0x96, 0xdd, 0x52, 0xb3, 0x12, 0xdc, 0x31, 0x9c, 0x57, 0x9a, 0x12, 0x1a, 0xc6, 0xa1, 0x26, 0xb4,
	0x7f, 0x96, 0x60, 0x2d, 0x34, 0x06, 0xb0, 0x63, 0xd8, 0x98, 0x5f, 0x4d, 0xd6, 0x98, 0xda, 0xb3,
	0xfc, 0xbb, 0xe4, 0xed, 0xdc, 0xc0, 0xb0, 0x8e, 0x47, 0x50, 0xfb, 0x90, 0xf6, 0x39, 0x61, 0x47,
	0x1c, 0xe9, 0xa9, 0x8d, 0x57, 0xcc, 0xb7, 0xd6, 0x2d, 0x2b, 0xe6, 0xf2, 0xbb, 0x37, 0x72, 0xec,
	0x03, 0x1d, 0xf8, 0x77, 0x6e, 0x1f, 0xd9, 0xf6, 0x55, 0x5b, 0xcb, 0xbf, 0x28, 0x5e, 0xe3, 0x7a,
	0x82, 0x55, 0xfd, 0xea, 0xc0, 0xce, 0x6f, 0x57, 0x83, 0x3d, 0x2e, 0xea, 0xdc, 0x6a, 0x4f, 0xbd,
	0xf6, 0x9f, 0x5c, 0xb1, 0xcd, 0xbc, 0x87, 0xf5, 0xd9, 0x3c, 0xb3, 0x07, 0x45, 0x95, 0xc5, 0x0d,
	0xf0, 0xb6, 0xaf, 0xad, 0x1b, 0xc9, 0xee, 0x5f, 0xfa, 0x4f, 0xe7, 0xc9, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1a, 0x8f, 0xbb, 0x2b, 0xf1, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageClient interface {
	// Persist transactions. Called by Execution when either syncing nodes or
	// committing blocks during normal operation.
	SaveTransactions(ctx context.Context, in *SaveTransactionsRequest, opts ...grpc.CallOption) (*SaveTransactionsResponse, error)
	// Used to get a piece of data and return the proof of it. If the client
	// knows and trusts a ledger info at version v, it should pass v in as the
	// client_known_version and we will return the latest ledger info together
	// with the proof that it derives from v.
	UpdateToLatestLedger(ctx context.Context, in *UpdateToLatestLedgerRequest, opts ...grpc.CallOption) (*UpdateToLatestLedgerResponse, error)
	// When we receive a request from a peer validator asking a list of
	// transactions for state synchronization, this API can be used to serve the
	// request. Note that the peer should specify a ledger version and all proofs
	// in the response will be relative to this given ledger version.
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
	GetAccountStateWithProofByVersion(ctx context.Context, in *GetAccountStateWithProofByVersionRequest, opts ...grpc.CallOption) (*GetAccountStateWithProofByVersionResponse, error)
	// Returns information needed for libra core to start up.
	GetStartupInfo(ctx context.Context, in *GetStartupInfoRequest, opts ...grpc.CallOption) (*GetStartupInfoResponse, error)
}

type storageClient struct {
	cc *grpc.ClientConn
}

func NewStorageClient(cc *grpc.ClientConn) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) SaveTransactions(ctx context.Context, in *SaveTransactionsRequest, opts ...grpc.CallOption) (*SaveTransactionsResponse, error) {
	out := new(SaveTransactionsResponse)
	err := c.cc.Invoke(ctx, "/storage.Storage/SaveTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) UpdateToLatestLedger(ctx context.Context, in *UpdateToLatestLedgerRequest, opts ...grpc.CallOption) (*UpdateToLatestLedgerResponse, error) {
	out := new(UpdateToLatestLedgerResponse)
	err := c.cc.Invoke(ctx, "/storage.Storage/UpdateToLatestLedger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, "/storage.Storage/GetTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetAccountStateWithProofByVersion(ctx context.Context, in *GetAccountStateWithProofByVersionRequest, opts ...grpc.CallOption) (*GetAccountStateWithProofByVersionResponse, error) {
	out := new(GetAccountStateWithProofByVersionResponse)
	err := c.cc.Invoke(ctx, "/storage.Storage/GetAccountStateWithProofByVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetStartupInfo(ctx context.Context, in *GetStartupInfoRequest, opts ...grpc.CallOption) (*GetStartupInfoResponse, error) {
	out := new(GetStartupInfoResponse)
	err := c.cc.Invoke(ctx, "/storage.Storage/GetStartupInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
type StorageServer interface {
	// Persist transactions. Called by Execution when either syncing nodes or
	// committing blocks during normal operation.
	SaveTransactions(context.Context, *SaveTransactionsRequest) (*SaveTransactionsResponse, error)
	// Used to get a piece of data and return the proof of it. If the client
	// knows and trusts a ledger info at version v, it should pass v in as the
	// client_known_version and we will return the latest ledger info together
	// with the proof that it derives from v.
	UpdateToLatestLedger(context.Context, *UpdateToLatestLedgerRequest) (*UpdateToLatestLedgerResponse, error)
	// When we receive a request from a peer validator asking a list of
	// transactions for state synchronization, this API can be used to serve the
	// request. Note that the peer should specify a ledger version and all proofs
	// in the response will be relative to this given ledger version.
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
	GetAccountStateWithProofByVersion(context.Context, *GetAccountStateWithProofByVersionRequest) (*GetAccountStateWithProofByVersionResponse, error)
	// Returns information needed for libra core to start up.
	GetStartupInfo(context.Context, *GetStartupInfoRequest) (*GetStartupInfoResponse, error)
}

// UnimplementedStorageServer can be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (*UnimplementedStorageServer) SaveTransactions(ctx context.Context, req *SaveTransactionsRequest) (*SaveTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTransactions not implemented")
}
func (*UnimplementedStorageServer) UpdateToLatestLedger(ctx context.Context, req *UpdateToLatestLedgerRequest) (*UpdateToLatestLedgerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToLatestLedger not implemented")
}
func (*UnimplementedStorageServer) GetTransactions(ctx context.Context, req *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (*UnimplementedStorageServer) GetAccountStateWithProofByVersion(ctx context.Context, req *GetAccountStateWithProofByVersionRequest) (*GetAccountStateWithProofByVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountStateWithProofByVersion not implemented")
}
func (*UnimplementedStorageServer) GetStartupInfo(ctx context.Context, req *GetStartupInfoRequest) (*GetStartupInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStartupInfo not implemented")
}

func RegisterStorageServer(s *grpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_SaveTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SaveTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/SaveTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SaveTransactions(ctx, req.(*SaveTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_UpdateToLatestLedger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateToLatestLedgerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).UpdateToLatestLedger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/UpdateToLatestLedger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).UpdateToLatestLedger(ctx, req.(*UpdateToLatestLedgerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/GetTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetAccountStateWithProofByVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountStateWithProofByVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetAccountStateWithProofByVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/GetAccountStateWithProofByVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetAccountStateWithProofByVersion(ctx, req.(*GetAccountStateWithProofByVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetStartupInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStartupInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetStartupInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/GetStartupInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetStartupInfo(ctx, req.(*GetStartupInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveTransactions",
			Handler:    _Storage_SaveTransactions_Handler,
		},
		{
			MethodName: "UpdateToLatestLedger",
			Handler:    _Storage_UpdateToLatestLedger_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _Storage_GetTransactions_Handler,
		},
		{
			MethodName: "GetAccountStateWithProofByVersion",
			Handler:    _Storage_GetAccountStateWithProofByVersion_Handler,
		},
		{
			MethodName: "GetStartupInfo",
			Handler:    _Storage_GetStartupInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}
