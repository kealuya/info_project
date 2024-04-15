<template>
  <taskCard :title="title" :taskTitle="taskTitle" :picSrc="picSrc" :target="target" :content="content" :isComplete="isComplete"></taskCard>
<!--  <van-nav-bar-->
<!--      title="任务详情"-->
<!--      left-text="返回"-->
<!--      left-arrow-->
<!--      @click-left="onClickLeft"-->
<!--  />-->
<!--    <div class="task  f_white">-->
<!--&lt;!&ndash;      <div class="task_card">&ndash;&gt;-->
<!--&lt;!&ndash;        <div>&ndash;&gt;-->
<!--&lt;!&ndash;          医患诊断采集&ndash;&gt;-->
<!--&lt;!&ndash;        </div>&ndash;&gt;-->
<!--&lt;!&ndash;        <div class="content-font">&ndash;&gt;-->
<!--&lt;!&ndash;          医患诊断采集是医疗过程中非常重要的一环，准确的信息采集和综合分析有助于医生做出正确的诊断和治疗决策，提高患者的治疗效果和生活质量。&ndash;&gt;-->
<!--&lt;!&ndash;        </div>&ndash;&gt;-->
<!--&lt;!&ndash;      </div>&ndash;&gt;-->
<!--&lt;!&ndash;      <div class="btn">&ndash;&gt;-->
<!--&lt;!&ndash;        立即参与&ndash;&gt;-->
<!--&lt;!&ndash;      </div>&ndash;&gt;-->
<!--    </div>-->
<!--    <div class="content">-->
<!--      <div class="title_one m-b-10" >新药上市情况跟踪</div>-->
<!--      <div class="title_two f-z-14 f-w-550  m-b-10">-->
<!--        <span>任务目标</span>-->
<!--&lt;!&ndash;        <span class="m-l">发布日期</span>&ndash;&gt;-->
<!--      </div>-->
<!--      <div class="f-z-14  m-b-10">-->
<!--        及时获取同行业新上市药品的信息。-->
<!--      </div>-->
<!--      <div class="title_two f-z-14 f-w-550  m-b-10">-->
<!--        <span>任务内容</span>-->
<!--      </div>-->
<!--      <div class="f-z-14 m-b-10">-->
<!--        关注行业内新药上市动态,及时采集新药的适应症、功能主治、使用说明、销售情况等资料。-->
<!--      </div>-->
<!--&lt;!&ndash;      <div class="f-z-14 m-b-10">&ndash;&gt;-->
<!--&lt;!&ndash;        病史采集：医生会询问患者有关疾病发生的时间、症状的持续时间、疼痛的性质、病史（包括既往病史、家族史等）等信息。患者提供详细的病史对于正确诊断疾病非常重要。&ndash;&gt;-->
<!--&lt;!&ndash;      </div>&ndash;&gt;-->
<!--&lt;!&ndash;      <div class="f-z-14 m-b-10">&ndash;&gt;-->
<!--&lt;!&ndash;        症状采集：医生会询问患者当前的症状，比如疼痛、咳嗽、发热等。患者描述症状的特点和变化有助于医生判断疾病的可能性。&ndash;&gt;-->
<!--&lt;!&ndash;      </div>&ndash;&gt;-->
<!--&lt;!&ndash;      <div class="f-z-14 m-b-10">&ndash;&gt;-->
<!--&lt;!&ndash;        体征检查：医生会进行体格检查，包括测量体温、血压、心率、听诊心肺等。体征检查可以提供重要的诊断线索。&ndash;&gt;-->
<!--&lt;!&ndash;      </div>&ndash;&gt;-->
<!--&lt;!&ndash;      <div class="f-z-14 m-b-10">&ndash;&gt;-->
<!--&lt;!&ndash;        实验室检查：医生可能会要求患者进行实验室检查，包括血液检查、尿液检查、影像学检查（如X光、CT、MRI等），以获取更多的诊断信息。&ndash;&gt;-->
<!--&lt;!&ndash;      </div>&ndash;&gt;-->
<!--&lt;!&ndash;      <div class="f-z-14 m-b-10">&ndash;&gt;-->
<!--&lt;!&ndash;        诊断过程：医生根据患者提供的信息、体征和检查结果，进行综合分析和判断，最终做出诊断并制定治疗方案。&ndash;&gt;-->
<!--&lt;!&ndash;      </div>&ndash;&gt;-->
<!--&lt;!&ndash;      <div class="f-z-14 m-b-10">&ndash;&gt;-->
<!--&lt;!&ndash;        医患诊断采集是医疗过程中非常重要的一环，准确的信息采集和综合分析有助于医生做出正确的诊断和治疗决策，提高患者的治疗效果和生活质量。在现代医疗中，随着信息技术的发展，电子病历系统等工具也广泛应用于医患诊断采集过程，提高了信息的准确性和可追溯性。&ndash;&gt;-->
<!--&lt;!&ndash;      </div>&ndash;&gt;-->
<!--    </div>-->
<!--    <div class="footer_btn">-->
<!--      <van-button block type="primary" @click="taskProcessingHandle">-->
<!--        立即参与-->
<!--      </van-button>-->
<!--    </div>-->
<!--  <div class="box"></div>-->
</template>
<script lang="ts" setup>
import {useRouter} from "vue-router";
// import { showSuccessToast, showFailToast } from 'vant';
import {showSuccessToast} from "vant";
import bgc from "../../assets/img/home_04.png";
import {ref} from "vue";
const router = useRouter()
const isComplete = ref<Boolean>(false)
const title = ref<string>('任务详情')
const taskTitle = ref<string>('内容营销长期策略')
const picSrc = ref<string>(bgc)
const target = ref<string>('持续发布新内容可以吸引用户回访，保持用户对品牌的关注度。不断推出有价值的内容可以加强品牌在目标受众中的认知和印象。')
const content = ref<string>('随着时间推移，行业知识和信息可能会发生变化，持续更新内容可确保信息的准确性和有效性。\n' +
    '因此，持续更新和改进内容是内容营销成功的关键，有助于建立稳固的品牌声誉、吸引用户并保持用户忠诚度。\n')
const onClickLeft = () => {
  // history.go(-2)
  router.replace('/homeNew')
}
const taskProcessingHandle = ()=>{
  showSuccessToast('参与成功');
  setTimeout(() => {
    router.replace('/homeNew')
  }, 2000);
}
</script>
<style lang="less" scoped>
.task{
  //margin-top: 2vw;
  //border-radius: 10px;
  //padding: 0 2vw;
  width:96vw;
  margin: 2vw;
  height: 160px;
  background: url('../../assets/img/home_04.png') no-repeat center;
  background-size: cover;
  //box-sizing: border-box;
  //margin-left: 2vw;
  .task_card{
    width: 62vw;
    background-color: #e9f8f1;
    opacity: 0.5; /* 设置透明度为50% */
    color:#000000;
    margin: 2vw;
    padding:2vw;
    font-size: 1em;
    border-radius: 10px;
    .content-font{
      font-size: 0.75em;
      margin-top:1vw;
    }
  }
  .btn{
    border-radius: 10px;
    background-color: #253334;
    width:42vw;
    margin-left: 2vw;
    padding:1vw;
    text-align: center;
    font-size: 1em;
  }
}
.box{
  height:50px;
}
.content{
  margin:2vw;
  line-height:3vh;
.title_one{
  text-align:center;
  font-weight: 550;
  font-size: 0.875em;
}
 .f-z-14{
   font-size: 0.75em;
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
}
.footer_btn{
  width:92vw;
  position: fixed;
  bottom:60px;
  padding:2vw;
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
<script setup lang="ts">
</script>