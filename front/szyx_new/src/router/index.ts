import {createRouter, createWebHashHistory} from 'vue-router'
import InvoiceList from '../pages/invoiceFolder/invoice/index.vue'
import error from '../pages/404/index.vue';
import {userInfoData} from "../store";

// @ts-ignore
const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/', redirect: '/login' },
    {
      path: '/home',
      name: '首页',
      component: () => import('../HomeSetting.vue'),
      children: [
        // {
        //   path:'/home',
        //   name: 'home',
        //   component: () => import('../pages/home/index.vue'),
        // },
        {
          path:'/homeNew',
          name: 'homeNew',
          component: () => import('../pages/homeMetting/home.vue'),
        },
          // 业务页面
        {
          path: '/business',
          name: 'business',
          component: () => import('../pages/business/index.vue'),
        },
          //任务详情页面
        {
          path: '/taskDetail',
          name: 'taskDetail',
          component: () => import('../pages/taskDetail/index.vue'),
        },
        {
          path: '/taskDetail_two',
          name: 'taskDetail_two',
          component: () => import('../pages/taskDetail_two/index.vue'),
        },
        {
          path: '/taskDetail_three',
          name: 'taskDetail_three',
          component: () => import('../pages/taskDetail_three/index.vue'),
        },
        {
          path: '/taskDetail_four',
          name: 'taskDetail_four',
          component: () => import('../pages/taskDetail_four/index.vue'),
        },
        //完成任务
        {
          path: '/addTasks',
          name: 'addTasks',
          component: () => import('../pages/taskDetail/addTasks.vue'),
        },  {
          path: '/valueDetails',
          name: 'valueDetails',
          component: () => import('../pages/valueDetails/index.vue'),
        },
          //业务详情
        {
          path: '/businessDetail',
          name: 'businessDetail',
          component: () => import('../pages/businessDetail/index.vue'),
        },
          //task processing
        {
          path: '/taskList',
          name: 'taskList',
          component: () => import('../pages/taskList/index.vue'),
        },
         //任务处理
        {
          path: '/taskProcessing',
          name: 'taskProcessing',
          component: () => import('../pages/taskProcessing/index.vue'),
        },

        {
          path: '/task',
          name: 'taskDeatilD',
          component: () => import('../pages/taskDeatilD/index.vue'),
        },
          //价值申请
        // application
        {
          path: '/applicationValue',
          name: 'applicationValue',
          component: () => import('../pages/applicationValue/index.vue'),
        },
          //有录音的页面
        {
          path: '/speech',
          name: 'speech',
          component: () => import('../pages/speech/index.vue'),
        },
          //文档上传页面
        {
          path: '/documentUpload',
          name: 'documentUpload',
          component: () => import('../pages/documentUpload/index.vue'),
        },
        {
          path: '/personal',
          name: 'personal',
          component: () => import('../pages/personal/index.vue'),
        },
        {
          path: '/message',
          name: 'message',
          meta:{
            keepAlive:true
          },
          component: () => import('../pages/message/index.vue'),
        },
        {
          path: '/system',
          name: 'system',
          component: () => import('../pages/system/index.vue'),
        },
        {
          path: '/user',
          name: 'user',
          component: () => import('../pages/user/index.vue'),
        },
        {
          path: '/book',
          name: 'book',
          component: () => import('../pages/accountBook/add.vue'),
        },
      ],
      // meta: Promise
    },

    // 第三方业务-------机、火、酒、心意选
    {
      path:'/thirdPartyBusiness/:applicationType',
      name: 'thirdPartyBusiness',
      component: () => import('../pages/thirdPartyBusiness/businessTrip.vue'),
    },

    // 预算管理列表
    {
      path:'/budget',
      name: 'budget',
      component: () => import('../pages/budget/index.vue'),
    },
    // 预算详情
    {
      path:'/budgetDetail',
      name: 'budgetDetail',
      component: () => import('../pages/budget/detail/index.vue'),
    },
    // 成本中心
    {
      path:'/costCenter',
      name: 'costCenter',
      component: () => import('../pages/costCenter/index.vue'),
    },
    {
      path:'/dataTest',
      name: 'dataTest',
      component: () => import('../pages/data/test.vue'),
    },
    {
      path:'/about',
      name: 'about',
      component: () => import('../pages/about/index.vue'),
    },
    //个人使用统计
    {
      path:'/personalUsageStatistics',
      name: 'personalUsageStatistics',
      component: () => import('../pages/personalUsagesStatistics/index.vue'),
    },
    // {
    //   path:'/login',
    //   name: 'login',
    //   component: () => import('../Login.vue'),
    //   meta: {
    //     keepAlive: true  // 需要缓存的页面配置
    //   }
    // },
      //登录（新）
    {
      path:'/login',
      name: 'login',
      component: () => import('../pages/loginMetting/login.vue'),
      meta: {
        keepAlive: true  // 需要缓存的页面配置
      }
    },
    // 个人考勤统计
    {
      path:'/personalAttendance',
      name: 'personalAttendance',
      component: () => import('../pages/personalAttendance/index.vue'),
      meta: {
        keepAlive: true  // 需要缓存的页面配置
      }
    },
    // 修改用户信息
    {
      path: '/personalInfo',
      name: 'personalInfo',
      component: () => import('../pages/personalInfo/index.vue'),
    },
    // 更改用户手机号
    {
      path: '/changePhone',
      name: 'changePhone',
      component: () => import('../pages/changePhone/index.vue'),
    },
      // 申请单详情
    {
      path:'/applicationDetail',
      name: 'applicationDetail',
      component: () => import('../pages/application/detail/index.vue'),
    },
    // 隐私协议
    {
      path:'/privacyPolicy',
      name: 'privacyPolicy',
      component: () => import('../pages/privacyPolicy/index.vue'),
    },
    {
      path: '/sysMessageDetail/:htmlMessage/:newTitle/:newType/:createTime/:dept',
      name: 'sysMessageDetail',
      component: () => import('../pages/message/sysMessageDetail.vue'),
    },
    {
      path: '/approvalDetailsIframe/:url/:realHeight/:titleName',
      name: 'approvalDetailsIframe',
      component: () => import('../pages/approvalIframe/approvalDetailsIframe.vue'),
    },
    {
      path: '/dailyReimbursement/:bill',
      name: 'dailyReimbursement',
      component: () => import('../pages/accountBook/dailyReimbursement.vue'),
    },
      //账本----费用列表
    {
      path: '/costInformationSelectionPage/:bill',
      name: 'costInformationSelectionPage',
      component: () => import('../pages/accountBook/costInformationSelectionPage.vue'),
    },
      //账本-----发票列表
    {
      path: '/invoiceInformationList/:invoice',
      name: 'invoiceInformationList',
      component: () => import('../pages/accountBook/InvoiceInformationList.vue'),
    },
      //记账单
    {
      path: '/book',
      name: 'book',
      component: () => import('../pages/accountBook/add.vue'),
    },
    //记账单---列表
    {
      path: '/expressDeliveryOrCarPage/:applicationType',
      name: 'expressDeliveryOrCarPage',
      component: () => import('../pages/applicationFormList/expressDeliveryOrCarPage.vue'),
    },
    {
      path: '/IframePage/:barTitle/:iframeUrl',
      name: 'IframePage',
      component: () => import('../pages/iframe/businessIframePage.vue'),
    },
      //账户管理-----列表页
    {
      path: '/manageAccounts/:bankType',
      name: 'manageAccounts',
      component: () => import('../pages/manageAccounts/index.vue'),
    },
      //账户管理----绑定银行卡
    {
      path: '/bindBankCard/:accountObj/:isComponents',
      name: 'bindBankCard',
      component: () => import('../pages/manageAccounts/add.vue'),
    },
    {
      path: '/myDelegation',
      name: 'myDelegation',
      component: () => import('../pages/myDelegation/index.vue'),
    },
    {
      path: '/addMyDelegation/:typeName/:commissionDeadline',
      name: 'addMyDelegation',
      component: () => import('../pages/myDelegation/add.vue'),
    },
    {
      path: '/helpFeedback',
      name: 'helpFeedback',
      component: () => import('../pages/helpFeedback/index.vue'),
    },
      // 报销单
    {
      path: '/reimburse',
      name: 'reimburse',
      component: () => import('../pages/reimburse/index.vue'),
    },
    // 报销单
    {
      path: '/reimburse/detail',
      name: 'reimburse/detail',
      component: () => import('../pages/reimburse/detail/index.vue'),
    },
    // 新增日常报销
    {
      path: '/addDailyBx',
      name: 'addDailyBx',
      component: () => import('../pages/reimburse/addDailyBx/index.vue'),
      meta: {
        keepAlive: true  // 需要缓存的页面配置
      }
    },
    {
      path: '/addLoan/:selectAccountInformation',
      name: 'addLoan',
      component: () => import('../pages/reimburse/addLoan.vue'),
    },
    {
      path: '/addTravelSubsidies',
      name: 'addTravelSubsidies',
      component: () => import('../pages/reimburse/addTravelSubsidies.vue'),
    },
    {
      path: '/travelSubsidiesDetail/:applyNo',
      name: 'travelSubsidiesDetail',
      component: () => import('../pages/reimburse/travelSubsidiesDetail.vue'),
    },
    {
      path: '/loanDetail/:loanId',
      name: 'loanDetail',
      component: () => import('../pages/reimburse/loanDetail.vue'),
    },
    {
      // 申请单
      path:'/application',
      name: 'application',
      component: () => import('../pages/application/index.vue'),
    },
    {
      // 单据动态
      path:'/documentDynamics',
      name: 'documentDynamics',
      component: () => import('../pages/documentDynamics/index.vue'),
    },
    {
      // 请假申请单
      path:'/leaveApply',
      name: 'leaveApply',
      component: () => import('../pages/application/form/leave/index.vue'),
    },
      // 加班申请单
    {
      path:'/overWork',
      name: 'overWork',
      component: () => import('../pages/application/form/overWork/index.vue'),
    },
    // 用章申请单
    {
      path:'/useSeal',
      name: 'useSeal',
      component: () => import('../pages/application/form/useSeal/index.vue'),
    },
    // 考勤申请单
    {
      path:'/openCard',
      name: 'openCard',
      component: () => import('../pages/application/form/openCard/index.vue'),
    },
      // 物品领用
    {
      path:'/requisition',
      name: 'requisition',
      component: () => import('../pages/application/form/requisition/index.vue'),
    },
    // 外出办公
    {
      path:'/goingOutWork',
      name: 'goingOutWork',
      component: () => import('../pages/application/form/goingOutWork/index.vue'),
    },
    // 董事签批
    {
      path:'/director',
      name: 'director',
      component: () => import('../pages/application/form/director/index.vue'),
    },
      // 客户宴请
    {
      path:'/entertain',
      name: 'entertain',
      //@ts-ignore
      component: () => import('../pages/application/form/entertain/index.vue'),
    },
      // 会议记录
    {
      path:'/meeting',
      name: 'meeting',
      //@ts-ignore
      component: () => import('../pages/application/form/meeting/index.vue'),
    },
    // 物品申购
    {
      path:'/subscribe',
      name: 'subscribe',
      //@ts-ignore
      component: () => import('../pages/application/form/subscribe/index.vue'),
    },
    // 用车申请
    {
      path:'/useCar',
      name: 'useCar',
      //@ts-ignore
      component: () => import('../pages/application/form/useCar/index.vue'),
    },
    // 物流申请
    {
      path:'/logistics',
      name: 'logistics',
      //@ts-ignore
      component: () => import('../pages/application/form/logistics/index.vue'),
    },
    {
      // 企业数据统计
      path:'/enterpriseStatistics',
      name: 'enterpriseStatistics',
      //@ts-ignore
      component: () => import('../pages/enterpriseStatistics/index.vue'),
    },
      // 费用申请
    {
      path:'/cost',
      name: 'cost',
      //@ts-ignore
      component: () => import('../pages/application/form/cost/index.vue'),
    },
      // 差旅补助报销
    {
      path:'/travel',
      name: 'travel',
      //@ts-ignore
      component: () => import('../pages/application/form/travel/index.vue'),
    },
    {
      // 报销单
      path:'/reimburse',
      name: 'reimburse',
      component: () => import('../pages/reimburse/index.vue'),
    },
      // 用户协议
    {
      path:'/userAgreement',
      name: 'userAgreement',
      component: () => import('../pages/userAgreement/index.vue'),
    },
    {
      path: '/404',
      name: '404',
      component:error,
    },
    // 发票抬头列表
    {
      path: '/invoiceBilling',
      name: 'invoiceBilling',
      component: () => import('../pages/invoiceFolder/invoiceBillng/index.vue'),
    },
      // 发票抬头信息
    {
      path: '/billingInfo',
      name: 'billingInfo',
      component: () => import('../pages/invoiceFolder/billingInfo/index.vue'),
    },
    {
      path: '/invoiceList',
      name: 'invoiceList',
      component: () => import('../pages/invoiceFolder/invoice/index.vue'),
    },
    {
      path: '/addList',
      name: 'addList',
      component: () => import('../pages/invoiceFolder/addInvoice/index.vue'),
    },
    {
      path: '/addInvoiceFrom',
      name: 'addInvoiceFrom',
      component: () => import('../pages/invoiceFolder/addInvoiceFrom/index.vue'),
    },
    {
      path: '/detail',
      name: 'detail',
      component: () => import('../pages/invoiceFolder/invoiceDetail/index.vue'),
    },
    {
      path: '/iframe',
      name: 'iframe',
      component: () => import('../pages/invoiceFolder/invoiceIframe/index.vue'),
    },
    {
      path: '/accountBook',
      name: 'accountBook',
      component:()=> import('../pages/accountBook/index.vue'),
    }
  ]
})

// router.beforeEach((to,from,next)=>{
//   // 模拟url登陆地址导航栏截取
//   // const url='http://122.9.41.215/fk_move/MapprList?userid=1815&school=szht6666'
//   // let urlOne = url.split('?');
//   // let userInfo=urlOne[1].split('&');
//   // let userId =userInfo[0].replaceAll('userid=','');
//   // let qyId=userInfo[1].replaceAll('school=','');
//   const userInfoState: any = userInfoData();
//   const store = userInfoData()
//   if (to.path === '/login' || to.path === '/404'||to.path ==="/privacyPolicy"||to.path==='/userAgreement') {
//     next();
//   } else {
//     if (Object.keys(userInfoState.userInfo).length === 0) {
//       next('/login');
//     } else {
//       next();
//     }
//   }
//
// })

export default router
