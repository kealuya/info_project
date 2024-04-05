package com.szhtjykj.speech.xfyun.knowledge;

import com.szhtjykj.speech.xfyun.knowledge.dto.UploadResp;
import com.szhtjykj.speech.xfyun.knowledge.util.ChatDocUtil;

import java.io.IOException;
import java.util.concurrent.CompletableFuture;

/**
 * Test
 * 详细接口文档请查看 <a href="https://chatdoc.xfyun.cn/docs">文档知识库API</a>
 *
 * @author ydwang16
 * @version 2023/09/06 13:59
 **/
public class Main {
    private static final String uploadUrl = "https://chatdoc.xfyun.cn/openapi/fileUpload";
    private static final String fileStatusUrl = "https://chatdoc.xfyun.cn/openapi/fileStatus";
    private static final String chatUrl = "wss://chatdoc.xfyun.cn/openapi/chat";
    private static final String startSummaryUrl = "https://chatdoc.xfyun.cn/openapi/startSummary";
    private static final String fileSummaryUrl = "https://chatdoc.xfyun.cn/openapi/fileSummary";
    private static final String appId = "60f7d982";
    private static final String secret = "NGM3ZDYwMzdiNWRhMTIyMTg0OWZkNzUw";

    public static void main(String[] args) throws IOException {
//        test1();
        ChatDocUtil chatDocUtil = new ChatDocUtil();
//        String file = "8738c347720640fbb32d201ef76e1e16";
//         1、上传
//        UploadResp uploadResp = chatDocUtil.upload("C:\\Users\\kealuya\\myCoding\\info_project\\back\\kdxf_speech\\src\\main\\output\\knowledge.txt", uploadUrl, appId, secret);
//        System.out.println("请求sid=" + uploadResp.getSid());
//        System.out.println("文件id=" + uploadResp.getData().getFileId());

        // 2、问答，上传文件状态为vectored时才可以问答，文件状态可以调用【文档状态查询】接口查询

        String fileId = "dc20dceba1794998b00af3e4d896a077";
            String question = "请按照会议纪要的格式返回详细的会议内容，并且不需要其他描述语言，只需要会议纪要";
        CompletableFuture<String> cf=  chatDocUtil.chat(chatUrl, fileId, question, appId, secret);
        System.out.print("回答內容::");
        System.out.println(cf.join());

//        String file = uploadResp.getData().getFileId();


//        UploadResp ur1 = chatDocUtil.summary(startSummaryUrl, file, appId, secret);
//        System.out.print("startSummaryUrl::");
//        System.out.println(ur1);


//        UploadResp ur2 = chatDocUtil.summary(fileSummaryUrl, file, appId, secret);
//        System.out.print("fileSummaryUrl::");
//        System.out.println(ur2.getData());
/*
{"flag":true,"code":0,"desc":null,"data":{"summaryStatus":"summarying"},"sid":"56c93cac012d48a68fb2307fa726fb48"}
 */
/*
{"flag":false,"code":-1,"desc":"当前文件状态不可做总结","data":null,"sid":"104e871fdb77405aa9c0f089a806cf32"}
 */




    }
}