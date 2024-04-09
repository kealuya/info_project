<template>
  <van-nav-bar
      title="业务详情"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
  />
  <div class="content">
    <div class="title_one m-b-10" >聚醚酮酮数字化乳牙早失间隙保持器的临床效果</div>
    <div class="title_two f-z-14 f-w-550  m-b-10">
      <span>会议摘要</span>
    </div>
    <div class="f-z-14 m-b-10">
      医患诊断采集是指医生或医疗机构在诊疗过程中收集患者的相关信息，包括病史、症状、体征、检查结果等，以便进行正确的诊断和治疗。这个过程通常包括以下几个方面：
    </div>
    <div class="f-z-14 m-b-10">
      病史采集：医生会询问患者有关疾病发生的时间、症状的持续时间、疼痛的性质、病史（包括既往病史、家族史等）等信息。患者提供详细的病史对于正确诊断疾病非常重要。
    </div>
    <div class="f-z-14 m-b-10">
      症状采集：医生会询问患者当前的症状，比如疼痛、咳嗽、发热等。患者描述症状的特点和变化有助于医生判断疾病的可能性。
    </div>
    <div class="f-z-14 m-b-10">
      体征检查：医生会进行体格检查，包括测量体温、血压、心率、听诊心肺等。体征检查可以提供重要的诊断线索。
    </div>
    <div class="f-z-14 m-b-10">
      实验室检查：医生可能会要求患者进行实验室检查，包括血液检查、尿液检查、影像学检查（如X光、CT、MRI等），以获取更多的诊断信息。
    </div>
    <div class="f-z-14 m-b-10">
      诊断过程：医生根据患者提供的信息、体征和检查结果，进行综合分析和判断，最终做出诊断并制定治疗方案。
    </div>
    <div class="f-z-14  m-b-10 space">
      <span class="f-w-550">会议记录</span>
      <div class="right">
        <van-image :src="Ai"  width="20" height="20">
        </van-image>
        <span class="c-blue">生成</span>
      </div>
    </div>
    <div class="metting">
      <van-image :src="Word"  width="30" height="30">
      </van-image>
      <div class="c-blue f-z-14">20240408-聚醚酮酮数字化临床效果会议.word</div>
    </div>
    <div class="f-z-14  m-b-10 space">
      <span class="f-w-550">会议脑图</span>
      <div class="right">
        <van-image :src="Ai"  width="20" height="20">
        </van-image>
        <span class="c-blue">生成</span>
      </div>
    </div>
    <van-image :src="naotu">
    </van-image>
    <div class="metting">
      <van-image :src="Xmind"  width="30" height="30">
      </van-image>
      <div class="c-blue f-z-14">20240408-聚醚酮酮数字化临床效果会议.xmind</div>
    </div>
    <div class="title_two f-z-14 f-w-550  m-b-10">
      <span>会议信息</span>
    </div>
    <div class="f-z-14">会议地址：天津市南开区红旗南路一中心总医院B座</div>
    <div class="f-z-14">会议时长：1小时53分钟</div>
  </div>
  <div class="box"></div>
</template>
<script setup lang="ts">
import naotu from '../../assets/img/naotu.png'
import {useRouter} from "vue-router";
import Ai from '../../assets/img/ai.png'

import Word from '../../assets/img/word.png'
import Xmind from '../../assets/img/xmind.png'

const router = useRouter()
const onClickLeft = () => {
router.push('/business')
}
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
    margin-bottom:10px;
  }
  .space{
    display:flex;
    justify-content: space-between;
  }
  .metting{
    display:flex;
    align-items: center;
    margin-bottom: 20px;
  }
  .right{
    display: flex;
    vertical-align: middle;
  }
  .c-blue{
    color:#1989fa;
  }

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
</style>
