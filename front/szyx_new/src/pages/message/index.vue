<template>
  <div>
    <!-- 上半部分的header     -->
    <div style="background:#0080FF;height: 30%;padding-bottom:20px">
      <van-row class="pt10 pl20">
        <van-col span="23" style="text-align: center">
          <span class="f_white">消息通知</span>
        </van-col>
        <van-col span="1">
          <van-icon style="float: right;margin-right: 20px;"
                    v-if="activeName==0"
                    :name="filterIdentification == 0?screening_01:screening_02"
                    size="23"
                    @click="filterBtnClick"/>
        </van-col>
      </van-row>
    </div>

    <div class="messageCss">
      <van-tabs v-model:active="activeName">
        <van-tab title="推送消息" name="0">
          <div style="background: #FFFFFF;min-height: 40px">
            <van-row justify="space-between" align="center">
              <van-col>
                <div @click="clearReadMth" class="clearIcon">
                  <van-row type="flex" justify="space-between" color="#0088ff" align="center">
                    <van-col>
                      <van-icon :name="qkIcon" class="pl5" size="12"/>
                    </van-col>
                    <van-col>
                      <div class="font_blue ml5 mr5 f12">清空未读</div>
                    </van-col>
                  </van-row>
                </div>
              </van-col>
              <van-col class="headTagCss">
                <van-row justify="space-around" align="center" type="flex">
                  <van-col v-for="item in approveTypeList">
                    <div class="typeCard"
                         :class="approvalType==item.value?'checkCard'||filterObj.formState==item.value:''"
                         @click="changeFiledMth(item.value)">{{ item.text }}
                    </div>
                  </van-col>
                </van-row>

              </van-col>
            </van-row>
          </div>
          <div :style="{height:clientHeight+'px',overflow:'scroll'}">
            <div class="content_center mt20" v-if="cardPullList.length<1">
              <van-image width="220" height="220" :src="default_img_03"/>
            </div>
            <van-pull-refresh v-model="refreshing" @refresh="onPullMessageRefresh"
                              :style="{minHeight:clientHeight+'px'}">
              <van-list
                  v-model:loading="pullMessageLoading"
                  :finished="pullMessageFinished"
                  finished-text="没有更多了"
                  offset="100"
                  @load="onLoad"
                  :immediate-check="false"
              >
                <van-cell v-if="cardPullList.length>0" v-for="item in cardPullList" :key="item"
                          @click="queryApprovalDetail(item)"
                          class="mt10">
                  <div class="cardCss">
                    <van-row>
                      <van-col span="10" class="f16 cardTitle">
                        <div class="pb10"
                             style="display: flex;align-items: center;justify-content: left;">
                          <van-badge :dot=item.dotReadState position="top-right">
                            <van-icon class="mr10 iconSizeCss" size="20"
                                      :name="message_icon_01"/>
                          </van-badge>

                          <span> {{ item.apply_type }}</span>

                        </div>
                      </van-col>
                      <van-col span="3"></van-col>
                      <van-col span="11" class="f14 fc">{{ item.create_time }}</van-col>
                    </van-row>
                    <van-row class="f14 ">
                      <van-col class="fc fw ta-l" span="5">
                        <span>推送通知</span>
                      </van-col>
                      <van-col span="7"></van-col>
                      <van-col span="12" class="f16"
                               :style="{color: item.state=='已拒绝'?'red':item.state=='待审批'?'#0080FF':item.state=='已同意'?'green':item.state=='已撤销'?'red':'#898989'}"
                               v-if="item.state!=undefined&&item.state!=''">
                        <div>{{ item.state }}</div>
                      </van-col>
                    </van-row>
                    <van-row class="f14 fc">
                      <van-col>申请事由:</van-col>
                      <van-col>
                        <div class="ellipsis">{{ item.type_matter }}</div>
                      </van-col>
                    </van-row>
                    <van-row class="f14 fc">
                      <span>申请人: {{ item.apply_user }}</span>
                    </van-row>
                    <van-row class="f14 fc">
                      <span>审核人: {{ item.appr_user }}</span>
                    </van-row>
                    <van-row class="f14 fc">
                      <span>申请单号: {{ item.ywno }}</span>
                    </van-row>
                    <van-row class="f14 fc fw" v-if="item.state == '待审批'">
                      <span style="color: #FF9200">您有一个新的审批通知，请进行审批操作！</span>
                    </van-row>
                    <van-divider/>
                    <van-row class="f14 fc">
                      <van-col span="22">
                        <div style="float:left;margin-bottom: 5px">查看详情</div>
                      </van-col>
                      <van-col span="2">
                        <van-icon name="arrow"/>
                      </van-col>
                    </van-row>
                  </div>
                </van-cell>
              </van-list>
            </van-pull-refresh>
          </div>

        </van-tab>
        <van-tab title="系统消息" name="1">
          <div :style="{minHeight:clientHeight+'px'}">
            <div class="content_center mt20" v-if="cardNewList.length<1">
              <van-image width="220" height="220" :src="default_img_03"/>
            </div>
            <van-pull-refresh v-model="delegatedSettingsRefresh" @refresh="onSysMessageRefresh"
                              style="min-height: 800px">
              <van-list
                  v-model:loading="messageLoading"
                  :finished="messageFinished"
                  finished-text="没有更多了"
                  offset="100"
                  @load="onLoadMessage">
                <van-cell v-if="cardNewList.length>0" v-for="item in cardNewList" :key="item"
                          @click="queryDetail(item)" class="mt10 pl10">
                  <div class="cardCss">
                    <van-row>
                      <van-col span="8" class="f16 cardTitle">
                        <div style="float: left">
                          <van-icon class="mr10 iconSizeCss"
                                    :name="message_icon_01"/>
                          <span> {{ item.newType }}</span>
                        </div>
                      </van-col>
                      <van-col span="6"></van-col>
                      <van-col span="10" class="f14 fc">{{ item.createTime }}</van-col>
                    </van-row>
                    <van-row class="f14 pt10 fc">
                      <span>{{ item.newsTitle }}</span>
                    </van-row>
                    <van-divider/>
                    <van-row class="f14 fc">
                      <van-col span="22">
                        <div style="float:left;margin-bottom: 5px">查看详情</div>
                      </van-col>
                      <van-col span="2">
                        <van-icon name="arrow"/>
                      </van-col>
                    </van-row>
                  </div>
                </van-cell>
              </van-list>
            </van-pull-refresh>
          </div>
        </van-tab>
      </van-tabs>
    </div>
    <van-action-sheet v-model:show="show" title="筛选">
      <div style="margin-left: 10px;height: 400px;overflow: scroll">
        <van-row class="titleCss">关键字检索</van-row>
        <van-row class="f10">
          <van-field v-model="filterObj.queryKeywords" placeholder="请输入申请事由等关键字"/>
          <van-divider/>
        </van-row>
        <van-row class="titleCss">申请单类型</van-row>
        <van-row>
          <div class="btnCss">
            <van-row class="grayTitleCss">全部</van-row>
            <van-row>
              <van-button size="small" @click="changeType('')"
                          :type="filterObj.requisitionFormType==''?'primary':'default'">全部
              </van-button>
            </van-row>
          </div>
        </van-row>
        <div>
          <van-row class="grayTitleCss">oa办公</van-row>
          <div class="btnCss">
            <van-row>
              <van-col>
                <van-button size="small" @click="changeType('请假申请')"
                            :type="filterObj.requisitionFormType=='请假申请'?'primary':'default'">请假申请
                </van-button>
              </van-col>
              <van-col>
                <van-button size="small" @click="changeType('加班申请')"
                            :type="filterObj.requisitionFormType=='加班申请'?'primary':'default'">加班申请
                </van-button>
              </van-col>
              <van-col>
                <van-button size="small" @click="changeType('用章申请')"
                            :type="filterObj.requisitionFormType=='用章申请'?'primary':'default'">用章申请
                </van-button>
              </van-col>
            </van-row>
            <van-row>
              <van-col>
                <van-button size="small" @click="changeType('考勤申请')"
                            :type="filterObj.requisitionFormType=='考勤申请'?'primary':'default'">考勤申请
                </van-button>
              </van-col>
              <van-col>
                <van-button size="small" @click="changeType('物品领用')"
                            :type="filterObj.requisitionFormType=='物品领用'?'primary':'default'">物品领用
                </van-button>
              </van-col>
              <van-col>
                <van-button size="small" @click="changeType('董事文件签批')"
                            :type="filterObj.requisitionFormType=='董事文件签批'?'primary':'default'">董事签批
                </van-button>
              </van-col>
            </van-row>
            <van-row>
              <van-col>
                <van-button size="small" @click="changeType('通用审批')"
                            :type="filterObj.requisitionFormType=='通用审批'?'primary':'default'">通用审批
                </van-button>
              </van-col>
              <van-col>
                <van-button size="small" @click="changeType('记录便签')"
                            :type="filterObj.requisitionFormType=='记录便签'?'primary':'default'">记录便签
                </van-button>
              </van-col>
            </van-row>
          </div>
        </div>
        <div class="btnCss">
          <van-row class="grayTitleCss">业务场景</van-row>
          <van-row>
            <van-col>
              <van-button size="small" @click="changeType('差旅申请')"
                          :type="filterObj.requisitionFormType=='差旅申请'?'primary':'default'">差旅申请
              </van-button>
            </van-col>
            <van-col>
              <van-button size="small" @click="changeType('用车申请')"
                          :type="filterObj.requisitionFormType=='用车申请'?'primary':'default'">用车申请
              </van-button>
            </van-col>
            <van-col>
              <van-button size="small" @click="changeType('客户宴请')"
                          :type="filterObj.requisitionFormType=='客户宴请'?'primary':'default'">客户宴请
              </van-button>
            </van-col>
          </van-row>
          <van-row>
            <van-col>
              <van-button size="small" @click="changeType('外出办公')"
                          :type="filterObj.requisitionFormType=='外出办公'?'primary':'default'">外出办公
              </van-button>
            </van-col>
            <van-col>
              <van-button size="small" @click="changeType('会议记录')"
                          :type="filterObj.requisitionFormType=='会议记录'?'primary':'default'">会议记录
              </van-button>
            </van-col>
            <van-col>
              <van-button size="small" @click="changeType('快递寄送')"
                          :type="filterObj.requisitionFormType=='快递寄送'?'primary':'default'">快递寄送
              </van-button>
            </van-col>
          </van-row>
          <van-row>
            <van-col>
              <van-button size="small" @click="changeType('物品申购')"
                          :type="filterObj.requisitionFormType=='物品申购'?'primary':'default'">物品申购
              </van-button>
            </van-col>
            <van-col>
              <van-button size="small" @click="changeType('订餐申请')"
                          :type="filterObj.requisitionFormType=='订餐申请'?'primary':'default'">订餐申请
              </van-button>
            </van-col>
            <van-col>
              <van-button size="small" @click="changeType('合同申请')"
                          :type="filterObj.requisitionFormType=='合同申请'?'primary':'default'">合同申请
              </van-button>
            </van-col>
          </van-row>
          <van-row>
            <van-col>
              <van-button size="small" @click="changeType('基金报销')"
                          :type="filterObj.requisitionFormType=='基金报销'?'primary':'default'">基金报销
              </van-button>
            </van-col>
            <van-col>
              <van-button size="small" @click="changeType('培训申请')"
                          :type="filterObj.requisitionFormType=='培训申请'?'primary':'default'">培训申请
              </van-button>
            </van-col>
          </van-row>
        </div>
        <van-row class="grayTitleCss">报销相关</van-row>
        <van-row class="btnCss">
          <van-col>
            <van-button size="small" @click="changeType('日常报销')"
                        :type="filterObj.requisitionFormType=='日常报销'?'primary':'default'">日常报销
            </van-button>
          </van-col>
          <van-col>
            <van-button size="small" @click="changeType('差旅补助')"
                        :type="filterObj.requisitionFormType=='差旅补助'?'primary':'default'">差旅补助
            </van-button>
          </van-col>
          <van-col>
            <van-button size="small" @click="changeType('借款申请')"
                        :type="filterObj.requisitionFormType=='借款申请'?'primary':'default'">借款申请
            </van-button>
          </van-col>
        </van-row>
        <div class="btnCss">
          <van-row class="titleCss">申请状态</van-row>
          <van-row>
            <van-col>
              <van-button size="small" @click="changeState('')"
                          :type="filterObj.formState==''?'primary':'default'">全部
              </van-button>
            </van-col>
            <van-col>
              <van-button size="small" @click="changeState('已同意')"
                          :type="filterObj.formState=='已同意'?'primary':'default'">已同意
              </van-button>
            </van-col>
            <van-col>
              <van-button size="small" @click="changeState('已拒绝')"
                          :type="filterObj.formState=='已拒绝'?'primary':'default'">已拒绝
              </van-button>
            </van-col>
          </van-row>
          <van-row>
            <van-col>
              <van-button size="small" @click="changeState('已撤销')"
                          :type="filterObj.formState=='已撤销'?'primary':'default'">已撤销
              </van-button>
            </van-col>
            <van-col>
              <van-button size="small" @click="changeState('抄送通知')"
                          :type="filterObj.formState=='抄送通知'?'primary':'default'">抄送通知
              </van-button>
            </van-col>
            <van-col>
              <van-button size="small" @click="changeState('抄送结果通知')"
                          :type="filterObj.formState=='抄送结果通知'?'primary':'default'">抄送结果通知
              </van-button>
            </van-col>
          </van-row>
        </div>
      </div>
      <van-divider/>
      <div>
        <van-row style="margin-bottom: 15px">
          <van-col span="4">
            <van-button style="margin-left: 20px;width: 100%" size="small" @click="resetFilter">重置
            </van-button>
          </van-col>
          <van-col span="6"/>
          <van-col span="13">
            <van-button style="width: 100%" type="primary" size="small" @click="submitBtn">确定</van-button>
          </van-col>
          <van-col span="1"/>
        </van-row>
      </div>
    </van-action-sheet>
  </div>
