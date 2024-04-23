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

// const width = ref<string>('auto')
// const height = ref<string>('100%')
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
  options.value = {
    legend: {},
    tooltip: {},
    dataset: {
      source: [
        ['product', '2015', '2016', '2017'],
        [' ', 43.3, 85.8, 93.7],
        ['   ', 83.1, 73.4, 55.1],
        [' ', 86.4, 65.2, 82.5],
        ['  Brownie', 72.4, 53.9, 39.1]
      ]
    },
    xAxis: { type: 'category' },
    yAxis: {},
    // Declare several bar series, each will be mapped
    // to a column of dataset.source by default.
    series: [{ type: 'bar' }, { type: 'bar' }, { type: 'bar' }]
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