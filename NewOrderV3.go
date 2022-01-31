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

// NewOrderV3 is the `NewOrderV3` instruction.
type NewOrderV3 struct {
	Args *NewOrderInstructionV3

	// [0] = [WRITE] market
	// ··········· the market
	//
	// [1] = [WRITE] openOrders
	// ··········· the OpenOrders account to use
	//
	// [2] = [WRITE] requestQueue
	// ··········· the request queue
	//
	// [3] = [WRITE] eventQueue
	// ··········· the event queue
	//
	// [4] = [WRITE] bids
	// ··········· bids
	//
	// [5] = [WRITE] asks
	// ··········· asks
	//
	// [6] = [WRITE] orderPayer
	// ··········· the (coin or price currency) account paying for the order
	//
	// [7] = [SIGNER] owner
	// ··········· owner of the OpenOrders account
	//
	// [8] = [WRITE] coinVault
	// ··········· coin vault
	//
	// [9] = [WRITE] pcVault
	// ··········· pc vault
	//
	// [10] = [] splTokenProgram
	// ··········· spl token program
	//
	// [11] = [] rentSysvar
	// ··········· the rent sysvar
	//
	// [12] = [] feeDiscounts
	// ··········· (optional) the (M)SRM account used for fee discounts
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewNewOrderV3InstructionBuilder creates a new `NewOrderV3` instruction builder.
func NewNewOrderV3InstructionBuilder() *NewOrderV3 {
	nd := &NewOrderV3{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *NewOrderV3) SetArgs(args NewOrderInstructionV3) *NewOrderV3 {
	inst.Args = &args
	return inst
}

// SetMarketAccount sets the "market" account.
// the market
func (inst *NewOrderV3) SetMarketAccount(market ag_solanago.PublicKey) *NewOrderV3 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
// the market
func (inst *NewOrderV3) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetOpenOrdersAccount sets the "openOrders" account.
// the OpenOrders account to use
func (inst *NewOrderV3) SetOpenOrdersAccount(openOrders ag_solanago.PublicKey) *NewOrderV3 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(openOrders).WRITE()
	return inst
}

// GetOpenOrdersAccount gets the "openOrders" account.
// the OpenOrders account to use
func (inst *NewOrderV3) GetOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetRequestQueueAccount sets the "requestQueue" account.
// the request queue
func (inst *NewOrderV3) SetRequestQueueAccount(requestQueue ag_solanago.PublicKey) *NewOrderV3 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(requestQueue).WRITE()
	return inst
}

// GetRequestQueueAccount gets the "requestQueue" account.
// the request queue
func (inst *NewOrderV3) GetRequestQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetEventQueueAccount sets the "eventQueue" account.
// the event queue
func (inst *NewOrderV3) SetEventQueueAccount(eventQueue ag_solanago.PublicKey) *NewOrderV3 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(eventQueue).WRITE()
	return inst
}

// GetEventQueueAccount gets the "eventQueue" account.
// the event queue
func (inst *NewOrderV3) GetEventQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetBidsAccount sets the "bids" account.
// bids
func (inst *NewOrderV3) SetBidsAccount(bids ag_solanago.PublicKey) *NewOrderV3 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
// bids
func (inst *NewOrderV3) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetAsksAccount sets the "asks" account.
// asks
func (inst *NewOrderV3) SetAsksAccount(asks ag_solanago.PublicKey) *NewOrderV3 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(asks).WRITE()
	return inst
}

// GetAsksAccount gets the "asks" account.
// asks
func (inst *NewOrderV3) GetAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetOrderPayerAccount sets the "orderPayer" account.
// the (coin or price currency) account paying for the order
func (inst *NewOrderV3) SetOrderPayerAccount(orderPayer ag_solanago.PublicKey) *NewOrderV3 {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(orderPayer).WRITE()
	return inst
}

