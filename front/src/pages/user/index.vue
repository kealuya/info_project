<template>
    <div class="headBg">
    <van-sticky>
      <div style="padding-top: 50px;padding-left: 10px">
        <van-row>
          <van-col :span="7" style="padding-top:10px;padding-left: 20px">
            <div style="background: #ffffff; width:70px; height:70px;border-radius: 10px">
              <van-image :src="logoIcon" width="70px" height="70px"  />
            </div>
          </van-col>
          <van-col :span="16">
            <van-row>
              <van-col :span="8">
                <span class="f20 f_white">{{userInfo.userInfo.userName}}</span>
              </van-col>
              <van-col :span="14" class="tr">
                <span class="f14 f_white">{{userInfo.userInfo.jobName}}</span>
              </van-col>
            </van-row>
            <div class="mt5 mb5">
              <van-row>
                <span class="f_white f14">{{userInfo.userInfo.deptName }}</span>
              </van-row>
              <van-row>
                <span  class="f_white f14">所属: {{userInfo.userInfo.corpName}}</span>
              </van-row>
            </div>
            <div>
              <van-row>
                <van-col :span="16">
                  <div  class="f_white f14">Tel: {{userInfo.userInfo.phone}}</div>
                  <div  class="f_white f14">No: {{userInfo.userInfo.empCode}}</div>
                </van-col>
                <van-col :span="8">
                  <van-image :src="userInfoIcon"  width="50px" height="50px" />
                </van-col>
              </van-row>
            </div>
          </van-col>
        </van-row>
      </div>
    </van-sticky>
    </div>

  <div class="dataTip">
    <van-row justify="center" align="center">
      <van-col span="8" @click="gotoPath('trip')">
        <div>{{tripNumber}}次</div>
        <div>累计出行</div>
      </van-col>
      <van-col span="8" @click="gotoPath('app')">
        <div>{{appNumber}}单</div>
        <div>累计申请</div>
      </van-col>
      <van-col span="8" @click="gotoPath('bxMoney')">
        <div>{{parseInt(bxMoneyAmount)}}元</div>
        <div>累计报销</div>
      </van-col>
    </van-row>
  </div>
    <div>
        <van-cell title="修改个人信息" is-link to="personalInfo">
            <template #icon>
                <van-image :src="meIcon1" class="cellImage"/>
            </template>
        </van-cell>
        <van-cell title="个人考勤统计" is-link to="personalAttendance">
            <template #icon>
                <van-image :src="meIcon10" class="cellImage"/>
            </template>
        </van-cell>
        <van-cell title="企业数据统计" is-link to="enterpriseStatistics">
            <template #icon>
                <van-image :src="meIcon3" class="cellImage"/>
            </template>
        </van-cell>
        <van-cell title="管理账户" is-link :to="{name:'manageAccounts',params:{bankType:'个人账户'}}">
            <template #icon>
                <van-image :src="meIcon2" class="cellImage"/>
            </template>
        </van-cell>
        <van-cell title="我的委托" is-link to="myDelegation">
            <template #icon>
                <van-image :src="meIcon7" class="cellImage"/>
            </template>
        </van-cell>
        <van-cell title="帮助反馈" is-link to="helpFeedback">
            <template #icon>
                <van-image :src="meIcon5" class="cellImage"/>
            </template>
        </van-cell>
        <van-cell title="退出登录" is-link @click="escUser">
            <template #icon>
                <van-image :src="meIcon6" class="cellImage"/>
            </template>
        </van-cell>
    </div>
</template>

<script setup lang="ts">
//  logo头像
import logoIcon from '../../assets/icon/szht_logo.png';
import meIcon1 from '../../assets/icon/me_icon_01.png';
import meIcon2 from '../../assets/icon/me_icon_02.png';
import meIcon3 from '../../assets/icon/me_icon_03.png';
import meIcon4 from '../../assets/icon/me_icon_04.png';
import meIcon5 from '../../assets/icon/me_icon_05.png';
import meIcon6 from '../../assets/icon/me_icon_06.png';
import meIcon7 from '../../assets/icon/me_icon_07.png';
import meIcon9 from '../../assets/icon/me_icon_09.png';
import meIcon10 from '../../assets/icon/me_icon_10.png';
import meIcon11 from '../../assets/icon/me_icon_11.png';
import moment from "moment";
import {useRouter} from "vue-router";
// 用户信息图标
import userInfoIcon from '../../assets/icon/userInfo.png'
import {inject, onBeforeMount,reactive,ref} from "vue";
import {showConfirmDialog} from "vant";
import {countAmountNumber, modifyUserInfo} from "../../services/user";
let userInfo=reactive<any>({});
let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
let nowDate=new Date(); // 获取当前时间
let tripNumber=ref('0'); // 出行次数
let appNumber=ref('0'); // 累计申请
let bxMoneyAmount=ref('0'); // 累计申请
let travelMoney=ref('0'); // 累计出行金额
const router = useRouter(); //使用路由跳转
// 退出登陆
const escUser=()=>{
  showConfirmDialog({
    title: '退出',
    message:
        '您确定要退出吗？',
  })
      .then(() => {
        localStorage.clear()
        sessionStorage.clear();
        location.reload()
        // on confirm
      })
      .catch(() => {
        // on cancel
      });
}
// 调用获取用户个人图表信息 根据时间范围统计个人出行次数、累计报销金额、累计申请次数
const getAmountCheck=async()=>{
  let minDate = new Date(nowDate.getFullYear()-1, nowDate.getMonth(), nowDate.getUTCDate()); //最小时间是当前日期的前半年
  let da={
  'corpCode':userInfoData.userInfo.corpCode,
  'empCode':userInfoData.userInfo.empCode,
  'userName':userInfoData.userInfo.userName,
  'startTime':moment(minDate).format('yyyy-MM-DD'),
  'endTime':moment(nowDate).format('yyyy-MM-DD'),
  }
  let result: any = await countAmountNumber(da);
  if(result.success==true){
    tripNumber.value=result.data.tripNumber;
    appNumber.value=result.data.appNumber;
    bxMoneyAmount.value=result.data.bxMoneyAmount;
    travelMoney.value=result.data.travelMoney;
  }else{

  }
}
// 根据code跳转下一页时返回对应效果
const gotoPath=(code:any)=>{
  let pathInfo={
    code:code,
    tripNumber:tripNumber.value,
    appNumber:appNumber.value,
    bxMoneyAmount:bxMoneyAmount.value,
    travelMoney:travelMoney.value,
  };
  sessionStorage.setItem('personUserChat', JSON.stringify(pathInfo))
  router.push('personalUsageStatistics');
}
onBeforeMount(async ()=>{
  let userInfoJson:any=sessionStorage.getItem('userInfo');
  userInfo=JSON.parse(userInfoJson);
  getAmountCheck();
})
</script>

<style lang="less" scoped>
.headBg {
  width: 100%;
  height: 200px;
  background: url('../../assets/img/bg_01.png') no-repeat center;
  background-size: 100vw 200px;
}
.dataTip{
  margin: 8px;
  padding: 15px 10px;
  text-align: center;
  color:#ffffff;
  background:  url('../../assets/img/bg.png') no-repeat center;
  background-size: 100% 80px;
  border-radius: 10px;
}
.cellImage {
    width: 20px;
    height: 20px;
    margin-right: 5px;
}
</style>
