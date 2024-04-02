<template>
  <div class="header">
    <div class="left">
      <van-image :src="logoIcon" width="40px" height="40px" round />
      <div class="information">
        <div class="f-w-550">L。</div>
        <div class="amount-12">个人账号</div>
      </div>
    </div>
    <div class="right"><van-icon name="chat-o" dot size="20" /></div>
    <!-- <div>
      <van-image :src="logoIcon" width="40px" height="40px" round/>
      <div class="information">
        <div class="f-w-550">L。</div>
        <div class="amount-12">个人账号</div>
      </div>
    </div>
      <div><van-icon name="chat-o" dot /></div> -->
  </div>
  <!-- 按钮部分 -->
  <div class="metting">
    <div @click="handelMettingList">
      <van-image :src="mettingList" width="60px" height="60px" />
      <div>会议列表</div>
    </div>
    <div>
      <van-image :src="mettingList" width="60px" height="60px" />
      <div>xxx</div>
    </div>
    <div>
      <van-image :src="mettingList" width="60px" height="60px" />
      <div>xxxx</div>
    </div>

    <div>
      <van-image :src="mettingList" width="60px" height="60px" />
      <div>xxxxx</div>
    </div>
  </div>
  <!-- <van-tabs v-model:active="active" sticky> -->
  <!-- <van-tab title="会议记录"> -->
  <!-- //class="container" -->
  <div>
      <div style="padding-top: 5px">
        <div style="vertical-align: bottom">
          <div
                  style="height: 70px; width: 92%; margin: 0 20px"
                  class="recwave"
          ></div>
        </div>
        <!-- //灰色的 计时器 -->
        <div
                style="
            margin: 5px 20px;
            height: 20px;
            width: 92%;
            display: inline-block;
            vertical-align: bottom;
            position: relative;
          "
        >
          <div
                  class="recpowerx"
                  style="height: 20px; background: #07c160; position: absolute"
          ></div>
          <div
                  class="recpowert"
                  style="padding-left: 50px; line-height: 20px; position: relative"
          ></div>
        </div>
      </div>

    <div v-if="list.length > 0" class="scrollable-div">
      <div class="list" v-for="item in list">
        <div class="left">
          <div class="time">{{ item.time }}</div>
          <div>{{ item.date }}</div>
        </div>
        <div class="right" style="line-height: 20px">
          <div>{{ item.long }}</div>
          <van-image
                  :src="start"
                  width="20px"
                  height="20px"
                  style="vertical-align: middle; margin-left: 10px"
          ></van-image>
        </div>
      </div>
    </div>
    <div style="text-align: center" v-if="list.length == 0 && isHide">
      <van-image :src="zanwupiaoju" width="300px" height="300px"></van-image>
      <div>暂无录音</div>
    </div>
    <!-- <van-card
      num="2"
      tag="标签"
      price="2.00"
      desc="描述信息"
      title="商品标题"
      thumb="https://fastly.jsdelivr.net/npm/@vant/assets/ipad.jpeg"
      origin-price="10.00"
    /> -->

    <!-- <div class="cardTable">
      <van-button type="primary" plain block round onclick="recPlay()"
        >播放</van-button
      >
      <van-divider />
      <van-button type="primary" plain block round onclick="recUpload()"
        >上传</van-button
      >
      <van-divider />
      <van-button type="primary" plain block round onclick="recLocalDown()"
        >本地下载</van-button
      >
    </div> -->
    <!-- 按钮部分 -->
    <div class="round-button">
      <!-- <div onclick="recOpen()">
        <van-image :src="limits" width="50px" height="50px"></van-image>
        <div>打开权限</div>
      </div> -->
      <div onclick="recStart()">
        <van-image :src="start" width="50px" height="50px"></van-image>
        <div>开始录音</div>
      </div>
      <div onclick="recStop()">
        <van-image :src="end" width="50px" height="50px"></van-image>
        <div>结束录音</div>
      </div>
      <div onclick="recPause()">
        <van-image :src="stop" width="50px" height="50px"></van-image>
        <div>暂停录音</div>
      </div>
    </div>
    <div class="cardTable" @click="handelMetting">
      <van-button type="primary" block round style="width: 92%; margin: 20px"
      >会议结束</van-button
      >
    </div>
    <!-- <div>
      <van-row>
        <van-col span="8">
          <van-button
            plain
            icon="stop-circle-o"
            type="primary"
            round
            onclick="recOpen()"
            >打开权限</van-button
          >

          <van-button
            plain
            icon="stop-circle-o"
            type="primary"
            round
            onclick="recClose()"
            >关闭权限</van-button
          >
        </van-col>
        <van-col span="8">
          <van-button
            plain
            icon="stop-circle-o"
            type="primary"
            round
            onclick="recStart()"
            >开始录制</van-button
          >

          <van-button
            plain
            icon="stop-circle-o"
            type="primary"
            round
            onclick="recStop()"
            >结束录制</van-button
          >
        </van-col>
        <van-col span="8">
          <van-button
            plain
            icon="stop-circle-o"
            type="primary"
            round
            onclick="recPause()"
            >暂停</van-button
          >
          <van-button
            plain
            icon="stop-circle-o"
            type="primary"
            round
            onclick="recResume()"
            >恢复</van-button
          >
        </van-col>
      </van-row>
    </div> -->

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
    <van-dialog
            v-model:show="show"
            :show-cancel-button="false"
            :showConfirmButton="false"
    >
      <van-form @submit="onSubmit" style="padding: 20px" ref="formRef">
        <van-cell-group inset>
          <van-field
                  v-model="username"
                  name="会议标题"
                  label="会议标题"
                  placeholder="会议标题"
                  :rules="[{ required: true, message: '请填写会议标题' }]"
          />
        </van-cell-group>
        <div style="margin: 16px">
          <van-button round block type="primary" native-type="submit">
            提交
          </van-button>
        </div>
      </van-form>
    </van-dialog>
  </div>
  <!-- </van-tab> -->
  <!-- </van-tabs> -->
