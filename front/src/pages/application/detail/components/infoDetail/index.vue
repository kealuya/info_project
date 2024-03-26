<template>
  <div class="formCard">
    <div class="m10" v-for="item in list">
      <van-row v-if="item.isImportant=='1'||item.value!==''">
        <van-col class="f14 graytxt">{{item.label}}：</van-col>
        <van-col class="f14">{{item.value}}</van-col>
      </van-row>
    </div>
  </div>
</template>

<script lang="ts" setup>
interface Props{
  list?:[];

}
defineProps<Props>()
</script>

<style scoped>
.formCard {
  margin: 10px 5px;
  padding-top: 10px;
  padding-bottom: 5px;
  background: white;
  border-radius: 10px;
}
</style>
