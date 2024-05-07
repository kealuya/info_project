<template>
 <div class="container">
   <!-- 上半部分的header -->
   <van-nav-bar
       title="业务列表"
       left-text="返回"
       left-arrow
   />
   <div class="header">
     <van-dropdown-menu class="menuItems">
       <!--    <van-dropdown-item v-model="value1" :options="option1" />-->
       <van-dropdown-item v-model="value2" :options="option2" @change="dropdownChange"/>
     </van-dropdown-menu>
     <div class="menuItems" @click="open">{{ timer ? timer : '时间范围' }}</div>
     <van-calendar v-model:show="showDate" :max-date="maxDate" :min-date="minDate" type="range" @confirm="onConfirm" ref="calendarRef"/>
     <div class="menuItems1" @click="resetHandel">
       重置时间
     </div>
   </div>
   <van-pull-refresh v-model="isLoading" success-text="刷新成功" @refresh="onRefresh" :disabled="isShowImg || pullRefreshDisabled">
     <div v-if="searchLoading && ywList.length==0&&!isPull" class="h-67">
       <van-loading type="spinner" color="#1989fa" :vertical="true">加载中...</van-loading>
     </div>
   <van-list
       @scroll="divScroll"
       v-if="!isShowImg"
       class="list"
       v-model:loading="loading"
       :finished="finished"
       :immediate-check="false"
       :offset="30"
       finished-text="没有更多了"
       loading-text="正在加载中请稍后"
       @load="getListLoad"
   >
       <div v-for="(item,index) in ywList" class="list_card" @click="businessDetailHandle(item.meetingType,item.meetingId)">
         <div class="flex-space m-b-10">
           <div class="list_title">
             <van-tag v-if="item.meetingType=='audio'" type="primary">音频会议</van-tag>
             <van-tag v-else type="success">文档记录</van-tag>
             {{ item.meetingTitle }}{{index+1}}
           </div>
         </div>
         <van-divider/>
         <div class="f-z-12 m-b-10">会议时长：{{ item.time }}</div>
         <div class="f-z-12 m-b-10">会议地址：{{ item.meetingCity }}{{item.meetingAddress}}</div>
         <div class="f-z-12 m-b-10">结束时间：{{ item.createTime }}</div>
<!--         <div class="f-z-12 m-b-10">-->
<!--           <span v-if="item.hyType" style="color: #0080FF">已关联任务：RW202404080001</span>-->
<!--         </div>-->

       </div>
   </van-list>
     <van-row v-else-if="isShowImg" align="center" class="tc" justify="center">
       <div class="list">
         <img src="../../assets/img/zanwupiaoju.png" style="height:50vw;width:80vw;object-fit: contain"/>
         <div class="banner">暂无数据</div>
       </div>
     </van-row>
   </van-pull-refresh>
<!--   <div v-if="searchLoading && ywList.length==0&&loading==false" class="h-67">-->
<!--     <van-loading :vertical="true" color="#1989fa" type="spinner">加载中...</van-loading>-->
<!--   </div>-->
   <van-row :getter="24" align="center" justify="space-around" class="btn-footer">
     <van-col span="11">
       <div class="mainButton f_white" @click="speechHandle">
         <div>
           <van-row align="center" justify="space-between">
             <van-col>
               <p class="f16">音频记录</p>
               <p class="f12">会议录制 准确还原</p>
             </van-col>
           </van-row>
         </div>
       </div>
     </van-col>
     <van-col span="11">
       <div class="mainButton_sec  f_white" @click="documentUpload">
         <div>
           <van-row align="center" justify="space-between">
             <van-col>
               <p class="f16">文档上传</p>
               <p class="f12">集中存储 轻松共享</p>
             </van-col>
           </van-row>
         </div>
       </div>
     </van-col>
   </van-row>
 </div>
  <!--  <div class="mb10"></div>-->
