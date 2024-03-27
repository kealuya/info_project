package com.szhtjykj.speech.controller;

import com.google.gson.Gson;
import org.noear.solon.annotation.Controller;
import org.noear.solon.annotation.Mapping;
import org.noear.solon.annotation.Post;
import org.noear.solon.core.handle.UploadedFile;

import java.io.File;
import java.io.IOException;
import java.net.URISyntaxException;
import java.text.SimpleDateFormat;
import java.util.Date;

import static org.apache.http.client.methods.RequestBuilder.delete;

@Controller
public class SpeechController {

    private static String AUDIO_FILE_PATH;

    private static final Gson gson = new Gson();

    static {
        try {
            AUDIO_FILE_PATH = SpeechController.class.getResource("/").toURI().getPath() + "/audio/";
        } catch (URISyntaxException e) {
            e.printStackTrace();
        }
    }

    //文件上传
    @Post
    @Mapping("/uploadAudio")
    public String uploadAudio(UploadedFile file) throws IOException { //表单变量名要跟参数名对上

        // 设置时间戳格式
        SimpleDateFormat sdf = new SimpleDateFormat("yyyyMMddHHmmssSSS");

        // 格式化时间戳
        String timestamp = sdf.format(new Date());
        String tempFileName = "audio_" + timestamp.toString() + ".mp3";
        try {
            file.transferTo(new File(AUDIO_FILE_PATH + tempFileName)); //把它转为本地文件
        } finally {
            //用完之后，删除"可能的"临时文件 //v2.7.2 后支持
            file.delete();
        }

        return file.getName();
    }


}
