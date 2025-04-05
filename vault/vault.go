package vault

import "github.com/ethereum/go-ethereum/common"

// WhitelistedVault represents vault information
type WhitelistedVault struct {
	Address  common.Address
	Metadata string
}
