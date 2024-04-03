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


//    //文件上传
//    @Post
//    @Mapping("/uploadAudio")
//    public Map uploadAudio(Context ctx) {
//        String orderId = ctx.param("orderId");
//        Map<String, Object> returnMap = new HashMap<>();
//        try {
//            String fileId = xfyunKnowledgeService.uploadFile(orderId);
//
//
//        } catch (Exception e) {
//            log.error("发生错误::", e);
//
//            returnMap.put("success", "false");
//            returnMap.put("msg", e.toString());
//
//            return returnMap;
//        }
//    }


}
