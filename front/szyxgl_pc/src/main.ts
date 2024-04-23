import { createApp } from "vue";
import "./styles/index.scss";
import App from "./App.vue";
// 引入elementplus
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
//@ts-ignore忽略当前文件ts类型的检测否则有红色提示(打包会失败)
import zhCn from "element-plus/dist/locale/zh-cn.mjs";
// import router from "./router";
// 引入仓库
import pinia from "./store";
//全局引入icon
// import * as ElementPlusIconsVue from "@element-plus/icons-vue";
//引入router
import router from "./router";
//引入echarts
import * as echarts from 'echarts'
// 全局注册svg
import "virtual:svg-icons-register";
import gloablComponent from "@/components/index";
//暗黑模式
import 'element-plus/theme-chalk/dark/css-vars.css'
import '../src/styles/index.scss';
// import './styles/element/index.scss'z
const app = createApp(App);

//引入pinia
app.use(pinia);
//引入ElementPlus
app.use(ElementPlus, {
  locale: zhCn,
});
//引入svg
app.use(gloablComponent);
// 引入路由
app.use(router);
app.mount("#app");
