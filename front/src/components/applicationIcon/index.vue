<template>
  <div v-for="item in applicationType">
    <van-image width="30" height="30" v-show="type==item.key" :src="item.value"/>
  </div>
</template>

<script lang="ts" setup>
import  trainIcon from '../../assets/icon/application_icon_01.png';
import  rcbxIcon from '../../assets/icon/application_icon_11.png';
import  jksqIcon from '../../assets/icon/application_icon_13.png';
import  clbxIcon from '../../assets/icon/application_icon_12.png';
import  otherIcon from '../../assets/icon/application_icon_12.png';
import  travelBxIcon from '../../assets/icon/application_icon_13.png';
import  leaveIcon from '../../assets/icon/application_icon_01.png';
import  workIcon from '../../assets/icon/application_icon_02.png';
import  costIcon from '../../assets/icon/application_icon_14.png';
import  userSealIcon from '../../assets/icon/application_icon_07.png';
import  userCarIcon from '../../assets/icon/application_icon_15.png';
import  subscriptionIcon from '../../assets/icon/application_icon_06.png';
import  DSQpIcon from '../../assets/icon/application_icon_11.png';
import  cardIcon from '../../assets/icon/application_icon_04.png';
import  itemCollectionIcon from '../../assets/icon/application_icon_05.png';
import  jiJinBaoXiaoIcon from '../../assets/icon/application_icon_10.png';
import  entertainIcon from '../../assets/icon/application_icon_08.png';
import  meetingIcon from '../../assets/icon/application_icon_09.png';
import  outWorkIcon from '../../assets/icon/application_icon_03.png';
import  logIcon from '../../assets/icon/application_icon_05.png';
import {ref} from "vue";
interface Props{
  type?:String;

}
defineProps<Props>()
const applicationType=ref([
  {key:"travel",value:travelBxIcon},
  {key:"leave",value:leaveIcon},
  {key:"work",value:workIcon},
  {key:"reimbursement",value:costIcon},
  {key:"seal",value:userSealIcon},
  {key:"car",value:userCarIcon},
  {key:"shop",value:subscriptionIcon},
  {key:"sign",value:DSQpIcon},
  {key:"card",value:cardIcon},
  {key:"itemCollection",value:itemCollectionIcon},
  {key:"fund",value:jiJinBaoXiaoIcon},
  {key:"train",value:trainIcon},
  {key:"entertain",value:entertainIcon},
  {key:"meeting",value:meetingIcon},
  {key:"outWork",value:outWorkIcon},
  {key:"log",value:logIcon},
  {key:"rcbx",value:rcbxIcon},
  {key:"jksq",value:jksqIcon},
  {key:"clbx",value:clbxIcon},
  {key:"other",value:otherIcon},
]);
</script>

<style scoped>

</style>
