package ossuni

import (
	"io"

	"fmt"

	"github.com/lily-lee/qcloud-cos-go-sdk/cos"
)

// qcloudCos is qcloud cos
type qcloudCos struct {
	client *cos.Client
	config *Config
}

func (qc *qcloudCos) Init(config Config) error {
	if qc.config == nil {
		qc.config = &config
	}
	if qc.client == nil {
		c := config
		client, err := cos.NewClient(c.AppID, c.SecretID, c.SecretKey, c.Scheme, c.AuthExpired)
		if err != nil {
			return err
		}
		qc.client = client
	}

	return nil

}

func (qc *qcloudCos) GetAuthToken() {

}

// TODO
func (qc *qcloudCos) STSCertificate(param STSParam) (interface{}, error) {
	return nil, nil
}

func (qc *qcloudCos) GetService() {

}

func (qc *qcloudCos) PutObject(param ObjectParam, reader io.Reader) error {
	b, err := qc.client.NewBucket(param.Bucket, param.Region)
	if err != nil {
		return err
	}
	return b.PutObject(param.Key, reader, nil)
}

func (qc *qcloudCos) PutObjectFromFile(param ObjectParam, filepath string) error {
	b, err := qc.client.NewBucket(param.Bucket, param.Region)
	if err != nil {
		return err
	}
	return b.PutObjectFromFile(param.Key, filepath, nil)
}

func (qc *qcloudCos) InitMultipartUpload(param ObjectParam) (InitMultipartUploadResult, error) {
	var result InitMultipartUploadResult
	b, err := qc.client.NewBucket(param.Bucket, param.Region)
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
	result.Region = param.Region

	return result, nil
}

func (qc *qcloudCos) AbortMultipartUpload(imur InitMultipartUploadResult) error {
	b, err := qc.client.NewBucket(imur.Bucket, imur.Region)
	if err != nil {
		return err
	}

	return b.AbortMultipartUpload(cos.InitiateMultipartUploadResult{
		Bucket:   imur.Bucket,
		Key:      imur.Key,
		UploadID: imur.UploadID,
	})
}

func (qc *qcloudCos) CompleteMultipartUpload(imur InitMultipartUploadResult, parts []UploadPart) (CompleteMultipartUploadResult, error) {
	var result CompleteMultipartUploadResult
	b, err := qc.client.NewBucket(imur.Bucket, imur.Region)
	if err != nil {
		return result, err
	}

	cosParts := make([]cos.Part, len(parts))
	for i, part := range parts {
		cosParts[i].ETag = part.ETag
		cosParts[i].PartNumber = part.PartNumber
		cosParts[i].Size = part.Size
		cosParts[i].LastModified = part.LastModified
	}

	r, err := b.CompleteMultipartUpload(cos.InitiateMultipartUploadResult{
		Bucket:   imur.Bucket,
		Key:      imur.Key,
		UploadID: imur.UploadID,
	}, cosParts)
	if err != nil {
		return result, err
	}

	result.Location = r.Location.String()
	result.ETag = r.ETag
	result.Key = r.Key
	result.Bucket = r.Bucket

	return result, nil
}

func (qc *qcloudCos) UploadPart(imur InitMultipartUploadResult, reader io.Reader, size int64, partNumber int) (UploadPart, error) {
	var result UploadPart
	b, err := qc.client.NewBucket(imur.Bucket, imur.Region)
	if err != nil {
		return result, err
	}

	r, err := b.UploadPart(cos.InitiateMultipartUploadResult{
		Bucket:   imur.Bucket,
		Key:      imur.Key,
		UploadID: imur.UploadID,
	}, reader, size, partNumber)
	if err != nil {
		return result, err
	}

	result.PartNumber = partNumber
	result.ETag = r.ETag
	result.Size = r.Size
	result.LastModified = r.LastModified

	return result, nil
}

func (qc *qcloudCos) DeleteObject(param ObjectParam) error {
	b, err := qc.client.NewBucket(param.Bucket, param.Region)
	if err != nil {
		return err
	}

	return b.DeleteObject(param.Key)
}

func (qc *qcloudCos) MultiPartUpload() {

}

func (qc *qcloudCos) GetAttachmentURL(param ObjectParam) string {
	return fmt.Sprintf("%s://%s.cos.%s.myqcloud.com/%s", qc.config.Scheme, param.Bucket, param.Region, param.Key)
}
