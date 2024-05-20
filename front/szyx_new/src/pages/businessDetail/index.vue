<template>
    <van-nav-bar
            left-arrow
            left-text="返回"
            title="业务详情"
            @click-left="onClickLeft"
    />

    <div class="content">
        <div class="card-item">
            <div class="m-b-10" style="display:flex;align-items: center">
                <van-image :src="meetingPic" class="m-r-2" fit="contain" height="30"
                           width="30"></van-image>
              <span class="f-z-16">会议详情内容</span>
            </div>
            <div class="item-card">
                <span class="f-z-14-c">会议ID</span>
                <div class="f-z-14-black">{{ detailInfo.meetingId ? detailInfo.meetingId : '--' }}</div>
            </div>
            <div class="item-card">
                <span class="f-z-14-c">会议标题</span>
                <div class="f-z-14-black">{{ detailInfo.meetingTitle ? detailInfo.meetingTitle : '--' }}</div>
            </div>
            <div class="item-card">
                <span class="f-z-14-c">会议时间</span>
                <div class="f-z-14-black">{{ detailInfo.createTime ? detailInfo.createTime : '--' }}</div>
            </div>
            <div class="item-card">
                <span class="f-z-14-c">会议地址</span>
                <div class="f-z-14-black">{{ detailInfo.meetingCity}}-{{detailInfo.meetingAddress}}</div>
            </div>
            <div class="item-card">
                <span class="f-z-14-c">参会人员</span>
                <div class="f-z-14-black">{{ detailInfo.meetingPeople ? detailInfo.meetingPeople : '--' }}</div>
            </div>
        </div>
        <!--     //会议音频-->
        <div class="card-item">
            <div class="title_two m-b-10 space">
                <span class="m-r-2 f-z-16">{{ detailInfo.meetingType == 'document' ? '文档' : '会议' }}音频</span>
                <div class="file-box-right" v-if="detailInfo.translationState=='0'">
                    <van-image :src="progress" width="25" fit="contain" class="audioPic"></van-image>
                    <div class="f-z-12 f-w-550 textTag">音频转译中</div>
                </div>
                <div class="file-box-right" v-if="detailInfo.translationState=='1'">
                    <van-image :src="progressSuccess" width="25" fit="contain" class="audioPic"></van-image>
                    <div class="f-z-12 f-w-550 textTagSuccess">音频转译成功</div>
                </div>
                <div class="file-box-right" v-if="detailInfo.translationState=='2'">
                    <van-image :src="progressFail" width="25" fit="contain" class="audioPic"></van-image>
                    <div class="f-z-12 f-w-550 textTagFail">音频转译失败</div>
                </div>
            </div>
<!--            //音频文件部分-->

                <div class="file-box-left">
                    <div v-for="item in detailInfo.meetingFile" class="file">
                        <div>
                            <van-image :src="Mp3" height="23" width="23" v-if="item.fileType=='mp3'"></van-image>
                            <van-image :src="Word" height="23" width="23" v-if="item.fileType=='word'"></van-image>
                            <van-image :src="Xmind" height="23" width="23" v-if="item.fileType=='xmind'"></van-image>
                            <span class="fileName">{{ item.fileName }}</span>

                        </div>
                        <div @click="downloadFile(item.fileUrl)" class="fileName f-z-14-c down"><van-image :src="downIcon" width="20"></van-image>下载附件</div>
                    </div>

            </div>
        </div>
        <!--      //摘要-->
        <div class="card-item">
            <div class="title_two m-b-10">
              <div class="title-ai">
                  <div class="f-z-16">{{ detailInfo.meetingType == 'document' ? '文档' : '会议' }}摘要</div>
                  <div class="file-box-right" v-if="detailInfo.meetingSummaryState=='1'">
                      <van-image :src="Abstract" width="25" fit="contain" height="25"></van-image>
                      <div class="textTag f-w-550 f-z-12">该内容Ai生成</div>
                  </div>
              </div>
                <!--                //音频转译完成之后 才显示 摘要以及 按钮-->
                <div v-if="detailInfo.meetingSummaryState == '0'">
                    <div class="center-content">
                        <van-image :src="Generating" width="40"></van-image>
                        <div class="text">会议摘要生成中...</div>
                    </div>
                    <div class="generating">
                        <van-button type="primary" size="small" @click="refreshResults">刷新结果</van-button>
                    </div>

                </div>
                <div v-if="detailInfo.translationState!='1'" class="generating">
                    <div class="center-content">
                        <van-image :src="zanwupiaoju" width="150"></van-image>
                        <div class="text">暂无摘要</div>
                    </div>
                </div>
                <!--                //音频转译成功 显示-->
                <div class="generate" @click="aiHandle" v-if="detailInfo.translationState=='1'&& detailInfo.meetingSummaryState==''">
                    <van-image :src="Ai" width="30"></van-image>
                    生成会议摘要
                </div>
            </div>
