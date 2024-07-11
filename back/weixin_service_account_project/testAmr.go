package main

import (
	"bytes"
	"fmt"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func convertAMRToWAV(amrData []byte) ([]byte, error) {
	inputReader := bytes.NewReader(amrData)
	outputBuffer := bytes.NewBuffer(nil)

	// 添加更多的 FFmpeg 参数
	err := ffmpeg.Input("pipe:0").
		Output("pipe:1", ffmpeg.KwArgs{
			"f":      "s16le",
			"acodec": "pcm_s16le",
			"ar":     "8000", // AMR 通常是 8kHz
			"ac":     "1",    // 单声道
		}).
		WithInput(inputReader).
		WithOutput(outputBuffer).
		Run()

	if err != nil {
		return nil, fmt.Errorf("FFmpeg conversion failed: %v", err)
	}

	if outputBuffer.Len() == 0 {
		return nil, fmt.Errorf("output buffer is empty")
	}

	return outputBuffer.Bytes(), nil
}

func main() {
	// 读取 AMR 文件内容
	amrData, err := os.ReadFile("vvv.amr")
	if err != nil {
		fmt.Printf("Error reading AMR file: %v\n", err)
		return
	}

	fmt.Printf("AMR data size: %d bytes\n", len(amrData))

	// 转换 AMR 到 WAV
	wavData, err := convertAMRToWAV(amrData)
	if err != nil {
		fmt.Printf("Conversion error: %v\n", err)
		return
	}

	fmt.Printf("WAV data size: %d bytes\n", len(wavData))

	// 将转换后的 WAV 数据写入文件
	err = os.WriteFile("output.wav", wavData, 0644)
	if err != nil {
		fmt.Printf("Error writing WAV file: %v\n", err)
		return
	}

	fmt.Println("Conversion completed successfully.")

	// 验证输出文件
	outputFileInfo, err := os.Stat("output.wav")
	if err != nil {
		fmt.Printf("Error checking output file: %v\n", err)
		return
	}
	fmt.Printf("Output file size: %d bytes\n", outputFileInfo.Size())
}
