
<template>
    <van-image :src="bgsy" width="100%" height="180"></van-image>
  <van-row  justify="space-around" :getter="10" align="center" class="mt10">
    <van-col span="11">
      <div class="mainButton_zsk f_white" @click="hangyezhishiku">
        <van-row justify="space-between" align="center" >
          <van-col>
            <p class="f16">行业知识库</p>
            <p class="f12">资源共享 轻松查找共享信息</p>
          </van-col>
<!--          <van-col>-->
<!--            <van-icon :name="shenqingdanIcon" size="3em"   />-->
<!--          </van-col>-->
        </van-row>
      </div>
    </van-col>
    <van-col span="11">
      <div class="mainButton f_white">
        <van-row justify="space-between" align="center" @click="taskProcessingHandle">
          <van-col>
            <p class="f16">任务列表</p>
            <p class="f12">最新发布 参与流程</p>
          </van-col>
<!--          <van-col>-->
<!--            <van-icon :name="shenqingdanIcon" size="3em"   />-->
<!--          </van-col>-->
        </van-row>
      </div>
      <div class="mainButton_sec  f_white">
        <div>
          <van-row justify="space-between" align="center" @click="applicationHandle">
            <van-col>
              <p class="f16">价值申请</p>
              <p class="f12">一键申请  随时随地 </p>
            </van-col>
<!--            <van-col>-->
<!--              <van-icon :name="danjuIcon" size="3em"  />-->
<!--            </van-col>-->
          </van-row>
        </div>
      </div>
    </van-col>
  </van-row>
  <div class="mt5">
    <van-row class="mt10 card_padding" align="center">
      <span class="title">任务发布</span> <span class="banner">更多海量任务、为营销提供清晰的方向</span>
    </van-row>
    <div>
    </div>
  </div>
    <van-list
        v-model="loading"
        :finished="finished"
        :immediate-check="false"
        loading-text="正在加载中请稍后"
        finished-text="没有更多了"
    >
      <van-row v-if="list.length>0">
        <div class="task2  f_white" v-for="item in list" :key="item.taskId">
          <van-image :src="item.taskImg" class="task_bgc">
          </van-image>
          <div class="task_card">
            <div class="task_title">
              {{item.taskTitle}}
            </div>
            <div class="content-font">
              {{item.taskContent}}
            </div>
          </div>
          <div class="btn"  @click="toDetail(item)">
            <van-icon name="play-circle-o" /> 任务详情
          </div>
        </div>
      </van-row>
      <van-row justify="center" align="center" class="tc" v-else-if="isShowImg">
        <div>
          <img style="height:50vw;width:80vw;object-fit: contain" src="../../assets/img/zanwupiaoju.png" />
          <div class="banner">暂无数据</div>
        </div>
      </van-row>
     <van-row justify="center" align="center">
       <van-button type="primary" class="more_btn" @click="moreTasks" v-if="list.length>0">查看更多</van-button>
     </van-row>
    </van-list>
  <div v-if="searchLoading" class="h-67">
    <van-loading type="spinner" color="#1989fa" :vertical="true">加载中...</van-loading>
  </div>
  <div class="box"></div>
</template>

<script setup lang="ts">
import  test from '../../assets/img/task_bgc_01.png'
  import bgsy from '../../assets/img/bgsy.png';
import { showSuccessToast, showFailToast, showToast } from 'vant';
import shenqingdanIcon from  '../../assets/img/shenqingdan.png';
import danjuIcon from  '../../assets/img/danju.png';
import tabbar from '../../Tabbar.vue'
import {useRouter} from "vue-router";
  import ApplicationCard from "../home/component/applicationCard.vue";
  import {onMounted, ref} from "vue";
  import {getDocStateList} from "../../services/home";
  import {getTaskPoolList} from "../../services/task-processing/index";
  //定义列表数据类型
  interface listType {
    currentPage:Number, //当前页
    pageSize:Number, //每页条数
    searchKey:String,
    corpCode:String,//企业code
    status:String,	//是	状态，0未下架，1已下架；当前接口传0
    taskList:[],
    totalCount:0
  }
  //定义后端返给list数组的数据类型
  interface  LISTTYPE{
    bz1: String,
    bz2: String,
    bz3: String,
    corpCode:String,
    corpName:String,
    createTime:String,
    creater:String,
    taskContent:String,
    taskId:String,
    taskImg:String,
    taskStatus:String,
    taskTarget:String,
    taskTitle:String,
    taskType:String
  }
