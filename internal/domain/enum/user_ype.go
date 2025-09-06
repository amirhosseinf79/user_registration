package enum

import (
	"database/sql/driver"
	"fmt"
)

// UserType represents the type of user.
type UserType uint

const (
	Unknown UserType = 0
	Admin   UserType = 1
	User    UserType = 2
)

func (m *UserType) IsValid() bool {
	switch *m {
	case Admin, User:
		return true
	default:
		return false
	}
}

func (m *UserType) String() string {
	switch *m {
	case Admin:
		return "Admin"
	case User:
		return "User"
	default:
		return "Unknown"
	}
}

// Value implements driver.Valuer so GORM/DB can store the enum as integer.
func (m *UserType) Value() (driver.Value, error) {
	if !m.IsValid() {
		return nil, fmt.Errorf("invalid userType: %d", m)
	}
	return int64(*m), nil
}

// Scan implements sql.Scanner so GORM/DB can read the integer into the enum.
func (m *UserType) Scan(value any) error {
	if value == nil {
		*m = Unknown
		return nil
	}
	switch v := value.(type) {
	case int:
		*m = UserType(v)
	case int32:
		*m = UserType(v)
	case int64:
		*m = UserType(v)
	case uint64:
		*m = UserType(v)
	case string:
		switch v {
		case "1", "Admin", "admin":
			*m = Admin
		case "2", "User", "user":
			*m = User
		default:
			return fmt.Errorf("cannot scan UserType from string: %s", v)
		}
	default:
		return fmt.Errorf("unsupported scan type for UserType: %T", value)
	}
	if !m.IsValid() {
		return fmt.Errorf("scanned invalid UserType value: %d", *m)
	}
	return nil
}
