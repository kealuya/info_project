<template>
  <!--  //最外层盒子必须加上100vh-->
    <div class="page-container">
        <van-nav-bar
                left-arrow
                left-text="返回"
                title="任务列表"
                @click-left="onClickLeft"
        />
        <van-tabs v-model:active="active" @change="changeList">
            <van-tab title="我的任务">

            </van-tab>
            <van-tab title="已完成">

            </van-tab>

        </van-tabs>
        <div :class="{ 'list1': tabIndex }" class="list">
            <van-pull-refresh v-model="isLoading"  class="refresh" success-text="刷新成功"
                              @refresh="onRefresh">
                <div v-if="searchLoading && rwList.length==0" class="h-67">
                    <van-loading :vertical="true" color="#1989fa" type="spinner">加载中...</van-loading>
                </div>
                <van-list
                        v-if="!isShowImg"
                        v-model:loading="loading"
                        :finished="finished"
                        :immediate-check="false"
                        :offset="30"
                        finished-text="没有更多了"
                        loading-text="加载中..."
                        @load="getListLoad"
                >
                    <div>
                        <div v-for="item in rwList" class="list_cardtask">
                            <div class="m-b-10 icon">
                                <van-image
                                        :src="businessIcon"
                                        height="20"
                                        width="20"
                                />
                                <div class="list_title">{{ item.taskTitle }}</div>
                                <!--                  //放弃任务的 图标-->
                                <van-icon name="delete-o" color="#e96352" size="20" @click.stop="deleteHandle(item.taskId)"/>
                            </div>
                            <van-divider/>
                            <div class="m-b-10 f-z-12"><span style="color: #4BA3FB">任务ID：</span>{{
                                    item.taskId
                                }}
                            </div>
                            <div @click="businessDetailHandle(item)">
                                <div class="m-b-10 f-z-12"><span style="color: #4BA3FB">任务介绍：</span>{{
                                    item.taskData
                                    }}
                                </div>
                                <div class="m-b-10 f-z-12" style="display: flex;justify-content: space-between">
                                    <div>
                                        <span style="color: #4BA3FB">创建时间：</span>{{ item.createTime }}
                                    </div>
                                    <van-tag :type="item.flag=='0'?'primary':'success'">
                                        {{ item.flag == '0' ? '待处理' : '已完成' }}
                                    </van-tag>
                                </div>
                            </div>
                        </div>
                    </div>
                    <!--          <div class="box"></div>-->
                </van-list>
                <van-row v-else-if="isShowImg" align="center" class="tc" justify="center">
                    <div>
                        <img src="../../assets/img/zanwupiaoju.png" style="height:50vw;width:80vw;object-fit: contain"/>
                        <div class="banner">暂无数据</div>
                    </div>
                </van-row>
            </van-pull-refresh>
        </div>
        <div class="list">
            <van-pull-refresh v-model="isLoading"  success-text="刷新成功" @refresh="onRefresh">
                <div v-if="searchLoading && rwList.length==0" class="h-67">
                    <van-loading :vertical="true" color="#1989fa" type="spinner">加载中...</van-loading>
                </div>
                <van-list
                        v-if="!isShowImg"
                        v-model:loading="loading"
                        :finished="finished"
                        :immediate-check="false"
                        :offset="30"
                        finished-text="没有更多了"
                        loading-text="加载中..."
                        @load="getListLoad"

                >
                    <div v-for="item in rwList" class="list_card" @click="finishDetail(item.taskId)">
                        <div class="flex-space m-b-10">
                            <van-image
                                    :src="businessIcon"
                                    height="20"
                                    width="20"
                            />
                            <div class="list_title">{{ item.taskTitle }}</div>
                            <van-image :src="yiwancheng"
                                       height="60"
                                       style="position:absolute;right:8vw;top:4vw" width="72"></van-image>
                        </div>
                        <van-divider/>
                        <div class="m-b-10 f-z-12"><span style="color: #4BA3FB">任务ID：</span>{{ item.taskId}}
                        </div>
                        <div class="m-b-10 f-z-12"><span style="color: #4BA3FB">任务介绍：</span>{{ item.taskData }}
                        </div>
                        <div class="m-b-10 f-z-12"><span style="color: #4BA3FB">创建时间：</span>{{ item.createTime }}
                        </div>

                        <div class="m-b-10 f-z-12"><span style="color: #4BA3FB">完成时间：</span>{{ item.finishTime }}
                        </div>
                        <!--          <van-image  width="70"-->
                        <!--                      v-if="item.isShow"-->
                        <!--                      height="70"-->
                        <!--                      :src="daiwancheng" style="position:absolute;right:8vw;top:4vw"></van-image>-->
                    </div>
                </van-list>
                <van-row v-else-if="isShowImg" align="center" class="tc" justify="center">
                    <div>
                        <img src="../../assets/img/zanwupiaoju.png" style="height:50vw;width:80vw;object-fit: contain"/>
                        <div class="banner">暂无数据</div>
                    </div>
                </van-row>

            </van-pull-refresh>
        </div>
    </div>
