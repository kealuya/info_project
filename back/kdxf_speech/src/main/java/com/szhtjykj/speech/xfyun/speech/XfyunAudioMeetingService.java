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
            }
        }


    }




    public void getResult(String orderId, String fileName, String orderId2,String meetingId) throws Exception {
        HashMap<String, Object> map = new HashMap<>(16);
        // 此处的orderid 是 实际请求音频的订单，需要通过该订单获取新上传音频的翻译
        map.put("orderId", orderId);

        LfasrSignature lfasrSignature = new LfasrSignature(appid, keySecret);
        map.put("signa", lfasrSignature.getSigna());
        map.put("ts", lfasrSignature.getTs());
        map.put("appId", appid);
        map.put("resultType", "transfer,predict");
        String paramString = HttpUtil.parseMapToPathParam(map);
        String url = HOST + "/v2/api/getResult" + "?" + paramString;
//        System.out.println("\nget_result_url:" + url);
        log.info("\nget_result_url:" + url);

        // 将语音转换信息存入db


        KdxfSpeech ks = new KdxfSpeech();
        Integer maxNo = 0;
        if (orderId2 == null) {
            ks.setOrder_id(orderId);
            ks.setNo(0);
        } else {
            // 处理多语音归属同一个会议的情况
            log.info("***::" + "多语音记录，原始orderId::" + orderId + ",归属orderId2::" + orderId2);
            SQLReady sqlReady = new SQLReady("select max(no) from kdxf_speech where order_id = ?", new Object[]{orderId2});
            List<Integer> maxNoList = kdxfSpeechDao.getSQLManager().execute(sqlReady, Integer.class);
            if (maxNoList.size() != 1) {
                throw new Exception("数据不正确::orderId::" + orderId + "::orderId2::" + orderId2);
            }
            maxNo = maxNoList.get(0) + 1;
            ks.setNo(maxNo);
            ks.setOrder_id(orderId2);
        }
        ks.setDatetime(new Date());
        ks.setFile_name(fileName);
        ks.setState(0); // 进行中
        ks.setMeetingId(meetingId);
        kdxfSpeechDao.insert(ks);

        while (true) {
            String response = HttpUtil.iflyrecGet(url);
            JsonParse jsonParse = gson.fromJson(response, JsonParse.class);
            try {
                // 很可能返回接口会发生错误，这部分忽略掉
                Object j = jsonParse.content.orderInfo.status;
            } catch (Exception e) {
                log.warn("getResult接口调用错误::" + response, e);
                continue;
            }


            if (jsonParse.content.orderInfo.status == 4 || jsonParse.content.orderInfo.status == -1) {
//                System.out.println("订单完成:" + response);
                log.info("订单完成:" + orderId + "::" + getContentForNormal(jsonParse.content.orderResult));

                // 将语音转换信息存入db
                KdxfSpeech ksForUpdate;
                Object[] sqlParam;
                if (orderId2 == null) {
                    sqlParam = new Object[]{orderId, 0};
                } else {
                    sqlParam = new Object[]{orderId2, maxNo};
                }
                SQLReady sqlReady = new SQLReady("select * from kdxf_speech where order_id = ? and no = ?", sqlParam);
                List<KdxfSpeech> listForSpeech = kdxfSpeechDao.getSQLManager().execute(sqlReady, KdxfSpeech.class);
                if (listForSpeech.size() != 1) {
                    throw new Exception("listForSpeech数据不正确::orderId::" + orderId + "::orderId2::" + orderId2);
                }
                ksForUpdate = listForSpeech.get(0);

                ksForUpdate.setState(1); // 完成
                ksForUpdate.setContent(getContentForNormal(jsonParse.content.orderResult));
                ksForUpdate.setReal_duration(jsonParse.content.orderInfo.realDuration);
                // 复合主键更新 cao
                kdxfSpeechDao.updateTemplateById(ksForUpdate);
//                System.out.println(getContentForNormal(jsonParse.content.orderResult));
//                log.info("response:" + response);
//                response:{"code":"000000","descInfo":"success","content":{"orderInfo":{"orderId":"DKHJQ20240328145032273GTx6Tc2agamHqLso","failType":11,"status":-1,"originalDuration":200,"realDuration":98026,"expireTime":1711867724584},"orderResult":""
                return;
            } else {
                log.info("orderId::" + orderId + "::进行中...，状态为:" + jsonParse.content.orderInfo.status);
                //建议使用回调的方式查询结果，查询接口有请求频率限制
                Thread.sleep(7000);

                /* todo 在没有回调之前，此处需要有错误处理机制，在多少次请求后，如果还请求不到需要报错
                    或者提供一个 接口，一旦发生错误（就是无法返回翻译 后的文本），那么可以通过该接口直接请求url返回内容
                */
            }
        }
    }


    public Map manualGetResult(String orderId, String orderId2) throws Exception {
        HashMap<String, Object> map = new HashMap<>(16);
        // 此处的orderid 是 实际请求音频的订单，需要通过该订单获取新上传音频的翻译
        map.put("orderId", orderId);

        LfasrSignature lfasrSignature = new LfasrSignature(appid, keySecret);
        map.put("signa", lfasrSignature.getSigna());
        map.put("ts", lfasrSignature.getTs());
        map.put("appId", appid);
        map.put("resultType", "transfer,predict");
        String paramString = HttpUtil.parseMapToPathParam(map);
        String url = HOST + "/v2/api/getResult" + "?" + paramString;
//        System.out.println("\nget_result_url:" + url);
        log.info("\nget_result_url:" + url);


        String response = HttpUtil.iflyrecGet(url);
        JsonParse jsonParse = gson.fromJson(response, JsonParse.class);
        if (jsonParse.content.orderInfo.status == 4 || jsonParse.content.orderInfo.status == -1) {
            ;
        } else {
            Map<String, Object> m = new HashMap<>();
            m.put("success", false);
            m.put("msg", "进行中");
            return m;
        }
        Integer maxNo = 0;
        if (orderId2 == null) {
            ;
        } else {
            SQLReady sqlReady = new SQLReady("select max(no) from kdxf_speech where order_id = ?", new Object[]{orderId2});
            List<Integer> maxNoList = kdxfSpeechDao.getSQLManager().execute(sqlReady, Integer.class);
            maxNo = maxNoList.get(0) + 1;
        }


        // 将语音转换信息存入db
        KdxfSpeech ksForUpdate;
        Object[] sqlParam;
        if (orderId2 == null) {
            sqlParam = new Object[]{orderId, 0};
        } else {
            sqlParam = new Object[]{orderId2, maxNo};
        }
        SQLReady sqlReady = new SQLReady("select * from kdxf_speech where order_id = ? and no = ?", sqlParam);
        List<KdxfSpeech> listForSpeech = kdxfSpeechDao.getSQLManager().execute(sqlReady, KdxfSpeech.class);
        ksForUpdate = listForSpeech.get(0);

        ksForUpdate.setState(1); // 完成
        ksForUpdate.setContent(getContentForNormal(jsonParse.content.orderResult));
        ksForUpdate.setReal_duration(jsonParse.content.orderInfo.realDuration);
        // 复合主键更新 cao
        kdxfSpeechDao.updateTemplateById(ksForUpdate);
        Map<String, Object> m = new HashMap<>();
        m.put("success", true);
        m.put("msg", "成功存入");
        return m;
    }


    public static void write(String resp) throws IOException {
        //将写入转化为流的形式
        BufferedWriter bw = new BufferedWriter(new FileWriter("src\\main\\resources\\output\\test.txt"));
        String ss = resp;
        bw.write(ss);
        //关闭流
        bw.close();
        System.out.println("写入txt成功");
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

    public static String getContentForNormal(String sb) {

//        System.err.println("回调函数被调用..." + sb);
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






