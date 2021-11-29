package global

import (
	"crypto/rand"
	"fmt"
)

//  RandomString ===> 隨機生成亂數字串
func RandomString() (pwd string) {

	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	pwd = fmt.Sprintf("%X", b)
	return pwd
}
