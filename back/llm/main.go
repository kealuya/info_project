package main

import (
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
	"llm/chat"
	"os"
)

//go:embed  chat.gob
var chatId string

func main() {
	Execute()
}

var rootCmd = &cobra.Command{
	Use:   "chat",
	Short: "大模型",
	Long:  "大模型程序内集成",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("需要1个参数，也就是询问大模型的问题.")
			return
		}
		query := args[0]

		c := chat.NewInstance(string(chatId))
		br := c.Conversation(query)
		chatEncode(br.ConversationId)
		fmt.Printf("%s", br.Answer)
	},
}

func chatEncode(conversationId string) {
	// 创建或打开一个文件用于写入
	file, err := os.OpenFile("chat.gob", os.O_RDWR, 0755)
	if err != nil {
		fmt.Println("Failed to open chat.gob:", err)
		return
	}
	defer file.Close()

	// 将数据编码并写入文件
	_, err = file.WriteString(conversationId)
	if err != nil {
		fmt.Println("Failed to encode data:", err)
		return
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
