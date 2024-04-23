<template>
  <canvas id="canvas"></canvas>
  <div class="lonContainer">

    <el-row>
      <!-- <el-col :span="4" :xs="0"> -->
        <!-- <div class="leftTitle">欢迎使用VueAdmin</div> -->
      <!-- </el-col> -->
      <el-col :span="24" :xs="24">
        <!-- <el-form class="loginForm" :model="loginForm">
          <h1>Hello</h1>
          <h2>欢迎使用VueAdmin</h2>
          <el-form-item>
            <el-input v-model="loginForm.username" />
          </el-form-item>
          <el-form-item>
            <el-input type="password" show-password v-model="loginForm.password" />
          </el-form-item>
          <el-form-item>
            <el-button class="loginBtn" type="primary" size="default" @click="login">登录</el-button>
          </el-form-item>
        </el-form> -->
        
	<div class="section">
		<div class="container">
			<div class="row full-height justify-content-center">
				<div class="col-12 text-center align-self-center py-5">
					<div class="section pb-5 pt-5 pt-sm-2 text-center">
						<h6 class="mb-0 pb-3"><span>登 录 </span><span>注 册</span></h6>
			          	<input class="checkbox" type="checkbox" id="reg-log" name="reg-log"/>
			          	<label for="reg-log"></label>
						<div class="card-3d-wrap mx-auto">
							<div class="card-3d-wrapper">
								<div class="card-front">
									<div class="center-wrap">
										<div class="section text-center">
											<h4 class="mb-4 pb-3">登 录 </h4>
											<div class="form-group">
												<input type="email" name="logemail" class="form-style" placeholder="账号" id="logemail" autocomplete="off">
                        <i class="input-icon uil uil-user"></i>
											</div>	
											<div class="form-group mt-2">
												<input type="password" name="logpass1" class="form-style" placeholder="密码" id="logpass1" autocomplete="off">
												<i class="input-icon uil uil-lock-alt"></i>
											</div>
											<a href="#" class="btn mt-4"  @click="login">登 录</a>
                            				<p class="mb-0 mt-4 text-center"><a href="#0" class="link">忘记密码?</a></p>
				      					</div>
			      					</div>
			      				</div>
								<div class="card-back">
									<div class="center-wrap">
										<div class="section text-center">
											<h4 class="mb-4 pb-3">注册</h4>
											<div class="form-group">
												<input type="text" name="logname" class="form-style" placeholder="请输入账户名" id="logname" autocomplete="off">
													<i class="input-icon uil uil-user"></i>
											</div>	
											<!-- <div class="form-group mt-2">
												<input type="email" name="logemail" class="form-style" placeholder="Your Email" id="logemail" autocomplete="off">
												<i class="input-icon uil uil-at"></i>
											</div>	 -->
											<div class="form-group mt-2">
												<input type="password" name="logpass" class="form-style" placeholder="请输入密码" id="logpass" autocomplete="off">
												<i class="input-icon uil uil-lock-alt"></i>
											</div>
											<div class="form-group mt-2">
												<input type="password" name="logpass3" class="form-style" placeholder="请再次输入密码" id="logpass3" autocomplete="off">
												<i class="input-icon uil uil-lock-alt"></i>
											</div>
											<a href="#" class="btn mt-4">注 册</a>
				      					</div>
			      					</div>
			      				</div>
			      			</div>
			      		</div>
			      	</div>
		      	</div>
	      	</div>
	    </div>
	</div>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref, onMounted, onBeforeUnmount } from 'vue'
import useUserStore from '@/store/module/user';

let userStore = useUserStore()

const timeoutId = ref<number>();

let loginForm = reactive<LoginForm>({
  username: 'admin',
  password: '123',
})

const login = () => {
  // console.log(loginForm.value)
  console.log('aaa')
  clearTimeout(timeoutId.value);
  userStore.userLogin(loginForm)
  userStore.getUserInfo(loginForm.username,loginForm.password)
  // console.log('aa')
}
onMounted(() => {
  // console.log('aaa');

  // init()
})


// const ctx = ref('')
const ctx = ref('')
const canvas = reactive({})
// console.log(canvasWH.width);


