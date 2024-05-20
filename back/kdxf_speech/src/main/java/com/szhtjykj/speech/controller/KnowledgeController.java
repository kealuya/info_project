package com.szhtjykj.speech.controller;

import com.szhtjykj.speech.xfyun.knowledge.XfyunKnowledgeService;
import org.noear.solon.annotation.Controller;
import org.noear.solon.annotation.Inject;
import org.noear.solon.annotation.Mapping;
import org.noear.solon.annotation.Post;
import org.noear.solon.core.handle.Context;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.HashMap;
import java.util.Map;

import static com.szhtjykj.speech.xfyun.knowledge.XfyunKnowledgeService.QUESTION_MEETING_MUNIT;

@Controller
public class KnowledgeController {

    static Logger log = LoggerFactory.getLogger(KnowledgeController.class);


    @Inject
    XfyunKnowledgeService xfyunKnowledgeService;


    /**
     * 录音  文件 生成会议纪要
     * 通过获取已经翻译成文字的会议内容，生成概要
     * @param ctx
     * @return
     */
    @Post
    @Mapping("/convertAudioToSummary")
    public Map convertAudioToSummary(Context ctx) {
        String orderId = ctx.param("orderId");
        String meetingId = ctx.param("meetingId");
        Map<String, Object> returnMap = new HashMap<>();
        try {
            String fileId = xfyunKnowledgeService.uploadFile(orderId,meetingId);
            xfyunKnowledgeService.makeSummaryAndMeeting(orderId,fileId);

            returnMap.put("success", true);
            returnMap.put("msg", "申请成功");

            return returnMap;
        } catch (Exception e) {
            log.error("发生错误::", e);

            returnMap.put("success", false);
            returnMap.put("msg", e.toString());

            return returnMap;
        }
    }


    /**
     * 录音  文件 生成会议纪要
     * @param ctx
     * @return
     */
    @Post
    @Mapping("/convertAudioToMeetingMinute")
    public Map convertAudioToMeetingMinute(Context ctx) {
        String orderId = ctx.param("orderId");
        Map<String, Object> returnMap = new HashMap<>();
        String meetingMinute =  xfyunKnowledgeService.makeMeetingMinuteByOrderId(orderId,QUESTION_MEETING_MUNIT);

        returnMap.put("success", true);
        returnMap.put("msg","" );
        returnMap.put("meetingMinute",meetingMinute);

        return returnMap;
    }


    /**
     * 录音文件 生成脑图
     * @param ctx
     * @return
     */
    @Post
    @Mapping("/convertAudioToBrainMap")
    public Map convertAudioToBrainMap(Context ctx) {
        String orderId = ctx.param("orderId");
        Map<String, Object> returnMap = new HashMap<>();
        String brainMap = xfyunKnowledgeService.makeBrainMapByOrderId(orderId);

        returnMap.put("success", true);
        returnMap.put("msg","" );
        returnMap.put("brainMap",brainMap);

        return returnMap;
    }

    //============================================文档处理=========================

    /**
     * 通过文档内容，生成概要
     * @param
     * @return
     */
    @Post
    @Mapping("/convertDocumentToSummary")
    public Map convertDocumentToSummary(String meetingId,String fileUrl){

        Map<String, Object> returnMap = new HashMap<>();
        try {

            String fileId = xfyunKnowledgeService.uploadDocumentFile(fileUrl,meetingId);
            xfyunKnowledgeService.makeSummaryAndMeeting(meetingId,fileId);

            returnMap.put("fileId",fileId);
            returnMap.put("success", true);
            returnMap.put("msg", "申请成功");

            return returnMap;
        } catch (Exception e) {
            log.error("发生错误::", e);

            returnMap.put("success", false);
            returnMap.put("msg", e.toString());

            return returnMap;
        }
    }
}
