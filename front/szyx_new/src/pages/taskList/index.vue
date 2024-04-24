<template>
  <div class="listBox">
    <van-nav-bar
        title="业务内容"
        left-text="返回"
        left-arrow
        @click-left="onClickLeft"
    />
    <van-tabs v-model:active="active" @change="changeList">
     <div :class="{ 'list1': tabIndex }">
       <div class="list">
         <van-tab title="会议文件">
           <van-list
               :offset="80"
               v-model="loading"
               :finished="finished"
               :immediate-check="false"
               finished-text="没有更多了"
               loading-text="正在加载中请稍后"
               style="height: 100%"
               @load="getMeetingListLoad"
           >
             <div class="list_card" v-for="item in meetingList" :key="item.id">
               <div class="flex-space m-b-10">
                 <div style="display: flex;align-items: center">
                   <van-tag type="primary" v-if="item.meetingType == 'audio'">音频会议</van-tag>
                   <van-tag type="success" v-else>文档记录</van-tag>
                   <div class="list_title">{{item.meetingTitle}}-{{item.isCheck}}</div>
                 </div>
                 <van-checkbox v-model="item.isCheck" @change="changeCheck"></van-checkbox>
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
             <!--          <businessCard :meetingList="meetingList" @changeCheck="changeCheck"></businessCard>-->
           </van-list>
         </van-tab>
       </div>
       <div class="footer_btn">
         <van-button block type="primary" @click="handleSubmit">
           提交
         </van-button>
       </div>
     </div>
      <div class="listUse" :class="{ 'list1': tabIndex==false}">
        <van-tab title="已使用">
          <van-list
              :offset="10"
              v-model="loading"
              :finished="finished"
              :immediate-check="false"
              finished-text="没有更多了"
              loading-text="正在加载中请稍后"
              style="height: 100%"
              @load="getMeetingListLoad"
          >
            <div class="list_card" v-for="item in meetingList" :key="item.id" :class="{ 'list1': !tabIndex }">
              <div class="flex-space m-b-10">
                <div style="display: flex;align-items: center">
                  <van-tag type="primary" v-if="item.meetingType == 'audio'">音频会议</van-tag>
                  <van-tag type="success" v-else>文档记录</van-tag>
                  <div class="list_title">{{item.meetingTitle}}-{{item.isCheck}}</div>
                </div>
<!--                <van-checkbox v-model="item.isCheck" @change="changeCheck"></van-checkbox>-->
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
            <!--          <businessCard :meetingList="meetingList" @changeCheck="changeCheck"></businessCard>-->
          </van-list>
        </van-tab>
      </div>
    </van-tabs>
  </div>
</template>
<script setup lang="ts">
import hyTitle from '../../assets/img/hy_title.png'
import businessCard from '../../components/businessCard/index.vue'
import word from '../../assets/img/word.png'
import xmind from '../../assets/img/xmind.png'
import mp3 from '../../assets/img/mp3.png'

import businessIcon from '../../assets/img/business_icon.png';
import {ref, computed, onMounted, inject} from 'vue';
// import userRouter from 'user'
import { useRouter } from "vue-router";
import {getMeetingList} from "../../services/task-processing";
import {meetingData} from "../../store/index"
import {paramsTye} from "../../services/task-processing/types";
const checked = ref(false);
const active = ref(0);
const router = useRouter()
const meetingStore = meetingData()
const meetingList = ref<any>([])
const checkedItemsList = ref<any>()   //会议列表 勾选的数组
const loading = ref<Boolean>(false)
const finished = ref<Boolean>(false)
const tabIndex = ref(false)
interface pageType{
  currentPage:number, //	当前页
  pageSize:	number,//	Integer	是	每页显示条数
  // searchKey:string,//	否	筛选条件，暂时未用
  corpCode:string | null//	是	企业code
  userId:	string	//是	用户id
  meetingFlag:string
}
let params = ref<pageType>(
    {
      currentPage: 1,
      pageSize: 10,
      corpCode: '',
      userId: '',
      meetingFlag:'0'
    }
)
const onClickLeft = () => {
  router.replace('/addTasks')
}

const changeCheck = (checkedItems:any)=>{
  checkedItemsList.value = meetingList.value.filter((item:any) => item.isCheck);
}
// tab 切换
const changeList = async (name: number) => {
  if (name == 1) {
    params.value.currentPage = 1
    meetingList.value = []
    loading.value = true
    finished.value = false
    // isLoading.value = true
    await getList()
    tabIndex.value = true
  } else {
    params.value.currentPage = 1
    meetingList.value = []
    loading.value = true
    finished.value = false
    await getList()
    tabIndex.value = false
  }
}
//点击了提交按钮
const handleSubmit = async ()=>{
 await meetingStore.setMeetingData(checkedItemsList.value)
  router.replace('/addTasks')
}
//触底函数
const totalCount = ref<number>(0)
const pageCount = ref<number>(0)
const getMeetingListLoad = async()=>{
  if (meetingList.value.length < totalCount.value && params.value.currentPage < pageCount.value) {
    params.value.currentPage = params.value.currentPage + 1
    await getList()
  }
 // params.value.currentPage =  params.value.currentPage + 1
 //  await getList()
}
const getList = ()=>{
  params.value.meetingFlag = String(active.value)
  getMeetingList(params.value).then((res:any)=>{
    meetingList.value = meetingList.value.concat(res.data.MeetingList)
    totalCount.value = res.data.totalCount
    pageCount.value = res.data.pageCount
    meetingList.value.forEach((item:any)=>{
      item.isCheck = false
    })
  })
}
onMounted(async()=>{
  let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
  params.value.corpCode = localStorage.getItem('corpCode') //从本地获取corpCode
  //从本地获取corpCode
  params.value.userId = userInfoData.userInfo.userId
  meetingList.value.length = 0
  params.value.currentPage = 1
  await getList()
  // getMeetingList(params.value).then((res:any)=>{
  //   meetingList.value = res.data.MeetingList
  //   meetingList.value.forEach((item:any)=>{
  //     item.isCheck = false
  //   })
  // })
})
</script>
<style lang="less" scoped>
.listBox{
  height:100vh;
}
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
  width:96vw;
  padding:0 2vw;
}
.box{
  height:50px;
}
.list{
  height: calc(100vh - var(--van-nav-bar-height) - var(--van-tabs-line-height) - var(--van-tabbar-height) - var(--van-button-default-height));
  overflow-y: auto;
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
    .f-z-12{
      font-size: 0.75em;
      color:#4B5563;
    }
    .m-b-10{
      margin-bottom: 2vw;
    }
  }
}
.listUse{
  height: calc(100vh - var(--van-nav-bar-height) - var(--van-tabs-line-height) - var(--van-tabbar-height));
  overflow-y: auto;
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
    .f-z-12{
      font-size: 0.75em;
      color:#4B5563;
    }
    .m-b-10{
      margin-bottom: 2vw;
    }
  }
}
.list1{
  display: none;
}
.list_title1{
  font-size: 0.75em;
  //font-weight: 550;
  margin-left: 2vw;
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
