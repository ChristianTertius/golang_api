package tools

import (
	"time"
)

type MockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123abc",
		Username:  "alex",
	},

	"budi": {
		AuthToken: "budi token",
		Username:  "budi",
	},
	"sinta": {
		AuthToken: "sinta token",
		Username:  "sinta",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"budi": {
		Coins:    200,
		Username: "budi",
	},
	"sinta": {
		Coins:    300,
		Username: "sinta",
	},
}

func (d *MockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *MockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *MockDB) SetupDatabase() error {
	return nil
}
