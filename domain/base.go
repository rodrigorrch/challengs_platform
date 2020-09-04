package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
	// DeletedAt time.Time `json:"deleted_at" valid:"-" sql:"index"`
}

// func (base *Base) BeforeCreate(scope *gorm.Scope) error {
// 	err := scope.SetColumn("CreatedAt", time.Now())

// 	if err != nil {
// 		log.Fatalf("Error during obj creating %v", err)
// 	}

// 	err = scope.SetColumn("ID", uuid.NewV4().String())

// 	if err != nil {
// 		log.Fatalf("Error during obj creating %v", err)
// 	}

// 	return nil
// }
