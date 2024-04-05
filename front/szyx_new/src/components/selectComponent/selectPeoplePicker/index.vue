<!--
  选择人员组件（可多选）
  使用方法如下： 可参考差旅申请（同行人员）
  注：该模块可以通过传入的数目对可选的人数进行限制
-->
<template>
  <div class="selectPersonPicker">
    <van-sticky>
      <NavBar  title="选择人员" left-arrow left-text="返回" @left-event="$emit('selectOver',[])"  />
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
      <van-index-bar :index-list="value==''?indexList:[]" :offset-bottom="90"  :sticky="false">
        <div v-for="item in  arrFilter" :key="item.type">
          <div :id="'md'+item.type">
            <van-index-anchor :index="item.type" v-if="value==''"  />
            <div :class="'content'+ item.type">
              <van-cell v-for="data in item.list" clickable :key="data.empcode" border @click="selectUserList(data)" >
                <template #title>
                  <van-row justify="start" align="center">
                    <van-col :span="4">
                      <van-checkbox v-model="data.checked"  ></van-checkbox>
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
      <van-row class="mt10 ml10">
        <van-col span="16">
          <div style="height: 80px;overflow-x:scroll">
           <span v-for="item in checkUserList">
             <van-button type="default" size="small" plain class="mr5 mt5 mb5"  >{{item.empName}}</van-button>
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
import NavBar from '../../../components/NavBar.vue'
let userInfoData: any = inject("userInfo"); //获取审批数组
// @ts-ignore
import Pinyin from 'js-pinyin';
import {showFailToast} from "vant/lib/toast/function-call";
const value = ref('');
const userList=<any>ref([]); // 用户列表
const staticUserList=<any>ref([]); // 用户列表
const indexList=<any>ref([]); // 整理过后的用户列表
let checkUserList=<any>ref([]); // 选中的员工
const isAllPass=ref(''); //或审1 会审2
const emit=defineEmits(['selectOver']) // 返回方法
interface Props{};
const props=defineProps<{
  personType:String,// 使用类型
  list:[] // 回传list
  maxNumber:Number,
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

// 提交时方法处理
const SelectUserOver =()=>{
  let selectList:any=[];
  selectList=checkUserList.value;
  emit('selectOver',selectList);
  isAllPass.value='';
  value.value='';
  checkUserList.value=[];
  userList.value=[];
  resetUser();
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
    // 此处利用设置的最大数目进行判断
    let num;
    if(props.maxNumber){
      num=props.maxNumber;
    }else{
      num=5
    }
  // 判断是否超过五个人
    if(checkUserList.value.length<num){
      data.checked=true;
      checkUserList.value.push(data);
    }else{
      data.checked=false;
      showDialog({
        message: `最多选${num}人`,
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
  // 判断用户数组是否有数据，没有再去调用，以减少接口调用
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
.selectPersonPicker{
  height: 100%;
  display: flex;
  flex-direction: column;
}
.bottomBar{
  position: fixed;
  bottom:0;
  width: 100%;
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
.listBody{
  overflow: auto;
}
</style>