</template>
<script lang="ts" setup>
import businessIcon from '../../assets/img/business_icon.png';
import {ref, computed, onMounted, inject} from 'vue';
// import userRouter from 'user'
import {useRouter} from "vue-router";
import {getMeetingList} from '../../services/task-processing/index'
import empty from '../../assets/img/zanwupiaoju.png'
import {showSuccessToast,showFailToast} from "vant";
import Calendar from "echarts/types/src/coord/calendar/Calendar";
const router = useRouter()
const value2 = ref('meet');
const timer = ref<string>()
const showDate = ref<boolean>(false)
const isLoading = ref<boolean>(false)  //控制 下拉刷新
const loading = ref<boolean>(false)
const finished = ref<boolean>(false)
const searchLoading = ref<Boolean>(false)
const isShowImg = ref<boolean>(false)  //数据为空的时候 控制图片的显示隐藏
let lastRefreshTime = 0;
const refreshInterval = 15000; // 15秒
const isPull = ref<boolean>(false)
const pullRefreshDisabled = ref(false)
const calendarRef = ref<Calendar | null>(); // 创建一个 ref 用于存储 calendar 实例

const divScroll = (e: any) => {
  if (e.target.scrollTop == 0) {
    pullRefreshDisabled.value = false
  } else {
    pullRefreshDisabled.value = true
  }
}
interface  paramsType{
  currentPage: number, //当前页
  pageSize: number, //每页显示条数
  // searchKey: "", //筛选条件，暂时未用
  corpCode: string | null, //企业code
  startTime: string, //初始时间
  endTime: string, //末端时间
  meetingType: string, //会议类型
  userId: string, //人员ID
  meetingFlag: string //会议是否使用 0：未使用   1：已使用
}
let params = ref<paramsType>({
  currentPage: 1, //当前页
  pageSize: 10, //每页显示条数
  // searchKey: "", //筛选条件，暂时未用
  corpCode: '', //企业code
  startTime: '', //初始时间
  endTime: '', //末端时间
  meetingType: '', //会议类型
  userId: '', //人员ID
  meetingFlag: '0' //会议是否使用 0：未使用   1：已使用
})
const  resetHandel = ()=>{
  showDate.value = false;
  params.value.startTime = ''
  params.value.endTime = ''
  getList()
}
const open =()=>{
  showDate.value = true
  calendarRef.value?.reset();

}
//业务数据列表
const ywList = ref([])
const option2 = [
  {text: '全部会议', value: 'meet'},
  {text: '音频记录', value: 'audio'},
  {text: '文档记录', value: 'document'},
];
//切换了会议类型
const dropdownChange = async(value:string)=>{
  console.log('value',value)
  if(value !=='meet'){
    params.value.meetingType = value
    params.value.currentPage = 1
    ywList.value.length = 0
    loading.value = false
    finished.value = false
    await getList()
  }else{
    params.value.meetingType = ''
    params.value.currentPage = 1
    ywList.value.length = 0
    loading.value = false
    finished.value = false
    await getList()
  }
}
//下拉刷新
const onRefresh = async()=>{
  isPull.value = true
  isLoading.value = true   //打开下拉刷新的 loading
  params.value.currentPage = 1 //重置页码
  ywList.value = []//清空表单
  await getList() //重新加载数据
}
// const onRefresh = async () => {
//   const currentTime = Date.now();
//   if (currentTime - lastRefreshTime < refreshInterval) {
//     showFailToast("刷新操作太频繁，请稍后再试！");
//     isLoading.value = false; // 关闭下拉刷新的 loading
//     return; // 直接返回，不执行加载数据的操作
//   }
//   lastRefreshTime = currentTime;
//   isLoading.value = true; // 打开下拉刷新的 loading
//   params.value.currentPage = 1; // 重置页码
//   ywList.value = []; // 清空表单
//   await getList(); // 重新加载数据
//   showSuccessToast("刷新成功！");
//   isLoading.value = false; // 关闭下拉刷新的 loading
// };
const thisYear = new Date().getFullYear();

const minDate = computed(() => {
  return new Date(thisYear, 0); // 今年一月
});

