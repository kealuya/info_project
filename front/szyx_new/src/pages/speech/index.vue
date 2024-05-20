<template>
    <van-nav-bar
            left-arrow
            left-text="返回"
            title="会议记录"
            @click-left="onClickLeft"
    />
    <div class="speech-box">
        <!--    头部标题-->
        <div class="header">
            <div class="header_title">
                <van-image :src="speech" height="35" width="35"></van-image>
                <span>音频记录</span>
            </div>
            <div class="f-z-14-c">会议ID</div>
            <div class="f-z-16">{{ meetingId }}</div>
                  <div class="f-z-14-c m-t-10">会议标题</div>
                  <div class="f-z-16">{{reasonForBorrowing}}</div>
        </div>
        <!--    动画-->
        <div class="animation-box">
            <div style="vertical-align: bottom">
                <div
                        class="recwave"
                        style="height: 50px; width: 92%; margin: 0 20px"
                ></div>
            </div>
            <!-- //灰色的 计时器 -->
            <div
                    style="
            margin: 5px 15px;
            height: 10vh;
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
        <!--    列表-->
        <div class="speech-list-box">
            <div v-if="list.length > 0" class="speech-list-box" style="flex: 1;">
                <div v-for="(item,index) in list" :key="index" class="list">
                    <div class="left">
                        <div class="time">{{ item.time }}</div>
                        <div>{{ item.date }}</div>
                    </div>
                    <div class="right" style="line-height: 20px">
                        <div>{{ item.long }}</div>
                        <!--            <van-image-->
                        <!--                :src="start1"-->
                        <!--                height="20px"-->
                        <!--                style="vertical-align: middle; margin-left: 10px"-->
                        <!--                width="20px"-->
                        <!--            ></van-image>-->
                        <van-image
                                :src="deletePng"
                                height="18px"
                                style="vertical-align: middle; margin-left: 10px"
                                width="18px"
                                @click="deleteItem(item)"
                        ></van-image>
                    </div>
                </div>
            </div>
            <div v-if="list.length == 0" style="text-align: center">
                <van-image :src="zanwupiaoju" height="auto" width="180px" class="speech-list-box"></van-image>
                <div class="empty">暂无录音</div>
            </div>
        </div>
        <!-- 按钮部分 -->
        <div class="btn-box">
            <div class="round-button">
                <!-- <div onclick="recOpen()">
                  <van-image :src="limits" width="50px" height="50px"></van-image>
                  <div>打开权限</div>
                </div> -->
                <div onclick="recStart()" v-if="isShowBtn">
                    <van-image :src="start" height="50px" width="50px"></van-image>
                    <div>开始录音</div>
                </div>
                <div onclick="recPause()" v-if="isShowHf">
                    <van-image :src="stop" height="50px" width="50px"></van-image>
                    <div>暂停录音</div>
                </div>
                <div onclick="recResume()" v-if="isShowZt">
                    <van-image :src="start" height="50px" width="50px"></van-image>
                    <div>恢复录音</div>
                </div>
                <div onclick="recStop()">
                    <van-image :src="end" height="50px" width="50px"></van-image>
                    <div>结束录音</div>
                </div>
            </div>
            <div class="cardTable">
                <van-button block style="width: 92%; margin: 20px" type="primary" onclick="recClose()" :disabled="isDisabled"
                >会议结束
                </van-button
                >
            </div>
        </div>
    </div>
    <van-dialog
        title="会议信息"
            v-model:show="show"
            :show-cancel-button="false"
            :showConfirmButton="false"
    >
        <van-form ref="formRef" @submit="onSubmit" style="margin: 2vw">
            <van-cell-group inset>
                <van-field
                    class="m-b-10"
                    label="会议标题"
                        v-model="reasonForBorrowing" :rules="[{ required: true, message: '请填写会议标题' }]"
                        name="reasonForBorrowing"
                        placeholder="请输入会议标题"
                />
                <van-field
                    class="m-b-10"
                    label="会议时间"
                    v-model="timeer"
                    name="timeer"
                    readonly
                />
                <van-field
                    class="m-b-10"
                    label="参会人"
                    v-model="meetingPeople"
                    name="meetingPeople"
                    placeholder="请输入参会人"
                />
                <van-field
                    class="m-b-10"
                    label="会议地址"
                    v-model="meetingAddress"
                    name="meetingAddress"
                    placeholder="请输入会议城市"
                />
            </van-cell-group>
            <div style="margin: 16px;display:flex">
                <van-button block size="small" style="margin-right: 10px" @click="cancel">
                    取消
                </van-button>
                <van-button block native-type="submit" size="small" type="primary" :loading="isLoading" loading-text="提交中请稍后">
                    提交
                </van-button>
            </div>
        </van-form>
    </van-dialog>
    <van-number-keyboard safe-area-inset-bottom/>
  <!-- </van-tab> -->
  <!-- </van-tabs> -->
