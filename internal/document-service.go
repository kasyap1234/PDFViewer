package service

import (
	
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/kasyap1234/pdfviewer/proto/api/gen/document"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"
	"gorm.io/gorm"
)

type Document struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Path string `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

}
type DocumentService struct {
	pb.UnimplementedDocumentManagerServer
	db *gorm.DB 
	minioClient *minio.Client
	bucketName string 
}

// pb.unimplementedDocumentManagerServer
//used when some methods are not implemented 

func NewminioClient(endpoint string, accessKeyID string, secretAccessKey string, useSSL bool) (*minio.Client, error) {
	minioClient,err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL, 
	})
	if err !=nil {
		return nil,err; 

	}
	return minioClient,nil;

	
}

func NewDocumentService(db*gorm.DB, minioClient *minio.Client, bucketName string) *DocumentService {
	return &DocumentService{
		db: db,
		minioClient: minioClient,
		bucketName: bucketName,
	}
}
func makeBucket(minioClient *minio.Client,bucketName string , ctx context.Context){
	exists,err :=minioClient.BucketExists(ctx,bucketName); 
	if err !=nil {
		return
	}
	if !exists {
		
		exists :=minioClient.MakeBucket(ctx,bucketName,minio.MakeBucketOptions{})
		log.Printf("Created bucket %s\n", bucketName)
		fmt.Print(exists)


}
}

func (s *DocumentService)UploadDocument(stream pb.DocumentManager_UploadDocumentServer)error {
	var fileBytes []byte 
	var fileName string 
	for {
		req, err := stream.Recv()
		if err ==io.EOF{
			break
		}
		if err !=nil {
			return status.Errorf(codes.Internal, "failed to receive file chunk: %v", err)
		}
		fileName=req.GetMetadata().Title
		fileBytes = append(fileBytes, req.GetChunk()...)
	}

	// upload file to minio ; 
	_,err :=s.minioClient.PutObject(context.Background(),s.bucketName,fileName,io.NopCloser(bytes.NewReader(fileBytes)),int64(len(fileBytes)),minio.PutObjectOptions{})
 if err !=nil {
	return status.Errorf(codes.Internal, "failed to upload file to minio: %v", err)

 }

 // save metadata to the database 
 doc := &Document{
	Name: fileName,
	Path: s.bucketName+"/"+fileName,

 }
 if err :=s.db.Create(&doc).Error; err !=nil {
	return status.Errorf(codes.Internal, "failed to save metadata to the database: %v", err)

}



}