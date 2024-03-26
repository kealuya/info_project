<!--
  选择成本中心组件（单选）
  使用方法如下： 可参考差旅申请
-->
<template>
  <div class="selectPersonPicker">
    <van-sticky>
      <navbar title="选择成本中心" left-arrow left-text="返回" @leftEvent="$emit('selectOver',[])" />
      <form action="/">
        <van-row justify="center" align="center" style=" width:100%; background: #ffffff">
          <van-col span="18">
            <van-search
                v-model="value"
                placeholder="请输入成本中心名称"
                shape="round"
                @search="onSearch"
            />
          </van-col>
          <van-col  span="6" class="tc" @click="show=true">
              <span class="f14">{{costType}}</span>
              <van-icon class='ml5' name="arrow-down" />
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
          <div style="text-align:center;color: rgba(23,26,29,0.40)">暂无可以用的成本中心</div>
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
            :immediate-check="false"
            loading-text="正在加载中请稍后"
            finished-text="没有更多了"
            @load="onLoad"
            :offset="10"
        >
          <div class="ticket_content" v-for="(item,index) in list" :key="item">
            <cost-center-card :item="item"  @click="costCenterOK(item)" :check-key="costCostData" />
          </div>
        </van-list>
      </van-pull-refresh>
    </div>
    <van-action-sheet v-model:show="show" :actions="actions"   description="请点击选择" @select="onSelect" />
    <load-component v-if="loading"  />
  </div>
</template>

<script lang="ts" setup>
import CostCenterCard from './components/costCenterCard.vue'
import LoadComponent from '../../loading/index.vue'
import Navbar from '../../NavBar.vue'
const show = ref(false);
const actions = [
  { name: '全部',value:""},
  { name: '部门',value:"department"},
  { name: '项目',value:"cost"},
];
import {inject, ref, onMounted, defineEmits, onBeforeUpdate,} from "vue";
import {Toast,} from "vant";
import {getCostAndBudgetDeptList, } from "../../../services/current";
let userInfoData: any = inject("userInfo"); //获取审批数组
// @ts-ignore
import Pinyin from 'js-pinyin';
import {showFailToast} from "vant/lib/toast/function-call";
const value = ref('');
let list = ref<any>([]);//加载loading
let loading = ref(true); //加载效果
let refreshingOnly = ref(false); //加载效果只有下拉变化
let finished = ref(false); // 结束
const refreshing = ref(false); // 重新加载还是追加
const costType=ref('全部'); // 右侧切换状态样式
const emit=defineEmits(['selectOver'])
let payload = {
  "costType":"",
  "costDeptId":userInfoData.userInfo.deptId,
  "corpCode":userInfoData.userInfo.corpCode,
  "page":1,
  "pageSize":5,
  "keyWord":"",
  "equipmentType":""}
// 选中成本中心类型
const onSelect = (item:any) => {
  // 默认情况下点击选项时不会自动收起
  payload.costType=item.value;
  costType.value=item.name;
  refresh();
  show.value = false;
};
interface Props{};
const props=defineProps<{
  data:{
    xmbh:''
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
  loading.value=true;
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
  refresh();
}
// 选中后提交
const costCenterOK=(item:any)=>{
  costCostData.value=item.xmbh;
  emit('selectOver',item);
}

// 调用人员列表
const getCostCenterList=async(payload:any,code:any)=>{
  let result: any = await getCostAndBudgetDeptList(payload);
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
  costCostData.value=props.data.xmbh;
  refresh();
})
onBeforeUpdate(() => {
  // 文本内容应该与当前的 `count.value` 一致
  costCostData.value=props.data.xmbh;
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
