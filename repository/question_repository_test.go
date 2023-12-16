package repository

import (
	"database/sql"
	"final-project-kelompok-1/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type QuestionRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    QuestionRepository
}

func (suite *QuestionRepositoryTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.sqlmock = sqlmock
	suite.repo = NewQuestionRepository(suite.mockDb)
}

func TestQuestionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(QuestionRepositoryTestSuite))
}

func (suite *QuestionRepositoryTestSuite) TestCreateQuestion() {

	dummy := model.Question{
		QuestionID:  "1212uywewewr2",
		SessionID:   "123d2423",
		StudentID:   "234eedwerwer",
		TrainerID:   "234234fjehe",
		Title:       "go db",
		Description: "error saat debug bagaimana caranya?",
		CourseID:    "4234423423423",
		Answer:      "belum",
		Image:       "yweiuryeiuyriwerwerwerwer.jpg",
		Status:      "belum terjawab",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}

	suite.sqlmock.ExpectBegin()

	rows := sqlmock.NewRows([]string{"question_id", "session_id", "student_id", "trainer_id", "title", "description", "course_id", "image", "answer", "status", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.QuestionID, dummy.SessionID, dummy.StudentID, dummy.TrainerID, dummy.Title, dummy.Description, dummy.CourseID, dummy.Image, dummy.Answer, dummy.Status, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)

	suite.sqlmock.ExpectQuery("insert into question").WillReturnRows(rows)
	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.Create(dummy)
	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), dummy.QuestionID, actual.QuestionID, "QuestionID should match")
	assert.Equal(suite.T(), dummy.SessionID, actual.SessionID, "SessionID should match")
	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID, "StudentID should match")
	assert.Equal(suite.T(), dummy.TrainerID, actual.TrainerID, "TrainerID should match")
	assert.Equal(suite.T(), dummy.Title, actual.Title, "Title should match")
	assert.Equal(suite.T(), dummy.Description, actual.Description, "Description should match")
	assert.Equal(suite.T(), dummy.CourseID, actual.CourseID, "CourseID should match")
	assert.Equal(suite.T(), dummy.Image, actual.Image, "Image should match")
	assert.Equal(suite.T(), dummy.Answer, actual.Answer, "Answer should match")
	assert.Equal(suite.T(), dummy.Status, actual.Status, "Status should match")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")
}


