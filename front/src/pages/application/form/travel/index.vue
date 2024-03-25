<template>
  <div>
    <nav-bar-component title="差旅申请"/>
    <van-form @submit="onSubmit">
      <div class="formCard">
        <div class="pl5 f16"><span class="require">*</span>申请单号</div>
        <van-field
            v-model="approveNo"
            name="approveNo"
            placeholder="请输入申请单号"
            readonly
            required
            :rules="[{ required: true, message: '请输入申请单号' }]"
        />
      </div>
      <div class="formCard">
        <div class="pl5 f16"><span class="require">*</span>行程事由</div>
        <van-field
            v-model="approveReason"
            name="approveReason"
            placeholder="请输入行程事由"
            required
            :rules="[{ required: true, message: '请输入行程事由' }]"
        />
      </div>
      <div class="formCard">
        <div class="pl5 f16">
          <span v-if='costCenterRequired=="0"' class="require">*</span>
          成本中心
        </div>
        <van-field
            v-model="costCenterInfo"
            name="costCenterInfo"
            placeholder="请选择成本中心"
            @click="selectCostCenterOpen"
            readonly
            is-link
            required
            :rules="[{ required: true, message: '请选择成本中心' }]"
        />
      </div>
      <div class="formCard" v-if=" approfeMethods != '' && approfeMethods == '0'">
        <van-row justify="space-between" align="center">
          <van-col>
            <div class="pl5 f16"><span class="require">*</span>部门领导审批流程</div>
          </van-col>
          <van-col class="mr10" v-if="flowName!=''"  @click="seeApproveInfoPicker=true">
            <span class="f12">查看审批流程信息</span>
            <van-icon name="info-o" class="ml5" size="15" />
          </van-col>
        </van-row>
        <van-field
            v-model="flowName"
            name="flowName"
            placeholder="请选择部门领导审批流程"
            readonly
            @click="selectApprovePicker=true"
            is-link
            required
            :rules="[{ required: true, message: '请选择部门领导审批流程' }]"
        />
        <van-popup v-model:show="selectApprovePicker" position="bottom">
          <van-picker
              :columns="approveList"
              @confirm="onApproveConfirm"
              :columns-field-names="customApproveFieldName"
              @cancel="selectApprovePicker = false"
          />
        </van-popup>
      </div>
      <div class="formCard" v-else-if="approfeMethods != '0'&&flowList.length!=0">
        <van-row justify="space-between" align="center">
          <van-col>
            <div class="pl5 f16"><span class="require">*</span>成本中心审批流程</div>
          </van-col>
          <van-col class="mr10" v-if="flowName!=''"  @click="seeApproveInfoPicker=true">
            <span class="f12">查看审批流程信息</span>
            <van-icon name="info-o" class="ml5" size="15" />
          </van-col>
        </van-row>
        <van-field
            v-model="flowName"
            name="flowName"
            placeholder="请选择成本中心审批流程"
            @click="selectCostCenterFlowPicker=true"
            readonly
            is-link
            required
            :rules="[{ required: true, message: '请选择成本中心审批流程' }]"
        />
        <van-popup v-model:show="selectCostCenterFlowPicker" position="bottom">
          <van-picker
              :columns="flowList"
              @confirm="onConfirm"
              :columns-field-names="customFieldName"
              @cancel="selectCostCenterFlowPicker = false"
          />
        </van-popup>
      </div>
      <div class="formCard">
        <div class="pl5 f16"><span class="require">*</span>出差人员（最多可选5人）</div>
        <van-field
            v-model="companionTitle"
            name="companionTitle"
            placeholder="请选择陪同人员"
            readonly
            @click="selectPeoplePicker=true"
            is-link
            required
            :rules="[{ required: true, message: '请输入陪同人员' }]"
        />
      </div>
      <div class="formCard">
        <div class="pl5 f16"><span class="require">*</span>行程时间</div>
        <van-field
            v-model="leaveDate"
            name="leaveDate"
            placeholder="请选择行程时间"
            is-link
            required
            @click="showDatePicker = true"
            readonly
            :rules="[{ required: true, message: '请选择开始结束时间' }]"
        />
        <van-popup v-model:show="showDatePicker" position="bottom">
          <van-calendar v-model:show="showDatePicker" type="range" @confirm="onConfirmDate"/>
        </van-popup>
      </div>
      <!--   出行种类选择 -->
      <div class="formCard">
        <div class="pl5 f16"><span class="require">*</span>出行种类（可多选）</div>
        <div>
          <van-row>
            <van-col span="6">
              <div :class="travelType.indexOf('train')==-1?'uncheckTip':'checkTip'" @click="toggleType('train')">火车
              </div>
            </van-col>
            <van-col span="6">
              <div :class="travelType.indexOf('airplane')==-1?'uncheckTip':'checkTip'" @click="toggleType('airplane')">
                飞机
              </div>
            </van-col>
            <van-col span="6">
              <div :class="travelType.indexOf('hotel')==-1?'uncheckTip':'checkTip'" @click="toggleType('hotel')">酒店
              </div>
            </van-col>
          </van-row>
        </div>
        <div v-if="travelType.indexOf('train')!=-1">
          <div class="travelTypeCard">
            <van-row justify="space-between" align="center">
              <van-col class="font_gray f12 ml5"> 火车信息</van-col>
              <van-col>
                <van-image :src="ShenQingAdd" @click="addTrain" width="70px"/>
              </van-col>
            </van-row>
          </div>
          <div class="m10 pb10" v-for="(item,index) in travelTrainList">
            <div v-if="item.travelBy=='train'">
              <van-row justify="space-between" align="center">
                <van-col>
                  <div class="f14 pl5 font_gray">行程{{ index + 1 }}</div>
                </van-col>
                <van-col v-if="index!=0">
                  <div class="f14 font_blue" @click="delTravel(index)">删除</div>
                </van-col>
              </van-row>
              <div class="mb5 mt5 cityPicker">
                <van-row class="mt5 ml5 mb5" justify="center" align="center">
                  <van-col class="f12 tc" span="10" @click="startLocation(item,index)">
                    <span>{{ item.fromCity }}</span>
                  </van-col>
                  <van-col>
                    <van-image
                        :class="{ 'arrowTransform': !item.flag, 'arrowTransformReturn': item.flag}"
                        :src="ShenQingToggle"
                        @click="toggleLocation(item)"
                        width="30px" height="30px"
                    />
                  </van-col>
                  <van-col class="f12 tc" span="10" @click="endLocation(item,index)">
                    <span>{{ item.toCity }}</span>
                  </van-col>
                </van-row>
              </div>
            </div>
          </div>
        </div>
        <div v-if="travelType.indexOf('airplane')!=-1">
          <div class="travelTypeCard">
            <van-row justify="space-between" align="center">
              <van-col class="font_gray f12 ml5"> 飞机信息</van-col>
              <van-col>
                <van-image :src="ShenQingAdd" @click="addPlane" width="70px"/>
              </van-col>
            </van-row>
          </div>
          <div class="m10 pb10" v-for="(item,index) in travelPlaneList">
            <div v-if="item.travelBy=='air'">
              <van-row justify="space-between" align="center">
                <van-col>
                  <div class="f14 pl5 font_gray">行程{{ index + 1 }}</div>
                </van-col>
                <van-col v-if="index!=0">
                  <div class="f14 font_blue" @click="delPlaneTravel(index)">删除</div>
                </van-col>
              </van-row>
              <div class="mb5 mt5 cityPicker">
                <van-row class="mt5 ml5 mb5" justify="center" align="center">
                  <van-col class="f12 tc" span="10" @click="startLocation(item,index)">
                    <span>{{ item.fromCity }}</span>
                  </van-col>
                  <van-col>
                    <van-image
                        :class="{ 'arrowTransform': !flag, 'arrowTransformReturn': flag}"
                        :src="ShenQingToggle"
                        @click="toggleLocation(item)"
                        width="30px" height="30px"
                    />
                  </van-col>
                  <van-col class="f12 tc" span="10" @click="endLocation(item,index)">
                    <span>{{ item.toCity }}</span>
                  </van-col>
                </van-row>
              </div>
            </div>
          </div>
        </div>
        <div v-if="travelType.indexOf('hotel')!=-1">
          <div class="travelTypeCard">
            <van-row justify="space-between" align="center">
              <van-col class="font_gray f12 ml5"> 酒店信息</van-col>
              <van-col>
                <van-image :src="ShenQingAdd" @click="addHotel" width="70px"/>
              </van-col>
            </van-row>
          </div>
          <div class="m10 pb10" v-for="(item,index) in travelHotelList">
            <div v-if="item.sleepBy=='hotel'">
              <van-row justify="space-between" align="center">
                <van-col>
                  <div class="f14 pl5 font_gray">行程{{ index + 1 }}</div>
                </van-col>
                <van-col v-if="index!=0">
                  <div class="f14 font_blue" @click="delHotelTravel(index)">删除</div>
                </van-col>
              </van-row>
              <div class="mb5 mt5 cityPicker">
                <van-row class="mt5 ml5 mb5" justify="space-between" align="center" @click="endLocation(item,index)">
                  <van-col class="f12 tc" span="10">
                    <span>{{ item.toCity }}</span>
                  </van-col>
                  <van-col>
                    <van-icon color="#999999" name="arrow"/>
                  </van-col>
                </van-row>
              </div>
            </div>
          </div>
        </div>

      </div>
      <div class="formCard">
        <div class="pl5 f16"><span class="require">*</span>行程预算</div>
        <van-row justify="start" align="center" class="ml10">
          <van-col>¥</van-col>
          <van-col>
            <van-field
                v-model="money"
                name="money"
                type="number"
                placeholder="请输入行程预算"
                :rules="[{ required: true, message: '请输入行程预算' },{ pattern, message: '金额格式输入错误请重新输入' }]"
            />
          </van-col>
        </van-row>

      </div>
      <div class="formCard">
        <div class="pl5 f16"><span class="require">*</span>经办人</div>
        <van-field
            v-model="creater"
            name="creater"
            placeholder="请输入经办人"
            readdonly
            required
            :rules="[{ required: true, message: '请输入申请单号' }]"
        />
      </div>
      <div class="formCard">
        <div class="pl5 f16">行程备注</div>
        <van-field
            v-model="bz1"
            name="approveReason"
            placeholder="请输入行程备注"
            rows="3"
            type="textarea"
            autosize
            :rules="[{ required: true, message: '请输入行程备注' }]"
        />
      </div>
      <div class="btnStyle">
        <van-button block type="primary" native-type="submit">
          提交
        </van-button>
      </div>
    </van-form>
  </div>
  <!--  选择出差人员-->
  <van-popup v-model:show="selectPeoplePicker" style="width:100vw;height:100vh" position="bottom">
    <select-people-picker @select-over="selectPeople" ref="selectPeopleOver" :list="accompanyList" max-number="5" />
  </van-popup>
  <!-- 选择成本中心 -->
  <van-popup v-model:show="selectCostCenterPicker" style="width:100vw; height:100vh" position="bottom">
    <select-cost-center-picker @select-over="selectCostCenter" ref="selectCostCenterOver" :data="costCenterData"/>
  </van-popup>
  <!-- 选择城市 -->
  <van-popup v-model:show="selectCityPicker" style="width:100vw;height:100vh" position="bottom">
    <select-city-picker @select-over="selectCity" ref="selectCostCenterOver" :type="cityType"/>
  </van-popup>
  <!-- 审批显示 -->
  <van-popup v-model:show="seeApproveInfoPicker" style="width:100vw;height:100vh" position="bottom">
    <approval-info  @select-over="seeApproveInfoPicker=false" :approve-info-list="getApproveInfoList" :code="getApproveNum"  />
  </van-popup>