<!--            //显示会议摘要的内容-->
            <div v-if="detailInfo.meetingSummaryState=='1'&&detailInfo.translationState=='1'" class="f-z-14 m-b-10">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{{ detailInfo.meetingSummary }}
            </div>
        </div>
<!--        //此处是 生成会议纪要 和会议脑图的 按钮-->
            <div class="generate" @click="mapHandle" v-if="detailInfo.meetingSummaryState=='1'&& detailInfo.meetingBrainMap ==''&& detailInfo.meetingBrainMap==''">
                <van-image :src="Ai" width="30"></van-image>
                生成纪要、脑图
        </div>
        <!--      //会议纪要-->
        <div class="card-item">
            <div class="m-b-10 space">
                <div class="f-z-16">{{ detailInfo.meetingType == 'document' ? '文档' : '会议' }}纪要</div>
                <div class="file-box-right" v-if="detailInfo.meetingSummaryState=='1'">
                    <van-image :src="Abstract" width="25" fit="contain" height="25"></van-image>
                    <div class="textTag f-w-550 f-z-12">该内容Ai生成</div>
                </div>
            </div>
            <div v-if="searchLoading" class="h-67">
                <van-loading type="spinner" color="#1989fa" :vertical="true">正在获取会议纪要、脑图...</van-loading>
            </div>
            <van-text-ellipsis
                v-if="detailInfo.meetingMinutes"
                   class="f-z-14"
                    :content="detailInfo.meetingMinutes"
                    collapse-text="收起纪要"
                    expand-text="展开纪要"
                    rows="10"
            />
            <div class="generating" v-else>
                <div class="center-content">
                    <van-image :src="zanwupiaoju" width="150"></van-image>
                    <div class="text">暂无纪要</div>
                </div>
            </div>
        </div>
        <!--        会议脑图-->
        <div  class="card-item">
            <div class="m-b-10 space">
                <div class="f-z-16">{{ meetingType == 'document' ? '文档' : '会议' }}脑图</div>
                <div class="file-box-right" v-if="detailInfo.meetingSummaryState=='1'">
                    <van-image :src="Abstract" width="25" fit="contain" height="25"></van-image>
                    <div class="textTag f-w-550 f-z-12">该内容Ai生成</div>
                </div>
            </div>
            <div v-if="detailInfo.meetingBrainMap==''" class="generating">
                <div class="center-content">
                    <van-image :src="zanwupiaoju" width="150"></van-image>
                    <div class="text">暂无摘要</div>
                </div>
            </div>
            <svg ref="mapRef" class="flex-1"/>
        </div>
    </div>
    <div class="box"></div>
