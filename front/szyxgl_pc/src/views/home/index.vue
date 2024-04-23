<template>
  <div class="home-block">
    <el-card class="header">
      <el-row class="m-b-10 f-w-550">销售数据</el-row>
      <el-row :gutter="20" class="divider-row home_header">
        <el-col :span="3" class="divider-col">
          <div class="grid-content ep-bg-purple m-b-10">
            商机成单数
          </div>
          <div class="m-b-10 f-z-20">1319</div>
          <div>较昨日</div>
        </el-col>
        <el-col :span="1" class="divider-col">
          <el-divider direction="vertical" class="vertical-divider"></el-divider>
        </el-col>
        <el-col :span="3" class="divider-col col_card_center">
          <div class="grid-content ep-bg-purple-light m-b-10">
            线索购买数
          </div>
          <div class="f-z-20 m-b-10">1319</div>
          <div>较昨日</div>
        </el-col>
        <el-col :span="1" class="divider-col col_card_center">
          <el-divider direction="vertical" class="vertical-divider"></el-divider>
        </el-col>
        <el-col :span="3" class="divider-col col_card_center">
          <div class="grid-content ep-bg-purple m-b-10">
            临时商机新增数
          </div>
          <div class="f-z-20 m-b-10">1319</div>
          <div>较昨日</div>
        </el-col>
        <el-col :span="1" class="divider-col col_card_center">
          <el-divider direction="vertical" class="vertical-divider"></el-divider>
        </el-col>
        <el-col :span="3" class="divider-col col_card_center">
          <div class="grid-content ep-bg-purple-light m-b-10">
            个人商机新增数
          </div>
          <div class="f-z-20 m-b-10">1319</div>
          <div>较昨日</div>
        </el-col>
        <el-col :span="1" class="divider-col col_card_center">
          <el-divider direction="vertical" class="vertical-divider"></el-divider>
        </el-col>
        <el-col :span="3" class="divider-col centered-column col_card_center">
          <div class="grid-content ep-bg-purple m-b-10">
            联系人新增数
          </div>
          <div class="f-z-20 m-b-10">1319</div>
          <div>较昨日</div>
        </el-col>
        <el-col :span="1" class="divider-col">
          <el-divider direction="vertical" class="vertical-divider"></el-divider>
        </el-col>
        <el-col :span="2" class="divider-col col_card">
          <div class="grid-content ep-bg-purple-light m-b-10">
            跟进数
          </div>
          <div class="f-z-20 m-b-10">1319</div>
          <div>较昨日</div>
        </el-col>
      </el-row>
    </el-card>
    <el-row :gutter="10" class="middle">
      <el-col :span="16"><div class="grid-content ep-bg-purple" />
        <el-card>
          <div class="f-z-14 f-w-550">任务</div>
          <el-row :gutter="10">
            <el-col :span="10" class="calendar-box">
              <el-calendar v-model="value" class="calendar"/>
            </el-col>
            <el-divider direction="vertical" class="custom-divider"/>
            <!--          内容区-->
            <el-col :span="13">
              <div class="content-box">
                <div class="content-box-left">
                  <el-tag type="primary">已完成</el-tag>
                  <div class="content-m-10">13:00</div>
                  <div>拜访北京城市网邻信息技术有限公司</div>
                </div>
                <div class="content-box-right">
                  <el-image :src="FLAG" class="flag"></el-image>
                  <span>写跟进</span>
                </div>
              </div>
              <div class="content-box">
                <div class="content-box-left">
                  <el-tag type="primary">已完成</el-tag>
                  <div class="content-m-10">13:00</div>
                  <div>拜访北京城市网邻信息技术有限公司</div>
                </div>
                <div class="content-box-right">
                  <el-image :src="FLAG" class="flag"></el-image>
                  <span>写跟进</span>
                </div>
              </div>
              <div class="content-box">
                <div class="content-box-left">
                  <el-tag type="primary">已完成</el-tag>
                  <div class="content-m-10">13:00</div>
                  <div>拜访北京城市网邻信息技术有限公司</div>
                </div>
                <div class="content-box-right">
                  <el-image :src="FLAG" class="flag"></el-image>
                  <span>写跟进</span>
                </div>
              </div>
              <div class="content-box">
                <div class="content-box-left">
                  <el-tag type="danger">已完成</el-tag>
                  <div class="content-m-10">13:00</div>
                  <div>拜访北京城市网邻信息技术有限公司</div>
                </div>
                <div class="content-box-right">
                  <el-image :src="FLAG" class="flag"></el-image>
                  <span>写跟进</span>
                </div>
              </div>
              <div class="content-box">
                <div class="content-box-left">
                  <el-tag type="danger">已完成</el-tag>
                  <div class="content-m-10">13:00</div>
                  <div>拜访北京城市网邻信息技术有限公司</div>
                </div>
                <div class="content-box-right">
                  <el-image :src="FLAG" class="flag"></el-image>
                  <span>写跟进</span>
                </div>
              </div>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
      <el-col :span="8"><div class="grid-content ep-bg-purple" />
        <el-card class="clue-card">
          <div class="clue-box">
            <div class="clue-box-title">
              <div>最新更新线索</div>
              <div class="qy">更新企业 <div class="f-c-blue">4706家</div></div>
            </div>
            <div class="clue-box-content" v-for="item in 4">
              <div class="clue-box-content-left">
                <el-image :src="company" class="company-icon"></el-image>
              </div>
              <div class="clue-box-content-right">
                <div>
                  <div class="m-b-10">北京和立互联网科技有限公司</div>
                  <div class="f-z-12-abb8cb">胡伟 <span>人力资源</span></div></div>
                <div>5分钟前</div>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <!--    统计图部分-->
    <el-row :gutter="10" class="bottom-box">
      <el-col :span="16">
        <el-card class="notice-card">
          <div>任务接办效率</div>
          <div ref="chart" style="width: 100%; height: 350px;"></div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card class="notice-card">
          <div>通知公告</div>
          <div class="notice" v-for="item in 6">
            <div class="notice-title"><div class="dot"></div>此处为通知公告名称此处为通知公告名称此处为..</div>
            <div>2023-04-05</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

  </div>
