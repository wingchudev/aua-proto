// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/enum.proto

package user // import "gitlab.bull-b.com/aua/aua-proto/user"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 联系类型
type ContactType int32

const (
	ContactType_LoginType_Invalid ContactType = 0
	ContactType_Mobile            ContactType = 1
	ContactType_Email             ContactType = 2
	ContactType_UserName          ContactType = 3
)

var ContactType_name = map[int32]string{
	0: "LoginType_Invalid",
	1: "Mobile",
	2: "Email",
	3: "UserName",
}
var ContactType_value = map[string]int32{
	"LoginType_Invalid": 0,
	"Mobile":            1,
	"Email":             2,
	"UserName":          3,
}

func (x ContactType) String() string {
	return proto.EnumName(ContactType_name, int32(x))
}
func (ContactType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{0}
}

// 用户状态
type UserStatus int32

const (
	UserStatus_UserStatus_Invalid UserStatus = 0
	UserStatus_Normal             UserStatus = 1
	UserStatus_ForbitLogin        UserStatus = 2
)

var UserStatus_name = map[int32]string{
	0: "UserStatus_Invalid",
	1: "Normal",
	2: "ForbitLogin",
}
var UserStatus_value = map[string]int32{
	"UserStatus_Invalid": 0,
	"Normal":             1,
	"ForbitLogin":        2,
}

func (x UserStatus) String() string {
	return proto.EnumName(UserStatus_name, int32(x))
}
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{1}
}

// 通知状态
type NotificationStatus int32

const (
	NotificationStatus_NotificationStatus_Invalid NotificationStatus = 0
	NotificationStatus_SendSuccess                NotificationStatus = 1
	NotificationStatus_SendFailure                NotificationStatus = 2
	NotificationStatus_SendWaiting                NotificationStatus = 3
	NotificationStatus_IsSending                  NotificationStatus = 4
)

var NotificationStatus_name = map[int32]string{
	0: "NotificationStatus_Invalid",
	1: "SendSuccess",
	2: "SendFailure",
	3: "SendWaiting",
	4: "IsSending",
}
var NotificationStatus_value = map[string]int32{
	"NotificationStatus_Invalid": 0,
	"SendSuccess":                1,
	"SendFailure":                2,
	"SendWaiting":                3,
	"IsSending":                  4,
}

func (x NotificationStatus) String() string {
	return proto.EnumName(NotificationStatus_name, int32(x))
}
func (NotificationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{2}
}

// 通知类型
type NotificationType int32

const (
	NotificationType_NotificationType_Invalid NotificationType = 0
	NotificationType_NotificationType_SMS     NotificationType = 1
	NotificationType_NotificationType_Email   NotificationType = 2
)

var NotificationType_name = map[int32]string{
	0: "NotificationType_Invalid",
	1: "NotificationType_SMS",
	2: "NotificationType_Email",
}
var NotificationType_value = map[string]int32{
	"NotificationType_Invalid": 0,
	"NotificationType_SMS":     1,
	"NotificationType_Email":   2,
}

func (x NotificationType) String() string {
	return proto.EnumName(NotificationType_name, int32(x))
}
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{3}
}

// 在线状态
type OnlineStatus int32

const (
	OnlineStatus_OnlineStatus_Invalid OnlineStatus = 0
	OnlineStatus_Logon                OnlineStatus = 1
	OnlineStatus_Logout               OnlineStatus = 2
)

var OnlineStatus_name = map[int32]string{
	0: "OnlineStatus_Invalid",
	1: "Logon",
	2: "Logout",
}
var OnlineStatus_value = map[string]int32{
	"OnlineStatus_Invalid": 0,
	"Logon":                1,
	"Logout":               2,
}

func (x OnlineStatus) String() string {
	return proto.EnumName(OnlineStatus_name, int32(x))
}
func (OnlineStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{4}
}

type ReceiptAddressType int32

const (
	ReceiptAddressType_ReceiptAddressType_Invalid ReceiptAddressType = 0
	ReceiptAddressType_AddAddress                 ReceiptAddressType = 1
	ReceiptAddressType_ChangeAddress              ReceiptAddressType = 2
)

var ReceiptAddressType_name = map[int32]string{
	0: "ReceiptAddressType_Invalid",
	1: "AddAddress",
	2: "ChangeAddress",
}
var ReceiptAddressType_value = map[string]int32{
	"ReceiptAddressType_Invalid": 0,
	"AddAddress":                 1,
	"ChangeAddress":              2,
}

