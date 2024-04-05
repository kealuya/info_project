<!--预算科目卡片-->
<template>
  <div class="approvalCard">
    <van-row justify="space-between" align="center">
      <van-row  align="center">
        <van-col >
          <span class="f18 bold">{{getData.subjectName}}</span>
        </van-col>
        <van-col>
          <div class="subjectTips">{{getData.subjectId}}</div>
        </van-col>
      </van-row>
    </van-row>
    <van-row>
     <div class="mt5">
       <span  style="color: #c9c9c9;font-size: 12px">创建时间：{{getData.creationTime}}</span>
     </div>
    </van-row>
  </div>
</template>

<script setup lang="ts">
interface Props{
  getData?:Object;

}
defineProps<Props>()
</script>

<style scoped>
.approvalCard{
  margin: 10px;
  padding: 10px;
  border-radius: 10px;
  background: #ffffff;
}
.subjectTips{
  margin: 0px 5px;
  padding: 5px 10px;
  border-radius: 8px;
  background: #B6DBFF;
  color: #0080ff;
  font-size: 14px;
}
</style>
