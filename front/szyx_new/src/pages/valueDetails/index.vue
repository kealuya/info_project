<template>
  <van-nav-bar
      title="价值详情"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
  />
    <div class="content">
      <div class="title_one m-b-10" >2024年数字化转型互联网营销驱动销售会议</div>
      <div class="title_two f-z-14 f-w-550  m-b-10">
        <span>价值概况</span>
<!--        <span class="m-l">发布日期</span>-->
      </div>
      <div class="f-z-14  m-b-10">
        &nbsp; &nbsp;拓展市场空间意味着风险分散，一旦某个市场遇到挑战，企业仍能依靠其他市场维持业务运营通过优化销售策略，企业可以更好地了解市场需求，制定针对性的销售计划，提高竞争力并抢占市场份额。
      </div>
      <div class="title_two f-z-14 f-w-550  m-b-10">
        <span>申请价值内容</span>
      </div>
      <div class="f-z-14 m-b-10">
        &nbsp; &nbsp;有效的销售策略和市场拓展将为企业创造更多商机和增长机会，推动业务快速发展，综合来看，拓展市场空间和优化销售策略能够为企业带来多方面的价值，包括增长、创新、竞争力提升等，是实现企业长期可持续发展的关键因素。

      </div>
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
    <div class="card_small">
      <div class="left">
        <van-image :src="pingfen" width="25" height="25"></van-image>
        <div>价值申请评分</div>
      </div>
      <div class="right"> <div>87分</div></div>
    </div>
  <div class="card_small">
    <div class="left">
      <van-image :src="jine" width="25" height="25"></van-image>
      <div>价值申请金额</div>

    </div>
    <div class="right"> <div>32,000元</div></div>
  </div>
  <div style="wdith:96vw;margin:2vw;text-align: center;background-color:#fff;">
    <div><van-image :src="erweima" width="230" height="200"></van-image></div>
  </div>

    <div class="footer_btn">
<!--      <van-button block type="primary" @click="taskProcessingHandle">-->
<!--        立即参与-->
<!--      </van-button>-->
    </div>
  <div class="box"></div>
</template>
<script lang="ts" setup>
import erweima from '../../assets/img/erweima.jpg'
import pingfen from '../../assets/img/pf001.png'
import jine from '../../assets/img/jz001.png'
import {useRouter} from "vue-router";
// import { showSuccessToast, showFailToast } from 'vant';
import {showSuccessToast} from "vant";
const router = useRouter()
const onClickLeft = () => {
  // history.go(-2)
  router.replace('/applicationValue')
}
const taskProcessingHandle = ()=>{
  showSuccessToast('参与成功');

  router.replace('/homeNew')
}
</script>
<style lang="less" scoped>
.card_small{
  width: 92vw;
  margin: 2vw;
  padding:2vw;
  background-color: #fff;
  border-radius: 6px;
  display:flex;
  justify-content: space-between;
  font-size:0.875em;
  font-weight: 550;
.left{
  display:flex;
  align-items: center;
  div{
    margin-left: 2vw;
  }
}
}
.task{
  //margin-top: 2vw;
  //border-radius: 10px;
  //padding: 0 2vw;
  width:96vw;
  margin: 2vw;
  height: 160px;
  background: url('../../assets/img/home_05.png') no-repeat center;
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
