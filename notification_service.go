package xyz_notif
import "fmt"


type Transaction struct {
  Id string
  TransactionTime string
  BankMid string
  GrossAmount int32
  CustomerName string
  CustomerEmail string
  Status string
  MerchantId string
}

func SendNotification(url string, trx *Transaction) bool {
  fmt.Printf("SEND NOTIFICATION to:  %s ; trx-id : %s\n", url, trx.Id)
  return true
}
