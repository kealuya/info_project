package com.szhtjykj.speech.xfyun.knowledge;

import com.szhtjykj.speech.dao.KdxfKnowledgeDao;
import com.szhtjykj.speech.dao.KdxfSpeechDao;
import com.szhtjykj.speech.model.KdxfKnowledge;
import com.szhtjykj.speech.model.KdxfSpeech;
import com.szhtjykj.speech.xfyun.knowledge.dto.UploadResp;
import com.szhtjykj.speech.xfyun.knowledge.util.ChatDocUtil;
import org.beetl.sql.solon.annotation.Db;
import org.noear.solon.Solon;
import org.noear.solon.annotation.Component;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.*;
import java.net.URL;
import java.util.Date;
import java.util.List;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

@Component
public class XfyunKnowledgeService {



    public static final String QUESTION_MEETING_MUNIT = "请按照会议纪要的格式返回详细的会议内容，并且不需要其他描述语言，只需要会议纪要";
    //    public static final String QUESTION_MEETING_MUNIT = "请按照会议纪要的格式返回详细的会议内容，并且不需要其他描述语言，只需要会议纪要，" +
//        "纪要按照 '1. XXX\n" +
//        "- XXXXXXX\n" +
//        "- XXXXXXX\n"+ " 2. XXX\n" +
//            "- XXXXXXX\n'" +" 的形式返回。生成之后，再生成一份一样内容的markdown格式文本，两段文字中间用========分隔开";
    @Db
    KdxfSpeechDao kdxfSpeechDao;
    @Db
    KdxfKnowledgeDao kdxfKnowledgeDao;
    static Logger log = LoggerFactory.getLogger(XfyunKnowledgeService.class);

    static {
        String AUDIO_FILE_PATH = Solon.cfg().get("server.audio.save.path");
    }

    private static final ExecutorService executor = Executors.newFixedThreadPool(10); // 创建一个固定大小为10的线程池

    public void makeSummaryAndMeeting(String orderId, String fileId) throws Exception {
        List<KdxfSpeech> speechList =  kdxfSpeechDao.getSpeechDetailByOrderId(orderId);
        String meetingId = speechList.get(0).getMeetingId();
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

                    if (ur.isFlag() && ur.getCode() == 0) {
                        log.info("orderId::" + orderId + "  fileId::" + fileId + "::进行中...，状态为:提交文档总结成功...");
                        break;
                    }
                }

                while (true) {

                    UploadResp ur = chatDocUtil.summary(fileSummaryUrl, fileId, appId, secret);

                    if (ur.isFlag() && "unsummary".equals(ur.getData().getSummaryStatus())) {
                        log.info("orderId::" + orderId + "  fileId::" + fileId + "::进行中...，状态为:文档没有被申请总结...");
                        //建议使用回调的方式查询结果，查询接口有请求频率限制
                        Thread.sleep(5000);
                        continue;
                    }
                    if (ur.isFlag() && "summarying".equals(ur.getData().getSummaryStatus())) {
                        log.info("orderId::" + orderId + "  fileId::" + fileId + "::进行中...，状态为:文档总结中...");
                        //建议使用回调的方式查询结果，查询接口有请求频率限制
                        Thread.sleep(5000);
                        continue;
                    }

                    if (ur.isFlag() && "done".equals(ur.getData().getSummaryStatus())) {
                        log.info("orderId::" + orderId + "  fileId::" + fileId + "::处理完成，状态为:文档总结完成...");

                        KdxfKnowledge knowledge = new KdxfKnowledge();
                        knowledge.setOrder_id(orderId);
                        knowledge.setFile_id(fileId);
                        knowledge.setState(1);
                        knowledge.setSummary(ur.getData().getSummary());
                        knowledge.setMeetingId(meetingId);
                        kdxfKnowledgeDao.updateTemplateById(knowledge);
                        return;
                    }
                }
            } catch (Exception e) {
                log.error("获取文档总结发生错误::", e);

                KdxfKnowledge knowledge = new KdxfKnowledge();
                knowledge.setOrder_id(orderId);
                knowledge.setFile_id(fileId);
                knowledge.setState(2);
                knowledge.setComment(e.toString());
                kdxfKnowledgeDao.updateTemplateById(knowledge);
            }
        });
    }


    public String uploadFile(String orderId,String meetingId) throws Exception {

        List<KdxfSpeech> ksList = kdxfSpeechDao.getSpeechByOrderId(orderId);
        if (ksList.isEmpty()) {
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
                knowledge.setMeetingId(meetingId);
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


    /**
     * 文件信息 上传
     * @param fileUrl
     * @return
     * @throws Exception
     */
    public String uploadDocumentFile(String fileUrl,String meetingId) throws Exception {

            // 创建一个指向网络资源的URL
            URL url = new URL(fileUrl);
            // 打开到此 URL 引用的资源的连接
            InputStream inputStream = url.openStream();

            // 创建临时文件
            File tempFile = null;
            try {
                tempFile = File.createTempFile("tempfile" + System.currentTimeMillis(), ".txt");
                copyInputStreamToFile(inputStream,tempFile);

                UploadResp uploadResp = chatDocUtil.upload(tempFile.getAbsolutePath(), uploadUrl, appId, secret);

                if (uploadResp.isFlag() && uploadResp.getCode() == 0) {


                    KdxfKnowledge knowledge = new KdxfKnowledge();
                    knowledge.setFile_id(uploadResp.getData().getFileId());
                    knowledge.setDatetime(new Date());
                    knowledge.setOrder_id(meetingId);
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
                inputStream.close();
            }
    }


    // 流转文件   InputStream -> File
    private static void copyInputStreamToFile(InputStream inputStream, File file) throws IOException {

        try (FileOutputStream outputStream = new FileOutputStream(file)) {

            int read;
            byte[] bytes = new byte[1024];

            while ((read = inputStream.read(bytes)) != -1) {
                outputStream.write(bytes, 0, read);
            }

        }

    }




    public String makeMeetingMinuteByOrderId(String orderId, String question) {

        KdxfKnowledge kk = kdxfKnowledgeDao.unique(orderId);
        CompletableFuture<String> cf = chatDocUtil.chat(chatUrl, kk.getFile_id(), question, appId, secret);
        String mm = cf.join();
        kk.setMeeting_minutes(mm);
        System.out.println(mm);

        kdxfKnowledgeDao.updateTemplateById(kk);

        return mm;
    }


    public String makeBrainMapByOrderId(String orderId) {

        KdxfKnowledge kk = kdxfKnowledgeDao.unique(orderId);
        String mm = kk.getMeeting_minutes();
        String[] mms = mm.split("\n");
        for (int i = 0; i < mms.length; i++) {
            if (i == 0) {
                mms[0] = ("# " + mms[0].trim()).replace(":", "");
                continue;
            }
            mms[i] = mms[i].trim();
            if (mms[i].matches("^\\d\\..*")) {
                // 定义正则表达式
                String regex = "(^\\d\\. ?)";
                // 编译正则表达式
                Pattern pattern = Pattern.compile(regex);
                // 创建 Matcher 对象
                Matcher matcher = pattern.matcher(mms[i]);
                // 进行替换
                String replacedString = matcher.replaceAll("## ");
                mms[i] = replacedString;
            }
        }
        String mdString = String.join("\n", mms);

        kk.setBrain(mdString);
        System.out.println(mdString);

        kdxfKnowledgeDao.updateTemplateById(kk);

        return mdString;
    }


}
