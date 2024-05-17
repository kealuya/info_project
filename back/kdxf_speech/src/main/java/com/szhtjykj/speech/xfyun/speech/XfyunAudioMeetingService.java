package com.szhtjykj.speech.xfyun.speech;

import cn.hutool.json.JSONUtil;
import com.google.gson.Gson;
import com.szhtjykj.speech.dao.KdxfSpeechDao;
import com.szhtjykj.speech.dao.MeetingDao;
import com.szhtjykj.speech.dao.MeetingFileDao;
import com.szhtjykj.speech.model.KdxfSpeech;
import com.szhtjykj.speech.model.MeetingFile;
import com.szhtjykj.speech.xfyun.speech.sign.LfasrSignature;
import com.szhtjykj.speech.xfyun.speech.utils.HttpUtil;
import org.apache.commons.lang.StringEscapeUtils;
import org.beetl.sql.core.SQLReady;
import org.beetl.sql.solon.annotation.Db;
import org.noear.solon.annotation.Component;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.*;
import java.net.URL;
import java.net.URLConnection;
import java.security.SignatureException;
import java.util.Date;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

/**
 * 语音会议文件处理 新添加 根据meetingId处理
 * @program: kdxf_speech
 * @description:
 * @packagename: com.szhtjykj.speech.xfyun.speech
 * @author: zhanbaohua
 * @date: 2024-05-15 13:50
 **/
@Component
public class XfyunAudioMeetingService {

    private static final String HOST = "";
    private static final String appid = "";
    private static final String keySecret = "";

    private static final Gson gson = new Gson();
    static Logger log = LoggerFactory.getLogger(XfyunSpeechService.class);


    private static ExecutorService executor = Executors.newFixedThreadPool(10); // 创建一个固定大小为10的线程池
    @Db
    MeetingFileDao meetingFileDao;
    @Db
    KdxfSpeechDao kdxfSpeechDao;
    @Db
    MeetingDao meetingDao;


    /**
     *
     * 改为下面的，自查库的方法 入参为meetingId
     * @param meetingId
     * @return
     * @throws SignatureException
     * @throws IOException
     */
    public String audioFileTranslation(String meetingId) throws Exception {
        //SDK 调用map
        HashMap<String, Object> map = new HashMap<>(16);
        //1 根据会议ID，得到会议文件list  从而得到语音网络地址
        List<MeetingFile> fileList = meetingFileDao.getMeetingFileByMeetingId(meetingId);

        for (MeetingFile fileDto : fileList){
            if (!fileDto.getFileType().equals("mp3")) {
                return "文件类型不正确";
            }
            String fileUrl = fileDto.getFileUrl();

            try {
                URL file_Url = new URL(fileUrl);
                URLConnection connection = file_Url.openConnection();
                InputStream inputStream = connection.getInputStream();
                long fileSize = inputStream.available();
                map.put("appId", appid);
                map.put("fileSize", fileSize);
                map.put("fileName", fileDto.getFileName());
                map.put("duration", "200");// 音频真实时长.当前未验证，可随机传一个数字
                LfasrSignature lfasrSignature = new LfasrSignature(appid, keySecret);
                map.put("signa", lfasrSignature.getSigna());
                map.put("ts", lfasrSignature.getTs());
//          此处后续追加 mp3 判断，需要追加接口参数 lame ：：暂时不考虑，因为普通语音转写不需要格式区别。
                String paramString = HttpUtil.parseMapToPathParam(map);
                log.info("upload paramString__kdxf调用map参数:" + paramString);
                String url = HOST + "/v2/api/upload" + "?" + paramString;
                log.info("upload_url__kdxf调用地址:" + url);
                String response = HttpUtil.iflyrecUpload(url, inputStream);

                log.info("upload response__kdxf返回值:" + response);
                String jsonStr = StringEscapeUtils.unescapeJavaScript(response);
                String orderId = String.valueOf(JSONUtil.getByPath(JSONUtil.parse(jsonStr), "content.orderId"));

                //修改会议转译状态为1转译中
                meetingDao.updateMeetingTranslationStatus(fileDto.getMeetingId(), "1");
                //修改文件转译状态为1转译中,并添加orderId
                meetingFileDao.updateMeetingFileTranslationStatus(orderId, "1", fileDto.getId());
                //存储kdxf_speech表
                KdxfSpeech ks = new KdxfSpeech();
                ks.setOrder_id(orderId);
                ks.setState(0); // 进行中
                ks.setMeetingId(meetingId);
                ks.setFile_name(fileDto.getFileName());
                ks.setDatetime(new Date());
                ks.setFileId(fileDto.getId());
                kdxfSpeechDao.insert(ks);

//                // 提交任务给线程池执行
//                executor.submit(() -> {
//                    // 处理文件的逻辑
//                    try {
//                        //orderId 当前orderId  ,orderId2 第一个文件的orderId
//                        getResult(orderId, fileDto.getFileName(), "", fileDto.getMeetingId());
//                    } catch (Exception e) {
//                        log.error("获取音频翻译发生错误::", e);
//
//                        // 将语音转换信息存入db
//                        KdxfSpeech ksForUpdate = new KdxfSpeech();
//                        ksForUpdate.setOrder_id(orderId);
//                        ksForUpdate.setState(2); // 错误
//                        ksForUpdate.setMeetingId(meetingId);
//                        ksForUpdate.setComment(e.getMessage());
//                        kdxfSpeechDao.updateById(ksForUpdate);
//                    }
//                });


//                // 关闭inputStream
                inputStream.close();
            } catch (IOException e) {
                e.printStackTrace();
                log.error("获取音频翻译发生错误::", e);
            }
        }
        return "接口调用完成";
    }


