<template>
<!--  <van-sticky>-->
<!--    <van-nav-bar-->
<!--        title="申请单"-->
<!--        left-arrow-->
<!--        @click-left="onClickLeft"-->
<!--    />-->
<!--  </van-sticky>-->
  <NavBar title="申请单" />
  <div class="approvalCard">
    <span class="mb5 f14">办公申请</span>
    <div class="ml10 mt10 ">
      <van-row justify="space-around" align="center" >
      <van-col span="6" class="tc" @click.stop="goLeaveApply">
        <van-image :src="leaveIcon"  width="50px" height="50px" position="center" />
        <div class="f12">请假申请</div>
      </van-col>
      <van-col  span="6"  class="tc">
        <van-image :src="workIcon"  @click.stop="goOverWork" width="50px" height="50px" position="center" />
        <div class="f12">加班申请</div>
      </van-col>
      <van-col  span="6"  class="tc">
        <van-image :src="userSealIcon" @click.stop="goUseSeal"  width="50px" height="50px" position="center" />
        <div class="f12">用章申请</div>
      </van-col>
      <van-col  span="6"  class="tc">
        <van-image :src="cardIcon"  width="50px" @click.stop="goOpenCard" height="50px" position="center" />
        <div class="f12">考勤申请</div>
      </van-col>
    </van-row>
    </div>
    <div class="ml10 mt10">
      <van-row justify="space-around" align="center" >
      <van-col span="6" class="tc" >
        <van-image :src="itemCollectionIcon"  @click.stop="goRequisition" width="50px" height="50px" position="center" />
        <div class="f12">物品领用</div>
      </van-col>
      <van-col  span="6"  class="tc">
        <van-image :src="outWorkIcon"  @click.stop="goingOutWork" width="50px" height="50px" position="center" />
        <div class="f12">外出办公</div>
      </van-col>
      <van-col  span="6"  class="tc">
        <van-image :src="DSQpIcon" @click.stop="goDirector" width="50px" height="50px" position="center" />
        <div class="f12">董事签批</div>
      </van-col>
      <van-col  span="6"  class="tc">
        <van-image :src="entertainIcon"  @click.stop="goEntertain" width="50px" height="50px" position="center" />
        <div class="f12">客户宴请</div>
      </van-col>
    </van-row>
    </div>
    <div class="m10">
      <van-row justify="start" align="center" class="mt20">
      <van-col span="6" class="tc">
        <van-image :src="meetingIcon" @click.stop="goMeeting"  width="50px" height="50px" position="center" />
        <div class="f12">会议记录</div>
      </van-col>
      <van-col  span="6"  class="tc">
        <van-image :src="subscriptionIcon" @click.stop="goSubscribe" width="50px" height="50px" position="center" />
        <div class="f12">物品申购</div>
      </van-col>
    </van-row>
    </div>
  </div>
  <div class="approvalCard">
    <span class="mb5 f14">业务申请</span>
    <div class="ml10 mt10 mr10">
      <van-row justify="space-around" align="center">
      <van-col span="6" class="tc">
        <van-image :src="travelBxIcon" @click.stop="goTravel"  width="50px" height="50px" position="center" />
        <div class="f12">差旅申请</div>
      </van-col>
      <van-col  span="6"  class="tc">
        <van-image :src="userCarIcon" @click.stop="goUseCar" width="50px" height="50px" position="center" />
        <div class="f12">用车申请</div>
      </van-col>
      <van-col  span="6"  class="tc">
        <van-image :src="logIcon" @click.stop="goLogistics"  width="50px" height="50px" position="center" />
        <div class="f12">快递物流</div>
      </van-col>
      <van-col  span="6"  class="tc">
        <van-image :src="workIcon" @click.stop="goCost"  width="50px" height="50px" position="center" />
        <div class="f12">费用申请</div>
      </van-col>
    </van-row>
    </div>

  </div>
</template>

<script lang="ts" setup>
import NavBar from '../../components/navBar/index.vue'
import  travelBxIcon from '../../assets/icon/application_icon_13.png';
import  leaveIcon from '../../assets/icon/application_icon_01.png';
import  workIcon from '../../assets/icon/application_icon_02.png';
import  userSealIcon from '../../assets/icon/application_icon_07.png';
import  userCarIcon from '../../assets/icon/application_icon_15.png';
import  subscriptionIcon from '../../assets/icon/application_icon_06.png';
import  DSQpIcon from '../../assets/icon/application_icon_11.png';
import  cardIcon from '../../assets/icon/application_icon_04.png';
import  itemCollectionIcon from '../../assets/icon/application_icon_05.png';
import  entertainIcon from '../../assets/icon/application_icon_08.png';
import  meetingIcon from '../../assets/icon/application_icon_09.png';
import  outWorkIcon from '../../assets/icon/application_icon_03.png';
import  logIcon from '../../assets/icon/application_icon_05.png';
import {useRouter} from "vue-router";

const router = useRouter();
// 退出
const onClickLeft = () =>  history.back();
// 跳到请假申请
const goLeaveApply=()=> router.push('/leaveApply')
// 跳到加班申请
const goOverWork=()=> router.push('/overWork')
// 跳到用章申请
const goUseSeal=()=> router.push('/useSeal')
// 跳到打卡申请
const goOpenCard=()=> router.push('/openCard')
// 跳转到物品领用
const goRequisition =()=> router.push('/requisition')
// 跳转到外出办公
const goingOutWork =()=> router.push('/goingOutWork')
// 跳转到董事签批
const goDirector =()=> router.push('/director')
// 跳转到客户宴请
const goEntertain =()=> router.push('/entertain')
// 跳转到会议记录
const goMeeting =()=> router.push('/meeting')
// 跳转到申购
const goSubscribe =()=> router.push('/subscribe')
// 跳转到用车
const goUseCar =()=> router.push('/useCar')
// 跳转到物流
const goLogistics =()=> router.push('/logistics')
// 跳转到费用
const goCost =()=> router.push('/cost')
// 跳转到差旅申请
const goTravel =() => router.push('/travel')
</script>

<style scoped>
.approvalCard{
  margin: 10px;
  padding: 10px;
  background: #FFFFFF;
  border-radius: 8px;
}
</style>
