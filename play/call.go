package play

const (
	MandarinPath = "play/sounds/mandarin"
	EnglishPath  = "play/sounds/english"
	HokkienPath  = "play/sounds/hokkien"
	HakkaPath    = "play/sounds/hakka"
)

func Mandarin() error {

	PlayAudio(MandarinRules(83, 3))

	return nil
}
