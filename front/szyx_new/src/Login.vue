<template>
  <div class="headLogin">
    <div class="haedTitle">
      <div class="headline">欢迎使用「源易控」</div>
      <span class="headContext">高效、合规、降本，提升企业管理价值</span>
    </div>
  </div>
  <div class="mt10">
    <van-form @submit="onSubmit" ref="loginForm" autocomplete="off">
      <van-cell-group inset>
        <van-field
            v-model="phone"
            name="phone"
            label="手机号"
            placeholder="请填写手机号"
            :rules="[{ required: true, message: '请填写手机号' }]"
            autocomplete="off"
        />
        <van-field
            v-model="password"
            name="validateCode"
            label="验证码"
            placeholder="请填写验证码"
            length="6"
            :rules="[{ required: true, message: '请填写验证码' }]"
            autocomplete="off"
        >
          <template #button>
            <van-button v-if="timer" size="small" type="primary" @click="getMessage">获取验证码</van-button>
            <van-button v-else size="small" type="primary" disabled > {{ses}} 秒后重试</van-button>
          </template>
        </van-field>
      </van-cell-group>
      <div style="margin: 20px;">
        <van-button block type="primary" native-type="submit">
          登录
        </van-button>
      </div>
    </van-form>
    <div class="mt10 tc">
      <span class="f12">
        点击登录，即表示同意
        <span @click.stop="onPolicy" style="color:#0088FF;">《隐私政策》</span>和
        <span @click.stop="onUserAgreement"  style="color:#0088FF;">《用户协议》</span>
      </span>
    </div>
<!--    <div class="mt15 tc">-->
<!--      <span style="font-size: 14px;color: #999999">-&#45;&#45;&#45;&#45;&#45;&#45;&#45;&#45;&#45;&#45; 未开通账号 -&#45;&#45;&#45;&#45;&#45;&#45;&#45;&#45;&#45;&#45;</span>-->
<!--    </div>-->
<!--    <div style="margin: 16px;">-->
<!--      <van-button round plain block type="primary" native-type="submit">-->
<!--        体验账号-->
<!--      </van-button>-->
<!--    </div>-->
  </div>
</template>

<script setup lang="ts">
import {ref} from "vue";
import {useRouter} from "vue-router";
import {userInfoData} from "./store";
import {getMessageCaptcha, messageLogin} from "./services/home";
import {showToast, showSuccessToast} from "vant";
const staticCodeIconList=ref([
  'plane','hotel','invoice','car','train','buy','log',
]);
const router = useRouter();
const phone = ref('');
const password = ref('');
const timer =ref(true);
const ses=ref(0);
let loginForm = ref();
// 是否点击为
const onSubmit = async (values: any) => {
  let result: any = await messageLogin(values);
  if (result.msg!='error') {
      let data = result.data;
      await localStorage.setItem('token', data.jwt);
      await localStorage.setItem('refreshToken', data.refreshToken);
      const userInfoState: any = userInfoData();
        userInfoState.setData(data.user[0]);
      let functionConfigObj: any = {};
      data.functionConfig.map((item: any) => {
          functionConfigObj[item.ywTypeID] = item;
      });
      userInfoState.setFunctionConfig(functionConfigObj);
      userInfoState.setUserTravelConfig(data.userTravelConfig);
      showSuccessToast('登录成功');
    sessionStorage.setItem('myIconCodeList', JSON.stringify(staticCodeIconList.value));
      await router.replace('/home');
  } else {
      showToast(result.data);
  }
};
const onPolicy = (values: any) => {
  router.replace('/privacyPolicy');
};
const getMessage = () => {
    loginForm.value.validate('phone').then(async () => {
        let result: any = await getMessageCaptcha({
            'phone': phone.value
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
    })
};
const onUserAgreement = (values: any) => {
  router.replace('/userAgreement');
};
</script>

<style lang="less" scoped>
.headLogin {
  width: 100vw;
  height: 230px;
  background: url('./assets/img/login_bg.png') no-repeat center;
  background-size: 100vw 250px;

  .haedTitle {
    position: absolute;
    top: 155px;
    left: 20px;

    .headline {
      font-size: 26px;
      padding-bottom: 5px;
      font-weight: bold;
    }

    .headContext {
      color: #999999;
      margin-top: 20px;
      font-size: 12px;
    }
  }
}
</style>