</template>

<script setup lang="ts">
  // import logoIcon from "../../assets/img/avatar.jpeg";
  import logoIcon from "../../assets/icon/szht_logo.png";
  import zanwupiaoju from "../../assets/img/zanwupiaoju.png";
  import mettingList from "../../assets/img/mettingList.png";
  import start from "../../assets/img/start.png";
  import end from "../../assets/img/end.png";
  import stop from "../../assets/img/stop.png";
  import limits from "../../assets/img/limits.png";

  import {ref, inject, onMounted} from "vue";
  import Recorder from "recorder-core"; //已包含recorder-core和mp3格式支持
  //可选的插件支持项 波形绘制的插件支持
  import "recorder-core/src/extensions/waveview";
  //引入相应格式支持文件；如果需要多个格式支持，把这些格式的编码引擎js文件放到后面统统引入进来即可
  import "recorder-core/src/engine/mp3";
  import "recorder-core/src/engine/mp3-engine";
  // import router from "../../router";
  import { useRouter } from "vue-router";
  const isShow = ref<Boolean>(false);
  const isHide = ref<Boolean>(true);
  const show = ref(false);
  const username = ref("");
  const formRef = ref<any>(null); // 定义表单引用
  const onSubmit = (values) => {
    const valid = formRef.value.validate();
    if (valid) {
      // 如果验证通过，则执行页面跳转
      // 请确保您已经设置好了Vue Router
      router.push({
        name: "metting",
      });
    }
  };
  interface ListItem {
    time: string;
    date: string;
    long: string;
  }

  const list = ref<Array<ListItem>>([]);

  // const over = () => {
  //   list.value.push({
  //     time: "20240329_170029",
  //     date: "2024/3/29",
  //     long: "00:00:14",
  //   });
  //   console.log(list.value);
  // };
  const active = ref(0);
  let userInfoData: any = inject("userInfo");

  const win = window as any;
  const doc = document as any;
  const router = useRouter();
  let rec: any, wave: any, recBlob: any;
  //点击跳转到会议列表页面
  const handelMetting = () => {
    show.value = true;
    // console.log("点击了");
    // router.push({
    //   name: "metting",
    // });
  };
  const handelMettingList =()=>{
    router.push({
      name: "metting",
    });
  }


  onMounted(() => {
    //打开页面，获取权限
    win.recOpen()
  })



  /**调用open打开录音请求好录音权限**/
  win.recOpen = function () {
    //一般在显示出录音按钮或相关的录音界面时进行此方法调用，后面用户点击开始录音时就能畅通无阻了
    rec = null;
    wave = null;
    recBlob = null;
    var newRec = Recorder({
      type: "mp3",
      sampleRate: 16000,
      bitRate: 16, //mp3格式，指定采样率hz、比特率kbps，其他参数使用默认配置；注意：是数字的参数必须提供数字，不要用字符串；需要使用的type类型，需提前把格式支持文件加载进来，比如使用wav格式需要提前加载wav.js编码引擎
      onProcess: function (
              buffers: any,
              powerLevel: any,
              bufferDuration: any,
              bufferSampleRate: any,
              newBufferIdx: any,
              asyncEnd: any
      ) {
        //录音实时回调，大约1秒调用12次本回调
        doc.querySelector(".recpowerx").style.width = powerLevel + "%";
        doc.querySelector(".recpowert").innerText =
                formatMs(bufferDuration, 1) + " / " + powerLevel;

        //可视化图形绘制
        wave.input(buffers[buffers.length - 1], powerLevel, bufferSampleRate);
      },
    });

    newRec.open(
            function () {
              //打开麦克风授权获得相关资源
              rec = newRec;

              //此处创建这些音频可视化图形绘制浏览器支持妥妥的
              wave = Recorder.WaveView({ elem: ".recwave" });

              console.log("已打开录音，可以点击录制开始录音了", 2);
            },
            function (msg: any, isUserNotAllow: any) {
              //用户拒绝未授权或不支持
              console.log(
                      (isUserNotAllow ? "UserNotAllow，" : "") + "打开录音失败：" + msg,
                      1
              );
            }
    );
  };

  /**关闭录音，释放资源**/
  win.recClose = function () {
    if (rec) {
      rec.close();
      console.log("已关闭");
    } else {
      console.log("未打开录音", 1);
    }
  };

  /**开始录音**/
  win.recStart = function () {
    isShow.value = true;
    isHide.value = false;
    //打开了录音后才能进行start、stop调用
    if (rec && Recorder.IsOpen()) {
      recBlob = null;
      rec.start();
      console.log("已开始录音...");
    } else {
      console.log("未打开录音", 1);
    }
  };

  /**暂停录音**/
  win.recPause = function () {
    if (rec && Recorder.IsOpen()) {
      rec.pause();
    } else {
      console.log("未打开录音", 1);
    }
  };
  /**恢复录音**/
  win.recResume = function () {
    if (rec && Recorder.IsOpen()) {
      rec.resume();
    } else {
      console.log("未打开录音", 1);
    }
  };

  /**结束录音，得到音频文件**/
  win.recStop = function () {
    list.value.push({
      time: "20240329_170029",
      date: "2024/3/29",
      long: "00:00:14",
    });
    isShow.value = false;
    if (!(rec && Recorder.IsOpen())) {
      console.log("未打开录音", 1);
      return;
    }
    rec.stop(
            function (blob: any, duration: any) {
              console.log(
                      blob,
                      (win.URL || webkitURL).createObjectURL(blob),
                      "时长:" + duration + "ms"
              );

              recBlob = blob;
              console.log(
                      "已录制mp3：" +
                      formatMs(duration) +
                      "ms " +
                      blob.size +
                      "字节，可以点击播放、上传了",
                      2
              );
              win.recUpload()
              console.log("自动上传文件啦=======");
            },
            function (msg: any) {
              console.log("录音失败:" + msg, 1);
            }
    );

  };

  /**播放**/
  win.recPlay = function () {
    if (!recBlob) {
      console.log("请先录音，然后停止后再播放", 1);
      return;
    }
    var cls = ("a" + Math.random()).replace(".", "");
    console.log('播放中: <span class="' + cls + '"></span>');
    var audio = doc.createElement("audio");
    audio.controls = true;
    doc.querySelector("." + cls).appendChild(audio);
    //简单利用URL生成播放地址，注意不用了时需要revokeObjectURL，否则霸占内存
    audio.src = (win.URL || webkitURL).createObjectURL(recBlob);
    audio.play();

    setTimeout(function () {
      (win.URL || webkitURL).revokeObjectURL(audio.src);
    }, 5000);
  };

  /**上传**/
  win.recUpload = function () {
    var blob = recBlob;
    if (!blob) {
      console.log("请先录音，然后停止后再上传", 1);
      return;
    }

    //本例子假设使用原始XMLHttpRequest请求方式，实际使用中自行调整为自己的请求方式
    //录音结束时拿到了blob文件对象，可以用FileReader读取出内容，或者用FormData上传
    var api = "http://localhost:7001/v4/meeting/uploadMeetingAudioFile";
    var onreadystatechange = function (title: any) {
      return function () {
        if (xhr.readyState == 4) {
          if (xhr.status == 200) {
            console.log(title + "上传成功", 2);
          } else {
            console.log(
                    title +
                    "没有完成上传，演示上传地址无需关注上传结果，只要浏览器控制台内Network面板内看到的请求数据结构是预期的就ok了。",
                    "#d8c1a0"
            );

            console.error(title + "上传失败", xhr.status, xhr.responseText);
          }
        }
      };
    };
    console.log("开始上传到" + api + "，请求稍后...");

    // /***方式一：将blob文件转成base64纯文本编码，使用普通application/x-www-form-urlencoded表单上传***/
    // var reader = new win.FileReader();
    // reader.onloadend = function () {
    //   var postData = "";
    //   postData += "mime=" + encodeURIComponent(blob.type); //告诉后端，这个录音是什么格式的，可能前后端都固定的mp3可以不用写
    //   postData +=
    //           "&upfile_b64=" +
    //           encodeURIComponent(
    //                   (/.+;\s*base64\s*,\s*(.+)$/i.exec(reader.result) || [])[1]
    //           ); //录音文件内容，后端进行base64解码成二进制
    //   //...其他表单参数
    //
    //   var xhr = new XMLHttpRequest();
    //   xhr.open("POST", api);
    //   xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    //   xhr.onreadystatechange = onreadystatechange("上传方式一【Base64】");
    //   xhr.send(postData);
    // };
    // reader.readAsDataURL(blob);

    /***方式二：使用FormData用multipart/form-data表单上传文件***/
    let form=new FormData();
    let filename = "会议ID_"+new Date().getTime()+".mp3"
    form.append("file",blob,filename); //和普通form表单并无二致，后端接收到file参数的文件，文件名为 "会议ID_"+new Date()+".mp3"
    //...其他表单参数

    let xhr=new XMLHttpRequest();
    xhr.open("POST", api);
    xhr.onreadystatechange=onreadystatechange("上传方式二【FormData】");
    xhr.send(form);
  };

  /**本地下载  Local download**/
  win.recLocalDown = function () {
    if (!recBlob) {
      console.log("请先录音，然后停止后再下载", 1);
      return;
    }
    var cls = ("a" + Math.random()).replace(".", "");
    win.recdown64.lastCls = cls;
    console.log(
            '点击 <span class="' +
            cls +
            '"></span> 下载，或复制文本' +
            "<button onclick=\"recdown64('" +
            cls +
            '\')">生成Base64文本</button><span class="' +
            cls +
            '_b64"></span>'
    );

    var fileName = "recorder-" + Date.now() + ".mp3";
    var downA = doc.createElement("A");
    downA.innerHTML = "下载 " + fileName;
    downA.href = (win.URL || webkitURL).createObjectURL(recBlob);
    downA.download = fileName;
    doc.querySelector("." + cls).appendChild(downA);
    if (/mobile/i.test(navigator.userAgent)) {
      alert(
              "因移动端绝大部分国产浏览器未适配Blob Url的下载，所以本demo代码在移动端未调用downA.click()。请尝试点击日志中显示的下载链接下载"
      );
    } else {
      downA.click();
    }

    //不用了时需要revokeObjectURL，否则霸占内存
    //(win.URL||webkitURL).revokeObjectURL(downA.href);
  };
  win.recdown64 = function (cls: any) {
    var el = doc.querySelector("." + cls + "_b64");
    if (win.recdown64.lastCls != cls) {
      el.innerHTML =
              '<span style="color:red">老的数据没有保存，只支持最新的一条</span>';
      return;
    }
    var reader = new FileReader();
    reader.onloadend = function () {
      el.innerHTML = "<textarea></textarea>";
      el.querySelector("textarea").value = reader.result;
    };
    reader.readAsDataURL(recBlob);
  };

  var formatMs = function (ms: any, all?: any) {
    var f = Math.floor(ms / 60000),
            m = Math.floor(ms / 1000) % 60;
    var s =
            (all || f > 0 ? (f < 10 ? "0" : "") + f + ":" : "") +
            (all || f > 0 || m > 0 ? ("0" + m).substr(-2) + "″" : "") +
            ("00" + (ms % 1000)).substr(-3);
    return s;
  };