</template>

<script lang="ts" setup>
import {onMounted, ref, inject} from "vue";
import {orderNumberM} from "../ts/orderNumberM";
import SelectCostCenterPicker from '../../../../components/selectComponent/selectCostCenterPicker/index.vue'; // 引入选择成本中心的组件
import SelectPeoplePicker from '../../../../components/selectComponent/selectPeoplePicker/index.vue' // 引入选择出差人员的组件
import SelectCityPicker from '../../../../components/selectComponent/selectCityPicker/index.vue' // 引入选择城市的组件
import NavBarComponent from "../../../../components/navBar/index.vue"; //  引入通用导航栏
import ApprovalInfo from "../components/approveInfo/index.vue";
import {
  getApprovalFlow,
  getApprovalFlowHistory,
  getApprovalFlowList,
} from '../../../../services/current'; // 导入需要用到的接口
import ShenQingAdd from '../../../../assets/icon/shenqing_icon_02.png'; // 新增按钮
import ShenQingToggle from '../../../../assets/icon/shenqing_icon_01.png' // 旋转箭头
import {showSuccessToast} from "vant";
import {showFailToast} from "vant/lib/toast/function-call"; // 弹窗
import moment from "moment/moment"; // 日期格式插件引入
import {createTravelApply} from "../../../../services/apply"; // 提交差旅申请
import router from "../../../../router"; // 导入格式化日期文件
const flag = ref(false); // 旋转控制的变量
const approveNo = ref(''); // 申请单号
const approveReason = ref(''); // 行程事由
const showDatePicker = ref(false); // 控制开始结束时间
const companionTitle = ref(''); //出差人员提交展示
const companion = ref([]); // 存放出差人员的数组
let userInfoData: any = inject("userInfo"); //获取审批数组
const accompanyList = ref([]); // 陪同人员
const creater = ref(''); // 经办人
const leaveDate = ref(''); //时间展示
const startDate = ref(''); //出发时间
const endDate = ref(''); //结束时间
const bz1=ref(''); // 备注1
const selectPeoplePicker = ref(false); // 控制选择出差人员
const selectCostCenterFlowPicker = ref(false); // 控制选择成本中心审批的组件显隐
const selectCostCenterPicker = ref(false); // 控制选择成本中心
const selectCityPicker = ref(false); // 控制选择城市（用于出发地点和到达地点）
const seeApproveInfoPicker=ref(false) // 查看审批详情
const selectApprovePicker = ref(false); //控制选择部门审批流审批
const getApproveInfoList=ref([]); // 用于获取接口中的审批数组
const getapproveCCInfo=ref([]);  // 用于获取抄送的审批数组
const getApproveNum=ref(); // 获取审批数组
const travelType = ref(['train']); // 判断出行种类的类别
const travelList = ref([]); // 用于存储填写提交前的出行方式种类的数组
const travelTrainList = ref([]); // 用于存储出行方式种类选择为火车数组
const travelPlaneList = ref([]); // 用于存储出行方式种类选择为飞机的数组
const travelHotelList = ref([]); // 用于存储出行方式种类选择为酒店的数组
const costCenterData = ref({
  'xmmc': '',
  'xmye': '',
  'xmbh': '',
  'budGetingId': '',
}); // 将获取回来的成本中心数据存放在该对象中
const costCenterInfo = ref('') //用于展示在表单页面中的选中后的表单数据
const flowId = ref('') // 审批流程
const costCenterRequired = ref(''); // 控制成本中心是否必填项
const approfeMethods = ref(''); //审批方式
const flowList = ref([]); // 成本中心审批流数组
const flowName = ref(''); // 审批流数组
const cityData = ref({}); // 选中得到的城市
const startOrEnd = ref('0') // 出发还是结束 出发：0/ 结束：1
const cityType = ref('') //将key传递至组件以改变  飞机 air 火车 train 酒店 hotel
const cityIndex = ref(0) // 用于记录出行方式数组下标
const money = ref('')// 预算金额
const approveList = ref([]); // 用于部门审批流审批
const pattern = /^[0-9]+(\.[0-9]{1,2}|)$/; // 保留两位小数的金额校验
// 根据审批流回显审批流人员信息
const getFlowHistory = async (value: any) => {
  let payload = {
    "school": userInfoData.userInfo.corpCode,
    'flowId': value,
    "corpCode": userInfoData.userInfo.corpCode,
    'empCode': userInfoData.userInfo.empCode,
    'equipmentType': ""
  }
  let result: any = await getApprovalFlowHistory(payload);
  let flowHistory: any = JSON.parse(result.data);
  getApproveInfoList.value = flowHistory.data.toUserList;
  getapproveCCInfo.value = flowHistory.data.ccUserList;
  getApproveNum.value = flowHistory.data.isVote;
  const object: any = {};
  // ccUserList.value=flowHistory.data.ccUserList;
  for (let i = 0; i <= getApproveInfoList.value.length - 1; i++) {
    let no = getApproveInfoList.value[i]['no'];
    let isAllPass = getApproveInfoList.value[i]['isAllPass'];
    if (!object[no]) {
      object[no] = {
        no,
        isAllPass,
        children: []
      }
    }
    if (no === object[no].no) {
      object[no].children.push(getApproveInfoList.value[i])
    }
  }
  getApproveInfoList.value = [];
  Object.values(object).forEach((item) => {
    return getApproveInfoList.value.push(item as never);
  })
  let userName: any = [];
  flowHistory.data.ccUserList.forEach((item: any) => userName.push(`${item.ccUserName}(${item.toRole})`));
}
// 打开成本中心组件
const selectCostCenterOpen = () => {
  costCenterData.value = costCenterData.value;
  selectCostCenterPicker.value = true;
}
// 选择成本中心
const selectCostCenter = (data: any) => {
  if (data.xmmc) {
    costCenterData.value = data;
    costCenterInfo.value = `${data.budGetingYear}-${data.xmmc}   ¥ ${data.xmye}`;
    GetCostCenterFlow(data.xmbh);
  }
  selectCostCenterPicker.value = false;
}
// 根据获取到成本中心选择可选择的审批流
const GetCostCenterFlow = async (key: any) => {
  let payLoad = {
    "applyTypeId": "",
    "xmbh": key,
    "deptId": userInfoData.userInfo.deptId,
    "corpCode": userInfoData.userInfo.corpCode,
    "approvalFlowType": "costCenterFlow",
    "equipmentType": ""
  }
  let result: any = await getApprovalFlow(payLoad);
  flowName.value = '';
  flowId.value = '';
  if (result.data[0].length != 0) {
    flowList.value = result.data;
  } else {
    if(approfeMethods.value != '0'){
      showFailToast('该成本中心下没有相关审批流，请联系管理员');
    }
  }
}
// 用于规范选择框类型(成本中心审批流)
const customFieldName = {
  text: 'xmmc',
  value: 'flowid',
};
// 用于规范选择框类型(部门领导审批流)
const customApproveFieldName = {
  text: 'flow_name',
  value: 'flow_id',
};
// 选择部门领导审批流
const onApproveConfirm = (value: any) => {
  flowName.value = value.selectedOptions[0].flow_name;
  flowId.value = value.selectedOptions[0].flow_id;
  getFlowHistory(flowId.value);
  selectApprovePicker.value = false;
};
// 选择成本中心审批流
const onConfirm = (value: any) => {
  flowName.value = value.selectedOptions[0].xmmc;
  flowId.value = value.selectedOptions[0].flowid;
  getFlowHistory(flowId.value);
  selectCostCenterFlowPicker.value = false;
}
// 选择陪同人员
const selectPeople = (list: any) => {
  let toSelectCCUserList: any = [];
  if (list.length != 0) {
    companionTitle.value = '';
    companion.value = [];
    list.forEach((item: any) => {
      companionTitle.value = companionTitle.value + item.empName + '，';
      return toSelectCCUserList.push(item);
    });
    companionTitle.value = companionTitle.value.substring(0, companionTitle.value.length - 1);
    companion.value = toSelectCCUserList;
  }
  selectPeoplePicker.value = false;
}
// 切换(反转)出发地点和到达地点
const toggleLocation = (item: any) => {
  item.flag = !item.flag;
  let fromCity = item.fromCity;
  let fromCityCode = item.fromCityCode;
  item.fromCity = item.toCity;
  item.fromCityCode = item.toCityCode;
  item.toCity = fromCity;
  item.toCityCode = fromCityCode;
}
//出行种类选择
const toggleType = (code: any) => {
  if (travelType.value.indexOf(code) == -1) {
    travelType.value.push(code);
    if (code == 'train') {
      addTrain()
    } else if (code == 'airplane') {
      addPlane()
    } else {
      addHotel()
    }
  } else {
    let key = travelType.value.indexOf(code);
    travelType.value.splice(key, 1);
    if (code == 'train') {
      travelTrainList.value = [];
    } else if (code == 'airplane') {
      travelPlaneList.value = [];
    } else {
      travelHotelList.value = []
    }
  }
}
// 删除出行信息(火车)
const delTravel = (index: any) => {
  travelTrainList.value.splice(index, 1)
}
// 删除出行信息(飞机)
const delPlaneTravel = (index: any) => {
  travelPlaneList.value.splice(index, 1)
}
// 删除出行信息(酒店)
const delHotelTravel = (index: any) => {
  travelHotelList.value.splice(index, 1)
}
// 组件返回来之后的值
const selectCity = (item: any) => {
  if (item.label) {
    let travelVNList: any[]; // 新建拷贝数组用于赋值处理 (中间起名取自VNode)
    if (cityType.value == 'train') {
      travelVNList = travelTrainList.value; // 拷贝一份数组用于赋值处理(火车)
    } else if (cityType.value == 'air') {
      travelVNList = travelPlaneList.value; // 拷贝一份数组用于赋值处理(飞机)
    } else {
      travelVNList = travelHotelList.value; // 拷贝一份数组用于赋值处理(酒店)
    }
    let handleVnData = travelVNList[cityIndex.value]; //准确拿到数组的值用于更新
    if (startOrEnd.value == '0') { // 判断是出发还是到达 来进行赋值操作
      handleVnData.fromCity = item.label;
      handleVnData.fromCityCode = item.code;
    } else {
      handleVnData.toCity = item.label;
      handleVnData.toCityCode = item.code;
    }
  }
  selectCityPicker.value = false;
}
// 添加出行方式-火车出行
const addTrain = () => {
  const data = {
    'index': (travelList.value.length + 1).toString(),
    'travelBy': 'train',
    'fromCity': '出发地点',
    'toCity': '到达地点',
    'firstKey': 'fromCity',
    'flag':false,
    'secondKey': 'toCity',
    'fromCityCode': '',
    'toCityCode': '',
    'turns': 0.00,
  }
  travelTrainList.value.push(data as never);
}
// 添加出行方式-飞机出行
const addPlane = () => {
  const data = {
    'index': (travelList.value.length + 1).toString(),
    'travelBy': 'air',
    'fromCity': '出发地点',
    'toCity': '到达地点',
    'firstKey': 'fromCity',
    'flag':false,
    'secondKey': 'toCity',
    'fromCityCode': '',
    'toCityCode': '',
    'turns': 0.00,
  }
  travelPlaneList.value.push(data as never);
}
// 添加出行方式-酒店出行
const addHotel = () => {
  const data = {
    'index': (travelList.value.length + 1).toString(),
    'sleepBy': 'hotel',
    'toCity': '出差城市',
    'toCityCode': '',
  }
  travelHotelList.value.push(data as never);
}
// 选择出发地点
const startLocation = (item: any, index: any) => {
  startOrEnd.value = '0'; // 用于记录是出发地点还是到达地点， 0为出发地点 1为到达地点
  cityType.value = ''; // 将所需要重新渲染的城市列表组件按照type重新渲染
  cityIndex.value = 0; // 将记录传递下标重制
  cityType.value = item.travelBy; // 将数组类型赋给type值
  cityIndex.value = index; // 记录修改参数下标
  selectCityPicker.value = true;
}
// 选择到达地点
const endLocation = (item: any, index: any) => {
  startOrEnd.value = '1';
  cityType.value = ''; // 将所需要重新渲染的城市列表组件按照type重新渲染
  cityIndex.value = 0; // 将记录传递下标重制
  cityType.value = item.travelBy ?? item.sleepBy; // 将数组类型赋给type值
  cityIndex.value = index; // 记录修改参数下标
  selectCityPicker.value = true;
}
// 在处于部门审批流程的时候
const deptApproveFlow = async () => {
  let payload = {
    "corpCode": userInfoData.userInfo.corpCode,
    'deptId': userInfoData.userInfo.deptId,
    "typeId": "FK_CL_SQ",
    'isOpen': '1',
    'empCode': userInfoData.userInfo.empCode,
    'equipmentType': "pc"
  }
  let result: any = await getApprovalFlowList(payload);
  approveList.value = result.data;
}
// 选择时间
const onConfirmDate = (values: any) => {
  const [start, end] = values;
  showDatePicker.value = false;
  startDate.value = moment(start).format("yyyy-MM-DD");
  endDate.value = moment(end).format("yyyy-MM-DD");
  leaveDate.value = `${startDate.value} - ${endDate.value}`;
};
// 根据获取的值判断是部门审批还是成本中心审批
const approveMotheds = async () => {
  // 判断
  let rootList = userInfoData.functionConfig;
  // 成本中心审批
  for (var item in rootList) {
    if (rootList[item].ywTypeID == 'fk007') {
      approfeMethods.value = rootList[item].configValue;
    }
    if (rootList[item].ywTypeID == 'fk009') {
      costCenterRequired.value = rootList[item].configValue;
    }
    if (costCenterRequired.value != '0') {
      deptApproveFlow()
    }
  }
};
// 提交差旅申请（此处携带校验）
const onSubmit = async (values: any) => {
  // 校验出行方式数组是否为空
  if (travelType.value.length == 0) {
    showSuccessToast('请选择出行方式，以及提交相关出行信息');
    return false
  }
  // 校验火车中是否完全填写
  let trainKey: any = travelTrainList.value.find((e: { fromCity: any, toCity: any }) => e.fromCity == '出发地点' || e.toCity == '到达地点');
  // 校验飞机中是否完全填写
  let airPlaneKey: any = travelPlaneList.value.find((e: { fromCity: any, toCity: any }) => e.fromCity == '出发地点' || e.toCity == '到达地点');
  // 校验酒店中是否完全填写
  let hotelKey: any = travelPlaneList.value.find((e: { fromCity: any, toCity: any }) => e.toCity == '到达地点');
  if (trainKey) {
    showSuccessToast('请将火车出行方式具体填写');
    return false
  }
  if (airPlaneKey) {
    showSuccessToast('请将飞机出行方式具体填写');
    return false
  }
  if (hotelKey) {
    showSuccessToast('请将酒店中的出差城市具体填写');
    return false
  }
  if (money.value=='0') {
    showSuccessToast('行程预算不能为0');
    return false
  }
  postRequire()
};

