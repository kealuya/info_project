<template>
  <van-nav-bar title="新增发票" left-arrow left-text="返回" @click-left="onClickLeft"/>
  <van-form @submit="onSubmit" ref="form">
    <van-field
        v-model="fplxName"
        is-link
        readonly
        name="fplx"
        label="发票类别"
        required
        placeholder="请选择发票类别"
        @click="showPickerFplx = true"
    />
    <template v-for="item in formList">
      <van-field
          v-if="item.inputType === 'text'"
          v-model="data[item.inputKey]"
          :name="item.inputKey"
          :label="item.inputLabel"
          :maxlength="item.maxLangth??80"
          :placeholder="item.placeholder2"
          :autofocus="true"
          :required="item.inputImportant === 1"
          :rules="[{ required: item.inputImportant === 1, message: item.placeholder2 }]"
          autocomplete="off"
      />
      <van-field
          v-if="item.inputType === 'code'"
          v-model="data[item.inputKey]"
          :name="item.inputKey"
          :label="item.inputLabel"
          maxlength="12"
          :placeholder="item.placeholder2"
          :autofocus="true"
          :required="item.inputImportant === 1"
          :rules="[
              { required: item.inputImportant === 1, message: item.placeholder2 },
              {  pattern:  /^[\da-z]+$/i, message: '请输入正确的发票代码' },
              ]"
          autocomplete="off"
      />
      <van-field
          v-if="item.inputType === 'money' || item.inputType === 'number'"
          v-model="data[item.inputKey]"
          :name="item.inputKey"
          :label="item.inputLabel"
          :placeholder="item.placeholder2"
          :autofocus="true"
          :required="item.inputImportant === 1"
          maxlength="10"
          type="number"
          :rules="[{ required: item.inputImportant === 1, message: item.placeholder2 }]"
          autocomplete="off"
      />
      <van-field
          v-if="item.inputType === 'textarea'"
          v-model="data[item.inputKey]"
          :name="item.inputKey"
          :label="item.inputLabel"
          :autofocus="true"
          rows="5"
          :placeholder="item.placeholder2"
          :required="item.inputImportant === 1"
          :rules="[{ required: item.inputImportant === 1, message: item.placeholder2 }]"
          autosize
          autocomplete="off"
      />
      <van-field
          v-if="item.inputType === 'date'&&item.inputKey=='kprq'"
          v-model="data[item.inputKey]"
          is-link
          readonly
          :name="item.inputKey"
          :autofocus="true"
          :label="item.inputLabel"
          :required="item.inputImportant === 1"
          :placeholder="item.placeholder2"
          @click="showCalendar = true"
      />
      <van-field
          v-if="item.inputType === 'date'&&item.inputKey=='cfrq'"
          v-model="data[item.inputKey]"
          is-link
          readonly
          :name="item.inputKey"
          :autofocus="true"
          :label="item.inputLabel"
          :required="item.inputImportant === 1"
          :placeholder="item.placeholder2"
          @click="showCalendarCfrq = true"
      />
      <vue-hash-calendar
          v-if="item.inputKey=='kprq'"
          picker-type="date"
          :scroll-change-date="false"
          :disabled-week-view="true"
          model="dialog" v-model:visible="showCalendar" @confirm="onConfirmDate" :min-date="minDate"
          :default-datetime="data.kprq?new Date(data.kprq):defaultDate"
          :max-date="maxDate"/>
      <vue-hash-calendar
          v-if="item.inputKey=='cfrq'"
          picker-type="date"
          :scroll-change-date="false"
          :disabled-week-view="true"
          model="dialog" v-model:visible="showCalendarCfrq" @confirm="onConfirmDateCFsj" :min-date="minDate"
          :default-datetime="data.cfrq?new Date(data.cfrq):defaultDate"
          :max-date="maxDate"/>
      <!--      <van-calendar :max-date="maxDate" v-model:show="showCalendar" @confirm="onConfirmDate" ref="calendar"-->
      <!--                    :min-date="minDate" :default-date="data[item.inputKey]?new Date(data[item.inputKey]):defaultDate" />-->
      <!--      //  -->
      <van-field
          v-if="item.inputType === 'select'"
          v-model="data[item.inputKey]"
          is-link
          readonly
          :autofocus="true"
          :name="item.inputKey"
          :label="item.inputLabel"
          :placeholder="item.placeholder2"
          @click="showPickerFylb = true"
          :required="true" :rules="[{ required: true, message: '请选择费用类别',trigger: 'blur' }]"
      />
      <van-popup v-if="item.inputType === 'select'" v-model:show="showPickerFylb" position="bottom">
        <van-picker
            :columns="item.selectData"
            @confirm="onConfirmFylb"
            @cancel="showPickerFylb = false"
        />
      </van-popup>
    </template>
    <div v-if="imgOnly==true">
      <van-field name="uploader" label="相关图片" :required="true"
                 :rules="[{ required: true, message: '请上传图片',trigger: 'blur' }]">
        <template #input>
          <van-uploader v-model="imgValue" max-count="1"/>
        </template>
      </van-field>
    </div>
    <div style="margin: 16px;">
      <van-button round block type="primary" native-type="submit">
        提交
      </van-button>
    </div>
  </van-form>
  <van-popup v-model:show="showPickerFplx" round position="bottom">
    <van-picker
        v-model="fplx"
        :columns="typeList"
        title="发票类别"
        :show-toolbar="true"
        @confirm="onConfirmFplx"
        @cancel="showPickerFplx = false"
    />
  </van-popup>
  <div v-show="loading">
    <loadingcomponent :title="'提交中'"/>
  </div>
