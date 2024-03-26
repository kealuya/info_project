<template>
  <div>
    <div>
      <NavBar :leftText="'返回'" :leftArrow="true" :title="'账户管理'" @leftEvent="back()"></NavBar>
    </div>
    <div>
      <div style="display: flex">
        <van-cell title="账户类型" style="color: #969799;font-weight: 600" is-link arrow-direction="down"
                  :value="bankType" @click.stop="show=true"/>
        <van-action-sheet v-model:show="show" :actions="actions" @select="onSelect"/>
      </div>
      <van-pull-refresh v-model="refreshing" @refresh="onRefresh"
                        :style="{'height':realHeight+'px','overflow':'scroll'}">
        <van-list v-model:loading="loading" :finished="finished" finished-text="没有更多了" offset="100" @load="onLoad"
                  >
          <div align="center" v-for="(item,index) in accountBookList" :key="item"
               @click.stop="!isComponents? queryDetail(item): selectItem(item,index)">
            <div :class="item.isdefaultpayee=='false'?'cardCss':'cardDefaultCss'">
              <van-row class="mt5" style="height: 20px">
                <van-col span="17" align="left">
                  <div class="ellipsis">{{ item.payeeName }}</div>
                </van-col>
                <van-col span="7" v-if="bankType=='个人账户' && !isComponents ">
                  <van-button type="default" size="mini" v-if="item.isdefaultpayee=='true'"
                              style="background-color: unset;color: #FFFFFF"
                              @click.stop="cancelDefault(item)">
                    <van-icon name="edit"/>
                    取消默认
                  </van-button>
                  <van-button type="default" size="mini"
                              v-if="item.isdefaultpayee=='false'"
                              style="background-color: unset;color: #FFFFFF"
                              @click.stop="setUpDefault(item)">
                    <van-icon name="edit"/>
                    设置默认
                  </van-button>

                </van-col>
                <van-col span="7" v-if="isComponents">
                  <van-icon name="passed" size="25" v-if="item.selectIconStatus"/>
                </van-col>
              </van-row>
              <van-row class="mb10 pt10">
                {{ item.bankcardnum }}
              </van-row>
              <div class="pt10 pb10">
                <van-row>
                  <van-col span="4">{{ item.type == 'personalAccount' ? '个人账户' : '对公账户' }}
                  </van-col>
                  <van-col span="1">
                    <van-divider vertical/>
                  </van-col>
                  <van-col span="6">
                    <div class="ellipsisCss">
                      {{ item.bankaccountname }}
                    </div>
                  </van-col>
                  <van-col span="1">
                    <van-divider vertical/>
                  </van-col>
                  <van-col span="12">
                    <div class="ellipsisCss">
                      {{ item.accountnetwork }}
                    </div>
                  </van-col>
                </van-row>
              </div>
            </div>
          </div>
        </van-list>
      </van-pull-refresh>
    </div>
    <div align="center" style="background: #FFFFFF;">
      <van-row align="center" class="addBtnCss">
        <van-button plain hairline type="primary" color="#969799" size="normal"
                    @click="addBank">
          <van-icon name="plus"/>
          添加账户
        </van-button>
      </van-row>
    </div>
    <!--        绑定银行卡弹窗-->
    <van-popup v-model:show="addDialogVisible" v-if="addDialogVisible" style="width:100vw;height:100vh"
               position="bottom">
      <BindBankCard @close="closeBindBankCard" v-if="addDialogVisible" :bankCardParams="params"/>
    </van-popup>
  </div>
</template>

<script setup lang="ts">
import {inject, ref, onMounted, defineProps, defineEmits} from 'vue'
import router from "../../router";
import {payeeList, updatePayeeInfo, getDefaultPay} from "../../services/manageAccounts"
import {showConfirmDialog, showSuccessToast, showImagePreview, showFailToast} from "vant";
import bg_card from '../../assets/img/bg_card.png'
import bg_card_check from '../../assets/img/bg_card_check.png'
import screening_01 from '../../assets/icon/screening_01.png'
import screening_02 from '../../assets/icon/screening_02.png'
import BindBankCard from './add.vue'
import moment from "moment";
import {userInfoData} from "../../store/index";