func (x ReceiptAddressType) String() string {
	return proto.EnumName(ReceiptAddressType_name, int32(x))
}
func (ReceiptAddressType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{5}
}

// 实名状态
type IdentifyStatus int32

const (
	IdentifyStatus_IdentifyStatus_Invalid IdentifyStatus = 0
	IdentifyStatus_Identifying            IdentifyStatus = 1
	IdentifyStatus_PassIdentify           IdentifyStatus = 2
	IdentifyStatus_IdentifyFail           IdentifyStatus = 3
	IdentifyStatus_NoIdentify             IdentifyStatus = 4
	IdentifyStatus_OnProcess              IdentifyStatus = 5
)

var IdentifyStatus_name = map[int32]string{
	0: "IdentifyStatus_Invalid",
	1: "Identifying",
	2: "PassIdentify",
	3: "IdentifyFail",
	4: "NoIdentify",
	5: "OnProcess",
}
var IdentifyStatus_value = map[string]int32{
	"IdentifyStatus_Invalid": 0,
	"Identifying":            1,
	"PassIdentify":           2,
	"IdentifyFail":           3,
	"NoIdentify":             4,
	"OnProcess":              5,
}

func (x IdentifyStatus) String() string {
	return proto.EnumName(IdentifyStatus_name, int32(x))
}
func (IdentifyStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{6}
}

// 实名类型
type IdentifyType int32

const (
	IdentifyType_IdentifyType_Invalid IdentifyType = 0
	IdentifyType_IDCard               IdentifyType = 1
	IdentifyType_Company              IdentifyType = 2
	IdentifyType_Passport             IdentifyType = 3
	IdentifyType_Organization         IdentifyType = 4
)

var IdentifyType_name = map[int32]string{
	0: "IdentifyType_Invalid",
	1: "IDCard",
	2: "Company",
	3: "Passport",
	4: "Organization",
}
var IdentifyType_value = map[string]int32{
	"IdentifyType_Invalid": 0,
	"IDCard":               1,
	"Company":              2,
	"Passport":             3,
	"Organization":         4,
}

func (x IdentifyType) String() string {
	return proto.EnumName(IdentifyType_name, int32(x))
}
func (IdentifyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{7}
}

type UserBankStatus int32

const (
	UserBankStatus_Bank_Invalid  UserBankStatus = 0
	UserBankStatus_Bank_Pending  UserBankStatus = 1
	UserBankStatus_Bank_Active   UserBankStatus = 2
	UserBankStatus_Bank_Inactive UserBankStatus = 3
)

var UserBankStatus_name = map[int32]string{
	0: "Bank_Invalid",
	1: "Bank_Pending",
	2: "Bank_Active",
	3: "Bank_Inactive",
}
var UserBankStatus_value = map[string]int32{
	"Bank_Invalid":  0,
	"Bank_Pending":  1,
	"Bank_Active":   2,
	"Bank_Inactive": 3,
}

func (x UserBankStatus) String() string {
	return proto.EnumName(UserBankStatus_name, int32(x))
}
func (UserBankStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{8}
}

// 验证码类型
type VerifyCodeType int32

const (
	VerifyCodeType_VerifyCodeType_Invalid VerifyCodeType = 0
	VerifyCodeType_Regist                 VerifyCodeType = 1
	VerifyCodeType_Login                  VerifyCodeType = 2
	VerifyCodeType_ResetPassword          VerifyCodeType = 3
	VerifyCodeType_ChangePassword         VerifyCodeType = 4
	VerifyCodeType_UnbindMobile           VerifyCodeType = 5
	VerifyCodeType_BindMobile             VerifyCodeType = 6
	VerifyCodeType_UnbindEmail            VerifyCodeType = 7
	VerifyCodeType_BindEmail              VerifyCodeType = 8
	VerifyCodeType_ChangeTradePassword    VerifyCodeType = 9
	VerifyCodeType_ChangeNickName         VerifyCodeType = 10
	VerifyCodeType_BindPayment            VerifyCodeType = 11
	VerifyCodeType_DeletePayment          VerifyCodeType = 12
	VerifyCodeType_BindEOSAccount         VerifyCodeType = 13
	VerifyCodeType_CRMLogin               VerifyCodeType = 14
)

