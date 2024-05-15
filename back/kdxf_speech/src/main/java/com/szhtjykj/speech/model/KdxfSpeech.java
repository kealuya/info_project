package com.szhtjykj.speech.model;

import lombok.Data;
import org.beetl.sql.annotation.entity.AssignID;
import org.beetl.sql.annotation.entity.Table;

import java.util.Date;

@Data
@Table(name = "kdxf_speech")
public class KdxfSpeech {
    @AssignID
    private String order_id;
    @AssignID
    private int no;
    private String file_name;
    private Integer state;
    private Date datetime;
    private String content;
    private String comment;
    private Long real_duration;
    //添加字段
    private String meetingId; //会议ID
    private String fileId;  //文件ID
}