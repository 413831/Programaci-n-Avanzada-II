package domain

type AudioRepositoryStub struct {
	audios []Audio
}

func (s AudioRepositoryStub) FindAll() ([]Audio, error) {
	return s.audios, nil
}

func InitAudioRepoStub() AudioRepositoryStub {
	audios := []Audio{
		{
			ID:         "1001",
			Name:       "Jazz Blues Vocals",
			Creator:    "MC Tasty",
			Length:     "35",
			BPM:        90,
			Tone:       "C",
			Category:   "acapella",
			Genre:      "jazz blues",
			UploadDate: "24/09/2024",
			Filename:   "1001",
			ModerateBy: "A001",
			Status:     "uploaded",
		},
		{
			ID:         "1001",
			Name:       "Jazz Blues Vocals",
			Creator:    "MC Tasty",
			Length:     "35",
			BPM:        90,
			Tone:       "C",
			Category:   "acapella",
			Genre:      "jazz blues",
			UploadDate: "24/09/2024",
			Filename:   "1001",
			ModerateBy: "A001",
			Status:     "uploaded",
		},
		{
			ID:         "1001",
			Name:       "Jazz Blues Vocals",
			Creator:    "MC Tasty",
			Length:     "35",
			BPM:        90,
			Tone:       "C",
			Category:   "acapella",
			Genre:      "jazz blues",
			UploadDate: "24/09/2024",
			Filename:   "1001",
			ModerateBy: "A001",
			Status:     "uploaded",
		},
	}

	return AudioRepositoryStub{audios: audios}
}