</template>

<script setup name="message" lang="ts">
import {ref, toRef, reactive, onMounted, inject} from 'vue'
import {getMessageList, getMessagePull, updateMsgPullState} from "../../services/message";
import router from "../../router";
import {userInfoData} from '../../store'
import {showConfirmDialog} from 'vant';
import {applyType} from '../home/component/ts/applyType';
import screening_01 from '../../assets/icon/screening_01.png'
import screening_02 from '../../assets/icon/screening_02.png'
import qkIcon from '../../assets/icon/清空.png'
import default_img_03 from '../../assets/icon/default_img_03.png'
import message_icon_01 from '../../assets/icon/message_icon_01.png'

let userInfo: any = inject("userInfo");
const store = userInfoData();
const activeName = ref(0);
// 申请状态list
const approveTypeList = ref([
  {text: '全部', value: ''},
  {text: '已同意', value: '已同意'},
  {text: '已拒绝', value: '已拒绝'},
  {text: '抄送', value: '抄送通知'},
]);
// 单据类型list
const documentTypeList=ref([
  {text: '全部', value: '',type:'all'},
  {text: '请假申请', value: '请假申请',type:'office'},
  {text: '加班申请', value: '加班申请',type:'office'},
  {text: '用章申请', value: '用章申请',type:'office'},
  {text: '考勤申请', value: '考勤申请',type:'office'},
  {text: '物品领用', value: '物品领用',type:'office'},
  {text: '董事签批', value: '董事签批',type:'office'},
  {text: '通用审批', value: '通用审批',type:'office'},
  {text: '记录便签', value: '记录便签',type:'office'},
  {text: '差旅申请', value: '差旅申请',type:'business'},
  {text: '用车申请', value: '用车申请',type:'business'},
  {text: '客户宴请', value: '客户宴请',type:'business'},
  {text: '外出办公', value: '外出办公',type:'business'},
  {text: '会议记录', value: '会议记录',type:'business'},
  {text: '快递寄送', value: '快递寄送',type:'business'},
  {text: '物品申购', value: '物品申购',type:'business'},
  {text: '订餐申请', value: '订餐申请',type:'business'},
  {text: '合同申请', value: '合同申请',type:'business'},
  {text: '基金报销', value: '基金报销',type:'business'},
  {text: '培训申请', value: '培训申请',type:'business'},
  {text: '日常报销', value: '订餐申请',type:'expense'},
  {text: '差旅补助', value: '差旅补助',type:'expense'},
  {text: '借款申请', value: '借款申请',type:'expense'},
])
//系統消息list
let cardNewList: any = ref([])
//推送消息list
let cardPullList:any = ref([]);
//弹出筛选框状态
const show = ref(false)
//推送消息
const pullMessageLoading = ref(false)
const pullMessageFinished = ref(false);
const pullMesasageTotle = ref(0)
const pullMessagePage = ref(0)
const pullMessagePageSize = ref(10)
//系统消息
const messageLoading = ref(false)
const messageFinished = ref(false);
const mesasageTotle = ref(0)
const messagePage = ref(0)
const messagePageSize = ref(10)
//推送消息审批状态
const approvalType = ref('')
//筛选弹窗参数
const filterObj = ref({
  queryKeywords: '',
  requisitionFormType: '',
  formState: ''
});
// 次数
const filterIdentification = ref(0)
const state = reactive({
  count: 1,
  color: "black"
})
const colorRef = toRef(state, "color")
//获取高度
const clientHeight = ref(document.documentElement.offsetHeight - 185)
const refreshing:any = ref(false);
const delegatedSettingsRefresh = ref(false)

