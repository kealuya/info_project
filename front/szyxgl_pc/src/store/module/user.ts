import { defineStore } from "pinia";
import { login } from "@/api";
import { ElMessage } from "element-plus";
import router from "~/router";
//引入常量路由
import { constantRoutes } from "@/router/router";

//创建小仓库
let useUserStore = defineStore("User", {
  //存储数据的地方
  state: () => {
    return {
      // routes:constantRoutes, //路由表
      menuRouters: constantRoutes,

      userName:'',
      userPass:'',
      token:''
    };
  },
  //异步,逻辑区
  actions: {
    userLogin(loginForm: any) {
      // console.log('123',loginForm.username)
      if (loginForm.username == "admin" && loginForm.password == "123") {
        // console.log("登录成功");
        login().then((res) => {
          // console.log("res", res);
          localStorage.setItem("token", res.data.token);
          this.token = res.data.token
          // console.log('aaa',this.menuRouters)
          router.push("/home");
        });
      } else {
        console.log("密码错误");
        ElMessage.error("账号密码错误");
      }
    },
    getUserInfo(name:string,pass:string){
      this.userName = name 
      this.userPass = pass
    }
  },

  persist: {
    enabled: true, // 开启数据缓存
    strategies: [
      {
        storage: localStorage,
        paths: ['userName', 'userPass','token'],
      },
    ],
  },
  
});
export default useUserStore;
