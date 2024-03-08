package controllers

import (
	"encoding/json"
	"net/http"
	m "uts21/models"
)

func sendGetPopularSongResponse (w http.ResponseWriter, message string, data []m.PopularSong) {
	var response m.PopularSongsResponse
	response.Status = 200
	response.Message = message
	response.Data = data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendGetRecommendedSongResponse (w http.ResponseWriter, message string, data m.Song) {
	var response m.RecommendedSongResponse
	response.Status = 200
	response.Message = message
	response.Data = data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse (w http.ResponseWriter, message string) {
	var response m.GeneralResponse
	response.Status = 200
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse (w http.ResponseWriter, status int, message string) {
	var response m.GeneralResponse
	response.Status = status 
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

