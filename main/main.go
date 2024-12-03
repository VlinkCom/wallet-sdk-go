package main

import (
	"awesomeProject/ex_wallet"
	"log"
)

func main() {
	testGetUnUsedAddress()
	//testAddressVerify()
	//testGetAllocatedAddress()
	//testAddressBalanceList()
	//testListCoinConfig()
	//testCreateWithdraw()
	//testCancelWithdraw()
}

func testGetUnUsedAddress() {
	impl := ex_wallet.NewWalletRestfulProxyImpl(
		"https://waalll.ydykknd.com/rest",
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0Fg6/ZB+h6JL6EyjhXVMt7iAxxQ+BDrmMozBXbtNxyYtYdAcSu9zfDhloCug5shUzLrla4Avwrl5gfa31EGTSE1rKiGOU/5POoCEqYiuB4+k7rvhYUl+D2xUBRDba+u4CFKKcJ2tzIhXfJ8XcNyXnBHwW0EYpTFzQQc4LxOoAJQIDAQAB",
		"MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBALQWDr9kH6HokvoTKOFdUy3uIDHFD4EOuYyjMFdu03HJi1h0BxK73N8OGWgK6DmyFTMuuVrgC/CuXmB9rfUQZNITWsqIY5T/k86gISpiK4Hj6Tuu+FhSX4PbFQFENtr67gIUopwna3MiFd8nxdw3JecEfBbQRilMXNBBzgvE6gAlAgMBAAECgYAZ1HAZE4tV9b3fEJB0gqZDmDwV52Xp5OrIOT/dyo+ZaocKENpS3Y84jxUlzyqv29MenD0Jw8jypI/01xH0bH2JRkkk0hTEtZlKu22DtlUlrdUjE1AkYNap2UTmqRJAT9nHbSuILBHkCmu8J+4/GLlF0fGoLfcj3jpbmrxRdcD1oQJBAMKdvNn8PAYdyZqN3SfXOwvYRynTTpk4TRAih6bTXFh9OQpQzsbTo4PqeddMaHpl9BlmgQ5PkqkKW4iFYAdLc7cCQQDs4xi0dfpad/ixQIiaK2zj+PbniQy/K7dW2vJfDWW9ZljdW3YFd1hTAnl6BcnTLjth1sepJPH/h8jm0kB9t4MDAkAp/55ksHFHpKAAMYM7eNAuQzNATHCW0jaXN8xsbQptskBgAIZPFBcifbRjiQK2/0+JN7y8GOH+htSwBpBtopPLAkAVST7SEHvjnz4wp0zfFaHeRmPhYFZN417BJkF7Oxf2UfkQQwSy7wt3whQqW9TL7i6V1xfv4gsA3DG4VBJJebEDAkBpXc/gB94ahr3y6k3br+zDKrcyy34E3nX338oj28/+aTGsGWdgsz7uel8/dAeVWd5hmVM9mj2pzO6sEijgp3TK")
	address, _ := impl.GetUnUsedAddress(ex_wallet.AddressParam{
		Coin: "TON",
		Size: 1,
	})
	log.Println(address)
}

