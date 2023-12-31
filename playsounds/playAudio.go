package playsounds

import (
	"bytes"
	"embed"
	"errors"
	"io"
	"io/fs"
	"time"

	"github.com/go-audio/wav"
	"github.com/hajimehoshi/oto/v2"
)

// 預嵌入的音頻文件
//
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
	var context *oto.Context
	var readyChan chan struct{}
	var err error

	for _, file := range files {
		if context == nil {
			context, readyChan, err = initContext(language, file)
			if err != nil {
				return err
			}
			<-readyChan // 等待音頻設備準備就緒
		}
		if err := PlayFile(context, language, file); err != nil {
			return err
		}
	}

	return nil
}

func initContext(language string, filename string) (*oto.Context, chan struct{}, error) {
	fs := soundMap[language]
	f, err := fs.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, nil, err
	}

	d := wav.NewDecoder(bytes.NewReader(buf))
	if d == nil {
		return nil, nil, errors.New("invalid WAV file")
	}
	if _, err := d.FullPCMBuffer(); err != nil {
		return nil, nil, err
	}

	context, readyChan, err := oto.NewContext(int(d.SampleRate), int(d.NumChans), 2)
	if err != nil {
		return nil, nil, err
	}

	return context, readyChan, nil
}

func PlayFile(context *oto.Context, language string, filename string) error {
	fs := soundMap[language]
	f, err := fs.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	buf := bytes.NewReader(content)
	d := wav.NewDecoder(buf)
	if d == nil {
		return errors.New("invalid WAV file")
	}
	audioBuf, err := d.FullPCMBuffer()
	if err != nil {
		return err
	}

	player := context.NewPlayer(bytes.NewReader(intToBytes(audioBuf.Data)))
	defer player.Close()

	player.Play()
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
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
