package models

type CurrencyPair struct {
	ID   int64 `gorm:"primaryKey"`
	Name string
}

type CurrencyHistory struct {
	ID         int64 `gorm:"primaryKey"`
	CurrencyID int64
	Value      float64 //TODO new structure
}

type CurrencyData struct {
	Meta         CurrencyMeta
	CurrencyPair string // TODO remove this and use Meta
	Timestamp    string
	Open         float64
	High         float64
	Low          float64
	Close        float64
}

type CurrencyMeta struct {
	BaseCurrency  string
	QuoteCurrency string
}