func testAddressVerify() {
	impl := ex_wallet.NewWalletRestfulProxyImpl(
		"https://waalll.ydykknd.com/rest",
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0Fg6/ZB+h6JL6EyjhXVMt7iAxxQ+BDrmMozBXbtNxyYtYdAcSu9zfDhloCug5shUzLrla4Avwrl5gfa31EGTSE1rKiGOU/5POoCEqYiuB4+k7rvhYUl+D2xUBRDba+u4CFKKcJ2tzIhXfJ8XcNyXnBHwW0EYpTFzQQc4LxOoAJQIDAQAB",
		"MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBALQWDr9kH6HokvoTKOFdUy3uIDHFD4EOuYyjMFdu03HJi1h0BxK73N8OGWgK6DmyFTMuuVrgC/CuXmB9rfUQZNITWsqIY5T/k86gISpiK4Hj6Tuu+FhSX4PbFQFENtr67gIUopwna3MiFd8nxdw3JecEfBbQRilMXNBBzgvE6gAlAgMBAAECgYAZ1HAZE4tV9b3fEJB0gqZDmDwV52Xp5OrIOT/dyo+ZaocKENpS3Y84jxUlzyqv29MenD0Jw8jypI/01xH0bH2JRkkk0hTEtZlKu22DtlUlrdUjE1AkYNap2UTmqRJAT9nHbSuILBHkCmu8J+4/GLlF0fGoLfcj3jpbmrxRdcD1oQJBAMKdvNn8PAYdyZqN3SfXOwvYRynTTpk4TRAih6bTXFh9OQpQzsbTo4PqeddMaHpl9BlmgQ5PkqkKW4iFYAdLc7cCQQDs4xi0dfpad/ixQIiaK2zj+PbniQy/K7dW2vJfDWW9ZljdW3YFd1hTAnl6BcnTLjth1sepJPH/h8jm0kB9t4MDAkAp/55ksHFHpKAAMYM7eNAuQzNATHCW0jaXN8xsbQptskBgAIZPFBcifbRjiQK2/0+JN7y8GOH+htSwBpBtopPLAkAVST7SEHvjnz4wp0zfFaHeRmPhYFZN417BJkF7Oxf2UfkQQwSy7wt3whQqW9TL7i6V1xfv4gsA3DG4VBJJebEDAkBpXc/gB94ahr3y6k3br+zDKrcyy34E3nX338oj28/+aTGsGWdgsz7uel8/dAeVWd5hmVM9mj2pzO6sEijgp3TK")
	address, _ := impl.AddressVerify(ex_wallet.AddressVerifyParam{
		Address:  "EQBQbT6vP94VRct8wWsqmLk9jVw_G6dLPViQbnndL2dBYOAA@$10563",
		CoinType: "TON",
	})
	log.Println(address)
}

func testGetAllocatedAddress() {
	impl := ex_wallet.NewWalletRestfulProxyImpl(
		"https://waalll.ydykknd.com/rest",
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0Fg6/ZB+h6JL6EyjhXVMt7iAxxQ+BDrmMozBXbtNxyYtYdAcSu9zfDhloCug5shUzLrla4Avwrl5gfa31EGTSE1rKiGOU/5POoCEqYiuB4+k7rvhYUl+D2xUBRDba+u4CFKKcJ2tzIhXfJ8XcNyXnBHwW0EYpTFzQQc4LxOoAJQIDAQAB",
		"MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBALQWDr9kH6HokvoTKOFdUy3uIDHFD4EOuYyjMFdu03HJi1h0BxK73N8OGWgK6DmyFTMuuVrgC/CuXmB9rfUQZNITWsqIY5T/k86gISpiK4Hj6Tuu+FhSX4PbFQFENtr67gIUopwna3MiFd8nxdw3JecEfBbQRilMXNBBzgvE6gAlAgMBAAECgYAZ1HAZE4tV9b3fEJB0gqZDmDwV52Xp5OrIOT/dyo+ZaocKENpS3Y84jxUlzyqv29MenD0Jw8jypI/01xH0bH2JRkkk0hTEtZlKu22DtlUlrdUjE1AkYNap2UTmqRJAT9nHbSuILBHkCmu8J+4/GLlF0fGoLfcj3jpbmrxRdcD1oQJBAMKdvNn8PAYdyZqN3SfXOwvYRynTTpk4TRAih6bTXFh9OQpQzsbTo4PqeddMaHpl9BlmgQ5PkqkKW4iFYAdLc7cCQQDs4xi0dfpad/ixQIiaK2zj+PbniQy/K7dW2vJfDWW9ZljdW3YFd1hTAnl6BcnTLjth1sepJPH/h8jm0kB9t4MDAkAp/55ksHFHpKAAMYM7eNAuQzNATHCW0jaXN8xsbQptskBgAIZPFBcifbRjiQK2/0+JN7y8GOH+htSwBpBtopPLAkAVST7SEHvjnz4wp0zfFaHeRmPhYFZN417BJkF7Oxf2UfkQQwSy7wt3whQqW9TL7i6V1xfv4gsA3DG4VBJJebEDAkBpXc/gB94ahr3y6k3br+zDKrcyy34E3nX338oj28/+aTGsGWdgsz7uel8/dAeVWd5hmVM9mj2pzO6sEijgp3TK")
	address, _ := impl.GetAllocatedAddress(ex_wallet.AddressParam{
		Coin: "TON",
		Size: 1,
	})
	log.Println(address)
}

