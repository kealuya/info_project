import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import { constantRoutes } from "./router";
//引入进度条插件
// @ts-ignore
import NProgress from "nprogress";
import "nprogress/nprogress.css";
let router = createRouter({
  history: createWebHashHistory(),
  routes: constantRoutes,
  // scrollBehavior() {
  //   return {
  //     left: 0,
  //     top: 0,
  //   };
  // },
});
// 进度条配置
NProgress.configure({
  easing: "ease", // 动画方式
  speed: 500, // 递增进度条的速度
  showSpinner: false, // 是否显示加载 icon
  trickleSpeed: 200, // 自动递增间隔
  minimum: 0.3, // 初始化时的最小百分比
});
//路由守卫
router.beforeEach((to, from, next) => {
  // to and from are both route objects. must call `next`.
  NProgress.start(); // 进度条开始
  //根据token判断是否登录
  const token: string | null = localStorage.getItem("token");
  if (token) {
    if (to.path == "/login") {
      next();
    } else {
      next();
    }
  } else {
    //没有token
    if (to.path === "/login") {
      next();
      return;
    }
    next({ path: "/login" });
  }
});
router.afterEach(() => {
  NProgress.done(); // 进度条结束
});

export default router;
