// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package serum_dex

import (
	"encoding/binary"
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CancelOrder is the `CancelOrder` instruction.
type CancelOrder struct {
	Args *CancelOrderInstructionV2

	// [0] = [] market
	// ··········· market
	//
	// [1] = [WRITE] openOrders
	// ··········· OpenOrders
	//
	// [2] = [WRITE] requestQueue
	// ··········· the request queue
	//
	// [3] = [SIGNER] owner
	// ··········· the OpenOrders owner
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCancelOrderInstructionBuilder creates a new `CancelOrder` instruction builder.
func NewCancelOrderInstructionBuilder() *CancelOrder {
	nd := &CancelOrder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *CancelOrder) SetArgs(args CancelOrderInstructionV2) *CancelOrder {
	inst.Args = &args
	return inst
}

// SetMarketAccount sets the "market" account.
// market
func (inst *CancelOrder) SetMarketAccount(market ag_solanago.PublicKey) *CancelOrder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(market)
	return inst
}

// GetMarketAccount gets the "market" account.
// market
func (inst *CancelOrder) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetOpenOrdersAccount sets the "openOrders" account.
// OpenOrders
func (inst *CancelOrder) SetOpenOrdersAccount(openOrders ag_solanago.PublicKey) *CancelOrder {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(openOrders).WRITE()
	return inst
}

// GetOpenOrdersAccount gets the "openOrders" account.
// OpenOrders
func (inst *CancelOrder) GetOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetRequestQueueAccount sets the "requestQueue" account.
// the request queue
func (inst *CancelOrder) SetRequestQueueAccount(requestQueue ag_solanago.PublicKey) *CancelOrder {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(requestQueue).WRITE()
	return inst
}

// GetRequestQueueAccount gets the "requestQueue" account.
// the request queue
func (inst *CancelOrder) GetRequestQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetOwnerAccount sets the "owner" account.
// the OpenOrders owner
func (inst *CancelOrder) SetOwnerAccount(owner ag_solanago.PublicKey) *CancelOrder {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
// the OpenOrders owner
func (inst *CancelOrder) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst CancelOrder) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_CancelOrder, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CancelOrder) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CancelOrder) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OpenOrders is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.RequestQueue is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Owner is not set")
		}
	}
	return nil
}

func (inst *CancelOrder) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CancelOrder")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      market", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("  openOrders", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("requestQueue", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("       owner", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj CancelOrder) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CancelOrder) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewCancelOrderInstruction declares a new CancelOrder instruction with the provided parameters and accounts.
func NewCancelOrderInstruction(
	// Parameters:
	args CancelOrderInstructionV2,
	// Accounts:
	market ag_solanago.PublicKey,
	openOrders ag_solanago.PublicKey,
	requestQueue ag_solanago.PublicKey,
	owner ag_solanago.PublicKey) *CancelOrder {
	return NewCancelOrderInstructionBuilder().
		SetArgs(args).
		SetMarketAccount(market).
		SetOpenOrdersAccount(openOrders).
		SetRequestQueueAccount(requestQueue).
		SetOwnerAccount(owner)
}