onMounted(() => {
  let pullMessage: any = store.pullMessage
  if (pullMessage != "" && pullMessage) {
    cardPullList.value = pullMessage.list as never
    pullMesasageTotle.value = pullMessage.total as never
    pullMessagePage.value = 1
  }
})

const onPullMessageRefresh = () => {
  pullMessageFinished.value = false;
  // 重新加载数据
  // 将 loading 设置为 true，表示处于加载状态
  pullMessageLoading.value = true;
  pullMessagePage.value = 0
  onLoad();

};
const onSysMessageRefresh = () => {
  messageFinished.value = false;
  // 重新加载数据
  // 将 loading 设置为 true，表示处于加载状态
  messageLoading.value = true;
  messagePage.value = 0
  onLoadMessage();

};

function onLoadMessage() {
  messageLoading.value = true
  messagePage.value = messagePage.value + 1
  getMessageListMth()
}

//系统消息查看详情
function queryDetail(e: any) {
  router.push({
    name: 'sysMessageDetail',
    params: {
      htmlMessage: e.newContent,
      newTitle: e.newsTitle,
      newType: e.newsTitle,
      createTime: e.createTime,
      dept: e.receiveDeptName
    }
  })
}

function onLoad() {
  pullMessagePage.value = pullMessagePage.value + 1
  pullMessageLoading.value = true
  getMessagePullListMth(filterIdentification.value)
};

