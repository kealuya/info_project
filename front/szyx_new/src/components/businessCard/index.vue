<template>
  <div class="list">
<!--    <van-list-->
<!--        :offset="30"-->
<!--        v-model="loading"-->
<!--        :finished="finished"-->
<!--        :immediate-check="false"-->
<!--        finished-text="没有更多了"-->
<!--        loading-text="正在加载中请稍后"-->
<!--        style="height: 100%"-->
<!--        @load="getListLoad"-->
<!--    >-->
    <div class="list_card" v-for="item in meetingList" :key="item.id">
      <div class="flex-space m-b-10">
        <div style="display: flex;align-items: center">
          <van-tag type="primary" v-if="item.meetingType == 'audio'">音频会议</van-tag>
          <van-tag type="success" v-else>文档记录</van-tag>
          <div class="list_title">{{item.meetingTitle}}-{{item.isCheck}}</div>
        </div>
        <van-checkbox v-model="item.isCheck" v-if="!isCollapse" @change="changeCheck"></van-checkbox>
        <!--            <van-image-->
        <!--                width="20"-->
        <!--                height="20"-->
        <!--                :src="businessIcon"-->
        <!--            />-->
      </div>
      <van-divider />
      <!--          //display:flex;align-items: center;-->
      <div style="margin-bottom: 10px; margin-left: 10px;display:flex;align-items: center;" v-for="i in item.meetingFile">
                   <van-image :src="word" width="23" height="23" v-if="i.fileType=='word'"></van-image>
                   <van-image :src="xmind" width="23" height="23" v-if="i.fileType=='xmind'"></van-image>
                   <van-image :src="mp3" width="23" height="23" v-if="i.fileType=='mp3'"></van-image>
        <div class="list_title1">{{i.fileName}}</div>
<!--        <div class="list_title1">{{item.meetingBrainMapName?item.meetingBrainMapName:''}}</div>-->
<!--        <div class="list_title1">{{item.meetingMminutesName?item.meetingMminutesName:''}}</div>-->
      </div>
      <div>
        <!--            <div class="f-z-12 m-b-10" >所属会议：{{item.sshy}}</div>-->
        <!--<div class="f-z-12 m-b-10" v-if="item.hyType!=='word'">会议地点：{{item.address}}</div>-->
        <div class="f-z-12 m-b-10" style="margin-left: 10px;">创建时间：{{item.createTime}}</div>

      </div>
    </div>
<!--    </van-list>-->
  </div>
</template>
<script setup lang="ts">
import hyTitle from '../../assets/img/hy_title.png'
import word from '../../assets/img/word.png'
import xmind from '../../assets/img/xmind.png'
import mp3 from '../../assets/img/mp3.png'

import businessIcon from '../../assets/img/business_icon.png';
import {ref, computed, onMounted} from 'vue';
// import userRouter from 'user'
import { useRouter } from "vue-router";
const checked = ref(false);
const active = ref(0);
const router = useRouter()
const value2 = ref('a');
const timer = ref<string>()
const showDate = ref<boolean>(false)
const loading = ref<Boolean>(false)
const finished = ref<Boolean>(false)
const emit = defineEmits<{ (e: 'changeCheck', checkedItems: any): void }>()
// const emitMeeting = defineEmits<{ (e: 'meeting', meetingData: any): void }>()

const props = defineProps({
  isCollapse: {
    type: Boolean,
    default: false
  },
  meetingList:{
    type:Array,
    default:[]
  }
})
const option2 = [
  { text: '文档', value: 'a' },
  { text: '文档1', value: 'b' },
  { text: '文档2', value: 'c' },
];
//用户勾选了 把值传给父组件
const changeCheck = (names: any[])=>{
  const checkedItems = props.meetingList.filter((item:any) => item.isCheck);
  emit('changeCheck',checkedItems)
  console.log('checkedItems',checkedItems)
}
const getListLoad = ()=>{
  console.log(1111)
  // emitMeeting('meeting',true)
}
const businessDetailHandle = ()=>{
  router.replace('/businessDetail')
}
const onClickLeft = () => {
  router.replace('/addTasks')
}
const handleSubmit = ()=>{
  router.replace('/addTasks')
}
</script>
<style lang="less" scoped>
.header {
  position: fixed;
  z-index: 99;
  width: 100%;
  height: 60px;
  background-color: #fff;
  // padding: 0 10px;
  font-size: 14px;
  display: flex;
  text-align: center;
  height: 50px;
  line-height: 60px;
  border-bottom: 1px solid #f0f0f0;
  //margin-bottom: 50px;
  .menuItems {
    font-size: 14px;
    width: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    // border: 1px solid red;
  }
}
.footer_btn{
  width:92vw;
  padding:2vw;
  margin: 2vw;
  //position: fixed;
  //bottom:60px;
}
.box{
  height:50px;
}
.list{
  //max-height: 75vh;
  //overflow: auto;
  .list_card{
    width:92vw;
    margin:2vw;
    padding:2vw;
    background-color: #FFFFFF;
    border-radius: 10px;
    .flex-space{
      display: flex;
      justify-content: space-between;
    }
    .list_title{
      font-size: 0.875em;
      font-weight: 550;
      margin-left: 2vw;
    }
    .list_title1{
      font-size: 0.75em;
      //font-weight: 550;
      margin-left: 2vw;
    }
    .f-z-12{
      font-size: 0.75em;
      color:#4B5563;
    }
    .m-b-10{
      margin-bottom: 2vw;
    }
  }
}
.mainButton {
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 50px;
  background: url('../../assets/img/radio.png') no-repeat center;
  background-size: 100vw 10vh;
  //background-image: linear-gradient(to right bottom, #0080FF, #6AB5FF);
}

.mainButton_sec {
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 50px;
  background: url('../../assets/img/radio.png') no-repeat center;
  background-size: 100vw 10vh;
  //background-image: linear-gradient(to right bottom, #33B2AD, #3BD7D3);
}
:deep(.van-dropdown-menu__bar){
  box-shadow: none;
}
:deep(.van-divider){
  margin: 2vw;
}
::v-deep(.van-nav-bar__content) {
  background-color: #0088ff;
}

::v-deep(.van-nav-bar__title) {
  color: #ffffff;
}

::v-deep(.van-nav-bar__text) {
  color: #ffffff;
}

::v-deep(.van-icon-arrow-left:before) {
  color: #ffffff;
}
</style>
