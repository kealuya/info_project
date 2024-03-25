<template>
  <div class="tr">
    <div v-for="item in stateIcon">
      <van-image width="80" height="80" v-if="type==item.key" :src="item.value"/>
    </div>

  </div>
</template>

<script lang="ts" setup>
import  waitshenpi from '../../assets/img/waitshenpi_icon.png';
import mianshen from '../../assets/img/mianshen.png';
import yichexiao from  '../../assets/img/yichexiao_icon.png';
import yijujue from '../../assets/img/yijujue_icon.png';
import yitongyi from  '../../assets/img/yitongyi_icon.png'
import {ref} from "vue";
const stateIcon=ref([
  {key:"待审批",value:waitshenpi},
  {key:"免审",value:mianshen},
  {key:"已撤销",value:yichexiao},
  {key:"已拒绝",value:yijujue},
  {key:"已同意",value:yitongyi},
]);
interface Props{
  type?:String;

}
defineProps<Props>()
</script>

<style scoped>

</style>
