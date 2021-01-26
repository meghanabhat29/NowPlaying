package models

import (
	"database/sql"
	"entities"
)

//SongModel ...
type SongModel struct {
	Db *sql.DB
}

//FindAll ...
func (songModel SongModel) FindAll() (song []entities.Songs, err error) {
	rows, err := songModel.Db.Query("select * from songs")

	if err != nil {
		return nil, err
	} else {
		var songs []entities.Songs
		for rows.Next() {
			var id int64
			var name string
			var artist string
			var album string
			var genre string
			var year int64

			err2 := rows.Scan(&id, &name, &artist, &album, &genre, &year)
			if err2 != nil {
				return nil, err2
			} else {
				song := entities.Songs{
					ID:     id,
					Name:   name,
					Artist: artist,
					Album:  album,
					Genre:  genre,
					Year:   year,
				}
				songs = append(songs, song)
			}
		}

		return songs, nil
	}

}

//Search ...
func (songModel SongModel) Search(keyword string) (song []entities.Songs, err error) {
	rows, err := songModel.Db.Query("select * from songs where Name like ?", "%"+keyword+
		"%")

	if err != nil {
		return nil, err
	} else {
		var songs []entities.Songs
		for rows.Next() {
			var id int64
			var name string
			var artist string
			var album string
			var genre string
			var year int64

			err2 := rows.Scan(&id, &name, &artist, &album, &genre, &year)
			if err2 != nil {
				return nil, err2
			} else {
				song := entities.Songs{
					ID:     id,
					Name:   name,
					Artist: artist,
					Album:  album,
					Genre:  genre,
					Year:   year,
				}
				songs = append(songs, song)
			}
		}

		return songs, nil
	}

}

//SearchYear ...
/*func (songModel SongModel) SearchYear(min, max int64) (song []entities.Songs, err error) {
	rows, err := songModel.Db.Query("select * from songs where Year > ? and < ?", min, max)

	if err != nil {
		return nil, err
	} else {
		var songs []entities.Songs
		for rows.Next() {
			var id int64
			var name string
			var artist string
			var album string
			var genre string
			var year int64

			err2 := rows.Scan(&id, &name, &artist, &album, &genre, &year)
			if err2 != nil {
				return nil, err2
			} else {
				song := entities.Songs{
					ID:     id,
					Name:   name,
					Artist: artist,
					Album:  album,
					Genre:  genre,
					Year:   year,
				}
				songs = append(songs, song)
			}
		}

		return songs, nil
	}

}*/

//Create ...
func (songModel SongModel) Create(song *entities.Songs) (err error) {
	result, err := songModel.Db.Exec("insert into songs(Name, Artist,Album,Genre,Year) values(?,?,?,?,?)", song.Name, song.Artist, song.Album, song.Genre, song.Year)

	if err != nil {
		return err
	} else {
		song.ID, _ = result.LastInsertId()
		return nil
	}

}

//Update ...
func (songModel SongModel) Update(song *entities.Songs) (int64, error) {
	result, err := songModel.Db.Exec("update songs set Name=?,Artist=?,Album=?,Genre=?,Year=? where Id=?", song.Name, song.Artist, song.Album, song.Genre, song.Year, song.ID)

	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

//Delete ...
func (songModel SongModel) Delete(id int64) (int64, error) {
	result, err := songModel.Db.Exec("delete from songs where Id=?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}
