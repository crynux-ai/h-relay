package tests

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"h_relay/api"
	"h_relay/config"
	"h_relay/migrate"
	"os"
	"strings"
)

var Application *gin.Engine = nil

func init() {

	if err := config.InitConfig(""); err != nil {
		print(err.Error())
		os.Exit(1)
	}

	testAppConfig := config.GetConfig()

	if err := config.InitLog(testAppConfig); err != nil {
		print(err.Error())
		os.Exit(1)
	}

	err := config.InitDB(testAppConfig)
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}

	appConfig := config.GetConfig()

	appConfig.Test.RootPrivateKey = GetTestPrivateKey()

	err = checkPrivateKeyAndAddress(appConfig.Test.RootAddress, appConfig.Test.RootPrivateKey)
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}

	migrate.InitMigration(config.GetDB())

	if err := migrate.Migrate(); err != nil {
		print(err.Error())
		os.Exit(1)
	}

	Application = api.GetHttpApplication(testAppConfig)
}

func checkPrivateKeyAndAddress(address, privateKey string) error {

	if address == "" {
		return errors.New("address not set")
	}

	if privateKey == "" {
		return errors.New("private key not set")
	}

	var testPk string
	if strings.HasPrefix(privateKey, "0x") {
		testPk = privateKey[2:]
	} else {
		testPk = privateKey
	}

	testRootPrivateKey, err := crypto.HexToECDSA(testPk)
	if err != nil {
		return err
	}

	testRootPublicKey := testRootPrivateKey.Public()

	testRootPublicKeyECDSA, ok := testRootPublicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting test root public key to ECDSA")
	}

	testRootAddress := crypto.PubkeyToAddress(*testRootPublicKeyECDSA).Hex()

	if testRootAddress != address {
		return errors.New("account address and private key mismatch")
	}

	return nil
}
