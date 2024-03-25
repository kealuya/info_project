<template>
  <!--  此处是预算编制列表 用于index的type切换-->
  <van-pull-refresh v-model="refreshing" @refresh="onRefresh" style="min-height: 80vh">
    <template #success>
      <span v-if="refreshingOnly==true"><van-icon name="success" /> 刷新成功</span>
    </template>
    <div v-if="loading==false && list.length==0" style="height: 500px">
      <van-row justify="center" align="center" class="tc">
        <div>
          <img style="height:50vh;width:80vw;object-fit: contain" src="../../../../assets/img/zanwupiaoju.png" v-cloak/>
        </div>
      </van-row>
      <div style="color: rgba(23,26,29,0.40);text-align: center">暂无成本中心</div>
    </div>
    <van-list
        v-model="loading"
        v-else
        :finished="finished"
        loading-text="正在加载中请稍后"
        finished-text="没有更多了"
        :immediate-check="false"
        @load="onLoad"
        style="min-height:80vh"
        :offset="10"
    >
      <div class="ticket_content" v-for="(item,index) in list" :key="item">
        <cost-center-card  :get-data="item" />
      </div>
    </van-list>
  </van-pull-refresh>
  <div v-show="loadingComponent">
    <loading-component :title="'提交中'"/>
  </div>
</template>
<script lang="ts" setup>
import {getCostCenterList} from "../../../../services/budget";
import LoadingComponent from '../../../../components/loading/index.vue';
import CostCenterCard from '../costCenterCard/index.vue'
import { ref,onMounted,inject} from "vue";
const loadingComponent=ref(false); //loading组件控制显示隐藏
let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
let list = ref<any>([]);//加载loading
let loading = ref(true); //加载效果]
let refreshingOnly = ref(false); //加载效果只有下拉变化
let finished = ref(false); // 结束
const refreshing = ref(false); // 重新加载还是追加
const nowDate = new Date(); // 获取当前时间
// 提交参数
let payload={
  'corpCode': userInfoData.userInfo.corpCode,
  "keyWord":"",
  "budGetingYear":`${nowDate.getFullYear()}`,
  "page":1,
  "pageSize":10,
  "equipmentType":""
};
// 处理预算编制的接口
const getBudget=async(payload:any,code:any)=>{
  let result: any = await getCostCenterList(payload);
  if(code=='refresh'){
    list.value = [];
  }
  list.value.push(...result.data.records);
  loadingComponent.value=false;
  if (list.value.length.toString() === result.data.total){
    loading.value = false;
    finished.value = true;
  } else {
    payload.page = payload.page + 1;
    loading.value = false;
  }
  refreshing.value = false;
}
// 追加onload
const onLoad = async (code:any) => {
  if (refreshing.value) {
    return;
  }
  if(finished.value==true){
    return
  }
  await getBudget(payload,'loading')
};
// 刷新列表
const refresh=async()=>{
  loadingComponent.value=true;
  if (refreshing.value) {
  }
  if(finished.value==true){
    return
  }
  payload.page = 1;
  await getBudget(payload,'refresh')
}
// 首次刷新
const onRefresh = (code:any) => {
  payload.page = 1;
  if(code!='1'){
    refreshing.value = true;
    refreshingOnly.value = true;
  }else{
    refreshing.value = false;
    refreshingOnly.value = false;
  }
  finished.value = false;
  loading.value = true;
  refresh();
};
defineExpose({
  getBudget,onRefresh,refresh,payload
})
onMounted(() => {
  onRefresh('1');
});
</script>

<style scoped>

</style>
