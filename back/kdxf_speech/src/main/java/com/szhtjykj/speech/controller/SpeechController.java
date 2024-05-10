package com.szhtjykj.speech.controller;

import com.google.gson.Gson;
import com.szhtjykj.speech.dao.KdxfSpeechDao;
import com.szhtjykj.speech.model.KdxfSpeech;
import com.szhtjykj.speech.xfyun.speech.XfyunSpeechService;
import org.beetl.sql.solon.annotation.Db;
import org.noear.solon.Solon;
import org.noear.solon.annotation.*;
import org.noear.solon.core.handle.Context;
import org.noear.solon.core.handle.UploadedFile;
import org.noear.solon.data.annotation.Tran;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.File;
import java.io.IOException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;


@Controller
public class SpeechController {
    static Logger log = LoggerFactory.getLogger(SpeechController.class);
    private static String AUDIO_FILE_PATH;

    private static final Gson gson = new Gson();

    static {
        AUDIO_FILE_PATH = Solon.cfg().get("server.audio.save.path");
    }

    @Inject
    XfyunSpeechService xfyunService;


    //文件上传
    @Post
    @Mapping("/uploadAudio")
    public Map uploadAudio(Context ctx, UploadedFile file) {
        //表单变量名要跟参数名对上

        // 在传orderId的场合，视为多语音属于共同会议
        String orderId2 = ctx.param("orderId");

        Map<String, Object> returnMap = new HashMap<>();
        // 设置时间戳格式
        SimpleDateFormat sdf = new SimpleDateFormat("yyyyMMddHHmm_ssSSS");
        String orderId;
        // 格式化时间戳
        String timestamp = sdf.format(new Date());
        String tempFileName = "audio_" + timestamp + "." + file.getExtension();
        try {

            File newFile = new File(AUDIO_FILE_PATH + tempFileName);
            // todo 后续提供obs上传，然后通过url识别
            file.transferTo(newFile); //把它转为本地文件
            orderId = xfyunService.upload(newFile, orderId2);

        } catch (Exception e) {
            log.error("音频上传发生错误::", e);

            returnMap.put("success", false);
            returnMap.put("msg", e.toString());

            return returnMap;
        } finally {
            //用完之后，删除"可能的"临时文件 //v2.7.2 后支持
            try {
                file.delete();
            } catch (IOException e) {
                ;//无需处理
            }
        }
        returnMap.put("success", true);
        returnMap.put("msg", "");
        returnMap.put("orderId", orderId);
        returnMap.put("fileName", tempFileName);

        return returnMap;
    }


    @Post
    @Mapping("/getResult")
    public Map getResult(String orderId, String orderId2) throws Exception {

        return xfyunService.manualGetResult(orderId, orderId2);
    }

    @Get
    @Mapping("/callbackMsg")
    public void callbackMsg(String orderId, String status) throws Exception {
        //TODO 回调信息方法
    }


    @Db
    KdxfSpeechDao kdxfSpeechDao;

    @Tran
    @Mapping("/test")
    @Post
    public String test(Context ctx) throws IOException {


//        kdxfSpeechDao.unique("DKHJQ20240329133151293huiOKZj2CY3pnEV5");
//
//
//        System.out.println(ctx.param("param1"));
//        System.out.println(ctx.param("param2"));
//
//        SQLReady sqlReady = new SQLReady("select max(no) from kdxf_speech where order_id = ?", new Object[]{"66666"});
//        List<Integer> kl33 = kdxfSpeechDao.getSQLManager().execute(sqlReady, Integer.class);
//        System.out.println(kl33.get(0));


        KdxfSpeech ks = new KdxfSpeech();
        ks.setOrder_id("66666");
        ks.setDatetime(new Date());
        ks.setFile_name("sdfsdf");
//        ks.setNo(Long.valueOf(0));
        ks.setReal_duration(new Long(343434));
        ks.setState(1);

        kdxfSpeechDao.insert(ks);

//        List<KdxfSpeech> kl = kdxfSpeechDao.execute("select * from kdxf_speech");
//
//        for (KdxfSpeech kk : kl) {
//            System.out.println(kk.getOrder_id());
//        }


        return "ok";
    }

}
