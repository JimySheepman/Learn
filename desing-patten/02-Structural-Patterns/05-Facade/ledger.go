// Complex subsystem parts
package main

import "fmt"

type ledger struct {
}

func (s *ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for account %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}
