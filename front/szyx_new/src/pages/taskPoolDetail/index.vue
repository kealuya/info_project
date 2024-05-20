<template>
    <van-nav-bar
            :title="title"
            left-text="返回"
            left-arrow
            @click-left="onClickLeft"
    />
    <div class="task_box">
        <div class="task f_white">
            <van-image :src="picSrc" height="20vh" width="96vw"></van-image>
        </div>
        <div class="height_contain">
            <div class="content">
                <div class="title_one m-b-10">{{ taskTitle }}</div>
                <div class="title_two f-z-14 f-w-550 m-b-10">
                    <span>任务Id</span>
                    <!--        <span class="m-l">发布日期</span>-->
                </div>
                <div class="f-z-14 m-b-10">&nbsp; &nbsp;{{ taskId }}</div>
                <div class="title_two f-z-14 f-w-550 m-b-10">
                    <span>任务介绍</span>
                    <!--        <span class="m-l">发布日期</span>-->
                </div>
                <div class="f-z-14 m-b-10">&nbsp; &nbsp;{{ target }}</div>
                <div class="title_two f-z-14 f-w-550 m-b-10">
                    <span>任务内容</span>
                </div>
                <div class="f-z-14 m-b-10">&nbsp;{{ content }}</div>
            </div>
        </div>
        <!--  <div class="footer_btn">-->
        <!--    <van-button block type="primary" @click="taskProcessingHandle">-->
        <!--      去完成-->
        <!--    </van-button>-->
        <!--  </div>-->
    </div>
    <div class="footer_btn" v-if="userJoinState=='No'">
        <van-button block type="primary" @click="taskProcessingHandle1" :loading="isShow" loading-text="立即参与">
            立即参与
        </van-button>
    </div>
  <!--    <div v-if="searchLoading" class="h-67">-->
  <!--        <van-loading type="spinner" color="#1989fa" :vertical="true"-->
  <!--        >加载中...-->
  <!--        </van-loading-->
  <!--        >-->
  <!--    </div>-->
</template>
<script lang="ts" setup>
import bgc from '../../assets/img/task_bgc_01.png'
import {ref, onMounted, inject} from 'vue';
import taskCard from '../../components/taskCard/index.vue'
import {useRouter, useRoute} from "vue-router";
import {taskData} from '../../store/index'
import {applyJoinTask, myTaskDetails, taskPoolDetails} from '../../services/task-processing/index'
// import { showSuccessToast, showFailToast } from 'vant';
import {setToastDefaultOptions, showConfirmDialog, showFailToast, showSuccessToast} from "vant";
// import {meetingListType} from '../../services/task-processing/types/index'
const router = useRouter()
const route = useRoute()
const toPool = ref<string>('poolDetail')
const isComplete = ref<Boolean>(false)
const title = ref<string>('任务池详情')
const taskTitle = ref<string>('')
const picSrc = ref<string>()
const target = ref<string>()
const content = ref<string>()
const taskId = ref<string>()
const isShow = ref<boolean>(false)
const userId = ref<string | undefined>()
const corpCode = ref<string | null>()
const userName = ref<string | null>()
const userMobile = ref<string | null>()

// const meetingList = ref<any>()
// const taskId = ref<any>()
let finished = ref<any>()
const taskStore = taskData()
const userJoinState = ref<string | undefined>()

interface dataType {
    taskTitle: string | undefined,
    content: string | undefined
}

