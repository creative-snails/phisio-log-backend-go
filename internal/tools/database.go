package tools

type LoginDetails struct {
	AuthToken string
	Username string
}

type CoinDetails struct {
	Coins int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

// func NewDatabase() (*DatabaseInterface, error) {
// 	var database DatabaseInterface = &mockDB{}

// 	var err error = database.SetupDatabase()
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}

// 	return &database, nil
// }

// func NewDatabase() (*mockDB, error) {
// 	database := &mockDB{}

// 	var err error = database.SetupDatabase()
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}

// 	return database, nil
// }

// func NewDatabase() (DatabaseInterface, error) {
// 	database := &mockDB{}

	
// 	if err := database.SetupDatabase(); err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}

// 	return database, nil
// }

func NewDatabase() (DatabaseInterface, error) {
	database := &mockDB{}
	err := database.SetupDatabase()
	return database, err
}