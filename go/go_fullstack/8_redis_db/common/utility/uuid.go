package utility

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateUUID() string {

	/*	f, err := os.Open("/dev/urandom")
		if err != nil {
			log.Println("Encountered the following error when attempting to generate an UUID: ", err)
			return ""
		}
		b := make([]byte, 16)
		f.Read(b)
		f.Close()
	*/

	// todo bug 版本为了
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b_str := []byte(str)
	b := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 16; i++ {
		b = append(b, byte(r.Intn(len(b_str))))
	}
	//string(result)

	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
