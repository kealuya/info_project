<template>
  <van-tabs v-model:active="active" sticky>
    <van-tab title="会议记录">
      <div class="container">

        <!-- 波形绘制区域 -->
        <div>
          <div style="padding-top:5px">
            <div style="vertical-align:bottom;background: #3BD7D3">
              <div style="height:100px;width:100%;" class="recwave"></div>
            </div>
            <div style="height:20px;width:100%;margin-top:5px;display:inline-block;vertical-align:bottom;background:#999;position:relative;">
              <div class="recpowerx" style="height:40px;background:#0B1;position:absolute;"></div>
              <div class="recpowert" style="padding-left:50px; line-height:40px; position: relative;"></div>
            </div>
          </div>
        </div>


        <van-card
            num="2"
            tag="标签"
            price="2.00"
            desc="描述信息"
            title="商品标题"
            thumb="https://fastly.jsdelivr.net/npm/@vant/assets/ipad.jpeg"
            origin-price="10.00"
        />
        <van-card
            num="2"
            tag="标签"
            price="2.00"
            desc="描述信息"
            title="商品标题"
            thumb="https://fastly.jsdelivr.net/npm/@vant/assets/ipad.jpeg"
            origin-price="10.00"
        />


        <div class="cardTable">
          <van-button type="primary" plain block round>播放</van-button>
          <van-divider />
          <van-button type="primary" plain block round>上传</van-button>
          <van-divider />
          <van-button type="primary" plain block round>本地下载</van-button>
        </div>

        <div class="bottom_box" >
          <van-row>
            <van-col span="8">
              <van-button plain icon="stop-circle-o" type="primary" round="true"/>
            </van-col>
            <van-col span="8">
              <van-button plain icon="stop-circle-o" type="primary" round="true"/>
            </van-col>
            <van-col span="8">
              <van-button plain icon="stop-circle-o" type="primary" round="true"/>
            </van-col>
          </van-row>
        </div>



<!--        <div class="mainBox">-->
<!--          &lt;!&ndash; 按钮控制区域 &ndash;&gt;-->
<!--          <div class="pd btns">-->
<!--            <div>-->
<!--              <button onclick="recOpen()" style="margin-right:10px">打开录音,请求权限</button>-->
<!--              <button onclick="recClose()" style="margin-right:0">关闭录音,释放资源</button>-->
<!--            </div>-->

<!--            <button onclick="recStart()">录制</button>-->
<!--            <button onclick="recStop()" style="margin-right:80px">停止</button>-->

<!--            <span style="display: inline-block;">-->
<!--				<button onclick="recPause()">暂停</button>-->
<!--				<button onclick="recResume()">继续</button>-->
<!--			</span>-->
<!--            <span style="display: inline-block;">-->
<!--				<button onclick="recPlay()">播放</button>-->
<!--				<button onclick="recUpload()">上传</button>-->
<!--				<button onclick="recLocalDown()">本地下载</button>-->
<!--			</span>-->
<!--          </div>-->





      </div>
    </van-tab>
  </van-tabs>
</template>

<script setup lang="ts">
import {ref, inject} from "vue";
import Recorder from 'recorder-core' //已包含recorder-core和mp3格式支持

const active = ref(0);
let userInfoData: any = inject("userInfo");

const win = window as any;
const doc = document as any;