</template>

<script lang="ts" setup>
import {LISTTYPE, paramsTye} from '../../services/task-processing/types/index'
import {getTaskList,giveUpTask} from '../../services/task-processing/index'
import {inject, onMounted, ref} from 'vue';
import yiwancheng from '../../assets/img/yiwancheng.png'
import {useRouter,useRoute} from "vue-router";
import businessIcon from "../../assets/img/business_icon.png";
import {taskData} from '../../store/index'
import empty from "../../assets/img/zanwupiaoju.png";
import {showFailToast} from "vant/es";
import {showConfirmDialog, showSuccessToast} from "vant";
const route = useRoute()
const isLoading = ref<boolean>(false)  //控制 下拉刷新
const active = ref(0);
const router = useRouter()
const tabIndex = ref(false)
const totalCount = ref<number>(0) //总条数
const pageCount = ref<number>(0)  //总页码
const isShowImg = ref<boolean>(false)  //数据为空的时候 控制图片的显示隐藏
const finishText = ref<string>('没有更多了')
const searchLoading = ref<Boolean>(false)
const loading = ref(false);
const finished = ref(false);
const taskStore = taskData()
const disabled = ref<boolean>(false)
let lastRefreshTime = 0;
const refreshInterval = 15000; // 15秒
const tag = ref<number>(0)
const isFinished = ref<boolean>(false)   //控制 已完成按钮的显示隐藏
interface itemType {
    taskContent: string,
    taskData: string,
    taskId: string,
    taskImg: string,
    taskStatus: string,
    taskTarget: string
    taskTitle: string
}

