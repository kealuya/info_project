package com.szhtjykj.speech.model;

import lombok.Data;
import org.beetl.sql.annotation.entity.AssignID;
import org.beetl.sql.annotation.entity.Table;


/**
 * @program: kdxf_speech
 * @description:
 * @packagename: com.szhtjykj.speech.model
 * @author: zhanbaohua
 * @date: 2024-05-15 14:00
 **/
    @Data
    @Table(name = "kdxf_meeting")
    public class Meeting {
        @AssignID
        private String  id;// id,主键
        private String  meetingId;// 会议id" `
        private String  meetingType;// 会议类型" ` //audio：音频会议    document：文档会议
        private String  meetingTitle;// 会议标题" `
        private String  meetingTime;// 会议时间  年-月-日  时：分：秒" `
        private String  meetingCity;// 会议城市" `
        private String  meetingAddress;// 会议地址" `
        private String  meetingPeople;// 参会人" `
        private String  meetingFileUrl;// 会议文件地址" `         //完成任务是选择会议，添加会议文件。手动上传用
        private String  taskId;// 任务id，任务与会议为 1对多的关系，故而在此添加关联字段" `
        private String  meetingFlag;// 会议是否使用" ` //0：未使用   1：已使用   完成任务需要选择会议， 要区分会议是否被使用
        private String  corpName;// 企业名称" `
        private String  corpCode;// 企业code"`
        private String  createTime;// 创建时间、结束时间"`
        private String  creater;// 创建人"`
        private String  translationStatus;// 是否转译 0：未转译 1：转义中 2：转译完成
        private String  bz1;// 备注1"`
        private String  bz2;// 备注2"`
        private String  bz3;// 备注3"`
    }

