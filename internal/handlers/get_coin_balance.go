package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/creative-snails/phisio-log-backend-go/api"
	"github.com/creative-snails/phisio-log-backend-go/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	if err = decoder.Decode(&params, r.URL.Query()); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database tools.DatabaseInterface
	if database, err = tools.NewDatabase(); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// var tokenDetails *tools.CoinDetails
	tokenDetails := database.GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		// Balance: (*tokenDetails).Coins,
		Balance: tokenDetails.Coins,
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	// err = json.NewEncoder(w).Encode(response)
	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}