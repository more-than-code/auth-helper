package helper

import (
	"context"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestAuth(t *testing.T) {
	h, _ := NewHelper(&Config{TtlMinute: 0, TtlHour: 0, TtlDay: 1, Secret: []byte("test123")})

	var u1 = User{Name: "John", Age: 18}

	str, _ := h.Authenticate(context.TODO(), &u1)

	t.Log(str)

	var u2 User
	_ = h.ParseTokenString(str, &u2)

	t.Log(u2)

	if u1 != u2 {
		t.Fail()
	}
}
