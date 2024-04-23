<template>
  <div class="line_box">
    <CommonEcharts :width="width" :height="height" :options="options"></CommonEcharts>
  </div>
</template>
<script lang="ts" setup>
  // import Bar from "~/assets/images/echarts/bar.png";
  import { ref, watch } from "vue";
  interface BackData {
    data: object
  }

  const props = withDefaults(defineProps<BackData>(), {
    data: () => ({})
  })

  watch(() => props.data, (newValue: any) => {
    setOptions(newValue)
  }, { deep: true })

  const width = ref<string>('auto')
  const height = ref<string>('100%')
  const options = ref<object | any>({})
  const getMaxData = (data: any) => {
    let max = data[0];
    let array = []
    for (let i = 1; i < data.length; i++) {
      if (max < data[i]) {
        max = data[i]
      }
    }
    for (let j = 0; j < data.length; j++) {
      array.push(max)
    }
    return array;
  }
  const setOptions = (val: { xData: Object; barData: any }) => {
    let { xData, barData } = val;
    const bgData = getMaxData(barData.data)
    options.value ={
  // title: {
  //   text: 'Stacked Line'
  // },
  tooltip: {
    trigger: 'axis'
  },
  legend: {
    data: ['Email', 'Union Ads', 'Video Ads', 'Direct', 'Search Engine']
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  // toolbox: {
  //   feature: {
  //     saveAsImage: {}
  //   }
  // },
  xAxis: {
    type: 'category',
    boundaryGap: false,
    data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
  },
  yAxis: {
    type: 'value'
  },
  series: [
    {
      name: 'Email',
      type: 'line',
      stack: 'Total',
      data: [120, 132, 101, 134, 90, 230, 210]
    },
    {
      name: 'Union Ads',
      type: 'line',
      stack: 'Total',
      data: [220, 182, 191, 234, 290, 330, 310]
    },
    {
      name: 'Video Ads',
      type: 'line',
      stack: 'Total',
      data: [150, 232, 201, 154, 190, 330, 410]
    },
    {
      name: 'Direct',
      type: 'line',
      stack: 'Total',
      data: [320, 332, 301, 334, 390, 330, 320]
    },
    {
      name: 'Search Engine',
      type: 'line',
      stack: 'Total',
      data: [820, 932, 901, 934, 1290, 1330, 1320]
    }
  ]
};
  }

</script>
<style lang="scss" scoped>
  .line_box {
    flex: 1;
    // width: 400px;
    // height: 300px;
  
  }
</style>