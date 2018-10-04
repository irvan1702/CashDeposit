package customer

type Customer struct {
	CustomerName       string `bson:"customer_name" json:"customer_name"`
	AccountNumber      string `bson:"account_number" json:"account_number"`
	CashDeposit        int    `bson:"cash_deposit" json:"cash_deposit"`
	CashDepositHistory []int  `bson:"cash_deposit_history" json:"cash_deposit_history"`
	TotalCashDeposit   int    `bson:"total_cash_deposit" json:"total_cash_deposit"`
}
