<template>
  <div>
    <navbar title="个人信息" />
      <div class="userInfoCard">
        <van-skeleton title :row="12" :loading="loading">
        <van-form>
          <van-cell-group inset>
            <van-field
                v-model="userInfo.chineseName"
                name="姓名"
                label="姓名"
                required
                placeholder="姓名"
                :rules="[{ required: true, message: '请填写姓名' }]"
            />
            <van-field
                v-model="userInfo.idCard"
                name="身份证号"
                label="身份证号"
                required
                placeholder="身份证号"
                @click-input="desenInputText"
                :rules="[{ required: true, message: '请填写身份证号' },{validator, message: '身份证号码格式错误！'}]"
            />
            <van-field
                v-model="userInfo.phone"
                name="手机号"
                label="手机号"
                required
                center
                clearable
                :formatter="phoneFormatter"
                format-trigger="onBlur"
                placeholder="手机号"
                disabled
                :rules="[{ required: true, message: '请填写手机号' }]"
            >
              <template #button>
                <van-button size="small" type="primary" @click="goChangePhone">变更手机</van-button>
              </template>
            </van-field>
            <van-field
                v-model="userInfo.deptName"
                name="所属部门"
                label="所属部门"
                placeholder="所属部门"
                disabled
            />
            <van-field
                v-model="userInfo.jobName"
                name="职位名称"
                label="职位名称"
                placeholder="职位名称"
                disabled
            />
            <van-field
                v-model="userInfo.empCode"
                name="工号"
                label="工号"
                placeholder="工号"
                disabled
            />
            <van-field
                v-model="userInfo.email"
                name="邮箱"
                label="邮箱"
                placeholder="邮箱"
            />
            <van-field
                v-model="userInfo.bz"
                name="个人备注"
                rows="1"
                autosize
                type="textarea"
                label="个人备注"
                placeholder="个人备注"
            />
          </van-cell-group>
          <div style="margin: 16px;">
            <van-button  block type="primary" size="normal" native-type="submit" @click="updateUser">
              提交
            </van-button>
          </div>
        </van-form>
        </van-skeleton>
      </div>
    <div v-show="loading">
      <loading-component :title="'提交中'"/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Navbar from './../../components/navBar/index.vue'
import {inject,ref,onMounted} from "vue";
import {getUserInfo, updateUserInfo} from "../../services/user";
import LoadingComponent from '../../components/loading/index.vue'
import {showSuccessToast} from "vant";
import router from "../../router";
import {showFailToast} from "vant/lib/toast/function-call";
let userInfoData: any = inject("userInfo"); //用户信息
let loading=ref(false); // 控制loading显隐
let  idCard=ref(''); // 存储的身份证
let phone =ref(''); // 存储的手机号
const phoneFormatter = (value:any) => value.replace(/^(.{3})(?:\d+)(.{4})$/, "$1****$2");
let userInfo=ref({
  'idCard':'',
  'bz':'',
  'jobName':'',
  'phone':'',
  'jobTitleId':'',
  'corpName':'',
  'email':'',
  'chineseName':'',
  'empCode':'',
  'englishName':'',
  'deptName':''
}); //从接口中获得的用户数据
// 获取用户信息
const getUser=async()=>{
  let userLoad={
    'corpCode':userInfoData.userInfo.corpCode,
    'empCode':userInfoData.userInfo.empCode,
    'equipmentType':''
  }
  let result: any = await getUserInfo(userLoad);
  if(result.success==true){
    loading.value=false;
    userInfo.value=result.data;
    idCard.value=result.data.idCard;
    phone.value=result.data.phone;
    userInfo.value.idCard=userInfo.value.idCard.replace(/^(.{8})(?:\d+)(.{4})$/, "$1******$2");
  }else{
    loading.value=false
    showFailToast('获取个人信息失败，请检查网络或联系管理员');
  }

}
// 身份证号脱敏
const desenInputText=()=>{
    userInfo.value.idCard=idCard.value;
}
// 修改用户信息
const updateUser=async()=>{
  loading.value=true;
  userInfo.value.phone=phone.value;
  let postdata={
    "uuid": userInfoData.userInfo.empCode,
    "corpName": userInfo.value.corpName,
    "deptId": "242",
    "chineseName":  userInfo.value.chineseName,
    "userName": userInfo.value.empCode,
    "englishName": userInfo.value.englishName,
    "deptName": userInfo.value.deptName,
    "phone": userInfo.value.phone,
    "jobName": userInfo.value.jobName,
    "email": userInfo.value.email,
    "idCard": userInfo.value.idCard,
    "postbandId": userInfo.value.jobTitleId,
    "bz": userInfo.value.bz,
    "empCode": userInfoData.userInfo.empCode,
    "corpCode": userInfoData.userInfo.corpCode,
    "equipmentType": ""
  }
  let result: any = await updateUserInfo(postdata);
  if(result.success==true){
    loading.value=false;
    showSuccessToast('修改成功');
    router.back();
  }else{
     loading.value=false
    showFailToast(result.msg);
  }
}
// 校验身份证号
const validator= (val:any)=> {
  const card15 = /^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}[0-9Xx]$/
  const card18 = /^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$/
  return card15.test(val) || card18.test(val)
}
// 初始化需要处理的信息
onMounted(() => {
  loading.value=true;
  getUser();
});
// 跳转至变更手机号码界面
const goChangePhone=()=>router.push('/changePhone');
</script>

<style scoped>
.userInfoCard{
  margin: 20px 10px 10px;
  padding: 10px;
  background: #ffffff;
  border-radius: 5px;
}
</style>