</template>
<script lang="ts" setup>
import meetingPic from '../../assets/img/meeting.png'
import Ai from '../../assets/img/ai.png';
import zanwupiaoju from '../../assets/img/zanwupiaoju.png'
import Generating from '../../assets/img/generating.png'
import Word from '../../assets/img/word.png'
import Xmind from '../../assets/img/xmind.png'
import Abstract from '../../assets/img/abstract-ai.png'
import Mp3 from '../../assets/img/mp3.png'
import progress from '../../assets/img/progress.png'
import progressSuccess from '../../assets/img/progress-success.png'
import progressFail from '../../assets/img/progress-fail.png'
import downIcon from '../../assets/img/down-icon.png'
import {useRouter, useRoute} from "vue-router";
import {getMeetingDetails, audioMeeting, audioMeeting_Ai_Summary_BrainMap} from '../../services/task-processing/index'
import {inject, onMounted, ref} from "vue";
import finish from "../../assets/img/finish.png";
import {Markmap} from "markmap-view";
import {Transformer} from 'markmap-lib';
import {IMeetingDetailTypeResponse, ResponseType} from "../../services/task-processing/types";
import {closeToast, showFailToast, showLoadingToast, showSuccessToast} from "vant";
const isMap = ref<boolean>(false)
const isGenerating = ref<boolean>(false)
const searchLoading = ref<boolean>(false)
const detailInfo = ref<IMeetingDetailTypeResponse>({
    bz1: "",
    bz2: "",
    bz3: "",
    corpCode: "",
    corpName: "",
    createTime: "",
    creater: "",
    meetingAddress: "",
    meetingAudioText: "",
    meetingBrainMap: "",
    meetingCity: "",
    meetingFile: [],
    meetingFileUrl: "",
    meetingFlag: "",
    meetingId: "",
    meetingMinutes: "",
    meetingPeople: "",
    meetingSummary: "",
    meetingSummaryState: "",
    meetingTime: "",
    meetingTitle: "",
    meetingType: "",
    taskId: "",
    translationState: ""
})
const router = useRouter()
const route = useRoute()
let meetingAudioMinutes = ref()
let meetingAudioSummary = ref()
let wordFiles = ref()
let xmindFiles = ref()
let meetingType = ref()
const createTime = ref()
const isShow = ref<boolean>(false)
const isShowPic = ref<boolean>(false)
const meetingId = ref<string | undefined>()
const userId = ref<string | undefined>()
const isShowAbstract = ref<boolean>(false)

const onClickLeft = () => {
    router.push('/business')
}

//调用生成的方法
interface audioType {
    data: any,
    msg: string,
    success: boolean
}

//点击了 摘要 生成 按钮
const aiHandle = () => {
    let audioParams = {
        meetingId: meetingId.value
    }
    showLoadingToast({
        message: '加载中...',
        forbidClick: true,
    });
    audioMeeting(audioParams).then(async res => {
        const response: audioType = res as audioType; // 显式类型转换
        if (response.success) {
            closeToast()
            isShowPic.value = true
            showSuccessToast(response.msg)
            await getMeetingDetailsList()
            // location.reload()


        } else {
            isShowPic.value = false
            showFailToast(response.msg)
        }
    })
}
const refreshResults = async()=>{
    await getMeetingDetailsList()
}
//下载附件 功能
const downloadFile = async (value:string) => {  //参数接受的是文件地址
    // console.log(value)
    window.location.href = value; //
};
//生成会议、脑图的函数
const mapHandle = ()=>{
    let params = {
        meetingId:meetingId.value
    }
    searchLoading.value = true
    audioMeeting_Ai_Summary_BrainMap(params).then(async res => {
        const response: audioType = res as audioType; // 显式类型转换
        if (response.success) {
            searchLoading.value = false
            // closeToast()
            showSuccessToast(response.msg)
            await getMeetingDetailsList()
            // location.reload()
        } else {
            showFailToast(response.msg)
        }
    })
}
let initValue = ''
const getMeetingDetailsList = async () => {
    let params = {
        meetingId: meetingId.value,
        creater: userId.value
    }
    // const id = route.query.meetingId
    await getMeetingDetails(params).then((res:any) => {

        if (res.success) {
            console.log(111)
            detailInfo.value = {...res.data}
            if(detailInfo.value.translationState =='1' && detailInfo.value.meetingSummaryState=='1'){
                isShow.value = true
                isMap.value = true
                initValue = detailInfo.value.meetingBrainMap
                mapValue.value = Markmap.create(mapRef.value) //有用
                update()//有用
            }
            // initValue = ''

            // if (detailInfo.value.meetingSummaryState == '0') {   //说明  正在 ai 转译中  显示 正在生成中  请稍后
            //     // console.log(2)
            //     isGenerating.value = true
            //
            //
            // } else if (detailInfo.value.meetingSummaryState == '1') {
            //     isMap.value = true   //脑图
            // }else{
            //     isGenerating.value = false
            //
            // }
            // if(detailInfo.value.translationState =='0'){
            //     isShowAbstract.value = true
            //     isGenerating.value = false
            // }else {
            //     isShowAbstract.value = false
            // }
            //
            // if(detailInfo.value.translationState =='1'){
            //     isGenerating.value = false
            // }else{
            //     isGenerating.value = true
            // }
            // if(detailInfo.value.translationState =='1' && detailInfo.value.meetingSummaryState=='1'){
            //     isShow.value = true
            //     isMap.value = true
            //     initValue = detailInfo.value.meetingBrainMap
            //     mapValue.value = Markmap.create(mapRef.value) //有用
            //     update()//有用
            // }
        }
    })

}
const mapRef = ref()
const mapValue = ref<any>()
const transformer = new Transformer()
//生成 脑图的方法
const update = () => {
    const {root} = transformer.transform(initValue);
    mapValue.value.setData(root);
    mapValue.value.fit();
}