</template>

<script setup lang="ts">
import * as echarts from 'echarts';
import company from '../../assets/images/home/company.png'
import { ref,onMounted } from 'vue'
import FLAG from '../../assets/images/home/flag.png'
const value = ref(new Date())
const chart = ref<HTMLElement | null>(null);

const option =ref( {
  xAxis: {
    type: 'category',
    data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
  },
  yAxis: {
    type: 'value'
  },
  series: [
    {
      data: [120, 200, 150, 80, 70, 110, 130],
      type: 'bar',
      itemStyle: {
        color: '#007aff', // 设置柱状图颜色为红色
      },
    }
  ]
})
onMounted(() => {
  if (chart.value) {
    const myChart = echarts.init(chart.value);
    myChart.setOption(option.value);
  }
});


</script>
<style lang="scss" scoped>
.home-block {
  //width: 100%;
  overflow: hidden;
  padding: 10px;
  font-size: 14px;
  .header{
    border-radius: 10px;
    min-height: 150px;
  }
  .home_header{
    display: flex;
    justify-content: space-between;
    align-items: center;
    .vertical-divider {
      height: 50px; /* 设置分割线的高度 */
      margin: auto; /* 垂直居中 */
    }
    .col_card_center{
      text-align: center;
    }
    .col_card{
      text-align: center;
    }

  }
  .middle{
    margin-top: 10px;
    .calendar{
      width: 100%;
      height: 100%;
      :deep(.el-calendar-table .el-calendar-day){
        width: 100%;
        height: 30px;
        text-align: center;
      }
    }
    .custom-divider{
      height: 280px;
    }
    .content-box{
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 20px;
      padding-left: 16px;
      .content-box-left{
        display:flex;
        align-items: center;
      }
      .content-box-right{
        display: flex;
        .flag{
          width: 20px;
        }
      }
      .content-m-10{
        margin: 0 20px;
      }

    }
    .clue-card{
      max-height:344px;
      min-height: 344px;
      overflow-y: auto;
      .clue-box{
        font-size: 14px;
        .clue-box-title{
          display: flex;
          justify-content: space-between;
          .qy{
            display: flex;
            align-items: center;
          }
        }
        .clue-box-content{
          display: flex;
          align-items: center;
          margin-top: 20px;
          .clue-box-content-left{
            display: flex;
            .company-icon{
              width: 30px;
              height: 30px;
              margin-right: 10px;
            }
          }
          .clue-box-content-right{
            display: flex;
            justify-content: space-between;
            flex: 1;
            align-items: center;
          }
        }
      }
    }

  }
  //统计图部分
  .bottom-box{
    margin-top: 20px;
    .notice-card{
      max-height: 410px;
      min-height: 410px;
      overflow-y: auto;
      .notice{
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-top: 20px;
        .notice-title{
          display: flex;
          align-items: center;
          .dot{
            width: 10px;
            height: 10px;
            background-color: #a6d1ff;
            border-radius: 50%;
            margin-right: 10px;
          }
        }
      }
    }

  }
  .m-b-10{
    margin-bottom: 10px;
    color: #000000;
  }
  .f-z-20{
    font-size: 22px;
  }
  .f-w-550{
    font-weight: 550;
  }
  .f-z-12-abb8cb{
    color: #ccc;
    font-size: 12px;
  }
  .f-c-blue{
    color: #007aff;
    margin-left: 10px;
    //color: ;
  }

}
</style>

