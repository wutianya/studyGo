/*
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
	endpoint := ""
	accessKeyId := ""
	accessKeySecret := ""
	bucketName := ""

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
*/
/*
// oss列举所有的存储空间
package main

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {
	endpoint := ""
	accessKeyId := ""
	accessKeySecret := ""
	cli, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}

	marker := ""
	for {
		lsRes, err := cli.ListBuckets(oss.Marker(marker))
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		for _, bucker := range lsRes.Buckets {
			fmt.Println("Bucket: ", bucker.Name)
		}
		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
}
*/
/*
// 删除文件, 获取指定前缀的文件
package main

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func HandleError(err error) {
	fmt.Println("error: ", err)
	os.Exit(-1)
}
func main() {
	endpoint := ""
	accessKeyId := ""
	accessKeySecret := ""

	cli, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		HandleError(err)
	}

	bucketName := ""
	// 获取存储空间
	bucket, err := cli.Bucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	// 返回删除成功的文件
	// delRes, err := bucket.DeleteObjects([]string{"AliyunSLBHealthCheckLogs/20180310/10","AliyunSLBHealthCheckLogs/20180312/11"})
	// if err != nil {
	// 	HandleError(err)
	// }
	// fmt.Println("Deleted Objects:", delRes.DeletedObjects)
	// 不返回删除的结果
	// _, err = bucket.DeleteObjects([]string{"AliyunSLBHealthCheckLogs/20180312/14","AliyunSLBHealthCheckLogs/20180312/17"})
	// oss.DeleteObjectsQuiet(true)
	// if err != nil {
	// 	HandleError(err)
	// }
	// 列举包含指定前缀的文件,默认列举100个文件
	lsRes, err := bucket.ListObjects(oss.Prefix("AliyunSLBHealthCheckLogs/2019"),oss.MaxKeys(50))
	if err != nil {
		HandleError(err)
	}
	var data []string
	// fmt.Println("Objects: ", lsRes.Objects)
	for _, object := range lsRes.Objects {
		data = append(data,object.Key)
		// fmt.Println("object:", object.Key)
	}
	// fmt.Println(len(data))
	// 返回删除成功的文件
	delRes, err := bucket.DeleteObjects(data)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Deleted Objects:", delRes.DeletedObjects)
}
*/
/*
// stable version
package main

import (
	"fmt"
	"os"
	"errors"
	"flag"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var h bool
var d,s string
func init() {
	flag.BoolVar(&h, "h", false, "Show help.")
	flag.StringVar(&d, "d", "", "File path after uploading to Aliyun OSS.")
	flag.StringVar(&s, "s", "", "Specify the upload file path")
	flag.Usage = Usage
}
func Usage() {
	fmt.Fprintf(os.Stdout, `ossCli version: ossCli/v1.0.0
First add four environment variables:
	endpoint,accessKeyId,accessKeySecret,bucketName
Usage: ossCli [-d filepath] [-s filepath]
Optios:
`)
	flag.PrintDefaults()
}
// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		fmt.Printf("Transfer Started, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferDataEvent:
		fmt.Printf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}

func HandleError(err error) {
	fmt.Println("error: ", err)
	os.Exit(1)
}
func checkSecret(secret []string)(err error) {
	for _, v := range secret {
		if v == "" {
			// str := fmt.Sprintf(`The %v value in null !`, v)
			err = errors.New("The value of the injected environment variable is empty")
			return 
		}
	}
	return 
}
func main() {
	flag.Parse()
	if h || len(os.Args[1:]) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	if len(os.Args[1:]) != 4 {
		fmt.Println("Must specify -d \"xxx\" -s  \"xxx\" args!")
		os.Exit(0)
	}
	
	endpoint := os.Getenv("endpoint")
	accessKeyId := os.Getenv("accessKeyId")
	accessKeySecret := os.Getenv("accessKeySecret")
	bucketName := os.Getenv("bucketName")
	err := checkSecret([]string{endpoint,accessKeyId,accessKeySecret,bucketName})
	if err != nil {
		HandleError(err)
	}

	cli, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		HandleError(err)
	}

	// 获取存储空间
	bucket, err := cli.Bucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	err = bucket.PutObjectFromFile(d, s, oss.Progress(&OssProgressListener{}))
	if err != nil {
		HandleError(err)
	}
}
*/

// 断点续传
package main

import (
	"fmt"
	"os"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)
// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		fmt.Printf("Transfer Started, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferDataEvent:
		fmt.Printf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}

func main() {
	endpoint := os.Getenv("endpoint")
	accessKeyId := os.Getenv("accessKeyId")
	accessKeySecret := os.Getenv("accessKeySecret")
	bucketName := os.Getenv("bucketName")

	cli, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bucket, err := cli.Bucket(bucketName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = bucket.UploadFile("google/k8s/v1.14.3/office 2013.zip","F:\\software\\office 2013.zip", 100000*1024,
							oss.Routines(3),oss.Checkpoint(true, ""),oss.Progress(&OssProgressListener{}))
	                        //  google/k8s/v1.14.3
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}