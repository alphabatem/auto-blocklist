// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package shadow_drive_user_staking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// DeleteAccount is the `deleteAccount` instruction.
type DeleteAccount struct {

	// [0] = [WRITE] storageConfig
	//
	// [1] = [WRITE] userInfo
	//
	// [2] = [WRITE] storageAccount
	//
	// [3] = [WRITE] stakeAccount
	//
	// [4] = [WRITE] owner
	//
	// [5] = [WRITE] shdwPayer
	//
	// [6] = [SIGNER] uploader
	//
	// [7] = [WRITE] emissionsWallet
	//
	// [8] = [] tokenMint
	//
	// [9] = [] systemProgram
	//
	// [10] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDeleteAccountInstructionBuilder creates a new `DeleteAccount` instruction builder.
func NewDeleteAccountInstructionBuilder() *DeleteAccount {
	nd := &DeleteAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetStorageConfigAccount sets the "storageConfig" account.
func (inst *DeleteAccount) SetStorageConfigAccount(storageConfig ag_solanago.PublicKey) *DeleteAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(storageConfig).WRITE()
	return inst
}

// GetStorageConfigAccount gets the "storageConfig" account.
func (inst *DeleteAccount) GetStorageConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUserInfoAccount sets the "userInfo" account.
func (inst *DeleteAccount) SetUserInfoAccount(userInfo ag_solanago.PublicKey) *DeleteAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(userInfo).WRITE()
	return inst
}

// GetUserInfoAccount gets the "userInfo" account.
func (inst *DeleteAccount) GetUserInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetStorageAccountAccount sets the "storageAccount" account.
func (inst *DeleteAccount) SetStorageAccountAccount(storageAccount ag_solanago.PublicKey) *DeleteAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(storageAccount).WRITE()
	return inst
}

// GetStorageAccountAccount gets the "storageAccount" account.
func (inst *DeleteAccount) GetStorageAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetStakeAccountAccount sets the "stakeAccount" account.
func (inst *DeleteAccount) SetStakeAccountAccount(stakeAccount ag_solanago.PublicKey) *DeleteAccount {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(stakeAccount).WRITE()
	return inst
}

// GetStakeAccountAccount gets the "stakeAccount" account.
func (inst *DeleteAccount) GetStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetOwnerAccount sets the "owner" account.
func (inst *DeleteAccount) SetOwnerAccount(owner ag_solanago.PublicKey) *DeleteAccount {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(owner).WRITE()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *DeleteAccount) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetShdwPayerAccount sets the "shdwPayer" account.
func (inst *DeleteAccount) SetShdwPayerAccount(shdwPayer ag_solanago.PublicKey) *DeleteAccount {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(shdwPayer).WRITE()
	return inst
}

// GetShdwPayerAccount gets the "shdwPayer" account.
func (inst *DeleteAccount) GetShdwPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetUploaderAccount sets the "uploader" account.
func (inst *DeleteAccount) SetUploaderAccount(uploader ag_solanago.PublicKey) *DeleteAccount {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(uploader).SIGNER()
	return inst
}

// GetUploaderAccount gets the "uploader" account.
func (inst *DeleteAccount) GetUploaderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetEmissionsWalletAccount sets the "emissionsWallet" account.
func (inst *DeleteAccount) SetEmissionsWalletAccount(emissionsWallet ag_solanago.PublicKey) *DeleteAccount {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(emissionsWallet).WRITE()
	return inst
}

// GetEmissionsWalletAccount gets the "emissionsWallet" account.
func (inst *DeleteAccount) GetEmissionsWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenMintAccount sets the "tokenMint" account.
func (inst *DeleteAccount) SetTokenMintAccount(tokenMint ag_solanago.PublicKey) *DeleteAccount {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenMint)
	return inst
}

// GetTokenMintAccount gets the "tokenMint" account.
func (inst *DeleteAccount) GetTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *DeleteAccount) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *DeleteAccount {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *DeleteAccount) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *DeleteAccount) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *DeleteAccount {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *DeleteAccount) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst DeleteAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DeleteAccount,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DeleteAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DeleteAccount) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.StorageConfig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.UserInfo is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.StorageAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.StakeAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ShdwPayer is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Uploader is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.EmissionsWallet is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenMint is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *DeleteAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DeleteAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  storageConfig", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       userInfo", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        storage", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          stake", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("          owner", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("      shdwPayer", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("       uploader", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("emissionsWallet", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("      tokenMint", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("  systemProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("   tokenProgram", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj DeleteAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *DeleteAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewDeleteAccountInstruction declares a new DeleteAccount instruction with the provided parameters and accounts.
func NewDeleteAccountInstruction(
	// Accounts:
	storageConfig ag_solanago.PublicKey,
	userInfo ag_solanago.PublicKey,
	storageAccount ag_solanago.PublicKey,
	stakeAccount ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	shdwPayer ag_solanago.PublicKey,
	uploader ag_solanago.PublicKey,
	emissionsWallet ag_solanago.PublicKey,
	tokenMint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *DeleteAccount {
	return NewDeleteAccountInstructionBuilder().
		SetStorageConfigAccount(storageConfig).
		SetUserInfoAccount(userInfo).
		SetStorageAccountAccount(storageAccount).
		SetStakeAccountAccount(stakeAccount).
		SetOwnerAccount(owner).
		SetShdwPayerAccount(shdwPayer).
		SetUploaderAccount(uploader).
		SetEmissionsWalletAccount(emissionsWallet).
		SetTokenMintAccount(tokenMint).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram)
}