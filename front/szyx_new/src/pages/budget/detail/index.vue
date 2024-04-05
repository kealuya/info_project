<template>
<nav-bar title="预算编制管理"  />
  <div class="formCard">
    <van-row align="center">
      <van-col>
        <div class="ml10 bold">{{budgetInfo.budGetingName}}</div>
      </van-col>
    </van-row>
    <div class="mt10">
      <van-row justify="space-between" align="center">
        <van-col>
          <div class="ml10  graytxt">预算编号:</div>
          <div class="ml10">{{budgetInfo.budGetingId}}</div>
          <div class="ml10  graytxt">所属负责人:</div>
          <div class="ml10 ">{{budgetInfo.creation}}</div>
          <div class="ml10  graytxt">审定金额:</div>
          <div class="ml10 orange f18 bold">¥ {{ Number(budgetInfo.budGetingAmount).toFixed(2)}}</div>
        </van-col>
        <van-col>
          <div class="mr10">
            <approval-status  :type="budgetInfo.approvalResult=='1'?'已同意':'已拒绝'" />
          </div>
        </van-col>
      </van-row>
    </div>
  </div>
  <div class="formCard">
    <div class="mt5 ml10">
      <van-row align="center">
        <van-col>
          <span class="graytxt">所属年度：</span>
        </van-col>
        <van-col>{{budgetInfo.budGetingYear}}年度</van-col>
      </van-row>
    </div>
    <div class="mt5 ml10">
      <van-row align="center">
        <van-col>
          <span class="graytxt">开始时间：</span>
        </van-col>
        <van-col>{{budgetInfo.budGetingStartTime}}</van-col>
      </van-row>
    </div>
    <div class="mt5 ml10">
      <van-row align="center">
        <van-col>
          <span class="graytxt">结束时间：</span>
        </van-col>
        <van-col>{{budgetInfo.budGetingEndTime}}</van-col>
      </van-row>
    </div>
    <div class="mt5 ml10">
      <van-row align="center">
        <van-col>
          <span class="graytxt">创建时间：</span>
        </van-col>
        <van-col>{{budgetInfo.creationTime}}</van-col>
      </van-row>
    </div>
  </div>
  <div class="formCard">
    <van-row justify="space-between" align="center">
      <van-col>
        <span class="ml10 bold">预算详情</span>
      </van-col>
      <van-col>
          <div  class="mr10" @click="showPicker=true">
            <span style="color:#0088ff" class="f14">{{infoType}}</span>
            <van-icon name="arrow-down" class="ml5" size="10" />
          </div>
      </van-col>
    </van-row>
    <div v-if="infoType!='项目预算'">
      <div v-for="item in deptbudgetList">
        <div class="tips">
          <van-row align="center">
            <van-col> <span class="bold f18"> {{item.deptName}}</span> </van-col>
          </van-row>
          <van-row align="center">
            <van-col span="6"> <div class="typeTips">部门预算</div></van-col>
          </van-row>
          <div class="mt5">
            <van-row>
              <van-col  class="graytxt f14 bold">预算额度：</van-col>
              <van-col><span class="orange f14 bold">¥ {{Number(item.deptBudGetingAmount).toFixed(2)}}</span></van-col>
            </van-row>
            <van-row>
              <van-col  class="graytxt f14 bold">预算额度：</van-col>
              <van-col><span class="orange f14 bold">¥ {{Number(item.deptBudGetingYe).toFixed(2)}}</span></van-col>
            </van-row>
          </div>
        </div>
      </div>
    </div>
    <div v-if="infoType!='部门预算'">
      <div v-for="item in costRelationBudget">
        <div class="tips">
          <van-row align="center">
            <van-col> <span class="bold f18"> {{item.xmmc}}</span> </van-col>
          </van-row>
          <van-row align='center'>
            <van-col span="6"><div class="typeTips">项目预算</div></van-col>
            <van-col span="12" offset="1"> <div class="typeTips">{{item.xmbh}}</div></van-col>
          </van-row>
          <div class="mt5">
            <van-row>
              <van-col span="6" class="graytxt f14 bold">预算额度：</van-col>
              <van-col span="18"><span class="orange f14 bold">¥ {{Number(item.xmje).toFixed(2)}}</span></van-col>
            </van-row>
            <van-row>
              <van-col span="6"  class="graytxt f14 bold">预算余度：</van-col>
              <van-col  span="18"><span class="orange f14 bold">¥ {{Number(item.xmye).toFixed(2)}}</span></van-col>
            </van-row>
            <van-row>
              <van-col  span="6" class="graytxt f14 bold">消费科目：</van-col>
              <van-col  span="18"><span class="graytxt f14 bold">{{item.subjectNameList}}</span></van-col>
            </van-row>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div v-show="loading">
    <loading-component :title="'加载中'"/>
  </div>
  <van-popup v-model:show="showPicker" round position="bottom">
    <van-picker
        :columns="infoList"
        @cancel="showPicker = false"
        @confirm="onConfirm"
    />
  </van-popup>
</template>

<script lang="ts" setup>
import NavBar from '../../../components/navBar/index.vue'
import {inject, onBeforeMount, onUnmounted, ref} from "vue";
import LoadingComponent from '../../../components/loading/index.vue'
import ApprovalStatus from '../../../components/approvalStatus/index.vue'
import {getBudgetDetail} from "../../../services/budget";
const loading = ref<boolean>(false); //load加载
const apprInfo = ref({}); // 审批详情信息
let userInfoData: any = inject("userInfo"); // 用户信息
let budgetInfo=ref({}); // 用于接收接口传递过来的预算编制详情
let infoType=ref('全部'); // 用于切换显示对应的预算编制详情信息
let deptbudgetList=ref([]); // 用于接收接口传递过来的预算编制部门信息列表
let costRelationBudget=ref([]); // 用于接收接口传递过来的预算编制成本中心信息列表
const showPicker=ref(false); // 切换详情类型
const infoList=ref([{text:'全部',value: ''},{text:'部门预算',value: 'dept'},{text:'项目预算',value:'cost'}]); // 预算编制类型
const  onConfirm=(value:any)=>{
  infoType.value=value.selectedOptions[0].text;
  showPicker.value=false;
} //切换预算编制
//获取预算编制详情
const getLoanApply = async (data: any) => {
  const payload: any = {
    "budGetingId": data.approveNo,
    "approveType": data.approveType,
    "corpCode": userInfoData.userInfo.corpCode,
    "empCode": userInfoData.userInfo.empCode,
    "equipmentType": ""
  }
  let result: any = await getBudgetDetail(payload);
  budgetInfo.value=result.data.budgeting;
  deptbudgetList.value=result.data.deptbudgetList;
  costRelationBudget.value=result.data.costRelationBudget;
  loading.value = false;
}
onBeforeMount(async () => {
  let applicationDetail: any = sessionStorage.getItem('budgetDetail');
  loading.value = true;
  apprInfo.value = JSON.parse(applicationDetail);
  await getLoanApply(apprInfo.value);
})
onUnmounted(()=>{
  sessionStorage.removeItem('budgetDetail');
})
</script>

<style scoped>
.formCard {
  margin: 10px 5px;
  padding-top: 10px;
  padding-bottom: 5px;
  background: white;
  border-radius: 10px;
}
.tips{
  padding: 5px;
  margin: 10px;
  border: 1px solid #0088ff;
  border-radius: 10px;
}
.typeTips{
  padding: 5px;
  margin-top: 10px;
  border: 1px solid #0088ff;
  color:#0088ff;
  font-size:12px;
  text-align: center;
  border-radius: 10px;
}
</style>
