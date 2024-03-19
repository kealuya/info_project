import { createRouter, createWebHistory } from 'vue-router'
import InvoiceList from '../pages/invoice/index.vue'
import error from '../pages/404/index.vue';
import {userInfoData} from "../store";
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {path: '/',redirect:'/home'},
    {
      path:'/home',
      name: 'home',
      component: InvoiceList,
    },
    {
      path: '/404',
      name: '404',
      component:error,
    },
    {
      path: '/addList',
      name: 'addList',
      component: () => import('../pages/addInvoice/index.vue'),
    },
    {
      path: '/addInvoiceFrom',
      name: 'addInvoiceFrom',
      component: () => import('../pages/addInvoiceFrom/index.vue'),
    },
    {
      path: '/detail',
      name: 'detail',
      component: () => import('../pages/invoiceDetail/index.vue'),
    },
    {
      path: '/iframe',
      name: 'iframe',
      component: () => import('../pages/invoiceIframe/index.vue'),
    },
  ]
})
router.beforeEach((to,from,next)=>{
  // 模拟url登陆地址导航栏截取
  //FIXME 模拟登录地址
  const url=''
  let urlOne = url.split('?');
  let userInfo=urlOne[1].split('&');
  let userId =userInfo[0].replaceAll('userid=','');
  let qyId=userInfo[1].replaceAll('school=','');
  const userInfoState: any = userInfoData();
  userInfoState.setData({
    empCode:userId,
    corpCode:qyId
  });
  // 使用 watch 方法监听 userInfoState 对象的变化
  next();
})

export default router