func testAddressBalanceList() {
	impl := ex_wallet.NewWalletRestfulProxyImpl(
		"https://waalll.ydykknd.com/rest",
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0Fg6/ZB+h6JL6EyjhXVMt7iAxxQ+BDrmMozBXbtNxyYtYdAcSu9zfDhloCug5shUzLrla4Avwrl5gfa31EGTSE1rKiGOU/5POoCEqYiuB4+k7rvhYUl+D2xUBRDba+u4CFKKcJ2tzIhXfJ8XcNyXnBHwW0EYpTFzQQc4LxOoAJQIDAQAB",
		"MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBALQWDr9kH6HokvoTKOFdUy3uIDHFD4EOuYyjMFdu03HJi1h0BxK73N8OGWgK6DmyFTMuuVrgC/CuXmB9rfUQZNITWsqIY5T/k86gISpiK4Hj6Tuu+FhSX4PbFQFENtr67gIUopwna3MiFd8nxdw3JecEfBbQRilMXNBBzgvE6gAlAgMBAAECgYAZ1HAZE4tV9b3fEJB0gqZDmDwV52Xp5OrIOT/dyo+ZaocKENpS3Y84jxUlzyqv29MenD0Jw8jypI/01xH0bH2JRkkk0hTEtZlKu22DtlUlrdUjE1AkYNap2UTmqRJAT9nHbSuILBHkCmu8J+4/GLlF0fGoLfcj3jpbmrxRdcD1oQJBAMKdvNn8PAYdyZqN3SfXOwvYRynTTpk4TRAih6bTXFh9OQpQzsbTo4PqeddMaHpl9BlmgQ5PkqkKW4iFYAdLc7cCQQDs4xi0dfpad/ixQIiaK2zj+PbniQy/K7dW2vJfDWW9ZljdW3YFd1hTAnl6BcnTLjth1sepJPH/h8jm0kB9t4MDAkAp/55ksHFHpKAAMYM7eNAuQzNATHCW0jaXN8xsbQptskBgAIZPFBcifbRjiQK2/0+JN7y8GOH+htSwBpBtopPLAkAVST7SEHvjnz4wp0zfFaHeRmPhYFZN417BJkF7Oxf2UfkQQwSy7wt3whQqW9TL7i6V1xfv4gsA3DG4VBJJebEDAkBpXc/gB94ahr3y6k3br+zDKrcyy34E3nX338oj28/+aTGsGWdgsz7uel8/dAeVWd5hmVM9mj2pzO6sEijgp3TK")
	address, _ := impl.AddressBalanceList(ex_wallet.AddressBalanceParam{
		Address: "EQBQbT6vP94VRct8wWsqmLk9jVw_G6dLPViQbnndL2dBYOAA@$10501]",
		Id:      "123456789",
		Limit:   "1",
	})
	log.Println(address)
}

