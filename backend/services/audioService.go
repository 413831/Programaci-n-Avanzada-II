package services

import "backend/domain"

type AudioService interface {
	GetAllAudios() ([]domain.Audio, error)
}

type DefaultAudioService struct {
	repo domain.AudioRepository
}

func (s DefaultAudioService) GetAllAudios() ([]domain.Audio, error) {
	return s.repo.FindAll()
}

func InitAudioService(repository domain.AudioRepository) DefaultAudioService {
	return DefaultAudioService{repository}
}
