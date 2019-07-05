package main

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)
func HandleError(err error) {
	fmt.Println("Error: ", err)
	os.Exit(-1)
}

func main() {
	fmt.Println("OSS Go SDK Version: ", oss.Version)
	endpoint := "oss-cn-beijing.aliyuncs.com"
	accessKeyId := "LTAIteODziT24mn6"
	accessKeySecret := "VnJrHR7Sgy4IvaFSWIxvuKV7Y2StXo"
	bucketName := "smpt-public"

	cli, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	bucket, err := cli.Bucket(bucketName)

	if err != nil {
		HandleError(err)
	}

	// list file
	marker := ""
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			HandleError(err)
		}
		//  打印列举文件，默认情况下一次返回100条记录
		for _, object := range lsRes.Objects {
			fmt.Println("Bucket: ", object.Key)
		}
		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		}else {
			break
		}
	}
}