func testListCoinConfig() {
	impl := ex_wallet.NewWalletRestfulProxyImpl(
		"https://waalll.ydykknd.com/rest",
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0Fg6/ZB+h6JL6EyjhXVMt7iAxxQ+BDrmMozBXbtNxyYtYdAcSu9zfDhloCug5shUzLrla4Avwrl5gfa31EGTSE1rKiGOU/5POoCEqYiuB4+k7rvhYUl+D2xUBRDba+u4CFKKcJ2tzIhXfJ8XcNyXnBHwW0EYpTFzQQc4LxOoAJQIDAQAB",
		"MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBALQWDr9kH6HokvoTKOFdUy3uIDHFD4EOuYyjMFdu03HJi1h0BxK73N8OGWgK6DmyFTMuuVrgC/CuXmB9rfUQZNITWsqIY5T/k86gISpiK4Hj6Tuu+FhSX4PbFQFENtr67gIUopwna3MiFd8nxdw3JecEfBbQRilMXNBBzgvE6gAlAgMBAAECgYAZ1HAZE4tV9b3fEJB0gqZDmDwV52Xp5OrIOT/dyo+ZaocKENpS3Y84jxUlzyqv29MenD0Jw8jypI/01xH0bH2JRkkk0hTEtZlKu22DtlUlrdUjE1AkYNap2UTmqRJAT9nHbSuILBHkCmu8J+4/GLlF0fGoLfcj3jpbmrxRdcD1oQJBAMKdvNn8PAYdyZqN3SfXOwvYRynTTpk4TRAih6bTXFh9OQpQzsbTo4PqeddMaHpl9BlmgQ5PkqkKW4iFYAdLc7cCQQDs4xi0dfpad/ixQIiaK2zj+PbniQy/K7dW2vJfDWW9ZljdW3YFd1hTAnl6BcnTLjth1sepJPH/h8jm0kB9t4MDAkAp/55ksHFHpKAAMYM7eNAuQzNATHCW0jaXN8xsbQptskBgAIZPFBcifbRjiQK2/0+JN7y8GOH+htSwBpBtopPLAkAVST7SEHvjnz4wp0zfFaHeRmPhYFZN417BJkF7Oxf2UfkQQwSy7wt3whQqW9TL7i6V1xfv4gsA3DG4VBJJebEDAkBpXc/gB94ahr3y6k3br+zDKrcyy34E3nX338oj28/+aTGsGWdgsz7uel8/dAeVWd5hmVM9mj2pzO6sEijgp3TK")
	address, _ := impl.ListCoinConfig()
	log.Println(address)
}

// BizID        string `json:"bizId"`        // Business ID
//
//	TxCoin       string `json:"txCoin"`       // Coin type
//	Chain        string `json:"chain"`        // Blockchain chain
//	UserID       string `json:"userId"`       // User ID
//	TxFromWallet string `json:"txFromWallet"` // Wallet 'from' address
//	TxToWallet   string `json:"txToWallet"`   // Wallet 'to' address
//	TxAmount     string `json:"txAmount"`     // Withdrawal amount
//	TxFee        string `json:"txFee"`        // Withdrawal fee
//	Val          string `json:"val"`          // Encrypted parameter
//	TransferType string `json:"transferType"` // Transfer type
func testCreateWithdraw() {
	impl := ex_wallet.NewWalletRestfulProxyImpl(
		"https://waalll.ydykknd.com/rest",
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0Fg6/ZB+h6JL6EyjhXVMt7iAxxQ+BDrmMozBXbtNxyYtYdAcSu9zfDhloCug5shUzLrla4Avwrl5gfa31EGTSE1rKiGOU/5POoCEqYiuB4+k7rvhYUl+D2xUBRDba+u4CFKKcJ2tzIhXfJ8XcNyXnBHwW0EYpTFzQQc4LxOoAJQIDAQAB",
		"MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBALQWDr9kH6HokvoTKOFdUy3uIDHFD4EOuYyjMFdu03HJi1h0BxK73N8OGWgK6DmyFTMuuVrgC/CuXmB9rfUQZNITWsqIY5T/k86gISpiK4Hj6Tuu+FhSX4PbFQFENtr67gIUopwna3MiFd8nxdw3JecEfBbQRilMXNBBzgvE6gAlAgMBAAECgYAZ1HAZE4tV9b3fEJB0gqZDmDwV52Xp5OrIOT/dyo+ZaocKENpS3Y84jxUlzyqv29MenD0Jw8jypI/01xH0bH2JRkkk0hTEtZlKu22DtlUlrdUjE1AkYNap2UTmqRJAT9nHbSuILBHkCmu8J+4/GLlF0fGoLfcj3jpbmrxRdcD1oQJBAMKdvNn8PAYdyZqN3SfXOwvYRynTTpk4TRAih6bTXFh9OQpQzsbTo4PqeddMaHpl9BlmgQ5PkqkKW4iFYAdLc7cCQQDs4xi0dfpad/ixQIiaK2zj+PbniQy/K7dW2vJfDWW9ZljdW3YFd1hTAnl6BcnTLjth1sepJPH/h8jm0kB9t4MDAkAp/55ksHFHpKAAMYM7eNAuQzNATHCW0jaXN8xsbQptskBgAIZPFBcifbRjiQK2/0+JN7y8GOH+htSwBpBtopPLAkAVST7SEHvjnz4wp0zfFaHeRmPhYFZN417BJkF7Oxf2UfkQQwSy7wt3whQqW9TL7i6V1xfv4gsA3DG4VBJJebEDAkBpXc/gB94ahr3y6k3br+zDKrcyy34E3nX338oj28/+aTGsGWdgsz7uel8/dAeVWd5hmVM9mj2pzO6sEijgp3TK")
	address, _ := impl.CreateWithdraw(ex_wallet.WithdrawCreateParam{
		BizID:        "12345667891",
		TxCoin:       "USDT",
		Chain:        "TON",
		UserID:       "8001162",
		TxFromWallet: "0xd0c9a99c13ea9cbbd8a1c1c4788431ad00046a90",
		TxToWallet:   "0xbd93d06c71662a9e499340aef6448101923258d0",
		TxAmount:     "20",
		TxFee:        "0",
	})
	log.Println(address)
}

