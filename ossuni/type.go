package ossuni

import (
	"time"
)

type InitMultipartUploadResult struct {
	Bucket   string `xml:"Bucket"`   // Bucket名称
	Region   string `xml:"region"`   // Region
	Key      string `xml:"Key"`      // 上传Object名称
	UploadID string `xml:"UploadId"` // 生成的UploadId
}

type UploadPart struct {
	PartNumber   int       `xml:"PartNumber"` // Part编号
	ETag         string    `xml:"ETag"`       // ETag缓存码
	Size         int64     `xml:"Size"`
	LastModified time.Time `xml:"LastModified"`
}

type CompleteMultipartUploadResult struct {
	Location string `xml:"Location"` // Object的URL
	Bucket   string `xml:"Bucket"`   // Bucket名称
	ETag     string `xml:"ETag"`     // Object的ETag
	Key      string `xml:"Key"`      // Object的名字
}

type ObjectParam struct {
	Bucket string `json:"bucket"`
	Region string `json:"region"`
	Key    string `json:"key"`
}

type AliyunSTSRequest struct {
	Action          string
	RoleArn         string
	RoleSessionName string
	Policy          string
	DurationSeconds int

	// public param
	Format           string // JSON
	Version          string // 2015-04-01
	AccessKeyId      string
	Signature        string
	SignatureMethod  string // HMAC-SHA1
	SignatureVersion string // 1.0
	SignatureNonce   string // 随机数
	Timestamp        string // 2012-06-01T12:00:00Z
}

type Policy struct {
	Version    string
	Statements []Statement
}

type Statement struct {
	Effect    string
	Action    []string
	Resource  []string
	Condition string
}

type STSResult struct {
	RequestId       string
	Credentials     Credentials
	AssumedRoleUser AssumedRoleUser
}

type Credentials struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
	Expiration      string
}

type AssumedRoleUser struct {
	Arn               string
	AssumedRoleUserId string
}

type STSParam struct {
	SubAccessKeyId     string
	SubAccessKeySecret string
	RoleArn            string
	RoleSessionName    string
	ExpiredTime        uint
}
