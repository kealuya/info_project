<template>
    <div class="listBox">
        <van-nav-bar
                left-arrow
                left-text="返回"
                title="业务内容"
                @click-left="onClickLeft"
        />
        <van-tabs v-model:active="active" @change="changeList">
            <div :class="{ 'list1': tabIndex }">
                <van-tab title="会议文件">
                    <van-pull-refresh v-model="isLoading" success-text="刷新成功" @refresh="onRefresh">
                        <div v-if="searchLoading" class="h-67">
                            <van-loading :vertical="true" color="#1989fa" type="spinner">加载中...</van-loading>
                        </div>
                        <van-list
                                v-if="!isShowImg"
                                v-model:loading="loading"
                                :finished="finished"
                                :immediate-check="false"
                                :offset="80"
                                class="list"
                                finished-text="没有更多了"
                                loading-text="加载中..."
                                @load="getMeetingListLoad"
                        >
                            <div v-for="item in meetingList" :key="item.meetingId" class="list_card">
                                <div class="flex-space m-b-10">
                                    <div style="display: flex;align-items: center">
                                        <van-tag v-if="item.meetingType == 'audio'" type="primary">音频会议</van-tag>
                                        <van-tag v-else type="success">文档记录</van-tag>
                                        <div class="list_title">{{ item.meetingTitle }}</div>
                                    </div>
                                    <van-checkbox v-model="item.isCheck" @change="changeCheck"></van-checkbox>
                                    <!--            <van-image-->
                                    <!--                width="20"-->
                                    <!--                height="20"-->
                                    <!--                :src="businessIcon"-->
                                    <!--            />-->
                                </div>
                                <van-divider/>
                                <!--          //display:flex;align-items: center;-->
                                <div v-for="i in item.meetingFile"
                                     style="margin-bottom: 10px;display:flex;align-items: center;">
                                    <van-image v-if="i.fileType=='word'" :src="word" height="23" width="23"></van-image>
                                    <van-image v-if="i.fileType=='xmind'" :src="xmind" height="23"
                                               width="23"></van-image>
                                    <van-image v-if="i.fileType=='mp3'" :src="mp3" height="23" width="23"></van-image>
                                    <div class="list_title1">{{ i.fileName }}</div>
                                    <!--        <div class="list_title1">{{item.meetingBrainMapName?item.meetingBrainMapName:''}}</div>-->
                                    <!--        <div class="list_title1">{{item.meetingMminutesName?item.meetingMminutesName:''}}</div>-->
                                </div>
                                <div>
                                    <!--            <div class="f-z-12 m-b-10" >所属会议：{{item.sshy}}</div>-->
                                    <!--<div class="f-z-12 m-b-10" v-if="item.hyType!=='word'">会议地点：{{item.address}}</div>-->
                                    <div class="f-z-12 m-b-10">会议ID:{{ item.meetingId }}</div>
                                    <div class="f-z-12 m-b-10">创建时间：{{ item.createTime }}</div>

                                </div>
                            </div>
                            <!--          <businessCard :meetingList="meetingList" @changeCheck="changeCheck"></businessCard>-->
                        </van-list>
                        <van-row v-else-if="isShowImg" align="center" class="tc" justify="center">
                            <div>
                                <img src="../../assets/img/zanwupiaoju.png"
                                     style="height:50vw;width:80vw;object-fit: contain"/>
                                <div class="banner">暂无数据</div>
                            </div>
                        </van-row>
                        <div v-if="isShowButton" class="footer_btn">
                            <van-button block type="primary" @click="handleSubmit">
                                提交
                            </van-button>
                        </div>
                    </van-pull-refresh>
                </van-tab>
            </div>
            <div :class="{ 'list1': tabIndex==false}" class="listUse">
                <van-tab title="已使用">
                    <van-pull-refresh v-model="isLoading" success-text="刷新成功" @refresh="onRefresh">
                        <div v-if="searchLoading" class="h-67">
                            <van-loading :vertical="true" color="#1989fa" type="spinner">加载中...</van-loading>
                        </div>
                        <van-list
                                v-if="!isShowImg"
                                v-model:loading="loading"
                                :finished="finished"
                                :immediate-check="false"
                                :offset="10"
                                finished-text="没有更多了"
                                loading-text="加载中..."
                                style="height: 100%"
                                @load="getMeetingListLoad"
                        >
                            <div v-for="item in meetingList" :key="item.meetingId" :class="{ 'list1': !tabIndex }"
                                 class="list_card">
                                <div class="flex-space m-b-10">
                                    <div style="display: flex;align-items: center">
                                        <van-tag v-if="item.meetingType == 'audio'" type="primary">音频会议</van-tag>
                                        <van-tag v-else type="success">文档记录</van-tag>
                                        <div class="list_title">{{ item.meetingTitle }}</div>
                                    </div>
                                </div>
                                <van-divider/>
                                <!--          //display:flex;align-items: center;-->
                                <div v-for="i in item.meetingFile"
                                     style="margin-bottom: 10px;display:flex;align-items: center;">
                                    <van-image v-if="i.fileType=='word'" :src="word" height="23" width="23"></van-image>
                                    <van-image v-if="i.fileType=='xmind'" :src="xmind" height="23"
                                               width="23"></van-image>
                                    <van-image v-if="i.fileType=='mp3'" :src="mp3" height="23" width="23"></van-image>
                                    <div class="list_title1">{{ i.fileName }}</div>
                                    <!--        <div class="list_title1">{{item.meetingBrainMapName?item.meetingBrainMapName:''}}</div>-->
                                    <!--        <div class="list_title1">{{item.meetingMminutesName?item.meetingMminutesName:''}}</div>-->
                                </div>
                                <div>
                                    <div class="f-z-12 m-b-10">会议ID:{{ item.meetingId }}</div>
                                    <div class="f-z-12 m-b-10">创建时间：{{ item.createTime }}</div>

                                </div>
                            </div>
                            <!--          <businessCard :meetingList="meetingList" @changeCheck="changeCheck"></businessCard>-->
                        </van-list>
                        <van-row v-else-if="isShowImg" align="center" class="tc" justify="center">
                            <div>
                                <img src="../../assets/img/zanwupiaoju.png"
                                     style="height:50vw;width:80vw;object-fit: contain"/>
                                <div class="banner">暂无数据</div>
                            </div>
                        </van-row>
                    </van-pull-refresh>
                </van-tab>
            </div>
        </van-tabs>
    </div>
