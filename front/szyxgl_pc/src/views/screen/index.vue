<template>
  <div class="container">
    <div class="screen" ref="screen">
      <div class="top">
        <div class="topLeft" @click="router.replace('home')">
          <div class="topLeftBox">
            <img class="topLeftImg" src="../../assets/images/screen/dataScreen-header-btn-bg-l.png" alt="">
            <span class="topLeftTitle">首页</span>
          </div>
        </div>
        <div class="topModdle">
          <div class="topModdleBox">
            <img class="topModdleImg" src="../../assets/images/screen/dataScreen-header-center-bg.png" alt="">
            <span class="topModdleTitle">智慧可视化大数据展示平台</span>
          </div>
        </div>
        <div class="topRight"></div>
      </div>
      <div class="bottom">
        <div class="screenLeft">
          <div class="slOne">
            <div class="sEchartsItem">
              <div class="chartsTitle">
                <div>实时游客统计</div>
                <img src="../../assets/images/screen/dataScreen-title.png" alt="">
              </div>
              <div class="slChart">
                <screenLeftOne :data="BarData"></screenLeftOne>
              </div>
            </div>
            <!-- <AsyncLiquidfil :allPeopleNum="allPeopleNum" :peopleNum="peopleNum"></AsyncLiquidfil> -->
          </div>
          <div class="slOne">
            <div class="sEchartsItem">
              <div class="chartsTitle">
                <div>业绩季度图</div>
                <img src="../../assets/images/screen/dataScreen-title.png" alt="">
              </div>
              <div class="slChart">
                <screenLeftTwo :data="BarData"></screenLeftTwo>
              </div>
            </div>
          </div>
          <div class="slOne">
            <div class="sEchartsItem">
              <div class="chartsTitle">
                <div>实时游客统计1</div>
                <img src="../../assets/images/screen/dataScreen-title.png" alt="">
              </div>
              <div class="slChart">
                <Line :data="BarData"></Line>
              </div>
            </div>
          </div>

        </div>
        <div class="screenMiddle">
          <div class="smOne"></div>
          <div class="smTwo"></div>

        </div>
        <div class="screenLeft">
          <div class="slOne">
            <div class="sEchartsItem">
              <div class="chartsTitle">
                <div>业绩季度图1</div>
                <img src="../../assets/images/screen/dataScreen-title.png" alt="">
              </div>
              <div class="slChart">
                <Line :data="BarData"></Line>
              </div>
            </div>
            <!-- <AsyncLiquidfil :allPeopleNum="allPeopleNum" :peopleNum="peopleNum"></AsyncLiquidfil> -->
          </div>
          <div class="slOne">
            <div class="sEchartsItem">
              <div class="chartsTitle">
                <div>业绩季度图2</div>
                <img src="../../assets/images/screen/dataScreen-title.png" alt="">
              </div>
              <div class="slChart">
                <screenRightTwo :data="BarData"></screenRightTwo>
              </div>
            </div>
          </div>
          <div class="slOne">
            <div class="sEchartsItem">
              <div class="chartsTitle">
                <div>业绩季度图3</div>
                <img src="../../assets/images/screen/dataScreen-title.png" alt="">
              </div>
              <div class="slChart">
                <Pie :data="BarData"></Pie>
              </div>
            </div>
          </div>

        </div>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import {   ref, onMounted } from 'vue'
import router from "~/router";

//人数相关

//获取数据大屏展示内容盒子的DOM元素
let screen = ref()
onMounted(() => {
  screen.value.style.transform = `scale(${getScale()}) translate(-50%,-50%)`
  getBarData()
})
const BarData = ref<object>({})
// const PieData = ref<object>({})
const getBarData = () => {
  setTimeout(() => {
    BarData.value = {
      xData: [
        ['product', '2015', '2016', '2017'],
        [' ', 43.3, 85.8, 93.7],
        ['   ', 83.1, 73.4, 55.1],
        [' ', 86.4, 65.2, 82.5],
        ['  Brownie', 72.4, 53.9, 39.1]
      ],
      barData: {
        name: '企业景气指数',
        data: [120, 23, 45, 43, 87, 62],
        unit: "单位：亿元"
      }
    }
  }, 200)
}

