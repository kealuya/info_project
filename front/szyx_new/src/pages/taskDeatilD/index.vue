<template>
    <taskCard :title="title" :taskTitle="taskTitle" :picSrc="picSrc" :target="target" :content="content" :taskId="taskId"
              :isComplete="isComplete" :finished="finished"></taskCard>
</template>
<script lang="ts" setup>
import bgc from '../../assets/img/task_bgc_01.png'
import {ref, onMounted, inject} from 'vue';
import taskCard from '../../components/taskCard/index.vue'
import {useRouter, useRoute} from "vue-router";
import {taskData} from '../../store/index'
import {myTaskDetails} from '../../services/task-processing/index'
// import { showSuccessToast, showFailToast } from 'vant';
import {showSuccessToast} from "vant";
// import {meetingListType} from '../../services/task-processing/types/index'
const router = useRouter()
const route = useRoute()
const isComplete = ref<Boolean>(true)
const title = ref<string>('任务详情')
const taskTitle = ref<string>('')
const picSrc = ref<string>()
const target = ref<string>()
const content = ref<string>()
const taskId = ref<string>()
// const meetingList = ref<any>()
// const taskId = ref<any>()
let finished = ref<any>()
const taskStore = taskData()

interface dataType {
    taskTitle: string | undefined,
    content: string | undefined
}

let data = ref<dataType>()
const onClickLeft = () => {
    // history.go(-2)
    router.replace('/homeNew')
}

// const taskProcessingHandle = ()=>{
//   // 将 ref 对象转换为普通 JavaScript 对象
//   const queryData: dataType = {
//     taskTitle: data.value?.taskTitle,
//     content: data.value?.content
//   };
//   console.log('queryData',queryData)
//   // 使用转换后的对象作为查询参数传递给路由器
//   router.replace({ path: '/addTasks', query: queryData as any });
// }
onMounted(() => {
    finished.value = route.query.isFinished
    console.log(' finished.value', finished.value)
    console.log('taskStore', taskStore.getTaskId())
    let corpCode = localStorage.getItem('corpCode')
    let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
    let userId = userInfoData.userInfo.userId
    // taskId.value = taskStore.getTaskId()
    let params = {
        taskId: taskStore.getTaskId(),
        corpCode: corpCode,
        userId: userId
    }
    myTaskDetails(params).then((res: any) => {
        if (res.success) {
            taskTitle.value = res.data.taskTitle
            target.value = res.data.taskData
            content.value = res.data.taskContent
            picSrc.value = res.data.taskImg
            taskId.value = res.data.taskId
            console.log(res.data.meetingList)
            // meetingList.value = res.data.meetingList
        }
    })
// console.log(route.query)
//   taskTitle.value = route.query.taskTitle as string
//   target.value = route.query.taskTarget as string
//   content.value = route.query.taskContent as string
})
</script>