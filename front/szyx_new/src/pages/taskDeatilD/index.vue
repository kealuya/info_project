<template>
<!--  <van-nav-bar-->
<!--      title="任务详情"-->
<!--      left-text="返回"-->
<!--      left-arrow-->
<!--      @click-left="onClickLeft"-->
<!--  />-->
  <taskCard :title="title" :taskTitle="taskTitle" :picSrc="picSrc" :target="target" :content="content" :isComplete="isComplete"></taskCard>
<!--<div class="task_box">-->
<!--  <div class="task  f_white">-->
<!--  </div>-->
<!--      <div class="height_contain">-->
<!--        <div class="content">-->
<!--          <div class="title_one m-b-10" >生物标志物作为高血压患者风险评估的有效指标</div>-->
<!--          <div class="title_two f-z-14 f-w-550  m-b-10">-->
<!--            <span>任务目标</span>-->
<!--            &lt;!&ndash;        <span class="m-l">发布日期</span>&ndash;&gt;-->
<!--          </div>-->
<!--          <div class="f-z-14  m-b-10">-->
<!--            &nbsp; &nbsp;间隙保持器有助于维持乳牙缺失后的间隙空间安装间隙保持架后，间隙可以保持在一定程度上。恒磨牙萌出后，取下间隙固位体，最大限度地保持牙齿的咬合平衡。-->
<!--          </div>-->
<!--          <div class="title_two f-z-14 f-w-550  m-b-10">-->
<!--            <span>任务内容</span>-->
<!--          </div>-->
<!--          <div class="f-z-14 m-b-10">-->
<!--            &nbsp; &nbsp;探讨聚醚酮酮(PEKK)数字化乳牙早失间隙保持器的临床应用效果。方法从2019年4—5月在南京大学医学院附属口腔医院儿童口腔科就诊的6~10岁儿童中选取下颌单侧乳磨牙早失患者共5例,其中男2例、女3例,经数字化口内光学印模制取,间隙保持器数字化设计,PEKK基材间隙保持器的数字化制作,口内试戴调整,粘接。结果5例患儿的数字化乳牙早失间隙保持器重量只有传统金属带环丝圈保持器的25%,设计和制作时间只需1 h,5例数字化间隙保持器口内观察与基牙吻合度高。结论数字化乳牙早失间隙保持器采用数字化口内印模法比传统藻酸盐印模法提高了患儿的舒适度,可当日设计完成并佩戴,并使用PEKK材料,实现了间隙保持器去金属化,无金属过敏等风险,兼顾了美观的效果,具有较好的临床应用前景。-->
<!--          </div>-->
<!--          <div class="f-z-14 m-b-10">-->
<!--            &nbsp; &nbsp;探讨聚醚酮酮(PEKK)数字化乳牙早失间隙保持器的临床应用效果。方法从2019年4—5月在南京大学医学院附属口腔医院儿童口腔科就诊的6~10岁儿童中选取下颌单侧乳磨牙早失患者共5例,其中男2例、女3例,经数字化口内光学印模制取,间隙保持器数字化设计,PEKK基材间隙保持器的数字化制作,口内试戴调整,粘接。结果5例患儿的数字化乳牙早失间隙保持器重量只有传统金属带环丝圈保持器的25%,设计和制作时间只需1 h,5例数字化间隙保持器口内观察与基牙吻合度高。结论数字化乳牙早失间隙保持器采用数字化口内印模法比传统藻酸盐印模法提高了患儿的舒适度,可当日设计完成并佩戴,并使用PEKK材料,实现了间隙保持器去金属化,无金属过敏等风险,兼顾了美观的效果,具有较好的临床应用前景。-->
<!--          </div>-->
<!--        </div>-->
<!--      </div>-->
<!--&lt;!&ndash;  <div class="footer_btn">&ndash;&gt;-->
<!--&lt;!&ndash;    <van-button block type="primary" @click="taskProcessingHandle">&ndash;&gt;-->
<!--&lt;!&ndash;      去完成&ndash;&gt;-->
<!--&lt;!&ndash;    </van-button>&ndash;&gt;-->
<!--&lt;!&ndash;  </div>&ndash;&gt;-->
<!--</div>-->
<!--  <div class="footer_btn">-->
<!--    <van-button block type="primary" @click="taskProcessingHandle">-->
<!--      去完成-->
<!--    </van-button>-->
<!--  </div>-->
<!--  <div class="box"></div>-->

