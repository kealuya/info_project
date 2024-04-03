package com.szhtjykj.speech.model;

import lombok.Data;
import org.beetl.sql.annotation.entity.AssignID;
import org.beetl.sql.annotation.entity.Table;

import java.util.Date;

@Data
@Table(name = "kdxf_knowledge")
public class KdxfKnowledge {
    @AssignID
    private String order_id;
    private String summary;
    private Integer state;
    private Date datetime;
    private String meeting_minutes;
    private String comment;
    private String brain;
    private String chat1;
    private String chat2;
    private String file_id;
}
