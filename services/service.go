package services

import (
	"log"

	"github.com/temur-shamshidinov/Api_gateway/genproto/content_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Services () content_service.ContentServiceClient {

	conn, err := grpc.NewClient("localhost:8000",grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println(err)
		return nil
	}

	contentService  := content_service.NewContentServiceClient(conn)

	return contentService
}