package entity

type Role uint8

const (
	UserRole=iota+1
	AdminRole
)

func (r Role) strings() string{
	switch r{
	case UserRole:
		return  "user"
	case AdminRole:
		return "admin"
		
	}
	return ""
	
	
} 
