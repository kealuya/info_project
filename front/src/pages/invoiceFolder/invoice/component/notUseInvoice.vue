<template>
  <van-pull-refresh v-model="refreshing" @refresh="onRefresh" style="min-height:80vh">
    <template #success>
      <span v-if="refreshingOnly==true"><van-icon name="success"/> 刷新成功</span>
    </template>
    <div v-if="loading==false && list.length==0" style="height: 500px">
      <van-row justify="center" align="center" class="tc">
        <div>
          <img style="height:40vh;width:80vw;object-fit: contain" src="../../../../assets/img/zanwupiaoju.png"
               alt="暂无可报销的发票" v-cloak/>
        </div>
      </van-row>
      <div style="color: rgba(23,26,29,0.40);text-align: center">暂无可报销的发票</div>
      <van-row justify="center" align="center" class="tc mt5">
        <div>
          <van-button square type="primary" @click="onRefresh" class="mt10"
                      style="width: 90px; padding: 10px 5px;border-radius: 8px">刷新
          </van-button>
        </div>
      </van-row>

    </div>
    <van-list
        v-model="loading"
        v-else
        :finished="finished"
        loading-text="正在加载中请稍后"
        finished-text="没有更多了"
        @load="onLoad"
        :offset="10"
        style="min-height:80vh"
    >
      <div class="ticket_content" v-for="(item) in list" :key="item">
        <van-swipe-cell>
          <airplane-component v-if="item.type=='5'" :get-data="item" @select-invoice="checkInvoice"/>
          <train-component v-else-if="item.type=='8'" :get-data="item" @select-invoice="checkInvoice"/>
          <invoice-component v-else :get-data="item" @select-invoice="checkInvoice"/>
          <template #right>
            <van-button square text="删除" type="danger" class="delete-button mr5" @click="delInv(item)"/>
          </template>
        </van-swipe-cell>
        <!--                <div>发票: {{ item.invName }}</div>-->
      </div>
    </van-list>
  </van-pull-refresh>
</template>

<script setup lang="ts">
import {getInvoiceWbList, deleteInv} from "../../../../services/invoice";
import {ref,  inject, defineEmits} from "vue";
// @ts-ignore
import InvoiceComponent from "./invoiceComponent.vue";
// @ts-ignore
import AirplaneComponent from "./airplaneComponent.vue";
// @ts-ignore
import TrainComponent from "./trainComponent.vue";

const emit = defineEmits(['selectInvoice']); // emit事件方法
import {showConfirmDialog, showSuccessToast} from 'vant';
import {showFailToast} from "vant/lib/toast/function-call";

let list = ref<any>([]);//加载loading
let loading = ref(true); //加载效果
let refreshingOnly = ref(false); //加载效果只有下拉变化
let finished = ref(false); // 结束
const refreshing = ref(false); // 重新加载还是追加
let userInfoData: any = inject("userInfo");
let payload = {
  'companyCode': userInfoData.userInfo.corpCode,
  'corpCode': userInfoData.userInfo.corpCode,
  'empCode': userInfoData.userInfo.empCode,
  'row': '10',
  'page': '1',
  'zt': '0',
  'type': '',
}
const getInvoiceList = async (payload: any, code: any) => {
  let result: any = await getInvoiceWbList(payload);
  if (code == 'refresh') {
    list.value = [];
  }
  list.value.push(...result.data.data.data);
  if (list.value.length.toString() === result.data.data.count) {
    loading.value = false;
    finished.value = true;
  } else {
    let page = Number(payload.page) + 1;
    payload.page = page.toString();
    loading.value = false;
  }
  refreshing.value = false;
}
const onLoad = async (code: any) => {
  if (code == '1') {
    payload.page = "1";
  }
  if (refreshing.value) {
    return;
  }
  if (finished.value == true) {
    return
  }
  await getInvoiceList(payload, 'loading')
};
const refresh = async () => {
  if (refreshing.value) {
  }
  if (finished.value == true) {
    return
  }
  payload.page = "1";
  await getInvoiceList(payload, 'refresh')
}
const onRefresh = (code: any) => {
  payload.page = "1";
  if (code != '1') {
    refreshing.value = true;
    refreshingOnly.value = true;
  } else {
    refreshing.value = false;
    refreshingOnly.value = false;
  }
  finished.value = false;
  loading.value = true;
  refresh();
};

//删除发票
const delInv = async (item: any) => {
  showConfirmDialog({
    title: '您确定要删除此发票吗',
    confirmButtonColor: '#4BA3FB'
  }).then(async () => {
    let payload = {
      corpCode: userInfoData.userInfo.corpCode,
      dzfph: item.dzfph,
      empCode: userInfoData.userInfo.empCode
    }
    let res: any = await deleteInv(payload);
    console.log(res);
    if (res.success === true) {
      showSuccessToast('删除成功');
      await onRefresh('0')
    } else {
      showFailToast('删除失败');
    }
  })

}
// 将选中的发票传递至上一级父组件
const checkInvoice = (item: any) => {
  emit('selectInvoice', item);
}


defineExpose({
  getInvoiceList, onLoad, onRefresh, refresh, payload, list
})
</script>

<style lang="less" scoped>
.van-button {
  height: 100%;
  border-radius: 12px 12px 12px 12px;
  font-size: 15px;
  width: 70px;
}
</style>
