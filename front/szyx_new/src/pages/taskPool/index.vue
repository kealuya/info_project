<template>
  <div class="container">
    <van-nav-bar
        title="任务池列表"
        left-text="返回"
        left-arrow
        @click-left="onClickLeft"
    />
    <div class="list">
      <van-pull-refresh  v-model="isLoading" success-text="刷新成功" @refresh="onRefresh" :disabled="isShowImg">
        <van-list
            :finished-text="finishText"
            v-model:loading="loading"
            :finished="finished"
            :offset="30"
            finished-text="没有更多了"
            :immediate-check="false"
            loading-text="正在加载中请稍后"
            @load="getListLoad"

        >
          <div v-if="!isShowImg">
            <div class="list_cardtask" v-for="item in listData" @click="toDetail(item)">
              <div class="m-b-10 icon">
                <van-image
                    width="20"
                    height="20"
                    :src="businessIcon"
                />
                <div class="list_title">{{item.taskTitle}}</div>
              </div>
              <!--              <van-divider />-->
              <div class="m-b-10 f-z-12"><span style="color: #4BA3FB" >任务目标：</span>{{item.taskData}}</div>
<!--              <div class="m-b-10 f-z-12"><span style="color: #4BA3FB" >任务类型：</span>{{item.taskType}}</div>-->
              <div class="m-b-10 f-z-12" style="display: flex;justify-content: space-between">
                <div>
                  <span style="color: #4BA3FB" >创建时间：</span>{{item.createTime}}
                </div>
                <van-tag type="success" v-if="item.taskStatus==1">已完成</van-tag>
                <van-tag type="warning" v-else-if="item.taskStatus==0">未参与</van-tag>
              </div>
              <!--          <div class="m-b-10 f-z-12"><span style="font-size: 12px;color: #4BA3FB">发布时间：<span>{{item.time}}</span></span></div>-->
            </div>
          </div>
          <div v-else>
            <img :src="empty" alt="" class="img">
          </div>
        </van-list>
      </van-pull-refresh>
    </div>
    </div>
</template>

<script setup lang="ts">
import {onBeforeMount, onMounted, ref,nextTick } from 'vue';
import { showFailToast } from 'vant';
import {getTaskPoolList} from "../../services/task-processing/index";
import empty from '../../assets/img/zanwupiaoju.png'
import {useRouter} from "vue-router";
import businessIcon from "../../assets/img/business_icon.png";
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
const pageCount = ref<number>(0)
const totalCount = ref<number>(0)
let params = ref<any>({
  currentPage:1,
  pageSize:10,
  status:'0',
  corpCode: "",
})
const businessDetailHandle = ()=>{
  router.replace('/task')

}
//点击跳转到详情页面
const toDetail=(row:any)=>{
  router.replace({name:'taskDetail',query:row})
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
      totalCount.value = res.data.totalCount
      pageCount.value = res.data.pageCount
      if (listData.value.length == totalCount.value) {
        loading.value = false   //关闭下拉刷新的加载
        finished.value = true
      }else{
        loading.value = false
        finished.value = false
      }
      if(res.data.totalCount==0){
        isShowImg.value = true
        finishText.value = ''
      }else{
        isShowImg.value = false
        finishText.value = '没有更多数据了'
      }
    }

  })
}
//触底事件
const getListLoad = async ()=>{
  console.log('111111111111111111',listData.value.length,totalCount.value)
  if (listData.value.length < totalCount.value && params.value.currentPage < pageCount.value) {
    console.log(0)
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
//下拉刷新的方法
const onRefresh = async()=>{
  isLoading.value = true   //打开下拉刷新的 loading
  listData.value = []
  params.value.currentPage = 1 //重置页码
  await getList() //重新加载数据
}
onBeforeMount(async () => {
  params.value.corpCode = await localStorage.getItem('corpCode') //从本地获取corpCode
  getList()
})


</script>

<style lang="less" scoped>
.container{
  height:100vh;
  .list{
    height: calc(100vh - var(--van-nav-bar-height) - var(--van-tabbar-height));
    overflow-y: auto;
    .list_cardtask{
      width:92vw;
      margin: 2vw;
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

}
//:deep(.van-divider){
//  margin: 10px;
//}
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
