package ossuni

type OssType string

const (
	ALIYUN  OssType = "aliyun"
	TENCENT OssType = "qcloud"
)

type Config struct {
	OssType OssType `json:"ossType"`
	Scheme  string  `json:"scheme"`

	// qcloud
	AppID       string `json:"appId"`       // qcloud appId
	SecretID    string `json:"secretId"`    // qcloud secretId
	SecretKey   string `json:"secretKey"`   // qcloud secretKey
	AuthExpired int64  `json:"authExpired"` // qcloud authExpired

	// aliyun
	AccessKeyID     string `json:"accessKeyID"`     //  aliyun accessKeyID
	AccessKeySecret string `json:"accessKeySecret"` // aliyun accessKeySecret
	Endpoint        string `json:"endpoint"`        // aliyun

	CdnURL string `json:"cdnUrl"`
}
