package main

import (
	"fmt"
	"strings"
)

type TextProcessorInterface interface {
	Process(text string) string
}

// PlainText 是一种基本文本处理器
type PlainText struct{}

func (p *PlainText) Process(text string) string {
	return text
}

// TextDecorator 是装饰器的基础类型
//type TextDecorator struct {
//	TextProcessorInterface
//}

// NewlineDecorator 为文本增加换行符
type NewlineDecorator struct {
	TextProcessorInterface
}

func (d *NewlineDecorator) Process(text string) string {
	return d.TextProcessorInterface.Process(text) + "\n"
}

// UppercaseDecorator 将文本转换为大写
type UppercaseDecorator struct {
	TextProcessorInterface
}

func (d *UppercaseDecorator) Process(text string) string {
	processedText := d.TextProcessorInterface.Process(text)
	return strings.ToUpper(processedText)
}

func main() {
	plain := &PlainText{}
	textWithNewline := &NewlineDecorator{plain}
	textWithDecorations := &UppercaseDecorator{textWithNewline}

	result := textWithDecorations.Process("Hello, Decorator!")
	fmt.Print(result) // Outputs: "HELLO, DECORATOR!\n"
}
