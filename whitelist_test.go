package befi_test

import (
	"testing"

	"github.com/TropicalDog17/befi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestWhitelist(t *testing.T) {
	rpcURL := "https://berachain-rpc.publicnode.com"
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	berachefAddress := common.HexToAddress("0xdf960E8F3F19C481dDE769edEDD439ea1a63426a")
	berachef, err := befi.NewBeraChef(berachefAddress, client)
	if err != nil {
		t.Fatalf("Failed to create BeraChef client: %v", err)
	}

	whileListedVaultAddress := common.HexToAddress("0x66eb42c499372e897929efbf6026821b0a148119")
	if err != nil {
		t.Fatalf("Failed to create vault address: %v", err)
	}
	vault, err := berachef.IsWhitelistedVault(&bind.CallOpts{}, whileListedVaultAddress)
	if err != nil {
		t.Fatalf("Failed to check if vault is whitelisted: %v", err)
	}
	if !vault {
		t.Fatalf("Expected vault to be whitelisted, but it is not")
	}

}
