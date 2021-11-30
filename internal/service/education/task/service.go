package task

import (
	"context"
	"fmt"
	"log"

	edu_task_api "github.com/ozonmp/edu-task-api/pkg/edu-task-api"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/ozonmp/omp-bot/internal/model/education"
	internal_errors "github.com/ozonmp/omp-bot/internal/pkg/errors"
	mwclient "github.com/ozonmp/omp-bot/internal/pkg/mw/client"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type TaskService interface {
	Describe(taskID uint64) (*education.Task, error)
	List(cursor uint64, limit uint64) ([]education.Task, error)
	Create(education.Task) (uint64, error)
	Update(task map[string]interface{}) error
	Remove(taskID uint64) (bool, error)
}

type DummyTaskService struct {
	grpcClient edu_task_api.TaskApiServiceClient
}

func NewDummyTaskService() *DummyTaskService {

	cfg := config.GetConfigInstance()
	categoryServiceConn, err := grpc.DialContext(
		context.Background(),
		cfg.GrpcServerAddr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(mwclient.AddAppInfoUnary),
	)

	if err != nil {
		log.Fatal("failed to create client")
	}

	grpcClient := edu_task_api.NewTaskApiServiceClient(categoryServiceConn)

	education.TaskEntitiesInit()

	return &DummyTaskService{
		grpcClient: grpcClient,
	}
}

func (s *DummyTaskService) Describe(taskID uint64) (*education.Task, error) {

	msg := edu_task_api.DescribeTaskRequest{
		Id: taskID,
	}

	responce, err := s.grpcClient.DescribeTask(context.TODO(), &msg)

	if err != nil {
		return &education.Task{}, err
	}

	task := education.Task{
		Id:            responce.Task.GetId(),
		Championat_id: responce.Task.GetChampionatId(),
		Difficulty:    responce.Task.GetDifficulty(),
		Title:         responce.Task.GetTitle(),
		Description:   responce.Task.GetDescription(),
	}

	return &task, err
}

func (s *DummyTaskService) List(offset, limit uint64) ([]education.Task, error) {

	msg := edu_task_api.ListTaskRequest{
		Limit:  limit,
		Offset: offset,
	}

	response, err := s.grpcClient.ListTask(context.TODO(), &msg)

	if err != nil {
		return nil, err
	}

	result := make([]education.Task, 0, len(response.Tasks))

	for _, v := range response.Tasks {
		result = append(result, education.Task{
			Id:            v.GetId(),
			Championat_id: v.GetChampionatId(),
			Difficulty:    v.GetDifficulty(),
			Title:         v.GetTitle(),
			Description:   v.GetDescription(),
		})
	}

	return result, nil
}

func (s *DummyTaskService) Create(Task education.Task) (uint64, error) {

	msg := edu_task_api.CreateTaskRequest{
		ChampionatId: Task.Championat_id,
		Difficulty:   Task.Difficulty,
		Title:        Task.Title,
		Description:  Task.Description,
	}

	if err := msg.Validate(); err != nil {
		return 0, err
	}

	resp, err := s.grpcClient.CreateTask(context.TODO(), &msg)
	if err != nil {
		return 0, err
	}

	return resp.GetId(), nil
}

func (s *DummyTaskService) Update(task map[string]interface{}) error {

	msg := edu_task_api.UpdateTaskRequest{}
	msg.Task = &edu_task_api.TaskNull{}

	for i, v := range task {

		if i == "id" || i == "championatID" || i == "difficulty" {
			if _, ok := v.(float64); !ok {
				return fmt.Errorf("%s not type uint64", i)
			}
		} else {
			if _, ok := v.(string); !ok {
				return fmt.Errorf("%s not type string", i)
			}
		}

		switch i {
		case "id":
			msg.Task.Id = uint64(v.(float64))
		case "championatID":
			msg.Task.ChampionatId = wrapperspb.UInt64(uint64(v.(float64)))
		case "difficulty":
			msg.Task.Difficulty = wrapperspb.UInt64(uint64(v.(float64)))
		case "title":
			msg.Task.Title = wrapperspb.String(v.(string))
		case "description":
			msg.Task.Description = wrapperspb.String(v.(string))
		}
	}

	request, err := s.grpcClient.UpdateTask(context.TODO(), &msg)
	if err != nil {
		return err
	}

	if request.Result == edu_task_api.UpdateTaskResponse_RESULT_NOT_FOUND {
		return internal_errors.ErrNotFound
	}

	return nil
}

func (s *DummyTaskService) Remove(taskID uint64) (bool, error) {

	msg := edu_task_api.RemoveTaskRequest{
		Id: taskID,
	}

	request, err := s.grpcClient.RemoveTask(context.TODO(), &msg)

	if err != nil {
		return false, err
	}

	if request.Result == edu_task_api.RemoveTaskResponse_RESULT_NOT_FOUND {
		return false, internal_errors.ErrNotFound
	}

	return true, nil
}
