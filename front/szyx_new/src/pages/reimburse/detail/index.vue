<template>
  <div>
    <nav-bar :title="title" />
    <!--  详情  -->
    <div class="formCard">
      <head-card :title="title" :data="dailyDetailDto"  />
    </div>
    <div class="formCard p10">
      <info-card :data="dailyDetailDto" :type="reimburseType" />
    </div>
    <div class="formCard p10">
      <div class="f18">费用信息</div>
      <div v-for="item in billList">
        <cost-card :data="item" />
      </div>
    </div>
    <div class="formCard p10">
      <div class="f18">发票信息</div>
      <div v-for="item in spendDetailList">
        <invoice-card :data="item" />
      </div>
    </div>
    <div class="formCard p10">
      <div class="f18">冻结记录</div>
      <van-steps direction="vertical" :active="0">
        <div v-for="item in freezingList">
          <van-step>
            <p class="bold">{{item.creationTime}}</p>
            <freezing-card :data="item" />
          </van-step>
        </div>
      </van-steps>
    </div>
    <div class="formCard p10" v-if="loanDetail!=null">
      <loan-card :data="loanDetail" />
    </div>
    <div class="formCard p10" v-else>
      <!--  收款信息    -->
      <pay-card :data="dailyDetailDto" />
    </div>

    <div class="formCard">
      <div >
        <van-row align="center">
          <span class="ml10">审批进度</span>
        </van-row>
        <div style="margin-top: 10px;margin-bottom: 10px">
          <van-steps direction="vertical" :active="nextApprIndex">
            <van-step>
              <van-row>
                <van-col :span="4">
                  <div class="f14 graytxt bold">发起</div>
                </van-col>
                <van-col :span="20">
                  <div style="padding-top: 3px;padding-bottom: 3px;background: #f6f6f6">
                    <div class="approveCard">
                      <van-row>
                        <van-col :span="6" class="f12 ml5">{{ dailyDetailDto.empname }}</van-col>
                        <van-col :span="16" class="f12 tr ml5">{{ dailyDetailDto.createtime }}</van-col>
                      </van-row>
                    </div>
                  </div>
                </van-col>
              </van-row>
            </van-step>
            <van-step v-for="item in approveInfoList">
              <van-row>
                <van-col :span="4">
                  <van-row align="center">
                    <van-col>
                      <div class="f14 graytxt bold">
                        <div v-if="item.list.length==1">单审</div>
                        <div v-else>{{ item.isAllPass == '1' ? '或审' : '会审' }}</div>
                        <div>
                          <span v-if="item.result=='ok'" class="ok f12 bold">通过</span>
                          <span v-else-if="item.result=='ng'" class="ng f12 bold">已拒绝</span>
                          <span v-else class="graytxt f12 bold">待审批</span>
                        </div>
                      </div>
                    </van-col>
                    <van-col>
                      <van-icon name="info-o" />
                    </van-col>
                  </van-row>
                </van-col>
                <van-col :span="20">
                  <div style="padding-top: 3px;padding-bottom: 3px;background: #f6f6f6">
                    <div v-for="key in item.list">
                      <div class="approveCard">
                        <van-row>
                          <van-col :span="12" class="f12 ml5">{{ key.toUser }}({{ key.toRole }})</van-col>
                          <van-col :span="10" class="f12 tr ml5">
                            <span class="ok bold" v-if="key.apprResult=='ok'">{{
                                stateLoad(key.state, key.apprResult)
                              }}</span>
                            <span class="ng bold"
                                  v-else-if="key.apprResult=='ng'">{{ stateLoad(key.state, key.apprResult) }}</span>
                            <span class="graytxt bold" v-else>{{ stateLoad(key.state, key.apprResult) }}</span>
                          </van-col>
                        </van-row>
                        <div v-if="key.state=='over'" class="f12 ml5 mt5">审批意见：{{ key.apprContent }}</div>
                        <div v-if="key.state=='over'" class="f12 ml5 mt5 graytxt">{{ key.modifyTime }}</div>
                      </div>
                    </div>
                  </div>
                </van-col>
              </van-row>
            </van-step>
          </van-steps>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import NavBar from "../../../components/navBar/index.vue";