</template>
<script setup lang="ts">
import moment from "moment";
import {Notify, Uploader} from 'vant';
import {inject, reactive, ref, onBeforeUnmount, watch,} from "vue";
import {invoiceTypeList, typeList} from "./componets/invoiceTypeData";
import Loadingcomponent from '../../components/loading/index.vue';
//@ts-ignore
import DataFormatWithFpMixin from '../../utils/DataFormatWithFpMixin.js';
import {invVerifyWb} from "../../services/invoice";
import {UpLoadImg1} from "../../components/hw_obs_input";
import {useRouter} from "vue-router";

const loading = ref<boolean>(false);
let imgOnly = ref<boolean>(false); //根据上传文件类型决定是否显示上传文件
let calendar = ref()
const router = useRouter();
const onClickLeft = () => {
  sessionStorage.clear();
  invoiceData.$reset();
  history.back()
};
//提交时添加loading效果
let invoiceData: any = inject("invoiceData");
const showCalendar = ref<boolean>(false);
const showCalendarCfrq = ref<boolean>(false);
const showPickerFplx = ref<boolean>(false);
const showPickerFylb = ref<boolean>(false);
let userInfoData: any = inject("userInfo");
let formList = reactive<any>([]);
const imgValue = ref<any>([]);
let maxDate = new Date();
const form = ref()
// 默认赋值
let data: any = reactive({
  fplx: '5',
  kprq: '',
  cfrq: '',
  ph: '',
  cfdd: '',
  dddd: '',
  cc: '',
  pb: '',
  fylb: ''
})
//最小日期范围
let minDate = reactive({});
//默认日期
let defaultDate = reactive({})
let date = new Date();
let year = date.getFullYear();
let month = date.getMonth();
let day = date.getDate();
// 设置最小时间为2000年1月1日
minDate = new Date(2000, 0, 1);
// 最大时间限制为当天
defaultDate = new Date(year, month, day)
invoiceTypeList[0].list.map((item: any) => {
  data[item.inputKey] = "";
});
if ((invoiceData.invoiceData.data || {}).fplx) {
  data = invoiceData.invoiceData.data;
  if (data.fplx == '1') {
    data.cfdd = invoiceData.invoiceData.data.fz ?? '';
    data.dddd = invoiceData.invoiceData.data.dz ?? '';
    data.pb = invoiceData.invoiceData.data.xwlx ?? '';
    
    if (invoiceData.invoiceData.data.rq != '') {
      let kprq = moment(invoiceData.invoiceData.data.rq).format('YYYY/MM/DD');
      data.cfrq = kprq;
      data.kprq = kprq;
    }
    data.fpje = invoiceData.invoiceData.data.je ?? '';
  } else if (data.fplx == '0' || data.fplx == '6' || data.fplx == '7' || data.fplx == '9' || data.fplx == '8' || data.fplx == '10' || data.fplx == '16' || data.fplx == '22' || data.fplx == '23' || data.fplx == '27' || data.fplx == '28') {
    if (data.fpjym != null) {
      data.jym = data.fpjym.substring(data.fpjym.length - 6, data.fpjym.length);
    }
    if (data.kprq != '') {
      let newKprq = moment(data.kprq).format('YYYY/MM/DD');
      data.kprq = newKprq;
    }
    if (data.lb) {
      let dataList = data.lb;
      let dataIndex = ''
      if (dataList.length == 1) {
        data.fwmc = dataList[0].lbmc;
      } else {
        for (let i = 0; i < dataList.length; i++) {
          dataIndex = dataIndex + dataList[i].lbmc + ','
        }
        data.fwmc = dataIndex;
      }
    }
    
  } else if (data.fplx == '3') {
    data.xm = invoiceData.invoiceData.data.cjrxm ?? '';
    data.hbh = invoiceData.invoiceData.data.list[0].cc ?? '';
    data.cfdd = invoiceData.invoiceData.data.list[0].fz ?? '';
    data.dddd = invoiceData.invoiceData.data.list[0].dz ?? '';
    data.pb = invoiceData.invoiceData.data.xwlx ?? '';
    data.fpje = invoiceData.invoiceData.data.je ?? '';
  } else if (data.fplx == '12') {
    data.fpjym = data.fpjym.substring(data.fpjym.length - 6, data.fpjym.length);
    data.fpjetax = data.fpje;
    data.fpje = data.sqje;
    if (data.kprq != '') {
      let kprq = moment(data.kprq).format('YYYY/MM/DD');
      data.kprq = kprq;
    }
    let dataList = data.lb;
    let dataIndex = ''
    if (dataList.length == 1) {
      data.fwmc = dataList[0].lbmc;
    } else {
      for (let i = 0; i < dataList.length; i++) {
        dataIndex = dataIndex + dataList[i].lbmc + ','
      }
      data.fwmc = dataIndex;
    }
  } else {
    if (data.kprq != '') {
      let nowDate = moment(data.kprq).format('YYYY/MM/DD');
      data.kprq = nowDate;
    }
  }
} else if ((invoiceData.invoiceData.data || {}).resArr) {
  data = invoiceData.invoiceData.data.resArr[0];
  let kprq = moment(invoiceData.invoiceData.data.resArr[0].kprq).format('YYYY/MM/DD');
  if (data.lb) {
    let dataList = data.lb;
    let dataIndex = ''
    if (dataList.length == 1) {
      data.fwmc = dataList[0].lbmc;
    } else {
      for (let i = 0; i < dataList.length; i++) {
        dataIndex = dataIndex + dataList[i].lbmc + ','
      }
      data.fwmc = dataIndex;
    }
  }
  
  data.kprq = kprq;
  data.fpjym = data.fpjym.substring(data.fpjym.length - 6, data.fpjym.length);
  
}

