// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package shadow_drive_user_staking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RequestDeleteAccount2 is the `requestDeleteAccount2` instruction.
type RequestDeleteAccount2 struct {

	// [0] = [] storageConfig
	//
	// [1] = [WRITE] storageAccount
	//
	// [2] = [WRITE, SIGNER] owner
	//
	// [3] = [] tokenMint
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRequestDeleteAccount2InstructionBuilder creates a new `RequestDeleteAccount2` instruction builder.
func NewRequestDeleteAccount2InstructionBuilder() *RequestDeleteAccount2 {
	nd := &RequestDeleteAccount2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetStorageConfigAccount sets the "storageConfig" account.
func (inst *RequestDeleteAccount2) SetStorageConfigAccount(storageConfig ag_solanago.PublicKey) *RequestDeleteAccount2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(storageConfig)
	return inst
}

// GetStorageConfigAccount gets the "storageConfig" account.
func (inst *RequestDeleteAccount2) GetStorageConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStorageAccountAccount sets the "storageAccount" account.
func (inst *RequestDeleteAccount2) SetStorageAccountAccount(storageAccount ag_solanago.PublicKey) *RequestDeleteAccount2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(storageAccount).WRITE()
	return inst
}

// GetStorageAccountAccount gets the "storageAccount" account.
func (inst *RequestDeleteAccount2) GetStorageAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOwnerAccount sets the "owner" account.
func (inst *RequestDeleteAccount2) SetOwnerAccount(owner ag_solanago.PublicKey) *RequestDeleteAccount2 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *RequestDeleteAccount2) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenMintAccount sets the "tokenMint" account.
func (inst *RequestDeleteAccount2) SetTokenMintAccount(tokenMint ag_solanago.PublicKey) *RequestDeleteAccount2 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenMint)
	return inst
}

// GetTokenMintAccount gets the "tokenMint" account.
func (inst *RequestDeleteAccount2) GetTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *RequestDeleteAccount2) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *RequestDeleteAccount2 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *RequestDeleteAccount2) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst RequestDeleteAccount2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RequestDeleteAccount2,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RequestDeleteAccount2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RequestDeleteAccount2) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.StorageConfig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.StorageAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *RequestDeleteAccount2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RequestDeleteAccount2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("storageConfig", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      storage", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        owner", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    tokenMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj RequestDeleteAccount2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RequestDeleteAccount2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRequestDeleteAccount2Instruction declares a new RequestDeleteAccount2 instruction with the provided parameters and accounts.
func NewRequestDeleteAccount2Instruction(
	// Accounts:
	storageConfig ag_solanago.PublicKey,
	storageAccount ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	tokenMint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *RequestDeleteAccount2 {
	return NewRequestDeleteAccount2InstructionBuilder().
		SetStorageConfigAccount(storageConfig).
		SetStorageAccountAccount(storageAccount).
		SetOwnerAccount(owner).
		SetTokenMintAccount(tokenMint).
		SetSystemProgramAccount(systemProgram)
}