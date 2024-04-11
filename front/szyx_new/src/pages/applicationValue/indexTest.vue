<template>
  <svg class="flex-1" ref="svgRef" />
</template>

<script>
  import { ref, onMounted, onUpdated } from 'vue';
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
`;

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
