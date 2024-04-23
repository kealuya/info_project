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
  options.value = {
  tooltip: {
    trigger: 'item'
  },
  legend: {
    top: '5%',
    left: 'center'
  },
  series: [
    {
      name: 'Access From',
      type: 'pie',
      radius: ['40%', '70%'],
      avoidLabelOverlap: false,
      itemStyle: {
        borderRadius: 10,
        borderColor: '#fff',
        borderWidth: 2
      },
      label: {
        show: false,
        position: 'center'
      },
      emphasis: {
        label: {
          show: true,
          fontSize: 40,
          fontWeight: 'bold'
        }
      },
      labelLine: {
        show: false
      },
      data: [
        { value: 1048, name: 'Search Engine' },
        { value: 735, name: 'Direct' },
        { value: 580, name: 'Email' },
        { value: 484, name: 'Union Ads' },
        { value: 300, name: 'Video Ads' }
      ]
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