// GetOrderPayerAccount gets the "orderPayer" account.
// the (coin or price currency) account paying for the order
func (inst *NewOrderV3) GetOrderPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetOwnerAccount sets the "owner" account.
// owner of the OpenOrders account
func (inst *NewOrderV3) SetOwnerAccount(owner ag_solanago.PublicKey) *NewOrderV3 {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
// owner of the OpenOrders account
func (inst *NewOrderV3) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetCoinVaultAccount sets the "coinVault" account.
// coin vault
func (inst *NewOrderV3) SetCoinVaultAccount(coinVault ag_solanago.PublicKey) *NewOrderV3 {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(coinVault).WRITE()
	return inst
}

// GetCoinVaultAccount gets the "coinVault" account.
// coin vault
func (inst *NewOrderV3) GetCoinVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetPcVaultAccount sets the "pcVault" account.
// pc vault
func (inst *NewOrderV3) SetPcVaultAccount(pcVault ag_solanago.PublicKey) *NewOrderV3 {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(pcVault).WRITE()
	return inst
}

// GetPcVaultAccount gets the "pcVault" account.
// pc vault
func (inst *NewOrderV3) GetPcVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

// SetSplTokenProgramAccount sets the "splTokenProgram" account.
// spl token program
func (inst *NewOrderV3) SetSplTokenProgramAccount(splTokenProgram ag_solanago.PublicKey) *NewOrderV3 {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(splTokenProgram)
	return inst
}

// GetSplTokenProgramAccount gets the "splTokenProgram" account.
// spl token program
func (inst *NewOrderV3) GetSplTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[10]
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// the rent sysvar
func (inst *NewOrderV3) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *NewOrderV3 {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// the rent sysvar
func (inst *NewOrderV3) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[11]
}

// SetFeeDiscountsAccount sets the "feeDiscounts" account.
// (optional) the (M)SRM account used for fee discounts
func (inst *NewOrderV3) SetFeeDiscountsAccount(feeDiscounts ag_solanago.PublicKey) *NewOrderV3 {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(feeDiscounts)
	return inst
}

// GetFeeDiscountsAccount gets the "feeDiscounts" account.
// (optional) the (M)SRM account used for fee discounts
func (inst *NewOrderV3) GetFeeDiscountsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[12]
}

func (inst NewOrderV3) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_NewOrderV3, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst NewOrderV3) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *NewOrderV3) Validate() error {
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
			return errors.New("accounts.EventQueue is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Bids is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Asks is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.OrderPayer is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.CoinVault is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.PcVault is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.SplTokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}

		// [12] = FeeDiscounts is optional

	}
	return nil
}

func (inst *NewOrderV3) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("NewOrderV3")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         market", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("     openOrders", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("   requestQueue", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("     eventQueue", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("           bids", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("           asks", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("     orderPayer", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("          owner", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("      coinVault", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("        pcVault", inst.AccountMetaSlice[9]))
						accountsBranch.Child(ag_format.Meta("splTokenProgram", inst.AccountMetaSlice[10]))
						accountsBranch.Child(ag_format.Meta("     rentSysvar", inst.AccountMetaSlice[11]))
						accountsBranch.Child(ag_format.Meta("   feeDiscounts", inst.AccountMetaSlice[12]))
					})
				})
		})
}

func (obj NewOrderV3) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *NewOrderV3) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewNewOrderV3Instruction declares a new NewOrderV3 instruction with the provided parameters and accounts.
func NewNewOrderV3Instruction(
	// Parameters:
	args NewOrderInstructionV3,
	// Accounts:
	market ag_solanago.PublicKey,
	openOrders ag_solanago.PublicKey,
	requestQueue ag_solanago.PublicKey,
	eventQueue ag_solanago.PublicKey,
	bids ag_solanago.PublicKey,
	asks ag_solanago.PublicKey,
	orderPayer ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	coinVault ag_solanago.PublicKey,
	pcVault ag_solanago.PublicKey,
	splTokenProgram ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey,
	feeDiscounts ag_solanago.PublicKey) *NewOrderV3 {
	return NewNewOrderV3InstructionBuilder().
		SetArgs(args).
		SetMarketAccount(market).
		SetOpenOrdersAccount(openOrders).
		SetRequestQueueAccount(requestQueue).
		SetEventQueueAccount(eventQueue).
		SetBidsAccount(bids).
		SetAsksAccount(asks).
		SetOrderPayerAccount(orderPayer).
		SetOwnerAccount(owner).
		SetCoinVaultAccount(coinVault).
		SetPcVaultAccount(pcVault).
		SetSplTokenProgramAccount(splTokenProgram).
		SetRentSysvarAccount(rentSysvar).
		SetFeeDiscountsAccount(feeDiscounts)
}