let rec: any, wave: any, recBlob: any;
/**调用open打开录音请求好录音权限**/
win.recOpen=function(){//一般在显示出录音按钮或相关的录音界面时进行此方法调用，后面用户点击开始录音时就能畅通无阻了
  rec=null;
  wave=null;
  recBlob=null;
  var newRec=Recorder({
    type:"mp3",sampleRate:16000,bitRate:16 //mp3格式，指定采样率hz、比特率kbps，其他参数使用默认配置；注意：是数字的参数必须提供数字，不要用字符串；需要使用的type类型，需提前把格式支持文件加载进来，比如使用wav格式需要提前加载wav.js编码引擎
    ,onProcess:function(buffers:any,powerLevel:any,bufferDuration:any,bufferSampleRate:any,newBufferIdx:any,asyncEnd:any){
      //录音实时回调，大约1秒调用12次本回调
      doc.querySelector(".recpowerx").style.width=powerLevel+"%";
      doc.querySelector(".recpowert").innerText=formatMs(bufferDuration,1)+" / "+powerLevel;

      //可视化图形绘制
      wave.input(buffers[buffers.length-1],powerLevel,bufferSampleRate);
    }
  });

  newRec.open(function(){//打开麦克风授权获得相关资源
    rec=newRec;

    //此处创建这些音频可视化图形绘制浏览器支持妥妥的
    wave=Recorder.WaveView({elem:".recwave"});

    console.log("已打开录音，可以点击录制开始录音了",2);
  },function(msg:any,isUserNotAllow:any){//用户拒绝未授权或不支持
    console.log((isUserNotAllow?"UserNotAllow，":"")+"打开录音失败："+msg,1);
  });
};



/**关闭录音，释放资源**/
win.recClose=function(){
  if(rec){
    rec.close();
    console.log("已关闭");
  }else{
    console.log("未打开录音",1);
  };
};



/**开始录音**/
win.recStart=function(){//打开了录音后才能进行start、stop调用
  if(rec&&Recorder.IsOpen()){
    recBlob=null;
    rec.start();
    console.log("已开始录音...");
  }else{
    console.log("未打开录音",1);
  };
};

/**暂停录音**/
win.recPause=function(){
  if(rec&&Recorder.IsOpen()){
    rec.pause();
  }else{
    console.log("未打开录音",1);
  };
};
/**恢复录音**/
win.recResume=function(){
  if(rec&&Recorder.IsOpen()){
    rec.resume();
  }else{
    console.log("未打开录音",1);
  };
};

/**结束录音，得到音频文件**/
win.recStop=function(){
  if(!(rec&&Recorder.IsOpen())){
    console.log("未打开录音",1);
    return;
  };
  rec.stop(function(blob:any,duration:any){
    console.log(blob,(win.URL||webkitURL).createObjectURL(blob),"时长:"+duration+"ms");

    recBlob=blob;
    console.log("已录制mp3："+formatMs(duration)+"ms "+blob.size+"字节，可以点击播放、上传了",2);
  },function(msg:any){
    console.log("录音失败:"+msg,1);
  });
};









/**播放**/
win.recPlay=function(){
  if(!recBlob){
    console.log("请先录音，然后停止后再播放",1);
    return;
  };
  var cls=("a"+Math.random()).replace(".","");
  console.log('播放中: <span class="'+cls+'"></span>');
  var audio=doc.createElement("audio");
  audio.controls=true;
  doc.querySelector("."+cls).appendChild(audio);
  //简单利用URL生成播放地址，注意不用了时需要revokeObjectURL，否则霸占内存
  audio.src=(win.URL||webkitURL).createObjectURL(recBlob);
  audio.play();

  setTimeout(function(){
    (win.URL||webkitURL).revokeObjectURL(audio.src);
  },5000);
};

