package main
import (
	"crypto/tls"
	"fmt"
  "github.com/barockok/xyz_notif"
  "github.com/barockok/xyz_notif/xyz"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type NotificationServiceHandler struct {}
func (this *NotificationServiceHandler) Send(url string, trx *xyz.Transaction ) (result string, err error) {

  xyzNotifTrx := xyz_notif.Transaction{}

  xyzNotifTrx.Id              = trx.ID
  xyzNotifTrx.GrossAmount     = *trx.GrossAmount
  xyzNotifTrx.MerchantId      = *trx.MerchantID

  if xyz_notif.SendNotification(url, &xyzNotifTrx) {
    return "OK", nil
  }else{
    return "NOT OK", nil
  }
}

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TServerTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		if cert, err := tls.LoadX509KeyPair("keys/server.crt", "keys/server.key"); err == nil {
			cfg.Certificates = append(cfg.Certificates, cert)
		} else {
			return err
		}
		transport, err = thrift.NewTSSLServerSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTServerSocket(addr)
	}

	if err != nil {
		return err
	}
	fmt.Printf("%T\n", transport)
	handler := &NotificationServiceHandler{}
	processor := xyz.NewNotificationServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}