let fplx = ref<any>(['5']);
let fplxName = ref<string>('定额发票');
invoiceTypeList.map((item) => {
  if ((data || {}).fplx) {
    if (item.ocrType == Number(data.fplx)) {
      fplx.value = [data.fplx];
      fplxName.value = item.label;
      if (invoiceData.invoiceData.imgUrl) {
        imgValue.value = [{url: invoiceData.invoiceData.imgUrl, deletable: false,}];
      }
      console.log(invoiceData.invoiceData.imgUrl)
      let fileType;
      if (invoiceData.invoiceData.imgUrl) {
        fileType = invoiceData.invoiceData.imgUrl.slice(invoiceData.invoiceData.imgUrl.length - 3, invoiceData.invoiceData.imgUrl.length);
      }
      if (fileType == 'pdf') {
        imgOnly.value = false;
      } else {
        imgOnly.value = true;
      }
      formList.push(...item.list);
    }
    
  } else {
    let res = data.resArr[0];
    if (item.ocrType == Number(res.fplx)) {
      fplx.value = [res.fplx];
      fplxName.value = item.label;
      if (invoiceData.invoiceData.imgUrl) {
        imgValue.value = [{url: invoiceData.invoiceData.imgUrl, deletable: false,}];
      }
      formList.push(...item.list);
    }
  }
});
// 中国标准时间 转换成 年月日

