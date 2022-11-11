package movie

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/storage_implementation/internal/domain"
	"github.com/stretchr/testify/assert"
)

var (
	ERRORFORZADO = errors.New("error forzado")
)
var movie_test = domain.Movie{
	ID:           1,
	Created_at:   time.Now(),
	Updated_at:   time.Now(),
	Title:        "Cars 1",
	Rating:       4,
	Awards:       2,
	Release_date: time.Layout,
	// Length:       0,
	// Genre_id:     0,
}

func TestSave(t *testing.T) {
	// (db) es un puntero de sql.DB, por ende lo podemos reemplazar en donde hagamos uso del mismo.
	// El siguiente parámetro “mock” es quien nos permite agregar mediante regex respuestas tanto para Queries como para Exec.
	// err queda != nil si hubo algun problema con la inicialización de la mock db.
	// Se utiliza la instancia de mock para establecer qué realizar al cumplir con cada consulta sql.
	db, mock, err := sqlmock.New()
  	assert.NoError(t, err)
  	defer db.Close()

	// test ok
	t.Run("Store Ok", func(t *testing.T) {

		// Arrange
		mock.ExpectPrepare(regexp.QuoteMeta(SAVE_MOVIE))
		// sqlmock.NewResult(lastInsertID int64, rowsAffected int64)
		mock.ExpectExec(regexp.QuoteMeta(SAVE_MOVIE)).WillReturnResult(sqlmock.NewResult(1, 1))

		// indico en string las columnas que va a tener
		columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
		// creo las rows
		rows := sqlmock.NewRows(columns)
		// indico que dato va a tener cada campo
		rows.AddRow(movie_test.ID, movie_test.Title, movie_test.Rating, movie_test.Awards, movie_test.Length, movie_test.Genre_id)
		
		// WithArgs will match given expected args to actual database query arguments
		// WillReturnRows specifies the set of resulting rows that will be returned by the triggered query
		mock.ExpectQuery(regexp.QuoteMeta(GET_MOVIE)).WithArgs(1).WillReturnRows(rows)

		repository := NewRepository(db)
		ctx := context.TODO()

		// Act
		newID, err := repository.Save(ctx, movie_test)
		assert.NoError(t, err)

		movieResult, err := repository.Get(ctx, int(newID))
		assert.NoError(t, err)

		// Assert
		assert.NotNil(t, movieResult)
		assert.Equal(t, movie_test.ID, movieResult.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Store Fail", func(t *testing.T) {
		// Arrange
		mock.ExpectPrepare(regexp.QuoteMeta(SAVE_MOVIE))
		mock.ExpectExec(regexp.QuoteMeta(SAVE_MOVIE)).WillReturnError(ERRORFORZADO)

		respository := NewRepository(db)
		ctx := context.TODO()

		// Act
		id, err := respository.Save(ctx, movie_test)
		
		// Assertions
		assert.EqualError(t, err, ERRORFORZADO.Error())
		assert.Empty(t, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("GetOne Ok", func(t *testing.T) {
		// Arrange
		columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(movie_test.ID, movie_test.Title, movie_test.Rating, movie_test.Awards, movie_test.Length, movie_test.Genre_id)
		mock.ExpectQuery(regexp.QuoteMeta(GET_MOVIE)).WithArgs(movie_test.ID).WillReturnRows(rows)

		repository := NewRepository(db)
		ctx := context.TODO()
		
		// Act
		movieResult, err := repository.Get(ctx, movie_test.ID)
		
		//Assert
		assert.NoError(t, err)
		assert.Equal(t, movie_test.Title, movieResult.Title)
		assert.Equal(t, movie_test.ID, movieResult.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("GetAll Ok", func(t *testing.T) {
		// Arrange
		columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
		rows := sqlmock.NewRows(columns)
		// movies a devolver
		movies := []domain.Movie{{ID: 1, Title: "Avatar", Rating: 22, Awards: 99}, {ID: 2, Title: "Simpson", Rating: 33, Awards: 11}}

		for _, movie := range movies {
			rows.AddRow(movie.ID, movie.Title, movie.Rating, movie.Awards, movie.Length, movie.Genre_id)
		}

		mock.ExpectQuery(regexp.QuoteMeta(GET_ALL_MOVIES)).WillReturnRows(rows)

		repo := NewRepository(db)
		// Act
		resultMovies, err := repo.GetAll(context.TODO())

		//Assert
		assert.NoError(t, err)
		assert.Equal(t, movies, resultMovies)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUpdate(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Update Ok", func(t *testing.T) {
		// Arrange
		m := movie_test
		mock.ExpectPrepare(regexp.QuoteMeta(UPDATE_MOVIE))
		mock.ExpectExec(regexp.QuoteMeta(UPDATE_MOVIE)).WithArgs(m.Title, m.Rating, m.Awards, m.Length, m.Genre_id, m.ID).WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewRepository(db)
		
		// Act
		err := repo.Update(context.TODO(), m, 1)

		//Assert
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		// poderia verificar que el dato se haya modificado?

	})

}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()


	t.Run("Delete Ok", func(t *testing.T) {
		// Arrange
		mock.ExpectPrepare(regexp.QuoteMeta(DELETE_MOVIE))
		mock.ExpectExec(regexp.QuoteMeta(DELETE_MOVIE)).WithArgs(movie_test.ID).WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewRepository(db)

		// Act
		err = repo.Delete(context.TODO(), 1)
		
		//Assert
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		
	})
	
}