<template>
  <div>
    <navbar title="新增发票" />
    <div class="card">
      <van-row align="center">
        <van-col>
          <div class="f14 ml5 bold">录入方式</div>
        </van-col>
      </van-row>
      <div class="mt5 " style="padding-top: 10px;">
        <van-row justify="space-around" align="center">
          <van-col>
            <van-uploader accept="image/*"   :after-read="getFile" upload-text="拍照上传/选取照片" ref="uploaderRef">
              <div>
                <div class="ml5">
                  <van-image :src="cameraIcon" width="40" height="40" style="position: relative;z-index: 0;"/>
                </div>
                <div class="f12">拍照上传</div>
              </div>
            </van-uploader>
          </van-col>

          <van-col @click="linkToForm">
            <div class="ml5">
              <van-image :src="printIcon" width="40" height="40"/>
            </div>
            <div class="f12">手动录入</div>
          </van-col>
          <van-col>
            <van-uploader accept="application/pdf" :after-read="getOcrPdf">
              <div>
                <div class="ml5">
                  <van-image :src="fileIcon" width="40" height="40"/>
                </div>
                <div class="f12">文件上传</div>
              </div>
            </van-uploader>
          </van-col>
          <van-col @click="linkToWx">
            <div class="ml5">
              <van-image :src="wechatIcon" width="40" height="40"/>
            </div>
            <div class="f12">微信导入</div>
          </van-col>
        </van-row>
      </div>
      <div class="mt5 mb10" style="padding-top: 10px;">
        <van-row justify="space-around" align="center">
          <van-col @click="linkToMessage()">
            <div class="ml5">
              <van-image :src="messageIcon" width="40" height="40"/>
            </div>
            <div class="f12">短信导入</div>
          </van-col>
          <van-col @click="linkToEmail()">
            <div class="ml5">
              <van-image :src="emailIcon" width="40" height="40"/>
            </div>
            <div class="f12">邮箱导入</div>
          </van-col>
          <van-col @click="linkToClould()">
            <div class="ml5">
              <van-image :src="clouldboxIcon" width="40" height="40"/>
            </div>
            <div class="f12">云票据箱</div>
          </van-col>
          <van-col @click="linkToEmailRPA()">
            <div class="ml5">
              <van-image :src="emailRPAIcon" width="40" height="40"/>
            </div>
            <div class="f12">邮箱RPA</div>
          </van-col>
        </van-row>
      </div>
      <!--   拍照上传 -->
      <!--                  <input-->
      <!--                      type="file"-->
      <!--                      id="file"-->
      <!--                      accept="image/*"-->
      <!--                      @change="getPicture($event)"-->
      <!--                  />-->
    </div>
    <div v-show="loading">
      <loadingcomponent :title="'识别中'" />
    </div>
  </div>
  <!--  <camera ref="camera"></camera>-->
</template>

<script setup lang="ts">
import wx from "weixin-js-sdk-ts";
import { Notify} from 'vant';
import {UpLoadImg} from "../../../components/hw_obs";
import Loadingcomponent from '../../../components/loading/index.vue';
import Navbar from '../../../components/navBar/index.vue'
import { useRouter} from "vue-router";
let date = new Date();
let year = date.getFullYear();
let month = date.getMonth();
let day = date.getDate();
// import Camera from './components/camera.vue';
// const wx = window['wx'];
// const camera = ref();
/***
 * 使用时仅需添加一下两端代码即可，
 * 适用于在手机上查看log信息
 * 注：打包前必须注释掉
 * import Vconsole from 'vconsole'
 * const vconsole = new Vconsole();
 */
// TODO ：可以在移动端浏览器查看控制台信息
const show = ref(false)
const loading = ref<boolean>(false);
const router =useRouter();
import {
  getInvOcrResultByPdfWb,
  getInvOcrResultWb,
  getToken,
  getWxSignToken,
  getWxTokenTicket
} from "../../../services/invoice";
// 拍照
import cameraIcon from '../../../assets/icon/invoice_icon_05.png'
//  微信导入
import wechatIcon from '../../../assets/icon/invoice_icon_03.png'
// 手动导入
import printIcon from '../../../assets/icon/invoice_icon_04.png'
// 文件上传
import fileIcon from '../../../assets/icon/invoice_icon_02.png'
// 短信导入
import messageIcon from '../../../assets/icon/invoice_icon_09.png'
// 邮箱RPA
import emailRPAIcon from '../../../assets/icon/invoice_icon_08.png'
//邮箱导入
import emailIcon from '../../../assets/icon/invoice_icon_10.png'
// 云票据箱
import clouldboxIcon from '../../../assets/icon/invoice_icon_11.png'
import { compress } from 'image-conversion'
import {inject, ref} from "vue";
import {showFailToast} from "vant/lib/toast/function-call";