//拼接提交参数
const postRequire = async () => {
  // 拼接提交要用到的数组
  let goCityPeople = companion.value.map((e: { deptName: any, empName: any, empCode: any }) => `${e.deptName}-${e.empName}(${e.empCode})`);
  let goCityPeopleList: { corpCode: any; deptId: any; deptName: any; empCode: any; empName: any; idCard: any; jobName: any; phone: any; postBandId: string; togetherPersonId: any; togetherPersonName: any; }[] = [];
  companion.value.map((e: { corpCode: any, deptName: any, empName: any, empCode: any, deptId: any, idCard: any, jobName: any, phone: any }) => goCityPeopleList.push({
    corpCode: e.corpCode,
    deptId: e.deptId,
    deptName: e.deptName,
    empCode: e.empCode,
    empName: e.empName,
    idCard: e.idCard,
    jobName: e.jobName,
    phone: e.phone,
    postBandId: "",
    togetherPersonId: e.empCode,
    togetherPersonName: e.empName
  }));
  let travelApplyDetail: any = {
    "corpCode": userInfoData.userInfo.corpCode,
    "routeType": "corporate", //提交为成本中心审批  corporate
    "approveNo": approveNo.value,
    "travelType": "corporate",
    "fromCity": "",
    "toCity": "",
    "startTime": startDate.value,
    "returnTime": endDate.value,
    "travelBudget": money.value,
    "creater": userInfoData.userInfo.userid,
    "createTime": "",
    "bz1": bz1.value,
    "flowId": flowId.value
  };
  let postTravelData: any = {
    "approveNo": approveNo.value,
    "approveReason": approveReason.value,
    "state": "",
    "creater": userInfoData.userInfo.userName,
    "deptId": userInfoData.userInfo.deptId,
    "deptName": userInfoData.userInfo.deptName,
    "createTime": "",
    "bz1": bz1.value,
    "appKey": "8rjvhuir778ghuisblq8uwu8bjusbvh",
    "appSceret": "fcb7b69e75b183e0a59726a7a7e64edb",
    "type": "travel",
    "travelApplyDetail": travelApplyDetail,
    "travelHotelDetail": travelHotelList.value,
    "travelTrainFightDetail": travelTrainList.value.concat(travelPlaneList.value),
    "togetherPerson": goCityPeopleList,
    "togetherStaffs": goCityPeople,
    "approvalFlowId": flowId.value,
    "flowId": "",
    "xmmc": costCenterData.value.xmmc,
    "xmye": costCenterData.value.xmye,
    "xmbh": costCenterData.value.xmbh,
    "costType": "cost",
    "budGetingId": costCenterData.value.budGetingId,
    "approvalMode": "0",  // 判断是否有审批流0，自定义审批流1
    "agentPeopleDetail": [],
    "ccUserList": [],
    "corpCode": userInfoData.userInfo.corpCode,
    "empCode": userInfoData.userInfo.empCode,
    "equipmentType": ''
  }
  let result: any = await createTravelApply(postTravelData);
  if (result.success == true) {
    showSuccessToast(result.msg);
    router.back();
  } else {
    showFailToast(result.msg);
  }

}
onMounted(() => {
  approveNo.value = orderNumberM();
  creater.value = `${userInfoData.userInfo.deptName} - ${userInfoData.userInfo.userName}(${userInfoData.userInfo.empCode})`
  approveMotheds();
  addTrain();
});
</script>

