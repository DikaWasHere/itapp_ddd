package notification
 
// Notifier adalah port (interface) untuk mengirim notifikasi.
// Implementasinya ada di infrastructure layer (HTTP client).
type Notifier interface {
	Send(username string, title string, message string) error
}
 
// Service adalah application service untuk Notification
type Service struct {
	notifier Notifier
	logger   Logger
}
 
// NewService constructor
func NewService(notifier Notifier, logger Logger) *Service {
	return &Service{notifier: notifier, logger: logger}
}
 
// Send mengirim notifikasi ke satu user (username = NIK)
func (s *Service) Send(username string, title string, message string) error {
	if username == "" {
		s.logger.Warn("Send skipped: username is empty")
		return nil
	}
 
	if err := s.notifier.Send(username, title, message); err != nil {
		s.logger.Error("Failed to send notification to %s", err, username)
		return err
	}
 
	s.logger.Info("Notification sent to %s | title: %s", username, title)
	return nil
}