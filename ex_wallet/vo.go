package ex_wallet

import (
	"time"
)

// AddressBalanceVo represents the address balance and related information.
type AddressBalanceVo struct {
	ID         string    `json:"id"`         // Serial number
	SId        int       `json:"sId"`        // Service ID
	Balance    string    `json:"balance"`    // Address balance
	Address    string    `json:"address"`    // Address
	CoinType   string    `json:"coinType"`   // Coin type
	Chain      string    `json:"chain"`      // Chain type
	UpdateTime time.Time `json:"updateTime"` // Last update time
}

// CoinConfigVo represents the configuration details of a coin.
type CoinConfigVo struct {
	Code                     string   `json:"code"`                     // Coin name
	Desc                     string   `json:"desc"`                     // Description
	BlockConfirm             int      `json:"blockConfirm"`             // Block confirmation count
	BlockTime                int      `json:"blockTime"`                // Block time
	MainCoinType             string   `json:"mainCoinType"`             // Main coin type
	IsSupportDeposit         bool     `json:"isSupportDeposit"`         // Whether deposit is supported
	IsSupportWithdraw        bool     `json:"isSupportWithdraw"`        // Whether withdrawal is supported
	IsSupportTrade           bool     `json:"isSupportTrade"`           // Whether trading is supported
	IsSupportRedPacket       bool     `json:"isSupportRedPacket"`       // Whether red packets are supported
	ShouldDisplay            bool     `json:"shouldDisplay"`            // Whether to display
	CoinFullName             string   `json:"coinFullName"`             // Full coin name
	ShowCoinName             string   `json:"showCoinName"`             // Display coin name
	MinWithdrawSingle        string   `json:"minWithdrawSingle"`        // Minimum withdrawal amount per transaction
	MaxWithdrawSingle        string   `json:"maxWithdrawSingle"`        // Maximum withdrawal amount per transaction
	MaxWithdrawOneDay        string   `json:"maxWithdrawOneDay"`        // Maximum withdrawal amount per day
	WithdrawFee              string   `json:"withdrawFee"`              // Withdrawal fee
	MinimumTradeAmount       string   `json:"minimumTradeAmount"`       // Minimum trade amount
	MinimumDepositAmount     string   `json:"minimumDepositAmount"`     // Minimum deposit amount
	MinRedPacketAmount       string   `json:"minRedPacketAmount"`       // Minimum red packet amount
	MaxRedPacketAmount       string   `json:"maxRedPacketAmount"`       // Maximum red packet amount
	RedPacketAmountPrecision int      `json:"redPacketAmountPrecision"` // Red packet precision
	InternalWithdrawFee      string   `json:"internalWithdrawFee"`      // Internal transfer fee
	MemoNeeded               bool     `json:"memoNeeded"`               // Whether memo is needed for this coin
	MaxPrecision             int      `json:"maxPrecision"`             // Maximum precision
	AreaType                 int      `json:"areaType"`                 // Area type (main, new, other)
	OrderNo                  int      `json:"orderNo"`                  // Order number for sorting
	Logo                     string   `json:"logo"`                     // Logo URL
	SupportAddrVerify        bool     `json:"supportAddrVerify"`        // Whether address verification is supported
	AddrRegexpExpress        string   `json:"addrRegexpExpress"`        // Address regex format
	SupportMemoVerify        bool     `json:"supportMemoVerify"`        // Whether memo verification is supported
	MemoRegexpExpress        string   `json:"memoRegexpExpress"`        // Memo regex format
	BlockChainHashBuilder    string   `json:"blockChainHashBuilder"`    // Blockchain hash builder URL
	BlockChainAddrBuilder    string   `json:"blockChainAddrBuilder"`    // Blockchain address builder URL
	IsPutOnShelves           bool     `json:"isPutOnShelves"`           // Whether the coin is on shelves
	CommunityURL             string   `json:"communityURL"`             // Community URL
	GroupId                  string   `json:"groupId"`                  // Chat group ID
	IsFrozenTransfer         bool     `json:"isFrozenTransfer"`         // Whether transfers are frozen
	Chains                   []string `json:"chains"`                   // Supported chains
	Source                   string   `json:"source"`                   // Coin source (wallet or cobo)
	AddressSource            string   `json:"addressSource"`            // Address source
	ConfirmNum               int      `json:"confirmNum"`               // Confirmation number
	Confirm                  bool     `json:"confirm"`                  // Whether the coin is confirmed
	UnExamineAmount          string   `json:"unExamineAmount"`          // Amount exempt from review
	UnExamineNum             int      `json:"unExamineNum"`             // Number of times exempt from review
	Serious                  bool     `json:"serious"`                  // Whether it is a token
	SeriousTypes             []string `json:"seriousTypes"`             // Token types
}