let params = ref<paramsTye>(
    {
        currentPage: 1,
        pageSize: 10,
        corpCode: '',
        flag: '0',
        userId: ''
    }
)
// tab 切换
const changeList = async (name: number) => {
    if (name == 1) {
        params.value.currentPage = 1
        rwList.value = []
        loading.value = false
        finished.value = false
        // searchLoading.value = false
        // isFinished.value = true
        // isLoading.value = true
        await getList()
        tabIndex.value = true
        tag.value = 1
    } else {
        // isFinished.value = false
        // console.log('进来了')
        params.value.currentPage = 1
        rwList.value = []
        loading.value = false
        finished.value = false
        await getList()
        tabIndex.value = false
        tag.value = 1

    }
}
//下拉刷新
const onRefresh = async () => {
    const currentTime = Date.now();
    if (currentTime - lastRefreshTime < refreshInterval) {
        showFailToast("刷新操作太频繁，请稍后再试！");
        isLoading.value = false; // 关闭下拉刷新的 loading
        return; // 直接返回，不执行加载数据的操作
    }
    lastRefreshTime = currentTime;
    isLoading.value = true; // 打开下拉刷新的 loading
    params.value.currentPage = 1; // 重置页码
    rwList.value = []; // 清空表单
    await getList(); // 重新加载数据
    showSuccessToast("刷新成功！");

    isLoading.value = false; // 关闭下拉刷新的 loading
};
//任务数据列表
const rwList = ref<LISTTYPE[]>([]);
const rwListComplete = ref<LISTTYPE[]>([]);
const toDetail = () => {
    router.replace('/taskDetail')
}
const businessDetailHandle = (item: any) => {
    taskStore.setTaskId(item.taskId)
    taskStore.setTaskTitle(item.taskTitle)
    taskStore.setTaskContent(item.taskContent)
    taskStore.setTaskId(item.taskId)
    taskStore.setTaskTime(item.createTime)
    router.replace({path: '/task', query: {isFinished: false}});

    // console.log(active.value)
    // if (active.value == 1) {
    //   //把数据存到 pinia
    //   router.replace({path: '/task', query: {isFinished: true}});
    // } else {
    //   router.replace({path: '/task', query: {isFinished: false}});
    // }


}
//已完成的 详情页 （单独加的）
const finishDetail = (id: number) => {
    router.replace({path: '/finishDetail', query: {id: id}})
}
const onClickLeft = () => {
    // history.go(-2)
    router.replace('/homeNew')
}
const getList = async () => {
    params.value.flag = String(active.value)
    // searchLoading.value = true
    // if( tag.value = 0){
    //     searchLoading.value = true
    // }else{
    //     searchLoading.value = false
    // }
    loading.value = true
    await getTaskList(params.value).then(async (res: any) => {

        if (res.success) {
            // rwList.value.length = 0
            // if (active.value == 0) {  //说明是我的任务
            rwList.value = rwList.value.concat(res.data.myTaskList)
            isLoading.value = false   //关闭下拉刷新的加载
            totalCount.value = res.data.totalCount  //总条数
            //todo 总页
            pageCount.value = res.data.pageCount //总页码
            if (rwList.value.length == totalCount.value) {
                loading.value = false   //关闭下拉刷新的加载
                finished.value = true
            }
            if (res.data.totalCount == 0) {
                isShowImg.value = true
                finishText.value = ''
            } else {
                isShowImg.value = false
                finishText.value = '没有更多数据了'
            }
        }
    }).finally(() => {
        searchLoading.value = false
        loading.value = false
    })
}
//触底方法
const getListLoad = async () => {
    // loading.value = true
    //这个判断必须加 否则会调用多次接口
    if (rwList.value.length < totalCount.value && params.value.currentPage < pageCount.value) {
        params.value.currentPage = params.value.currentPage + 1
        setTimeout(() => {
            getList()
        }, 500)

    }

}
//放弃任务接口
const deleteHandle = (id: string) => {
    let deleteParams={
        taskId: id, //任务ID
        corpCode: params.value.corpCode, //企业code
        creater: params.value.userId //人员id
    }
    showConfirmDialog({
        title: '删除',
        message:
            '你确定要删除吗？',
    })
        .then(() => {
            giveUpTask(deleteParams).then((res:any)=>{
                if(res.success){
                    showSuccessToast(res.msg)
                    //刷新列表数据
                    rwList.value.length = 0
                    params.value.currentPage = 1
                    getList()
                }else {
                    showFailToast(res.msg)
                }
            })
        })
        .catch(() => {
            // on cancel
        });
}
onMounted(async () => {
    // console.log(route.query.finish,'sahjcvdjsvjdsjkvjknvj ')
    let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
    params.value.corpCode = localStorage.getItem('corpCode') //从本地获取corpCode
    //从本地获取corpCode
    params.value.userId = userInfoData.userInfo.userId
    console.log('params', userInfoData.userInfo)
    //初始化
    rwList.value.length = 0
    params.value.currentPage = 1
    await getList()
    if(route.query.finish){
        active.value = 1
        params.value.currentPage = 1
        rwList.value = []
        loading.value = false
        finished.value = false
        await getList()
        tabIndex.value = true
        tag.value = 1
    }else{
        active.value = 0
    }
})
</script>

<style lang="less" scoped>
.page-container {
  height: 100vh;
  overflow: hidden;
}

.list1 {
  display: none;
}