let userInfo: any = inject("userInfo");
const imgUrl = ref('https://file-report-store.obs.cn-north-4.myhuaweicloud.com/')
const accountBookList = ref([]);
const loading = ref(false);
const finished = ref(false);
const refreshing = ref(false);
const total = ref(0)
const page = ref(0)
const pageSize = ref(8)
const queryObj = ref({
  keyWord: "",
  expenseTypeName: "",
  occurrenceTime: "",
  state: "0",
})
const iconNameUrl = ref(screening_01)
const realHeight = window.document.documentElement.offsetHeight - 155
const show = ref(false)
const bankType = ref("个人账户")
const actions = [
  {name: '个人账户'},
  {name: '对公账户'},
];
const defaultBank = ref()
const isComponents = ref(false)
const selectBank = ref({})
const bankIndex = ref(-1)
const addDialogVisible = ref(false)
const params = ref({
  title: '添加账户',
  accountBookDetail: {}
})
const props = defineProps({
  isComponentsStatus: {
    type: Boolean,
  },
  selectBankMesssage: {
    type: Object,
  }
})
const emit = defineEmits(["selectBankCardInformation"])

onMounted(() => {
  let type = router.currentRoute.value.params.bankType
  isComponents.value = props.isComponentsStatus
  if (type) {
    bankType.value = type as never
    let ite = {
      name: bankType.value
    }
    onSelect(ite)
  }
})

function closeBindBankCard(e: any) {
  addDialogVisible.value = false
  page.value = 1;
  let obj = {
    name: e
  }
  onSelect(obj)
}

function selectItem(e: any, index: any) {
  if (bankIndex.value != -1) {
    accountBookList.value[bankIndex.value].selectIconStatus = false
  }
    bankIndex.value = index
    accountBookList.value[bankIndex.value].selectIconStatus = true
  setTimeout(() => {
    emit("selectBankCardInformation", e)
  }, 150)
}

const onSelect = (item: any) => {
  show.value = false;
  bankType.value = item.name
  page.value = 0
  // refreshing.value = true
  finished.value = false;
  accountBookList.value = []
  onLoad()
};

function onLoad() {
  page.value = page.value + 1;
  loading.value = true
  getDataMth()

}

const onRefresh = () => {
  // 清空列表数据
  finished.value = false;
  loading.value = true
  page.value = 0
  onLoad();
};

const resetFilter = () => {
  queryObj.value.keyWord = "",
      queryObj.value.expenseTypeName = "",
      queryObj.value.occurrenceTime = "",
      queryObj.value.state = "0"

}

async function getDataMth() {
  let list = accountBookList.value
  setTimeout(async () => {
    if (refreshing.value) {
      list = []
      refreshing.value = false
    }
    getDefaultPayMth().then(() => {
      if (page.value == 1 && bankType.value == '个人账户' && defaultBank.value != null) {
        if (props.isComponentsStatus == true && defaultBank.value.bankcardnum == props.selectBankMesssage.bankcardnum) {
          bankIndex.value = 0
          defaultBank.value.selectIconStatus = true
        }
        list.unshift(defaultBank.value as never)
      }
    }).then(async () => {
      let param = {
        corpCode: userInfo.userInfo.corpCode,
        empCode: userInfo.userInfo.empCode,
        keyWord: "",
        page: page.value,
        pageSize: pageSize.value,
        type: bankType.value == '个人账户' ? 'personalAccount' : 'publicAccount',
      }
      let res: any = await payeeList(param)
      if (res.code == 200) {
        if (isComponents.value == true) {
          res.data.records.map((item: any) => {
            item.selectIconStatus = false
            if (props.selectBankMesssage && item.bankcardnum == props.selectBankMesssage.bankcardnum) {
              item.selectIconStatus = true
            }
            if (item.isdefaultpayee == 'false') {
              list.push(item as never)
            }
          })
        } else {
          res.data.records.map((item: any) => {
            item.selectIconStatus = false
            if (item.isdefaultpayee == 'false') {
              list.push(item as never)
            }
          })
        }
        loading.value = false;
        total.value = res.data.total
        if (list.length >= total.value) {
          finished.value = true;
        }
        accountBookList.value = list
      } else {
        list = []
      }
    })
  }, 1000)
}

function changeState(e: any) {
  queryObj.value.state = e
}

function changeType(e: any) {
  queryObj.value.expenseTypeName = e
}

function previewImage(item: string) {
  showImagePreview({
    images: [item],
    showIndex: false,
  })
}

