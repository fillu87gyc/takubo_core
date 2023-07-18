package repository

type IVoiceRepository interface {
	Speak(text string) error
}