//TODO 跳转审批详情(跳转不同详情页)
function queryApprovalDetail(e: any) {
  let approveType = ""
  applyType.map((item) => {
    if (e.apply_type == item.text) {
      approveType = item.value
    }
  })
  const data: any = {
    'approveType': approveType,
    'approveNo': e.ywno,
  };
  sessionStorage.setItem('applicationDetail', JSON.stringify(data))
  if (e.state == "已同意" || e.state == "已拒绝") {
    switch (e.apply_type) {
      case "日常报销":
        const data: any = {
          'bxType': 'rcbx',
          'approveNo': e.ywno,
          'isBx': e.isBx
        };
        sessionStorage.setItem('reimburseDetail', JSON.stringify(data))
        router.push('/reimburse/detail');
        break;
      case "差旅补助报销":
        router.push({
          name: 'travelSubsidiesDetail',
          params: {
            applyNo: e.ywno
          }
        })
        break;
      case "借款申请":
        router.push({
          name: 'loanDetail',
          params: {
            loanId: e.ywno
          }
        })
        break;
      default:
        router.push('/applicationDetail')
        break;
    }

  } else if (e.state == "抄送结果通知" || e.state == "抄送通知") {
    let url = `http://122.9.41.215/fk_move/MapprList?userid=${userInfo.userInfo.empCode}&username=${userInfo.userInfo.userName}&school=${userInfo.userInfo.corpCode}&type=fk&isDark=false#/mobile/fk/MapprFkDetail?state=wait&apprno=${e.appr_no}&page=Copy&apprResult&type=fk&isDark=false`
    router.push({
      name: 'approvalDetailsIframe',
      params: {
        url: url,
        realHeight: document.documentElement.offsetHeight,
        titleName: '抄送通知'
      }
    })

  } else if (e.state == "已撤销") {
    let url = `http://122.9.41.215/fk_move/MapprList?userid=${userInfo.userInfo.empCode}&username=${userInfo.userInfo.userName}&school=${userInfo.userInfo.corpCode}&type=fk&isDark=false#/mobileCommon/apprMineCommon`
    router.push({
      name: 'approvalDetailsIframe',
      params: {
        url: url,
        realHeight: document.documentElement.offsetHeight,
        titleName: '详情'
      }
    })

  } else if (e.state == "待审批") {
    let url = `http://122.9.41.215/fk_move/MapprList?userid=${userInfo.userInfo.empCode}&username=${userInfo.userInfo.userName}&school=${userInfo.userInfo.corpCode}&type=fk&isDark=false#/mobile/fk/MapprFkDetail?state=wait&apprno=${e.appr_no}&id=1&page=waitMine&type=fk&isDark=false`
    router.push({
      name: 'approvalDetailsIframe',
      params: {
        url: url,
        realHeight: document.documentElement.offsetHeight,
        titleName: '待审批'
      }
    })
  }
  updateReadState(e)
}


