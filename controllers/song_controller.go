package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	m "uts21/models"
	// "github.com/gorilla/mux"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)


func GetPopularSongs(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	password := r.Form.Get("password")
	userType := getUserType(w, db, password)

	// ini bisa dijadiin function tersendiri
	
	if userType != 2 {
		sendErrorResponse(w, 400, "You dont have access!")
		return
	}

	query := "SELECT s.songId, s.songTitle, s.songDuration, s.songSinger, SUM(dp.timePlayed) AS 'total_timeplayed' FROM detailplaylistsong dp JOIN songs s ON dp.songId = s.songId GROUP BY s.songId ORDER BY total_timeplayed DESC;"

	results, err := db.Query(query)
	if err != nil {
		sendErrorResponse(w, 500, "Internal Server Error! Login Fail")
		return
	} 

	var popularSong m.PopularSong
	var popularSongs []m.PopularSong

	for results.Next(){
		if err := results.Scan(
		&popularSong.Song.SongId, &popularSong.Song.SongTitle, &popularSong.Song.SongDuration, &popularSong.Song.SongSinger, &popularSong.TimePlayed); err != nil {
			print(err.Error())
			return
		} else {
			popularSongs = append(popularSongs, popularSong)
		}
	}

	sendGetPopularSongResponse(w, "Success", popularSongs)
}

func GetRecommendedSong(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	password := r.Form.Get("password")
	userType := getUserType(w, db, password)

	// ini bisa dijadiin function tersendiri
	
	if userType != 1 {
		sendErrorResponse(w, 400, "You dont have access!")
		return
	}

	query := "SELECT s.songId, s.songTitle, s.songDuration, s.songSinger, SUM(dp.timePlayed) AS 'total_timeplayed' FROM detailplaylistsong dp JOIN songs s ON dp.songId = s.songId JOIN playlists p ON dp.playlistId = p.playlistId JOIN users u ON p.userId = u.userId WHERE u.userPassword = ? GROUP BY s.songId ORDER BY total_timeplayed DESC LIMIT 1;"

	result, errResult := db.Query(query, password)
	if errResult != nil {
		fmt.Println(errResult)
		sendErrorResponse(w, 500, "Internal Server Error! Database Query Fail!")
		return
	}


	var recommendedSong m.Song 
	if result.Next() {
		errScan := result.Scan(&recommendedSong.SongId, &recommendedSong.SongTitle, &recommendedSong.SongDuration, &recommendedSong.SongSinger, &sql.NullInt64{})

		if errScan != nil {
			fmt.Println(errScan)
			sendErrorResponse(w, 500, "Internal Server Error! Fail to Scan!")
			return
		}
	}else{
		sendErrorResponse(w, 404, "account not found!")
		return
	}

	sendGetRecommendedSongResponse(w, "Success", recommendedSong)
}