<!--  <div class="footer_btn">-->
<!--    <van-button block type="primary" @click="taskProcessingHandle">-->
<!--      去完成-->
<!--    </van-button>-->
<!--  </div>-->

</template>
<script lang="ts" setup>
import bgc from '../../assets/img/task_bgc_01.png'
import { ref } from 'vue';
import taskCard from '../../components/taskCard/index.vue'
import {useRouter} from "vue-router";
// import { showSuccessToast, showFailToast } from 'vant';
import {showSuccessToast} from "vant";
const router = useRouter()
const isComplete = ref<Boolean>(true)
const title = ref<string>('任务详情')
const taskTitle = ref<string>('互联网销售的目标设定、客户分析\n')
const picSrc = ref<string>(bgc)
const target = ref<string>('确定推广的具体目标，如增加网站流量、提高转化率。\n' +
    '\n')
const content = ref<string>('增加销售额等利用搜索引擎广告（如谷歌广告）提高网站曝光度，\n' +
    '吸引潜在客户。分析目标客户群体的特征、偏好和行为，以便制定针对性的推广策略。关注留存率、重复购买率等指标，确保客户满意度并提高忠诚度。\n' +
    '分析客户的购买历史、浏览习惯、交互行为，了解他们的偏好和需求。确定客户的重要性和价值，重点关注高价值客户，并提供相应服务和优惠。\n')
const onClickLeft = () => {
  // history.go(-2)
  router.replace('/homeNew')
}
const taskProcessingHandle = ()=>{
  // showSuccessToast('参与成功');

  router.replace('/addTasks')
}
</script>
<!--<style lang="less" scoped>-->
<!--.task_box{-->
<!--  display: flex;-->
<!--  flex-direction: column;-->
<!--  align-items: center;-->
<!--  //justify-content: center;-->
<!--  padding: 2vw;-->
<!--  //height: calc(100vh - var(&#45;&#45;van-nav-bar-height - 4vh);-->
<!--  height: calc(100vh - var(&#45;&#45;van-nav-bar-height) - 14vh);-->
<!--  box-sizing: border-box;-->
<!--}-->
<!--.height_contain{-->
<!--  flex: 1;-->
<!--  width: 100%;-->
<!--  overflow-y: auto; /* 当内容超出高度时显示垂直滚动条 */-->
<!--}-->
<!--.task{-->
<!--  //margin-top: 2vw;-->
<!--  //border-radius: 10px;-->
<!--  //padding: 0 2vw;-->
<!--  width:96vw;-->
<!--  margin: 2vw;-->
<!--  height: 20vh;-->
<!--  background: url('../../assets/img/task_bgc_02.png') no-repeat center;-->
<!--  background-size: cover;-->
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
<!--//.box{-->
<!--//  height:60px;-->
<!--//}-->
<!--.content{-->
<!--  margin:2vw;-->
<!--  line-height:3vh;-->
<!--  //flex:1;-->
<!--  //overflow: auto;-->
<!--  .title_one{-->
<!--    text-align:center;-->
<!--    font-weight: 550;-->
<!--    font-size: 0.875em;-->
<!--  }-->
<!--  .f-z-14{-->
<!--    font-size: 0.75em;-->
<!--  }-->
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
<!-- width:100%;-->
<!--  height:10vh;-->
<!--  //border: 1px solid red;-->
<!--  //line-height: 10vh;-->
<!--  //position: fixed;-->
<!--  //bottom:60px;-->
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
<!--//.container {-->
<!--//  flex: 1;-->
<!--//  width: 100%;-->
<!--//  overflow-y: auto; /* 当内容超出高度时显示垂直滚动条 */-->
<!--//}-->

<!--</style>-->
<script setup lang="ts">
</script>