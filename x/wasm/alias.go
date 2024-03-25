// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/andromedaprotocol/wasmd/x/wasm/types
// ALIASGEN: github.com/andromedaprotocol/wasmd/x/wasm/keeper
package wasm

import (
	"github.com/andromedaprotocol/wasmd/x/wasm/keeper"
	"github.com/andromedaprotocol/wasmd/x/wasm/types"
)

const (
	// Deprecated: Do not use.
	ModuleName = types.ModuleName
	// Deprecated: Do not use.
	StoreKey = types.StoreKey
	// Deprecated: Do not use.
	TStoreKey = types.TStoreKey
	// Deprecated: Do not use.
	QuerierRoute = types.QuerierRoute
	// Deprecated: Do not use.
	RouterKey = types.RouterKey
	// Deprecated: Do not use.
	WasmModuleEventType = types.WasmModuleEventType
	// Deprecated: Do not use.
	AttributeKeyContractAddr = types.AttributeKeyContractAddr
	// Deprecated: Do not use.
	ProposalTypeStoreCode = types.ProposalTypeStoreCode
	// Deprecated: Do not use.
	ProposalTypeInstantiateContract = types.ProposalTypeInstantiateContract
	// Deprecated: Do not use.
	ProposalTypeMigrateContract = types.ProposalTypeMigrateContract
	// Deprecated: Do not use.
	ProposalTypeUpdateAdmin = types.ProposalTypeUpdateAdmin
	// Deprecated: Do not use.
	ProposalTypeClearAdmin = types.ProposalTypeClearAdmin
)

var (
	// functions aliases
	// Deprecated: Do not use.
	RegisterCodec = types.RegisterLegacyAminoCodec
	// Deprecated: Do not use.
	RegisterInterfaces = types.RegisterInterfaces
	// Deprecated: Do not use.
	ValidateGenesis = types.ValidateGenesis
	// Deprecated: Do not use.
	ConvertToProposals = types.ConvertToProposals
	// Deprecated: Do not use.
	GetCodeKey = types.GetCodeKey
	// Deprecated: Do not use.
	GetContractAddressKey = types.GetContractAddressKey
	// Deprecated: Do not use.
	GetContractStorePrefixKey = types.GetContractStorePrefix
	// Deprecated: Do not use.
	NewCodeInfo = types.NewCodeInfo
	// Deprecated: Do not use.
	NewAbsoluteTxPosition = types.NewAbsoluteTxPosition
	// Deprecated: Do not use.
	NewContractInfo = types.NewContractInfo
	// Deprecated: Do not use.
	NewEnv = types.NewEnv
	// Deprecated: Do not use.
	NewWasmCoins = types.NewWasmCoins
	// Deprecated: Do not use.
	DefaultWasmConfig = types.DefaultWasmConfig
	// Deprecated: Do not use.
	DefaultParams = types.DefaultParams
	// Deprecated: Do not use.
	InitGenesis = keeper.InitGenesis
	// Deprecated: Do not use.
	ExportGenesis = keeper.ExportGenesis
	// Deprecated: Do not use.
	NewMessageHandler = keeper.NewDefaultMessageHandler
	// Deprecated: Do not use.
	DefaultEncoders = keeper.DefaultEncoders
	// Deprecated: Do not use.
	EncodeBankMsg = keeper.EncodeBankMsg
	// Deprecated: Do not use.
	NoCustomMsg = keeper.NoCustomMsg
	// Deprecated: Do not use.
	EncodeStakingMsg = keeper.EncodeStakingMsg
	// Deprecated: Do not use.
	EncodeWasmMsg = keeper.EncodeWasmMsg
	// Deprecated: Do not use.
	NewKeeper = keeper.NewKeeper
	// Deprecated: Do not use.
	DefaultQueryPlugins = keeper.DefaultQueryPlugins
	// Deprecated: Do not use.
	BankQuerier = keeper.BankQuerier
	// Deprecated: Do not use.
	NoCustomQuerier = keeper.NoCustomQuerier
	// Deprecated: Do not use.
	StakingQuerier = keeper.StakingQuerier
	// Deprecated: Do not use.
	WasmQuerier = keeper.WasmQuerier
	// Deprecated: Do not use.
	CreateTestInput = keeper.CreateTestInput
	// Deprecated: Do not use.
	TestHandler = keeper.TestHandler
	// Deprecated: Do not use.
	NewWasmProposalHandler = keeper.NewWasmProposalHandler //nolint:staticcheck
	// Deprecated: Do not use.
	NewQuerier = keeper.Querier
	// Deprecated: Do not use.
	ContractFromPortID = keeper.ContractFromPortID
	// Deprecated: Do not use.
	WithWasmEngine = keeper.WithWasmEngine
	// Deprecated: Do not use.
	NewCountTXDecorator = keeper.NewCountTXDecorator

	// variable aliases
	// Deprecated: Do not use.
	ModuleCdc = types.ModuleCdc
	// Deprecated: Do not use.
	DefaultCodespace = types.DefaultCodespace
	// Deprecated: Do not use.
	ErrCreateFailed = types.ErrCreateFailed
	// Deprecated: Do not use.
	ErrAccountExists = types.ErrAccountExists
	// Deprecated: Do not use.
	ErrInstantiateFailed = types.ErrInstantiateFailed
	// Deprecated: Do not use.
	ErrExecuteFailed = types.ErrExecuteFailed
	// Deprecated: Do not use.
	ErrGasLimit = types.ErrGasLimit
	// Deprecated: Do not use.
	ErrInvalidGenesis = types.ErrInvalidGenesis
	// Deprecated: Do not use.
	ErrNotFound = types.ErrNotFound
	// Deprecated: Do not use.
	ErrQueryFailed = types.ErrQueryFailed
	// Deprecated: Do not use.
	ErrInvalidMsg = types.ErrInvalidMsg
	// Deprecated: Do not use.
	KeyLastCodeID = types.KeyLastCodeID
	// Deprecated: Do not use.
	KeyLastInstanceID = types.KeyLastInstanceID
	// Deprecated: Do not use.
	CodeKeyPrefix = types.CodeKeyPrefix
	// Deprecated: Do not use.
	ContractKeyPrefix = types.ContractKeyPrefix
	// Deprecated: Do not use.
	ContractStorePrefix = types.ContractStorePrefix
	// Deprecated: Do not use.
	EnableAllProposals = types.EnableAllProposals
	// Deprecated: Do not use.
	DisableAllProposals = types.DisableAllProposals
)

