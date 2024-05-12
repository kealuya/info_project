<template>
  <taskCard :title="title" :taskTitle="taskTitle" :picSrc="picSrc" :target="target" :content="content" :isComplete="isComplete" :params="params"></taskCard>
</template>
<script lang="ts" setup>
import taskCard from '../../components/taskCard/index.vue'
import {useRouter,useRoute} from "vue-router";
// import { showSuccessToast, showFailToast } from 'vant';
import {showSuccessToast} from "vant";
import {inject, onMounted, ref} from "vue";
import bgc from "../../assets/img/task_bgc_01.png";
const router = useRouter();
const route = useRoute();
interface paramsType{
  taskId:string, //任务ID
  taskTitle:string,//任务标题
  corpName:string,//企业名称
  corpCode:string | null | undefined,//企业code
  userId:string,//用户ID
  userName:string,//用户姓名
  userMobile:string,//用户手机号
}
// const route = useRoute();
// const isLoading = ref(false)
const isComplete = ref<Boolean>(false)
const title = ref<string>('任务详情')
let taskTitle = ref<string>('')
const picSrc = ref<string>()
const target = ref<string>()
const content = ref<string>()
let params = ref<paramsType>({
  taskId:'', //任务ID
  taskTitle:'',//任务标题
  corpName:'',//企业名称
  corpCode:'',//企业code
  userId:'',//用户ID
  userName:'',//用户姓名
  userMobile:'',//用户手机号
})
const onClickLeft = () => {
  router.replace('/homeNew')
}
// const taskProcessingHandle = ()=>{
//   showSuccessToast('参与成功');
//   setTimeout(() => {
//     router.replace('/homeNew')
//   }, 2000);
// }
onMounted(async ()=>{
  console.log(111)
  let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
  params.value.corpCode = await localStorage.getItem('corpCode') //从本地获取corpCode
  taskTitle.value = route.query.taskTitle as string
  target.value = route.query.taskData as string
  content.value = route.query.taskContent as string
    picSrc.value = route.query.taskImg as string
    // console.log(' picSrc.value', picSrc.value)
  params.value.taskId = route.query.taskId as string
  params.value.taskTitle = route.query.taskTitle as string
  params.value.corpName = route.query.corpName as string
  // console.log('userInfoData',userInfoData.userInfo)
  params.value.userId = userInfoData.userInfo.userId
  params.value.userName = userInfoData.userInfo.userName
  params.value.userMobile = userInfoData.userInfo.userMobile

})
</script>