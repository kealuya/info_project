<template>
  <van-dialog v-model:show="show" :title="title" show-cancel-button>
    <div  @click="confirm">确定</div>
  </van-dialog>
</template>

<script setup lang="ts">
import { ref,reactive } from "vue";
interface Props{
  title?:Object;
  show?:boolean;
}
defineProps<Props>();
import {Dialog} from 'vant';
const VanDialog = Dialog.Component;//这句是重点
function confirm (){
}
</script>

<style scoped>

</style>