</script>

<style lang="less" scoped>
  .cardTable {
    width: 100%;
    // margin-top: 20px;
    text-align: center;
    position: fixed;
    bottom: 60px;
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
  .header {
    background-color: #fff;
    // width: 96%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px;
    .left {
      display: flex;
      .information {
        margin-left: 14px;
        font-size: 14px;
        .amount-12 {
          font-size: 12px;
          // color: #ccc;
          margin-top: 6px;
        }
      }
    }
  }
  .f-w-550 {
    font-weight: 550;
  }
  .metting {
    display: flex;
    justify-content: space-around;
    margin: 30px 0;
    text-align: center;
    font-size: 12px;
    // justify-content: center;
  }
  .round-button {
    display: flex;
    justify-content: space-around;
    margin: 20px;
    text-align: center;
    font-size: 12px;
    position: fixed;
    bottom: 130px;
    width: 94%;
  }
  .footer {
    width: 100%;
  }
  .list {
    display: flex;
    justify-content: space-between;
    padding: 20px;
    margin: 20px;
    background-color: #fff;
    border-radius: 8px;
    align-items: center;
    font-size: 12px;
    color: #7b7b7b;
    // color: #ccc;
    .right {
      display: flex;
      text-align: center;
    }
    .time {
      font-size: 14px;
      font-weight: 600;
      margin-bottom: 10px;
      color: #000;
    }
  }
  .scrollable-div {
    max-height: 360px; /* 设置最大高度为300px */
    overflow-y: auto; /* 当内容超出高度时显示垂直滚动条 */
  }
</style>