var VerifyCodeType_name = map[int32]string{
	0:  "VerifyCodeType_Invalid",
	1:  "Regist",
	2:  "Login",
	3:  "ResetPassword",
	4:  "ChangePassword",
	5:  "UnbindMobile",
	6:  "BindMobile",
	7:  "UnbindEmail",
	8:  "BindEmail",
	9:  "ChangeTradePassword",
	10: "ChangeNickName",
	11: "BindPayment",
	12: "DeletePayment",
	13: "BindEOSAccount",
	14: "CRMLogin",
}
var VerifyCodeType_value = map[string]int32{
	"VerifyCodeType_Invalid": 0,
	"Regist":                 1,
	"Login":                  2,
	"ResetPassword":          3,
	"ChangePassword":         4,
	"UnbindMobile":           5,
	"BindMobile":             6,
	"UnbindEmail":            7,
	"BindEmail":              8,
	"ChangeTradePassword":    9,
	"ChangeNickName":         10,
	"BindPayment":            11,
	"DeletePayment":          12,
	"BindEOSAccount":         13,
	"CRMLogin":               14,
}

func (x VerifyCodeType) String() string {
	return proto.EnumName(VerifyCodeType_name, int32(x))
}
func (VerifyCodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{9}
}

// 验证码状态
type VerifyCodeStatus int32

const (
	VerifyCodeStatus_VerifyCode_Invalid VerifyCodeStatus = 0
	VerifyCodeStatus_VerifyCodeNotUse   VerifyCodeStatus = 1
	VerifyCodeStatus_VerifyCodeUsed     VerifyCodeStatus = 2
)

var VerifyCodeStatus_name = map[int32]string{
	0: "VerifyCode_Invalid",
	1: "VerifyCodeNotUse",
	2: "VerifyCodeUsed",
}
var VerifyCodeStatus_value = map[string]int32{
	"VerifyCode_Invalid": 0,
	"VerifyCodeNotUse":   1,
	"VerifyCodeUsed":     2,
}

func (x VerifyCodeStatus) String() string {
	return proto.EnumName(VerifyCodeStatus_name, int32(x))
}
func (VerifyCodeStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{10}
}

// 支付方式类型
type PaymentType int32

const (
	PaymentType_PaymentType_Invalid PaymentType = 0
	PaymentType_Bankcard            PaymentType = 1
	PaymentType_AliPay              PaymentType = 2
	PaymentType_Wechat              PaymentType = 3
)

var PaymentType_name = map[int32]string{
	0: "PaymentType_Invalid",
	1: "Bankcard",
	2: "AliPay",
	3: "Wechat",
}
var PaymentType_value = map[string]int32{
	"PaymentType_Invalid": 0,
	"Bankcard":            1,
	"AliPay":              2,
	"Wechat":              3,
}

func (x PaymentType) String() string {
	return proto.EnumName(PaymentType_name, int32(x))
}
func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{11}
}

// 通用状态
type NormalStatus int32

const (
	NormalStatus_NormalStatus_Invalid NormalStatus = 0
	NormalStatus_Enable               NormalStatus = 1
	NormalStatus_Disable              NormalStatus = 2
)

var NormalStatus_name = map[int32]string{
	0: "NormalStatus_Invalid",
	1: "Enable",
	2: "Disable",
}
var NormalStatus_value = map[string]int32{
	"NormalStatus_Invalid": 0,
	"Enable":               1,
	"Disable":              2,
}

func (x NormalStatus) String() string {
	return proto.EnumName(NormalStatus_name, int32(x))
}
func (NormalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{12}
}

// 通知type
type NotificationByType int32

