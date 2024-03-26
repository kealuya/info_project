<template>
  <div>
    <div>
      <NavBar :leftText="'返回'" :leftArrow="true" :title="'账本'" @leftEvent="back()" :iconName="iconNameUrl"
              :slortName="'right'"
              @iconEvent="showFilterMth">
      </NavBar>
    </div>
    <div>
      <van-pull-refresh v-model="refreshing" @refresh="onRefresh"
                        :style="{'height':realHeight+'px','overflow-y':'scroll'}">
        <van-list
            v-model:loading="loading"
            :finished="finished"
            finished-text="没有更多了"
            offset="100"
            @load="onLoad"
        >
          <div v-for="item in accountBookList" :key="item.billId">

            <van-swipe-cell>
              <div class="ml5 mt10 mr5 AccountBookcard">
                <van-row>
                  <van-col span="3" style="text-align: center">
                    <van-image class="iconCss" :src="item.imgUrl"/>
                  </van-col>
                  <van-col span="20" offset="1">
                    <van-row type="flex" justify="space-between" align="center">
                      <van-col class="ml5 graytxt tl f16 fc">
                        <span>{{ item.expenseTypeName }}</span>
                      </van-col>
                      <van-col class="tl graytxt f16 bold moneyColor">
                        <span>￥{{ moneyFormat(item.money) }}</span>
                      </van-col>
                    </van-row>
                    <van-row justify="space-between" align="center">
                      <van-col class="ml5 mt5 graytxt tl f12">
                        <span>类型：{{ item.expenseTypeName }}</span>
                        <van-row class="tl mt5 f12">
                          <span>费用日期：{{ item.createTime }}</span>
                        </van-row>
                      </van-col>
                      <van-col v-show="item.state=='0'">
                        <van-button size="small" style="width: 70px;" type="primary" @click="submitAnExpenseAccount(item)">
                          <van-icon name="balance-list-o" size="14"> 报销</van-icon>

                        </van-button>
                      </van-col>
                    </van-row>
                    <van-row class="ml5 mt5 graytxt tl f12 ellipsis">
                      <span>地点：{{
                          item.locationData.address == '' || item.locationData.address == null ? '-' : item.locationData.address
                        }}</span>
                    </van-row>
                    <van-row class="ml5 mt5 tl f12">
                      <van-col :span="5" v-for="(ite,index) in item.voucherList" class="pr5">
                        <van-image :src="imgUrl+ite" @click="previewImage(imgUrl+ite)" alt="加载失败"></van-image>
                      </van-col>
                    </van-row>
                  </van-col>
                </van-row>
              </div>
              <template #right>
                <van-button square text="删除" type="danger" class="delete-button"
                            @click="deleteAccountBook(item)"/>
              </template>
            </van-swipe-cell>


          </div>
        </van-list>
      </van-pull-refresh>
    </div>
    <div>
      <van-action-sheet v-model:show="show" title="筛选">
        <div style="margin-left: 10px;height: 350px;overflow: scroll">
          <div class="btnCss">
            <van-row class="grayTitleCss">报销状态</van-row>
            <van-row>
              <van-col span="8" v-for="item in ReimburseState">
                <div class="selectCard"   @click="changeState(item.state)" :class="queryObj.state==item.state?'check':'noCheck'">
                    {{item.text}}
                </div>
              </van-col>
            </van-row>
          </div>
          <div class="btnCss">
            <van-row class="grayTitleCss">消费类型</van-row>
            <van-row>
              <van-col span="8" v-for="item in consumpotionType">
                <div class="selectCard"   @click="changeType(item.state)" :class="queryObj.expenseTypeName==item.state?'check':'noCheck'">
                  {{item.text}}
                </div>
              </van-col>
            </van-row>
          </div>
        </div>
        <van-divider/>
        <div>
          <van-row style="margin-bottom: 15px">
            <van-col span="4" offset="1">
              <van-button block @click="resetFilter">重置
              </van-button>
            </van-col>
            <van-col span="17" offset="1">
              <van-button  color="#0088ff" block @click="submitFillter">确定
              </van-button>
            </van-col>
          </van-row>
        </div>
      </van-action-sheet>
    </div>
  </div>
