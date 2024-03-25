<!--
  选择城市组件（单选）
  使用方法如下： 可参考申请单模块下的差旅申请
-->
<template>
  <div class="selectPersonPicker">
    <van-sticky>
      <Navbar title="选择城市"  left-arrow left-text="返回" @leftEvent="$emit('selectOver',{})" />
      <form action="/">
        <van-search
            v-model="value"
            show-action
            placeholder="请输入城市名称"
        />
      </form>
    </van-sticky>
    <div class="listBody">
    <!--  城市数组列表展示    -->
      <van-index-bar :index-list="indexList" :offset-bottom="90" >
        <!--  城市热门城市    -->
        <div :id="'md'+'hot'">
          <van-index-anchor index="*" >* 热门城市</van-index-anchor>
          <div :class="'content'+ 'hot'">
            <van-row class="ml15">
                <span v-for="item in hotCityList">
                   <van-col class="hotCityCard f12"  @click="selectCity(item)">{{item.label}}</van-col>
                </span>

            </van-row>
          </div>
        </div>
        <div v-for="item in arrFilter" :key="item.type">
          <div :id="'md'+item.type">
            <van-index-anchor :index="item.type" />
            <div :class="'content'+ item.type">
              <van-cell v-for="data in item.list" clickable :key="data.value" border @click="selectCity(data)" >
                <template #title>
                  <van-row justify="start" align="center">
                    <van-col :span="15">
                      <span>{{data.label}}</span>
                    </van-col>
                  </van-row>
                </template>
              </van-cell>
            </div>
          </div>
        </div>
      </van-index-bar>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, defineEmits, computed,} from "vue";
import './ts/cityData' // 引入城市列表数组
import Navbar from '../../NavBar.vue';
// @ts-ignore
import {flightList, hotelList, trainList} from "./ts/cityData"; // 导入会用到的城市数组
const value = ref(''); // 搜索项
const cityList=<any>ref([]); // 城市列表
const staticCityList=<any>ref([]); // 减少调用接口时会用的静态存储城市数组
const indexList=<any>ref([]); // 字母索引数组
// let checkCityList=<any>ref([]); // 选中的城市（暂时不用，等待优化功能时可能会用）
const hotCityList=<any>ref([]); // 热门城市数组
const emit=defineEmits(['selectOver']); // emit事件方法
interface Props{}
const props=defineProps<{
  type:String,// 使用类型
}>()
// 通过计算属性处理数组
const arrFilter=computed(()=>{
  if(value.value==''){
    // 根据传递来的出行类型显示对应的可供选择的列表
    getcityList();
    return cityList.value
  }else{
    let nameList=staticCityList.value.filter((a:any)=> a.label.includes(value.value));
    let data={
      type:'*',
      list:nameList
    }
    return [data];
  }
})
// 调用人员列表
const getcityList = async () => {
  // 根据传递来的出行类型显示对应的可供选择的列表
  if(props.type=='train'){
    staticCityList.value=trainList;
  }
  else if(props.type=='air'){
    staticCityList.value=flightList;
  }else{
    staticCityList.value=hotelList;
  }
  resetUser();
}
// 此处是从选择人员页粘过来的，日后可作为回显数据列表做取选中使用（锦上添花）
const resetUser=()=>{
    hotCityList.value=[];
    cityList.value=[];
    indexList.value=[];
    let cityArray:any=[];
    staticCityList.value.forEach((item: { label: any; tagIndex:any})=>{
        if(item.tagIndex=='★'){
          hotCityList.value.push(item);
        }
       cityArray.push({...item});
    })
  // 用于右侧索引
  cityList.value=filterCity(cityArray);
}
// 处理城市数组
const filterCity = (list: any[]) => {
  let letterArray = [];
  for (let i = 65; i < 91; i++) {
    letterArray.push(String.fromCharCode(i));
  }
  let newCities: any[] = []
  let indexKeyList:any[]=[];
  letterArray.forEach(item => {
    newCities.push({
      type: `${item}`,
      list: list.filter(item1 => item1.pinyin.substring(0, 1).toUpperCase() === `${item}`)
    })
  })
  newCities = newCities.filter(item => item.list.length)
  newCities.forEach(item => {
    if(item.list.length!=0){
      indexKeyList.push(item.type);
    }
  })
  indexKeyList.unshift('*');
  indexList.value=indexKeyList;
  return newCities;
}
// 选择用户回显至原先的数组
const selectCity=(data:any)=>{
  emit('selectOver',data);
  cityList.value=[];
};
</script>

<style scoped>
.hotCityCard{
  margin: 10px;
  padding: 8px;
  text-align: center;
  border-radius: 8px;
  border:#9c9c9c 1px solid;
}
.selectPersonPicker{
  height: 100%;
  display: flex;
  flex-direction: column;
}
.approveBottomGroup div{
  display: inline-block;
  width: 70px;
  text-align: center;
  padding-top: 10px;
  color:#FFFFFF;
  font-size: 14px;
  height: 30px;
}
.listBody{
  overflow: auto;
}
</style>
