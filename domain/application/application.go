package application

import "time"
type Application struct {
	ApplicationID   uint
	ApplicationName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time}
 
type SSOApplication struct {
	ID     uint
	Module string
}
 
type GetListApplicationRequest struct {
	SessionToken string
}