//列表数据 数组
const list = ref<LISTTYPE[]>([]);
const router = useRouter();
const loading = ref(true);
const finished = ref(false);

const searchLoading = ref<Boolean>(false)
const isShowImg = ref<Boolean>(false)
const params = ref<any>({
  currentPage:1,
  pageSize:10,
  status:'0',
  corpCode: "",
})
  const getDocList=async(payload:any,code:any)=>{
  // console.log('payload',payload)
    searchLoading.value = true
  getTaskPoolList(payload).then((res:any)=>{
    list.value.push(...res.data.taskList.splice(0,5));
   list.value.length ==0?isShowImg.value=true:isShowImg.value=false
    // console.log('list.value',res.data.taskList)
    }).finally(()=>{
    searchLoading.value = false
  });
  }
const toDetail=(row:any)=>{
  router.replace({name:'taskDetail',query:row})
}
// const toDetail2=()=>{
//   router.replace('/taskDetail_two')
// }
// const toDetail3=()=>{
//   router.replace('/taskDetail_three')
// }
// const toDetail4=()=>{
//   router.replace('/taskDetail_four')
// }
const taskProcessingHandle = ()=>{
  router.replace('/taskProcessing')

}
const hangyezhishiku = ()=>{
  showToast('暂未开放，敬请期待！');
}
const applicationHandle = ()=>{
  router.replace('/applicationValue')
}
//点击了查看更多按钮 跳转到 任务池列表
const moreTasks = ()=>{
  router.push('/taskPool')
}
onMounted(async () => {
  params.value.corpCode = await localStorage.getItem('corpCode')
  getDocList(params.value, '')
})
</script>

<style lang="less" scoped>
.h-67{
  height:67vh;
  z-index:999;
}
.title_banner{
  font-size: 0.875em;
}
.headBg {
  width: 100vw;
  height: 22vh;
  background-color: #0080FF;
}
.card_padding{
  padding:2vw;
}
.title{
  color:#000000;
  font-size: 16px;
}
.more_btn{
  width:96vw;
  padding:2vw;
}
.mainButton {
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 50px;
  background: url('../../assets/img/shenqingdan_bg.png') no-repeat center;
  background-size: 48vw 10vh;
  margin-bottom: 10px;
  //background-image: linear-gradient(to right bottom, #0080FF, #6AB5FF);
}

.mainButton_sec {
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 50px;
  background: url('../../assets/img/tidanju_bg.png') no-repeat center;
  background-size: 48vw 10vh;
  //background-image: linear-gradient(to right bottom, #33B2AD, #3BD7D3);
}
.mainButton_zsk{
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 125px;
  background: url('../../assets/img/zhishiku_bg.png') no-repeat center;
  background-size: 48vw 20vh;
}
.banner{
  color: #bbc2cc;
  font-size: 12px;
  margin: 5px 10px 0
}
.box{
  height:80px;
}
.task_title{
  white-space: nowrap; /* 禁止文本换行 */
  overflow: hidden; /* 隐藏超出容器的部分文本 */
  text-overflow: ellipsis; /* 显示省略号 */
  width: 55vw; /* 设置容器宽度 */
  font-size: 1.1em;
  font-weight: 380;
}
.task2{
  padding:2vw;
  position: relative;
  border-radius: 10px;
  width:96vw;
  //height: 160px;
  .task_card{
    position: absolute;
    top:5vw;
    left:4vw;
    width: 58vw;
    background-color: #000000;
    opacity: 0.5; /* 设置透明度为50% */
    color:#ffffff;
    padding:2vw;
    font-size: 0.875em;
    border-radius: 10px;
    .content-font{
      font-size: 0.75em;
      margin-top:1vw;
      display: -webkit-box;
      -webkit-box-orient: vertical;
      -webkit-line-clamp: 3;
      overflow: hidden;
      text-overflow: ellipsis;
    }
  }
  .btn{
    position: absolute;
    top: 30vw;
    border-radius: 10px;
    background-color: #1768ed;
    width:42vw;
    margin-left: 2vw;
    padding:1vw;
    text-align: center;
    font-size: 0.875em;
  }
}

</style>
