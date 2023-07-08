package types_test

import (
	"encoding/json"
	"testing"

	"github.com/thirathawat/types"
	"go.mongodb.org/mongo-driver/bson"
)

func TestHashString(t *testing.T) {
	h := types.Hash("password")

	result := h.String()

	if result == "password" {
		t.Errorf("String() returned %s, expected hashed value", result)
	}
}

func TestHashCompare(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		h := types.Hash("password")

		result := h.Compare("password")

		if !result {
			t.Errorf("Compare() returned false, expected true")
		}
	})

	t.Run("Wrong", func(t *testing.T) {
		h := types.Hash("password")

		result := h.Compare("wrongpassword")

		if result {
			t.Errorf("Compare() returned true, expected false")
		}
	})
}

func TestHashMarshalJSON(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		h := types.Hash("password")

		result, err := json.Marshal(h)

		if err != nil {
			t.Errorf("MarshalJSON() returned error %s, expected nil", err)
		}

		if string(result) == "password" {
			t.Errorf("MarshalJSON() returned %s, expected hashed value", result)
		}
	})

	t.Run("Struct", func(t *testing.T) {
		h := struct {
			Password types.Hash `json:"password"`
		}{
			Password: "password",
		}

		result, err := json.Marshal(h)

		if err != nil {
			t.Errorf("MarshalJSON() returned error %s, expected nil", err)
		}

		if string(result) == "password" {
			t.Errorf("MarshalJSON() returned %s, expected hashed value", result)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		h := types.Hash("")

		result, err := json.Marshal(h)

		if err != nil {
			t.Errorf("MarshalJSON() returned error %s, expected nil", err)
		}

		if string(result) == "password" {
			t.Errorf("MarshalJSON() returned %s, expected hashed value", result)
		}
	})

	t.Run("Invalid string", func(t *testing.T) {
		h := types.Hash("password")

		result, err := json.Marshal(h)

		if err != nil {
			t.Errorf("MarshalJSON() returned error %s, expected nil", err)
		}

		if string(result) == "password" {
			t.Errorf("MarshalJSON() returned %s, expected hashed value", result)
		}
	})

	t.Run("Pointer", func(t *testing.T) {
		h := types.Hash("password")

		result, err := json.Marshal(&h)

		if err != nil {
			t.Errorf("MarshalJSON() returned error %s, expected nil", err)
		}

		if string(result) == "password" {
			t.Errorf("MarshalJSON() returned %s, expected hashed value", result)
		}
	})
}

func TestHashUnmarshalJSON(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		var h types.Hash

		err := json.Unmarshal([]byte(`"password"`), &h)

		if err != nil {
			t.Errorf("UnmarshalJSON() returned error %s, expected nil", err)
		}

		if !h.Compare("password") {
			t.Errorf("UnmarshalJSON() returned %s, expected hashed value", h)
		}
	})

	t.Run("Wrong", func(t *testing.T) {
		var h types.Hash

		err := json.Unmarshal([]byte(`"password"`), &h)

		if err != nil {
			t.Errorf("UnmarshalJSON() returned error %s, expected nil", err)
		}

		if h.Compare("wrongpassword") {
			t.Errorf("UnmarshalJSON() returned %s, expected hashed value", h)
		}
	})

	t.Run("Invalid string", func(t *testing.T) {
		var h types.Hash

		err := json.Unmarshal([]byte(`password`), &h)

		if err == nil {
			t.Errorf("UnmarshalJSON() returned nil, expected error")
		}

		if h.Compare("password") {
			t.Errorf("UnmarshalJSON() returned %s, expected hashed value", h)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		var h types.Hash

		err := json.Unmarshal([]byte(`""`), &h)

		if err != nil {
			t.Errorf("UnmarshalJSON() returned error %s, expected nil", err)
		}

		if h.Compare("password") {
			t.Errorf("UnmarshalJSON() returned %s, expected hashed value", h)
		}
	})

	t.Run("Struct", func(t *testing.T) {
		var h struct {
			Password types.Hash `json:"password"`
		}

		err := json.Unmarshal([]byte(`{"password": "password"}`), &h)

		if err != nil {
			t.Errorf("UnmarshalJSON() returned error %s, expected nil", err)
		}

		if !h.Password.Compare("password") {
			t.Errorf("UnmarshalJSON() returned %s, expected hashed value", h.Password)
		}
	})

	t.Run("Pointer", func(t *testing.T) {
		var h *types.Hash

		err := json.Unmarshal([]byte(`"password"`), &h)

		if err != nil {
			t.Errorf("UnmarshalJSON() returned error %s, expected nil", err)
		}

		if !h.Compare("password") {
			t.Errorf("UnmarshalJSON() returned %s, expected hashed value", h)
		}
	})
}

func TestHashMarshalBSON(t *testing.T) {
	h := struct {
		Password types.Hash `bson:"password"`
	}{
		Password: "password",
	}

	result, err := bson.Marshal(h)

	if err != nil {
		t.Errorf("MarshalBSON() returned error %s, expected nil", err)
	}

	if string(result) == "password" {
		t.Errorf("MarshalBSON() returned %s, expected hashed value", result)
	}
}

func TestHashUnmarshalBSON(t *testing.T) {
	h := struct {
		Password types.Hash `bson:"password"`
	}{
		Password: "password",
	}

	result, err := bson.Marshal(h)
	if err != nil {
		t.Errorf("MarshalBSON() returned error %s, expected nil", err)
	}

	var h2 struct {
		Password types.Hash `bson:"password"`
	}

	err = bson.Unmarshal(result, &h2)
	if err != nil {
		t.Errorf("UnmarshalBSON() returned error %s, expected nil", err)
	}

	if !h2.Password.Compare("password") {
		t.Errorf("UnmarshalBSON() returned %s, expected hashed value", h2.Password)
	}

	if h2.Password.String() == "password" {
		t.Errorf("UnmarshalBSON() returned %s, expected hashed value", h2.Password)
	}
}