<style scoped>
.van-cell__title van-field__label van-field__label--required {
  display: none;
}

.formCard {
  margin-top: 10px;
  padding-top: 10px;
  background: white;
}

.btnStyle {
  margin: 16px;
}

.require {
  color: red;
  font-size: 16px;
}

.checkTip {
  margin: 5px 10px;
  padding: 10px;
  text-align: center;
  border-radius: 8px;
  color: #ffffff;
  font-size: 12px;
  background-color: #0088ff;
}

.uncheckTip {
  margin: 5px 10px;
  padding: 10px;
  text-align: center;
  border-radius: 8px;
  color: #666666;
  font-size: 12px;
  background-color: #EDF4FA;
}

.number {
  display: inline-block;
  width: 20px;
  height: 20px;
  border: 2px solid #333333;
  border-radius: 15px;
  margin-left: 10px;
  margin-right: 10px;
  text-align: center;
}

.travelTypeCard {
  margin: 10px;
  border-radius: 5px;
  background: #EFF5FB;
}

.cityPicker {
  border-bottom: 1px rgba(153, 153, 153, 0.5) solid;
}

/*图片旋转*/
.arrowTransform {
  transition: 0.2s;
  transform-origin: center;
  transform: rotateZ(180deg);
}

.arrowTransformReturn {
  transition: 0.2s;
  transform-origin: center;
  transform: rotateZ(360deg);
}
</style>
