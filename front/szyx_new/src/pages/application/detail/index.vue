<template>
  <nav-bar title="单据申请" />
  <div>
    <head-detail :get-data="docDetail"/>
    <info-detail :list="rowShowJson" v-if="rowShowJson!=''"/>
    <div class="formCard" v-if="docDetail.approveType=='travel'">
      <div class="ml10">事由：{{ docDetail.approveReason }}</div>
      <div class="ml10 mt10">
        <div class="f14 graytxt">出差人员</div>
        <span v-for="(item,index) in docDetail.togetherPerson">
              <span class="f14 ">({{ item.deptName }}){{ item.togetherPersonName }} <span
                  v-if="docDetail.togetherPerson.length!=index+1">，</span></span>
           </span>
      </div>
      <div class="ml10 mt5">
        <van-row justify="center">
          <van-col :span="8">
            <div class="graytxt f14">出发时间</div>
            <div>{{ travelInfo.startTime }}</div>
          </van-col>
          <van-col :span="8">
            <div class="graytxt f14">结束时间</div>
            <div>{{ travelInfo.returnTime }}</div>
          </van-col>
          <van-col :span="8">
            <div class="graytxt f13">出行预算</div>
            <div>¥ {{ parseInt(travelInfo.travelBudget).toFixed(2) }}</div>
          </van-col>
        </van-row>
        <van-row justify="center">
          <van-col :span="8">
            <div class="graytxt f14">出发类型</div>
            <div>因公出行</div>
          </van-col>
          <van-col :span="8">
            <div class="graytxt f14">经办人</div>
            <div>{{ docDetail.creater }}</div>
          </van-col>
          <van-col :span="8">
            <div class="graytxt f13">所属部门</div>
            <div>{{ docDetail.deptName }}</div>
          </van-col>
        </van-row>
      </div>
    </div>
    <div class="formCard" v-if="docDetail.approveType=='travel'">
      <div class=" ml10 f16 bold">出行信息</div>
      <div class="travelCard" v-if="trainInfo.length!=0">
        <van-row align="center">
          <van-image :src="trainIcon" width="20" height="20"/>
          <span class="f14 bold ml5">火车出行</span>
        </van-row>
        <div v-for="item in trainInfo">
          <van-row justify="space-around" align="center" class="mt10 tc" v-if="item.travelBy=='train'">
            <van-col>
              <div>{{ item.fromCity }}</div>
              <div class="graytxt f14">出发城市</div>
            </van-col>
            <van-col>
              <van-image :src="arrowRight" width="100" height="10"/>
            </van-col>
            <van-col>
              <div>{{ item.toCity }}</div>
              <div class="graytxt f14">目的城市</div>
            </van-col>
          </van-row>
        </div>
      </div>
      <div class="travelCard" v-if="planeInfo.length!=0">
        <van-row align="center">
          <van-image :src="planeIcon" width="20" height="20"/>
          <span class="f14 bold ml5">飞机出行</span>
        </van-row>
        <div v-for="(item,index) in planeInfo">
          <van-row justify="space-around" align="center" class="mt10 tc" v-if="item.travelBy=='air'">
            <van-col>
              <div>{{ item.fromCity }}</div>
              <div class="graytxt f14">出发城市</div>
            </van-col>
            <van-col>
              <van-image :src="arrowRight" width="100" height="10"/>
            </van-col>
            <van-col>
              <div>{{ item.toCity }}</div>
              <div class="graytxt f14">目的城市</div>
            </van-col>
          </van-row>
        </div>
      </div>
      <div class="travelCard" v-if="hotelInfo.length!=0">
        <van-row align="center">
          <van-image :src="hotelIcon" width="20" height="20"/>
          <span class="f14 bold ml5">酒店</span>
        </van-row>
        <div v-for="item in hotelInfo">
          <van-row justify="space-around" class="mt10 tc" align="center" v-if="item.travelBy=='hotel'">
            <van-col>
              <div class="graytxt f14">入住城市</div>
            </van-col>
            <van-col>
              <div>{{ item.toCity }}</div>
            </van-col>
          </van-row>
        </div>
      </div>
    </div>
    <div class="formCard" v-if="docDetail.approveType=='travel'">
      <div class="f16 bold ml10">行程明细</div>
      <span v-for="item in docDetail.togetherPerson">
        <span class="f16 ml10" :class="orderIndex==item.togetherPersonId?'checkCard':'noCard'" @click="orderOver(item)">{{
            item.togetherPersonName
          }}</span>
      </span>
      <div>
        <div v-for="item in orderUserList">
          <div v-for="key in item.orderList">
            <div class="trainCard" v-if="key.type=='train'">
              <div class="trainCardHeader">
                <van-row justify="space-between">
                  <van-col class="ml10">火车票订单</van-col>
                  <van-col class="mr10">{{ key.orderno }}</van-col>
                </van-row>
              </div>
              <van-row justify="space-between" align="center">
                <van-col class="ml10 f14">{{ key.planBeginDate }} 出发</van-col>
                <van-col>
                  <div class="stateTip">{{ key.ticketStatus }}</div>
                </van-col>
              </van-row>
              <van-row class="ml10">
                <div class="f18 bold">{{ key.planBeginDate.substring(0, 11) }} {{ key.fromStation }} -
                  {{ key.planEndDate.substring(0, 11) }} {{ key.toStation }}
                </div>
              </van-row>
              <van-row class="ml10 mt5 mb5">
                <div class="f12">{{ key.trainNo }}{{ key.seatClass }}{{ key.seatNo }}</div>
              </van-row>
              <div class="mt5"></div>
            </div>
            <div class="trainCard" v-if="key.type=='hotel'">
              <div class="hotelCardHeader">
                <van-row justify="space-between">
                  <van-col class="ml10">酒店订单</van-col>
                  <van-col class="mr10">{{ key.orderno }}</van-col>
                </van-row>
              </div>
              <van-row justify="space-between" align="center">
                <van-col class="ml10 f14">{{ key.hotelName }}</van-col>
              </van-row>
              <van-row class="ml10">
                <div class="f18 bold">{{ key.gmtOccupancyTime }} 入住 - {{ key.gmtLeaveTime }} 离店</div>
              </van-row>
              <van-row class="ml10 mt5 mb5">
                <div class="f12">{{ key.roomName }}</div>
              </van-row>
              <div class="mt5"></div>
            </div>
            <div class="trainCard" v-if="key.type=='flight'">
              <div class="planeCardHeader">
                <van-row justify="space-between">
                  <van-col class="ml10">飞机票订单</van-col>
                  <van-col class="mr10">{{ key.orderno }}</van-col>
                </van-row>
              </div>
              <van-row justify="space-between" align="center">
                <van-col class="ml10 f14">{{ key.planBeginDate }} 出发</van-col>
                <van-col>
                  <div class="stateTip">{{ key.ticketStatus }}</div>
                </van-col>
              </van-row>
              <van-row class="ml10">
                <div class="f18 bold">{{ key.planBeginDate }} {{ key.startCityName }}{{ key.startPort }} -
                  {{ key.planEndDate }} 浦东T2
                </div>
              </van-row>
              <van-row class="ml10 mt5 mb5">
                <div class="f12">{{ key.airlineCompany }} {{ key.flightNo }} {{ key.seatClass }}
                  {{ key.flightDiscount }}
                </div>
              </van-row>
              <div class="mt5"></div>
            </div>
          </div>
        </div>
        <van-row justify="center" v-if="noKey">
          <van-col><span class="font_gray f14 mt5 mb5">暂无行程</span></van-col>
        </van-row>

      </div>
    </div>
    <div class="formCard">
      <div v-if="docDetail.approvalState!='免审'">
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
                        <van-col :span="6" class="f12 ml5">{{ docDetail.creater }}</van-col>
                        <van-col :span="16" class="f12 tr ml5">{{ docDetail.createTime }}</van-col>
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
                      <van-icon name="info-o" @click="approveToast(item.isAllPass,item.list.length)"/>
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
      <div v-if="apprCcList.length!=0">
        <span class=" f14 ml10">
         <span class="graytxt"> 抄送人：</span>
          <span v-for="item in apprCcList">
          <span class="f12"> {{ item.ccUser }}({{ item.toRole }})</span>
          </span>
        </span>
      </div>
    </div>

  </div>
  <div class="strck">
    <div style="margin: 0px;height: 60px;background: #FFFFFF;padding:10px 18px"
         v-if="docDetail.approvalState == '待审批'||docDetail.approvalState == '免审'">
      <van-button v-if="docDetail.approvalState == '待审批'" type="primary" @click="revokeDoc(docDetail)"
                  style="width: 90vw ">撤销
      </van-button>
      <van-button v-if="docDetail.approvalState == '免审'" @click="deleteDoc(docDetail)" type="primary"
                  style="width: 90vw ">删除
      </van-button>
    </div>
    <div style="margin: 0px;height: 60px;background: #FFFFFF;padding:10px 18px"
         v-if="docDetail.approvalState == '已通过'&&docDetail.approveType=='car'">
      <van-button type="primary" @click="goTakeTaxi(docDetail)" style="width: 90vw ">去打车
      </van-button>
    </div>

  </div>
  <div v-show="loading">
    <loadingcomponent :title="'加载中'"/>
  </div>
