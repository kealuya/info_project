<!--
  选择审批人组件（可多选）
  使用方法如下： 可参考申请单模块下考勤申请
-->
<template>
  <div class="selectPersonPicker">
      <van-sticky>
        <NavBar  title="选择审批人" left-arrow left-text="返回" @left-event="escOver"  />
        <form action="/">
          <van-search
              v-model="value"
              show-action
              placeholder="请输入中文名或工号"
              @search="onSearch"
          />
        </form>
      </van-sticky>

   <div class="listBody">
     <van-index-bar :index-list="value==''?indexList:[]" :offset-bottom="90" :sticky="false" >
       <div v-for="item in arrFilter" :key="item.type">
         <div :id="'md'+item.type">
           <van-index-anchor :index="item.type" />
           <div :class="'content'+ item.type">
             <van-cell v-for="data in item.list" clickable :key="data.empcode" border @click="selectUserList(data)" >
               <template #title>
                 <van-row justify="start" align="center">
                   <van-col :span="4">
                     <van-checkbox v-model="data.checked"></van-checkbox>
                   </van-col>
                   <van-col :span="8">
                     <span>{{data.empName}}</span>
                   </van-col>
                 </van-row>
               </template>
               <template #value>
                  <span class="mr10">{{data.jobName}}</span>
               </template>
             </van-cell>
           </div>
         </div>

       </div>
     </van-index-bar>
   </div>
   <div class="bottomBar">
     <van-row class="mt10 ml10" justify="center">
       <van-col span="16" class="tc">
         <div class="approveBottomGroup">
           <div :class="isAllPass==''?'btnChecked':'btnNoChecked'" @click="oneReview()" class="blc">单审</div>
           <div :class="isAllPass=='1'?'btnChecked':'btnNoChecked'" @click="orReview()">或审</div>
           <div :class="isAllPass=='2'?'btnChecked':'btnNoChecked'" @click="jointReview()"  class="brc">会审</div>
         </div>
       </van-col>
     </van-row>

     <van-row class="mt10 ml10">
       <van-col span="16">
         <div style="height: 100px;overflow-x:scroll">
           <span v-for="item in checkUserList">
             <van-button type="default" size="small" plain class="mr5 mt5"  >{{item.empName}}</van-button>
           </span>

         </div>

       </van-col>
       <van-col span="8" class="tr">
         <div class="mr10">
           <van-button type="primary"  @click="SelectUserOver">确认({{checkUserList.length}})</van-button>
         </div>
       </van-col>
     </van-row>
   </div>
  </div>
</template>

<script lang="ts" setup>
import {inject, ref, onMounted, defineEmits, onBeforeUpdate,computed} from "vue";
import {showDialog, Toast,} from "vant";
import {getEmpList} from "../../../services/current";
import NavBar from '../../NavBar.vue'
let userInfoData: any = inject("userInfo"); //获取审批数组
// @ts-ignore
import Pinyin from 'js-pinyin';
const value = ref('');
const userList=<any>ref([]); // 用户列表
const staticUserList=<any>ref([]); // 用户列表
const indexList=<any>ref([]);
let checkUserList=<any>ref([]); // 选中的员工
const isAllPass=ref(''); //或审1 会审2
const emit=defineEmits(['selectOver'])
interface Props{};
const props=defineProps<{
  list:[]
}>()
// 模糊查询
const arrFilter=computed(()=>{
  if(value.value==''){
    return userList.value
  }else{
    let nameList=staticUserList.value.filter((a:any)=> a.empName.includes(value.value));
    let codeList=staticUserList.value.filter((a:any)=> a.empCode.includes(value.value));
    let data={
      type:'*',
      list:codeList.concat(nameList)
    }
    return [data];
  }
})
// 或审
const oneReview=()=>{
  if(checkUserList.value.length<2){
    isAllPass.value='';
  }else{
    showDialog({
      message: '选择审批人人数等于1的时候才可以将审批流类型选为单审',
    }).then(() => {
      // on close
    });
  }

}
// 或审
const orReview=()=>{
  if(checkUserList.value.length>1){
    isAllPass.value='1';
  }else{
    showDialog({
      message: '选择审批人人数大于1的时候才可以将审批流类型选为或审',
    }).then(() => {
      // on close
    });
  }
}
// 会审
const jointReview=()=>{
  if(checkUserList.value.length>1){
    isAllPass.value='2'
  }else{
    showDialog({
      message: '选择审批人人数大于1的时候才可以将审批流类型选为会审',
    }).then(() => {
      // on close
    });
  }
}
const escOver=()=>{
  isAllPass.value='';
  checkUserList.value=[];
  value.value='';
  userList.value=[];
  resetUser();
  emit('selectOver',[]);
}
// 提交时方法处理
const SelectUserOver =()=>{
  let selectList:any=[];
  if(checkUserList.value.length>1 && isAllPass.value==''){
    showDialog({
      message: '选择人员数量大于1不可以选择单审',
    }).then(() => {
      // on close
    });
  }else{
    if(isAllPass.value=='1'){
      checkUserList.value.forEach((item: any)=>{
        return selectList.push({...item, isAllPass: '1'});
      })
      emit('selectOver',selectList);
      isAllPass.value='';
      checkUserList.value=[];
      value.value='';
      userList.value=[];
      resetUser();
    }else if(isAllPass.value=='2'){
      checkUserList.value.forEach((item: any)=>{
        return selectList.push({...item, isAllPass: '2'});
      })
      emit('selectOver',selectList);
      isAllPass.value='';
      value.value='';
      checkUserList.value=[];
      userList.value=[];
      resetUser();
    }else{
      selectList=checkUserList.value;
      emit('selectOver',selectList);
      value.value='';
      isAllPass.value='';
      checkUserList.value=[];
      userList.value=[];
      resetUser();
    }

  }
}
const onSearch = (val:any) => new Toast(val);