</template>
<script lang="ts" setup>
import hyTitle from '../../assets/img/hy_title.png'
import businessCard from '../../components/businessCard/index.vue'
import word from '../../assets/img/word.png'
import xmind from '../../assets/img/xmind.png'
import mp3 from '../../assets/img/mp3.png'

import businessIcon from '../../assets/img/business_icon.png';
import {ref, computed, onMounted, inject} from 'vue';
// import userRouter from 'user'
import {useRouter} from "vue-router";
import {getMeetingList} from "../../services/task-processing";
import {meetingData} from "../../store/index"
import {paramsTye} from "../../services/task-processing/types";

const checked = ref(false);
const isShowImg = ref<boolean>(false)  //数据为空的时候 控制图片的显示隐藏
const active = ref(0);
const router = useRouter()
const meetingStore = meetingData()
const meetingList = ref<meettingType[]>([])
const checkedItemsList = ref<any>()   //会议列表 勾选的数组
const loading = ref<Boolean>(false)
const finished = ref<Boolean>(false)
const finishText = ref<string>('没有更多了')
const isLoading = ref<boolean>(false)  //控制 下拉刷新
const tabIndex = ref(false)
const searchLoading = ref<Boolean>(false)  //控制 数据没请求回来的 loading
const isPull = ref<boolean>(false)
const isShowButton = ref<boolean>(false)
interface meettingType {
    createTime: string,
    meetingAddress: string,
    meetingAudioFileUrl: string,
    meetingBrainMapFileUrl: string,
    meetingCity: string,
    meetingFile: [{
        corpCode: string,
        corpName: string,
        createTime: string,
        creater: string,
        fileName: string,
        fileType: string,
        fileUrl: string,
        meetingId: string,
        meetingTitle: string
    }]
    meetingFlag: string,
    meetingId: string,
    meetingMminutesFileUrl: string,
    meetingPeople: string,
    meetingTime: string,
    meetingTitle: string
    meetingType: string
    taskId: string,
    isCheck: boolean
}

interface pageType {
    currentPage: number, //	当前页
    pageSize: number,//	Integer	是	每页显示条数
    // searchKey:string,//	否	筛选条件，暂时未用
    corpCode: string | null//	是	企业code
    userId: string	//是	用户id
    meetingFlag: string
}

let params = ref<pageType>(
    {
        currentPage: 1,
        pageSize: 10,
        corpCode: '',
        userId: '',
        meetingFlag: '0'
    }
)
const onClickLeft = () => {
    router.replace('/addTasks')
}