</template>

<script setup lang="ts">
import {inject, ref, onMounted} from 'vue'
import {getAccountBookList, deleteBillInfo} from "../../services/accountBook"
import {showConfirmDialog, showSuccessToast, showImagePreview} from "vant";
import cost_icon_01 from '../../assets/icon/cost_icon_01.png'
import cost_icon_02 from '../../assets/icon/cost_icon_02.png'
import cost_icon_03 from '../../assets/icon/cost_icon_03.png'
import cost_icon_04 from '../../assets/icon/cost_icon_04.png'
import cost_icon_05 from '../../assets/icon/cost_icon_05.png'
import cost_icon_06 from '../../assets/icon/cost_icon_06.png'
import cost_icon_07 from '../../assets/icon/cost_icon_07.png'
import cost_icon_08 from '../../assets/icon/cost_icon_08.png'
import cost_icon_09 from '../../assets/icon/cost_icon_09.png'
import cost_icon_10 from '../../assets/icon/cost_icon_10.png'
import cost_icon_11 from '../../assets/icon/cost_icon_11.png'
import cost_icon_12 from '../../assets/icon/cost_icon_12.png'
import cost_icon_13 from '../../assets/icon/cost_icon_13.png'
import cost_icon_14 from '../../assets/icon/cost_icon_14.png'
import cost_icon_15 from '../../assets/icon/cost_icon_15.png'
import screening_01 from '../../assets/icon/screening_01.png'
import screening_02 from '../../assets/icon/screening_02.png'
import {useRouter} from "vue-router";

const router = useRouter();
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
const realHeight = window.document.documentElement.offsetHeight - 50
const show = ref(false)
// 报销状态数组
const ReimburseState=ref([
    {text:'未报销',state:"0"},
    {text:'已报销',state:"1"},
    {text:'已使用',state:"2"},
])
// 消费类型数组
const consumpotionType=([
  {text:'全部',state:""},
  {text:'用车',state:"C"},
  {text:'机票',state:"F"},
  {text:'公交',state:"TR"},
  {text:'火车',state:"T"},
  {text:'娱乐',state:"E"},
  {text:'办公',state:"W"},
  {text:'会议',state:"M"},
  {text:'通讯',state:"R"},
  {text:'酒店',state:"H"},
  {text:'快递',state:"EX"},
  {text:'燃油',state:"FU"},
  {text:'档案',state:"A"},
  {text:'体检',state:"P"},
  {text:'签证',state:"V"},
  {text:'餐饮',state:"RE"},
])
function onLoad() {
  page.value = page.value + 1;
  loading.value = true
  getDataMth()

}

