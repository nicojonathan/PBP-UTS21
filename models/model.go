package models

import "time"

type User struct {
	UserId       int    `json:"id,omitempty" gorm:"primaryKey"`
	UserName     string `json:"name,omitempty"`
	UserEmail    string `json:"email,omitempty"`
	UserPassword string `json:"password,omitempty"`
	UserCountry  string `json:"country,omitempty"`
	// UserType     int    `json:"type,omitempty"`
}

type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data,omitempty"`
}

type Playlist struct {
	PlaylistId          int    `json:"id,omitempty" gorm:"primaryKey"`
	PlaylistName        string `json:"name,omitempty"`
	PlayListDateCreated time.Time`json:"date_created,omitempty"`
	PlayListState bool `json:"playlistState,omitempty"`
	UserId int `json:"userId,omitempty"`
}

type Song struct {
	SongId          int    `json:"id,omitempty" gorm:"primaryKey"`
	SongTitle        string `json:"title,omitempty"`
	SongDuration float64 `json:"duration,omitempty"`
	SongSinger       string `json:"singer,omitempty"`
}

type PopularSong struct {
	Song Song `json:"Song"`
	TimePlayed int `json:"time played"`
}


type PopularSongsResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    []PopularSong `json:"data"`
}

type RecommendedSongResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    Song `json:"data"`
}

type DetailPlaylistSong struct {
	PlayList Playlist `json:"playlist"`
	Song Song `json:"song"`
	TimePlayed int `json:"timeplayed"`
}

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}





    //type TransactionDetail struct {
	// 	ID       int     `json:"transactionID"`
	// 	User     User    `json:"user"`
	// 	Product  Product `json:"product"`
	// 	Quantity int     `json:"quantity"`
	// }
	
	// type TransactionsDetail struct {
	// 	Transaction []TransactionDetail `json:"transactions"`
	// }
	
	// type TransactionDetailResponse struct {
	// 	Status  int                `json:"status"`
	// 	Message string             `json:"message"`
	// 	Data    TransactionsDetail `json:"data"`
	// }