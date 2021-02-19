package model

type Stock struct {
	ID         int64   `json:"id"`
	SymbolID   string  `json:"symbol_id" gorm:"column:symbol_id"`
	SymbolName string  `json:"symbol_name" gorm:"column:symbol_name"`
	Price      float64 `json:"price" gorm:"column:price"`
}
