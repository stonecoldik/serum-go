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

// SendTake is the `SendTake` instruction.
type SendTake struct {
	Args *SendTakeInstruction

	// [0] = [WRITE] market
	// ··········· market
	//
	// [1] = [WRITE] requestQueue
	// ··········· the request queue
	//
	// [2] = [WRITE] eventQueue
	// ··········· the event queue
	//
	// [3] = [WRITE] bids
	// ··········· bids
	//
	// [4] = [WRITE] asks
	// ··········· asks
	//
	// [5] = [WRITE] coinCurrencyWallet
	// ··········· the coin currency wallet account
	//
	// [6] = [WRITE] priceCurrencyWallet
	// ··········· the price currency wallet account
	//
	// [7] = [] signer
	// ··········· signer
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
	// [11] = [] feeDiscounts
	// ··········· (optional) the (M)SRM account used for fee discounts
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSendTakeInstructionBuilder creates a new `SendTake` instruction builder.
func NewSendTakeInstructionBuilder() *SendTake {
	nd := &SendTake{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 12),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *SendTake) SetArgs(args SendTakeInstruction) *SendTake {
	inst.Args = &args
	return inst
}

// SetMarketAccount sets the "market" account.
// market
func (inst *SendTake) SetMarketAccount(market ag_solanago.PublicKey) *SendTake {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
// market
func (inst *SendTake) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetRequestQueueAccount sets the "requestQueue" account.
// the request queue
func (inst *SendTake) SetRequestQueueAccount(requestQueue ag_solanago.PublicKey) *SendTake {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(requestQueue).WRITE()
	return inst
}

// GetRequestQueueAccount gets the "requestQueue" account.
// the request queue
func (inst *SendTake) GetRequestQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetEventQueueAccount sets the "eventQueue" account.
// the event queue
func (inst *SendTake) SetEventQueueAccount(eventQueue ag_solanago.PublicKey) *SendTake {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(eventQueue).WRITE()
	return inst
}

// GetEventQueueAccount gets the "eventQueue" account.
// the event queue
func (inst *SendTake) GetEventQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetBidsAccount sets the "bids" account.
// bids
func (inst *SendTake) SetBidsAccount(bids ag_solanago.PublicKey) *SendTake {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
// bids
func (inst *SendTake) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetAsksAccount sets the "asks" account.
// asks
func (inst *SendTake) SetAsksAccount(asks ag_solanago.PublicKey) *SendTake {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(asks).WRITE()
	return inst
}

// GetAsksAccount gets the "asks" account.
// asks
func (inst *SendTake) GetAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetCoinCurrencyWalletAccount sets the "coinCurrencyWallet" account.
// the coin currency wallet account
func (inst *SendTake) SetCoinCurrencyWalletAccount(coinCurrencyWallet ag_solanago.PublicKey) *SendTake {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(coinCurrencyWallet).WRITE()
	return inst
}

// GetCoinCurrencyWalletAccount gets the "coinCurrencyWallet" account.
// the coin currency wallet account
func (inst *SendTake) GetCoinCurrencyWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetPriceCurrencyWalletAccount sets the "priceCurrencyWallet" account.
// the price currency wallet account
func (inst *SendTake) SetPriceCurrencyWalletAccount(priceCurrencyWallet ag_solanago.PublicKey) *SendTake {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(priceCurrencyWallet).WRITE()
	return inst
}

// GetPriceCurrencyWalletAccount gets the "priceCurrencyWallet" account.
// the price currency wallet account
func (inst *SendTake) GetPriceCurrencyWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetSignerAccount sets the "signer" account.
// signer
func (inst *SendTake) SetSignerAccount(signer ag_solanago.PublicKey) *SendTake {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(signer)
	return inst
}

// GetSignerAccount gets the "signer" account.
// signer
func (inst *SendTake) GetSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetCoinVaultAccount sets the "coinVault" account.
// coin vault
func (inst *SendTake) SetCoinVaultAccount(coinVault ag_solanago.PublicKey) *SendTake {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(coinVault).WRITE()
	return inst
}

// GetCoinVaultAccount gets the "coinVault" account.
// coin vault
func (inst *SendTake) GetCoinVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetPcVaultAccount sets the "pcVault" account.
// pc vault
func (inst *SendTake) SetPcVaultAccount(pcVault ag_solanago.PublicKey) *SendTake {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(pcVault).WRITE()
	return inst
}

// GetPcVaultAccount gets the "pcVault" account.
// pc vault
func (inst *SendTake) GetPcVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

// SetSplTokenProgramAccount sets the "splTokenProgram" account.
// spl token program
func (inst *SendTake) SetSplTokenProgramAccount(splTokenProgram ag_solanago.PublicKey) *SendTake {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(splTokenProgram)
	return inst
}

// GetSplTokenProgramAccount gets the "splTokenProgram" account.
// spl token program
func (inst *SendTake) GetSplTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[10]
}

// SetFeeDiscountsAccount sets the "feeDiscounts" account.
// (optional) the (M)SRM account used for fee discounts
func (inst *SendTake) SetFeeDiscountsAccount(feeDiscounts ag_solanago.PublicKey) *SendTake {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(feeDiscounts)
	return inst
}

// GetFeeDiscountsAccount gets the "feeDiscounts" account.
// (optional) the (M)SRM account used for fee discounts
func (inst *SendTake) GetFeeDiscountsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[11]
}

func (inst SendTake) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_SendTake, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SendTake) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SendTake) Validate() error {
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
			return errors.New("accounts.RequestQueue is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.EventQueue is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Bids is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Asks is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.CoinCurrencyWallet is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PriceCurrencyWallet is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Signer is not set")
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
			return errors.New("accounts.FeeDiscounts is not set")
		}
	}
	return nil
}

