package types

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

type Hash string

// MarshalJSON implements the json.Marshaler interface.
func (h Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (h *Hash) UnmarshalJSON(data []byte) error {
	var hashed string
	if err := json.Unmarshal(data, &hashed); err != nil {
		return err
	}

	*h = Hash(hashed)
	return nil
}

// String implements the fmt.Stringer interface.
func (h Hash) String() string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(string(h)), bcrypt.MinCost)
	return string(hashed)
}

// Compare compares the hashed value with a plain string.
func (h Hash) Compare(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h.String()), []byte(plain))
	return err == nil
}