func testCancelWithdraw() {
	impl := ex_wallet.NewWalletRestfulProxyImpl(
		"https://waalll.ydykknd.com/rest",
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0Fg6/ZB+h6JL6EyjhXVMt7iAxxQ+BDrmMozBXbtNxyYtYdAcSu9zfDhloCug5shUzLrla4Avwrl5gfa31EGTSE1rKiGOU/5POoCEqYiuB4+k7rvhYUl+D2xUBRDba+u4CFKKcJ2tzIhXfJ8XcNyXnBHwW0EYpTFzQQc4LxOoAJQIDAQAB",
		"MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBALQWDr9kH6HokvoTKOFdUy3uIDHFD4EOuYyjMFdu03HJi1h0BxK73N8OGWgK6DmyFTMuuVrgC/CuXmB9rfUQZNITWsqIY5T/k86gISpiK4Hj6Tuu+FhSX4PbFQFENtr67gIUopwna3MiFd8nxdw3JecEfBbQRilMXNBBzgvE6gAlAgMBAAECgYAZ1HAZE4tV9b3fEJB0gqZDmDwV52Xp5OrIOT/dyo+ZaocKENpS3Y84jxUlzyqv29MenD0Jw8jypI/01xH0bH2JRkkk0hTEtZlKu22DtlUlrdUjE1AkYNap2UTmqRJAT9nHbSuILBHkCmu8J+4/GLlF0fGoLfcj3jpbmrxRdcD1oQJBAMKdvNn8PAYdyZqN3SfXOwvYRynTTpk4TRAih6bTXFh9OQpQzsbTo4PqeddMaHpl9BlmgQ5PkqkKW4iFYAdLc7cCQQDs4xi0dfpad/ixQIiaK2zj+PbniQy/K7dW2vJfDWW9ZljdW3YFd1hTAnl6BcnTLjth1sepJPH/h8jm0kB9t4MDAkAp/55ksHFHpKAAMYM7eNAuQzNATHCW0jaXN8xsbQptskBgAIZPFBcifbRjiQK2/0+JN7y8GOH+htSwBpBtopPLAkAVST7SEHvjnz4wp0zfFaHeRmPhYFZN417BJkF7Oxf2UfkQQwSy7wt3whQqW9TL7i6V1xfv4gsA3DG4VBJJebEDAkBpXc/gB94ahr3y6k3br+zDKrcyy34E3nX338oj28/+aTGsGWdgsz7uel8/dAeVWd5hmVM9mj2pzO6sEijgp3TK")
	address, _ := impl.CancelWithdraw(ex_wallet.WithdrawCancelParam{
		BizID: "12345667891",
	})
	log.Println(address)
}
