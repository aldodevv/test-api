package structs

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User adalah representasi tabel users di database.
type User struct {
	// gorm.Model secara otomatis memberikan field ID, CreatedAt, UpdatedAt, dan DeletedAt
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"password" binding:"required,min=6"`
}

// BeforeSave adalah "Hook" dari GORM. Fungsi ini otomatis dipanggil sebelum data User disimpan ke DB.
// Kita menggunakannya untuk men-hash password demi keamanan, menggunakan bcrypt.
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// Jika password sedang diperbarui/baru, maka hash text aslinya
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return
}
