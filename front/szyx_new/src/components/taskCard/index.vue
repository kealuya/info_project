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
                <div class="f-z-14 m-b-10">&nbsp; &nbsp;{{taskId}}</div>
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
    <div class="footer_btn" v-if="isComplete ||isShowComplate ">
        <van-button
                block
                type="primary"
                @click="taskProcessingHandle"
                v-if="finished == 'false' || isShowFinish=='false'"
        >
            去完成
        </van-button>
    </div>
    <div class="footer_btn" v-else>
        <van-button block type="primary" @click="taskProcessingHandle1">
            立即参与
        </van-button>
    </div>
    <div v-if="searchLoading" class="h-67">
        <van-loading type="spinner" color="#1989fa" :vertical="true"
        >加载中...
        </van-loading
        >
    </div>
</template>
<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {useRouter, useRoute} from "vue-router";
import {
    applyJoinTask,
    checkIsJoinTask,
} from "../../services/task-processing/index";
import {showSuccessToast, showFailToast} from "vant";
import {showConfirmDialog} from 'vant';
const id = ref<string>()
const route = useRoute()
const props = defineProps({
    title: {
        type: String,
        default: "",
    },
    taskId:{
        type: String,
        default: "",
    },
    taskTitle: {
        type: String,
        default: "",
    },
    picSrc: {
        type: String,
        default: "",
    },
    target: {
        type: String,
        default: "",
    },
    content: {
        type: String,
        default: "",
    },
    isComplete: {
        type: Boolean,
        default: false,
    },
    params: {
        type: Object,
        default: {},
    },
    finished: {
        type: String,
        default: "",
    },
});
const searchLoading = ref<Boolean>(false);
const isShowComplate = ref<any>()
const isShowFinish = ref<any>()
const paramsData = ref<any>();
const router = useRouter();
const onClickLeft = () => {
    // history.go(-1)
    router.replace("/taskProcessing");
};
const taskProcessingHandle = () => {
    router.replace("/addTasks");
    // router.replace('/addTasks')
};
const taskProcessingHandle1 = () => {
    showConfirmDialog({
        title: '参与任务',
        message:
            '你确定要参与这项任务吗？',
    })
        .then(() => {
            searchLoading.value = true;
            applyJoinTask(paramsData.value).then((res: any) => {
                if (res.success) {
                    searchLoading.value = false;
                    //给提示，跳转到首页
                    showSuccessToast(res.msg);
                    setTimeout(() => {
                        router.replace("/homeNew");
                    }, 2000);
                } else {
                    showFailToast(res.msg);
                    searchLoading.value = false;
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
onMounted(() => {
    console.log(props.toPool)
    paramsData.value = props.params;
    if (route.query) {
        isShowComplate.value = route.query.isComplete
        isShowFinish.value = route.query.finished
        // id.value = paramsData.value.taskId
    }
    // console.log(' paramsData.value', paramsData.value)
});
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
.f-z-12-c{
    font-size: 0.75em;
    color: #4B5563;
}
.meeting-card{
    background-color: #FFFFFF;
    margin: 2vw 0;
    padding:2vw;
    border-radius: 2vw;
    .title{
        display: flex;
        align-items: center;
        span{
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
<script setup lang="ts"></script>