</template>

<script lang="ts" setup>
// import Vconsole from 'vconsole'
//
// const vconsole = new Vconsole();
// import {createMeeting} from '../../services/task-processing/index'
import {meetingData, useInvoiceData, userInfoData} from "../../store";
import zanwupiaoju from '../../assets/img/zanwupiaoju.png'
import speech from "../../assets/img/speech.png";
// import logoIcon from "../../assets/img/avatar.jpeg";
import logoIcon from "../../assets/icon/szht_logo.png";
// import zanwupiaoju from "../../assets/img/empty.png";
// import mettingList from "../../assets/img/mettingList.png";
import start from "../../assets/img/kaishi_btn.png";
import start1 from "../../assets/img/start.png";

import deletePng from '../../assets/img/delete.png'
import end from "../../assets/img/tingzhi_btn.png";
import stop from "../../assets/img/zanting_btn.png";
// import limits from "../../assets/img/limits.png";
import {showConfirmDialog, showFailToast, showLoadingToast, showSuccessToast, showToast, Toast} from 'vant';
import {ref, inject, onMounted, onBeforeMount} from "vue";
import Recorder from "recorder-core"; //已包含recorder-core和mp3格式支持
//可选的插件支持项 波形绘制的插件支持
import "recorder-core/src/extensions/waveview";
//引入相应格式支持文件；如果需要多个格式支持，把这些格式的编码引擎js文件放到后面统统引入进来即可
import "recorder-core/src/engine/mp3";
import "recorder-core/src/engine/mp3-engine";
// import router from "../../router";
import {useRouter} from "vue-router";
import huiyijilu from "../../assets/img/huiyijilu.png";
import {createMeeting, deleteAudioMeetingFile} from "../../services/task-processing";

const value = ref<string>('')
const isShow = ref<Boolean>(false);
const isHide = ref<Boolean>(true);
const show = ref(false);
const username = ref("");
const indexValue = ref()
const height = ref(0)
const formattedDate = ref<string>()
const date = ref<any>()
const userId = ref<string | undefined | null>()
const corpCode = ref<string | undefined | null>()
const isShowBtn = ref<boolean>(true)  //控制 显示 开始按钮 还是 暂停按钮
const isShowHf = ref<boolean>(false)
const isShowZt = ref<boolean>(false)
const isDisabled = ref<boolean>(true) //控制 会议记录按钮的 显示隐藏
const meetingId = ref<string>()
const isHideAnimation = ref<boolean>(false)
const formRef = ref<any>(null); // 定义表单引用
const personal = ref<string>('')
const timeer = ref<string>('')
const meetingPeople = ref<string>('')
const meetingAddress = ref<string>('')
const isLoading = ref<boolean>(false)
const idNumber = ref<number>(0)
interface ListItem {
    time: string;
    date: string;
    long: string;
}

const reasonForBorrowing = ref<string>('')
const list = ref<Array<ListItem>>([]);
const deleteItem = (item: any) => {
    let deleteParams = {
        meetingId:meetingId.value,
        fileName:item.time,
        corpCode: corpCode.value,
        creater:userId.value
    }
    showConfirmDialog({
        title: '提示',
        message:
            '你确定要删除吗？',
    })
        .then(() => {
            deleteAudioMeetingFile(deleteParams).then((res:any)=>{
             if(res.success){
              list.value= list.value.filter(i => i.time !==item.time );
             }else{
                 showFailToast(res.msg)
             }
            })
            // list.value.splice(indexValue.value, 1)
        })
        .catch(() => {
            // on cancel
        });
}
const onSubmit = (values: any) => {
    const valid = formRef.value.validate();
    if (valid) {
        let params = {
            meetingId:meetingId.value,
            meetingTitle:reasonForBorrowing.value,
            meetingType:'audio',
            meetingTime:timeer.value,
            meetingPeople:meetingPeople.value,
            meetingAddress:meetingAddress.value,
            corpCode:corpCode.value,
           creater:userId.value
        }
        isLoading.value = true
      createMeeting(params).then((res:any)=>{
        if(res.success){
            //关闭弹窗
            isShow.value = false
            showSuccessToast(res.msg)

        }else{
            showFailToast(res.msg)
        }
      }).finally(()=>{
          isLoading.value = false
          router.replace('/business');
      })

    }
};
const cancel = () => {
    show.value = false
}