onMounted(async () => {
    // mapValue.value = Markmap.create(mapRef.value) //有用
    // update()//有用
    let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
    userId.value = userInfoData.userInfo.userId
    meetingId.value = route.query.meetingId as string
    await getMeetingDetailsList()
})
</script>

<style lang="less" scoped>
.content {
  margin: 2vw;

  .title_one {
    text-align: center;
    font-weight: bold;
  }

  .f-z-14 {
    font-size: 0.875em;
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

  .m-b-10 {
    margin-bottom: 10px;
  }

  .space {
    display: flex;
    justify-content: space-between;
  }

  .metting {
    display: flex;
    align-items: center;
    padding: 3vw 0;
  }

  .right {
    display: flex;
    vertical-align: middle;
  }

  .c-blue {
    color: #000000;
    margin-left: 1vw;
    font-size: 0.78em;

    span {
      color: var(--van-text-ellipsis-action-color);
    }
  }

  .a-center {
    display: flex;
      flex-direction: column;
    //align-items: center;
  }

  .generate {
    background-color: #0088ff;
    color: #FFFFFF;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 0.875em;
    padding: 2vw 0;
    margin-top: 4vw;
    border-radius: 4px;
  }

  .generating {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 4vw;
  }

  .center-content {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .text {
    margin-top: 10px;
    color: #A4A4A4;
  }

}
.h-67{
    //height:67vh;
    z-index:999;
}
.flex-1 {
  width: 100%;
}
.title-ai{
    display: flex;
    justify-content: space-between;
    align-items: center;
}

::v-deep(.van-nav-bar__content) {
  background-color: #0088ff;
}

.card-item {
  padding: 2vw;
  background-color: #fff;
  border-radius: 10px;
  width: 92vw;
  margin: 2vw 0;

}
    .file-box-left{
        display:flex;
        flex-direction:column;
        padding: 0 2vw;
        .file {
            display: flex;
            justify-content: space-between;
            margin-bottom: 2vw;
        }
}
.file-box-right{
    display: flex;
    justify-content: center;
    align-items: center;
    color: #0088ff;
    .textTag{
        margin-left: 2vw;
    }
    .textTagSuccess{
        color:#16cb8d;
        margin-left: 2vw;
    }
    .textTagFail{
        color:#ff5b5b;
        margin-left: 2vw;
    }
    .audioPic{
        vertical-align: middle;
    }
}
.box{
    height: 10vh;
}
.fileName{
    margin-left: 1vw;
    font-size: 0.875em;
}
.down{
    display: flex;
    align-items: center;
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
::v-deep(.van-image){
    vertical-align: middle;
}
:deep(.van-text-ellipsis) {
  font-size: 0.75em;
  color: #000000;
}
.theme-color{
    color: #4BA3FB;
}

.f-z-12 {
  font-size: 0.75em;
}
.f-z-16{
    font-size: 1em;
    color: #000000;
}
.f-z-14 {
  font-size: 0.875em;
}

.m-r-2 {
  margin-right: 2vw;
}

.f-z-14-c {
  font-size: 0.875em;
  color: rgba(23, 26, 29, 0.4);
  //font-weight: 550;
}

.f-z-14-black {
  font-size: 0.875em;
  color: #000000;
}

.item-card {
  padding-bottom: 2vw;
}

.m-b-5 {
  margin-bottom: 2vw;
}

</style>
