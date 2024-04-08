package task

import "szyx_back/entity"

//任务
type Task struct {
	entity.Base
	TaskName string `json:"TaskName" description:"任务名称" `
}
