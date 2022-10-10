package handlers

import (
	artis "dumbsound/dto/artiss"
	dto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerArtis struct {
	ArtisRepository repositories.ArtisRepository
}

func HandlerArtis(ArtisRepository repositories.ArtisRepository) *handlerArtis {
	return &handlerArtis{ArtisRepository}
}

func (h *handlerArtis) FindArtis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	artis, err := h.ArtisRepository.FindArtis()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: artis}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArtis) GetArtis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	artis, err := h.ArtisRepository.GetArtis(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: artis}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArtis) CreateArtis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	old, _ := strconv.Atoi(r.FormValue("old"))
	career, _ := strconv.Atoi(r.FormValue("start_career"))

	request := artis.CreateArtisRequest{
		Name:         r.FormValue("name"),
		Old:          old,
		Type_Artis:   r.FormValue("type_artis"),
		Start_Career: career,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	artis := models.Artis{
		Name:         request.Name,
		Old:          request.Old,
		Type_Artis:   request.Type_Artis,
		Start_Career: request.Start_Career,
	}

	// err := mysql.DB.Create(&artis).Error
	artis, err = h.ArtisRepository.CreateArtis(artis)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	artis, _ = h.ArtisRepository.GetArtis(artis.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: artis}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArtis) UpdateArtis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(artis.UpdateArtisRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	artis, err := h.ArtisRepository.GetArtis(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// if request.Name != "" {
	// 	artis.Name = request.Name
	// }

	data, err := h.ArtisRepository.UpdateArtis(artis)
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

func (h *handlerArtis) DeleteArtis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	artis, err := h.ArtisRepository.GetArtis(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ArtisRepository.DeleteArtis(artis)
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