// const over = () => {
//   list.value.push({
//     time: "20240329_170029",
//     date: "2024/3/29",
//     long: "00:00:14",
//   });
//   console.log(list.value);
// };
const isMobile = ref(false);

const checkMobile = () => {
    const mediaQuery = window.matchMedia('(max-width: 768px)');
    isMobile.value = mediaQuery.matches;
};
onMounted(() => {
    //打开页面，获取权限
    win.recOpen()
    // const userInfoState: any = heightData();
    // height.value = userInfoState.height
    checkMobile()
})
const seeHeight = ref()
seeHeight.value = window.innerHeight
// onBeforeMount(()=>{
//
// })

const active = ref(0);
let userInfoData: any = inject("userInfo");

const win = window as any;
const doc = document as any;
const router = useRouter();
let rec: any, wave: any, recBlob: any;
const onClickLeft = () => {
    router.replace('/business')
    rec.close()
    // test()
}
// const test = ()=>{
//     win.recClose = function () {
//         if (rec) {
//             rec.close();
//             console.log("已关闭");
//         } else {
//             console.log("未打开录音", 1);
//         }
//     };
// }
//点击跳转到会议列表页面
const handelMetting = () => {
    show.value = true;
    // console.log("点击了");
    // router.push({
    //   name: "metting",
    // });
};
const handelMettingList = () => {
    router.replace({
        name: "metting",
    });
}
/**调用open打开录音请求好录音权限**/
win.recOpen = function () {
    isHideAnimation.value = true
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
            wave = Recorder.WaveView({elem: ".recwave"});

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
    show.value = true   //显示弹窗
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
    isShowHf.value = true
    isShowBtn.value = false;//隐藏开始录音按钮
    isDisabled.value = true
    console.log(' isShow.value ', isShow.value)
    isHide.value = false;
    //打开了录音后才能进行start、stop调用
    if (rec && Recorder.IsOpen()) {
        // debugger
        recBlob = null;
        rec.start();
        console.log("已开始录音...");
    } else {
        console.log("未打开录音", 1);
    }
};

/**暂停录音**/
win.recPause = function () {
    isShowHf.value = false
    isShowZt.value = true
    // isShowBtn.value = true
    // isShowHf.value = true
    if (rec && Recorder.IsOpen()) {
        rec.pause();
    } else {
        console.log("未打开录音", 1);
    }
};
/**恢复录音**/
win.recResume = function () {
    isShowHf.value = true
    isShowZt.value = false
    if (rec && Recorder.IsOpen()) {
        rec.resume();
    } else {
        console.log("未打开录音", 1);
    }
};

