<template>
  <van-nav-bar
      title="业务详情"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
  />
  <div class="content">
    <div class="title_one m-b-10 f-z-14" >支原体抗生素大环内酯类的研究会议</div>
    <div class="title_two f-z-14 f-w-550  m-b-10">
      <span>文档摘要</span>
    </div>
    <div class="f-z-12 m-b-10">
      &nbsp;&nbsp;抗菌药物（antibacterial drugs）来源于细菌、霉菌，或者通过化学的从头合成（de novo synthesis）制备。从技术上讲，“抗生素（antibiotics）”仅指那些源自细菌或霉菌的抗菌剂，但通常与“抗菌药物（antibacterial drugs，或antibacterials）”同义使用。
      抗生素有许多作用机制，包括：抑制细胞壁合成、增加细胞膜通透性、干扰蛋白质合成、核酸代谢和其他代谢过程（如叶酸合成）
    </div>
<!--    <div class="f-z-12 m-b-10">
      &nbsp;&nbsp;病史采集：医生会询问患者有关疾病发生的时间、症状的持续时间、疼痛的性质、病史（包括既往病史、家族史等）等信息。患者提供详细的病史对于正确诊断疾病非常重要。
    </div>-->
<!--    <div class="f-z-12 m-b-10">
      &nbsp;&nbsp;症状采集：医生会询问患者当前的症状，比如疼痛、咳嗽、发热等。患者描述症状的特点和变化有助于医生判断疾病的可能性。
    </div>-->
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
      <span>文档纪要</span>
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
      <div class="c-blue f-z-14">20240408- 支原体抗生素大环内酯类文档.word <span>&nbsp;下载附件</span></div>
    </div>
    <div class="f-z-14  m-b-10 space">
      <span class="f-w-550">文档脑图</span>
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
      <div class="c-blue f-z-14"> &nbsp;&nbsp;20240408-支原体抗生素大环内酯类脑图.xmind<span>&nbsp;下载附件</span></div>
    </div>

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
 `;
const text =
    ' 1. 报销流程讨论：\n' +
    '      - 报销员指定报销项目提交审批。\n' +
    '      - 项目负责人审批后进行报销。\n' +
    '      - 团委部门结算其负担的费用部分。\n' +
    '      2. 学生团队出行报销处理：\n' +
    '      - 存在院系层面的报销。\n' +
    '      - 需要按每个学生的明细检索并发起二次报销。\n' +
    '      - 二次报销选择院系所负担的项目，提交审批。\n' +
    '      - 记录第一次与第二次报销信息，进行第三次报销。\n' +
    '      - 所有报销层级需附上发票信息。\n' +
    '\n' +
    '      3. 报销流程复杂性讨论：\n' +
    '      - 当前流程复杂，考虑创建单独模块处理团委报销。\n' +
    '      - 对于团队出行的多个报销处理方式未确定。\n' +
    '      - 讨论是否应将团委报销单独处理。\n' +
    '\n' +
    '      4. 报销明细表设计讨论：\n' +
    '      - 设计了一个差旅报销汇总表。\n' +
    '      - 领队填写实际报销金额，计算应报和实报金额。\n' +
    '      - 多次选择同一行程申请单的控制方法未确定。\n' +
    '\n' +
    '      5. 功能开发与评估：\n' +
    '      - 老师希望前台集成多项功能。\n' +
    '      - 担心功能多而使用情况不佳，需做前评估。\n' +
    '      - 先梳理逻辑框架和业务流程，确保易用性。\n' +
    '\n' +
    '      6. 后续步骤讨论：\n' +
    '      - 确定先后顺序，先开发效果显著的部分。\n' +
    '      - 讨论整个逻辑框架的合理性和业务匹配度。\n' +
    '      - 确保老师能按照流程完成任务。\n' +
    '\n' +
    '      7. 行动计划：\n' +
    '      - 参与者需再次审视文档，梳理思路。\n' +
    '      - 下一步聚焦于确定可执行的部分和设计问题。\n' +
    '      - 决定哪些部分现在可做，哪些需调整或延后开发。';
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
