<template>
    <van-nav-bar
            title="任务详情"
            left-text="返回"
            left-arrow
            @click-left="onClickLeft"
    />
    <div class="contain-card">
        <div class="f-z-16-c m-b-10" style="display:flex;align-items: center">  <van-image :src="finish" width="40" fit="contain" class="m-r-2"
                                                  height="40"></van-image>已完成的任务详情</div>
        <div class="item-card">
            <span class="f-z-14-c">任务ID</span>
            <div class="f-z-14-black">{{taskId}}</div>
        </div>
     <div class="item-card">
         <span class="f-z-14-c">任务标题</span>
         <div class="f-z-14-black">{{taskTitle}}</div>
     </div>
        <div class="item-card">
            <span class="f-z-14-c">任务类型</span>
            <div class="f-z-14-black">{{taskType}}</div>
        </div>
        <van-image :src="yiwancheng" class="img" wdith="70"  fit="contain"
                   height="70"></van-image>
    </div>
    <div class="contain-card">
    <div class="item-card">
        <div class="f-z-16-c m-b-10">任务介绍</div>
        <span class="f-z-14-black">{{taskData}}</span>
    </div>
    </div>
    <div class="contain-card">
    <div class="item-card">
        <div class="f-z-16-c  m-b-10">任务类型</div>
        <span class="f-z-14-black">{{taskContent}}</span>
    </div>
    </div>

    <div class="contain-card">
        <div class="f-z-16-c m-b-10">关联的业务内容</div>
        <div v-for="item in meetingList" class="m-b-10">
           <div class="title m-b-10">
               <van-tag type="success" v-if="item.meetingType=='document'">文档记录</van-tag>
               <van-tag type="primary" v-else>会议记录</van-tag>
              <span class="f-z-14-black m-l-10">{{item.meetingTitle}}</span>
           </div>
            <div class="f-z-14-4B5563 m-b-10">会议id：{{item.meetingId}}</div>
            <div class="f-z-14-4B5563 m-b-10">创建时间：{{item.createTime}}</div>
            <div class="f-z-14-4B5563 m-b-10">会议地址：{{item.meetingCity}}-{{item.meetingAddress}}</div>
            <div class="f-z-14-4B5563">参会人员：{{item.meetingPeople}}</div>
            <van-divider />
        </div>
    </div>
    <div style="height: 20px"></div>
</template>
<script setup lang="ts">
import finish from '../../assets/img/finish.png'
  import {useRoute,useRouter} from "vue-router";
  import {inject, onMounted, ref} from "vue";
  import {myTaskDetails} from '../../services/task-processing/index'
  import yiwancheng from "../../assets/img/yiwancheng.png";
  const route = useRoute()
  const router = useRouter()
  const taskTitle = ref<string>()  //标题
  const taskData = ref<string>()   //目标
  const taskContent = ref<string>()   //内容
  const meetingList = ref<any>()
  const taskId = ref<string>()
  const taskType = ref<string>()
  const params = ref<any>({
      corpCode:'',
      userId:'',
      taskId:''
  })
  const onClickLeft = ()=>{
      router.replace('/taskProcessing')
  }
  onMounted(()=>{
      let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
      console.log(route.query.id)
      params.value.corpCode = localStorage.getItem('corpCode')
      params.value.userId = userInfoData.userInfo.userId
      params.value.taskId = route.query.id
          console.log('params',params)
      myTaskDetails(params.value).then((res:any)=>{
          console.log('res',res)
          if(res.success){
           //todo
              taskTitle.value = res.data.taskTitle
              taskData.value = res.data.taskData
              taskContent.value = res.data.taskContent
              taskId.value = res.data.taskId
              taskType.value = res.data.taskType
              meetingList.value = res.data.meetingList
          }
      })
  })
</script>
<style lang="less" scoped>

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
.contain-card{
  border-radius: 10px;
    position: relative;
    padding: 2vw;
    margin:2vw;
    background-color: #fff;
    width:92vw;
  .item-card{
    padding-bottom: 2vw;
  }
    .title{
        display:flex;
        align-items: center;
    }
    .img{
        position: absolute;
        top:12vw;
        right:8vw;
    }
}
.f-z-14-c {
  font-size: 0.875em;
  color: rgba(23, 26, 29, 0.4);
  //font-weight: 550;
}
.f-z-14-black{
  font-size: 0.875em;
  color:#000000;
}
.f-z-16-c{
  font-size: 1em;
  color:#000000;
}
.m-b-10{
  margin-bottom: 10px;
}
.m-l-10{
    margin-left: 2vw;
}
.m-r-2{
  margin-right: 2vw;
}
.f-z-14-4B5563{
    font-size: 0.75em;
    color:#4B5563;
}


</style>