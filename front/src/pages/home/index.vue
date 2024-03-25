<template>
  <van-sticky>
    <van-row>
      <div :class="weatherKey(weatherInfo.weather)?weachetCard:'default'" class=" headBg ">
        <weather-component
            :weather-info="weatherInfo"
            :week-time="weekTime"
            :now-time="nowTime"
            :account-book="accountBook"
            :msg="msg"
        />
      </div>
    </van-row>
  </van-sticky>
  <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
    <template #success>
      <span v-if="refreshingOnly==true"><van-icon name="success" /> 刷新成功</span>
    </template>
    <IconComponent/>
    <van-row  justify="space-around" :getter="10" align="center" class="mt10">
      <van-col span="11" @click.stop="openApplication">
        <div class="mainButton f_white">
         <div>
          <van-row justify="space-between" align="center">
            <van-col>
              <p class="f16">申请单</p>
              <p class="f12">快速申请 规范流程</p>
            </van-col>
            <van-col>
              <van-icon :name="shenqingdanIcon" size="3em"   />
            </van-col>
          </van-row>
         </div>
        </div>
      </van-col>
      <van-col span="11">
        <div class="mainButton_sec  f_white" @click="openReimburse">
          <div>
           <van-row justify="space-between" align="center">
             <van-col>
               <p class="f16">报销单</p>
               <p class="f12">移动报销 随时随地</p>
             </van-col>
             <van-col>
               <van-icon :name="danjuIcon" size="3em"  />
             </van-col>
           </van-row>
          </div>
        </div>
      </van-col>
    </van-row>
    <div class="mt5">
      <van-row class="p5" justify="space-between" align="center">
        <van-col span="8" class="tl ml5 f18">单据动态</van-col>
        <van-col span="8" class="tr f14 mr10" @click="about()">查看更多 ></van-col>
      </van-row>
      <div>
      </div>
    </div>
    <div v-if="loading==false && list.length==0" style="height: 500px">
      <van-row justify="center" align="center" class="tc">
        <div>
          <img style="height:50vw;width:80vw;object-fit: contain" src="../../assets/img/zanwupiaoju.png" v-cloak/>
        </div>
      </van-row>
      <div style="color: rgba(23,26,29,0.40);text-align: center">暂无单据动态</div>
    </div>
    <van-list
        v-model="loading"
        v-else
        :finished="finished"
        :immediate-check="false"
        loading-text="正在加载中请稍后"
        finished-text="没有更多了"
        :offset="10"
    >
      <div class="ticket_content" v-for="(item,index) in list" :key="item">
            <application-card  :get-data="item" />
      </div>
    </van-list>
    <div style="margin-bottom: 60px"></div>
  </van-pull-refresh>
</template>

<script setup lang="ts">
import ApplicationCard from "./component/applicationCard.vue";
import IconComponent from "./component/iconComponent.vue";
import {fetchLocation, fetchWeather} from "../../services/utils";
import {inject, onMounted, ref} from "vue";
import {getAccountBookList, getDocStateList, } from "../../services/home";
import {weatherTypeList } from './component/ts/weatherType';
import  {getApprCountTeacher} from '../../services/utils'
import {useRouter} from "vue-router";
import WeatherComponent from "./component/weatherComponent.vue";
import shenqingdanIcon from  '../../assets/img/shenqingdan.png';
import danjuIcon from  '../../assets/img/danju.png';
const locationName = ref("请选择地理位置");
class classDemo {
    approveType: string | undefined
}
let list = ref<any>([]) as classDemo[];//加载loading
const router = useRouter();
const loading = ref(true);
const finished = ref(false);
const refreshing = ref(false);
let refreshingOnly = ref(false); //加载效果只有下拉变化
let userInfoData: any = inject("userInfo");
let msg=ref('0');
let accountBook=ref('0');
let weachetCard=<any>ref('');
// 根据天气的key
const weatherKey=async (key:any)=>{
  let weather:any=await weatherTypeList.find((e)=>{
      return e.name==key;
  })
  if(weather){
    weachetCard.value=weather.style
  }
};
const refresh=async()=>{
  if (refreshing.value) {
  }
  if(finished.value==true){
    return
  }
  await getDocList(payload,'refresh');
  getWaitApprovalMsg();
  getAccountBook();
}
const about=()=>{
  router.push('/documentDynamics')
}
const onRefresh = () => {
  // 清空列表数据
  finished.value = false;
  // 重新加载数据
  // 将 loading 设置为 true，表示处于加载状态
  loading.value = true;
  getDocList(payload,'loading')

};
let payload = {
  // 'companyCode': userInfoData.userInfo.corpCode,
  'corpCode': userInfoData.userInfo.corpCode,
  'empCode': userInfoData.userInfo.empCode,
  'page': 1,
  'pageSize': 5,
  'approveType': [],
  'createTime': '',
  'keyWord': '',
  'approveReason': '',
  "approvalState": "",
}
// 获取审批和账本数
const getWaitApprovalMsg=async()=>{
  let resData={
    'type':'fk',
    'userid': userInfoData.userInfo.empCode,
    'corpCode': userInfoData.userInfo.corpCode,
    'school':userInfoData.userInfo.corpCode,
    'isManage':'',
    'deptId':userInfoData.userInfo.deptId,
    'reimbursedState':'0',
    'page':1,
    'pageSize':10
  };
  let result: any = await getApprCountTeacher(resData);
  msg.value=result.data;
}
const getAccountBook=async()=>{
  let resData={
    "page":1,
    "pageSize":8,
    "keyWord":"",
    "empCode": userInfoData.userInfo.empCode,
    "expenseTypeName":"",
    "occurrenceTime":"",
    "state":"0",
    "corpCode": userInfoData.userInfo.corpCode,
    "equipmentType":""
  };
  let result: any = await getAccountBookList(resData);
  accountBook.value=result.data.total
}
// 单据动态
const getDocList=async(payload:any,code:any)=>{
  let result: any = await getDocStateList(payload);
  if(code=='refresh'){
    list.value = [];
  }
  list.value.push(...result.data.list);
  if (list.value.length.toString() === result.data.total){
    loading.value = false;
    finished.value = true;
  } else {
    let page = Number(payload.page) + 1;
    payload.page = page.toString();
    loading.value = false;
  }
  refreshing.value = false;
}
// 跳转到申请单
const openApplication = () =>  { router.push('/application')};

