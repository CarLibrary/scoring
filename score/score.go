package score

import (
	"CarLibrary/score/model"
	"CarLibrary/score/serializer"
	"context"
	pb "github.com/CarLibrary/proto/score"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
type ScoreServiceServer struct {
	pb.UnsafeScoreServiceServer
}

//评分
func (s *ScoreServiceServer)MakeScore(ctx context.Context, score *pb.ScoreRequest) (*pb.ScoreResponse, error){

	//评分
	sc := &model.Score{
		CarBand:   score.GetCarBand(),
		CarSeries: score.GetCarSeries(),
		Score:     score.GetScore(),
		UserID:    uint(score.GetUserid()),
	}
	res,err:=sc.MakeScore()
	if err != nil {
		return new(pb.ScoreResponse),status.Error(codes.Aborted,err.Error())
	}
	return serializer.BuildScoreResponse(res), status.Error(codes.OK,"ok")
}

//修改评分
func (s *ScoreServiceServer)ModifyScore(ctx context.Context, sr *pb.ScoreRequest) (*pb.ScoreResponse, error){
	var score = &model.Score{
		CarBand:   sr.GetCarBand(),
		CarSeries: sr.GetCarSeries(),
		Score:     sr.GetScore(),
		UserID:    uint(sr.GetUserid()),
	}

	res,err:=score.ModifyScore()
	if err != nil {
		return new(pb.ScoreResponse), status.Error(codes.Aborted,err.Error())
	}
	return serializer.BuildScoreResponse(res), status.Error(codes.OK,"ok")

}

//查看评分
func (s *ScoreServiceServer)FindMYScore(ms *pb.MyScoresRequest, m pb.ScoreService_FindMYScoreServer) error{

	res,err:=model.FindMYScore(ms.GetUserid())
	if err != nil {
		return status.Error(codes.Aborted,err.Error())
	}
	for _,v:=range *res{
		v2:=serializer.BuildScoreResponse(&v)
		if err:=m.Send(v2);err!=nil{
			//todo
			continue
		}
	}
	return status.Error(codes.OK,"ok")
}