    /**
     * 根据会议ID，得到转译的返回值
     * @param meetingId
     * @throws Exception
     */
    public void getResultByMeetingId(String meetingId) throws Exception {

        //是否为最后一个文件，且转译已经完成
        Boolean flag = false;

        //根据会议ID，得到speech  list
        List<KdxfSpeech> kdxfSpeechList = kdxfSpeechDao.getSpeechByMeetingId(meetingId);

        //循环获取转译结果
        for (KdxfSpeech kdxfSpeech : kdxfSpeechList){
            HashMap<String, Object> map = new HashMap<>(16);
            // 此处的orderid 是 实际请求音频的订单，需要通过该订单获取新上传音频的翻译
            map.put("orderId", kdxfSpeech.getOrder_id());
            LfasrSignature lfasrSignature = new LfasrSignature(appid, keySecret);
            map.put("signa", lfasrSignature.getSigna());
            map.put("ts", lfasrSignature.getTs());
            map.put("appId", appid);
            map.put("resultType", "transfer,predict");
            String paramString = HttpUtil.parseMapToPathParam(map);
            String url = HOST + "/v2/api/getResult" + "?" + paramString;
            log.info("获取转译返回值的url:" + url);

            String response = HttpUtil.iflyrecGet(url);
            JsonParse jsonParse = gson.fromJson(response, JsonParse.class);
            try {
//                Object j = jsonParse.content.orderInfo.status;
                if (jsonParse.content.orderInfo.status == 4 || jsonParse.content.orderInfo.status == -1) {
                    log.info("订单完成:" + kdxfSpeech.getOrder_id() + "::" + getContentForNormal(jsonParse.content.orderResult));
                    KdxfSpeech ksForUpdate = kdxfSpeech;
                    ksForUpdate.setState(1); // 完成
                    ksForUpdate.setContent(getContentForNormal(jsonParse.content.orderResult));
                    ksForUpdate.setReal_duration(jsonParse.content.orderInfo.realDuration);
                    //更新数据
                    kdxfSpeechDao.updateTemplateById(ksForUpdate);
                }
            } catch (Exception e) {
                log.warn("getResult接口调用错误::" + response, e);
                continue;
            }
        }
    }

    /**
     * 根据会议ID,查看会议中的音频是否都转译完成并完成最后的拼接
     * @param meetingId
     * @throws Exception
     */
    public void getLastFileByMeetingId(String meetingId) throws Exception {
        //根据会议ID，得到speech  list
        List<KdxfSpeech> kdxfSpeechList = kdxfSpeechDao.getSpeechByMeetingId(meetingId);
        //查看最后一个文件是否处理完成
        KdxfSpeech ks = kdxfSpeechList.get(kdxfSpeechList.size()-1);
        if(ks.getState() == 1){
            // 根据fileId 将合成后的content 拼接到第一个文件中
            String content = "";
            for (KdxfSpeech kdxfSpeech : kdxfSpeechList){
                content = content+kdxfSpeech.getContent();
            }
            kdxfSpeechList.get(0).setContent(content);
            //更新数据 更新kdxf_speech
            kdxfSpeechDao.updateTemplateById(kdxfSpeechList.get(0));
            //则修改会议的转译状态为2完成
            //更新会议表状态
            meetingDao.updateMeetingTranslationStatus(kdxfSpeechList.get(0).getMeetingId(), "2");
        }
    }




    class JsonParse {
        Content content;
    }

    class Content {
        OrderInfo orderInfo;
        String orderResult;
    }

    class OrderInfo {
        Integer status;
        Long realDuration;
    }


    /**
     * 得到content
     * @param sb
     * @return
     */
    public static String getContentForNormal(String sb) {
        Result jsonParse = gson.fromJson(sb.toString(), Result.class);
        List<Lattice> latticeList = jsonParse.lattice;
        String allResult = "";
        for (int i = 0; i < latticeList.size(); i++) {
            Lattice tempLattice = latticeList.get(i);
            Json_1best json_1best = gson.fromJson(tempLattice.json_1best, Json_1best.class);
            String rl = json_1best.st.rl;
            List<Rt> rtList = json_1best.st.rt;
            for (int j = 0; j < rtList.size(); j++) {
                Rt tempRt = rtList.get(j);
                List<Ws> wsList = tempRt.ws;
                for (int k = 0; k < wsList.size(); k++) {
                    Ws tempWs = wsList.get(k);
                    List<Cw> cwList = tempWs.cw;
                    for (int l = 0; l < cwList.size(); l++) {
                        Cw tempCw = cwList.get(l);
                        System.out.print(tempCw.w);
                        allResult = allResult + tempCw.w;
                    }
                }
            }
        }

        return allResult;
    }
}