</template>

<script lang="ts" setup>
import {onBeforeMount, onUnmounted,inject, ref} from "vue";
import {approvalRevoke, deleteWorkInfo, getLoanApplyDetail, getPeopleOrderList} from "../../../services/apply";
import {approvalHistory} from '../../../services/current';
import HeadDetail from './components/headDetail/index.vue'
import InfoDetail from './components/infoDetail/index.vue'
import Loadingcomponent from '../../../components/loading/index.vue';
import 'vant/es/toast/style';
import {showSuccessToast} from "vant";
import {showConfirmDialog} from "vant";
import {showFailToast} from "vant/es";
import {getCarLoginUrl} from "../../../services/trirdPartyBusiness";
import trainIcon from '../../../assets/icon/icon_03.png';
import planeIcon from '../../../assets/icon/icon_01.png';
import hotelIcon from '../../../assets/icon/icon_02.png';
import arrowRight from '../../../assets/img/icon_arrow_right.png';
import router from "../../../router";
import NavBar from '../../../components/navBar/index.vue'
let userInfoData: any = inject("userInfo"); // 用户信息
const docDetail = ref({rowShowJson: '{}', approvalState: ''}); // 单据详情
const rowShowJson = ref([]); //中间显示
const apprApprovalCcList = ref({stateMap: {}}); // 审批人抄送人数组
const apprCcList = ref([]);
const approveInfoList: any = ref([]);
const nextApprIndex = ref('');
const apprInfo = ref({}); // 审批详情信息
const loading = ref<boolean>(false); //load加载
const travelInfo = ref({}); // 差旅的信息详情
const trainInfo = ref([]); // 差旅申请的火车出行详情
const planeInfo = ref([]); // 差旅申请的飞机出行申请
const hotelInfo = ref([]); // 差旅申请的酒店出行申请
const orderUserList = ref([]); // 用于存放行程明细
const orderIndex = ref(""); // 用于标识当前显示的人员相关的行程明细
const noKey = ref(false); // 如果没有相关的行程信息则显示为暂无行程信息
const payload: any = {
  "approveNo": '',
  "corpCode": userInfoData.userInfo.corpCode,
  "empCode": userInfoData.userInfo.empCode,
  "equipmentType": ""
}
// 根据获取来的差旅申请单号内容来获取对应的用户信息
const getPeopleOrderListA = async (data: any) => {
  payload.approveNo = data.approveNo;
  let result: any = await getPeopleOrderList(payload);
  orderUserList.value = result.data.data.list;
  if(result.data.data.list.length==0){
    noKey.value=true;
  }
};
// 根据选中的人员显示对应的明细列表
const orderOver = (item: any) => {
  orderIndex.value = item.togetherPersonId;
  payload.corpCode = item.corpCode;
  payload.empCode = item.togetherPersonId
}
// 撤销申请单
const revokeDoc = async (data: any) => {
  const payloadA: any = {
    'applyNo': data.approveNo,
    'corpCode': userInfoData.userInfo.corpCode,
  }
  showConfirmDialog({
    title: '撤销申请单',
    message:
        '此操作将撤回该条申请, 是否继续',
    confirmButtonColor: '#4BA3FB'
  }).then(async () => {
    loading.value = true;
    let result: any = await approvalRevoke(payloadA);
    if (result.success) {
      showSuccessToast('撤销成功');
      await getLoanApply(apprInfo.value);
      await getApprovalHistory(apprInfo.value);
    } else {
      showFailToast(result.msg);
    }
  })
}
// 删除申请单
const deleteDoc = async (data: any) => {
  const payload: any = {
    'approveNo': data.approveNo,
    'corpCode': userInfoData.userInfo.corpCode,
  }
  showConfirmDialog({
    title: '删除申请单',
    message:
        '是否确认删除这条免审申请单，确认后将删除',
    confirmButtonColor: '#4BA3FB'
  }).then(async () => {
    let result: any = await deleteWorkInfo(payload);
    if (result.success) {
      showSuccessToast('删除成功');
      await history.back();
    }
  })
}
const onClickLeft = () => history.back();
//获取申请单详情
const getLoanApply = async (data: any) => {
  const payload: any = {
    "approveNo": data.approveNo,
    "approveType": data.approveType,
    "corpCode": userInfoData.userInfo.corpCode,
    "empCode": userInfoData.userInfo.empCode,
    "equipmentType": ""
  }
  let result: any = await getLoanApplyDetail(payload);
  docDetail.value = result.data
  const rowJson: string = docDetail.value.rowShowJson;
  if (rowJson != '') {
    rowShowJson.value = JSON.parse(rowJson);
  }
  if (result.data.approveType == 'travel') {
    travelInfo.value = result.data.travelApplyDetail;
    travelInfoFormat(result.data.travelApplyDetail.epidemicSituationCity);
    getPeopleOrderListA(apprInfo.value);
  }
}
// 处理获取来的行程信息用于回显
const travelInfoFormat = (list: any) => {
  for (let item in list) {
    if (list[item].travelBy == 'air') {
      planeInfo.value.push(list[item] as never)
    } else if (list[item].travelBy == 'train') {
      trainInfo.value.push(list[item] as never)
    } else {
      hotelInfo.value.push(list[item] as never);
    }
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
    approveInfoList.value.push(apprData);
  }
  loading.value = false;
}
// 根据状态显示审批状态
const stateLoad = (key: String, state: String) => {
  if (state == 'ok' && key == 'over') {
    return '已通过'
  } else if (state == 'ng' && key == 'over') {
    return '已拒绝'
  } else if (docDetail.value.approvalState == '已撤销') {
    return '已撤销'
  } else {
    return '待审批'
  }

}
// 去打车
const goTakeTaxi = async (e: any) => {
  let param = {
    empCode: userInfoData.userInfo.empCode,
    corpCode: userInfoData.userInfo.corpCode,
    approveNo: e.approveNo,
    approveType: 'car',
  };
  let res: any = await getCarLoginUrl(param)
  if (res.success) {
    router.push({
      name: 'IframePage',
      params: {
        barTitle: '用车',
        iframeUrl: res.data.url
      }
    })
  }
}
onBeforeMount(async () => {
  let applicationDetail: any = sessionStorage.getItem('applicationDetail');
  loading.value = true;
  apprInfo.value = JSON.parse(applicationDetail);
  await getLoanApply(apprInfo.value);
  await getApprovalHistory(apprInfo.value);
})
onUnmounted(()=>{
  sessionStorage.removeItem('applicationDetail');
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

.strck {
  position: sticky;
  bottom: 0px;
}

.travelCard {
  background: #FCFCFC;
  padding: 10px;
  margin: 10px;
  border: 1px #999999;
  border-radius: 8px;
}

.approveCard {
  margin: 5px;
  padding-top: 10px;
  padding-bottom: 5px;
  background: white;
  border-radius: 5px;
}

.checkCard {
  color: #0088ff;
  border-bottom: 2px solid #0088ff;
}

.noCard {
  color: #333333;
}

.trainCard {
  margin: 10px;
  /*padding: 8px;*/
  background: #F6F8FA;
  border-radius: 8px;
}

.trainCardHeader {
  background: #0088ff;
  padding: 5px;
  color: #ffffff;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
}

.hotelCardHeader {
  background: #F5A623;
  padding: 5px;
  color: #ffffff;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
}

.planeCardHeader {
  background: #FA6131;
  padding: 5px;
  color: #ffffff;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
}

.ok {
  color: #0088ff;
}

.ng {
  color: #ff6257;
}

.stateTip {
  color: orange;
  background: rgba(242, 152, 66, 0.5);
  padding: 6px;
  margin: 10px;
  font-size: 12px;
  border-radius: 5px;
}
</style>
