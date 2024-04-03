package com.szhtjykj.speech.dao;

import com.szhtjykj.speech.model.KdxfSpeech;
import org.beetl.sql.mapper.BaseMapper;
import org.beetl.sql.mapper.annotation.Template;

import java.util.List;

/**
 * @author
 */
public interface KdxfSpeechDao extends BaseMapper<KdxfSpeech> {

    @Template("select content from kdxf_speech where order_id = #{orderId} order by no")
    List<KdxfSpeech> getSpeechByOrderId(String orderId);
}