func (inst *SendTake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SendTake")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=12]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             market", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("       requestQueue", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("         eventQueue", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("               bids", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("               asks", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta(" coinCurrencyWallet", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("priceCurrencyWallet", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("             signer", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("          coinVault", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("            pcVault", inst.AccountMetaSlice[9]))
						accountsBranch.Child(ag_format.Meta("    splTokenProgram", inst.AccountMetaSlice[10]))
						accountsBranch.Child(ag_format.Meta("       feeDiscounts", inst.AccountMetaSlice[11]))
					})
				})
		})
}

func (obj SendTake) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SendTake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewSendTakeInstruction declares a new SendTake instruction with the provided parameters and accounts.
func NewSendTakeInstruction(
	// Parameters:
	args SendTakeInstruction,
	// Accounts:
	market ag_solanago.PublicKey,
	requestQueue ag_solanago.PublicKey,
	eventQueue ag_solanago.PublicKey,
	bids ag_solanago.PublicKey,
	asks ag_solanago.PublicKey,
	coinCurrencyWallet ag_solanago.PublicKey,
	priceCurrencyWallet ag_solanago.PublicKey,
	signer ag_solanago.PublicKey,
	coinVault ag_solanago.PublicKey,
	pcVault ag_solanago.PublicKey,
	splTokenProgram ag_solanago.PublicKey,
	feeDiscounts ag_solanago.PublicKey) *SendTake {
	return NewSendTakeInstructionBuilder().
		SetArgs(args).
		SetMarketAccount(market).
		SetRequestQueueAccount(requestQueue).
		SetEventQueueAccount(eventQueue).
		SetBidsAccount(bids).
		SetAsksAccount(asks).
		SetCoinCurrencyWalletAccount(coinCurrencyWallet).
		SetPriceCurrencyWalletAccount(priceCurrencyWallet).
		SetSignerAccount(signer).
		SetCoinVaultAccount(coinVault).
		SetPcVaultAccount(pcVault).
		SetSplTokenProgramAccount(splTokenProgram).
		SetFeeDiscountsAccount(feeDiscounts)
}