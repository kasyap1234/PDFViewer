package service


import (
	"context"
pb "github.com/kasyap1234/pdfviewer/proto/api/gen/document"
"github.com/minio/minio-go/v7"
"gorm.io/gorm"

)


type DocumentService struct {
	pb.UnimplementedDocumentManagerServer
	db *gorm.DB 
	minioClient *minio.Client
	bucketName string 
}

// pb.unimplementedDocumentManagerServer
//used when some methods are not implemented 


func NewDocumentService(db*gorm.DB, minioClient *minio.Client, bucketName string) *DocumentService {
	return &DocumentService{
		db: db,
		minioClient: minioClient,
		bucketName: bucketName,
	}
}
func (s *DocumentService)UploadDocument(stream pb.UploadD)