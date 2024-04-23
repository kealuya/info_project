
<template>
  <van-nav-bar
      title="任务池列表"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
  />
  <van-pull-refresh   v-model="isLoading" success-text="刷新成功" @refresh="onRefresh" :disabled="isShowImg">
    <van-list
        :finished-text="finishText"
        v-model:loading="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load = "getList()"

    >
      <div class="list"   v-if="!isShowImg">
        <div class="list_cardtask" v-for="item in listData">
          <div class="m-b-10 icon">
            <van-image
                width="20"
                height="20"
                :src="businessIcon"
            />
            <div class="list_title">{{item.taskTitle}}</div>
          </div>
          <van-divider />
          <div class="m-b-10 f-z-12"><span style="color: #4BA3FB" >任务目标：</span>{{item.taskTarget}}</div>
          <div class="m-b-10 f-z-12"><span style="color: #4BA3FB" >任务类型：</span>{{item.taskType}}</div>
          <div class="m-b-10 f-z-12" style="display: flex;justify-content: space-between">
            <div>
              <span style="color: #4BA3FB" >创建时间：</span>{{item.createTime}}
            </div>
            <van-tag type="success" v-if="item.taskStatus==1">已完成</van-tag>
            <van-tag type="danger" v-else-if="item.taskStatus==0">未完成</van-tag>
          </div>
          <!--          <div class="m-b-10 f-z-12"><span style="font-size: 12px;color: #4BA3FB">发布时间：<span>{{item.time}}</span></span></div>-->
        </div>
      </div>
      <div v-else>
        <img :src="empty" alt="" class="img">
      </div>
    </van-list>
  </van-pull-refresh>
  <div style="height:60px"></div>
</template>

<script setup lang="ts">
import {onBeforeMount, onMounted, ref,nextTick } from 'vue';
import { showFailToast } from 'vant';
import {getTaskPoolList} from "../../services/task-processing/index";
import empty from '../../assets/img/zanwupiaoju.png'
import  daiwancheng from '../../assets/img/daiwancheng.png'
import  yiwancheng from '../../assets/img/yiwancheng.png'
import shenqingdanIcon from  '../../assets/img/shenqingdan.png';
import danjuIcon from  '../../assets/img/danju.png';
import tabbar from '../../Tabbar.vue'
import {useRouter} from "vue-router";
import businessIcon from "../../assets/img/business_icon.png";
import screening_01 from "../../assets/icon/screening_01.png";
import screening_02 from "../../assets/icon/screening_02.png";
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
  taskStatus:Number,
  taskTarget:String,
  taskTitle:String,
  taskType:String
}
const active = ref(0);
const router = useRouter()
const listData = ref<LISTTYPE[]>([]); //列表的数据
const loading = ref(false); //van-list 的loading
const finished = ref(false);
const finishText = ref<string>('没有更多了')
const isLoading = ref<boolean>(false)  //控制 下拉刷新
const isShowImg = ref<boolean>(false)  //数据为空的时候 控制图片的显示隐藏
let params = ref<any>({
  currentPage:1,
  pageSize:10,
  status:'0',
  corpCode: "",
})
const businessDetailHandle = ()=>{
  router.replace('/task')

}
const onClickLeft = () => {
  // history.go(-2)
  router.replace('/homeNew')
}
//获取列表数据
const getList = ()=>{
  getTaskPoolList(params.value).then((res:any) => {
    if(res.success==false){
      showFailToast(res.msg)
    }else{
      listData.value = listData.value.concat(res.data.taskList)
      isLoading.value =false   //关闭下拉刷新的加载
      let page=  params.value.currentPage+1
      params.value.currentPage = page
      // if(res.data.totalCount==list.value.length){
      //   finished.value = true
      // }

      loading.value = false
      if(res.data.totalCount==0){
        isShowImg.value = true
        finishText.value = ''
      }else{
        isShowImg.value = false
        finishText.value = '没有更多数据了'
      }
      console.log(listData.value.length,res.data.totalCount)
      if(listData.value.length == res.data.totalCount){
        console.log('相等')
        finished.value = true
        loading.value = true
        console.log('停止')
      }else{
        finished.value = false
        loading.value = false
      }
      // list.value.length == res.data.totalCount?finished.value = true:finished.value = false
    }

  })
}
//下拉刷新的方法
const onRefresh = async()=>{
  isLoading.value = true   //打开下拉刷新的 loading
  params.value.currentPage = 1 //重置页码
  await getList() //重新加载数据
}
onBeforeMount(async () => {
  params.value.corpCode = await localStorage.getItem('corpCode') //从本地获取corpCode
})


