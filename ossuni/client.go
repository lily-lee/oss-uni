package ossuni

func NewClient(ossType OssType) Clienter {
	if ossType == ALIYUN {
		return new(aliOss)
	} else {
		return new(qcloudCos)
	}
}

type Clienter interface {
	Init(config Config) error // first
	GetAuthToken()            // TODO
	STSCertificate(param STSParam) (interface{}, error)
	PutObject(param ObjectParam, reader io.Reader) error
	PutObjectFromFile(param ObjectParam, filePath string) error
	InitMultipartUpload(param ObjectParam) (InitMultipartUploadResult, error)
	AbortMultipartUpload(imur InitMultipartUploadResult) error
	CompleteMultipartUpload(imur InitMultipartUploadResult, parts []UploadPart) (CompleteMultipartUploadResult, error)
	UploadPart(imur InitMultipartUploadResult, reader io.Reader, size int64, partNumber int) (UploadPart, error)
	MultiPartUpload() // TODO
	GetAttachmentURL(param ObjectParam) string
	DeleteObject(param ObjectParam) error
}
