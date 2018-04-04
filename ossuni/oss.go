package ossuni

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"io"
)

// aliOss is aliyun oss
type aliOss struct {
	config *Config
	client *oss.Client
}

func (ao *aliOss) Init(config Config) error {
	if ao.config == nil {
		ao.config = &config
	}
	if ao.client == nil {
		c := config
		client, err := oss.New(c.Endpoint, c.AccessKeyID, c.AccessKeySecret)
		if err != nil {
			return err
		}
		ao.client = client
	}

	return nil
}

func (ao *aliOss) GetAuthToken() {

}

func (ao *aliOss) STSCertificate() {
	//param AliyunSTSRequest
	//reqUrl := "https://sts.aliyuncs.com"
	//params := map[string]string{
	//	"Action":"AssumeRole",
	//	"RoleArn":param.RoleArn,
	//	"RoleSessionName":param.RoleSessionName,
	//	"DurationSeconds":param.DurationSeconds,
	//	"Policy":
	//}
	//assumeRole := types.AssumeRole{}
	//sts.GetSecurityToken(ao.config.AccessKeyID, ao.config.AccessKeySecret, &assumeRole)
}

func (ao *aliOss) GetService() {

}

func (ao *aliOss) PutObject(param ObjectParam, reader io.Reader) error {
	b, err := ao.client.Bucket(param.Bucket)
	if err != nil {
		return err
	}

	return b.PutObject(param.Key, reader)
}

func (ao *aliOss) PutObjectFromFile(param ObjectParam, filePath string) error {
	b, err := ao.client.Bucket(param.Bucket)
	if err != nil {
		return err
	}

	return b.PutObjectFromFile(param.Key, filePath)
}

func (ao *aliOss) InitMultipartUpload(param ObjectParam) (InitMultipartUploadResult, error) {
	var result InitMultipartUploadResult
	b, err := ao.client.Bucket(param.Bucket)
	if err != nil {
		return result, err
	}

	r, err := b.InitiateMultipartUpload(param.Key)
	if err != nil {
		return result, err
	}

	result.Bucket = r.Bucket
	result.Key = r.Key
	result.UploadID = r.UploadID

	return result, nil
}

func (ao *aliOss) AbortMultipartUpload(imur InitMultipartUploadResult) error {
	b, err := ao.client.Bucket(imur.Bucket)
	if err != nil {
		return err
	}

	return b.AbortMultipartUpload(oss.InitiateMultipartUploadResult{
		Bucket:   imur.Bucket,
		Key:      imur.Key,
		UploadID: imur.UploadID,
	})
}

func (ao *aliOss) CompleteMultipartUpload(imur InitMultipartUploadResult, parts []UploadPart) (CompleteMultipartUploadResult, error) {
	var result CompleteMultipartUploadResult
	b, err := ao.client.Bucket(imur.Bucket)
	if err != nil {
		return result, err
	}

	ossParts := make([]oss.UploadPart, len(parts))
	for i, part := range parts {
		ossParts[i].ETag = part.ETag
		ossParts[i].PartNumber = part.PartNumber
	}

	r, err := b.CompleteMultipartUpload(oss.InitiateMultipartUploadResult{
		Bucket:   imur.Bucket,
		Key:      imur.Key,
		UploadID: imur.UploadID,
	}, ossParts)
	if err != nil {
		return result, err
	}

	result.Location = r.Location
	result.ETag = r.ETag
	result.Key = r.Key
	result.Bucket = r.Bucket

	return result, nil
}

func (ao *aliOss) UploadPart(imur InitMultipartUploadResult, reader io.Reader, size int64, partNumber int) (UploadPart, error) {
	var result UploadPart
	b, err := ao.client.Bucket(imur.Bucket)
	if err != nil {
		return result, err
	}

	r, err := b.UploadPart(oss.InitiateMultipartUploadResult{
		Bucket:   imur.Bucket,
		Key:      imur.Key,
		UploadID: imur.UploadID,
	}, reader, size, partNumber)
	if err != nil {
		return result, err
	}

	result.Size = size
	result.ETag = r.ETag
	result.PartNumber = r.PartNumber

	return result, nil
}

func (ao *aliOss) MultiPartUpload() {

}

func (ao *aliOss) DeleteObject(param ObjectParam) error {
	b, err := ao.client.Bucket(param.Bucket)
	if err != nil {
		return err
	}

	return b.DeleteObject(param.Key)
}

func (ao *aliOss) GetAttachmentURL(param ObjectParam) string {
	if ao.config.CdnURL != "" {
		return fmt.Sprintf("%s://%s/%s", ao.config.Scheme, ao.config.CdnURL, param.Key)
	}

	return fmt.Sprintf("%s://%s.%s/%s", ao.config.Scheme, param.Bucket, ao.config.Endpoint, param.Key)
}