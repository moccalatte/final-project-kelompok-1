package usecase

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
	"fmt"
)

type QuestionUseCase interface {
	AddQuestion(payload dto.QuestionRequestDto) (model.Question, error)
	FindQuestionById(id string) (model.Question, error)
	Update(payload dto.QuestionRequestDto, id string) (model.Question, error)
	Delete(id string) (model.Question, error)
	Answer(payload dto.QuestionRequestDto, id string) (model.Question, error)
}

type questionUseCase struct {
	repo repository.QuestionRepository
}

func (s *questionUseCase) AddQuestion(payload dto.QuestionRequestDto) (model.Question, error) {
	newSession := model.Question{
		SessionID:   model.Session{SessionID: payload.SessionID},
		StudentID:   model.Student{StudentID: payload.StudentID},
		TrainerID:   model.Users{UserID: payload.TrainerID},
		Title:       payload.Title,
		Description: payload.Description,
		CourseID:    model.Course{CourseID: payload.CourseID},
		Image:       payload.Image,
	}

	createsQuestion, err := s.repo.Create(newSession)

	if err != nil {
		return model.Question{}, fmt.Errorf("failed to save data : %s", err.Error())
	}
	return createsQuestion, nil
}

func (s *questionUseCase) FindQuestionById(id string) (model.Question, error) {
	Question, err := s.repo.GetById(id)

	if err != nil {
		return model.Question{}, fmt.Errorf("failed to get data by id : %s", err.Error())
	}
	return Question, nil
}

func (s *questionUseCase) Update(payload dto.QuestionRequestDto, id string) (model.Question, error) {
	question := model.Question{
		SessionID:   model.Session{SessionID: payload.SessionID},
		StudentID:   model.Student{StudentID: payload.StudentID},
		TrainerID:   model.Users{UserID: payload.TrainerID},
		Title:       payload.Title,
		Description: payload.Description,
		CourseID:    model.Course{CourseID: payload.CourseID},
		Image:       payload.Image,
	}

	question, err := s.repo.Update(question, id)

	if err != nil {
		return model.Question{}, fmt.Errorf("failed to update data : %s", err.Error())
	}

	return question, nil
}

func (s *questionUseCase) Delete(id string) (model.Question, error) {
	question, err := s.repo.Delete(id)

	if err != nil {
		return model.Question{}, fmt.Errorf("failed to delete data : %s", err.Error())
	}

	return question, nil
}

func (s *questionUseCase) Answer(payload dto.QuestionRequestDto, id string) (model.Question, error) {
	answer := model.Question{
		Answer: payload.Answer,
	}
	answered, err := s.repo.Answer(answer, id)

	if err != nil {
		return model.Question{}, fmt.Errorf("failed to answer : %s", err.Error())
	}

	return answered, nil

}

func NewQuestion(repo repository.QuestionRepository) QuestionUseCase {
	return &questionUseCase{repo: repo}
}