const paramsData = ref<any>();
let data = ref<dataType>()
const onClickLeft = () => {
    // history.go(-2)
    router.replace('/taskPool')
}
const taskProcessingHandle1 = () => {
    showConfirmDialog({
        title: '参与任务',
        message:
            '你确定要参与这项任务吗？',
    })
        .then(() => {
            isShow.value = true  //打开 按钮上的 loading
            console.log(paramsData.value)
            let params = {
                userId: userId.value,
                creater:userId.value,
                corpCode: corpCode.value,
                taskId:paramsData.value.taskId,
                userName:userName.value,
                userMobile:userMobile.value
            }
            // searchLoading.value = true;
            applyJoinTask(params).then((res: any) => {
                if (res.success) {
                    // searchLoading.value = false;
                    //给提示，跳转到首页
                    showSuccessToast(res.msg);

                    setTimeout(() => {
                        isShow.value = false  //打开 按钮上的 loading
                        router.replace("/homeNew");
                    }, 1500);
                } else {
                    showFailToast(res.msg);
                    // searchLoading.value = false;
                }
            });
            // checkIsJoinTask(props.params).then((res: any) => {
            //     if (res.data == 'noJoinTask') {
            //         applyJoinTask(paramsData.value).then((res: any) => {
            //             console.log("res", res);
            //             if (res.success) {
            //                 searchLoading.value = false;
            //                 //给提示，跳转到首页
            //                 showSuccessToast(res.msg);
            //                 router.replace("/homeNew");
            //             } else {
            //                 showFailToast(res.msg);
            //                 searchLoading.value = false;
            //             }
            //         });
            //     } else {
            //         showFailToast(res.msg);
            //     }
            // });
        })
        .catch(() => {
            // on cancel
        });


    // applyJoinTask()
    // // isLoading.value = true
    // showSuccessToast('参与成功');
    // setTimeout(() => {
    //   router.replace('/homeNew')
    // }, 500);
};
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
    setToastDefaultOptions({ duration: 1000 });

    // console.log(route.query.userJoinState)
    userJoinState.value = route.query.userJoinState as string
    // console.log(route.query.userJoinState)
    paramsData.value = route.query

    corpCode.value = localStorage.getItem('corpCode')
    let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
    userName.value = userInfoData.userInfo.userName
    userMobile.value = userInfoData.userInfo.userMobile
    if (userInfoData) {
        userId.value = userInfoData.userInfo.userId
    }
    // taskId.value = taskStore.getTaskId()
    let params = {
        taskId: route.query.taskId
    }
    taskPoolDetails(params).then((res: any) => {
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
<style lang="less" scoped>
.task_box {
  overflow: hidden;
  display: flex;
  flex-direction: column;
  align-items: center;
  //justify-content: center;
  padding: 2vw;
  //height: calc(100vh - var(--van-nav-bar-height - 4vh);
  height: calc(100vh - var(--van-nav-bar-height) - 14vh);
  box-sizing: border-box;
}

.height_contain {
  flex: 1;
  width: 100%;
  overflow-y: auto; /* 当内容超出高度时显示垂直滚动条 */
}

.task {
  //margin-top: 2vw;
  //border-radius: 10px;
  //padding: 0 2vw;
  width: 96vw;
  //margin: 2vw;
  height: 20vh;
  //background: url('../../assets/img/task_bgc_02.png') no-repeat center;
  //background-size: cover;
  //box-sizing: border-box;
  //margin-left: 2vw;
  .task_card {
    width: 62vw;
    background-color: #e9f8f1;
    opacity: 0.5; /* 设置透明度为50% */
    color: #000000;
    //margin: 2vw;
    padding: 2vw;
    font-size: 1em;
    border-radius: 10px;

    .content-font {
      font-size: 0.75em;
      margin-top: 1vw;
    }
  }

  .btn {
    border-radius: 10px;
    background-color: #253334;
    width: 42vw;
    margin-left: 2vw;
    padding: 1vw;
    text-align: center;
    font-size: 1em;
  }
}

//.box{
//  height:60px;
//}
.content {
  //margin:2vw;
  line-height: 3vh;
  padding-top: 2vw;
  //flex:1;
  //overflow: auto;
  .title_one {
    text-align: center;
    font-weight: 550;
    font-size: 0.875em;
  }

  .f-z-14 {
    font-size: 0.75em;
  }

  //.f-z-12{
  //  font-size: 0.75em;
  //}
  .f-w-550 {
    font-weight: 550;
  }

  .m-l {
    margin-left: 20vw;
  }


}

.m-b-10 {
  margin-bottom: 10px;
}

.f-z-12-c {
  font-size: 0.75em;
  color: #4B5563;
}

.meeting-card {
  background-color: #FFFFFF;
  margin: 2vw 0;
  padding: 2vw;
  border-radius: 2vw;

  .title {
    display: flex;
    align-items: center;

    span {
      margin-right: 2vw;
    }
  }
}

.footer_btn {
  width: 96vw;
  height: 10vh;
  padding: 0 2vw;
  //border: 1px solid red;
  //line-height: 10vh;
  //position: fixed;
  //bottom:60px;
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

//.container {
//  flex: 1;
//  width: 100%;
//  overflow-y: auto; /* 当内容超出高度时显示垂直滚动条 */
//}
</style>