async function updateReadState(item: any) {

  let param = {
    corpCode: userInfo.userInfo.corpCode,
    empCode: userInfo.userInfo.empCode,
    id: item.id,
    sign: ""
  }
  let res: any = await updateMsgPullState(param)
  if (res.success) {
    store.$patch({
      unReadMessageNum: res.data.count
    })
  }
}

//点击筛选状态按钮方法
function changeFiledMth(e: any) {
  pullMessageFinished.value = false
  filterIdentification.value = 0
  if (e == '') {
    approvalType.value = '';
    filterObj.formState = '';
    filterIdentification.value = 0;
    pullMessagePage.value = 0;
    cardPullList.value = [];
    onLoad()
  } else {
    approvalType.value = e;
    filterObj.value.formState=e;
    filterIdentification.value = 1;
    pullMessagePage.value = 0;
    cardPullList.value = [];
    onLoad()
  }
}

//筛选弹出
function filterBtnClick() {
  show.value = true
}

//筛选关闭
function close() {
  show.value = false
}

//重置
function resetFilter() {
  filterObj.value.queryKeywords = '',
      filterObj.value.requisitionFormType = '',
      filterObj.value.formState = ''
}

function changeType(e: any) {
  filterObj.value.requisitionFormType = e
}

function changeState(e: any) {
  filterObj.value.formState = e
}

