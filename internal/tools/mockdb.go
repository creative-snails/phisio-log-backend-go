package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username: "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username: "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username: "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 100,
		Username: "alex",
	},
	"jason": {
		Coins: 200,
		Username: "jason",
	},
	"marie": {
		Coins: 300,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	// var clientData = LoginDetails{} // not needed

	// clientData, ok := mockLoginDetails[username]
	// if !ok {
	// 	return nil
	// }
	// return &clientData

	if clientData, ok := mockLoginDetails[username]; !ok {
		return nil
	} else {
		return &clientData
	}
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	// var clientData = CoinDetails{} // not needed

	// clientData, ok := mockCoinDetails[username]
	// if !ok {
	// 	return nil
	// }

	// return &clientData

	
	if clientData, ok := mockCoinDetails[username]; ok {
		return &clientData
	}

	return nil
}

func (d *mockDB) SetupDatabase() error {
	return nil
}