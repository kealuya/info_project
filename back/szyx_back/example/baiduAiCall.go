package main

import (
	"fmt"
	. "github.com/baidubce/app-builder/go/appbuilder"
	"log"
)

/*

import appbuilder
import os

# 设置环境中的TOKEN，以下TOKEN为访问和QPS受限的试用TOKEN，正式使用请替换为您的个人TOKEN
os.environ["APPBUILDER_TOKEN"] = "bce-v3/ALTAK-n5AYUIUJMarF7F7iFXVeK/1bf65eed7c8c7efef9b11388524fa1087f90ea58"

# 从AppBuilder网页获取并传入应用ID，以下为地理小达人应用ID
app_id = "42eb211a-14b9-43d2-9fae-193c8760ef26"

app_builder_client = appbuilder.AppBuilderClient(app_id)
conversation_id = app_builder_client.create_conversation()

answer = app_builder_client.run(conversation_id, "中国的首都在哪里？春季天气怎么样？有什么适合玩的景点？")
print(answer.content)

*/

func main() {

	config, err := NewSDKConfig("", "bce-v3/ALTAK-8ZBCTrw4yPE50hWsHJd0q/47ca30c03c2899a9ad88db30c496b18121ed4302")
	if err != nil {
		log.Panicln(err)
	}

	appID := "573a3957-c6d3-42f1-b82c-7cf7cbfac681"
	client, err := NewAppBuilderClient(appID, config)
	if err != nil {
		log.Panicln(err)
	}
	conversationID, err := client.CreateConversation()
	if err != nil {
		log.Panicln(err)
	}
	i, err := client.Run(conversationID, "老师是院士，8月31日到9月2日去秦皇岛出差，住宿标准是多少", nil, true)
	if err != nil {
		log.Panicln(err)
	}

	totalAnswer := ""

	for answer, err := i.Next(); err == nil; answer, err = i.Next() {
		//fmt.Println(answer.Answer, answer.Events)
		totalAnswer = totalAnswer + answer.Answer
	}

	fmt.Println("answer:", totalAnswer)

}
