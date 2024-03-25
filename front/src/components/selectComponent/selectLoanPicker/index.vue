<!--
  选择借款单组件（单选）
  使用方法如下： 可参考日常报销
-->
<template>
  <div class="selectPersonPicker">
    <van-sticky>
      <Navbar  title="选择借款单" left-arrow left-text="返回" @leftEvent="$emit('selectOver',[])"  />
      <form action="/">
        <van-row justify="center" align="center" style=" width:100%; background: #ffffff">
          <van-col span="22">
            <van-search
                v-model="value"
                placeholder="请输入申请单名称"
                shape="round"
                @search="onSearch"
            />
          </van-col>
        </van-row>
      </form>
    </van-sticky>
    <div>
      <van-pull-refresh v-model="refreshing" @refresh="onRefresh" style="background:#f1f2f3; min-height: 100vh;">
        <template #success>
          <span v-if="refreshingOnly==true"><van-icon name="success" /> 刷新成功</span>
        </template>
        <div v-if="loading==false && list.length==0" style="height: 500px">
          <van-row justify="center" align="center" class="tc">
            <div>
              <img style="height:50vh;width:80vw;object-fit: contain" src="../../../assets/img/zanwupiaoju.png" v-cloak/>
            </div>
          </van-row>
          <div style="color: rgba(23,26,29,0.40);text-align: center">暂无可以用的借款单</div>
          <van-row justify="center" align="center" class="tc mt5">
          </van-row>

        </div>
        <van-list
            v-model="loading"
            v-else
            :finished="finished"
            :immediate-check="false"
            loading-text="正在加载中请稍后"
            finished-text="没有更多了"
            @load="onLoad"
            :offset="10"
        >
          <div class="ticket_content" v-for="(item,index) in list" :key="item">
            <loan-card :get-data="item"  @click="costCenterOK(item)" :check-key="costCostData" />
          </div>
        </van-list>
      </van-pull-refresh>
    </div>
    <load-component v-if="loading"  />
  </div>
</template>

<script lang="ts" setup>
 import LoanCard from './components/loanCard.vue'
 import LoadComponent from '../../loading/index.vue'
import {getLoanList} from "../../../services/reimburse";
 import Navbar from '../../../components/NavBar.vue'
const show = ref(false);
import {inject, ref, onMounted, defineEmits, onBeforeUpdate,} from "vue";
let userInfoData: any = inject("userInfo"); //获取审批数组
// @ts-ignore
import Pinyin from 'js-pinyin';
const value = ref('');
let list = ref<any>([]);//加载loading
let loading = ref(true); //加载效果
let refreshingOnly = ref(false); //加载效果只有下拉变化
let finished = ref(false); // 结束
const refreshing = ref(false); // 重新加载还是追加
const emit=defineEmits(['selectOver'])
let payload ={
  "page":1,
  "pageSize":5,
  "corpCode":"szht6666",
  "empCode":"1815",
  "keyWord":"",
  "state":"1",
  "ishk":"0",
  "relationState":"true",
  "equipmentType":"PC"
}
interface Props{};
const props=defineProps<{
  data:{
    applyno:''
  }
}>()
const costCostData =ref({});
// 分页加载
const onLoad = async (code:any) => {
  if(code=='1'){
    payload.page = 1;
  }
  if (refreshing.value) {
    return;
  }
  if(finished.value==true){
    return
  }
  await getCostCenterList(payload,'loading')
};
// 首次进入刷新
const refresh=async()=>{
  if (refreshing.value) {
  }
  if(finished.value==true){
    return
  }
  payload.page = 1;
  loading.value = true;
  await getCostCenterList(payload,'refresh')
}
// 刷新处理
const onRefresh = (code:any) => {
  payload.page = 1;
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
// 根据成本中心内容搜索相关内容
const onSearch = async ()=>{
  payload.keyWord=value.value;
  onRefresh('1');
}
// 选中后提交
const costCenterOK=(item:any)=>{
  costCostData.value=item.approveno;
  emit('selectOver',item);
}

// 调用人员列表
const getCostCenterList=async(payload:any,code:any)=>{
  let result: any = await getLoanList(payload);
  if(code=='refresh'){
    list.value = [];
  }
  list.value.push(...result.data.records);
  if (list.value.length.toString() === result.data.total){
    finished.value = true;
    loading.value = false;
  } else {
    let page = Number(payload.page) + 1;
    payload.page = page.toString();
    loading.value = false;
  }
  refreshing.value = false;
}
onMounted(()=> {
  costCostData.value=props.data.applyno;
  refresh();
})
onBeforeUpdate(() => {
  // 文本内容应该与当前的 `count.value` 一致
  costCostData.value=props.data.applyno;
})
</script>

<style scoped>
.selectPersonPicker{
  /*height: 100%;*/
  display: flex;
  flex-direction: column;
}
.approveBottomGroup div{
  display: inline-block;
  width: 70px;
  text-align: center;
  padding-top: 10px;
  color:#FFFFFF;
  font-size: 14px;
  height: 30px;
}
</style>
