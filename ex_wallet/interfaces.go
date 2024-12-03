package ex_wallet

type WalletRestfulProxy interface {
	GetUnUsedAddress(param AddressParam) (CommonResponse[[]string], error)
	AddressVerify(param AddressVerifyParam) (CommonResponse[bool], error)
	GetAllocatedAddress(param AddressParam) (CommonResponse[[]string], error)
	AddressBalanceList(param AddressBalanceParam) (CommonResponse[[]AddressBalanceVo], error)
	ListCoinConfig() (CommonResponse[[]CoinConfigVo], error)
	FlowTransactionRecords(param TransactionFlowParam) (CommonResponse[[]WalletTransactionVo], error)
	PageTransactionRecords(param TransactionPageParam) (CommonResponse[PageInfo[WalletTransactionVo]], error)
	FindByHash(param TransactionHashParam) (CommonResponse[WalletTransactionVo], error)
	CreateWithdraw(param WithdrawCreateParam) (CommonResponse[TxEntityVo], error)
	CancelWithdraw(param WithdrawCancelParam) (CommonResponse[any], error)
	GetWithdrawRecord(param WithdrawGetParam) (CommonResponse[TxEntityVo], error)
	FlowWithdrawRecords(param WithdrawFlowParam) (CommonResponse[[]WalletTransactionVo], error)
	PageWithdrawRecords(param WithdrawPageParam) (CommonResponse[PageInfo[WalletTransactionVo]], error)
	GetWalletPublicKey() (CommonResponse[string], error)
	UpdSidPublicKey(param UpdPublicKeyParam) (CommonResponse[any], error)
	WalletStatistic(param WalletStatisticParam) (CommonResponse[[]WalletStatisticVo], error)
}
