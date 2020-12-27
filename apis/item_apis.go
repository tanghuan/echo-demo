package apis

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func init() {
	fmt.Println("package apis init aliyun oss client ......")
	endpoint := "http://dev.oss-cn-chengdu.aliyuncs.com"
	accessKeyID := "asdf"
	accessKeySecret := "asdf"
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret, oss.UseCname(true))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println(client)
}