const (
	NotificationByType_NotificationByType_Invalid       NotificationByType = 0
	NotificationByType_Receipt                          NotificationByType = 1
	NotificationByType_IssueApply                       NotificationByType = 2
	NotificationByType_BindEOS                          NotificationByType = 3
	NotificationByType_IssuePass                        NotificationByType = 4
	NotificationByType_IssueReject                      NotificationByType = 5
	NotificationByType_IssueSuccess                     NotificationByType = 6
	NotificationByType_ContractApply                    NotificationByType = 7
	NotificationByType_Extract                          NotificationByType = 8
	NotificationByType_RegistSMSVerifyCode              NotificationByType = 9
	NotificationByType_RegistEmailVerifyCode            NotificationByType = 10
	NotificationByType_RegistSMSSuccess                 NotificationByType = 11
	NotificationByType_RegistEmailSuccess               NotificationByType = 12
	NotificationByType_BindMobileSMSVerifyCode          NotificationByType = 13
	NotificationByType_BindEmailEmailVerifyCode         NotificationByType = 14
	NotificationByType_ModifyPassWordSMSVerifyCode      NotificationByType = 15
	NotificationByType_ModifyPassWordEmailVerifyCode    NotificationByType = 16
	NotificationByType_FindTradePasswordSMSVerifyCode   NotificationByType = 17
	NotificationByType_FindTradePasswordEmailVerifyCode NotificationByType = 18
	NotificationByType_SetTradePasswordSMS              NotificationByType = 19
	NotificationByType_BindEOSEmail                     NotificationByType = 20
	NotificationByType_IssuePassEmail                   NotificationByType = 21
	NotificationByType_IssueRejectEmail                 NotificationByType = 22
	NotificationByType_ExtractEmail                     NotificationByType = 23
	NotificationByType_IssueNotification                NotificationByType = 24
	NotificationByType_OtherTemplete                    NotificationByType = 25
	NotificationByType_IdentifyApplySMS                 NotificationByType = 26
	NotificationByType_IdentifyApplyEmail               NotificationByType = 27
	NotificationByType_IdentifyPassSMS                  NotificationByType = 28
	NotificationByType_IdentifyPassEmail                NotificationByType = 29
	NotificationByType_IssueTokenSMS                    NotificationByType = 30
	NotificationByType_IssueTokenEmail                  NotificationByType = 31
	NotificationByType_ReceiveTokenSMS                  NotificationByType = 32
	NotificationByType_ReceiveTokenEmail                NotificationByType = 33
	NotificationByType_WithdrawToBankSMS                NotificationByType = 34
	NotificationByType_WithdrawToBankEmail              NotificationByType = 35
	NotificationByType_ReceiveTokenSettlementEmail      NotificationByType = 36
)

var NotificationByType_name = map[int32]string{
	0:  "NotificationByType_Invalid",
	1:  "Receipt",
	2:  "IssueApply",
	3:  "BindEOS",
	4:  "IssuePass",
	5:  "IssueReject",
	6:  "IssueSuccess",
	7:  "ContractApply",
	8:  "Extract",
	9:  "RegistSMSVerifyCode",
	10: "RegistEmailVerifyCode",
	11: "RegistSMSSuccess",
	12: "RegistEmailSuccess",
	13: "BindMobileSMSVerifyCode",
	14: "BindEmailEmailVerifyCode",
	15: "ModifyPassWordSMSVerifyCode",
	16: "ModifyPassWordEmailVerifyCode",
	17: "FindTradePasswordSMSVerifyCode",
	18: "FindTradePasswordEmailVerifyCode",
	19: "SetTradePasswordSMS",
	20: "BindEOSEmail",
	21: "IssuePassEmail",
	22: "IssueRejectEmail",
	23: "ExtractEmail",
	24: "IssueNotification",
	25: "OtherTemplete",
	26: "IdentifyApplySMS",
	27: "IdentifyApplyEmail",
	28: "IdentifyPassSMS",
	29: "IdentifyPassEmail",
	30: "IssueTokenSMS",
	31: "IssueTokenEmail",
	32: "ReceiveTokenSMS",
	33: "ReceiveTokenEmail",
	34: "WithdrawToBankSMS",
	35: "WithdrawToBankEmail",
	36: "ReceiveTokenSettlementEmail",
}
var NotificationByType_value = map[string]int32{
	"NotificationByType_Invalid":       0,
	"Receipt":                          1,
	"IssueApply":                       2,
	"BindEOS":                          3,
	"IssuePass":                        4,
	"IssueReject":                      5,
	"IssueSuccess":                     6,
	"ContractApply":                    7,
	"Extract":                          8,
	"RegistSMSVerifyCode":              9,
	"RegistEmailVerifyCode":            10,
	"RegistSMSSuccess":                 11,
	"RegistEmailSuccess":               12,
	"BindMobileSMSVerifyCode":          13,
	"BindEmailEmailVerifyCode":         14,
	"ModifyPassWordSMSVerifyCode":      15,
	"ModifyPassWordEmailVerifyCode":    16,
	"FindTradePasswordSMSVerifyCode":   17,
	"FindTradePasswordEmailVerifyCode": 18,
	"SetTradePasswordSMS":              19,
	"BindEOSEmail":                     20,
	"IssuePassEmail":                   21,
	"IssueRejectEmail":                 22,
	"ExtractEmail":                     23,
	"IssueNotification":                24,
	"OtherTemplete":                    25,
	"IdentifyApplySMS":                 26,
	"IdentifyApplyEmail":               27,
	"IdentifyPassSMS":                  28,
	"IdentifyPassEmail":                29,
	"IssueTokenSMS":                    30,
	"IssueTokenEmail":                  31,
	"ReceiveTokenSMS":                  32,
	"ReceiveTokenEmail":                33,
	"WithdrawToBankSMS":                34,
	"WithdrawToBankEmail":              35,
	"ReceiveTokenSettlementEmail":      36,
}

