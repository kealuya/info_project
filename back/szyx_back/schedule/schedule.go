package schedule

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
	"szyx_back/common"
	"szyx_back/db"
)

func InitTask() {
	//每5分钟执行一次，获取转译文件
	tk_Middle := toolbox.NewTask("getKdxfResult", "0 */5 * * * *", GetKdxfResult)
	toolbox.AddTask("getKdxfResult", tk_Middle)
}

func GetKdxfResult() error {
	logs.Info("========定时任务 获取Kdxf转译文件========")
	//获取库中，需要查询结果的会议 区间设计为 3天内的会议 ，超出的，kdxf 也不会保留转译的内容
	meetingList, err := db.GetMeetingListTranslation()

	for _, meeting := range meetingList {
		common.DoHttpPost_kdxf_audio_translation_getResult(meeting.MeetingId)
	}

	return err
}
