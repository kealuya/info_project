<template>
  <van-nav-bar
      title="完成任务"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"/>
  <div>
    <van-form>
      <div>
        <van-cell-group class="pb10 mt10">
          <van-field v-model="reasonForBorrowing"
                     name="reasonForBorrowing"
                     label="任务名称"
                     placeholder="请选择任务名称"
                     class="mt10 mb10"
                     label-width="7.5em"
                     maxlength="20"
                     readonly
                     label-align="top"/>
            <van-field v-model="taskId"
                       label="任务Id"
                       class="mt10 mb10"
                       label-width="7.5em"
                       maxlength="20"
                       readonly
                       label-align="top"/>
            <van-field v-model="taskTime"
                       label="创建时间"
                       class="mt10 mb10"
                       label-width="7.5em"
                       maxlength="20"
                       readonly
                       label-align="top"/>
          <van-notice-bar color="#1989fa" :scrollable="false" wrapable background="#ecf9ff" class="f10">
            任务内容：
           <p>{{contentTask}}</p>
          </van-notice-bar>

            <van-cell title="业务内容" is-link value="请选择业务内容" @click="taskList"/>


<!--          <div class="yw_content" @click="taskList">-->
<!--            业务内容<van-icon name="arrow" />-->
<!--          </div>-->
          <div class="list_card" v-for="item in meetingList" :key="item.meetingId">
            <div class="flex-space m-b-10">
              <div style="display: flex;align-items: center">
                <van-tag type="primary" v-if="item.meetingType == 'audio'">音频会议</van-tag>
                <van-tag type="success" v-else>文档记录</van-tag>
                <div class="list_title">{{item.meetingTitle}}</div>
              </div>
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
        </van-cell-group>
      </div>
    </van-form>
  </div>

  <div class="footer_btn">
    <van-button block type="primary" @click="showLoanStatus" :loading="isLoading" loading-text="提交中">
      提交任务
    </van-button>
  </div>
  <van-popup v-model:show="showPicker" round position="bottom">
    <van-picker
        :columns="columns"
        @cancel="showPicker = false"
        @confirm="onConfirm"
    />
  </van-popup>
  <div class="box"></div>
</template>
<script lang="ts" setup>
import {finishMyTask} from '../../services/task-processing/index'
import {taskData,meetingData} from '../../store/index'
import {inject, onMounted, ref} from 'vue'
import { useRouter,useRoute } from "vue-router"
import {showConfirmDialog, showSuccessToast,showFailToast} from "vant"
import {IMeetingDetailTypeResponse,finishTaskType} from '../../services/task-processing/types/index'
import word from "../../assets/img/word.png";
import xmind from "../../assets/img/xmind.png";
import mp3 from "../../assets/img/mp3.png";
const router = useRouter()
const route = useRoute()
const taskStore = taskData()
const meetingStore = meetingData()
const contentTask = ref<string>()
const isCollapse = ref<boolean>(true)
const taskId = ref<string>('')
const taskTime = ref<string>('')
const meetingList = ref<IMeetingDetailTypeResponse[]>([])
const isLoading = ref<boolean>(false)
const columns = [
  { text: '临床推广', value: '临床推广' },
  { text: '科研成果', value: '科研成果' } ,
  { text: '医药销售', value: '医药销售' }
];
const fieldValue = ref('');
const showPicker = ref(false);

const onConfirm = ({ selectedOptions}) => {
  showPicker.value = false;
  fieldValue.value = selectedOptions[0].text;
};

const onClickLeft = () => {
  // history.go(-2)
  router.replace({path:'/task',query:{isComplete:true,finished:false}})
}
const reasonForBorrowing = ref<string>()
const loanAmount = ref<string>('临床推广')

const taskProcessingHandle = ()=>{
  showSuccessToast('参与成功');

  router.replace('/homeNew')
}
const selectBankAccount = ()=>{
  router.replace('/taskList')
}

let params = ref<finishTaskType>({
  taskId:'',
  corpCode:'',
  userId:'',
  userName:'',
  userMobile:'',
  meetingIdList:''
})

//点击了提交 任务按钮
const showLoanStatus=()=> {
    if(meetingList.value.length==0){
        showFailToast('请选择业务内容！');
    }else{
        showConfirmDialog({
            title: '提交任务',
            message:
                '你确定要提交这几项任务吗？',
        })
            .then(() => {
                isLoading.value = true

                //以上是准备接口参数
                finishMyTask(params.value).then((res:any)=>{
                    if(res.success){
                        showSuccessToast(res.msg)
                        isLoading.value = false
                        router.replace({path:'/taskProcessing',query:{finish:true}})
                    }else{
                        showFailToast(res.msg)
                    }
                })
            })

    }
  // console.log('参数',params.value)



}
const taskList = ()=>{
  router.replace('/taskList')
}
onMounted(()=>{
  console.log(taskStore.getTaskTitle())
  reasonForBorrowing.value = taskStore.getTaskTitle()
  contentTask.value = taskStore.getTaskContent()
    taskId.value = taskStore.getTaskId()
    taskTime.value = taskStore.getTaskTime()
  meetingList.value =meetingStore.getMeetingData()  //从本地缓存中取 业务数据  勾选的
  if(meetingList.value){
    let arr =  meetingList.value.map((item:any)=>{
      return item.meetingId
    })
    params.value.meetingIdList = arr
  }
  params.value.taskId = taskStore.getTaskId()
  let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
  params.value.corpCode = localStorage.getItem('corpCode') //从本地获取corpCode
  params.value.userId = userInfoData.userInfo.userId
  params.value.userName = userInfoData.userInfo.userName
  params.value.userMobile = userInfoData.userInfo.userMobile
  router.beforeEach((to, from, next) => {
    // console.log('用户从哪个页面跳转到哪个页面');
    if(from.fullPath =='/addTasks'&&to.fullPath=='/taskList'){  //不清空
      next(); // 允许导航继续
    }else if(from.fullPath =='/taskList'&&to.fullPath=='/addTasks'){
      next(); // 允许导航继续
    }else{
      meetingStore.removeMeetingData()
      next(); // 允许导航继续
    }
  });
})
</script>
<style lang="less" scoped>
.yw_content{
  align-items: center;
  padding:10px;
  display:flex;
  justify-content: space-between;
}
.task{
  //margin-top: 2vw;
  //border-radius: 10px;
  //padding: 0 2vw;
  width:96vw;
  margin: 2vw;
  height: 160px;
  background: url('../../assets/img/home_01.png') no-repeat center;
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
    font-size: 1em;
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
    font-size: 1em;
  }
}
.box{
  height:50px;
}
.content{
  margin:2vw;
  .title_one{
    text-align:center;
    font-weight: 550;
    font-size: 0.875em;
  }
  .f-z-14{
    font-size: 0.75em;
  }
  //.f-z-12{
  //  font-size: 0.75em;
  //}
  .f-w-550{
    font-weight: 550;
  }
  .m-l{
    margin-left: 20vw;
  }
  .m-b-10{
    margin-bottom:10px;
  }
}
.footer_btn{
  width:95vw;
  padding:2vw;
}
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
.list_title1{
  font-size: 0.75em;
  //font-weight: 550;
  margin-left: 2vw;
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
<script setup lang="ts">
</script>
