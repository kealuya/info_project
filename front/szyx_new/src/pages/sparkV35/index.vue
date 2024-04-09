<template>
  <van-tabs v-model:active="active" sticky>
    <van-tab title="申请办公">
      <div class="container">

        <!--第一部分 推荐 -->
        <div>
        <van-grid direction="horizontal" :column-num="3" class="cardTable">
          <van-grid-item  >
            <div>外出办公</div>
          </van-grid-item>
          <van-grid-item  >
            <div>在线办公</div>
          </van-grid-item>
          <van-grid-item  >
            <div>申请工单</div>
          </van-grid-item><!--
          <van-grid-item  >
            <div>问用药</div>
          </van-grid-item>-->
        </van-grid>
        </div>

        <van-divider />

        <!--第二部分  问 -->
        <div gutter="1" >
          <van-row >
            <van-col span="2" style="text-align: center;">
                <img src="src/assets/img/customerPng.png" class="touxiangimg">
            </van-col>
            <van-col span="22">
              <van-cell-group inset>
                <van-field
                    v-model="requestDataStr"
                    rows="1"
                    autosize
                    type="textarea"
                    readonly

                />
              </van-cell-group>
            </van-col>
          </van-row>
        </div>

        <van-divider />

        <!--第二部分  答 -->
        <div gutter="1" >
          <van-row >
            <van-col span="2" style="text-align: center;">
              <img src="src/assets/img/customerPng.png" class="touxiangimg">
            </van-col>
            <van-col span="22">
              <van-cell-group inset>
                <van-field
                    v-model="responseDataStr"
                    rows="1"
                    autosize
                    type="textarea"
                    readonly

                />
              </van-cell-group>
            </van-col>
          </van-row>
        </div>




        <!--第三部分  提问输入框 -->
        <div >
        <van-cell-group inset class="bottom_box">
          <van-field
              v-model="requestDataStr"
              rows="1"
              autosize
              label="请提问"
              type="textarea"
              placeholder="在此输入您想了解的内容，Shift+Enter换行"
              @keydown.enter="alertMsg"
          />
        </van-cell-group>
        </div>

      </div>
    </van-tab>
  </van-tabs>
</template>

<script setup lang="ts" >
import {ref, inject} from "vue";

const active = ref(0);
let userInfoData: any = inject("userInfo");

const alertMsg = () => {
  alert("提交问题")
};

const requestDataStr =ref("办公需要的设备有哪些？")
const responseDataStr =ref("我不回答你不回答，大家都不回答")





</script>

<style lang="less" scoped>

.cardTable{
  margin-top: 10px;
}

.touxiangimg{
  width: 20px;
  height: 20px;
}

.container{
  position: relative;
  width: 100%;
  min-height: 780px;
}
.bottom_box{
  min-width: 92%;
  position: absolute;
  bottom: 0;
}


</style>




