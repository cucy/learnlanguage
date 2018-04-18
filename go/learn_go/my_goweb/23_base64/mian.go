package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := `"Love is but a song to sing Fear's the way we die You can make the mountains ring Or make the angels cry Though the bird is on the wing And you may not know why Come on people now Smile on your brother Everybody get together Try to love one another Right now"`

	s64 := base64.StdEncoding.EncodeToString([]byte(s)) // 加密
	fmt.Println(s64)                                    // 密文

	bs, err := base64.StdEncoding.DecodeString(s64) // 解密
	if err != nil {
		log.Fatalln("I'm giving her all she's got Captain!", err)
	}
	fmt.Println(string(bs))
}

// If for some reason you need double quotes in your cookies you can encode your cookie value to base64,
// and then decode the cookie value to get back your original double quoted string
// 如果因为某些原因你需要双引号在你的饼干你可以你的cookie值编码为base64，
// 然后解码Cookie值以返回原来的双引号字符串。
