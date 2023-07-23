package play

import (
	"embed"
	"os"

	"github.com/go-audio/wav"
	"github.com/hajimehoshi/oto"
)

//go:embed sounds/english/**.wav

var englishSounds embed.FS

func PlayAudio(files []string) {

	for _, file := range files {
		if err := PlayFile(file); err != nil {
			panic(err)
		}
	}
}

func PlayFile(filename string) error {
	// Open the WAV file
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Create a new WAV decoder
	d := wav.NewDecoder(f)

	// Decode the full audio
	buf, err := d.FullPCMBuffer()
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
	return playAudioData(buf.Data, p)
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
