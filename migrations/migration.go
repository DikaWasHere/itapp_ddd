package migration
 
import (
	"fmt"
	
 
	"itapp_ddd/infrastructure/persistence"
)
 
func RunMigration() error {
	if err := persistence.DB.Exec("CREATE SCHEMA IF NOT EXISTS trs").Error; err != nil {
		return fmt.Errorf("failed to create schema trs: %w", err)
	}
	if err := persistence.DB.Exec("CREATE SCHEMA IF NOT EXISTS mst").Error; err != nil {
		return fmt.Errorf("failed to create schema mst: %w", err)
	}
 
	if err := persistence.DB.AutoMigrate(
		&persistence.DocumentCrModel{},    
		&persistence.ReviewModel{},        
		&persistence.ApprovalModel{},      
		&persistence.PICModel{},           
		&persistence.DocumentCrLogModel{}, 
		&persistence.AttachmentModel{},    
	); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
 
	fmt.Println("Database Migrated")
	return nil
}

