package com.szhtjykj.speech.xfyun.knowledge;

import com.google.gson.Gson;
import com.szhtjykj.speech.dao.KdxfKnowledgeDao;
import com.szhtjykj.speech.dao.KdxfSpeechDao;
import com.szhtjykj.speech.model.KdxfKnowledge;
import com.szhtjykj.speech.model.KdxfSpeech;
import com.szhtjykj.speech.xfyun.knowledge.dto.UploadResp;
import com.szhtjykj.speech.xfyun.knowledge.util.ChatDocUtil;
import com.szhtjykj.speech.xfyun.speech.XfyunSpeechService;
import org.beetl.sql.solon.annotation.Db;
import org.noear.solon.Solon;
import org.noear.solon.annotation.Component;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.util.Date;
import java.util.List;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

@Component
public class XfyunKnowledgeService {

    private static final String uploadUrl = "https://chatdoc.xfyun.cn/openapi/fileUpload";
    //    private static final String fileStatusUrl = "https://chatdoc.xfyun.cn/openapi/fileStatus";
    private static final String chatUrl = "wss://chatdoc.xfyun.cn/openapi/chat";
    private static final String startSummaryUrl = "https://chatdoc.xfyun.cn/openapi/startSummary";
    private static final String fileSummaryUrl = "https://chatdoc.xfyun.cn/openapi/fileSummary";

    private static final String appId = "60f7d982";
    private static final String secret = "NGM3ZDYwMzdiNWRhMTIyMTg0OWZkNzUw";
    private static final ChatDocUtil chatDocUtil = new ChatDocUtil();
    @Db
    KdxfSpeechDao kdxfSpeechDao;
    @Db
    KdxfKnowledgeDao kdxfKnowledgeDao;
    private static String AUDIO_FILE_PATH;
    static Logger log = LoggerFactory.getLogger(XfyunKnowledgeService.class);

    static {
        AUDIO_FILE_PATH = Solon.cfg().get("server.audio.save.path");
    }

    private static ExecutorService executor = Executors.newFixedThreadPool(10); // 创建一个固定大小为10的线程池

    public void makeSummaryAndMeeting(String orderId, String fileId) throws Exception {


        // 提交任务给线程池执行
        executor.submit(() -> {
            // 处理文件的逻辑
            try {

                while (true) {

                    UploadResp ur = chatDocUtil.summary(startSummaryUrl, fileId, appId, secret);
                    if (!ur.isFlag() && ur.getCode() == -1) {
                        log.info("orderId::" + orderId + "  fileId::" + fileId + "::进行中...，状态为:向量化中...");
                        //建议使用回调的方式查询结果，查询接口有请求频率限制
                        Thread.sleep(5000);
                        continue;
                    }

                    if (ur.isFlag() && ur.getData().getSummaryStatus() == "summarying") {
                        log.info("orderId::" + orderId + "  fileId::" + fileId + "::进行中...，状态为:文档总结中...");
                        //建议使用回调的方式查询结果，查询接口有请求频率限制
                        Thread.sleep(5000);
                        continue;
                    }

                    if (ur.isFlag() && ur.getData().getSummaryStatus() == "done") {
                        log.info("orderId::" + orderId + "  fileId::" + fileId + "::处理完成，状态为:文档总结完成...");

                        KdxfKnowledge knowledge = new KdxfKnowledge();
                        knowledge.setOrder_id(orderId);
                        knowledge.setState(1);
                        knowledge.setSummary(ur.getData().getSummary());
                        kdxfKnowledgeDao.updateTemplateById(knowledge);

                        return;
                    }


                }


            } catch (Exception e) {
                log.error("获取音频翻译发生错误::", e);

                // 将语音转换信息存入db
                KdxfSpeech ksForUpdate = new KdxfSpeech();
                ksForUpdate.setOrder_id(orderId);
                ksForUpdate.setState(2); // 错误
                ksForUpdate.setComment(e.getMessage());
                kdxfSpeechDao.updateById(ksForUpdate);
            }
        });


    }


    public String uploadFile(String orderId) throws Exception {

        List<KdxfSpeech> ksList = kdxfSpeechDao.getSpeechByOrderId(orderId);
        if (ksList.size() == 0) {
            throw new Exception("orderId检索不到数据::" + orderId);
        }
        StringBuffer sb = new StringBuffer();
        for (KdxfSpeech ks : ksList) {
            sb.append(ks.getContent());
        }
        // 创建临时文件
        File tempFile = null;
        try {
            tempFile = File.createTempFile("tempfile" + System.currentTimeMillis(), ".txt");

            // 写入内容到临时文件
            FileWriter writer = new FileWriter(tempFile);
            writer.write(sb.toString());
            writer.close();

            UploadResp uploadResp = chatDocUtil.upload(tempFile.getAbsolutePath(), uploadUrl, appId, secret);

            if (uploadResp.isFlag() && uploadResp.getCode() == 0) {


                KdxfKnowledge knowledge = new KdxfKnowledge();
                knowledge.setFile_id(uploadResp.getData().getFileId());
                knowledge.setDatetime(new Date());
                knowledge.setOrder_id(orderId);
                knowledge.setState(0);
                kdxfKnowledgeDao.insert(knowledge);

                return uploadResp.getData().getFileId();
            } else {
                throw new Exception("请求接口错误::" + uploadResp.getCode() + "::" + uploadResp.getDesc());
            }

        } catch (Exception e) {
            throw new Exception(e);
        } finally {
            // 删除临时文件
            if (tempFile.exists()) {
                tempFile.delete();
            }
        }
    }

}
