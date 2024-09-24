package domain

type Audio struct {
	ID         string
	Name       string
	Creator    string
	Length     string
	BPM        float32
	Tone       string
	Category   string
	Genre      string
	UploadDate string
	Filename   string
	ModerateBy string
	Status     string
}

type AudioRepository interface {
	FindAll() ([]Audio, error)
}