func (x NotificationByType) String() string {
	return proto.EnumName(NotificationByType_name, int32(x))
}
func (NotificationByType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_99e672d53393dce3, []int{13}
}

func init() {
	proto.RegisterEnum("user.ContactType", ContactType_name, ContactType_value)
	proto.RegisterEnum("user.UserStatus", UserStatus_name, UserStatus_value)
	proto.RegisterEnum("user.NotificationStatus", NotificationStatus_name, NotificationStatus_value)
	proto.RegisterEnum("user.NotificationType", NotificationType_name, NotificationType_value)
	proto.RegisterEnum("user.OnlineStatus", OnlineStatus_name, OnlineStatus_value)
	proto.RegisterEnum("user.ReceiptAddressType", ReceiptAddressType_name, ReceiptAddressType_value)
	proto.RegisterEnum("user.IdentifyStatus", IdentifyStatus_name, IdentifyStatus_value)
	proto.RegisterEnum("user.IdentifyType", IdentifyType_name, IdentifyType_value)
	proto.RegisterEnum("user.UserBankStatus", UserBankStatus_name, UserBankStatus_value)
	proto.RegisterEnum("user.VerifyCodeType", VerifyCodeType_name, VerifyCodeType_value)
	proto.RegisterEnum("user.VerifyCodeStatus", VerifyCodeStatus_name, VerifyCodeStatus_value)
	proto.RegisterEnum("user.PaymentType", PaymentType_name, PaymentType_value)
	proto.RegisterEnum("user.NormalStatus", NormalStatus_name, NormalStatus_value)
	proto.RegisterEnum("user.NotificationByType", NotificationByType_name, NotificationByType_value)
}

func init() { proto.RegisterFile("user/enum.proto", fileDescriptor_enum_99e672d53393dce3) }

