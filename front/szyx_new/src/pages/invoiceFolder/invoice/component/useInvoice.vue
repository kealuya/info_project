<template>
    <van-pull-refresh v-model="refreshing" @refresh="onRefresh" >
      <template #success>
        <span v-if="refreshingOnly==true"> <van-icon name="success" />刷新成功</span>
      </template>
      <div  v-if="loading==false && list.length==0">
        <van-row justify="center" align="center" class="tc">
          <div>
            <img style="height:50vh;width:80vw;object-fit: contain" src="../../../../assets/img/zanwupiaoju.png" v-cloak/>
          </div>
        </van-row>
        <div style="color: rgba(23,26,29,0.40);text-align: center">您还没有已报销的发票</div>
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
            <div v-for="item in list" :key="item" @click="goInvoiceDetail(item)">
              <airplane-component v-if="item.type=='5'" :get-data="item" />
              <train-component   v-else-if="item.type=='8'" :get-data="item" />
              <invoice-component  v-else :get-data="item" />
            </div>
        </van-list>
    </van-pull-refresh>
</template>

<script setup lang="ts">
import {getInvoiceWbList} from "../../../../services/invoice";
import {onMounted, reactive, ref, onActivated, inject} from "vue";
import AirplaneComponent from "./airplaneComponent.vue";
import TrainComponent from "./trainComponent.vue";
import InvoiceComponent from "./invoiceComponent.vue";
import {useRouter} from "vue-router";
import {Notify} from "vant";
const router=useRouter();
let list = ref<any>([]);
let loading = ref(true);
let refreshingOnly = ref(false); //加载效果只有下拉变化
let show = ref(false); //加载loading
let finished = ref(false);
let userInfoData: any = inject("userInfo");
let payload = {
    'companyCode': userInfoData.userInfo.corpCode,
    'corpCode': userInfoData.userInfo.corpCode,
    'empCode': userInfoData.userInfo.empCode,
    'row': '10',
    'page': '1',
    'zt': '1',
    'type': '',
};
const getInvoiceUseList=async(payload:any,code:any)=>{
  let result: any = await getInvoiceWbList(payload);
  if(code=='refresh'){
    list.value = [];
  }
  list.value.push(...result.data.data.data);
  if (list.value.length.toString() === result.data.data.count) {
    loading.value = false;
    finished.value = true;
  } else {
    let page = Number(payload.page) + 1;
    payload.page = page.toString();
    loading.value = false;
  }
  refreshing.value = false;
}
const refreshing = ref(false);
const onLoad = async () => {
  if (refreshing.value) {
    list.value = [];

  }
  if(finished.value==true){
    return
  }
  show.value=true;
  await getInvoiceUseList(payload,'loading')
};
const refresh=async()=>{
  if (refreshing.value) {
  }
  if(finished.value==true){
    return
  }
  payload.page = "1";
  show.value=true;
  await getInvoiceUseList(payload,'refresh')
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

const goInvoiceDetail = (item:any) => {
    sessionStorage.setItem('invoiceDetail',JSON.stringify(item))
    router.push('./detail');
};
    defineExpose({
        getInvoiceUseList,onRefresh,payload
    })
</script>

<style lang="less" scoped>

</style>