type (
	// Deprecated: Do not use.
	ProposalType = types.ProposalType
	// Deprecated: Do not use.
	GenesisState = types.GenesisState
	// Deprecated: Do not use.
	Code = types.Code
	// Deprecated: Do not use.
	Contract = types.Contract
	// Deprecated: Do not use.
	MsgStoreCode = types.MsgStoreCode
	// Deprecated: Do not use.
	MsgStoreCodeResponse = types.MsgStoreCodeResponse
	// Deprecated: Do not use.
	MsgInstantiateContract = types.MsgInstantiateContract
	// Deprecated: Do not use.
	MsgInstantiateContract2 = types.MsgInstantiateContract2
	// Deprecated: Do not use.
	MsgInstantiateContractResponse = types.MsgInstantiateContractResponse
	// Deprecated: Do not use.
	MsgExecuteContract = types.MsgExecuteContract
	// Deprecated: Do not use.
	MsgExecuteContractResponse = types.MsgExecuteContractResponse
	// Deprecated: Do not use.
	MsgMigrateContract = types.MsgMigrateContract
	// Deprecated: Do not use.
	MsgMigrateContractResponse = types.MsgMigrateContractResponse
	// Deprecated: Do not use.
	MsgUpdateAdmin = types.MsgUpdateAdmin
	// Deprecated: Do not use.
	MsgUpdateAdminResponse = types.MsgUpdateAdminResponse
	// Deprecated: Do not use.
	MsgClearAdmin = types.MsgClearAdmin
	// Deprecated: Do not use.
	MsgWasmIBCCall = types.MsgIBCSend
	// Deprecated: Do not use.
	MsgClearAdminResponse = types.MsgClearAdminResponse
	// Deprecated: Do not use.
	MsgServer = types.MsgServer
	// Deprecated: Do not use.
	Model = types.Model
	// Deprecated: Do not use.
	CodeInfo = types.CodeInfo
	// Deprecated: Do not use.
	ContractInfo = types.ContractInfo
	// Deprecated: Do not use.
	CreatedAt = types.AbsoluteTxPosition
	// Deprecated: Do not use.
	Config = types.WasmConfig
	// Deprecated: Do not use.
	CodeInfoResponse = types.CodeInfoResponse
	// Deprecated: Do not use.
	MessageHandler = keeper.SDKMessageHandler
	// Deprecated: Do not use.
	BankEncoder = keeper.BankEncoder
	// Deprecated: Do not use.
	CustomEncoder = keeper.CustomEncoder
	// Deprecated: Do not use.
	StakingEncoder = keeper.StakingEncoder
	// Deprecated: Do not use.
	WasmEncoder = keeper.WasmEncoder
	// Deprecated: Do not use.
	MessageEncoders = keeper.MessageEncoders
	// Deprecated: Do not use.
	Keeper = keeper.Keeper
	// Deprecated: Do not use.
	QueryHandler = keeper.QueryHandler
	// Deprecated: Do not use.
	CustomQuerier = keeper.CustomQuerier
	// Deprecated: Do not use.
	QueryPlugins = keeper.QueryPlugins
	// Deprecated: Do not use.
	Option = keeper.Option
)
