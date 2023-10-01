package playsounds

import (
	"bytes"
	"embed"
	"io"
	"io/fs"

	"github.com/go-audio/wav"
	"github.com/hajimehoshi/oto"
)

//go:embed sounds/english/*.wav
var englishSounds embed.FS

//go:embed sounds/mandarin/*.wav
var mandarinSounds embed.FS

//go:embed sounds/hokkien/*.wav
var hokkienSounds embed.FS

//go:embed sounds/hakka/*.wav
var hakkaSounds embed.FS

//go:embed sounds/other/*.wav
var otherSounds embed.FS

var soundMap = map[string]fs.FS{
	"english":  englishSounds,
	"mandarin": mandarinSounds,
	"hokkien":  hokkienSounds,
	"hakka":    hakkaSounds,
	"other":    otherSounds,
}

func PlayAudio(language string, files []string) error {
	for _, file := range files {
		if err := PlayFile(language, file); err != nil {
			return err
		}
	}

	return nil
}

func PlayFile(language string, filename string) error {
	// Retrieve the correct FS based on the language
	fs := soundMap[language]
	// Open the WAV file from the embedded file system
	f, err := fs.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Read file content
	content, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	// Create a bytes reader
	buf := bytes.NewReader(content)

	// Create a new WAV decoder
	d := wav.NewDecoder(buf)

	// Decode the full audio
	audioBuf, err := d.FullPCMBuffer()
	if err != nil {
		return err
	}
	// Create the player with correct sample rate and number of channels
	c, err := oto.NewContext(int(d.SampleRate), int(d.NumChans), 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	// Convert and play
	return playAudioData(audioBuf.Data, p)
}

func playAudioData(data []int, player *oto.Player) error {
	// Convert the int audio data to byte
	bytesData := intToBytes(data)

	// Write the byte data to the player
	if _, err := player.Write(bytesData); err != nil {
		return err
	}

	return nil
}

func intToBytes(data []int) []byte {
	int16Data := make([]int16, len(data))
	for i, v := range data {
		int16Data[i] = int16(v)
	}

	size := len(int16Data) * 2
	buf := make([]byte, size)
	for i, val := range int16Data {
		buf[i*2] = byte(val & 0xff)
		buf[i*2+1] = byte((val >> 8) & 0xff)
	}

	return buf
}
