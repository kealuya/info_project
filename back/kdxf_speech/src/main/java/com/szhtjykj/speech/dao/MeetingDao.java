package com.szhtjykj.speech.dao;

import com.szhtjykj.speech.model.Meeting;
import com.szhtjykj.speech.model.MeetingFile;
import org.beetl.sql.mapper.BaseMapper;
import org.beetl.sql.mapper.annotation.Template;
import org.beetl.sql.mapper.annotation.Update;

import java.util.List;

/**
 * @program: kdxf_speech
 * @description:
 * @packagename: com.szhtjykj.speech.dao
 * @author: zhanbaohua
 * @date: 2024-05-16 10:44
 **/
public interface MeetingDao extends BaseMapper<Meeting> {
    @Update
    @Template("update  kdxf_meeting set translationStatus = #{status} where meetingId = #{meetingId}")
    Integer updateMeetingTranslationStatus(String meetingId,String status);


}
