package movie

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/storage_implementation/internal/domain"
)

// queries
const (
	GET_ALL_MOVIES = "SELECT id, title, rating, awards, length, genre_id FROM movies"

	GET_MOVIE = "SELECT id, title, rating, awards, length, genre_id FROM movies WHERE id=?;"

	UPDATE_MOVIE = "UPDATE movies SET title=?, rating=?, awards=?, lenght=?, genre_id=? where id=?;"

	SAVE_MOVIE = "INSERT INTO movies SET (title, rating, awards, lenght, genre_id) VALUES (?,?,?,?,?);"

	DELETE_MOVIE = "DElETE FROM movies WHERE id=?;"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Movie, error)
	Get(ctx context.Context, id int) (domain.Movie, error)
	Save(ctx context.Context, b domain.Movie) (int64, error)
	Exists(ctx context.Context, id int) bool
	Update(ctx context.Context, b domain.Movie, id int) error
	Delete(ctx context.Context, id int64) error
}

// el repo tiene como componente a un db sql
type repository struct {
	db *sql.DB
}

// recibe la bd y devuelve la interface Repository
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r * repository) Get(ctx context.Context, id int) (domain.Movie, error){

	// ejecutar la query
	// devuelve una sola row
	row := r.db.QueryRow(GET_MOVIE, id)

	// variable para guardar la movie que leyo de db
	var movie domain.Movie

	// seteo el resultado de la query en la variable
	if err := row.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.Genre_id); err != nil {		
		return domain.Movie{}, nil
	}

	return movie, nil

}

func (r * repository) Save(ctx context.Context, m domain.Movie) (int64, error){
	// preparamos la query
	stm, err := r.db.Prepare(SAVE_MOVIE)
	if err != nil {
		return 0, err
	}
	// cerramos para no perder memoria
	defer stm.Close()

	// ejecutamos la query con los valores a del record a guardar
	result, err := stm.Exec(m.Title, m.Rating, m.Awards, m.Length, m.Genre_id)
	if err != nil {
		return 0, err
	}

	// guardamos el ultimo id que puso (o sea, en esta ultima ejecucion) para devolverlo
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r * repository) Exists(ctx context.Context, id int) bool{
	return false
}

func (r * repository) Update(ctx context.Context, m domain.Movie, id int) error{
	// preparamos la query
	stm, err := r.db.Prepare(UPDATE_MOVIE)
	if err != nil {
		return err
	}
	// cerramos para no perder memoria
	defer stm.Close()

	// ejecutamos la query con los valores a reemplazar en el record
	result, err := stm.Exec(m.Title, m.Rating, m.Awards, m.Length, m.Genre_id, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("error: no affected rows, id no encontrado")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	// preparamos la query
	stm, err := r.db.Prepare(DELETE_MOVIE)
	if err != nil {
		return err
	}
	// cerramos para no perder memoria
	defer stm.Close()

	// ejecutamos la query con los valores a reemplazar en el record
	result, err := stm.Exec(id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return errors.New("error: no affected rows, no se encontro id")
	}

	return nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Movie, error) {

	// creo el array para guardar todas las rows devueltas por la query
	var movies []domain.Movie

	// ejecuto la query
	rows, err := r.db.Query(GET_ALL_MOVIES)
	if err != nil {
		return []domain.Movie{}, err
	}

	//recorro todas las rows para guardarlas en el array de movies
	for rows.Next() {
		var movie domain.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.Genre_id)
		if err != nil {
			return []domain.Movie{}, err
		}
		// por cada row leida, la appendeo en el arrays de movies
		movies = append(movies, movie)
	}

	return movies, nil
}