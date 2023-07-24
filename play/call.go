package play

const (
	MandarinPath = "play/sounds/mandarin"
	EnglishPath  = "play/sounds/english"
	HokkienPath  = "play/sounds/hokkien"
	HakkaPath    = "play/sounds/hakka"
)

func Mandarin(call, counter int) error {

	return PlayAudio(MandarinRules(call, counter))

}

func English(call, counter int) error {

	return PlayAudio(EnglishRules(call, counter))
}

func Hokkien(call, counter int) error {

	return PlayAudio(HokkienRules(call, counter))
}

func Hakka(call, counter int) error {

	return PlayAudio(HakkaRules(call, counter))
}
