<template>
  <div>
    <div>
      <van-row class="p10" justify="start" align="center">
                <van-col span="6" class="tc pt10" v-for="item in IconList" @click="goPath(item.path,item.code)">
                  <van-image :src="item.icon" width="50px" height="50px"/>
                  <div class="f12">{{item.title}}</div>
                </van-col>
      </van-row>
    </div>
  </div>
</template>

<script lang="ts" setup>
// 更多
import aboutIcon from '../../../assets/icon/icon_08.png'
import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {showSuccessToast} from "vant";
import {staticIconListA} from "../../about/ts/iconList";
const router = useRouter(); // 路由
// 根据key跳转到对应页
const goPath=(path:any,code:any)=>{
  // 跳转至飞机购票页
  if (path==''){
    showSuccessToast('暂未开放');
    return false
  }else if(path=='expressDeliveryOrCarPage'){
    router.push({name:path,params:{applicationType:code}})
  }else{
    router.push(path)
  }
}
// 图标数组
let IconList:any=ref([]);
// 处理首页功能
const indexIconFormat = (myIconList:any) => {
  let codeList = myIconList;
  let comeIconList = [];
  for (let i = 0; i < codeList.length; i++) {
    let codeItem = staticIconListA.find((e: any) => codeList[i] == e.code);
    comeIconList.push(codeItem);
  }
  IconList.value = comeIconList;
  IconList.value.push(aboutObject as never)
}
// 更多功能
let aboutObject={icon:aboutIcon,title:'更多',code:'about',path:'about'}
onMounted(() => {
  let myIconList:any=sessionStorage.getItem('myIconCodeList');
  myIconList = JSON.parse(myIconList);
  indexIconFormat(myIconList);
});
</script>

<style scoped>

</style>
