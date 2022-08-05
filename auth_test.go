package helper

import (
	"encoding/json"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestAuth(t *testing.T) {
	h, _ := NewHelper(&Config{TtlMinute: 0, TtlHour: 0, TtlDay: 1, Secret: []byte("test123")})

	var u1 = User{Name: "John", Age: 18}

	b, _ := json.Marshal(u1)
	str1 := string(b)
	tokenStr, _ := h.Authenticate(str1)

	str2, _ := h.ParseTokenString(tokenStr)

	t.Log(str1)
	t.Log(str2)

	if str1 != str2 {
		t.Fail()
	}
}
