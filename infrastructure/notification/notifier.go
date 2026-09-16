package notification
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
 
	"github.com/joho/godotenv"
)
 
// HTTPNotifier mengirim notifikasi via HTTP POST ke endpoint URL_NOTIF.
// Port hasil migrasi dari go_itapp/controller/notification.controller.go
type HTTPNotifier struct {
	client *http.Client
}
 
// NewHTTPNotifier constructor
func NewHTTPNotifier() *HTTPNotifier {
	return &HTTPNotifier{client: &http.Client{}}
}
 
// Send mengirim notifikasi (implementasi dari port Notifier)
func (n *HTTPNotifier) Send(username string, title string, message string) error {
	// Data notifikasi
	notificationData := map[string]interface{}{
		"username": username,
		"fcm":      true,
		"title":    title,
		"message":  message,
		"module":   "IT APPLICATION",
	}
 
	// Konversi ke JSON
	jsonData, err := json.Marshal(notificationData)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}
 
	_ = godotenv.Load() // Mengabaikan error jika file .env absen (sama seperti db.go)
 
	// Buat request HTTP
	req, err := http.NewRequest("POST", os.Getenv("URL_NOTIF"), bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
 
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", "D884B5333413F1225CB067736991C5F4")
 
	// Kirim request
	resp, err := n.client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
 
	// Baca response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
 
	// Tampilkan status dan isi response
	fmt.Println("Notification response status:", resp.Status)
	fmt.Println("Response body:", string(body))
	return nil
}
