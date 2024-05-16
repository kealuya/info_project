package com.szhtjykj.speech.controller;

import com.google.gson.Gson;
import com.szhtjykj.speech.xfyun.speech.XfyunAudioMeetingService;
import org.noear.solon.Solon;
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
 * @date: 2024-05-15 13:37
 **/

@Controller
public class AudioMeetingController {
    static Logger log = LoggerFactory.getLogger(AudioMeetingController.class);
    private static String AUDIO_FILE_PATH;
    private static final Gson gson = new Gson();
    static {
        AUDIO_FILE_PATH = Solon.cfg().get("server.audio.save.path");
    }

    @Inject
    XfyunAudioMeetingService xfyunAudioMeetingService;

    /**
     * 语音文件转义
     * @param ctx
     * @return
     */
    @Post
    @Mapping("/audioMeetingTranslation")
    public Map audioMeetingTranslation(Context ctx) {
        String meetingId = ctx.param("meetingId");
        log.info("接收到的meetingId:"+meetingId);
        //返回参数map
        Map<String, Object> returnMap = new HashMap<>();
        try {
            String orderId = xfyunAudioMeetingService.audioFileTranslation(meetingId);
        } catch (Exception e) {
            log.error("音频上传发生错误::", e);

            returnMap.put("success", false);
            returnMap.put("msg", e.toString());

            return returnMap;
        }
        returnMap.put("success", true);
        returnMap.put("msg", "");
//        returnMap.put("orderId", orderId);
//        returnMap.put("fileName", tempFileName);
        return returnMap;
    }


    /**
     * 这个solon查了两个小时，没有实现定时任务，后面再看吧
     *
     * 定时任务，执行查询转译结果的方法
     * @param ctx
     * @return
     */
    @Post
    @Mapping("/getResultByMeetingId")
    public Map getResultByMeetingId(Context ctx) {
        String meetingId = ctx.param("meetingId");
        log.info("接收到的meetingId:"+meetingId);
        //返回参数map
        Map<String, Object> returnMap = new HashMap<>();
        try {
            xfyunAudioMeetingService.getResultByMeetingId(meetingId);
        } catch (Exception e) {
            log.error("定时方法执行::", e);
            returnMap.put("success", false);
            returnMap.put("msg", e.toString());
            return returnMap;
        }
        returnMap.put("success", true);
        returnMap.put("msg", "调用成功");
        return returnMap;
    }



}