const onConfirmDate = (date: any) => {
  data.kprq = moment(Number(date)).format('YYYY/MM/DD')
  
  showCalendar.value = false
};
const onConfirmDateCFsj = (date: any) => {
  data.cfrq = moment(Number(date)).format('YYYY/MM/DD')
  showCalendar.value = false
};
const onConfirmFplx = (value: any, index: any) => {
  formList = reactive([]);
  invoiceTypeList.map((item) => {
    if (item.ocrType == Number(value.value)) {
      data.fplx = (item.ocrType).toString();
      fplxName.value = item.label;
      if (invoiceData.invoiceData.imgUrl) {
        imgValue.value = [{url: invoiceData.invoiceData.imgUrl}];
      }
      formList.push(...item.list);
    }
  });
  showPickerFplx.value = false;
};
const onConfirmFylb = (value: any, index: any) => {
  data.fylb = value.text
  data.pb = value.text
  if (data.fplx == '28') {
    data.bz = value.text
  }
  showPickerFylb.value = false;
};
const onSubmit = async (values: any) => {
  loading.value = true;
  data.base64 = data.cutfile;
  let submitJson1 = DataFormatWithFpMixin.ocrDataToShowData(data, "hand");
  let submitJson2 = DataFormatWithFpMixin.showDataToVerifyData(submitJson1[0]);
  submitJson2.qyId = userInfoData.userInfo.corpCode;
  submitJson2.empCode = userInfoData.userInfo.empCode;
  submitJson2.fileName = DataFormatWithFpMixin.getUuid() + ".jpeg";
  if (data.kprq) {
    submitJson2.kprq = moment(data.kprq).format('YYYYMMDD');
  }
  if (data.cfrq) {
    submitJson2.cfrq = moment(data.cfrq).format('YYYYMMDD');
  }
  submitJson2.xm = data.xm;
  submitJson2.pb = data.pb;
  if (imgValue.value[0].file) {
    let result: any = await UpLoadImg1([imgValue.value[0].file]);
    submitJson2.filePath = result[0];
  } else {
    submitJson2.filePath = imgValue.value[0].url;
  }
  let res: any = await invVerifyWb(submitJson2);
  if (res.code == '1001') {
    Notify({type: 'success', message: '提交成功'});
    loading.value = false;
    invoiceData.$reset();
    router.replace({path: '/'});
  } else {
    Notify({type: 'danger', message: res.msg ?? '提交失败'});
    loading.value = false;
  }
};
onBeforeUnmount(() => {
  invoiceTypeList[0].list.map((item: any) => {
    data[item.inputKey] = "";
    invoiceData.invoiceData.imgUrl = ''
  });
  data.fplx = '5'
})
watch(() => fplxName.value, () => {
  invoiceTypeList[0].list.map((item: any) => {
    data[item.inputKey] = "";
    invoiceData.invoiceData.imgUrl = ''
  });
  data.kprq = '';
  data.fpjym = '';
  data.cfrq = '';
  data.cfdd = '';
  data.dddd = '';
  data.cc = '';
  data.ph = '';
  data.xm = '';
  data.pb = '';
  data.cph = '';
  data.zjh = '';
  data.xsfmc = '';
  data.xsfnsrsbh = '';
  data.gmfmc = '';
  data.gmfnsrsbh = '';
  data.fwmc = '';
  imgValue.value = []
  form.value.resetValidation()
})
</script>

<style lang="less" scoped>

/deep/ .hash-calendar {
  background: rgba(0, 0, 0, 0.1)
}

/deep/ .calendar_title {
  bottom: 260.406px !important;
}

/deep/ .calendar_content {
  height: 260.406px !important;
}
</style>
