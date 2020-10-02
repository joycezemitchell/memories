package server

import (
	"context"
	"fmt"

	config "allyapps.com/memories/config"
	models "allyapps.com/memories/models"
	memoriespb "allyapps.com/memories/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson"
)

// Server - gRPC memories server
type Server struct {
}

// UploadVideo Method
func (*Server) UploadVideo(ctx context.Context, req *memoriespb.UploadVideoRequest) (*memoriespb.UploadVideoResponse, error) {

	// Update database
	data := models.Videos{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Video:       req.GetVideo(),
		Thumbnail:   req.GetThumbnail(),
		CreatedAt:   req.GetCreatedAt(),
		UpdatedAt:   req.GetUpdatedAt(),
	}

	res, _ := config.Database.Collection("videos").InsertOne(context.Background(), data)
	oid, _ := res.InsertedID.(primitive.ObjectID)

	fmt.Println(oid)

	return &memoriespb.UploadVideoResponse{
		Memories: &memoriespb.Memories{
			Id:          oid.Hex(),
			Title:       req.GetTitle(),
			Description: req.GetDescription(),
			Video:       req.GetVideo(),
			Thumbnail:   req.GetThumbnail(),
			CreatedAt:   req.GetCreatedAt(),
			UpdatedAt:   req.GetUpdatedAt(),
		},
	}, nil

	// return nil, nil

}

// ListVideos Videos
func (*Server) ListVideos(req *memoriespb.ListVideosRequest, stream memoriespb.MemoriesService_ListVideosServer) error {

	// Display random
	// Reff: https://stackoverflow.com/questions/56082193/mongo-random-records-on-golang

	pipeline := []bson.D{bson.D{{"$sample", bson.D{{"size", 500}}}}}

	// cur, err := config.Database.Collection("videos").find(context.Background(), primitive.D{{}})
	cur, err := config.Database.Collection("videos").Aggregate(context.Background(), pipeline)

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &models.Videos{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)

		}

		stream.Send(&memoriespb.ListVideosResponse{
			Memories: &memoriespb.Memories{
				Id:          data.ID.Hex(),
				Title:       data.Title,
				Description: data.Description,
				Video:       data.Video,
				Thumbnail:   data.Thumbnail,
				CreatedAt:   data.CreatedAt,
				UpdatedAt:   data.UpdatedAt,
			},
		})

	}

	return nil

}
