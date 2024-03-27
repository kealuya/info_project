package com.example.demo;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;
import org.noear.solon.annotation.Controller;
import org.noear.solon.annotation.Mapping;
import org.noear.solon.annotation.Param;
import org.noear.solon.annotation.Post;
import org.noear.solon.core.handle.Context;
import org.noear.solon.core.handle.UploadedFile;

import java.io.File;
import java.io.IOException;
import java.util.Map;
import java.util.TreeMap;

import static org.apache.http.client.methods.RequestBuilder.delete;

@Controller
public class DemoController {

    private static final Gson gson = new Gson();
    @Mapping("/hello")
    @Post
    public Map hello(Context ctx) throws IOException {
        String jsonString = ctx.body();
        System.out.println(ctx.body());
        System.out.println(ctx.param("aa"));


        Map<String, Object> jsonMap = gson.fromJson(jsonString, new TypeToken<TreeMap<String, Object>>() {}.getType());
        jsonMap.put("666",666);
        TestClass tc = new TestClass();
        tc.age = 342;
        tc.username = "dfje";

        return jsonMap;
    }
    //文件上传
    @Post
    @Mapping("/upload1")
    public String upload(UploadedFile file) throws IOException { //表单变量名要跟参数名对上
        try{
            file.transferTo(new File("/Users/kealuya/Downloads/demo/logo.jpg")); //把它转为本地文件
        } finally {
            //用完之后，删除"可能的"临时文件 //v2.7.2 后支持
            UploadedFile:delete();
        }

        return file.getName();
    }
}

