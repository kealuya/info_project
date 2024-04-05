<!--
  选择费用信息组件（单选）
  使用方法如下： 可参考日常报销
-->
<template>
  <div class="selectPersonPicker">
    <van-sticky>
      <NavBar   title="选择账本" left-arrow right-text="确定" @leftEvent="escOver" @rightEvent="costCenterOK" />
    </van-sticky>
    <div>
      <van-pull-refresh v-model="refreshing" @refresh="onRefresh" style="background:#f1f2f3; min-height: 100vh;">
        <template #success>
          <span v-if="refreshingOnly==true"><van-icon name="success" /> 刷新成功</span>
        </template>
        <div v-if="loading==false && list.length==0" style="height: 500px">
          <van-row justify="center" align="center" class="tc">
            <div>
              <img style="height:50vh;width:80vw;object-fit: contain" src="../../../assets/img/zanwupiaoju.png" v-cloak/>
            </div>
          </van-row>
          <div style="color: rgba(23,26,29,0.40);text-align: center">暂无可以用的账本信息</div>
          <van-row justify="center" align="center" class="tc mt5">
            <div>
              <van-button square plain type="primary" @click="onRefresh" style="padding: 5px 5px;border-radius: 8px">刷新
              </van-button>
            </div>
          </van-row>

        </div>
        <van-list
            v-model="loading"
            v-else
            :finished="finished"
            loading-text="正在加载中请稍后"
            :immediate-check="false"
            finished-text="没有更多了"
            @load="onLoad"
            :offset="10"
        >
          <div class="ticket_content" v-for="(item,index) in list" :key="item">
            <div  @click="selectAccountList(item,index)">
              <van-row justify="space-between" align="center">
                <van-col :span="2" offset="1">
                  <van-checkbox v-model="item.checked"></van-checkbox>
                </van-col>
                <van-col :span="21">
                  <div class="accountBookCard">
                    <van-row justify="space-between" align="center">
                      <van-col class="bold">{{item.reason}}</van-col>
                      <van-col class="mr10 bold orange">¥ {{item.money}}</van-col>
                    </van-row>
                    <div class="f14 mt5">
                      <van-row align="center">
                        <van-col>
                         <van-row align="center">
                           <van-icon :name="vectorIcon" />
                           <span class="ml5"> {{item.expenseTypeName}}</span>
                         </van-row>
                        </van-col>
                        <van-col offset="1">
                          <van-row align="center">
                            <van-icon :name="timeIcon" />
                            <span class="ml5">{{item.createTime}}</span>
                          </van-row>
                        </van-col>
                      </van-row>
                    </div>
                  </div>
                </van-col>
              </van-row>
            </div>
          </div>
        </van-list>
      </van-pull-refresh>
      <div>
        <!--  浮动添加按钮      -->
        <van-sticky position="bottom">
          <div class="sticky" @click="onClick" >
            <van-icon name="records" color="#0088ff" size="30" />
          </div>
        </van-sticky>
      </div>
    </div>
    <load-component  v-if="loading" />
  </div>
</template>

