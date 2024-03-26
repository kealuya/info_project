<template>
  <div>
    <NavBar title="变更手机号"/>
    <div style="margin-top: 35px;margin-bottom: 20px">
      <van-row justify="center" align="center">
        <van-image :src="changePhone" width="150" height="150"/>
      </van-row>
    </div>
   <div v-if="phoneCode=='old'">
     <div class="tc">
       <span class="graytxt f14">为了您的账户安全，需要验证您的手机</span>
     </div>
     <div class="tc mt10">
       <van-form @submit="onSubmit">
         <van-row justify="center" align="center">
           <span class="f14">
          您的手机号：
          <span v-if="showPhone">
            {{ userInfoData.userInfo.phone.replace(/^(\d{3})\d{4}(\d{4})$/, '$1****$2') }}
          </span>
          <span v-else>
            {{ userInfoData.userInfo.phone}}
          </span>
        </span>
           <span class="ml10">
          <van-icon  @click="showPhoneOver('0')" v-if="showPhone" name="closed-eye" size="25"  />
          <van-icon  @click="showPhoneOver('1')" v-else name="eye-o" size="25"  color="#0088ff" />
        </span>
         </van-row>
         <div class="mt10">
           <van-cell-group inset>
             <van-field
                 v-model="code"
                 center
                 clearable
                 label="短信验证码"
                 placeholder="请输入短信验证码"
                 :rules="[{ required: true, message: '请输入短信验证码' }]"
             >
               <template #button>
                 <van-button size="small" v-if="timer" @click="getMessage" type="primary">发送验证码</van-button>
                 <van-button size="small" v-else disabled type="primary">{{ses}} 秒后重试</van-button>
               </template>
             </van-field>
           </van-cell-group>
         </div>
         <div style="margin: 30px;">
           <van-button  block type="primary" native-type="submit">
             提交
           </van-button>
         </div>
         <div class="tc f12">
           <span>(无法验证？ 请 <span @click="goHelpFeedback" class="linkTo">联系客服</span> )</span>
         </div>
       </van-form>
     </div>
   </div>
    <div v-else>
      <div class="tc">
        <span class="graytxt f14">更换手机号码后，下次登陆可以使用新手机号码进行登陆</span>
      </div>
      <div class="tc mt10">
        <van-form @submit="updatePhoneSubmit">
          <div class="mt10">
            <van-cell-group inset>
              <van-field
                  v-model="newPhone"
                  name="手机号码"
                  label="手机号码"
                  placeholder="请输入新手机号码"
                  :rules="[{ required: true, message: '请输入新手机号码' }]"
              />
              <van-field
                  v-model="newCode"
                  center
                  clearable
                  label="短信验证码"
                  placeholder="请输入短信验证码"
                  :rules="[{ required: true, message: '请输入短信验证码' }]"
              >
                <template #button>
                  <van-button size="small" v-if="timer" @click="getNewMessage" type="primary">获取验证码</van-button>
                  <van-button size="small" v-else disabled type="primary">{{newSes}} 秒后重试</van-button>
                </template>
              </van-field>
            </van-cell-group>
          </div>
          <div style="margin: 30px;">
            <van-button  block type="primary" native-type="submit" @click="updatePhoneSubmit">
              提交
            </van-button>
          </div>
          <div class="tc f12">
            <span>(无法验证？ 请 <span @click="goHelpFeedback" class="linkTo">联系客服</span> )</span>
          </div>
        </van-form>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import NavBar from '../../components/navBar/index.vue'
import changePhone from '../../assets/img/xiugaishouji.png'
import {inject, ref} from "vue";
import {getMessageCaptcha} from "../../services/home";
import {modifyPhoneNum, phoneVerification,modifyUserInfo} from '../../services/user'
import {showToast} from "vant";
import router from "../../router";
import {showFailToast} from "vant/lib/toast/function-call";
let userInfoData: any = inject("userInfo"); //用户信息
let showPhone = ref(true);  // 是否展示手机号
const timer =ref(true);
let code=ref(''); // 手机验证码
const ses=ref(0); // 记录短信验证码时间
const newSes=ref(0); // 记录新手机号验证码时间
let newCode=ref(''); // 新手机号的验证码
let phoneCode=ref('old'); // 手机验证码阶段控制  一开始 old 第一步校验完成后跳转至第二阶段 new
let newPhone=ref(''); // 要更换的新手机号
// 提交进行更新手机号
const updatePhoneSubmit=async()=>{
  let data={
    'phone':newPhone.value,
    'validateCode':newCode.value
  };
  let result: any = await modifyPhoneNum(data);
  if(result.success==true){
    showFailToast(`更改成功：${result.msg}`);
    phoneCode.value='old';
    localStorage.clear()
    sessionStorage.clear();
    location.reload()
  }else{
    showFailToast(`校验失败：${result.msg}`);
  }
};
// 提交跳转至更换新的手机号中
const onSubmit = async() => {
  let data={
    'phone':userInfoData.userInfo.phone,
    'validateCode':code.value
  };
  let result: any = await phoneVerification(data);
  if(result.success==true){
    phoneCode.value='new';
  }else{
    showFailToast(`校验失败：${result.msg}`);
  }
};
// 根据手机号获取验证码
const getMessage = async() => {
  let result: any = await getMessageCaptcha({
    'phone': userInfoData.userInfo.phone
  });
  if (result.success) {
    showToast(result.msg);
    timer.value=false;
    ses.value=59;
    setInterval(() => {
      if (ses.value > 0 && timer.value==false) {
        ses.value--;
      } else {
        timer.value= true;
        clearInterval(ses.value);
        ses.value=60;
      }
    }, 1000)

  } else {
    showToast(result.msg);
  }
};
// 根据新手机号获取验证码
const getNewMessage = async() => {
  if(newPhone.value!=''){
    // 17862842258
    let result: any = await getMessageCaptcha({
      'phone': newPhone.value
    });
    if (result.success) {
      showToast(result.msg);
      timer.value=false;
      newSes.value=59;
      setInterval(() => {
        if (ses.value > 0 && timer.value==false) {
          newSes.value--;
        } else {
          timer.value= true;
          clearInterval(ses.value);
          newSes.value=60;
        }
      }, 1000)

    } else {
      showToast(result.msg);
    }
  }else{
    showToast('请输入新的手机号再点击获取验证码哦');
  }

};
// 控制手机号是否显隐藏
const showPhoneOver =(key:any)=>{
  if(key=='0'){
    showPhone.value=false;
  }else{
    showPhone.value=true;
  }
}
// 跳转到联系客服页面
const goHelpFeedback =()=>router.push('/helpFeedback');
</script>

<style scoped>
  .linkTo{
    color:#0088ff;
    border-bottom:1px #0088ff;
  }
</style>