func(suite *QuestionRepositoryTestSuite)TestGetById(){
	dummy := model.Question{
		QuestionID: "12123123",
		SessionID: "0980980809",
		StudentID: "09798798798",
		TrainerID: "98787987978",
		Title: "hujan",
		Description: "yuisuydisyifywerwerwerwer",
		CourseID: "98797343324",
		Image: "gambarnaruto.jpg",
		Answer: "",
		Status: "belum terjawab",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
	}
	query := "select \\* from question where question_id = \\$1;"
	questionID := "12123123"

	rows := sqlmock.NewRows([]string{"question_id", "session_id", "student_id", "trainer_id", "title", "description", "course_id", "image", "answer", "status", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.QuestionID, dummy.SessionID, dummy.StudentID, dummy.TrainerID, dummy.Title, dummy.Description, dummy.CourseID, dummy.Image, dummy.Answer, dummy.Status, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)

		suite.sqlmock.ExpectQuery(query).WithArgs(questionID).WillReturnRows(rows)

		actual, err := suite.repo.GetById(questionID)

		assert.Nil(suite.T(), err, "Error should be nill")
		assert.Equal(suite.T(),dummy.QuestionID,actual.QuestionID, "QuestionID")
		assert.Equal(suite.T(),dummy.SessionID,actual.SessionID, "SessionID")
		assert.Equal(suite.T(),dummy.StudentID,actual.StudentID, "StudentID")
		assert.Equal(suite.T(),dummy.TrainerID,actual.TrainerID, "TrainerID")
		assert.Equal(suite.T(),dummy.Title,actual.Title, "Tittle")
		assert.Equal(suite.T(),dummy.Description,actual.Description, "Description")
		assert.Equal(suite.T(),dummy.CourseID,actual.CourseID, "CourseID")
		assert.Equal(suite.T(),dummy.Image,actual.Image, "Image")
		assert.Equal(suite.T(),dummy.Answer,actual.Answer, "Answer")
		assert.Equal(suite.T(),dummy.Status,actual.Status, "Status")
		assert.Equal(suite.T(),dummy.CreatedAt,actual.CreatedAt, "CreatedAt")
		assert.Equal(suite.T(),dummy.UpdatedAt,actual.UpdatedAt, "UpdatedAt")
		assert.Equal(suite.T(),dummy.IsDeleted,actual.IsDeleted, "IsDeleted")
}

// func (suite *QuestionRepositoryTestSuite) TestUpdateQuestion() {

// 	dummyPayload := model.Question{
// 		SessionID:   "123d2423",
// 		StudentID:   "234eedwerwer",
// 		TrainerID:   "234234fjehe",
// 		Title:       "go db",
// 		Description: "error saat debug bagaimana caranya?",
// 		CourseID:    "4234423423423",
// 		Answer:      "belum",
// 		Image:       "yweiuryeiuyriwerwerwerwer.jpg",
// 		Status:      "Terjawab",
// 		UpdatedAt:   time.Now(),
// 		IsDeleted:   false,
// 	}
// 	dummyResult := model.Question{
// 		QuestionID:  "1212uywewewr2",
// 		SessionID:   dummyPayload.SessionID,
// 		StudentID:   dummyPayload.StudentID,
// 		TrainerID:   dummyPayload.TrainerID,
// 		Title:       dummyPayload.Title,
// 		Description: dummyPayload.Description,
// 		CourseID:    dummyPayload.CourseID,
// 		Answer:      dummyPayload.Answer,
// 		Image:       dummyPayload.Image,
// 		Status:      dummyPayload.Status,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   dummyPayload.UpdatedAt,
// 		IsDeleted:   dummyPayload.IsDeleted,
// 	}

// 	suite.sqlmock.ExpectBegin()

// 	query := "update question set session_id = \\$1, student_id = \\$2, trainer_id = \\$3, title = \\$4, description = \\$5, course_id = \\ $6, image = \\$7, answer = \\$8, status = \\$9, updated_at = \\$10, is_deleted = \\$11 where question_id = \\$12 returning question_id, session_id, student_id, trainer_id, title, description, course_id, image, answer, status, created_at, updated_at, is_deleted;"

// 	rows := sqlmock.NewRows([]string{"question_id", "session_id", "student_id", "trainer_id", "title", "description", "course_id", "image", "answer", "status", "created_at", "updated_at", "is_deleted"}).
// 		AddRow(dummyResult.QuestionID, dummyResult.SessionID, dummyResult.StudentID, dummyResult.TrainerID, dummyResult.Title, dummyResult.Description, dummyResult.CourseID, dummyResult.Image, dummyResult.Answer, dummyResult.Status, dummyResult.CreatedAt, dummyResult.UpdatedAt, dummyResult.IsDeleted)

// 	suite.sqlmock.ExpectQuery(query).WillReturnRows(rows)
// 	suite.sqlmock.ExpectCommit()

// 	actual, err := suite.repo.Update(dummyPayload,dummyResult.QuestionID)
// 	assert.Nil(suite.T(), err, "Error should be nil")
// 	assert.Equal(suite.T(), dummyResult.QuestionID, actual.QuestionID, "QuestionID should match")
// 	assert.Equal(suite.T(), dummyPayload.SessionID, actual.SessionID, "SessionID should match")
// 	assert.Equal(suite.T(), dummyPayload.StudentID, actual.StudentID, "StudentID should match")
// 	assert.Equal(suite.T(), dummyPayload.TrainerID, actual.TrainerID, "TrainerID should match")
// 	assert.Equal(suite.T(), dummyPayload.Title, actual.Title, "Title should match")
// 	assert.Equal(suite.T(), dummyPayload.Description, actual.Description, "Description should match")
// 	assert.Equal(suite.T(), dummyPayload.CourseID, actual.CourseID, "CourseID should match")
// 	assert.Equal(suite.T(), dummyPayload.Image, actual.Image, "Image should match")
// 	assert.Equal(suite.T(), dummyPayload.Answer, actual.Answer, "Answer should match")
// 	assert.Equal(suite.T(), dummyPayload.Status, actual.Status, "Status should match")
// 	assert.Equal(suite.T(), dummyResult.CreatedAt, actual.CreatedAt, "CreatedAt should match")
// 	assert.Equal(suite.T(), dummyPayload.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
// 	assert.Equal(suite.T(), dummyPayload.IsDeleted, actual.IsDeleted, "IsDeleted should match")
// }