//获取系统消息
async function getMessageListMth() {
  let sysMsgList = cardNewList.value
  setTimeout(async () => {
    if (delegatedSettingsRefresh.value) {
      sysMsgList = []
      delegatedSettingsRefresh.value = false
    }
    let param = {
      page: messagePage.value,
      pageSize: messagePageSize.value,
      corpCode: userInfo.userInfo.corpCode,
      deptId: userInfo.userInfo.deptId,
      empCode: userInfo.userInfo.empCode,
    }
    let res: any = await getMessageList(param);
    if (res.success) {
      res.data.list.map((item: any) => {
        sysMsgList.push(item as never)
      })
      messageLoading.value = false;
      mesasageTotle.value = res.data.total
      if (sysMsgList.length >= mesasageTotle.value) {
        messageFinished.value = true;
      }
    } else {
      sysMsgList = []
    }
    cardNewList.value = sysMsgList
  }, 1000)

}

//获取推送消息
async function getMessagePullListMth(type: number) {
  let msgList = cardPullList.value
  setTimeout(async () => {
    if (refreshing.value) {
      msgList = []
      refreshing.value = false
    }
    let param = {}
    if (type == 1) {
      param = {
        page: pullMessagePage.value,
        pageSize: pullMessagePageSize.value,
        corpCode: userInfo.userInfo.corpCode,
        deptId: userInfo.userInfo.deptId,
        empCode: userInfo.userInfo.empCode,
        approvalType: filterObj.value.formState,
        ywType: filterObj.value.requisitionFormType,
        reason: filterObj.value.queryKeywords == '' ? '' : filterObj.value.queryKeywords
      }
    } else {
      param = {
        page: pullMessagePage.value,
        pageSize: pullMessagePageSize.value,
        corpCode: userInfo.userInfo.corpCode,
        deptId: userInfo.userInfo.deptId,
        empCode: userInfo.userInfo.empCode,
        approvalType: approvalType.value
      }
    }
    let res: any = await getMessagePull(param);
    if (res.success) {
      res.data.list.map((item: any) => {
        item.dotReadState = item.isRead == '0' ? true : false
        msgList.push(item as never)
      })
      pullMessageLoading.value = false;
      pullMesasageTotle.value = res.data.total
      if (msgList.length >= pullMesasageTotle.value) {
        pullMessageFinished.value = true;
      }
      store.$patch({
        unReadMessageNum: res.data.count
      })
    } else {
      msgList = []
    }
    cardPullList.value = msgList
  }, 500)

}

