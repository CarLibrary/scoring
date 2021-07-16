package serializer

import (
	"CarLibrary/score/model"
	pb "github.com/CarLibrary/proto/score"
)

func BuildScoreResponse(score *model.Score) *pb.ScoreResponse  {
	return &pb.ScoreResponse{
		CarBand:  score.CarBand,
		CarSeries: score.CarSeries,
		Score:     score.Score,
	}
}

