<template>
  <navbar title="开票信息"/>
  <van-pull-refresh v-model="refreshing" @refresh="onRefresh" style="min-height:80vw">
    <template #success>
      <span v-if="refreshingOnly==true"><van-icon name="success"/> 刷新成功</span>
    </template>
    <div v-if="loading==false && list.length==0" style="height: 500px">
      <van-row justify="center" align="center" class="tc">
        <div>
          <img style="height:50vh;width:80vw;object-fit: contain" src="../../../assets/img/zanwupiaoju.png" v-cloak/>
        </div>
      </van-row>
      <div style="color: rgba(23,26,29,0.40);text-align: center">暂无发票信息</div>
    </div>
    <van-list
        v-model="loading"
        v-else
        :finished="finished"
        loading-text="正在加载中请稍后"
        finished-text="没有更多了"
        @load="onLoad"
        :offset="10"
        style="min-height:80vw"
    >
      <div class="ticket_content" v-for="(item,index) in list" :key="item">
        <div class="card">
          <van-row justify="space-between" align="center">
            <van-col @click="billingInfo(item)">
              <div class="bold">{{ item.companyName }}</div>
              <div class="f14 mt5 graytxt">税号：{{ item.taxpayerNo }}</div>
            </van-col>
            <van-col>
              <van-button plain type="primary" size="small" @click="copyBillingInfo(item)">复制</van-button>
            </van-col>
          </van-row>
        </div>
      </div>
    </van-list>
  </van-pull-refresh>
</template>

<script lang="ts" setup>
import Navbar from '../../../components/navBar/index.vue'
import {ref, inject,} from "vue";
import {getInvoiceBillingList} from "../../../services/invoice";
import {showFailToast} from "vant/lib/toast/function-call";
import {useRouter} from "vue-router";
const router = useRouter();
let list = ref<any>([]);//加载loading
let loading = ref(true); //加载效果
let refreshingOnly = ref(false); //加载效果只有下拉变化
let finished = ref(false); // 结束
const refreshing = ref(false); // 重新加载还是追加
let userInfoData: any = inject("userInfo");
let payload = {
  'corpCode': userInfoData.userInfo.corpCode,
  'page': 1,
  'deptId': userInfoData.userInfo.deptId,
  "pageSize": 10,
  "keyWord": "",
  "equipmentType": ''
}
// 点击进入开票信息
const billingInfo=(item:any)=>{
  sessionStorage.setItem('billingInfo', JSON.stringify(item));
  router.push('billingInfo');
}
//复制开票信息
const copyBillingInfo = (item: any) => {
  // 创建输入框元素
  let oInput = document.createElement('input');
  // 将想要复制的值
  oInput.value =
      `公司名称：${item.companyName}，公司税号：${item.taxpayerNo}，
      联系电话：${item.companyTel}，地址：${item.companyAddress}`;
  // 页面底部追加输入框
  document.body.appendChild(oInput);
  // 选中输入框
  oInput.select();
  // 执行浏览器复制命令
  document.execCommand('Copy');
  // 弹出复制成功信息
  // 复制后移除输入框
  oInput.remove();
  showFailToast('已复制到剪贴板!');
}
// 获取发票抬头信息列表
const getInvoiceList = async (payload: any, code: any) => {
  let result: any = await getInvoiceBillingList(payload);
  if (code == 'refresh') {
    list.value = [];
  }
  list.value.push(...result.data.list);
  if (list.value.length.toString() === result.data.total) {
    loading.value = false;
    finished.value = true;
  } else {
    let page = Number(payload.page) + 1;
    payload.page = page.toString();
    loading.value = false;
  }
  refreshing.value = false;
}
// 下拉加载
const onLoad = async (code: any) => {
  if (code == '1') {
    payload.page = 1;
  }
  if (refreshing.value) {
    return;
  }
  if (finished.value == true) {
    return
  }
  await getInvoiceList(payload, 'loading')
};
// 首次刷新
const refresh = async () => {
  if (refreshing.value) {
  }
  if (finished.value == true) {
    return
  }
  await getInvoiceList(payload, 'refresh')
}
// 刷新
const onRefresh = (code: any) => {
  payload.page = 1;
  refreshing.value = true;
  if (code != '1') {
    refreshingOnly.value = true;
  } else {
    refreshingOnly.value = false;
  }
  finished.value = false;
  loading.value = true;
  refresh();
};
</script>

<style scoped>
.card {
  background: #fff;
  border-radius: 10px;
  padding: 15px;
  margin: 10px;
}
</style>
