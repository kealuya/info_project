package com.szhtjykj.speech.controller;

import com.szhtjykj.speech.xfyun.knowledge.XfyunKnowledgeService;
import com.szhtjykj.speech.xfyun.speech.XfyunSpeechService;
import org.noear.solon.annotation.Controller;
import org.noear.solon.annotation.Inject;
import org.noear.solon.annotation.Mapping;
import org.noear.solon.annotation.Post;
import org.noear.solon.core.handle.Context;
import org.noear.solon.core.handle.UploadedFile;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

@Controller
public class KnowledgeController {
    static Logger log = LoggerFactory.getLogger(KnowledgeController.class);


    @Inject
    XfyunKnowledgeService xfyunKnowledgeService;


    //通过获取已经翻译成文字的会议内容，生成概要
    @Post
    @Mapping("/convertAudioToSummary")
    public Map convertAudioToSummary(Context ctx) {
        String orderId = ctx.param("orderId");
        Map<String, Object> returnMap = new HashMap<>();
        try {
            String fileId = xfyunKnowledgeService.uploadFile(orderId);
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


}
