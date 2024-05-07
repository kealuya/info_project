<template>
  <van-nav-bar
      title="业务详情"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
  />
  <div class="content">
    <div class="title_one m-b-10 f-z-14" >{{meetingTitle}}</div>
    <div class="title_two f-z-14 f-w-550  m-b-10">
      <span>文档摘要</span>
    </div>
    <div class="f-z-12 m-b-10">
      &nbsp;&nbsp;互联网营销策略包括多种方法和技术，旨在利用互联网平台、工具和资源来推广产品或服务、增加品牌曝光、吸引潜在客户并提高销售业绩。
    </div>
    <div class="f-z-12 m-b-10">
      &nbsp; 网站优化（SEO）： 通过优化网站内容和结构，提高在搜索引擎中的排名，增加有机流量
      搜索引擎营销（SEM）： 利用付费广告在搜索引擎上展示，提高网站访问量和曝光度
    </div>
    <!--    <div class="f-z-12 m-b-10">-->
    <!--      &nbsp;&nbsp;症状采集：医生会询问患者当前的症状，比如疼痛、咳嗽、发热等。患者描述症状的特点和变化有助于医生判断疾病的可能性。-->
    <!--    </div>-->
    <!--    <div class="f-z-12 m-b-10">-->
    <!--      &nbsp;&nbsp;体征检查：医生会进行体格检查，包括测量体温、血压、心率、听诊心肺等。体征检查可以提供重要的诊断线索。-->
    <!--    </div>-->
    <!--    <div class="f-z-12 m-b-10">-->
    <!--      &nbsp;&nbsp; 实验室检查：医生可能会要求患者进行实验室检查，包括血液检查、尿液检查、影像学检查（如X光、CT、MRI等），以获取更多的诊断信息。-->
    <!--    </div>-->
    <!--    <div class="f-z-12 m-b-10">-->
    <!--      &nbsp;&nbsp;诊断过程：医生根据患者提供的信息、体征和检查结果，进行综合分析和判断，最终做出诊断并制定治疗方案。-->
    <!--    </div>-->
    <div class="f-z-14  m-b-10 space">
      <span class="f-w-550">文档记录</span>
      <div class="right">
        <van-image :src="Ai"  width="20" height="20">
        </van-image>
        <span class="c-blue">生成</span>
      </div>
    </div>
    <div class="f-z-12  m-b-10 space">
      <span>会议纪要</span>
    </div>
    <van-text-ellipsis
        rows="5"
        :content="text"
        expand-text="展开纪要"
        collapse-text="收起纪要"
    />
    <div class="metting">
      <van-image :src="Word"  width="23" height="23">
      </van-image>
      <div class="c-blue f-z-14">20240408-探索互联网销售的未来.word <span>&nbsp;下载附件</span></div>
    </div>
<!--    <div class="f-z-14  m-b-10 space">-->
<!--      <span class="f-w-550">脑图</span>-->
<!--      <div class="right">-->
<!--        <van-image :src="Ai"  width="23" height="23">-->
<!--        </van-image>-->
<!--        <span class="c-blue">生成</span>-->
<!--      </div>-->
<!--    </div>-->
<!--    <svg class="flex-1" ref="svgRef" />-->
    <div class="metting">
      <van-image :src="Xmind"  width="23" height="23">
      </van-image>
      <div class="c-blue f-z-14"> &nbsp;&nbsp;20240408-探索互联网销售的未来.xmind<span>&nbsp;下载附件</span></div>
    </div>
    <div class="title_two f-z-14 f-w-550  m-b-10">
      <span>会议信息</span>
    </div>
    <div class="f-z-12">会议地址：{{meetingCity}}{{meetingAddress}}</div>
    <div class="f-z-12">会议时长：1小时53分钟</div>
  </div>
  <div class="box"></div>
</template>
<script setup lang="ts">
import Ai from '../../assets/img/ai.png';
import {getMeetingDetails} from '../../services/task-processing/index'
import Word from '../../assets/img/word.png'
import Xmind from '../../assets/img/xmind.png'
import {useRouter,useRoute} from "vue-router";
import {onMounted,ref} from "vue";
const router = useRouter()
const route = useRoute()
const meetingTitle = ref<string>()
const meetingCity = ref<string>()
const meetingAddress = ref<string>()
const onClickLeft = () => {
  router.push('/business')
}
onMounted(()=>{
// console.log(route.query.meetingId)
  const id = route.query.meetingId
  getMeetingDetails({MeetingId:id}).then((res:any)=>{
    if(res.success){
      meetingTitle.value = res.data.meetingTitle
      meetingCity.value = res.data.meetingCity
      meetingAddress.value = res.data.meetingAddress
    }
  })
})
</script>
<style lang="less" scoped>
.content{
  margin:2vw;
  .title_one{
    text-align:center;
    font-weight: bold;
  }
  .f-z-14{
    font-size: 0.875em;
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
    margin-bottom:20px;
  }
  .space{
    display:flex;
    justify-content: space-between;
  }
  .metting{
    display:flex;
    align-items: center;
    padding: 3vw 0;
  }
  .right{
    display: flex;
    vertical-align: middle;
  }
  .c-blue{
    color:#000000;
    margin-left: 1vw;
    font-size: 0.78em;
    span{
      color: var(--van-text-ellipsis-action-color);
    }
  }

}
.flex-1{
  width: 100%;
}
.box{
  height:10vh;
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
:deep(.van-text-ellipsis){
  font-size: 0.75em;
  color:#000000;
}
.f-z-12{
  font-size: 0.75em;
}
.f-z-14{
  font-size: 0.875em;
}

</style>