/**上传**/
win.recUpload=function(){
  var blob=recBlob;
  if(!blob){
    console.log("请先录音，然后停止后再上传",1);
    return;
  };

  //本例子假设使用原始XMLHttpRequest请求方式，实际使用中自行调整为自己的请求方式
  //录音结束时拿到了blob文件对象，可以用FileReader读取出内容，或者用FormData上传
  var api="https://xx.xx/test_request";
  var onreadystatechange=function(title:any){
    return function(){
      if(xhr.readyState==4){
        if(xhr.status==200){
          console.log(title+"上传成功",2);
        }else{
          console.log(title+"没有完成上传，演示上传地址无需关注上传结果，只要浏览器控制台内Network面板内看到的请求数据结构是预期的就ok了。", "#d8c1a0");

          console.error(title+"上传失败",xhr.status,xhr.responseText);
        };
      };
    };
  };
  console.log("开始上传到"+api+"，请求稍后...");

  /***方式一：将blob文件转成base64纯文本编码，使用普通application/x-www-form-urlencoded表单上传***/
  var reader=new win.FileReader();
  reader.onloadend=function(){
    var postData="";
    postData+="mime="+encodeURIComponent(blob.type);//告诉后端，这个录音是什么格式的，可能前后端都固定的mp3可以不用写
    postData+="&upfile_b64="+encodeURIComponent((/.+;\s*base64\s*,\s*(.+)$/i.exec(reader.result)||[])[1]) //录音文件内容，后端进行base64解码成二进制
    //...其他表单参数

    var xhr=new XMLHttpRequest();
    xhr.open("POST", api);
    xhr.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
    xhr.onreadystatechange=onreadystatechange("上传方式一【Base64】");
    xhr.send(postData);
  };
  reader.readAsDataURL(blob);

  /***方式二：使用FormData用multipart/form-data表单上传文件***/
  var form=new FormData();
  form.append("upfile",blob,"recorder.mp3"); //和普通form表单并无二致，后端接收到upfile参数的文件，文件名为recorder.mp3
  //...其他表单参数

  var xhr=new XMLHttpRequest();
  xhr.open("POST", api);
  xhr.onreadystatechange=onreadystatechange("上传方式二【FormData】");
  xhr.send(form);
};


/**本地下载  Local download**/
win.recLocalDown=function(){
  if(!recBlob){
    console.log("请先录音，然后停止后再下载",1);
    return;
  };
  var cls=("a"+Math.random()).replace(".","");
  win.recdown64.lastCls=cls;
  console.log('点击 <span class="'+cls+'"></span> 下载，或复制文本'
      +'<button onclick="recdown64(\''+cls+'\')">生成Base64文本</button><span class="'+cls+'_b64"></span>');

  var fileName="recorder-"+Date.now()+".mp3";
  var downA=doc.createElement("A");
  downA.innerHTML="下载 "+fileName;
  downA.href=(win.URL||webkitURL).createObjectURL(recBlob);
  downA.download=fileName;
  doc.querySelector("."+cls).appendChild(downA);
  if(/mobile/i.test(navigator.userAgent)){
    alert("因移动端绝大部分国产浏览器未适配Blob Url的下载，所以本demo代码在移动端未调用downA.click()。请尝试点击日志中显示的下载链接下载");
  }else{
    downA.click();
  }

  //不用了时需要revokeObjectURL，否则霸占内存
  //(win.URL||webkitURL).revokeObjectURL(downA.href);
};
win.recdown64=function(cls:any){
  var el=doc.querySelector("."+cls+"_b64");
  if(win.recdown64.lastCls!=cls){
    el.innerHTML='<span style="color:red">老的数据没有保存，只支持最新的一条</span>';
    return;
  }
  var reader = new FileReader();
  reader.onloadend = function() {
    el.innerHTML='<textarea></textarea>';
    el.querySelector("textarea").value=reader.result;
  };
  reader.readAsDataURL(recBlob);
};









var formatMs=function(ms:any,all?:any){
  var f=Math.floor(ms/60000),m=Math.floor(ms/1000)%60;
  var s=(all||f>0?(f<10?"0":"")+f+":":"")
      +(all||f>0||m>0?("0"+m).substr(-2)+"″":"")
      +("00"+ms%1000).substr(-3);
  return s;
};








</script>

<style lang="less" scoped>

.cardTable {
  margin-top: 20px;
  text-align: center;
}

.touxiangimg {
  width: 20px;
  height: 20px;
}

.container {
  position: relative;
  width: 100%;
  min-height: 750px;
}

.bottom_box {
  min-width: 100%;
  position: absolute;
  bottom: 0;
  text-align: center;
}


</style>




