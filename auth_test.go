package helper

import (
	"context"
	"encoding/json"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestAuth(t *testing.T) {
	h, _ := NewHelper(&Config{ttl: TTL{minute: 0, hour: 0, day: 1}, secret: []byte("test123")})

	var u1 = User{Name: "John", Age: 18}

	str, _ := h.Authenticate(context.TODO(), &u1)

	t.Log(str)

	p, _ := h.ParseTokenString(str)

	var u2 User
	json.Unmarshal([]byte(p.(string)), &u2)

	t.Log(u2)

	if u1 != u2 {
		t.Fail()
	}
}