//筛选方法
function submitBtn() {
  close()
  pullMessageFinished.value = false
  if (filterObj.value.formState == '' && filterObj.value.requisitionFormType == '' && filterObj.value.queryKeywords == '') {
    filterIdentification.value = 0
  } else {
    filterIdentification.value = 1
  }
  approvalType.value = filterObj.value.formState
  pullMessagePage.value = 0
  cardPullList.value = []
  onLoad()
}

function clearReadMth() {
  showConfirmDialog({
    title: '您确认要清空未读消息吗?'
  }).then(() => {
    updateReadStatus().then(() => {
      filterIdentification.value = 0
      approvalType.value = ''
      pullMessagePage.value = 0
      cardPullList.value = []
      onLoad()
    })
  }).catch(() => {
    // on cancel
  });
}

async function updateReadStatus() {
  let param = {
    corpCode: userInfo.userInfo.corpCode,
    empCode: userInfo.userInfo.empCode,
    id: 0,
    sign: "1"
  }
  let res: any = await updateMsgPullState(param)
  if (res.success) {
    store.$patch({
      unReadMessageNum: res.data.count
    })
  }
}


</script>

<style lang="less" scoped>

.messageCss {
  margin-top: -1px;
}

/deep/ .messageCss .van-tabs .van-tabs__nav {
  background: #0080FF !important;
}

/deep/ .messageCss .van-tabs .van-tabs__nav {
  background: #0080FF !important;
  color: #FFFFFF;
}

/deep/ .messageCss .van-tab {
  color: #FFFFFF !important;

}

/deep/ .headTagCss .van-tag {
  /*color: #323233 !important;*/
  padding: 3px 5px;
  margin: 5px;
  height: 20px;
}

.cardCss {
  /*margin-left: 10px;*/
  margin-top: 15px;
  /*padding: 10px;*/
}

.btnCss .van-button {
  margin-left: 5px;
  margin-top: 10px;
  width: 90px;
}

.titleCss {
  font-size: 14px;
  margin-top: 6px;
  font-weight: 600;
  margin-left: 5px;
}

.cardTitle {
  color: #0080FF;
}

.grayTitleCss {
  font-size: 14px;
  color: #898989;
  margin-top: 10px;
  font-weight: 400;
  margin-left: 5px;
}

/deep/ .nullImg .van-icon .van-icon__image {
  width: 260px;
  height: 260px;
  align: center;
  padding: 40% 0 0 25%
}

.fc {
  color: #1a1a1a;
}

.fw {
  font-weight: 800;
}

/deep/ .fc_white .van-tag {
  /*color: #323233 !important;*/
  color: #FFFFFF !important;
}

/deep/ .fc_black .van-tag {
  color: #1a1a1a !important;
}

.iconSizeCss {
  width: 14px;
  height: 14px;
  padding-bottom: 2px;

}

.clearIcon {
  margin: 6px 5px;
  padding: 5px;
  background: rgba(0, 128, 255, 0.1);
  border-radius: 8px;
}

.ellipsis {
  width: 200px; /*宽度一定要设置*/
  overflow: hidden; /*文本超出隐藏*/
  text-overflow: ellipsis; /*文本超出显示省略号*/
  white-space: nowrap; /*超出的空白区域不换行*/
  text-align: left
}

.ta-l {
  text-align: left;
}

// 状态颜色
.typeCard {
  margin: 5px;
  padding: 5px;
  font-size: 12px;
}

.checkCard {
  background: rgba(0, 128, 255, 0.1);
  border-radius: 4px;
  color: #0088ff;
}
</style>