const onRefresh = () => {
  // 清空列表数据
  finished.value = false;
  // refreshing.value = true;
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
    let param = {
      page: page.value,
      pageSize: pageSize.value,
      keyWord: queryObj.value.keyWord,
      empCode: userInfo.userInfo.empCode,
      expenseTypeName: queryObj.value.expenseTypeName,
      occurrenceTime: queryObj.value.occurrenceTime,
      state: queryObj.value.state,
      corpCode: userInfo.userInfo.corpCode,
    }
    let res: any = await getAccountBookList(param)
    if (res.success) {
      res.data.list.map((item: any) => {
        // 尝试解析，进行数据解析
        try {
          var json = JSON.parse((item.locationData));
        } catch (error) {
          // 解析失败，变量不是JSON 抛出去
          return false;
        }
        item.money=parseFloat(item.money).toFixed(2)
        item.locationData = item.locationData ?json : {}
        item.imgUrl = getIconUrl(item.expenseType)
        list.push(item as never)
      })
      loading.value = false;
      total.value = res.data.total

      if (list.length >= total.value) {
        finished.value = true;
      }
    } else {
      list = []
    }
    accountBookList.value = list
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

function getIconUrl(type: any) {
  let imgUrl = ""
  switch (type) {
    case"C":
      imgUrl = cost_icon_01
      break;
    case"F":
      imgUrl = cost_icon_02
      break;
    case"TR":
      imgUrl = cost_icon_03
      break;
    case"T":
      imgUrl = cost_icon_04
      break;
    case"RE":
      imgUrl = cost_icon_05
      break;
    case"E":
      imgUrl = cost_icon_06
      break;
    case"W":
      imgUrl = cost_icon_07
      break;
    case"M":
      imgUrl = cost_icon_08
      break;
    case"R":
      imgUrl = cost_icon_09
      break;
    case"H":
      imgUrl = cost_icon_10
      break;
    case"EX":
      imgUrl = cost_icon_11
      break;
    case"FU":
      imgUrl = cost_icon_12
      break;
    case"A":
      imgUrl = cost_icon_13
      break;
    case"P":
      imgUrl = cost_icon_14
      break;
    case"V":
      imgUrl = cost_icon_14
      break;
  }
  return imgUrl
}

function deleteAccountBook(e: any) {
  showConfirmDialog({
    title: '您确认要删除吗?'
  }).then(async () => {
    let param = {
      billId: e.billId,
      empCode: userInfo.userInfo.empCode,
      corpCode: userInfo.userInfo.corpCode
    }
    let res: any = await deleteBillInfo(param)
    if (res.success) {
      page.value = 0
      finished.value = false;
      refreshing.value = true
      showSuccessToast('删除成功');
      onLoad()
    }
  }).catch(() => {
  });
}

function showFilterMth() {
  show.value = true
}

function back() {
  router.back();
}
// 规范金额显示添加千分位
const moneyFormat=(item:any)=>{
  return item.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",")
}
//
function submitAnExpenseAccount(item: any) {
  // 因为日常报销存在
  const data: any = {
    'bxType': 'rcbx',
    'approveNo': '',
    'isBx': '-'
  };
  sessionStorage.setItem('reimburseDetail', JSON.stringify(data))
  sessionStorage.setItem('billList', JSON.stringify(item))
  sessionStorage.setItem('invoiceList', JSON.stringify([]))
  router.push('addDailyBx');
}

</script>

<style lang="less" scoped>
.selectCard{
  margin: 5px;
  padding: 10px;
  color: #1a1a1a;
  border-radius: 10px;
  font-size: 12px;
  text-align: center;
}
.check{
  background: rgba(0,128,255,1);
  color: #ffffff;
}
.noCheck{
  background: rgba(207,208,217,0.2);
}
.AccountBookcard {
  padding: 8px;
  background: #fff;
  border-radius: 5px;
}

.fc {
  color: #1a1a1a;
}

.fw {
  font-weight: 800;
}

.iconCss {
  width: 45px;
  height:45px;
  /*float: left;*/
  /*margin-left: 20%;*/
}

.ellipsis {
  width: 260px; /*宽度一定要设置*/
  overflow: hidden; //溢出内容隐藏
  text-overflow: ellipsis; //文本溢出部分用省略号表示
  display: -webkit-box; //特别显示模式
  -webkit-line-clamp: 2; //行数
  line-clamp: 2;
  -webkit-box-orient: vertical; //盒子中内容竖直排列
}

.moneyColor {
  color: #0080FF;
}

.delete-button {
  height: 100%;
}

.grayTitleCss {
  font-size: 14px;
  color: #898989;
  margin-top: 10px;
  font-weight: 400;
  margin-left: 5px;
}
.btnCss{
  margin: 10px;
}
.btnCss .van-button {
  margin-left: 5px;
  margin-top: 10px;
  width: 90px;
}

</style>
