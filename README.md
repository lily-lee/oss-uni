# ossuni

阿里云与腾讯云存储的统一接口，

为了方便在两个云之间做切换

# Install
```bash
go get github.com/lily-lee/oss-uni/ossuni

```

# Usage
```go
package main

import (
    "github.com/lily-lee/oss-uni/ossuni"
    "log"
)

// example
func main() {
	ossType := ossuni.OssType(ossuni.ALIYUN)
	client := ossuni.NewClient(ossType)
	client.Init(ossuni.Config{
		OssType:     ossType,
		Scheme:      "https",
		AppID:       "QCloud APPID",
		SecretID:    "QCloud SecretID",
		SecretKey:   "QCloud SecretKey",
		AuthExpired: 600,

		Endpoint:        "Aliyun OSS endpoint",
		AccessKeyID:     "Aliyun OSS AccessKeyID",
		AccessKeySecret: "Aliyun OSS AccessKeySecret",
	})
	param := ossuni.ObjectParam{
		Bucket: "Your Bucket Name",
		Key:    "readme",
		Region: "Your Bucket Region", // only for qcloud cos
	}
	e := client.PutObjectFromFile(param, "../README.md")
	log.Println(e)
}

```
# LICENSE
[MIT License](https://github.com/lily-lee/oss-uni/blob/master/LICENSE)