const maxDate = computed(() => {
  return new Date(thisYear + 1, 0); // 明年一月
});
const formatDate = (date: any) => `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}`;  //格式化日期 （年-月-日）
const onConfirm = async (values: any) => {
  const [start, end] = values;
  // console.log(values,'zheli')
  showDate.value = false;
  let beginTime = formatDate(start);
  let endTime = formatDate(end);
  params.value.startTime = beginTime
  params.value.endTime = endTime
  params.value.currentPage = 1
  ywList.value.length = 0
  await getList()

};
const businessDetailHandle = (type: string,meetingId:string) => {
  console.log('type', type)
    router.replace({path:'/businessDetail',query:{meetingId}})
}
const speechHandle = () => {
  router.replace('/speech')

}
const documentUpload = () => {
  router.replace('/documentUpload')
}
const pageCount = ref<number>(0)
const totalCount = ref<number>(0)
const getList =()=>{
  searchLoading.value = true

  // if(loading.value==true){
  //   searchLoading.value = true
  // }
  getMeetingList(params.value).then((res:any)=>{
    console.log(res)
    ywList.value =  ywList.value.concat(res.data.MeetingList)
    isLoading.value =false   //关闭下拉刷新的加载
    totalCount.value = res.data.totalCount
    pageCount.value = res.data.pageCount
    if (res.data.totalCount == 0) {
      isShowImg.value = true
    } else {
      isShowImg.value = false
    }
    if (ywList.value.length == totalCount.value) {
      loading.value = false   //关闭下拉刷新的加载
      finished.value = true
    }else{
      loading.value = false
      finished.value = false
    }
  }).finally(()=>{
    searchLoading.value = false
  })
}
//触底事件
const getListLoad = async ()=>{
  if (ywList.value.length < totalCount.value && params.value.currentPage < pageCount.value) {
    params.value.currentPage = params.value.currentPage + 1
    loading.value = true
    //显示 触底的 loading 动画
    setTimeout(() => {
      getList()
    },500)
    // await getList()
  }else{
    loading.value = false
    finished.value = false
  }
}
onMounted(async()=>{
  let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
  params.value.corpCode = localStorage.getItem('corpCode') //从本地获取corpCode
  //从本地获取corpCode
  params.value.userId = userInfoData.userInfo.userId
 await getList()
})
</script>
<style lang="less" scoped>
.container{
  height:100vh;
}
.h-67 {
  height: 20vh;
  align-items: center;
  display: flex;
  justify-content: center;
}
.header {
  width: 100%;
  background-color: #fff;
  // padding: 0 10px;
  font-size: 14px;
  display: flex;
  text-align: center;
  height: 50px;
  line-height: 50px;
  border-bottom: 1px solid #f0f0f0;
  //margin-bottom: 50px;
  .menuItems {
    font-size: 14px;
    width: 40%;
    display: flex;
    justify-content: center;
    align-items: center;
    // border: 1px solid red;
  }
  .menuItems1{
    font-size: 14px;
    width: 20%;
    display: flex;
    justify-content: center;
    align-items: center;
    color: #0088ff;
    // border: 1px solid red;
  }
}

.btn-footer{
  height:10vh;
}

.list {
  height:calc(100vh - var(--van-nav-bar-height) - var(--van-tabs-line-height) - 50px - 10vh);
  overflow-y: auto;

  .list_card {
    width: 92vw;
    margin: 2vw;
    padding: 2vw;
    background-color: #FFFFFF;
    border-radius: 10px;

    .flex-space {
      display: flex;
      justify-content: space-between;
    }

    .list_title {
      font-size: 0.875em;
      font-weight: bold;
    }

    .f-z-12 {
      font-size: 0.75em;
      color: #4B5563;
    }

    .m-b-10 {
      margin-bottom: 2vw;
    }
  }
}

.mainButton {
  border-radius: 10px;
  padding: 2vw;
  height: 50px;
  background: url('../../assets/img/yinpin_bg.png') no-repeat center;
  background-size: 48vw 10vh;
  //background-image: linear-gradient(to right bottom, #0080FF, #6AB5FF);
}

.radio {
  position: fixed;
  bottom: 10vh;
}

.mainButton_sec {
  //margin-left: 2vw;
  border-radius: 10px;
  padding: 2vw;
  height: 50px;
  background: url('../../assets/img/wendangshangchuan_bg.png') no-repeat center;
  background-size: 48vw 10vh;
  //background-image: linear-gradient(to right bottom, #33B2AD, #3BD7D3);
}

:deep(.van-dropdown-menu__bar) {
  box-shadow: none;
}

:deep(.van-divider) {
  margin: 2vw;
}
:deep(.van-nav-bar__content) {
    background-color: #0088ff;
}
:deep(.van-nav-bar__title){
  color: #FFFFFF;
}
</style>
