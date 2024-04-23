// 对外暴露
import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
export const constantRoutes = [
  {
    path: "/login",
    name: "login",
    meta: {
      title: "登录",
      hidden: true,
      icon: "Document",
    },
    component: () => import("@/views/login/index.vue"),
  },

  {
    path: "/",
    name: "main",
    component: () => import("../views/index.vue"),
    redirect: { path: "/home" },
    meta: {
      title: "首页",
      hidden: false,
      icon: "house",
    },
    children: [
      {
        path: "/home",
        component: () => import("../views/home/index.vue"),
        name: "Home",
        meta: {
          title: "首页",
          hiddenInMenu: false,
          closeTagMenu: false,
          icon: "Refresh",
        },
      },
    ],
  },
  {
    path: "/userPage",
    name: "user",
    component: () => import("../views/index.vue"),
    redirect: { path: "/home" },
    meta: {
      title: "员工管理",
      hidden: false,
      icon: 'User',
    },
    children: [
      {
        path: "/roleAdmin",
        component: () => import("../views/user/roleAdmin.vue"),
        name: "RoleAdmin",
        meta: {
          title: "角色管理",
          hiddenInMenu: false,
          closeTagMenu: false,
          icon: "Refresh",
        },
      },
      {
        path: "/workData",
        component: () => import("../views/user/workData.vue"),
        name: "WorkData",
        meta: {
          title: "员工工作统计",
          hiddenInMenu: false,
          closeTagMenu: false,
          icon: "Connection",
        },
      }
    ],
  },
  {
    path: "/",
    name: "customer",
    component: () => import("../views/index.vue"),
    redirect: { path: "/customer" },
    meta: {
      title: "客户管理",
      hidden: false,
      icon: "Tickets",
    },
    children: [
      {
        path: "/customer",
        component: () => import("../views/customer/index.vue"),
        name: "Customer",
        meta: {
          title: "客户管理",
          hiddenInMenu: false,
          closeTagMenu: false,
          icon: "Refresh",
        },
      },
    ],
  },
  {
    path: "/",
    name: "marketing",
    component: () => import("../views/index.vue"),
    redirect: { path: "/home" },
    meta: {
      title: "营销管理",
      hidden: false,
      icon: "Connection",
    },
    children: [
      {
        path: "/marketing",
        component: () => import("../views/marketing/index.vue"),
        name: "Marketing",
        meta: {
          title: "营销管理",
          hiddenInMenu: false,
          closeTagMenu: false,
          icon: "Refresh",
        },
      },
    ],
  },
  {
    path: "/",
    name: "knowledg",
    component: () => import("../views/index.vue"),
    redirect: { path: "/knowledg" },
    meta: {
      title: "知识库管理",
      hidden: false,
      icon: "ZoomIn",
    },
    children: [
      {
        path: "/knowledg",
        component: () => import("../views/knowledg/index.vue"),
        name: "Knowledg",
        meta: {
          title: "知识库管理",
          hiddenInMenu: false,
          closeTagMenu: false,
          icon: "Refresh",
        },
      },
    ],
  },

  {
    path: "/404",
    name: "404",
    meta: {
      title: "404",
      hidden: true,
    },
    component: () => import("@/views/404/index.vue"),
  },
  {
    path: "/:pathMatch(.*)*",
    redirect: "/404",
    name: "Any",
    meta: {
      title: "错误路由",
      hidden: true,
    },
  },
];
