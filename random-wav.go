package main

import (
	"encoding/binary"
	"math/rand"
	"os"
	"time"
)

func main() {
	const (
		filename      = "noise.wav"
		sampleRate    = 44100            // 44.1 kHz
		durationSec   = 5                // durasi dalam detik
		numChannels   = 1                // mono (ubah ke 2 untuk stereo)
		bitsPerSample = 16               // 16-bit PCM
	)

	numSamples := sampleRate * durationSec
	bytesPerSample := bitsPerSample / 8
	dataSize := uint32(numSamples * numChannels * bytesPerSample)
	riffChunkSize := uint32(36) + dataSize // 36 + Subchunk2Size

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Seed random
	rand.Seed(time.Now().UnixNano())

	// RIFF header
	_, _ = f.Write([]byte("RIFF"))
	_ = binary.Write(f, binary.LittleEndian, riffChunkSize) // 4 bytes
	_, _ = f.Write([]byte("WAVE"))

	// fmt subchunk
	_, _ = f.Write([]byte("fmt "))                         // 4 bytes
	_ = binary.Write(f, binary.LittleEndian, uint32(16))   // Subchunk1Size (16 for PCM)
	_ = binary.Write(f, binary.LittleEndian, uint16(1))    // AudioFormat (1 = PCM)
	_ = binary.Write(f, binary.LittleEndian, uint16(numChannels))
	_ = binary.Write(f, binary.LittleEndian, uint32(sampleRate))
	byteRate := uint32(sampleRate * numChannels * bytesPerSample)
	_ = binary.Write(f, binary.LittleEndian, byteRate)
	blockAlign := uint16(numChannels * bytesPerSample)
	_ = binary.Write(f, binary.LittleEndian, blockAlign)
	_ = binary.Write(f, binary.LittleEndian, uint16(bitsPerSample))

	// data subchunk
	_, _ = f.Write([]byte("data"))
	_ = binary.Write(f, binary.LittleEndian, dataSize)

	// Tulis sampel acak (16-bit little-endian). Tulis per-buffer untuk performa.
	buf := make([]byte, 4096) // buffer byte, harus kelipatan bytesPerSample
	for i := 0; i < numSamples; {
		// Mengisi buffer dengan sampel int16 acak
		// jumlah sampel yang dapat ditampung buf:
		maxSamples := len(buf) / bytesPerSample
		samplesToWrite := maxSamples
		remaining := numSamples - i
		if remaining < samplesToWrite {
			samplesToWrite = remaining
		}

		// write samplesToWrite int16 into buf
		// little-endian: low byte first
		for s := 0; s < samplesToWrite; s++ {
			v := int16(rand.Intn(65536) - 32768)
			idx := s * 2
			buf[idx] = byte(v)
			buf[idx+1] = byte(v >> 8)
		}

		// tulis buf[0: samplesToWrite*2]
		nw, err := f.Write(buf[:samplesToWrite*bytesPerSample])
		if err != nil {
			panic(err)
		}
		if nw != samplesToWrite*bytesPerSample {
			panic("short write")
		}

		i += samplesToWrite
	}

	// file closed by defer
}
