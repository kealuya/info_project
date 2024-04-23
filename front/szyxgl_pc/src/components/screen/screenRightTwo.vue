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
  options.value ={
    xAxis: {
      type: 'category',
      data: ['星期一', '星期二', '星期三', '星期四', '星期五', '星期六', '星期日']
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        data: [120, 200, 150, 80, 70, 110, 130],
        type: 'bar'
      }
    ],
    textStyle:{
      color:'#fff'
    }
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