//监听视口变化 
window.onresize = () => {
  console.log('aaaa')
  screen.value.style.transform = `scale(${getScale()}) translate(-50%,-50%)`
}
//定义大屏缩放比例
const getScale = (w = 1920, h = 1080) => {
  const ww = window.innerWidth / w
  const wh = window.innerHeight / h
  console.log('', ww, wh);

  return ww < wh ? ww : wh
}
</script>
<style  lang='scss' >
.container {
  width: 100vw;
  height: 100vh;
  background: url('../../assets/images/screen/bg.webp') no-repeat;
  background-size: 100% 100%;
  // background: #000928;


  .screen {
    width: 1920px;
    height: 1080px;
    // background: red;
    // border: 1px solid red;
    // background: #0c246a;
    position: fixed;
    left: 50%;
    top: 50%;
    transform-origin: left top;
    // border: 1px solid red;
    // border: 1px solid red;

    .top {
      width: 100%;
      height: 74px;
      color: #59b4e6;
      font-size: 16px;
      text-align: center;
      line-height: 74px;
      display: flex;

      .topLeft {
        flex: 1;
        // background: red
        background: url(../../assets/images/screen/dataScreen-header-left-bg.png) no-repeat;
        background-size: cover;
        height: 32px;
        line-height: 42px;
        display: flex;
        justify-content: end;

        .topLeftBox {
          height: 40px;
          width: 130px;
          position: relative;
          // background: red;

          .topLeftImg {
            height: 40px;
            // width: 60px;
          }

          .topLeftTitle {
            width: 130px;
            position: absolute;
            right: 16%;
            // font-weight: bold;
            top: -4px;
            // transform: translate(-50%, -50%);
          }
        }
      }

      .topModdle {
        flex: 1.5;
        height: 70px;

        .topModdleBox {
          position: relative;
          width: 100%;
          height: 70px;

          .topModdleImg {
            width: 100%;
            height: 70px;
          }

          .topModdleTitle {
            display: block;
            position: absolute;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
            font-size: 32px;
          }
        }

        // margin-top: 10px;
        // background: blue
        // background: url(../../assets/images/screen/dataScreen-header-center-bg.png) no-repeat;
        // background-size: 100%, 100%;
        // background-attachment:fixed;


      }

      .topRight {
        flex: 1;
        height: 32px;
        // background: yellow
        background: url(../../assets/images/screen/dataScreen-header-left-bg.png) no-repeat;
        background-size: cover;

      }
    }

    .bottom {
      width: 100%;
      display: flex;
      height: calc(100% - 74px);
      // border: 1px solid yellow;

      // overflow: scroll;
      .screenLeft,
      .screenMiddle {
        // flex: 1;
        width: 25%;
        display: flex;
        flex-direction: column;
        //width: 500px;

        // border: 1px solid red;
      }

      .screenLeft {
        border: none;

        .slOne {
          display: flex;
          // height: 100%;
          width: 100%;
          // margin-top: 30px;
          padding: 10px;
          box-sizing: border-box;
          // border: 1px solid pink;

          .sEchartsItem {
            width: 100%;
            height: 310px;
            background: url("@/assets/images/screen/dataScreen-main-lt.png") top center no-repeat;
            background-size: cover;
            box-sizing: border-box;
            border-radius: $base-box-radius;
            // margin-right: 2%;
            box-shadow: $base-homeCard-shadow;
            // background: $base-color-default;;
            display: flex;
            flex-direction: column;
            justify-content: space-between;

            // display: flex;
            .chartsTitle {
              width: 100%;
              height: 52px;
              // border: 1px solid red;
              font-size: 20px;
              color: #59b4e6;
            }

            .slChart {
              display: flex;
              // height: 400px;
              width: 100%;
              height: 258px;
              // margin-top: 30px;
              // padding: $base-padding-20;
              // border: 1px solid yellow;
              box-sizing: border-box;

            }
          }

        }


      }

      .screenMiddle {
        width: 50%;

        .smOne {
          flex: 2;
          border: 1px solid red;
        }

        .smTwo {
          flex: 1;
          border: 1px solid red;

        }
      }

    }

  }
}
</style>