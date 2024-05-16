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
    @Table(name = "kdxf_meetingFile")
    public class MeetingFile {
        @AssignID
        private Integer  Id;// id,主键
        private String  MeetingId;// 关联的会议id
        private String  FileType;// 文件类型   音频：mp3   文档：word  脑图：xmind
        private String  FileName;// 文件名称
        private String  FileUrl;// 文件云地址
        private String  AudioTime;// 音频时长  只在fileType值为mp3 有作用
        private String  CorpName;// 企业名称
        private String  CorpCode;// 企业code
        private String  CreateTime;// 创建时间
        private String  Creater;// 创建人
        private String  orderId;//科大讯飞订单ID
        private String  translationStatus;// 是否转译 0：未转译 1：转义中 2：转译完成
        private String  kdxf_result;//科大讯飞转译结果
    }

