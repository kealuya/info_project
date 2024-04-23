<template>
  <div class="detailbox">
    <el-row :gutter="24">
      <el-col :span="18">
        <el-card class="detailbox-card">
          <div class="title">
            <div class="title-sag">
              <img src="@/assets/layouts/companyInfo/basic.png" alt="">
              <span>基础信息</span>
            </div>
          </div>
<!--          <div class="title-top">-->
<!--            <div class="info-left">企业名称：<span>苏州思客信息技术有限公司</span></div>-->
<!--            <div class="info-left">详细地址：<span>江苏省苏州市工业园区协鑫广场 1 幢 20F</span></div>-->
<!--            <div class="info-left">客户来源：<span>1111</span></div>-->
<!--          </div>-->
          <el-row :gutter="24" class="m-b-10 f-z-12">
            <el-col :span="8">
              <div class="info-left">企业名称：<span>北京森铭妙传媒科技有限公司</span></div>
            </el-col>
            <el-col :span="8">
              <div class="info-left">详细地址：<span>北京市大兴区永旺西路26号院</span></div>
            </el-col>
            <el-col :span="8">
              <div class="info-left">详细地址：<span> 网页端注册</span></div>
            </el-col>
          </el-row>

          <el-row :gutter="24" class="m-b-10 f-z-12">
            <el-col :span="8">
              <div class="info-left">联系人：<span>王星予</span></div>
            </el-col>
            <el-col :span="8">
              <div class="info-left">手机号：<span>13428091504</span></div>
            </el-col>
            <el-col :span="8">
              <div class="info-left">邮箱：<span>senming@163com</span></div>
            </el-col>
          </el-row>

          <el-row :gutter="24" class="m-b-10 f-z-12">
            <el-col :span="8">
              <div class="info-left">企业规模：<span>--</span></div>
            </el-col>
            <el-col :span="8">
              <div class="info-left">客户类型：<span>高等院校</span></div>
            </el-col>
            <el-col :span="8">
              <div class="info-left">所属行业：<span>新一代信息技术</span></div>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="detailbox-card">
          <div class="title">
            <div class="title-sag">
<!--              <img src="@/assets/layouts/companyInfo/cooperate.png" alt="">-->
              <span>合作信息</span>
            </div>
          </div>
          <el-row class="title-top">
            <el-col :span="12">
              <div class="info-right">企业状态：<span>签约完成</span></div>
            </el-col>
            <el-col :span="12">
              <div class="info-right">客户类型：<span>A类客户</span></div>
            </el-col>
          </el-row>
          <el-row :gutter="12">
            <el-col :span="12">
              <div class="days">
<!--                <img src="@/assets/layouts/companyInfo/days.png" alt="">-->
                <div class="days-info">
                  <span class="num">365</span>
                  <span class="num-bottom">上线天数</span>

                </div>

              </div>
            </el-col>
            <el-col :span="12">
              <div class="total">
<!--                <img src="@/assets/layouts/companyInfo/total.png" alt="">-->
                <div class="days-info">
                  <span class="num">12345</span>
                  <span class="num-bottom">累计Saas(元)</span>
                </div>
              </div>
            </el-col>
          </el-row>
          <el-row :gutter="12">
            <el-col :span="12">
              <div class="order">
<!--                <img src="@/assets/layouts/companyInfo/order.png" alt="">-->
                <div class="days-info">
                  <span class="num">12345</span>
                  <span class="num-bottom">累计订单(张)</span>
                </div>
              </div>
            </el-col>
            <el-col :span="12">
              <div class="consumption">
<!--                <img src="@/assets/layouts/companyInfo/sale.png" alt="">-->
                <div class="days-info">
                  <span class="num">88888</span>
                  <span class="num-bottom">累计消费(元)</span>
                </div>
              </div>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script setup lang="ts">
import { onMounted, ref,reactive,toRefs,toRef } from 'vue';
</script>
<style lang="scss" scoped>

</style>