import {inject, onMounted, ref} from "vue";
import {getDailyBxDetil,getFreezingThawingHistory} from "../../../services/reimburse";
import PayCard from './component/payCard/index.vue';
import HeadCard from './component/headDetail/index.vue';
import InfoCard from './component/infoCard/index.vue';
import CostCard from './component/costCard/index.vue';
import InvoiceCard from './component/invoiceCard/index.vue';
import FreezingCard from './component/freezingDetail/index.vue';
import LoanCard from './component/loanCard/index.vue';
import {approvalHistory} from "../../../services/current";
let title=ref('日常报销单'); // 详情
let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
let payload=ref({
  "empCode":userInfoData.userInfo.empCode,
  "corpCode":userInfoData.userInfo.corpCode,
  "applyNo":"",
  "equipmentType":""
}); // 提交信息
// 申请单详情
let reimburseInfo=ref({}); // 接收信息
// 日常报销信息
let dailyDetailDto=ref({
  reimburseType:''
}); // 日常报销详情
// 费用信息
let billList=ref([]);  // 费用信息
// 成本中心
let costcenter=ref({}); // 成本中心
let spendDetailList=ref([]); // 发票数组
let reimburseType=ref(''); // 报销类型
let freezingList=ref([]); // 冻结记录数组
let loanDetail=ref({}); // 借款申请单
let approveInfoList=ref([]); // 审批信息列表
let apprApprovalCcList=ref({stateMap:{}}); // 抄送人信息列表
let apprCcList=ref([]); //抄送信息列表
const nextApprIndex = ref(''); //记录当前节点用
//获取申请单详情
const getDailyBx = async (payload: any) => {
  let result: any = await getDailyBxDetil(payload.value);
  reimburseInfo.value=result.data;
  dailyDetailDto.value=result.data.dailyDetailDto;
  billList.value=result.data.billList;
  costcenter.value=result.data.costcenter;
  loanDetail.value=result.data.loanDetail;
  spendDetailList.value=result.data.spendDetailList;
}
//获取申请单冻结详情
const getFreezing = async (data:any) => {
  let resData={
    "corpCode":userInfoData.userInfo.corpCode,
    "approveNo":data.approveNo,
    "equipmentType":""
  }
  let result: any = await getFreezingThawingHistory(resData);
  freezingList.value=result.data;
}
// 根据申请单号以及类型判断该报销单类型
const applyType=(item:any)=>{
  var str = item.approveNo.substring(0, 2);
  if (str == 'HK') {
    reimburseType.value = '冲借款';
  } else {
    reimburseType.value = '日常报销';
  }
}
// 获取申请单审批流程历史
const getApprovalHistory = async (data: any) => {
  const payload: any = {
    "applyNo": data.approveNo,
    "corpCode": userInfoData.userInfo.corpCode,
    "empCode": userInfoData.userInfo.empCode,
    "equipmentType": ""
  }
  let result: any = await approvalHistory(payload);
  approveInfoList.value = [];
  apprApprovalCcList.value = {stateMap: {}};
  apprCcList.value = [];
  apprApprovalCcList.value = JSON.parse(result.data).data;
  apprCcList.value = JSON.parse(result.data).data.apprCcList;
  const apprInfo: any = apprApprovalCcList.value.stateMap
  for (let key in apprInfo) {
    const apprData = {
      no: key,
      list: apprInfo[key],
      isAllPass: apprInfo[key][0].isAllPass,
      state: apprInfo[key][0].state,
      result: apprInfo[key][0].apprResult,
    }
    if (apprInfo[key][0].state == 'wait') {
      nextApprIndex.value = key;
    }
    approveInfoList.value.push(apprData as never);
  }
  // loading.value = false;
}
onMounted(() => {
  let reimburseDetail: any = sessionStorage.getItem('reimburseDetail');
  reimburseDetail = JSON.parse(reimburseDetail);
  // 根据类型显示对应详情
  if(reimburseDetail.bxType=='rcbx'){
    title.value='日常报销单';
    payload.value.applyNo=reimburseDetail.approveNo
    getDailyBx(payload);
    getFreezing(reimburseDetail);
    applyType(reimburseDetail);
    getApprovalHistory(reimburseDetail);
  }else if(reimburseDetail.bxType=='travelbx'){
    title.value='差旅补助报销单'
  }else{
    title.value='借款申请单'
  }

});
// 根据状态显示审批状态
const stateLoad = (key: String, state: String) => {
  if (state == 'ok' && key == 'over') {
    return '已通过'
  } else if (state == 'ng' && key == 'over') {
    return '已拒绝'
  } else {
    return '待审批'
  }

}
</script>

<style scoped>
.formCard {
  margin: 10px 5px;
  padding-top: 10px;
  padding-bottom: 5px;
  background: white;
  border-radius: 10px;
}
.approveCard {
  margin: 5px;
  padding-top: 10px;
  padding-bottom: 5px;
  background: white;
  border-radius: 5px;
}
</style>
