package com.szhtjykj.speech.dao;

import org.beetl.sql.mapper.BaseMapper;
import org.beetl.sql.mapper.annotation.Template;
import com.szhtjykj.speech.model.MeetingFile;
import org.beetl.sql.mapper.annotation.Update;

import java.util.List;


/**
 * @program: kdxf_speech
 * @description:
 * @packagename: com.szhtjykj.speech.dao
 * @author: zhanbaohua
 * @date: 2024-05-15 13:59
 **/
public interface MeetingFileDao extends BaseMapper<MeetingFile> {
    @Template("select * from kdxf_meetingFile where meetingId = #{meetingId} order by id")
    List<MeetingFile> getMeetingFileByMeetingId(String meetingId);

    @Update
    @Template("update kdxf_meetingFile set translationStatus = '1',orderId = #{orderId} where id = #{id}")
    Integer updateMeetingFileTranslationStatus(String orderId, String status, int id);



}
