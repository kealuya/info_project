package com.szhtjykj.speech.xfyun.speech;

import com.szhtjykj.speech.dao.MeetingFileDao;
import com.szhtjykj.speech.model.MeetingFile;
import org.beetl.sql.solon.annotation.Db;
import org.noear.solon.annotation.Component;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.URL;
import java.net.URLConnection;
import java.util.List;

import org.apache.poi.xwpf.usermodel.XWPFDocument;
import org.apache.poi.xwpf.usermodel.XWPFParagraph;

/**
 * @program: kdxf_speech
 * @description:
 * @packagename: com.szhtjykj.speech.xfyun.speech
 * @author: zhanbaohua
 * @date: 2024-05-18 10:31
 **/
@Component
public class XfyunDocumentMeetingService {

    static Logger log = LoggerFactory.getLogger(XfyunSpeechService.class);

    @Db
    MeetingFileDao meetingFileDao;

    /**
     * 根据会议ID，得到文档转译存到库中
     * @param meetingId
     * @return
     * @throws Exception
     */
    public String documentFileTranslation(String meetingId) throws Exception {

        String returnStr = "";
        //根据会议ID，得到文件
        List<MeetingFile> fileList = meetingFileDao.getMeetingFileByMeetingId(meetingId);

        //判断文件类型  pdf,txt,word
        String fileName = fileList.get(0).getFileName();
        String [] fileTypeSZ = fileName.split("\\.");
        String fileType = fileTypeSZ[1];

        if(fileType.equals("txt")){
            //txt 处理
            URL file_Url = new URL(fileList.get(0).getFileUrl());
            URLConnection connection = file_Url.openConnection();
            InputStream inputStream = connection.getInputStream();

            // 使用InputStreamReader将字节流转换为字符流
            // 使用BufferedReader提供了读取文本行的高效方法
            BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream));
            StringBuilder stringBuilder = new StringBuilder();
            String line;
            while ((line = reader.readLine()) != null) {
                stringBuilder.append(line).append("\n");
            }
            log.info("得到的字符"+stringBuilder);
            returnStr = stringBuilder.toString();

            inputStream.close();

            return returnStr;

        }else if(fileType.equals("doc") || fileType.equals("docx")) {
            //word 处理
            URL file_Url = new URL(fileList.get(0).getFileUrl());
            URLConnection connection = file_Url.openConnection();
            InputStream inputStream = connection.getInputStream();

            XWPFDocument doc = new XWPFDocument(inputStream);

            // 读取文档中的段落
            List<XWPFParagraph> paragraphs = doc.getParagraphs();
            for (XWPFParagraph para : paragraphs) {
                returnStr = returnStr + para.getText();
            }
            log.info("得到的字符" + returnStr);
            inputStream.close();

        }else{
            //文件类型不正确
        }

        return returnStr;
    }
}