const changeCheck = (checkedItems: any) => {
    checkedItemsList.value = meetingList.value.filter((item: any) => item.isCheck);
    console.log('checkedItemsList.value', checkedItemsList.value)
}
// tab 切换
const changeList = async (name: number) => {
    if (name == 1) {
        params.value.currentPage = 1
        meetingList.value = []
        loading.value = false
        finished.value = false
        // isLoading.value = true
        isPull.value = false
        await getList()
        tabIndex.value = true
    } else {
        params.value.currentPage = 1
        meetingList.value = []
        loading.value = false
        // isLoading.value = true
        finished.value = false
        isPull.value = false
        await getList()
        tabIndex.value = false
    }
}
//点击了提交按钮
const handleSubmit = async () => {
    await meetingStore.setMeetingData(checkedItemsList.value)
    router.replace('/addTasks')
}
//触底函数
const totalCount = ref<number>(0)
const pageCount = ref<number>(0)
const getMeetingListLoad = async () => {
    if (meetingList.value.length < totalCount.value && params.value.currentPage < pageCount.value) {
        params.value.currentPage = params.value.currentPage + 1
        searchLoading.value = false
        setTimeout(() => {
            getList()
        }, 800)
    }
    // params.value.currentPage =  params.value.currentPage + 1
    //  await getList()
}
const getList = () => {
    params.value.meetingFlag = String(active.value)
    loading.value = true
    // searchLoading.value = true
    // if(isPull.value){
    //   console.log('刷新')
    //   searchLoading.value = false  //打开loading
    //   isLoading.value = true
    //   console.log(' searchLoading.value', searchLoading.value)
    // }else{
    //   console.log('不')
    //   searchLoading.value = true  //打开loading
    // }
    getMeetingList(params.value).then((res: any) => {
        if (res.success) {
            meetingList.value = meetingList.value.concat(res.data.MeetingList)
            //通过判断 控制 提交按钮的 显示隐藏
            if(meetingList.value .length>0){
                isShowButton.value = true
            }else{
                isShowButton.value = false
            }
            meetingList.value.forEach((item: any) => {
                item.isCheck = false
            })
            totalCount.value = res.data.totalCount
            pageCount.value = res.data.pageCount
            //反选功能（回显）
            let arr = meetingStore.getMeetingData();
            meetingList.value.forEach(item => {
                let index = arr.findIndex(meeting => item.meetingId === meeting.meetingId);

                if (index !== -1) {
                    // 如果在 arr 中找到相同的 meetingId，设置 isCheck 为 true
                    item.isCheck = true;
                }
            });
            //测试
            meetingList.value.forEach(item => {
                if (checkedItemsList.value) {
                    let index = checkedItemsList.value.findIndex(meeting => item.meetingId === meeting.meetingId);
                    if (index !== -1) {
                        // 如果在 arr 中找到相同的 meetingId，设置 isCheck 为 true
                        item.isCheck = true;
                    }
                }
            });
            setInterval(() => {
                isLoading.value = false
            }, 2000)
            // isLoading.value = false
            if (meetingList.value.length == totalCount.value) {
                loading.value = false   //关闭下拉刷新的加载
                finished.value = true
            } else {
                loading.value = false
                finished.value = false
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
        setTimeout(() => {
            isLoading.value = false
        }, 5000)
        // isLoading.value = false
    })
}
// 下拉刷新
const onRefresh = async () => {
    isLoading.value = true
    // isLoading.value = true
    isPull.value = true  //是刷新
    searchLoading.value = false
    loading.value = true
    finished.value = false
    params.value.currentPage = 1
    meetingList.value.length = 0
    await getList()
}
onMounted(async () => {
    let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
    params.value.corpCode = localStorage.getItem('corpCode') //从本地获取corpCode
    //从本地获取corpCode
    params.value.userId = userInfoData.userInfo.userId
    meetingList.value.length = 0
    params.value.currentPage = 1
    isPull.value = false
    await getList()
    // const selectedOptions = arr.map(meetingId => {
    //   return meetingList.value.find(item => item.meetingId === meetingId);
    // });
    // console.log(selectedOptions)
    // getMeetingList(params.value).then((res:any)=>{
    //   meetingList.value = res.data.MeetingList
    //   meetingList.value.forEach((item:any)=>{
    //     item.isCheck = false
    //   })
    // })
})
</script>
<style lang="less" scoped>
.listBox {
  height: 100vh;
}

.footer_btn {
  width: 96vw;
  padding: 0 2vw;
    margin-bottom: 10vw;
}

.h-67 {
  height: 20vh;
  align-items: center;
  display: flex;
  justify-content: center;
}

.list {
  height: calc(100vh - var(--van-nav-bar-height) - var(--van-tabs-line-height) - var(--van-button-default-height) - 10vw);
  overflow-y: auto;
  //max-height: 75vh;
  //overflow: auto;
  .list_card {
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
      margin-left: 2vw;
    }

    .f-z-12 {
      font-size: 0.75em;
      color: #4B5563;
    }

    .m-b-10 {
      margin-bottom: 2vw;
    }
  }
}

.listUse {
  height: calc(100vh - var(--van-nav-bar-height) - var(--van-tabs-line-height));
  overflow-y: auto;
  //max-height: 75vh;
  //overflow: auto;
  .list_card {
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
      margin-left: 2vw;
    }

    .f-z-12 {
      font-size: 0.75em;
      color: #4B5563;
    }

    .m-b-10 {
      margin-bottom: 2vw;
    }
  }
}

.list1 {
  display: none;
}

.list_title1 {
  font-size: 0.75em;
  //font-weight: 550;
  margin-left: 2vw;
}

.mainButton {
  border-radius: 10px;
  padding: 2vw;
  width: 42vw;
  height: 50px;
  background: url('../../assets/img/radio.png') no-repeat center;
  background-size: 100vw 10vh;
  //background-image: linear-gradient(to right bottom, #0080FF, #6AB5FF);
}

.mainButton_sec {
  border-radius: 10px;
  padding: 2vw;
  width: 42vw;
  height: 50px;
  background: url('../../assets/img/radio.png') no-repeat center;
  background-size: 100vw 10vh;
  //background-image: linear-gradient(to right bottom, #33B2AD, #3BD7D3);
}

:deep(.van-dropdown-menu__bar) {
  box-shadow: none;
}

:deep(.van-divider) {
  margin: 2vw;
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
