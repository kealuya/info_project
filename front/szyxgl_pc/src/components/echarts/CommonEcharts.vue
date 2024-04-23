<template>
  <div class="echarts-wrapper" ref="chartDom" :style="getChartStyle"></div>
</template>

<script setup lang="ts">
import { onDeactivated, onMounted, ref, watch, computed, nextTick } from 'vue'
import * as echarts from 'echarts'

interface Props {
  width?: string;
  height?: string;
  options: echarts.EChartsOption;
}

const props = withDefaults(defineProps<Props>(), {


  options: () => ({}),
  // width: () => ( ''),
  // height: () => ('' )

})
console.log('props',props);


const getChartStyle = computed(() => {
  return {
    width: `${props.width}px`,
    height: `${props.height}px`
  }
})

let myEcharts: echarts.ECharts
const chartDom = ref<HTMLDivElement>()
let _resizeHandler = ref()

const handlerResize = () => {
  myEcharts.resize()
}
const addResize = () => _resizeHandler.value = window.addEventListener('resize', handlerResize, { passive: true }) // TODO debounce
const removeResize = () => { if (_resizeHandler.value) window.removeEventListener('resize', _resizeHandler.value) }

watch(() => props.options, (newValue: echarts.EChartsOption) => {
  myEcharts.setOption(newValue)
}, { deep: true })

onMounted(() => {
  // console.log('props',props)
  nextTick().then(() => {
    myEcharts = echarts.init(chartDom.value as HTMLDivElement)
    myEcharts?.setOption(props.options, true)
    addResize()
  })
})

onDeactivated(() => removeResize())
</script>

<style lang="scss" scoped>
.echarts-wrapper {
  width: 100%;
  // padding:10px;
  height: 100%;
}
</style>
