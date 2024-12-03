package ex_wallet

type AddressParam struct {
	Coin string `json:"coin"` // The coin type
	Size int    `json:"size"` // The size of addresses to retrieve
}

type AddressVerifyParam struct {
	Address  string `json:"address"`  // The address to verify
	CoinType string `json:"coinType"` // The type of coin associated with the address
}

type AddressBalanceParam struct {
	Address string `json:"address"`
	Id      string `json:"id"`
	Limit   string `json:"limit"`
}

// TransactionFlowParam represents the parameters for querying transaction flows.
type TransactionFlowParam struct {
	ID          string `json:"id"`          // Unique identifier for the transaction
	CoinType    string `json:"coinType"`    // Type of the coin
	InnerStatus string `json:"innerStatus"` // Internal status of the transaction
	Chain       string `json:"chain"`       // Blockchain identifier
	Start       int64  `json:"start"`       // Start timestamp
	End         int64  `json:"end"`         // End timestamp
	Size        int64  `json:"size"`        // Number of records to fetch
}

// TransactionHashParam represents the parameters for querying a transaction by hash.
type TransactionHashParam struct {
	Hash string `json:"hash"` // Transaction hash
}

// TransactionPageParam represents the parameters for paginated transaction queries.
type TransactionPageParam struct {
	CoinType    string `json:"coinType"`    // Type of the coin
	InnerStatus string `json:"innerStatus"` // Internal status of the transaction
	Chain       string `json:"chain"`       // Blockchain identifier
	Start       int64  `json:"start"`       // Start timestamp
	End         int64  `json:"end"`         // End timestamp
	Size        int    `json:"size"`        // Number of records per page
	Page        int    `json:"page"`        // Page number
}

// UpdPublicKeyParam represents the parameters for updating a public key.
type UpdPublicKeyParam struct {
	PK string `json:"pk"` // The new public key
}

// WalletStatisticParam represents the parameters for querying wallet statistics.
type WalletStatisticParam struct {
	Time int64 `json:"time"` // Timestamp for the statistics
}

// WithdrawCancelParam represents the parameters for canceling a withdrawal.
type WithdrawCancelParam struct {
	BizID string `json:"bizId"` // Business ID
}

// WithdrawCreateParam represents the parameters for creating a withdrawal.
type WithdrawCreateParam struct {
	BizID        string `json:"bizId"`        // Business ID
	TxCoin       string `json:"txCoin"`       // Coin type
	Chain        string `json:"chain"`        // Blockchain chain
	UserID       string `json:"userId"`       // User ID
	TxFromWallet string `json:"txFromWallet"` // Wallet 'from' address
	TxToWallet   string `json:"txToWallet"`   // Wallet 'to' address
	TxAmount     string `json:"txAmount"`     // Withdrawal amount
	TxFee        string `json:"txFee"`        // Withdrawal fee
	Val          string `json:"val"`          // Encrypted parameter
	TransferType string `json:"transferType"` // Transfer type
}

// WithdrawFlowParam represents the parameters for querying withdrawal flows.
type WithdrawFlowParam struct {
	ID       int64  `json:"id"`       // Business ID
	CoinType string `json:"coinType"` // Coin type
	Chain    string `json:"chain"`    // Blockchain chain
	Start    int64  `json:"start"`    // Start time
	End      int64  `json:"end"`      // End time
	Size     int    `json:"size"`     // Size limit
}

// WithdrawGetParam represents the parameters for retrieving a withdrawal.
type WithdrawGetParam struct {
	BizID string `json:"bizId"` // Business ID
}

// WithdrawPageParam represents the parameters for paginated withdrawal queries.
type WithdrawPageParam struct {
	CoinType string `json:"coinType"` // Coin type
	Chain    string `json:"chain"`    // Blockchain chain
	Start    int64  `json:"start"`    // Start time
	End      int64  `json:"end"`      // End time
	Page     int    `json:"page"`     // Page number
	Size     int    `json:"size"`     // Page size
}
