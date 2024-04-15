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
<!--      <div class="title_one m-b-10" >药品疗效采访</div>-->
<!--      <div class="title_two f-z-14 f-w-550  m-b-10">-->
<!--        <span>任务目标</span>-->
<!--&lt;!&ndash;        <span class="m-l">发布日期</span>&ndash;&gt;-->
<!--      </div>-->
<!--      <div class="f-z-14  m-b-10">-->
<!--        &nbsp; &nbsp;了解特定药品在治疗某些疾病时的疗效和患者反馈。-->
<!--      </div>-->
<!--      <div class="title_two f-z-14 f-w-550  m-b-10">-->
<!--        <span>任务内容</span>-->
<!--      </div>-->
<!--      <div class="f-z-14 m-b-10">-->
<!--        &nbsp; &nbsp; 与医生、患者或相关人员进行访谈,收集某种药品在临床使用中的疗效数据、不良反应情况、用药体验等。可以录音或记录文字资料。-->
<!--      </div>-->
<!--    </div>-->
<!--    <div class="footer_btn">-->
<!--      <van-button block type="primary" @click="taskProcessingHandle">-->
<!--        立即参与-->
<!--      </van-button>-->
<!--&lt;!&ndash;      <van-loading color="#1989fa" v-if="isLoading"/>&ndash;&gt;-->
<!--    </div>-->
<!--  <div class="box"></div>-->
</template>
<script lang="ts" setup>
import taskCard from '../../components/taskCard/index.vue'
import {useRouter} from "vue-router";
// import { showSuccessToast, showFailToast } from 'vant';
import {showSuccessToast} from "vant";
import {onMounted, ref} from "vue";
import bgc from "../../assets/img/task_bgc_01.png";
const router = useRouter();
// const route = useRoute();
const isLoading = ref(false)
const isComplete = ref<Boolean>(false)
const title = ref<string>('任务详情')
const taskTitle = ref<string>('互联网销售的目标设定、客户分析\n')
const picSrc = ref<string>(bgc)
const target = ref<string>('确定推广的具体目标，如增加网站流量、提高转化率。')
const content = ref<string>('增加销售额等利用搜索引擎广告（如谷歌广告）提高网站曝光度，\n' +
    '吸引潜在客户。分析目标客户群体的特征、偏好和行为，以便制定针对性的推广策略。关注留存率、重复购买率等指标，确保客户满意度并提高忠诚度。\n' +
    '分析客户的购买历史、浏览习惯、交互行为，了解他们的偏好和需求。确定客户的重要性和价值，重点关注高价值客户，并提供相应服务和优惠。\n')
const onClickLeft = () => {
  // history.go(-2)
  router.replace('/homeNew')
}
const taskProcessingHandle = ()=>{
  // isLoading.value = true
  showSuccessToast('参与成功');
  setTimeout(() => {
    router.replace('/homeNew')
  }, 2000);
}
</script>
<!--<style lang="less" scoped>-->
<!--.task{-->
<!--  //margin-top: 2vw;-->
<!--  //border-radius: 10px;-->
<!--  //padding: 0 2vw;-->
<!--  width:96vw;-->
<!--  margin: 2vw;-->
<!--  height: 160px;-->
<!--  //background: url('../../assets/img/task_bgc_01.png') no-repeat center;-->
<!--  //background-size: cover;-->
<!--  //box-sizing: border-box;-->
<!--  //margin-left: 2vw;-->
<!--  .task_card{-->
<!--    width: 62vw;-->
<!--    background-color: #e9f8f1;-->
<!--    opacity: 0.5; /* 设置透明度为50% */-->
<!--    color:#000000;-->
<!--    margin: 2vw;-->
<!--    padding:2vw;-->
<!--    font-size: 1em;-->
<!--    border-radius: 10px;-->
<!--    .content-font{-->
<!--      font-size: 0.75em;-->
<!--      margin-top:1vw;-->
<!--    }-->
<!--  }-->
<!--  .btn{-->
<!--    border-radius: 10px;-->
<!--    background-color: #253334;-->
<!--    width:42vw;-->
<!--    margin-left: 2vw;-->
<!--    padding:1vw;-->
<!--    text-align: center;-->
<!--    font-size: 1em;-->
<!--  }-->
<!--}-->
<!--.box{-->
<!--  height:50px;-->
<!--}-->
<!--.content{-->
<!--  margin:2vw;-->
<!--  line-height:3vh;-->
<!--.title_one{-->
<!--  text-align:center;-->
<!--  font-weight: 550;-->
<!--  font-size: 0.875em;-->
<!--}-->
<!-- .f-z-14{-->
<!--   font-size: 0.75em;-->
<!-- }-->
<!--  //.f-z-12{-->
<!--  //  font-size: 0.75em;-->
<!--  //}-->
<!--  .f-w-550{-->
<!--    font-weight: 550;-->
<!--  }-->
<!--  .m-l{-->
<!--    margin-left: 20vw;-->
<!--  }-->
<!--  .m-b-10{-->
<!--    margin-bottom:10px;-->
<!--  }-->
<!--}-->
<!--.footer_btn{-->
<!--  width:92vw;-->
<!--  position: fixed;-->
<!--  bottom:60px;-->
<!--  padding:2vw;-->
<!--  margin: 2vw;-->
<!--}-->
<!--::v-deep(.van-nav-bar__content) {-->
<!--  background-color: #0088ff;-->
<!--}-->

<!--::v-deep(.van-nav-bar__title) {-->
<!--  color: #ffffff;-->
<!--}-->

<!--::v-deep(.van-nav-bar__text) {-->
<!--  color: #ffffff;-->
<!--}-->

<!--::v-deep(.van-icon-arrow-left:before) {-->
<!--  color: #ffffff;-->
<!--}-->
<!--</style>-->
<script setup lang="ts">
</script>