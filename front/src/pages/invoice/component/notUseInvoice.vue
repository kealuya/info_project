<template>
  <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
    <template #success>
      <span v-if="refreshingOnly==true"><van-icon name="success" /> 刷新成功</span>
    </template>
    <div v-if="loading==false && list.length==0" style="height: 500px">
      <van-row justify="center" align="center" class="tc">
        <div>
          <img style="height:50vh;width:80vw;object-fit: contain" src="../../../assets/img/zanwupiaoju.png" v-cloak/>
        </div>
        <div style="color: rgba(23,26,29,0.40)">暂无可报销的发票</div>
      </van-row>
      <van-row justify="center" align="center" class="tc mt5">
        <div>
          <van-button square plain type="primary" @click="onRefresh" style="padding: 5px 5px;border-radius: 8px">刷新
          </van-button>
        </div>
      </van-row>
    
    </div>
    <van-list
        v-model="loading"
        v-else
        :finished="finished"
        loading-text="正在加载中请稍后"
        finished-text="没有更多了"
        @load="onLoad"
        :offset="10"
    >
      <div class="ticket_content" v-for="(item,index) in list" :key="item">
        <van-swipe-cell>
          <airplane-component v-if="item.type=='5'" :get-data="item" @click="goInvoiceDetail(item)"/>
          <train-component v-else-if="item.type=='8'" :get-data="item" @click="goInvoiceDetail(item)"/>
          <invoice-component v-else :get-data="item" @click="goInvoiceDetail(item)"/>
          <template #right>
            <van-button square text="删除" type="danger" class="delete-button mr5" @click="delInv(item)"/>
          </template>
        </van-swipe-cell>
        <!--                <div>发票: {{ item.invName }}</div>-->
      </div>
    </van-list>
  </van-pull-refresh>
  
</template>

<script setup lang="ts">
import {getInvoiceWbList, deleteInv} from "../../../services/invoice";
import {ref, onActivated, reactive, onMounted, inject} from "vue";
// @ts-ignore
import InvoiceComponent from "./invoiceComponent.vue";
// @ts-ignore
import AirplaneComponent from "./airplaneComponent.vue";
// @ts-ignore
import TrainComponent from "./trainComponent.vue";
import {useRoute, useRouter} from "vue-router";
import {Dialog} from "vant";
import { Notify } from 'vant';
import Loadingcomponent from '../../../components/loading/index.vue'
let list = ref<any>([]);//加载loading
let loading = ref(true); //加载效果
let refreshingOnly = ref(false); //加载效果只有下拉变化
let finished = ref(false); // 结束
const refreshing = ref(false); // 重新加载还是追加
const router = useRouter();
let userInfoData: any = inject("userInfo");
let payload = {
  'companyCode': userInfoData.userInfo.corpCode,
  'corpCode': userInfoData.userInfo.corpCode,
  'empCode': userInfoData.userInfo.empCode,
  'row': '10',
  'page': '1',
  'zt': '0',
  'type': '',
}
const getInvoiceList=async(payload:any,code:any)=>{
  let result: any = await getInvoiceWbList(payload);
  if(code=='refresh'){
    list.value = [];
  }
  list.value.push(...result.data.data);
  if (list.value.length.toString() === result.data.count){
    loading.value = false;
    finished.value = true;
  } else {
    let page = Number(payload.page) + 1;
    payload.page = page.toString();
    loading.value = false;
  }
  refreshing.value = false;
}
const onLoad = async (code:any) => {
  if(code=='1'){
    payload.page = "1";
  }
  if (refreshing.value) {
    return;
  }
  if(finished.value==true){
    return
  }
  await getInvoiceList(payload,'loading')
};
const refresh=async()=>{
  if (refreshing.value) {
  }
  if(finished.value==true){
    return
  }
  await getInvoiceList(payload,'refresh')
}
const onRefresh = (code:any) => {
  payload.page = "1";
  refreshing.value = true;
  if(code!='1'){
    refreshingOnly.value = true;
  }else{
    refreshingOnly.value = false;
  }
  finished.value = false;
  loading.value = true;
  refresh();
};
//点击跳转到发票详情页
const goInvoiceDetail = (item: any) => {
  sessionStorage.setItem('invoiceDetail', JSON.stringify(item))
  router.push('/detail');
};
//删除发票
const delInv = async (item: any) => {
  Dialog.confirm({
    title: '您确定要删除此发票吗',
    confirmButtonColor:'#4BA3FB'
  }).then(async () => {
    let payload = {
      corpCode: userInfoData.userInfo.corpCode,
      dzfph: item.dzfph,
      empCode: userInfoData.userInfo.empCode
    }
    let res: any = await deleteInv(payload);
    if (JSON.parse(res).success === true) {
      Notify({
        type: 'success',
        message: '删除成功'
      })
      await onRefresh('0')
    } else {
      Notify({
        type: 'danger',
        message: '删除失败'
      })
    }
  })
  
}

defineExpose({
  getInvoiceList, onLoad, onRefresh,refresh, payload, list
})
</script>

<style lang="less" scoped>
.van-button {
  height: 100%;
  border-radius: 12px 12px 12px 12px;
  font-size: 15px;
  width: 70px;
}
</style>
