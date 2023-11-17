package global

import (
	"crypto/md5"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"os"
)

var (
	DB *dynamo.DB
)

var JwtLoginKey = MD5("sam-app-test-2023")
var Issuer = "zz-test"

func InitDynamodb() {
	os.Setenv("AWS_ACCESS_KEY_ID", "dummy1")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy2")
	os.Setenv("AWS_SESSION_TOKEN", "dummy3")
	creds := credentials.NewEnvCredentials()
	creds.Get()
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Endpoint:    aws.String("http://localhost:8000"),
		Credentials: creds,
	}))
	DB = dynamo.New(sess)
	if DB == nil {
		panic("connect db wrong")
	}
	fmt.Println("InitDynamoDB...")
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func ReverseSlice(s []float64) []float64 {
	// 获取切片的长度
	n := len(s)
	// 创建一个新切片，用于存储倒置后的数据
	reversed := make([]float64, n)
	// 倒置切片
	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		reversed[i] = s[j]
	}
	return reversed
}