function cancelDefault(e: any) {
  showConfirmDialog({
    title: '您确定要取消该默认收款账户吗?'
  }).then(async () => {
    updatePayeeInfoMth(e, false)
    refreshing.value = true
    // accountBookList.value = []
    page.value = 1
    getDataMth()
  })
}

function setUpDefault(e: any) {
  showConfirmDialog({
    title: '您确定要设置该账户为默认收款账户吗?'
  }).then(async () => {
    updatePayeeInfoMth(e, true)
    // accountBookList.value = []
    refreshing.value = true
    page.value = 1
    getDataMth()
  })
}

async function updatePayeeInfoMth(item: any, state: any) {
  let dateTime = new Date()
  let params = {
    empcode: userInfo.userInfo.empCode,
    type: bankType.value == '个人账户' ? 'personalAccount' : 'publicAccount',
    id: item.id,
    bankcardnum: item.bankNumber,
    bankaccountname: item.bankName,
    city: item.regionName,
    region: item.cityName,
    cityCode: item.cityCode,
    accountnetwork: item.openingBankBranch,
    isdefaultpayee: state,
    payeeName: item.nameOfPayee,
    createtime: moment(dateTime).format("YYYY-MM-DD"),
    corpCode: userInfo.userInfo.corpCode
  }
  let res: any = await updatePayeeInfo(params)
  if (res.code == 200) {
    showSuccessToast("设置成功")
  } else {
    showFailToast(res.message)
  }
}

function submitFillter() {
  if (queryObj.value.state == '0' && queryObj.value.expenseTypeName == '') {
    iconNameUrl.value = screening_01
  } else {
    iconNameUrl.value = screening_02
  }
  page.value = 0
  refreshing.value = true
  show.value = false
  onLoad();
}

function showFilterMth() {
  show.value = true
}

function back() {
  if (isComponents.value) {
    emit("selectBankCardInformation", props.selectBankMesssage)
  } else {
    router.push({
      name: 'user'
    })
  }
}

function submitAnExpenseAccount(item: any) {
  let params = {
    identifying: 'accountBook',
    bill: JSON.stringify(item)
  }
  router.push({
    name: 'dailyReimbursement',
    params: params,
  })
}

function addBank() {
  let param = {
    title: '添加账户',
    accountBookDetail: {}
  }
  params.value = param
  addDialogVisible.value = true
}

function queryDetail(e: any) {
  let param = {
    title: '编辑账户',
    accountBookDetail: e
  }
  params.value = param
  addDialogVisible.value = true

}

async function getDefaultPayMth() {
  let params = {corpCode: userInfo.userInfo.corpCode, empCode: userInfo.userInfo.empCode,}
  let res: any = await getDefaultPay(params)
  if (res.code == 200) {
    defaultBank.value = res.data
  }
}

</script>

<style lang="less" scoped>
.fc {
  color: #1a1a1a;
}

.fw {
  font-weight: 800;
}

.iconCss {
  width: 45px;
  height: 45px;
  /*float: left;*/
  /*margin-left: 20%;*/
}

.ellipsis {
  width: 230px; /*宽度一定要设置*/
  overflow: hidden; /*文本超出隐藏*/
  text-overflow: ellipsis; /*文本超出显示省略号*/
  white-space: nowrap; /*超出的空白区域不换行*/
  text-align: left
}

.ellipsisCss {
  white-space: nowrap; /*超出的空白区域不换行*/
  overflow: hidden; /*文本超出隐藏*/
  text-overflow: ellipsis; /*文本超出显示省略号*/
  text-align: left

}

.cardCss {
  background: url("../../assets/img/bg_card.png");
  /*从左到右，从上到下的缩放比例，单位px则为px缩放*/
  background-size: 100% 100%;
  margin-top: 5px;
  border-radius: 3px;
  width: 93%;
  padding-top: 10px;
  padding-left: 10px;
  color: #FFFFFF;
  font-size: 14px;
}

.cardDefaultCss {
  background: url("../../assets/img/bg_card_check.png");
  /*从左到右，从上到下的缩放比例，单位px则为px缩放*/
  background-size: 100% 100%;
  margin-top: 5px;
  border-radius: 3px;
  width: 93%;
  padding-top: 10px;
  padding-left: 10px;
  color: #FFFFFF;
  font-size: 14px;
}

.addBtnCss {
  padding-top: 10px;
  padding-bottom: 10px;
  padding-left: 10px;
  font-weight: 600;
}

.addBtnCss .van-button {
  width: 98%;
  font-size: 14px;
}


</style>