// 调用人员列表
const getUserList = async () => {
  let payload = {
    'equipmentType': "pc",
    "keyWord":"",
    'page':1,
    'pageSize':10000,
  }
  let result: any = await getEmpList(payload);
  staticUserList.value=result.data.list;
  resetUser();
}
const resetUser=()=>{
  if(checkUserList.value.length==0){
    staticUserList.value.forEach((item: { empName: any; })=>{
      return userList.value.push({...item, pinyin: Pinyin.getFullChars(item.empName),checked:false});
    })
    userList.value=filterCity(userList.value);
  }else{
    if(checkUserList.value[0].isAllPass){
      isAllPass.value=checkUserList.value[0].isAllPass;
    }else{
      isAllPass.value="";
    }
    userList.value=[];
    staticUserList.value.forEach((item: { empName: any; empCode:any})=>{
      let key=inList(item.empCode);
      return userList.value.push({...item, pinyin: Pinyin.getFullChars(item.empName),checked:key});

    })
    userList.value=filterCity(userList.value);
  }

}
// 判断是否存在于数组
const inList=(key:any)=>{
  let label=checkUserList.value.find((item: { empCode: any; }) => {
    return key==item.empCode;
  })
 if(label==undefined){
   return false
 }else{
   return  true
 }
}

// 选择用户
const selectUserList=(data:any)=>{
  /**
   *  选择用户组件
   *  1.判断是否超过五个人
   *  2.判断是数组中是否存在该属于
   *  */
      // 此处是获取数据以及其数组下标
  let index= checkUserList.value.findIndex((item: { empCode: any; }) => {
        return data.empCode==item.empCode;
      });
  // 判断是否数组中是否有下标 如果不存在的话会返回-1
  if(index==-1){
    // 判断是否超过五个人
    if(checkUserList.value.length<5){
      data.checked=true;
      checkUserList.value.push(data);
    }else{
      data.checked=false;
      showDialog({
        message: '选择审批人数量不超过5位',
      }).then(() => {
        // on close
      });
    }
  }else{
    // 二次点击实现删除
    data.checked=false;
    checkUserList.value.splice(index,1);
  }
};
// 遍历用户
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
  indexList.value=indexKeyList;
  return newCities;
}

onMounted(()=> {
  checkUserList.value=[];
  isAllPass.value='';
  userList.value=[];
  if(staticUserList.value.length==0){
    getUserList();
  }
})
onBeforeUpdate(() => {
  // 文本内容应该与当前的 `count.value` 一致
  checkUserList.value=[];
  if(props.list){
    checkUserList.value=props.list;
    userList.value=[];
    resetUser();
  }
})
</script>

<style scoped>
>>>.van-popup .van-popup--center van-toast .van-toast--middle .van-toast--fail{
  z-index:9999!important;
}
.selectPersonPicker{
  height: 100%;
  display: flex;
  flex-direction: column;
}
.bottomBar{
  position: fixed;
  bottom:0;
  width: 100%;
  height: 120px;
  border-top: 0.5px solid #c9c9c9;
  background: #ffffff;
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
.btnNoChecked{
  background: #c9c9c9;
}
.btnChecked{
  background: #0088ff;
}
.blc{
  border-top-left-radius: 5px;
  border-bottom-left-radius: 5px;
}
.brc{
  border-top-right-radius: 5px;
  border-bottom-right-radius: 5px;
}
.listBody{
  overflow: auto;
}
</style>
