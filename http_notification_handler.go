package xyz_notif

import (
 "fmt"
 "net/http"
 "io/ioutil"
 "encoding/json"
)
type JsonParams struct {
  Notification NotificationParams
}
type NotificationParams struct {
  Url string `json:"url"`
  Trx Transaction `json:"transaction"`
}

type HttpNotificationHandler struct {}

func (h *HttpNotificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic("panic")
  }
  fmt.Printf("BODY : %s", string(body))

  var jsonParams JsonParams
  err = json.Unmarshal(body, &jsonParams)
  if err != nil {
    panic("panic")
  }

  notificationParams := jsonParams.Notification
  if SendNotification(notificationParams.Url, &notificationParams.Trx) {
    fmt.Fprintf(w, "OK\n")
  }else{
    fmt.Fprintf(w, "NOT OK\n")
  }
}