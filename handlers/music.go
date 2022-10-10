package handlers

import (
	musicdto "dumbsound/dto/music"
	dto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerMusic struct {
	MusicRepository repositories.MusicRepository
}

// `path_file` Global variable here ...
var PathFile = os.Getenv("PATH_FILE")

func HandlerMusic(MusicRepository repositories.MusicRepository) *handlerMusic {
	return &handlerMusic{MusicRepository}
}

func (h *handlerMusic) FindMusics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	musics, err := h.MusicRepository.FindMusics()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image & Video property here ...
	for i, p := range musics {
		musics[i].ThumbnailMusic = os.Getenv("PATH_FILE") + p.ThumbnailMusic
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: musics}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerMusic) GetMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var music models.Music
	music, err := h.MusicRepository.GetMusic(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	music.ThumbnailMusic = os.Getenv("PATH_FILE") + music.ThumbnailMusic

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseMusic(music)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerMusic) CreateMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64))

	// Get dataFile from midleware and store to filename variable here ...
	dataContex := r.Context().Value("dataMusic")
	filename := dataContex.(string)

	dataAudio := r.Context().Value("dataAudio")
	filethumb := dataAudio.(string)

	year, _ := strconv.Atoi(r.FormValue("year"))
	artis_id, _ := strconv.Atoi(r.FormValue("artis_id"))

	request := musicdto.MusicRequest{
		Title:          r.FormValue("title"),
		Year:           year,
		ArtisID:        artis_id,
		ThumbnailMusic: r.FormValue("thumbnailMusic"),
		Attache:        r.FormValue("attache"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	music := models.Music{
		Title:          request.Title,
		ThumbnailMusic: filename,
		Year:           request.Year,
		ArtisID:        artis_id,
		Artis:          models.ArtisResponse{},
		Attache:        filethumb,
	}

	// err := mysql.DB.Create(&music).Error
	music, err = h.MusicRepository.CreateMusic(music)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	music, _ = h.MusicRepository.GetMusic(music.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: music}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerMusic) UpdateMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(musicdto.UpdateMusicRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	music, err := h.MusicRepository.GetMusic(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// if request.Title != "" {
	// 	music.Title = request.Title
	// }

	// if request.ThumbnailMusic != "" {
	// 	music.ThumbnailMusic = request.ThumbnailMusic
	// }

	// if request.LinkMusic != "" {
	// 	music.LinkMusic = request.LinkMusic
	// }

	// if request.Year != 0 {
	// 	music.Year = request.Year
	// }

	// if request.Desc != "" {
	// 	music.Desc = request.Desc
	// }

	data, err := h.MusicRepository.UpdateMusic(music)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerMusic) DeleteMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	music, err := h.MusicRepository.GetMusic(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.MusicRepository.DeleteMusic(music)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDelMusic(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseMusic(u models.Music) models.MusicResponse {
	return models.MusicResponse{
		ID:             u.ID,
		Title:          u.Title,
		ThumbnailMusic: u.ThumbnailMusic,
		Year:           u.Year,
		Artis:          u.Artis,
		Attache:        u.Attache,
	}
}

func convertResponseDelMusic(u models.Music) models.MusicResponse {
	return models.MusicResponse{
		ID:    u.ID,
		Title: u.Title,
		Year:  u.Year,
	}
}
