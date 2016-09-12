package xyz_notif
import "testing"
import "net/http/httptest"
import "bytes"

func TestHttpNotificationHandler_ServeHTTP(t *testing.T) {
    handler := HttpNotificationHandler{}
    rec := httptest.NewRecorder()
    var b []byte
    rec.Body = bytes.NewBuffer(b)
    handler.ServeHTTP(rec, nil)
    if(rec.Body.String() != "OK\n"){
      t.Errorf("unexpected response: %s", rec.Body.String())
    }
}