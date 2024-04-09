<template>
  <van-nav-bar
      title="任务详情"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
  />
  <div class="task  f_white">
    <!--      <div class="task_card">-->
    <!--        <div>-->
    <!--          医患诊断采集-->
    <!--        </div>-->
    <!--        <div class="content-font">-->
    <!--          医患诊断采集是医疗过程中非常重要的一环，准确的信息采集和综合分析有助于医生做出正确的诊断和治疗决策，提高患者的治疗效果和生活质量。-->
    <!--        </div>-->
    <!--      </div>-->
    <!--      <div class="btn">-->
    <!--        立即参与-->
    <!--      </div>-->
  </div>
  <div class="content">
    <div class="title_one m-b-10" >生物标志物作为高血压患者风险评估的有效指标</div>
    <div class="title_two f-z-14 f-w-550  m-b-10">
      <span>任务目标</span>
      <!--        <span class="m-l">发布日期</span>-->
    </div>
    <div class="f-z-14  m-b-10">
      &nbsp; &nbsp;间隙保持器有助于维持乳牙缺失后的间隙空间安装间隙保持架后，间隙可以保持在一定程度上。恒磨牙萌出后，取下间隙固位体，最大限度地保持牙齿的咬合平衡。
    </div>
    <div class="title_two f-z-14 f-w-550  m-b-10">
      <span>任务内容</span>
    </div>
    <div class="f-z-14 m-b-10">
      &nbsp; &nbsp;探讨聚醚酮酮(PEKK)数字化乳牙早失间隙保持器的临床应用效果。方法从2019年4—5月在南京大学医学院附属口腔医院儿童口腔科就诊的6~10岁儿童中选取下颌单侧乳磨牙早失患者共5例,其中男2例、女3例,经数字化口内光学印模制取,间隙保持器数字化设计,PEKK基材间隙保持器的数字化制作,口内试戴调整,粘接。结果5例患儿的数字化乳牙早失间隙保持器重量只有传统金属带环丝圈保持器的25%,设计和制作时间只需1 h,5例数字化间隙保持器口内观察与基牙吻合度高。结论数字化乳牙早失间隙保持器采用数字化口内印模法比传统藻酸盐印模法提高了患儿的舒适度,可当日设计完成并佩戴,并使用PEKK材料,实现了间隙保持器去金属化,无金属过敏等风险,兼顾了美观的效果,具有较好的临床应用前景。
    </div>
    <!--      <div class="f-z-14 m-b-10">-->
    <!--        病史采集：医生会询问患者有关疾病发生的时间、症状的持续时间、疼痛的性质、病史（包括既往病史、家族史等）等信息。患者提供详细的病史对于正确诊断疾病非常重要。-->
    <!--      </div>-->
    <!--      <div class="f-z-14 m-b-10">-->
    <!--        症状采集：医生会询问患者当前的症状，比如疼痛、咳嗽、发热等。患者描述症状的特点和变化有助于医生判断疾病的可能性。-->
    <!--      </div>-->
    <!--      <div class="f-z-14 m-b-10">-->
    <!--        体征检查：医生会进行体格检查，包括测量体温、血压、心率、听诊心肺等。体征检查可以提供重要的诊断线索。-->
    <!--      </div>-->
    <!--      <div class="f-z-14 m-b-10">-->
    <!--        实验室检查：医生可能会要求患者进行实验室检查，包括血液检查、尿液检查、影像学检查（如X光、CT、MRI等），以获取更多的诊断信息。-->
    <!--      </div>-->
    <!--      <div class="f-z-14 m-b-10">-->
    <!--        诊断过程：医生根据患者提供的信息、体征和检查结果，进行综合分析和判断，最终做出诊断并制定治疗方案。-->
    <!--      </div>-->
    <!--      <div class="f-z-14 m-b-10">-->
    <!--        医患诊断采集是医疗过程中非常重要的一环，准确的信息采集和综合分析有助于医生做出正确的诊断和治疗决策，提高患者的治疗效果和生活质量。在现代医疗中，随着信息技术的发展，电子病历系统等工具也广泛应用于医患诊断采集过程，提高了信息的准确性和可追溯性。-->
    <!--      </div>-->
  </div>
  <div class="footer_btn">
    <van-button block type="primary" @click="taskProcessingHandle">
      去完成
    </van-button>
  </div>
  <div class="box"></div>
</template>
<script lang="ts" setup>
import {useRouter} from "vue-router";
// import { showSuccessToast, showFailToast } from 'vant';
import {showSuccessToast} from "vant";
const router = useRouter()
const onClickLeft = () => {
  // history.go(-2)
  router.replace('/homeNew')
}
const taskProcessingHandle = ()=>{
  // showSuccessToast('参与成功');

  router.replace('/addTasks')
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
  background: url('../../assets/img/task_bgc_02.png') no-repeat center;
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