const showPickerCamera = ref<boolean>(false);
const showPickerFile = ref<boolean>(false);
const invoiceData: any = inject('invoiceData');
let userInfoData: any = inject("userInfo");
//存储上传文件
let fileList = ref([]);
const uploaderRef = ref();
const getFile = async (file: any) => {
  file.type=file.content.slice(5, 14)
  loading.value = true;
  if (/iphone|ipad|ipod|ios/i.test(navigator.userAgent)) {
    // 隐藏文件列表
    uploaderRef.value.showFileList = false;
  }
  let imageList = [];
  if (Array.isArray(file)) {
    for (let item of file) {
      imageList.push(item);
    }
  } else {
    imageList.push(file);
  }
  let obsResult = await UpLoadImg(imageList);
  if (obsResult.length !== 0) {
    let payload = {
      'zt': '1',
      'qyId': userInfoData.userInfo.corpCode,
      'empCode': userInfoData.userInfo.empCode,
      'isBase64': '1', //固定值
      'filePath': obsResult[0],
      'ocrInvType': '300',
    };
    let result: any = await getInvOcrResultWb(payload);
    if (result.success !== false) {
      // 存储图片地址
      loading.value = false;
      let okData =JSON.parse(result);
      if (okData.data && okData.data[0].code === 'normal'){
        okData.data[0].cutfile='';
        // // 因为日期识别不出来到后面转换格式 会出现Invalid date，所以如果识别不出来就默认去提交的当天
        // if(okData.data[0].kprq==''){
        //   okData.data[0].kprq= new Date(year, month, day);
        // }
        await invoiceData.setData({
          code:'identify',
          data: okData.data[0],
          imgUrl: obsResult[0]
        });
        await router.push('/addInvoiceFrom');
      }else{
        new Notify({type: 'warning', message: '图片未识别出结果，请尝试手动填写！'});
      }
    } else {
      loading.value = false;
      new Notify({type: 'danger', message: '请求超时，请您稍后重试'});
    }

  }
}
//上传文件
const getOcrPdf = async (file: any) => {
  let imageList = [];
  if (Array.isArray(file)) {
    for (let item of file) {
      imageList.push(item);
    }
  } else {
    imageList.push(file);
  }
  loading.value = true;
  let obsResult = await UpLoadImg(imageList);
  let payload = {
    'isBase64': false,
    'empCode': userInfoData.userInfo.empCode,
    //（固定）
    'isPdf': true,
    //（固定）
    'isLocalOrNetworkPdf': 0,
    //（固定）
    'pdfUrl': obsResult[0],
    //
    'qyId': userInfoData.userInfo.corpCode,
  }
  let result: any = await getInvOcrResultByPdfWb(payload);
  // 存储图片地址
  JSON.parse(result.data).base64='';
  console.log(result)
  let jsonInfo=JSON.parse(result.data);
  console.log(jsonInfo)
  await invoiceData.setData({
    code:'identify',
    data: JSON.parse(result.data),
    imgUrl: obsResult[0]
  });
  loading.value = false;
  await router.push('/addInvoiceFrom');
}
const onClickLeft = () => history.back();
/***
 * 使用时仅需添加一下两端代码即可，
 * 适用于在手机上查看log信息
 * 注：打包前必须注释掉
 * import Vconsole from 'vconsole'
 * const vconsole = new Vconsole();
 */
// TODO ：可以在移动端浏览器查看控制台信息
const linkToForm = async() => {
  await invoiceData.setData({
    data: {},
    code:'handMove',
    imgUrl:''
  });
  router.push('/addInvoiceFrom');
}
const linkToCameraOrChoose = () => {
  showPickerCamera.value = true;
}

const linkToWx = async () => {
    new Notify({type: 'warning', message: '暂未开放', duration: 1500,});
};

const payload = {
  'corpCode': userInfoData.userInfo.corpCode,
  'empCode': userInfoData.userInfo.empCode,
};
const messageUrl = ''; //短信导入url
const emailUrl = ''; //邮箱导入url
const couldUrl = ''; //云票据箱url
const rpAUrl = ''; //邮箱RPA
let url = '';

function accessTest(url: any) {
  var link = document.createElement('link')
  link.rel = "stylesheet"
  link.type = "text/css"
  // 这里设置需要检测的url
  link.href = "url"
  link.onload = function () {
  }
  link.onerror = function () {
    showFailToast('唤醒链接超时，请您稍后重试');
  }
  document.body.appendChild(link)
};
const onLoad = async (code: any) => {
  let result: any = await getToken(payload);
  if (code === 'cloud') {
    url = couldUrl + 'token=' + result.token
  } else if (code === 'message') {
    url = messageUrl + 'token=' + result.token
  } else if (code == 'email') {
    url = emailUrl + 'token=' + result.token
  } else {
    url = rpAUrl + 'token=' + result.token
  }
  sessionStorage.setItem('url', url)
};
// 短信导入
const linkToMessage = async () => {
  await onLoad('message');
  await router.push({
    name: 'iframe', query: {code: 'message'}
  });
};
// 邮箱导入 iframeCode='1'
const linkToEmail = async () => {
  await onLoad('email');
  await router.push({
    name: 'iframe', query: {code: 'email'},
  });
};
// 云票据箱
const linkToClould = async () => {
  await onLoad('cloud');
  await router.push({
    name: 'iframe', query: {code: 'cloud'},
  });
};
// 邮箱RPA
const linkToEmailRPA = async () => {
  await onLoad('emailRPA');
  await router.push({
    name: 'iframe', query: {code: 'emailRPA'}
  });
};
</script>

<style lang="less" scoped>
.loading_body {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}

</style>
