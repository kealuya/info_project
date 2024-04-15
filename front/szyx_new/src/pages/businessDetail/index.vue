<template>
  <van-nav-bar
      title="业务详情"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
  />
  <div class="content">
    <div class="title_one m-b-10 f-z-14" >针对当前互联网营销策略会议</div>
    <div class="title_two f-z-14 f-w-550  m-b-10">
      <span>会议摘要</span>
    </div>
    <div class="f-z-12 m-b-10">
      &nbsp;&nbsp;互联网营销策略包括多种方法和技术，旨在利用互联网平台、工具和资源来推广产品或服务、增加品牌曝光、吸引潜在客户并提高销售业绩。
    </div>
    <div class="f-z-12 m-b-10">
      &nbsp; 网站优化（SEO）： 通过优化网站内容和结构，提高在搜索引擎中的排名，增加有机流量
      搜索引擎营销（SEM）： 利用付费广告在搜索引擎上展示，提高网站访问量和曝光度
      内容营销： 制作有吸引力和有价值的内容，如博客文章、视频、社交媒体内容等，吸引目标受众并建立品牌认知度
      社交媒体营销： 通过社交媒体平台与潜在客户互动，推广品牌、产品和服务，增加用户参与和忠诚度
      影响营销： 与影响者合作，借助他们的影响力在社交媒体上推广产品或服务
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
      <span class="f-w-550">会议记录</span>
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
      <div class="c-blue f-z-14">20240408-聚醚酮酮数字化临床效果会议.word <span>&nbsp;下载附件</span></div>
    </div>
    <div class="f-z-14  m-b-10 space">
      <span class="f-w-550">会议脑图</span>
      <div class="right">
        <van-image :src="Ai"  width="23" height="23">
        </van-image>
        <span class="c-blue">生成</span>
      </div>
    </div>
    <svg class="flex-1" ref="svgRef" />
    <div class="metting">
      <van-image :src="Xmind"  width="23" height="23">
      </van-image>
      <div class="c-blue f-z-14"> &nbsp;&nbsp;20240408-聚醚酮酮数字化临床效果会议.xmind<span>&nbsp;下载附件</span></div>
    </div>
    <div class="title_two f-z-14 f-w-550  m-b-10">
      <span>会议信息</span>
    </div>
    <div class="f-z-12">会议地址：天津市南开区红旗南路大厦B座</div>
    <div class="f-z-12">会议时长：1小时53分钟</div>
  </div>
  <div class="box"></div>
</template>
<script setup>
import Ai from '../../assets/img/ai.png';

import Word from '../../assets/img/word.png'
import Xmind from '../../assets/img/xmind.png'
import {useRouter} from "vue-router";
const router = useRouter()
const onClickLeft = () => {
  router.push('/business')
}
</script>
<script>
import naotu from '../../assets/img/naotu.png'
import {useRouter} from "vue-router";

import { Transformer } from 'markmap-lib';
import { Markmap } from 'markmap-view';
const transformer = new Transformer();
const initValue = `# 会议纪要：

## 报销流程讨论：
- 报销员指定报销项目提交审批。
- 项目负责人审批后进行报销。
- 团委部门结算其负担的费用部分。

## 学生团队出行报销处理：
- 存在院系层面的报销。
- 需要按每个学生的明细检索并发起二次报销。
- 二次报销选择院系所负担的项目，提交审批。
- 记录第一次与第二次报销信息，进行第三次报销。
- 所有报销层级需附上发票信息。

## 报销流程复杂性讨论：
- 当前流程复杂，考虑创建单独模块处理团委报销。
- 对于团队出行的多个报销处理方式未确定。
- 讨论是否应将团委报销单独处理。`;
const text =
    ' 1. 开场致辞：\n' +
    '      - 欢迎词和会议介绍。\n' +
    '      2. 上季度营销活动回顾和效果分析\n' +
    '      - 讨论本季度目标和关键重点\n' +
    '      - 制定本季度营销策略。\n' +
    '      - 提出针对市场趋势的策略建议.\n' +
    '\n' +
    '      3. 各小组汇报本季度营销计划和执行方案\n' +
    '      - 讨论营销预算及资源分配。\n' +
    '      - 确定本季度预算分配及资源需求\n' +
    '      - 推进销售与营销协同\n。\n' +
    '\n' +
    '      4. 报销明细表设计讨论\n' +
    '      - 设计了一个差旅报销汇总表。\n' +
    '      - 领队填写实际报销金额，计算应报和实报金额。\n' +
    '      - 多次选择同一行程申请单的控制方法未确定。\n' +
    '\n' +
    '      5. 探讨如何提高销售与营销部门协同效能：\n' +
    '      - 集成多项功能。\n' +
    '      - 担心功能多而使用情况不佳，需做前评估。\n' +
    '      - 先梳理逻辑框架和业务流程，确保易用性。\n' +
    '\n' +
    '      6. 后续步骤讨论\n' +
    '      - 确定先后顺序，先开发效果显著的部分。\n' +
    '      - 讨论整个逻辑框架的合理性和业务匹配度。\n' +
    '      - 确保老师能按照流程完成任务。;'
export default {
  name: 'App',
  data() {
    return {
      value: initValue,
    };
  },
  watch: {
    value: 'update',
  },
  methods: {
    update() {
      const { root } = transformer.transform(this.value);
      this.mm.setData(root);
      this.mm.fit();
    },
  },
  mounted() {
    this.mm = Markmap.create(this.$refs.svgRef);
    this.update();
  },
};

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