<script lang="ts" setup>
import {getBillList} from "../../../services/accountBook";
import vectorIcon from '../../../assets/icon/Vector.png';
import {useRouter} from "vue-router";
import {inject, ref, onMounted, defineEmits,} from "vue";
import NavBar from '../../../components/NavBar.vue'
import  LoadComponent from '../../loading/index.vue'
import timeIcon from '../../../assets/icon/baoxiao_icon_02.png';
import Navbar from '../../../components/NavBar.vue'
const router = useRouter(); // 用于监控路由在此处声明一下路由参数
let userInfoData: any = inject("userInfo"); //获取审批数组
let list = ref<any>([]);//加载loading
let loading = ref(true); //加载效果
let refreshingOnly = ref(false); //加载效果只有下拉变化
let finished = ref(false); // 结束
const refreshing = ref(false); // 重新加载还是追加
const emit=defineEmits(['selectOver']) //emit 方法声明
// 提交获取账本列表的数组
let payload = {
  "empCode":userInfoData.userInfo.empCode,
  "corpCode":userInfoData.userInfo.corpCode,
  "page":1,
  "pageSize":10,
  "keyWord":"",
  "equipmentType":""
}
interface Props{};
const props=defineProps<{
  selectList:[]
}>()
// 用于存放选中的账单列表
let selectAccountBookList=ref([]);
// 退出清空
const escOver=()=>{
  selectAccountBookList.value=[];
  emit('selectOver',[]);
}
// 分页加载
const onLoad = async (code:any) => {
  if(code=='1'){
    payload.page = 1;
  }
  if (refreshing.value) {
    return;
  }
  if(finished.value==true){
    return
  }
  await getCostCenterList(payload,'loading')
};
// 首次进入刷新
const refresh=async()=>{
  if (refreshing.value) {
  }
  if(finished.value==true){
    return
  }
  payload.page = 1;
  loading.value = true;
  await getCostCenterList(payload,'refresh')
}
// 刷新处理
 const onRefresh = (code:any) => {
  payload.page = 1;
  refreshing.value = true;
  if(code!='1'){
    refreshingOnly.value = true;
  }else{
    refreshingOnly.value = false;
  }
  finished.value = false;
  loading.value = true;
  refresh();
};
// 选中后提交
const costCenterOK=()=>{
  emit('selectOver',selectAccountBookList.value);
  selectAccountBookList.value=[];
}
// 点击跳转到添加账本页
const onClick = () => {
  router.push({
    name: 'book',
    query: {
      key: 'add',
    }
  })
};
// 选择用户
const selectAccountList=(data:any,index:any)=>{
  /**
   *  选择账本组件
   *  */
      // 此处是获取数据以及其数组下标
  let indexNum= selectAccountBookList.value.findIndex((item: { billId: any; }) => {
        return data.billId==item.billId;
      });
  if(indexNum!=-1){
    list.value[index].checked=false;
    selectAccountBookList.value.splice(index,1);
  }else{
    list.value[index].checked=true;
    selectAccountBookList.value.push(data as never);
  }
};
// 调用账本列表
const getCostCenterList=async(payload:any,code:any)=>{
  let result: any = await getBillList(payload);
  if(code=='refresh'){
    list.value = [];
  }
  list.value.push(...result.data.records);
  listFormat()
  if (list.value.length === result.data.total){
    finished.value = true;
    loading.value = false;
  } else {
    let page = Number(payload.page) + 1;
    payload.page = page;
    loading.value = false;
  }
  refreshing.value = false;
}
// 将发票列表赋初值
const listFormat=()=>{
  if(props.selectList.length!=0){
    let billKeyList;
    billKeyList=props.selectList;
    selectAccountBookList.value=billKeyList as never;
    let billList:any=[];
    list.value.forEach((item:any)=>{
      let key=inList(item.billId);
      return billList.push({...item,checked:key});
    })
    list.value=billList;
  }

}
// 判断是否存在于数组
const inList=(key:any)=>{
  let label=selectAccountBookList.value.find((item: { billId: any; }) => {
    return key==item.billId;
  })
  if(label==undefined){
    return false
  }else{
    return  true
  }
}
onMounted(()=> {
  onRefresh('1');
})
// 放出外部可调用的方法
defineExpose({
  onLoad, onRefresh,refresh, payload, list
})
</script>

<style scoped>
.selectPersonPicker{
  /*height: 100%;*/
  display: flex;
  flex-direction: column;
}
.accountBookCard{
  margin: 10px;
  padding: 10px;
  border: 1px solid rgba(0,128,255,0.4);
  background: rgba(0, 128, 255, 0.15);
  border-radius:12px;
}
.sticky{
  position: fixed;
  right: 10px;
  bottom: 100px;
  width:40px;
  height: 30px;
  text-align: center;
  padding: 8px;
  border: 1px solid #0088ff;
  border-radius: 20px;
  background: #FFFFFF;
}
</style>