.mainButton {
  border-radius: 10px;
  padding: 2vw;
  width: 42vw;
  height: 50px;
  background: url('../../assets/img/shenqigndan_bg.png') no-repeat center;
  background-size: 100vw 10vh;
  //background-image: linear-gradient(to right bottom, #0080FF, #6AB5FF);
}

.task {
  //margin-top: 2vw;
  border-radius: 10px;
  //padding: 0 2vw;
  width: 100vw;
  height: 160px;
  background: url('../../assets/img/task_bgc_01.png') no-repeat center;
  background-size: cover;
  //box-sizing: border-box;
  //margin-left: 2vw;
  .task_card {
    width: 62vw;
    background-color: #e9f8f1;
    opacity: 0.5; /* 设置透明度为50% */
    color: #000000;
    margin: 2vw;
    padding: 2vw;
    font-size: 0.875em;
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
    font-size: 0.875em;
  }
}

.task2 {
  margin-top: 4vw;
  border-radius: 10px;
  //padding: 0 2vw;
  width: 96vw;
  height: 160px;
  background: url('../../assets/img/task_bgc_02.png') no-repeat center;
  background-size: cover;
  //box-sizing: border-box;
  margin-left: 2vw;

  .task_card {
    width: 62vw;
    background-color: #e9f8f1;
    opacity: 0.5; /* 设置透明度为50% */
    color: #000000;
    margin: 2vw;
    padding: 2vw;
    font-size: 0.875em;
    border-radius: 10px;

    .content-font {
      font-size: 0.75em;
      margin-top: 1vw;
    }
  }

  .btn {
    border-radius: 10px;
    background-color: #1768ed;
    width: 42vw;
    margin-left: 2vw;
    padding: 1vw;
    text-align: center;
    font-size: 0.875em;
  }

}

.refresh {
  min-height: calc(100vh - var(--van-nav-bar-height) - var(--van-tabs-line-height));
}

.list { //必须算
  //height: 100vh;
  height: calc(100vh - var(--van-nav-bar-height) - var(--van-tabs-line-height));
  overflow-y: auto;

  .list_card {
    position: relative;
    width: 92vw;
    margin: 2vw;
    padding: 2vw;
    background-color: #FFFFFF;
    border-radius: 10px;

    .flex-space {
      display: flex;
      justify-content: space-between;
    }

    .list_title {
      font-size: 0.875em;
      font-weight: 550;
      white-space: nowrap; /* 禁止文本换行 */
      overflow: hidden; /* 隐藏超出容器的部分文本 */
      text-overflow: ellipsis; /* 显示省略号 */
      width: 100vw; /* 设置容器宽度 */
      margin-left: 2vw;
    }

    .f-z-12 {
      font-size: 0.75em;
      color: #4B5563;
    }

    .m-b-10 {
      margin-bottom: 2vw;
    }

    .icon {
      display: flex;
      align-items: center;
    }

    .f-w {
      font-weight: 550;
      color: #000000;
    }
  }

  .list_cardtask {
    position: relative;
    width: 92vw;
    margin: 2vw;
    padding: 2vw;
    background-color: #FFFFFF;
    border-radius: 10px;

    .flex-space {
      display: flex;
      justify-content: space-between;
    }

    .list_title {
      font-size: 0.875em;
      font-weight: 550;
      white-space: nowrap; /* 禁止文本换行 */
      overflow: hidden; /* 隐藏超出容器的部分文本 */
      text-overflow: ellipsis; /* 显示省略号 */
      width: 100vw; /* 设置容器宽度 */
      margin-left: 2vw;
    }

    .f-z-12 {
      font-size: 0.75em;
      color: #4B5563;
    }

    .m-b-10 {
      margin-bottom: 2vw;
    }

    .icon {
      display: flex;
      align-items: center;
    }

    .f-w {
      font-weight: 550;
      color: #000000;
    }
  }
}


.box {
  height: 50px;
}

.h-67 {
  height: 20vh;
  align-items: center;
  display: flex;
  justify-content: center;
}

:deep(.van-divider) {
  margin: 10px;
}

//::v-deep(.van-pull-refresh){
//  height:100%;
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