// 跳转到报销单
const openReimburse = () => {

  router.push('/reimburse');
};
//
const complement = function (value: any) {
  return value < 10 ? `0${value}` : value;
};
const nowTime = ref("");
const weekTime = ref("")
// 转换日期
const formateDate = (date: any) => {
  const time = new Date(date);
  const month = complement(time.getMonth() + 1);
  const day = complement(time.getDate());
  const week = '星期' + '日一二三四五六'.charAt(time.getDay());
  weekTime.value = week;
  return `${month}-${day}`;
};
const weatherInfo:any = ref({
  weather:''
});
onMounted(() => {
  getLocation();
  setInterval(() => {
    nowTime.value = formateDate(new Date())
  });
  refresh();
  getWaitApprovalMsg();
  getAccountBook();

});
// 获取本地天气
const getLocation = async () => {
  let result: any = await fetchLocation();
  if (result.adcode) {
    let weatherResult: any = await fetchWeather(result.adcode);
    weatherInfo.value = weatherResult.lives[0]
  }
  locationName.value = result;
};

</script>

<style lang="less" scoped>
.headBg {
  width: 100vw;
  height: 22vh;
}
.sun{
  background: url('../../assets/img/sun.png') no-repeat top left;
  background-size: 100% 30vh;
}
.night{
  background: url('../../assets/img/night.png') no-repeat top left;
  background-size: 100% 30vh;
}
.fog{
  background: url('../../assets/img/fogDay.png') no-repeat top left;
  background-size: 100% 30vh;
}
.cloudy{
  background: url('../../assets/img/cloudyDay.png') no-repeat top left;
  background-size: 100% 30vh;
}
.rainy{
  background: url('../../assets/img/rainyDay.png') no-repeat top left;
  background-size: 100% 30vh;
}
.snow{
  background: url('../../assets/img/snowDay.png') no-repeat top left;
  background-size: 100% 30vh;
}
.wind{
  background: url('../../assets/img/wind.png') no-repeat top left;
  background-size: 100% 30vh;
}
.sleet{
  background: url('../../assets/img/sleet.png') no-repeat top left;
  background-size: 100% 30vh;
}
.default{
  background: url('../../assets/img/bg_01.png') no-repeat top left;
  background-size: 100% 30vh;
}
.mainButton {
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 50px;
  background: url('../../assets/img/shenqigndan_bg.png') no-repeat center;
  background-size: 100vw 10vh;
  //background-image: linear-gradient(to right bottom, #0080FF, #6AB5FF);
}

.mainButton_sec {
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 50px;
  background: url('../../assets/img/danju_bg.png') no-repeat center;
  background-size: 100vw 10vh;
  //background-image: linear-gradient(to right bottom, #33B2AD, #3BD7D3);
}

</style>
