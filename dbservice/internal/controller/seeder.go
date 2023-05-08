package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/palhaziadev/forex-dashboard/pkg/models"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		DB: db,
	}
}

// TODO put this to pkg
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (ctrl *Controller) SeedTestData(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var currencyPairs = []models.CurrencyPair{{Name: "EURUSD"}, {Name: "GBPUSD"}}

	switch r.Method {
	case http.MethodGet:

		ctrl.DB.Migrator().DropTable("testy_test")
		ctrl.DB.Table("testy_test").AutoMigrate(&models.CurrencyPair{})
		ctrl.DB.Table("testy_test").Create(&currencyPairs)

		log.Println("////////////////////")
		for _, curr := range currencyPairs {
			log.Println(curr.ID) // 1,2
		}
		log.Println("////////////////////")

		var currencyHistory = []models.CurrencyHistory{
			{CurrencyID: 1, Value: 1.06513},
			{CurrencyID: 1, Value: 1.06455},
			{CurrencyID: 1, Value: 1.06543},
			{CurrencyID: 1, Value: 1.06594},
			{CurrencyID: 1, Value: 1.06653},
			{CurrencyID: 1, Value: 1.06607},
			{CurrencyID: 1, Value: 1.06702},
			{CurrencyID: 1, Value: 1.0662},
			{CurrencyID: 1, Value: 1.06681},
			{CurrencyID: 1, Value: 1.06605},
			{CurrencyID: 1, Value: 1.06594},
			{CurrencyID: 1, Value: 1.06628},
			{CurrencyID: 1, Value: 1.06664},
			{CurrencyID: 1, Value: 1.06678},
			{CurrencyID: 1, Value: 1.06716},
			{CurrencyID: 1, Value: 1.06756},
			{CurrencyID: 1, Value: 1.06691},
			{CurrencyID: 1, Value: 1.06674},
			{CurrencyID: 1, Value: 1.06622},
			{CurrencyID: 1, Value: 1.06694},
		}

		ctrl.DB.Migrator().DropTable("history")
		ctrl.DB.Table("history").AutoMigrate(&models.CurrencyHistory{})
		ctrl.DB.Table("history").Create(&currencyHistory)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(currencyPairs)

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
