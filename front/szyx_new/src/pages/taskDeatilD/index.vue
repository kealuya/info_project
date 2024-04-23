<template>
  <taskCard :title="title" :taskTitle="taskTitle" :picSrc="picSrc" :target="target" :content="content" :isComplete="isComplete"></taskCard>
</template>
<script lang="ts" setup>
import bgc from '../../assets/img/task_bgc_01.png'
import { ref,onMounted } from 'vue';
import taskCard from '../../components/taskCard/index.vue'
import {useRouter,useRoute} from "vue-router";
// import { showSuccessToast, showFailToast } from 'vant';
import {showSuccessToast} from "vant";
const router = useRouter()
const route = useRoute()
const isComplete = ref<Boolean>(true)
const title = ref<string>('任务详情')
const taskTitle = ref<string>('')
const picSrc = ref<string>(bgc)
const target = ref<string>()
const content = ref<string>()
const onClickLeft = () => {
  // history.go(-2)
  router.replace('/homeNew')
}
const taskProcessingHandle = ()=>{
  // showSuccessToast('参与成功');

  router.replace('/addTasks')
}
onMounted(()=>{
// console.log(route.query)
  taskTitle.value = route.query.taskTitle as string
  target.value = route.query.taskTarget as string
  content.value = route.query.taskContent as string
})
</script>