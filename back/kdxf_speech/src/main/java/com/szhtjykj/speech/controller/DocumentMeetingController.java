package com.szhtjykj.speech.controller;

import com.google.gson.Gson;
import com.szhtjykj.speech.xfyun.speech.XfyunDocumentMeetingService;
import org.noear.solon.annotation.Controller;
import org.noear.solon.annotation.Inject;
import org.noear.solon.annotation.Mapping;
import org.noear.solon.annotation.Post;
import org.noear.solon.core.handle.Context;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.HashMap;
import java.util.Map;

/**
 * @program: kdxf_speech
 * @description:
 * @packagename: com.szhtjykj.speech.controller
 * @author: zhanbaohua
 * @date: 2024-05-18 10:30
 **/
@Controller
public class DocumentMeetingController {

    static Logger log = LoggerFactory.getLogger(DocumentMeetingController.class);
    private static final Gson gson = new Gson();

    @Inject
    XfyunDocumentMeetingService xfyunDocumentMeetingService;

    @Post
    @Mapping("/documentMeetingTranslation")
    public Map documentMeetingTranslation(Context ctx) {
        String meetingId = ctx.param("meetingId");
        log.info("接收到的meetingId:"+meetingId);
        //返回参数map
        Map<String, Object> returnMap = new HashMap<>();
        try {
            String orderId = xfyunDocumentMeetingService.documentFileTranslation(meetingId);
        } catch (Exception e) {
            log.error("文档写入数据库发生错误::", e);

            returnMap.put("success", false);
            returnMap.put("msg", e.toString());

            return returnMap;
        }
        returnMap.put("success", true);
        returnMap.put("msg", "");
        return returnMap;
    }
}
