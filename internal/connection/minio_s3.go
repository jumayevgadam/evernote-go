package connection

import (
	"github.com/minio/minio-go/v7"
)

type MinioClient struct {
	minio *minio.Client
}

func NewMinioClient(minio *minio.Client) *MinioClient {
	return &MinioClient{minio: minio}
}
