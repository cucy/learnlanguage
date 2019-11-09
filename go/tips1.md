
# for循环引用问题


```go
package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func parse_stu() {
	m := make(map[string]*Student)

	stu := []Student{
		{Name: "su", Age: 100},
		{Name: "li", Age: 12},
		{Name: "mo", Age: 16},
	}

	for _, s := range stu {
		fmt.Printf("s的地址: %p\n", &s)
		m[s.Name] = &s // 有bug

		/*  重新赋值才会申请新的内存 */
		//   s1:=s
		// m[s1.Name] = &s1
	}
	fmt.Println()
	fmt.Printf("%#v\n\n", m)
	for k, v := range m {
		fmt.Printf("k,v的地址: %p  %p\n", &k, &v)
		fmt.Println(k, v)
	}

}
func main() {
	parse_stu()
}



/*

出现此原因是在语句块, 只第一次申请内存, 而后再次循环只会赋新值, 而不会申请新的地址, 所有使用引用的时候要注意

s的地址: 0xc000050420
s的地址: 0xc000050420
s的地址: 0xc000050420

map[string]*main.Student{"li":(*main.Student)(0xc000050420), "mo":(*main.Student)(0xc000050420), "su":(*main.Student)(0xc000050420)}

k,v的地址: 0xc0000461f0  0xc000086020
li &{mo 16}
k,v的地址: 0xc0000461f0  0xc000086020
mo &{mo 16}
k,v的地址: 0xc0000461f0  0xc000086020
su &{mo 16}


*/

```


## 生成证书

```
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"os"
)

func main() {
	/*
		name := os.Args[1]
		user := os.Args[2]*/

	name := "client"
	user := "user-name"

	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	keyDer := x509.MarshalPKCS1PrivateKey(key)
	keyBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keyDer,
	}
	keyFile, err := os.Create(name + "-key.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(keyFile, &keyBlock)
	keyFile.Close()

	commonName := user
	// You may want to update these too
	emailAddress := "someone@myco.com"

	org := "My Co, Inc."
	orgUnit := "Widget Farmers"
	city := "Seattle"
	state := "WA"
	country := "US"

	subject := pkix.Name{
		CommonName:         commonName,
		Country:            []string{country},
		Locality:           []string{city},
		Organization:       []string{org},
		OrganizationalUnit: []string{orgUnit},
		Province:           []string{state},
	}

	asn1, err := asn1.Marshal(subject.ToRDNSequence())
	if err != nil {
		panic(err)
	}
	csr := x509.CertificateRequest{
		RawSubject:         asn1,
		EmailAddresses:     []string{emailAddress},
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	bytes, err := x509.CreateCertificateRequest(rand.Reader, &csr, key)
	if err != nil {
		panic(err)
	}
	csrFile, err := os.Create(name + ".csr")
	if err != nil {
		panic(err)
	}

	pem.Encode(csrFile, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: bytes})
	csrFile.Close()
}

```

