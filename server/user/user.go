package user

import (
	"fmt"
	"time"
)

// User type describes a system user like admin, librarian
type User struct {
	UserID    string
	FirstName string
	LastName  string
	Password  string
	Mobile    string
	DOB       time.Time
}

func (u *User) String() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%v", u.UserID, u.FirstName, u.LastName, u.Password, u.Mobile, u.DOB.Format(LayoutISO))
}

// Print prints user in a nice format
func (u *User) Print() {
	fmt.Printf("%s, %s, %s, %s, %s, %v\n", u.UserID, u.FirstName, u.LastName, u.Password, u.Mobile, u.DOB.Format(LayoutISO))
}
