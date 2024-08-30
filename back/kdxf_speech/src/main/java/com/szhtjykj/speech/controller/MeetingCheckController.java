package com.szhtjykj.speech.controller;

import com.szhtjykj.speech.xfyun.meetingCheck.XfyunMeetingCheckService;
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
 * PC 端 会议检测
 * @program: kdxf_speech
 * @description:
 * @packagename: com.szhtjykj.speech.controller
 * @author: zhanbaohua
 * @date: 2024-06-18 14:39
 **/

@Controller
public class MeetingCheckController {

    static Logger log = LoggerFactory.getLogger(MeetingCheckController.class);

    @Inject
    XfyunMeetingCheckService xfyunMeetingCheckService;

    /**
     * 会议录音  文件检测
     * @param ctx
     * @return
     */
    @Post
    @Mapping("/meetingAudioCheck")
    public Map meetingAudioCheck(Context ctx) {
        String meetingId = ctx.param("meetingId");
        String corpCode = ctx.param("corpCode");
        Map<String, Object> returnMap = new HashMap<>();
        boolean flag =  xfyunMeetingCheckService.meetingAudioCheck(meetingId,corpCode);

        if(flag){
            returnMap.put("success", true);
            returnMap.put("msg","分析成功" );
        }else{
            returnMap.put("success", false);
            returnMap.put("msg","分析失败" );
        }
        return returnMap;
    }

}
