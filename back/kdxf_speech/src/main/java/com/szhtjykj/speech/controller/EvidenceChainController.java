package com.szhtjykj.speech.controller;

import com.google.gson.Gson;
import com.szhtjykj.speech.generatePdf.evidenceChain.EvidenceChainService;
import com.szhtjykj.speech.generatePdf.evidenceChain.dto.PlanDto;
import org.noear.solon.annotation.Controller;
import org.noear.solon.annotation.Inject;
import org.noear.solon.annotation.Mapping;
import org.noear.solon.annotation.Post;
import org.noear.solon.core.handle.Context;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/**
 * 证据链
 * @program: kdxf_speech
 * @description:
 * @packagename: com.szhtjykj.speech.controller
 * @author: zhanbaohua
 * @date: 2024-05-30 08:56
 **/
@Controller
public class EvidenceChainController {

    static Logger log = LoggerFactory.getLogger(EvidenceChainController.class);
    private static final Gson gson = new Gson();

    @Inject
    EvidenceChainService chainService;

    @Post
    @Mapping("/generateEvidenceChainExcel")
    public Map generateEvidenceChainExcel(Context ctx) throws IOException {

        // 获取JSON入参
        String paramJsonStr = ctx.body();
        PlanDto paramDto = gson.fromJson(paramJsonStr,PlanDto.class);

        String orderId = "";
        //返回参数map
        Map<String, Object> returnMap = new HashMap<>();
        try {
            orderId = chainService.GenerateEvidenceChainExcel(paramDto);
        } catch (Exception e) {
            log.error("生成证据链excel发生错误::", e);

            returnMap.put("success", false);
            returnMap.put("msg", e.toString());

            return returnMap;
        }
        returnMap.put("success", true);
        returnMap.put("data", orderId);
        returnMap.put("msg", "excel文件生成成功");
        return returnMap;
    }



    @Post
    @Mapping("/generatePdf")
    public Map generatePdf(Context ctx) throws IOException {

        // 获取JSON入参
        String paramJsonStr = ctx.body();
        PlanDto paramDto = gson.fromJson(paramJsonStr,PlanDto.class);

        //返回参数map
        Map<String, Object> returnMap = new HashMap<>();
        try {
            String orderId = chainService.GenerateEvidenceChainPdf(paramDto);
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
