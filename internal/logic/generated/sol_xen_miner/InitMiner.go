// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package sol_xen_miner

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitMiner is the `initMiner` instruction.
type InitMiner struct {
	Kind *uint8

	// [0] = [WRITE, SIGNER] admin
	//
	// [1] = [WRITE] globalXnRecord
	//
	// [2] = [] systemProgram
	//
	// [3] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitMinerInstructionBuilder creates a new `InitMiner` instruction builder.
func NewInitMinerInstructionBuilder() *InitMiner {
	nd := &InitMiner{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetKind sets the "kind" parameter.
func (inst *InitMiner) SetKind(kind uint8) *InitMiner {
	inst.Kind = &kind
	return inst
}

// SetAdminAccount sets the "admin" account.
func (inst *InitMiner) SetAdminAccount(admin ag_solanago.PublicKey) *InitMiner {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *InitMiner) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetGlobalXnRecordAccount sets the "globalXnRecord" account.
func (inst *InitMiner) SetGlobalXnRecordAccount(globalXnRecord ag_solanago.PublicKey) *InitMiner {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(globalXnRecord).WRITE()
	return inst
}

// GetGlobalXnRecordAccount gets the "globalXnRecord" account.
func (inst *InitMiner) GetGlobalXnRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitMiner) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitMiner {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitMiner) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetRentAccount sets the "rent" account.
func (inst *InitMiner) SetRentAccount(rent ag_solanago.PublicKey) *InitMiner {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *InitMiner) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst InitMiner) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitMiner,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitMiner) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitMiner) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Kind == nil {
			return errors.New("Kind parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.GlobalXnRecord is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *InitMiner) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitMiner")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Kind", *inst.Kind))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         admin", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("globalXnRecord", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta(" systemProgram", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          rent", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj InitMiner) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Kind` param:
	err = encoder.Encode(obj.Kind)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitMiner) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Kind`:
	err = decoder.Decode(&obj.Kind)
	if err != nil {
		return err
	}
	return nil
}

// NewInitMinerInstruction declares a new InitMiner instruction with the provided parameters and accounts.
func NewInitMinerInstruction(
	// Parameters:
	kind uint8,
	// Accounts:
	admin ag_solanago.PublicKey,
	globalXnRecord ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *InitMiner {
	return NewInitMinerInstructionBuilder().
		SetKind(kind).
		SetAdminAccount(admin).
		SetGlobalXnRecordAccount(globalXnRecord).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
