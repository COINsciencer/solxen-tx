// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package sol_xen_miner

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MineHashes is the `mineHashes` instruction.
type MineHashes struct {
	EthAccount *EthAccount
	Kind       *uint8

	// [0] = [WRITE] globalXnRecord
	//
	// [1] = [WRITE] xnByEth
	//
	// [2] = [WRITE] xnBySol
	//
	// [3] = [WRITE, SIGNER] user
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMineHashesInstructionBuilder creates a new `MineHashes` instruction builder.
func NewMineHashesInstructionBuilder() *MineHashes {
	nd := &MineHashes{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetEthAccount sets the "ethAccount" parameter.
func (inst *MineHashes) SetEthAccount(ethAccount EthAccount) *MineHashes {
	inst.EthAccount = &ethAccount
	return inst
}

// SetKind sets the "kind" parameter.
func (inst *MineHashes) SetKind(kind uint8) *MineHashes {
	inst.Kind = &kind
	return inst
}

// SetGlobalXnRecordAccount sets the "globalXnRecord" account.
func (inst *MineHashes) SetGlobalXnRecordAccount(globalXnRecord ag_solanago.PublicKey) *MineHashes {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(globalXnRecord).WRITE()
	return inst
}

// GetGlobalXnRecordAccount gets the "globalXnRecord" account.
func (inst *MineHashes) GetGlobalXnRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetXnByEthAccount sets the "xnByEth" account.
func (inst *MineHashes) SetXnByEthAccount(xnByEth ag_solanago.PublicKey) *MineHashes {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(xnByEth).WRITE()
	return inst
}

// GetXnByEthAccount gets the "xnByEth" account.
func (inst *MineHashes) GetXnByEthAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetXnBySolAccount sets the "xnBySol" account.
func (inst *MineHashes) SetXnBySolAccount(xnBySol ag_solanago.PublicKey) *MineHashes {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(xnBySol).WRITE()
	return inst
}

// GetXnBySolAccount gets the "xnBySol" account.
func (inst *MineHashes) GetXnBySolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserAccount sets the "user" account.
func (inst *MineHashes) SetUserAccount(user ag_solanago.PublicKey) *MineHashes {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *MineHashes) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *MineHashes) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *MineHashes {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *MineHashes) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst MineHashes) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MineHashes,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MineHashes) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MineHashes) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.EthAccount == nil {
			return errors.New("EthAccount parameter is not set")
		}
		if inst.Kind == nil {
			return errors.New("Kind parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.GlobalXnRecord is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.XnByEth is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.XnBySol is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *MineHashes) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MineHashes")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("EthAccount", *inst.EthAccount))
						paramsBranch.Child(ag_format.Param("      Kind", *inst.Kind))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("globalXnRecord", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       xnByEth", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("       xnBySol", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          user", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta(" systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj MineHashes) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `EthAccount` param:
	err = encoder.Encode(obj.EthAccount)
	if err != nil {
		return err
	}
	// Serialize `Kind` param:
	err = encoder.Encode(obj.Kind)
	if err != nil {
		return err
	}
	return nil
}
func (obj *MineHashes) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `EthAccount`:
	err = decoder.Decode(&obj.EthAccount)
	if err != nil {
		return err
	}
	// Deserialize `Kind`:
	err = decoder.Decode(&obj.Kind)
	if err != nil {
		return err
	}
	return nil
}

// NewMineHashesInstruction declares a new MineHashes instruction with the provided parameters and accounts.
func NewMineHashesInstruction(
	// Parameters:
	ethAccount EthAccount,
	kind uint8,
	// Accounts:
	globalXnRecord ag_solanago.PublicKey,
	xnByEth ag_solanago.PublicKey,
	xnBySol ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *MineHashes {
	return NewMineHashesInstructionBuilder().
		SetEthAccount(ethAccount).
		SetKind(kind).
		SetGlobalXnRecordAccount(globalXnRecord).
		SetXnByEthAccount(xnByEth).
		SetXnBySolAccount(xnBySol).
		SetUserAccount(user).
		SetSystemProgramAccount(systemProgram)
}
