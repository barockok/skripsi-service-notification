package main
import (
 "net/http"
 "github.com/barockok/xyz_notif"
 "log"
 "fmt"
 "os"
)

func main() {
    port := os.Getenv("PORT")
    if(port == ""){
      port = "9090"
    }
    http.Handle("/notifications", &xyz_notif.HttpNotificationHandler{})
    fmt.Println("Starting xyz_http\n")
    log.Fatal(http.ListenAndServe(":" + port, nil))
}
