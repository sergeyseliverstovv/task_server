package handlers

import (
	"context"

	api1 "github.com/sergeyseliverstovv/api/gen/go"
	"google.golang.org/grpc"
)

type serverApi struct {
	api1.UnimplementedTaskManagerServer
}

func Register(gRPC *grpc.Server) {
	api1.RegisterTaskManagerServer(gRPC, &serverApi{})
}

func (s *serverApi) GetTask(ctx context.Context, req *api1.GetTaskRequest) (*api1.GetTaskResponse, error) {
	//panic("not implemented GetTask")
	return &api1.GetTaskResponse{TaskId: []int64{1, 2}, TaskList: []string{"dkdk", "jeje"}}, nil
}

func (s *serverApi) GetTaskId(ctx context.Context, req *api1.GetTaskIdRequest) (*api1.GetTaskIdResponse, error) {
	panic("not implemented GetTaskId")
}

func (s *serverApi) CreateTask(ctx context.Context, req *api1.CreateNewTaskRequest) (*api1.CreateNewTaskResponse, error) {
	panic("not implemented CreateTask")
}

func (s *serverApi) UpdateTask(ctx context.Context, req *api1.UpdateTaskRequest) (*api1.UpdateTaskResponse, error) {
	panic("not implemented UpdateTask")
}

func (s *serverApi) DeleteTask(ctx context.Context, req *api1.DeleteTaskRequest) (*api1.DeleteTaskResponse, error) {
	panic("not implemented DeleteTask")
}
