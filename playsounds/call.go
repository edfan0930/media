package playsounds

const (
	MandarinPath = "sounds/mandarin"
	EnglishPath  = "sounds/english"
	HokkienPath  = "sounds/hokkien"
	HakkaPath    = "sounds/hakka"
)

func Mandarin(call, counter int) error {

	return PlayAudio("mandarin", MandarinRules(call, counter))

}

func MandarinWithSingleSound(call int) error {
	err := PlayAudio("mandarin", SingleWithMandarin(call))

	return err
}

func English(call, counter int) error {

	return PlayAudio("english", EnglishRules(call, counter))
}

func Hokkien(call, counter int) error {

	return PlayAudio("hokkien", HokkienRules(call, counter))
}

func Hakka(call, counter int) error {

	return PlayAudio("hakka", HakkaRules(call, counter))
}

func Other(fileName string) error {

	fileName = "sounds/other/" + fileName + ".wav"

	err := PlayAudio("other", []string{fileName})
	return err
}