// TxEntityVo represents a transaction entity.
type TxEntityVo struct {
	ID              int64     `json:"id"`              // Transaction ID
	SId             int       `json:"sId"`             // Service ID
	BizID           string    `json:"bizId"`           // Business ID
	TxCoin          string    `json:"txCoin"`          // Coin type
	TxType          int       `json:"txType"`          // Transaction type
	TxFromWallet    string    `json:"txFromWallet"`    // From wallet address
	TxToWallet      string    `json:"txToWallet"`      // To wallet address
	TxAmount        string    `json:"txAmount"`        // Transaction amount
	TxFee           string    `json:"txFee"`           // Transaction fee
	TxStatus        int       `json:"txStatus"`        // Transaction status
	TxStatusInner   string    `json:"txStatusInner"`   // Internal transaction status
	TxDepositStatus int       `json:"txDepositStatus"` // Deposit status
	TxDepositDesc   string    `json:"txDepositDesc"`   // Deposit description
	TxNetworkId     string    `json:"txNetworkId"`     // Network ID
	HeightInit      int64     `json:"heightInit"`      // Initial height
	TxNetworkTime   time.Time `json:"txNetworkTime"`   // Network time
	TxNetworkFee    string    `json:"txNetworkFee"`    // Network fee
	UpdateTime      int64     `json:"updateTime"`      // Last update time
	CreatedTime     int64     `json:"createdTime"`     // Creation time
	ExtraInfo       string    `json:"extraInfo"`       // Extra information
	Remark          string    `json:"remark"`          // Remark
	Nonce           string    `json:"nonce"`           // Nonce
	Chain           string    `json:"chain"`           // Chain name
}

// WalletTransactionVo represents a wallet transaction.
type WalletTransactionVo struct {
	ID                    string    `json:"id"`                    // Transaction ID
	SId                   int       `json:"sId"`                   // Service ID (Exchange ID)
	Coin                  string    `json:"coin"`                  // Coin type
	Hash                  string    `json:"hash"`                  // Transaction hash
	BlockHash             string    `json:"blockHash"`             // Block hash
	BlockTime             time.Time `json:"blockTime"`             // Block time
	Amount                string    `json:"amount"`                // Transaction amount
	Fee                   string    `json:"fee"`                   // Transaction fee
	FromAccountOrAddress  string    `json:"fromAccountOrAddress"`  // From account or address
	ToAddress             string    `json:"toAddress"`             // To address
	PaymentCategory       string    `json:"paymentCategory"`       // Payment category
	Status                string    `json:"status"`                // Transaction status
	BlockNumberOrPosition int64     `json:"blockNumberOrPosition"` // Block number or position
	Memo                  string    `json:"memo"`                  // Memo (optional)
	Confirmations         int64     `json:"confirmations"`         // Number of confirmations
	UpdateTime            time.Time `json:"updateTime"`            // Last update time
	CreateTime            time.Time `json:"createTime"`            // Transaction creation time
	IsFinished            bool      `json:"isFinished"`            // Whether the transaction is finished
	DisposeStatusInner    string    `json:"disposeStatusInner"`    // Internal disposal status
	AuditStatus           int       `json:"auditStatus"`           // Audit status
	Source                string    `json:"source"`                // Source of the transaction
	Chain                 string    `json:"chain"`                 // Blockchain name
}

// WalletStatisticVo represents wallet statistics.
type WalletStatisticVo struct {
	ID             int    `json:"id"`             // Wallet statistic ID
	SId            int    `json:"sId"`            // Service ID
	CoinType       string `json:"coinType"`       // Coin type
	Chain          string `json:"chain"`          // Chain type
	HotAmount      string `json:"hotAmount"`      // Hot wallet amount
	ColdAmount     string `json:"coldAmount"`     // Cold wallet amount
	Fee            string `json:"fee"`            // Fee
	DepositCount   int    `json:"depositCount"`   // Deposit count
	WithdrawCount  int    `json:"withdrawCount"`  // Withdrawal count
	UseAddress     int    `json:"useAddress"`     // Used addresses count
	TodayAddress   int    `json:"todayAddress"`   // Addresses used today
	UserAmount     string `json:"userAmount"`     // User amount (balance)
	CreateTime     int64  `json:"createTime"`     // Creation time (Unix timestamp)
	UpdateTime     int64  `json:"updateTime"`     // Last update time (Unix timestamp)
	Time           int64  `json:"time"`           // Time (Unix timestamp)
	DepositAmount  string `json:"depositAmount"`  // Total deposit amount
	WithdrawAmount string `json:"withdrawAmount"` // Total withdrawal amount
	DepositPrice   string `json:"depositPrice"`   // Deposit price
	WithdrawPrice  string `json:"withdrawPrice"`  // Withdrawal price
	Gas            string `json:"gas"`            // Gas (transaction fee)
	ColdCount      int    `json:"coldCount"`      // Cold wallet count
	ColdPrice      string `json:"coldPrice"`      // Cold wallet price
	Price          string `json:"price"`          // Price of the coin
}

type PageInfo[T any] struct {
	PageNum           int   `json:"pageNum"`
	PageSize          int   `json:"pageSize"`
	Size              int   `json:"size"`
	StartRow          int   `json:"startRow"`
	EndRow            int   `json:"endRow"`
	Pages             int   `json:"pages"`
	PrePage           int   `json:"prePage"`
	NextPage          int   `json:"nextPage"`
	IsFirstPage       bool  `json:"isFirstPage"`
	IsLastPage        bool  `json:"isLastPage"`
	HasPreviousPage   bool  `json:"hasPreviousPage"`
	HasNextPage       bool  `json:"hasNextPage"`
	NavigatePages     int   `json:"navigatePages"`
	NavigatePageNums  []int `json:"navigatepageNums"`
	NavigateFirstPage int   `json:"navigateFirstPage"`
	NavigateLastPage  int   `json:"navigateLastPage"`
	Total             int64 `json:"total"` // Using int64 for total to support larger values
	List              []T   `json:"list"`  // 使用泛型 T 来表示 List 字段
}