var fileDescriptor_enum_99e672d53393dce3 = []byte{
	// 1042 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x56, 0xdd, 0x4e, 0xeb, 0x46,
	0x10, 0x6e, 0x7e, 0x48, 0x0e, 0xe3, 0x10, 0x96, 0xe5, 0xf7, 0x84, 0x73, 0x38, 0x3f, 0x45, 0xbd,
	0xb0, 0x74, 0xe0, 0xa2, 0x0f, 0x50, 0x85, 0x00, 0x52, 0x24, 0x12, 0x10, 0x0e, 0xa5, 0xea, 0x4d,
	0xb5, 0xb1, 0x97, 0xb0, 0xc5, 0xd9, 0x8d, 0xec, 0x35, 0xa7, 0xe9, 0x55, 0xdf, 0xa4, 0x0f, 0xd8,
	0x97, 0xa8, 0x66, 0xd7, 0x8e, 0xd7, 0xe4, 0x02, 0x29, 0xf3, 0xed, 0xec, 0x37, 0xb3, 0xf3, 0xcd,
	0x0c, 0x86, 0xed, 0x2c, 0xe5, 0xc9, 0x39, 0x97, 0xd9, 0xfc, 0x6c, 0x91, 0x28, 0xad, 0x68, 0x13,
	0x01, 0x7f, 0x08, 0xde, 0x40, 0x49, 0xcd, 0x42, 0x3d, 0x59, 0x2e, 0x38, 0xdd, 0x87, 0x9d, 0x1b,
	0x35, 0x13, 0x12, 0x8d, 0x3f, 0x86, 0xf2, 0x95, 0xc5, 0x22, 0x22, 0x3f, 0x50, 0x80, 0xd6, 0x48,
	0x4d, 0x45, 0xcc, 0x49, 0x8d, 0x6e, 0xc2, 0xc6, 0xd5, 0x9c, 0x89, 0x98, 0xd4, 0x69, 0x07, 0xde,
	0x3d, 0xa4, 0x3c, 0x19, 0xb3, 0x39, 0x27, 0x0d, 0xbf, 0x0f, 0x80, 0x56, 0xa0, 0x99, 0xce, 0x52,
	0x7a, 0x00, 0xb4, 0xb4, 0xaa, 0x54, 0x63, 0x95, 0xcc, 0x59, 0x4c, 0x6a, 0x74, 0x1b, 0xbc, 0x6b,
	0x95, 0x4c, 0x85, 0x36, 0x31, 0x49, 0xdd, 0x7f, 0x05, 0x3a, 0x56, 0x5a, 0x3c, 0x89, 0x90, 0x69,
	0xa1, 0x64, 0x4e, 0x75, 0x02, 0xbd, 0x75, 0xd4, 0xa1, 0xdc, 0x06, 0x2f, 0xe0, 0x32, 0x0a, 0xb2,
	0x30, 0xe4, 0x69, 0x6a, 0x79, 0x11, 0xb8, 0x66, 0x22, 0xce, 0x12, 0x4e, 0xea, 0x05, 0xf0, 0xc8,
	0x84, 0x16, 0x72, 0x46, 0x1a, 0x74, 0x0b, 0x36, 0x87, 0x29, 0x42, 0x68, 0x36, 0xfd, 0x27, 0x20,
	0x6e, 0x04, 0x53, 0x8a, 0x0f, 0x70, 0xf4, 0x16, 0x73, 0x62, 0x1e, 0xc1, 0xde, 0xda, 0x69, 0x30,
	0x0a, 0x48, 0x8d, 0xf6, 0xe0, 0x60, 0xed, 0x24, 0x2f, 0x98, 0xff, 0x0b, 0x74, 0x6e, 0x65, 0x2c,
	0x24, 0xcf, 0x5f, 0x76, 0x04, 0x7b, 0xae, 0xed, 0xf0, 0x6f, 0xc2, 0xc6, 0x8d, 0x9a, 0x29, 0x49,
	0x6a, 0x58, 0x31, 0xfc, 0x99, 0x69, 0x52, 0xf7, 0x1f, 0x81, 0xde, 0xf3, 0x90, 0x8b, 0x85, 0xee,
	0x47, 0x51, 0xc2, 0xd3, 0xd4, 0xa4, 0x7a, 0x02, 0xbd, 0x75, 0xd4, 0x21, 0xeb, 0x02, 0xf4, 0xa3,
	0x28, 0x3f, 0x23, 0x35, 0xba, 0x03, 0x5b, 0x83, 0x67, 0x26, 0x67, 0xbc, 0x80, 0xea, 0xfe, 0x3f,
	0x35, 0xe8, 0x0e, 0x23, 0x2e, 0xb5, 0x78, 0x5a, 0xe6, 0xc9, 0xf5, 0xe0, 0xa0, 0x8a, 0x54, 0x4b,
	0x5e, 0x9c, 0x61, 0x05, 0x6b, 0x94, 0x40, 0xe7, 0x8e, 0xa5, 0x69, 0x01, 0x92, 0x3a, 0x22, 0x85,
	0x85, 0x42, 0x90, 0x06, 0xa6, 0x31, 0x56, 0x2b, 0x8f, 0x26, 0x8a, 0x70, 0x2b, 0xef, 0x12, 0x65,
	0x54, 0xdb, 0xf0, 0x59, 0x79, 0xc1, 0xbc, 0xea, 0x08, 0xf6, 0x5c, 0xbb, 0xda, 0x43, 0xc3, 0xcb,
	0x01, 0x4b, 0x22, 0x52, 0xa3, 0x1e, 0xb4, 0x07, 0x6a, 0xbe, 0x60, 0x72, 0x69, 0x1b, 0x12, 0xb3,
	0x58, 0xa8, 0x44, 0x93, 0x06, 0x66, 0x70, 0x9b, 0xcc, 0x98, 0x14, 0x7f, 0x1b, 0x25, 0x48, 0xd3,
	0xff, 0x0d, 0xba, 0xd8, 0x94, 0x17, 0x4c, 0xbe, 0xe4, 0x8f, 0x24, 0xd0, 0x41, 0xcb, 0x21, 0x2f,
	0x90, 0xbb, 0xbc, 0x3b, 0x4c, 0x3b, 0x19, 0xa4, 0x1f, 0x6a, 0xf1, 0x8a, 0xed, 0xb4, 0x03, 0x5b,
	0xf9, 0x25, 0x66, 0xa1, 0x86, 0xff, 0x6f, 0x1d, 0xba, 0xbf, 0xf2, 0x44, 0x3c, 0x2d, 0x07, 0x2a,
	0xe2, 0x26, 0xff, 0x1e, 0x1c, 0x54, 0x91, 0xea, 0x0b, 0xee, 0xf9, 0x4c, 0xa4, 0xda, 0x0e, 0x54,
	0xde, 0xff, 0x48, 0x7c, 0xcf, 0x53, 0xae, 0xf1, 0x11, 0xdf, 0x55, 0x12, 0x91, 0x06, 0xa5, 0xd0,
	0xb5, 0x5a, 0xad, 0xb0, 0x26, 0xa6, 0xf8, 0x20, 0xa7, 0x42, 0x46, 0xf9, 0x50, 0x6e, 0x60, 0x69,
	0x2f, 0x4a, 0xbb, 0x85, 0x29, 0x5b, 0x0f, 0xdb, 0x79, 0x6d, 0xac, 0xf5, 0xc5, 0xca, 0x7c, 0x47,
	0x0f, 0x61, 0xd7, 0xb2, 0x4e, 0x12, 0x16, 0x95, 0xd4, 0x9b, 0x65, 0xb8, 0xb1, 0x08, 0x5f, 0xcc,
	0x60, 0x83, 0x79, 0xbf, 0x90, 0xd1, 0x1d, 0x5b, 0xce, 0xb9, 0xd4, 0xc4, 0xc3, 0x34, 0x2f, 0x79,
	0xcc, 0x35, 0x2f, 0xa0, 0x0e, 0xde, 0x33, 0xfc, 0xb7, 0x41, 0x3f, 0x0c, 0x55, 0x26, 0x35, 0xd9,
	0x42, 0x35, 0x06, 0xf7, 0x23, 0xfb, 0xb6, 0xae, 0x3f, 0x01, 0x52, 0x96, 0xa3, 0x5c, 0x12, 0x25,
	0xe6, 0x94, 0x67, 0xcf, 0xf5, 0x1d, 0x2b, 0xfd, 0x90, 0xe2, 0xe6, 0xa1, 0x6e, 0x89, 0x1f, 0x52,
	0x1e, 0x91, 0xba, 0x7f, 0x03, 0x5e, 0x9e, 0x84, 0xa9, 0xf9, 0x21, 0xec, 0x3a, 0xa6, 0xc3, 0xd8,
	0x81, 0x77, 0x28, 0x59, 0x68, 0x9b, 0x06, 0xa0, 0xd5, 0x8f, 0xc5, 0x1d, 0xc3, 0x9e, 0x01, 0x68,
	0x3d, 0xf2, 0xf0, 0x99, 0x69, 0xb3, 0xc2, 0x3a, 0x76, 0x39, 0x95, 0xf3, 0xe9, 0xda, 0x55, 0x01,
	0xaf, 0x24, 0x9b, 0x9a, 0x8d, 0xe8, 0x41, 0xfb, 0x52, 0xa4, 0xc6, 0xa8, 0xfb, 0xff, 0xb5, 0xaa,
	0x3b, 0xec, 0x62, 0x59, 0x8c, 0xe8, 0x3a, 0xea, 0xf0, 0x79, 0xd0, 0xce, 0x47, 0x98, 0xd4, 0x50,
	0xcd, 0x61, 0x9a, 0x66, 0xbc, 0xbf, 0x58, 0xc4, 0x98, 0xa2, 0x07, 0xed, 0xbc, 0xb8, 0xc5, 0xea,
	0x4a, 0x33, 0x23, 0x1a, 0x69, 0x9a, 0x49, 0x44, 0xf3, 0x9e, 0xff, 0xc9, 0x43, 0x4d, 0x36, 0xcc,
	0xdc, 0x21, 0x50, 0xac, 0xc3, 0x96, 0x19, 0x77, 0x25, 0x75, 0xc2, 0x42, 0x6d, 0x19, 0xdb, 0xc8,
	0x78, 0xf5, 0x97, 0x41, 0x6c, 0x33, 0xd8, 0x66, 0x0c, 0x46, 0x41, 0x59, 0x60, 0xb2, 0x49, 0xdf,
	0xc3, 0xbe, 0x3d, 0x30, 0x6d, 0xe3, 0x1c, 0x01, 0x2a, 0xb4, 0xba, 0x53, 0x44, 0xf2, 0x50, 0x4f,
	0xe7, 0x42, 0x81, 0x77, 0xe8, 0x31, 0x1c, 0x96, 0xed, 0x59, 0x8d, 0xb2, 0x85, 0x8b, 0x76, 0xd5,
	0x9a, 0x6f, 0x03, 0x75, 0xe9, 0x27, 0x38, 0x1e, 0xa9, 0x48, 0x3c, 0x2d, 0xf1, 0xbd, 0x8f, 0x2a,
	0x89, 0xaa, 0xd7, 0xb7, 0xe9, 0x17, 0xf8, 0x58, 0x75, 0x78, 0xcb, 0x41, 0xe8, 0x57, 0x38, 0xb9,
	0x16, 0x32, 0xaa, 0xf4, 0x7a, 0x95, 0x66, 0x87, 0x9e, 0xc2, 0xe7, 0x35, 0x9f, 0xb7, 0x4c, 0x14,
	0x4b, 0x15, 0x70, 0xfd, 0x96, 0x88, 0xec, 0x9a, 0xad, 0x61, 0x25, 0xb2, 0x23, 0xb6, 0x87, 0xdd,
	0xba, 0xd2, 0xc9, 0x62, 0xfb, 0x58, 0x35, 0x47, 0x2c, 0x8b, 0x1e, 0xe0, 0xdd, 0x5c, 0x0c, 0x8b,
	0x1c, 0xe2, 0xbf, 0x61, 0xe3, 0xe7, 0xb6, 0x0c, 0x39, 0x42, 0x21, 0x6f, 0xf5, 0x33, 0x4f, 0x26,
	0x7c, 0xbe, 0xc0, 0xf1, 0x23, 0xef, 0x0d, 0x63, 0xbe, 0x24, 0x8d, 0xb6, 0x98, 0x4d, 0x0f, 0x75,
	0xa8, 0xa0, 0x96, 0xf7, 0x98, 0xee, 0xc2, 0x76, 0x81, 0x63, 0x5a, 0xe8, 0xfc, 0xc1, 0x04, 0x73,
	0x40, 0xeb, 0xfb, 0x11, 0x83, 0x99, 0x1c, 0x26, 0xea, 0x85, 0x4b, 0xf4, 0x3c, 0x31, 0xd7, 0x57,
	0x90, 0xf5, 0xfb, 0x84, 0xa0, 0xe9, 0xdc, 0xd7, 0xd2, 0xf3, 0x33, 0x72, 0xba, 0xa0, 0xf5, 0xfd,
	0x82, 0xf0, 0xa3, 0xd0, 0xcf, 0x51, 0xc2, 0xbe, 0x4f, 0x94, 0xd9, 0xc2, 0xa3, 0x80, 0x7c, 0xc5,
	0xaa, 0x56, 0x61, 0xeb, 0xff, 0x23, 0x8a, 0x5f, 0xe1, 0xe6, 0x5a, 0xc7, 0x1c, 0xa7, 0xdb, 0x3a,
	0x9c, 0x5e, 0xfc, 0xf4, 0xfb, 0xe9, 0x4c, 0xe8, 0x98, 0x4d, 0xcf, 0xa6, 0x59, 0x1c, 0x7f, 0x9b,
	0x9e, 0x85, 0x6a, 0x7e, 0xce, 0x32, 0x86, 0x7f, 0xdf, 0xcc, 0x77, 0xce, 0x39, 0x7e, 0xe6, 0x4c,
	0x5b, 0xe6, 0xf7, 0xcf, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xaf, 0xa1, 0xe9, 0x06, 0x09,
	0x00, 0x00,
}