const init = () => {
  // console.log('aaa');
  var canvas = document.getElementById("canvas");
  var ctx = canvas.getContext("2d");

  canvas.width = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
  canvas.height = window.innerHeight || document.documentElement.clientHeight || document.body.clientHeight;

  // window.onresize = resize;
  console.log(' canvas.width', canvas.width);
  console.log(' canvas.height', canvas.height);
  var RAF = (function () {
    return window.requestAnimationFrame || window.webkitRequestAnimationFrame || window.mozRequestAnimationFrame || window.oRequestAnimationFrame || window.msRequestAnimationFrame || function (callback) {
      window.setTimeout(callback, 1000 / 60);
    };
  })();

  // 鼠标活动时，获取鼠标坐标
  var warea = { x: null, y: null, max: 20000 };
  window.onmousemove = function (e) {
    e = e || window.event;

    warea.x = e.clientX;
    warea.y = e.clientY;
  };
  window.onmouseout = function (e) {
    warea.x = null;
    warea.y = null;
  };

  // 添加粒子
  // x，y为粒子坐标，xa, ya为粒子xy轴加速度，max为连线的最大距离
  var dots = [];
  for (var i = 0; i < 300; i++) {
    var x = Math.random() * canvas.width;
    var y = Math.random() * canvas.height;
    var xa = Math.random() * 2 - 1;
    var ya = Math.random() * 2 - 1;

    dots.push({
      x: x,
      y: y,
      xa: xa,
      ya: ya,
      max: 6000
    })
  }

  // 延迟100秒开始执行动画，如果立即执行有时位置计算会出错
  timeoutId.value = setTimeout(function () {
    animate();
  }, 100);

  // 每一帧循环的逻辑
  const animate = () => {
    var canvas = document.getElementById("canvas");
    var ctx = canvas.getContext("2d");
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    // 将鼠标坐标添加进去，产生一个用于比对距离的点数组
    var ndots = [warea].concat(dots);

    dots.forEach(function (dot) {

      // 粒子位移
      dot.x += dot.xa;
      dot.y += dot.ya;

      // 遇到边界将加速度反向
      dot.xa *= (dot.x > canvas.width || dot.x < 0) ? -1 : 1;
      dot.ya *= (dot.y > canvas.height || dot.y < 0) ? -1 : 1;

      // 绘制点
      ctx.fillRect(dot.x - 0.5, dot.y - 0.5, 1, 1);

      // 循环比对粒子间的距离
      for (var i = 0; i < ndots.length; i++) {
        var d2 = ndots[i];

        if (dot === d2 || d2.x === null || d2.y === null) continue;

        var xc = dot.x - d2.x;
        var yc = dot.y - d2.y;

        // 两个粒子之间的距离
        var dis = xc * xc + yc * yc;

        // 距离比
        var ratio;

        // 如果两个粒子之间的距离小于粒子对象的max值，则在两个粒子间画线
        if (dis < d2.max) {

          // 如果是鼠标，则让粒子向鼠标的位置移动
          if (d2 === warea && dis > (d2.max / 2)) {
            dot.x -= xc * 0.03;
            dot.y -= yc * 0.03;
          }

          // 计算距离比
          ratio = (d2.max - dis) / d2.max;

          // 画线
          ctx.beginPath();
          ctx.lineWidth = ratio / 2;
          ctx.strokeStyle = 'rgba(255,255,255,' + (ratio + 0.2) + ')';
          ctx.moveTo(dot.x, dot.y);
          ctx.lineTo(d2.x, d2.y);
          ctx.stroke();
        }
      }

      // 将已经计算过的粒子从数组中删除
      ndots.splice(ndots.indexOf(dot), 1);
    });

    RAF(animate);
  }
}

// 在组件销毁前执行清理操作
onBeforeUnmount(() => {
  // 清除 setTimeout
  if (timeoutId.value) {
    clearTimeout(timeoutId.value);
  }
});

</script>
<style scoped lang='scss'>
@import './login.css';
@import url('https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.5.0/css/bootstrap.min.css');
@import url('https://unicons.iconscout.com/release/v2.1.9/css/unicons.css');
// canvas {
//   background-color: #000;
//   width: 100%;
//   height:100vh
// }
#canvas {
  width: 100%;
  height: 99vh;
  overflow: hidden;
  position: absolute;
  z-index: 1;
  background: url('../../assets/images/bg.gif') no-repeat;
  background-size: cover;


}

.lonContainer {
  width: 100%;
  height: 100vh;
  // z-index: 99;
  position: fixed;
  left: 0;
  top: 0;
  // background: $base-color;
  z-index: 2;

  .loginForm {
    position: relative;
    top: 30vh;
    width: 80%;
    background: url('@/assets/images/login_form.png') no-repeat;
    background-size: cover;
    padding: 40px;
    z-index: 2;

    h1 {
      color: #fff;
      font-size: 40px;
    }

    h2 {
      color: #fff;
      font-size: 20px;
      margin: 20px 0;
    }

    .loginBtn {
      width: 100%;
      z-index: 9;

    }
  }

  .leftTitle {
    width: 100%;
    height: 100vh;
    display: flex;
    // justify-content: center;
    margin-left: 48%;
    align-items: center;
    color: #FFA500;
    font-size: 48px;
  }
}
</style>