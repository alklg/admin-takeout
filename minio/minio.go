package minio

import (
	"github.com/minio/minio-go"
	"log"
)

var minioClient *minio.Client
var bucketName string

func InitMinio() {
	endpoint := "119.45.145.96:9000"
	accessKeyID := "h6g1IdtIBGIhgrACxapp"
	secretAccessKey := "3xfwtrJsYMVvQ3ILzwfpwFTb2sEPnWrTq1fFi9na"
	useSSL := true

	// accessKeyID is expiry in 2014/6/17 04:00:00 UTC

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)

	if err != nil {
		log.Printf("Init Minio error = %v\n", err)
		return
	}

	log.Printf("this is minioClient %v\n", minioClient)
}

func CreateBucket() {
	bucketName = "adminTakeout"
	location := "cn-north-1"

	exists, err := minioClient.BucketExists(bucketName)

	if exists == true && err != nil {
		log.Println("bucket already exists")
	}

	if !exists {
		err = minioClient.MakeBucket(bucketName, location)

		if err != nil {
			log.Printf("make bucket error = %v\n", err)
		}
	}

	log.Printf("bucket is %v", bucketName)
}
