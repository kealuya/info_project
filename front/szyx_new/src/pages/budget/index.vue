<template>
<nav-bar title="预算管理" />
  <van-row class="m10">
    <van-col>
      <div  @click="showPicker=true" >
        <span class="f14 " style="color:#0088ff"> {{budgetType}}</span>
        <van-icon name="arrow-down" color="#999999"  />
      </div>
    </van-col>
  </van-row>
  <div v-if="budgetType=='预算编制'">
    <organization-list/>
  </div>
  <div v-else>
    <subject-list/>
  </div>
  <van-popup v-model:show="showPicker" round position="bottom">
    <van-picker
        :columns="budgetList"
        @cancel="showPicker = false"
        @confirm="onConfirm"
    />
  </van-popup>
</template>

<script lang="ts" setup>
import NavBar from '../../components/navBar/index.vue';
import OrganizationList from './component/oeganizationList/index.vue';
import SubjectList from './component/subjectList/index.vue'
import { ref,} from "vue";
const budgetType=ref('预算编制'); //切换类型
const budgetList=ref([{text:'预算编制',value: 'bianZhi'},{text:'预算科目',value: 'keMu'}]); // 预算编制类型
const showPicker=ref(false); // 切换编制类型
const  onConfirm=(value:any)=>{
  budgetType.value=value.selectedOptions[0].text;
  showPicker.value=false;
} //切换预算编制

</script>

<style scoped>

</style>
