package testingg_test

import (
	"testing"

	"simplecrud/testingg"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	// Start mocking the database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Define the expected insert query and result

	expectedInsert := "INSERT INTO todo VALUES\\('test'\\);?"

	mock.ExpectExec(expectedInsert).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the function being tested
	testingg.CreateTask(db, "test", "50")

	// Verify that the expected queries were executed
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateTask(t *testing.T) {
	// Start mocking the database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Define the expected update query and result
	expectedUpdate := "UPDATE todo SET TASK = 'new_task', DUE = '70' WHERE TASK = 'eats';"
	mock.ExpectExec(expectedUpdate).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the function being tested
	testingg.UpdateTask(db, "eats", "70", "new_task")

	// Verify that the expected queries were executed
	assert.NoError(t, mock.ExpectationsWereMet())
}
