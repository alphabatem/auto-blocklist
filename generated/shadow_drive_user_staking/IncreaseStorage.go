// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package shadow_drive_user_staking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// IncreaseStorage is the `increaseStorage` instruction.
type IncreaseStorage struct {
	AdditionalStorage *uint64

	// [0] = [] storageConfig
	//
	// [1] = [WRITE] storageAccount
	//
	// [2] = [WRITE, SIGNER] owner
	//
	// [3] = [WRITE] ownerAta
	//
	// [4] = [WRITE] stakeAccount
	//
	// [5] = [SIGNER] uploader
	//
	// [6] = [] tokenMint
	//
	// [7] = [] systemProgram
	//
	// [8] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewIncreaseStorageInstructionBuilder creates a new `IncreaseStorage` instruction builder.
func NewIncreaseStorageInstructionBuilder() *IncreaseStorage {
	nd := &IncreaseStorage{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetAdditionalStorage sets the "additionalStorage" parameter.
func (inst *IncreaseStorage) SetAdditionalStorage(additionalStorage uint64) *IncreaseStorage {
	inst.AdditionalStorage = &additionalStorage
	return inst
}

// SetStorageConfigAccount sets the "storageConfig" account.
func (inst *IncreaseStorage) SetStorageConfigAccount(storageConfig ag_solanago.PublicKey) *IncreaseStorage {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(storageConfig)
	return inst
}

// GetStorageConfigAccount gets the "storageConfig" account.
func (inst *IncreaseStorage) GetStorageConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStorageAccountAccount sets the "storageAccount" account.
func (inst *IncreaseStorage) SetStorageAccountAccount(storageAccount ag_solanago.PublicKey) *IncreaseStorage {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(storageAccount).WRITE()
	return inst
}

// GetStorageAccountAccount gets the "storageAccount" account.
func (inst *IncreaseStorage) GetStorageAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOwnerAccount sets the "owner" account.
func (inst *IncreaseStorage) SetOwnerAccount(owner ag_solanago.PublicKey) *IncreaseStorage {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *IncreaseStorage) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOwnerAtaAccount sets the "ownerAta" account.
func (inst *IncreaseStorage) SetOwnerAtaAccount(ownerAta ag_solanago.PublicKey) *IncreaseStorage {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(ownerAta).WRITE()
	return inst
}

// GetOwnerAtaAccount gets the "ownerAta" account.
func (inst *IncreaseStorage) GetOwnerAtaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetStakeAccountAccount sets the "stakeAccount" account.
func (inst *IncreaseStorage) SetStakeAccountAccount(stakeAccount ag_solanago.PublicKey) *IncreaseStorage {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(stakeAccount).WRITE()
	return inst
}

// GetStakeAccountAccount gets the "stakeAccount" account.
func (inst *IncreaseStorage) GetStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetUploaderAccount sets the "uploader" account.
func (inst *IncreaseStorage) SetUploaderAccount(uploader ag_solanago.PublicKey) *IncreaseStorage {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(uploader).SIGNER()
	return inst
}

// GetUploaderAccount gets the "uploader" account.
func (inst *IncreaseStorage) GetUploaderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenMintAccount sets the "tokenMint" account.
func (inst *IncreaseStorage) SetTokenMintAccount(tokenMint ag_solanago.PublicKey) *IncreaseStorage {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenMint)
	return inst
}

// GetTokenMintAccount gets the "tokenMint" account.
func (inst *IncreaseStorage) GetTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *IncreaseStorage) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *IncreaseStorage {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *IncreaseStorage) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *IncreaseStorage) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *IncreaseStorage {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *IncreaseStorage) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst IncreaseStorage) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_IncreaseStorage,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst IncreaseStorage) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *IncreaseStorage) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.AdditionalStorage == nil {
			return errors.New("AdditionalStorage parameter is not set")
		}
	}

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
			return errors.New("accounts.OwnerAta is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.StakeAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Uploader is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *IncreaseStorage) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("IncreaseStorage")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("AdditionalStorage", *inst.AdditionalStorage))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("storageConfig", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      storage", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        owner", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("     ownerAta", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        stake", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("     uploader", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("    tokenMint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta(" tokenProgram", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj IncreaseStorage) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `AdditionalStorage` param:
	err = encoder.Encode(obj.AdditionalStorage)
	if err != nil {
		return err
	}
	return nil
}
func (obj *IncreaseStorage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `AdditionalStorage`:
	err = decoder.Decode(&obj.AdditionalStorage)
	if err != nil {
		return err
	}
	return nil
}

// NewIncreaseStorageInstruction declares a new IncreaseStorage instruction with the provided parameters and accounts.
func NewIncreaseStorageInstruction(
	// Parameters:
	additionalStorage uint64,
	// Accounts:
	storageConfig ag_solanago.PublicKey,
	storageAccount ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	ownerAta ag_solanago.PublicKey,
	stakeAccount ag_solanago.PublicKey,
	uploader ag_solanago.PublicKey,
	tokenMint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *IncreaseStorage {
	return NewIncreaseStorageInstructionBuilder().
		SetAdditionalStorage(additionalStorage).
		SetStorageConfigAccount(storageConfig).
		SetStorageAccountAccount(storageAccount).
		SetOwnerAccount(owner).
		SetOwnerAtaAccount(ownerAta).
		SetStakeAccountAccount(stakeAccount).
		SetUploaderAccount(uploader).
		SetTokenMintAccount(tokenMint).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram)
}
