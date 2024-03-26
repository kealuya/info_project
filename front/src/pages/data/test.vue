<template>
    <NavBar title="echarts 测试" />
    <div>
        <div id="main" class="testMain"></div>
    </div>
</template>

<script setup lang="ts">
import NavBar from '../../components/navBar/index.vue'
//1.导入echarts
import * as echarts from 'echarts';
import {onMounted} from 'vue'
onMounted(() => {
    init();
});
function init() {
    // 基于准备好的dom，初始化echarts实例
    let Chart = echarts.init(document.getElementById("main"));
    // 绘制图表
    let options = {
        title: {
            text: "ECharts 测试",
        },
        tooltip: {},
        xAxis: {
            data: ["衬衫", "羊毛衫", "雪纺衫", "裤子", "高跟鞋", "袜子"],
        },
        yAxis: {},
        series: [
            {
                name: "销量",
                type: "bar",
                data: [5, 20, 36, 10, 10, 20],
            },
        ],
    };

    // 渲染图表
    Chart.setOption(options);
}

</script>

<style lang="less" scoped>
.testMain {
    width: 100vw;
    height: 80vh;
}
</style>
