// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package serumgo

import (
	"encoding/binary"
	"errors"

	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CancelOrderByClientId is the `CancelOrderByClientId` instruction.
type CancelOrderByClientId struct {
	ClientOrderID *uint64

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

// NewCancelOrderByClientIdInstructionBuilder creates a new `CancelOrderByClientId` instruction builder.
func NewCancelOrderByClientIdInstructionBuilder() *CancelOrderByClientId {
	nd := &CancelOrderByClientId{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetClientOrderID sets the "clientOrderID" parameter.
func (inst *CancelOrderByClientId) SetClientOrderID(clientOrderID uint64) *CancelOrderByClientId {
	inst.ClientOrderID = &clientOrderID
	return inst
}

// SetMarketAccount sets the "market" account.
// market
func (inst *CancelOrderByClientId) SetMarketAccount(market ag_solanago.PublicKey) *CancelOrderByClientId {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(market)
	return inst
}

// GetMarketAccount gets the "market" account.
// market
func (inst *CancelOrderByClientId) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOpenOrdersAccount sets the "openOrders" account.
// OpenOrders
func (inst *CancelOrderByClientId) SetOpenOrdersAccount(openOrders ag_solanago.PublicKey) *CancelOrderByClientId {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(openOrders).WRITE()
	return inst
}

// GetOpenOrdersAccount gets the "openOrders" account.
// OpenOrders
func (inst *CancelOrderByClientId) GetOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetRequestQueueAccount sets the "requestQueue" account.
// the request queue
func (inst *CancelOrderByClientId) SetRequestQueueAccount(requestQueue ag_solanago.PublicKey) *CancelOrderByClientId {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(requestQueue).WRITE()
	return inst
}

// GetRequestQueueAccount gets the "requestQueue" account.
// the request queue
func (inst *CancelOrderByClientId) GetRequestQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOwnerAccount sets the "owner" account.
// the OpenOrders owner
func (inst *CancelOrderByClientId) SetOwnerAccount(owner ag_solanago.PublicKey) *CancelOrderByClientId {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
// the OpenOrders owner
func (inst *CancelOrderByClientId) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst CancelOrderByClientId) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_CancelOrderByClientId, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CancelOrderByClientId) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CancelOrderByClientId) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ClientOrderID == nil {
			return errors.New("ClientOrderID parameter is not set")
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

func (inst *CancelOrderByClientId) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CancelOrderByClientId")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ClientOrderID", *inst.ClientOrderID))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      market", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  openOrders", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("requestQueue", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       owner", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj CancelOrderByClientId) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ClientOrderID` param:
	err = encoder.Encode(obj.ClientOrderID)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CancelOrderByClientId) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ClientOrderID`:
	err = decoder.Decode(&obj.ClientOrderID)
	if err != nil {
		return err
	}
	return nil
}

// NewCancelOrderByClientIdInstruction declares a new CancelOrderByClientId instruction with the provided parameters and accounts.
func NewCancelOrderByClientIdInstruction(
	// Parameters:
	clientOrderID uint64,
	// Accounts:
	market ag_solanago.PublicKey,
	openOrders ag_solanago.PublicKey,
	requestQueue ag_solanago.PublicKey,
	owner ag_solanago.PublicKey) *CancelOrderByClientId {
	return NewCancelOrderByClientIdInstructionBuilder().
		SetClientOrderID(clientOrderID).
		SetMarketAccount(market).
		SetOpenOrdersAccount(openOrders).
		SetRequestQueueAccount(requestQueue).
		SetOwnerAccount(owner)
}
