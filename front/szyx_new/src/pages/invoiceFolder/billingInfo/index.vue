<template>
  <navbar title="开票信息详情" :icon="qrCodeIcon" @rightEvent="openQrCode(billInfoA)"/>
  <div class="card">
    <van-form>
      <van-field
          v-model="billInfoA.companyName"
          name="抬头名称"
          label="抬头名称"
          readonly
          placeholder="抬头名称"
      />
      <van-field
          v-model="billInfoA.taxpayerNo"
          name="税号"
          label="税号"
          readonly
          placeholder="税号"
      />
      <van-field
          v-model="billInfoA.companyAccount"
          name="开户银行"
          label="开户银行"
          readonly
          placeholder="开户银行"
      />
      <van-field
          v-model="billInfoA.companyNumber"
          name="银行卡号"
          label="银行卡号"
          readonly
          placeholder="银行卡号"
      />
      <van-field
          v-model="billInfoA.companyTel"
          name="联系电话"
          label="联系电话"
          readonly
          placeholder="联系电话"
      />
      <van-field
          v-model="billInfoA.companyAddress"
          type="textarea"
          name="注册地址"
          label="注册地址"
          rows="3"
          readonly
          placeholder="注册地址"
      />
    </van-form>
  </div>
  <van-overlay :show="show" @click="escQrCode">
    <div class="wrapper">
      <div class="block" @click.stop>
        <qrcode-vue :value="qrCode123" size="160" ref="qrcodeImg" id="qrcode-canvas" level="H" margin="10"
                    @click="downloadQrcode(billInfoA)"></qrcode-vue>
        <canvas/>
      </div>
      <p style="color:white" class="mt5 f14">点击保存至本地相册</p>
      <van-icon :name="closeIcon" size="25" color="#ffffff"  class="mt10" @click="escQrCode"/>
    </div>

  </van-overlay>
</template>

<script lang="ts" setup>
import Navbar from '../../../components/navBar/index.vue';
import {ref, onBeforeMount} from 'vue';
import qrCodeIcon from '../../../assets/icon/qrCode.png'
import closeIcon from '../../../assets/icon/close_Qrcode.png'
import QrcodeVue from 'qrcode.vue'
import {showSuccessToast} from "vant";
// 控制二维码显示隐藏
const show = ref(false);
// 调用保存二维码时可能会用到的组件
const qrcodeImg = ref();
// 开票信息
let billInfoA = ref({
  'companyName': '',
  'taxpayerNo': '',
  'companyAccount': '',
  'companyNumber': '',
  'companyAddress': '',
  'companyTel': ''
});
const qrCode123 = ref("")
// 生成二维码
const openQrCode = async (item: any) => {
  show.value = true;
  qrCode123.value = ` "抬头名称：${billInfoA.value.companyName}"
                              "税号：${billInfoA.value.taxpayerNo}"
                              "开户银行：${billInfoA.value.companyAccount}"
                              "银行账号：${billInfoA.value.companyNumber}"
                              "注册地址：${billInfoA.value.companyAddress}"
                              "电话号码：${billInfoA.value.companyTel}"`
}
// 保存二维码
const downloadQrcode = async (item: any) => {

  setTimeout(()=>{
    showSuccessToast('正在保存请稍后...');
    let canvas: any = document.getElementById('qrcode-canvas'); // 在模板中的canvas元素设置id为 qrcode-canvas
    var dataURL = canvas.toDataURL("image/png")
    let a = document.createElement('a')
    a.href = dataURL
    a.download = `${billInfoA.value.companyName}开票信息.png` // 报文到本地时的文件名称
    a.click();
    showSuccessToast(`${billInfoA.value.companyName}开票信息.png，已保存至您的手机相册`);
  },2000);

}
// 关闭二维码
const escQrCode = () => {
  show.value = false;
}
onBeforeMount(() => {
  let billInfo: any = sessionStorage.getItem('billingInfo');
  billInfo = JSON.parse(billInfo);
  billInfoA.value = billInfo;
})
</script>

<style scoped>
.wrapper {
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: center;
  height: 100%;
}

.block {
  width: 200px;
  height: 170px;
  text-align: center;
  padding-top: 10px;
  border-radius: 10px;
  background-color: #fff;
}

.escIcon {
  width: 30px;
  height: 30px;
  border: #FFFFFF 2px solid;
  border-radius: 20px;
}

.card {
  background: #fff;
  border-radius: 10px;
  padding: 15px;
  margin: 10px;
}
</style>