/**结束录音，得到音频文件**/
win.recStop = function () {
    idNumber.value += 1
    isShowBtn.value= true
    isShowHf.value = false
    isShowZt.value = false
    isDisabled.value = false
    // isHideAnimation.value = false
    isShow.value = false;
    // list.value.push({
    //     time: "会议ID_" + formattedDate.value+ ".mp3",
    //     date: date.value,
    //     long: '11',
    // });
    if (!(rec && Recorder.IsOpen())) {
        console.log("未打开录音", 1);
        return;
    } else {

    }
    // if(list.value.length==0){
    //     showFailToast('未打开录音!')
    // }else {
    //
    // }
    rec.stop(
        function (blob: any, duration: any) {
            console.log(
                blob,
                (win.URL || webkitURL).createObjectURL(blob),
                "时长:" + duration + "ms"
            );
            let timer = formatMilliseconds(duration)
            list.value.push({
                time:  meetingId.value+'_0'+formatter(idNumber.value) + '.mp3',
                date: date.value,
                long: timer,
            });
            recBlob = blob;
            console.log(
                "已录制mp3：" +
                formatMs(duration) +
                "ms " +
                blob.size +
                "字节，可以点击播放、上传了",
                2
            );
            //控制 动画和 进度条的关闭
            // doc.querySelector(".recpowerx").style.width = 0 + "%";
            doc.querySelector(".recpowert").innerText = '00:00"000 / 00';


            win.recUpload(duration)
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
win.recUpload = function (duration:any) {
    var blob = recBlob;
    if (!blob) {
        console.log("请先录音，然后停止后再上传", 1);
        return;
    }

    //本例子假设使用原始XMLHttpRequest请求方式，实际使用中自行调整为自己的请求方式
    //录音结束时拿到了blob文件对象，可以用FileReader读取出内容，或者用FormData上传
    var api = "https://szhtjykj.mynatapp.cc/v4/meeting/uploadMeetingAudioFile";
    var onreadystatechange = function (title: any) {
        return function () {
            console.log('xhr',xhr)
            if (xhr.readyState == 4) {
                if (xhr.status == 200) {
                    console.log(title + "上传成gong", 0);
                    const response = xhr.response;
                    console.log(JSON.parse(response).success,'1000000000000000');
                    if(JSON.parse(response).success){
                         showToast(JSON.parse(response).msg)

                } else {
                        showFailToast('上传失败，请联系管理员！')
                    }
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

    /***方式一：将blob文件转成base64纯文本编码，使用普通application/x-www-form-urlencoded表单上传***/
    /*var reader = new win.FileReader();
    reader.onloadend = function () {
      var postData = "";
      postData += "mime=" + encodeURIComponent(blob.type); //告诉后端，这个录音是什么格式的，可能前后端都固定的mp3可以不用写
      postData +=
        "&upfile_b64=" +
        encodeURIComponent(
          (/.+;\s*base64\s*,\s*(.+)$/i.exec(reader.result) || [])[1]
        ); //录音文件内容，后端进行base64解码成二进制
      //...其他表单参数

      var xhr = new XMLHttpRequest();
      xhr.open("POST", api);
      xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
      xhr.onreadystatechange = onreadystatechange("上传方式一【Base64】");
      xhr.send(postData);
    };
    reader.readAsDataURL(blob);*/

    /***方式二：使用FormData用multipart/form-data表单上传文件***/
    let form = new FormData();
    let filename = meetingId.value+'_0'+formatter(idNumber.value)+'.mp3'
    form.append("file", blob, filename); //和普通form表单并无二致，后端接收到file参数的文件，文件名为 "会议ID_"+new Date()+".mp3"
    //...其他表单参数
    form.append("meetingId", meetingId.value as string);
    form.append("meetingTitle", reasonForBorrowing.value);
    form.append("userId", userId.value as string);
    form.append("corpCode", corpCode.value as string);
    form.append("audioTime", formatMilliseconds(duration));


    let xhr = new XMLHttpRequest();
    xhr.open("POST", api);
    xhr.onreadystatechange = onreadystatechange("上传方式二【FormData】");
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
var formatter = (id:any)=>{
    if (id < 10) {
        return '0' + id;
    } else {
        return '' + id;
    }
}

function formatMilliseconds(milliseconds:number) {
    // 将毫秒数转换为秒数
    var totalSeconds = Math.floor(milliseconds / 1000);

    // 计算小时、分钟和秒数
    var hours = Math.floor(totalSeconds / 3600);
    var minutes = Math.floor((totalSeconds % 3600) / 60);
    var seconds = totalSeconds % 60;

    // 将时、分、秒数转换为两位数的字符串
    var formattedHours = hours.toString().padStart(2, '0');
    var formattedMinutes = minutes.toString().padStart(2, '0');
    var formattedSeconds = seconds.toString().padStart(2, '0');

    // 返回格式化后的时间字符串
    return formattedHours + ':' + formattedMinutes + ':' + formattedSeconds;
}

// 示例用法
var milliseconds = 6640; // 6640 毫秒
var formattedTime = formatMilliseconds(milliseconds);
console.log(formattedTime); // 输出：00:00:06
const timeHandle = () => {
    const currentDate = new Date();
    const year = currentDate.getFullYear();
    const month = String(currentDate.getMonth() + 1).padStart(2, '0');
    const day = String(currentDate.getDate()).padStart(2, '0');
    const hours = String(currentDate.getHours()).padStart(2, '0');
    const minutes = String(currentDate.getMinutes()).padStart(2, '0');
    const seconds = String(currentDate.getSeconds()).padStart(2, '0');
    date.value = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    let formattedDate = `${year}${month}${day}${hours}${minutes}${seconds}`;
    return formattedDate
}
//生成随机数函数
const generateTwoDigitNumber = () => {
    const randomNumber = Math.floor(Math.random() * 900) + 100; // 生成100到999之间的随机数
    return randomNumber.toString(); // 将随机数转换为字符串并返回
}

onMounted(() => {
    const currentDate = new Date();
    const year = currentDate.getFullYear();
    const month = String(currentDate.getMonth() + 1).padStart(2, '0');
    const day = String(currentDate.getDate()).padStart(2, '0');
    const hours = String(currentDate.getHours()).padStart(2, '0');
    const minutes = String(currentDate.getMinutes()).padStart(2, '0');
    const seconds = String(currentDate.getSeconds()).padStart(2, '0');
    const formattedDate = `${year}${month}${day}${hours}${minutes}${seconds}`;
    let number = generateTwoDigitNumber()
    meetingId.value = `YP` + formattedDate + number
    timeer.value = `${year}-${month}-${day}`
    date.value = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
    userId.value = userInfoData.userInfo.userId
    personal.value = userInfoData.userInfo.userName
    reasonForBorrowing.value = `${year}-${month}-${day} ${personal.value}的音频会议`;
    corpCode.value = localStorage.getItem('corpCode') //从本地获取corpCode
    doc.querySelector(".recpowerx").style.width = 0 + "%";
    doc.querySelector(".recpowert").innerText = '';
})

</script>

<style lang="less" scoped>
.speech-box {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 2vw;
  height: calc(100vh - var(--van-nav-bar-height) - 4vh);

  .animation-box {
    height: 10vh;
    padding: 2vw 0;
    width: 100%;
  }

  .speech-list-box {
    //padding-top: 2vh;
    //background-color: #fff;
    flex: 1;
    width: 100%;
    overflow-y: auto; /* 当内容超出高度时显示垂直滚动条 */
  }

  .btn-box {
    height: 20vh;
    width: 100%;
  }
}

.cardTable {
  width: 100%;
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
  width: 96%;
  border-radius: 10px;
  //height: 18vh;
  background-color: #fff;
  padding: 2vw;

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

  .header_title {
    display: flex;
    align-items: center;
    font-size: 1em;
    margin-bottom: 2vw;

    span {
      margin-left: 10px;
    }
  }
}
.form{
    display: flex;
}
.f-w-550 {
  font-weight: 550;
}

.f-z-14-c {
  color: rgba(23, 26, 29, 0.40);
  font-size: 0.875em;
}

.m-t-10 {
  margin-top: 4vw;
}

.f-z-16 {
  font-size: 16px;
}

.metting {
  display: flex;
  justify-content: space-around;
  text-align: center;
  font-size: 12px;
  // justify-content: center;
}

.round-button {
  display: flex;
  justify-content: space-around;
  text-align: center;
  font-size: 12px;
  padding-top: 2vw;
}

.footer {
  width: 100%;
}

.list {
  display: flex;
  justify-content: space-between;
  padding: 20px;
  //margin: 20px;
  background-color: #fff;
  //border-radius: 8px;
  border-bottom: 1px solid #ccc;
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

///* 小屏幕设备 */
//@media screen and (max-width: 100px) {
//  .scrollable-div {
//    height: 28vh; /* 设置列表高度为屏幕高度的60% */
//    overflow-y: auto;
//  }
//}
//
///* 大屏幕设备 */
//@media screen and (min-width: 101px) {
//  .scrollable-div {
//    height: 30vh; /* 设置列表高度为屏幕高度的80% */
//    overflow-y: auto;
//  }
//}
.empty {
  font-size: 12px;
  color: #999a9e;
}

::v-deep(.van-nav-bar__content) {
  background-color: #0088ff;
}

::v-deep(.van-nav-bar__title) {
  color: #ffffff;
}

::v-deep(.van-nav-bar__text) {
  color: #ffffff;
}

::v-deep(.van-icon-arrow-left:before) {
  color: #ffffff;
}

::v-deep(.van-cell) {
  padding: 0;
}
.m-b-10{
    margin-bottom: 10px;
}
</style>