</script>

<style lang="less" scoped>
.mainButton {
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 50px;
  background: url('../../assets/img/shenqigndan_bg.png') no-repeat center;
  background-size: 100vw 10vh;
  //background-image: linear-gradient(to right bottom, #0080FF, #6AB5FF);
}
.task{
  //margin-top: 2vw;
  border-radius: 10px;
  //padding: 0 2vw;
  width:100vw;
  height: 160px;
  background: url('../../assets/img/task_bgc_01.png') no-repeat center;
  background-size: cover;
  //box-sizing: border-box;
  //margin-left: 2vw;
  .task_card{
    width: 62vw;
    background-color: #e9f8f1;
    opacity: 0.5; /* 设置透明度为50% */
    color:#000000;
    margin: 2vw;
    padding:2vw;
    font-size: 0.875em;
    border-radius: 10px;
    .content-font{
      font-size: 0.75em;
      margin-top:1vw;
    }
  }
  .btn{
    border-radius: 10px;
    background-color: #253334;
    width:42vw;
    margin-left: 2vw;
    padding:1vw;
    text-align: center;
    font-size: 0.875em;
  }
}
.task2{
  margin-top: 4vw;
  border-radius: 10px;
  //padding: 0 2vw;
  width:96vw;
  height: 160px;
  background: url('../../assets/img/task_bgc_02.png') no-repeat center;
  background-size: cover;
  //box-sizing: border-box;
  margin-left: 2vw;
  .task_card{
    width: 62vw;
    background-color: #e9f8f1;
    opacity: 0.5; /* 设置透明度为50% */
    color:#000000;
    margin: 2vw;
    padding:2vw;
    font-size: 0.875em;
    border-radius: 10px;
    .content-font{
      font-size: 0.75em;
      margin-top:1vw;
    }
  }
  .btn{
    border-radius: 10px;
    background-color: #1768ed;
    width:42vw;
    margin-left: 2vw;
    padding:1vw;
    text-align: center;
    font-size: 0.875em;
  }

}
.list{
  //height: 100vh;
  overflow: auto;
  .list_card{
    position: relative;
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
      white-space: nowrap; /* 禁止文本换行 */
      overflow: hidden; /* 隐藏超出容器的部分文本 */
      text-overflow: ellipsis; /* 显示省略号 */
      width: 100vw; /* 设置容器宽度 */
      margin-left: 2vw;
    }
    .f-z-12{
      font-size: 0.75em;
      color:#4B5563;
    }
    .m-b-10{
      margin-bottom: 2vw;
    }
    .icon{
      display:flex;
      align-items: center;
    }
    .f-w{
      font-weight: 550;
      color:#000000;
    }
  }
  .list_cardtask{
    position: relative;
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
      white-space: nowrap; /* 禁止文本换行 */
      overflow: hidden; /* 隐藏超出容器的部分文本 */
      text-overflow: ellipsis; /* 显示省略号 */
      width: 100vw; /* 设置容器宽度 */
      margin-left: 2vw;
    }
    .f-z-12{
      font-size: 0.75em;
      color:#4B5563;
    }
    .m-b-10{
      margin-bottom: 2vw;
    }
    .icon{
      display:flex;
      align-items: center;
    }
    .f-w{
      font-weight: 550;
      color:#000000;
    }
  }
}


.box{
  height:50px;
}
:deep(.van-divider){
  margin: 10px;
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
