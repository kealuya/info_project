import { MockMethod } from "vite-plugin-mock"
 
const mock: Array<MockMethod> = [
    //登录
    {
      // 接口路径
      url: '/api/login',

      // 接口方法
      method: 'get',

      // 返回数据
      response: () => {
          return {
              status: 200,
              message: 'success',
              token:'123',
              data: 'Hello World'
          }
      }
  }, 
    /**
     * 路由数据接口
     */
    {
        url: '/api/routes',
        method: 'get',
        response: () => {
            // 路由
            const routes =   [
             {
             "parentId": 0,
             "id": 1,
             "icon": "Document",
             "title": "菜单",
             "name": "Menu",
             "path": "/menu",
             "component": "Layout",
             "hideInMenu": false,
             "children": [
               {
                 "parentId": 1,
                 "icon": "Document",
                 "title": "二级菜单",
                 "name": "menu1",
                 "path": "/menu/menu1",
                 "component": "/menu/menu1",
                 "hideInMenu": false,
                 "children": [],
                 "id": 11
               },
               {
                 "id": 12,
                 "parentId": 1,
                 "icon": "Document",
                 "title": "二级菜单",
                 "name": "Menu2",
                 "hideInMenu": false,
                 "path": "/menu",
                 "component": "Layout",
                 "children": [{
                   "id": 122,
                   "parentId": 12,
                   "icon": "Document",
                   "title": "三级菜单",
                   "name": "Menu12",
                   "hideInMenu": false,
                   "path": "/menu/menu2",
                   "component": "/menu/menu2",
                   "children": [],
                 },
               ],
               },
       
             ],
           }, {
             "id": 2,
             "parentId": 0,
             "icon": "Document",
             "title": "表格部分",
             "name": "Tables",
             "path": "/tables",
             "component": "Layout",
             "hideInMenu": false,
             "children": [{
               "parentId": 2,
               "icon": "Document",
               "hideInMenu": false,
               "title": "表格1",
               "name": "tables1",
               "path": "/tables/tables1",
               "component": "/tables/tables1",
               "children": [],
               "id": 21
             },
             {
               "parentId": 2,
               "icon": "Document",
               "title": "表格2",
               "name": "tables2",
               "path": "/tables/tables2",
               "hideInMenu": false,
               "component": "/tables/tables2",
               "children": [],
               "id": 22
             },
             {
               "parentId": 2,
               "icon": "Document",
               "title": "分页",
               "name": "pagination",
               "path": "/tables/pagination",
               "hideInMenu": false,
               "component": "/tables/pagination",
               "children": [],
               "id": 23
             },
             {
               "parentId": 2,
               "icon": "Document",
               "title": "导出excel",
               "name": "exportExcel",
               "path": "/tables/exportExcel",
               "hideInMenu": false,
               "component": "/tables/exportExcel",
               "children": [],
               "id": 24
             },
             {
               "parentId": 2,
               "icon": "Document",
               "title": "导入excel",
               "name": "importExcel",
               "path": "/tables/importExcel",
               "hideInMenu": false,
               "component": "/tables/importExcel",
               "children": [],
               "id": 25
             },
       
             ],
       
           },
           {
             "id": 3,
             "parentId": 0,
             "icon": "Document",
             "title": "父子传参",
             "name": "FCcommunicate",
             "path": "/communicate",
             "component": "Layout",
             "hideInMenu": false,
             "children": [{
               "parentId": 3,
               "icon": "Document",
               "hideInMenu": false,
               "title": "父传子",
               "name": "father",
               "path": "/communicate/fathers",
               "component": "/communicate/fathers",
               "children": [],
               "id": 31
             },
             {
               "parentId": 3,
               "icon": "Document",
               "title": "子传父",
               "name": "child",
               "path": "/communicate/childs",
               "hideInMenu": false,
               "component": "/communicate/childs",
               "children": [],
               "id": 32
             },
             ],
       
       
           },
           {
             "id": 4,
             "parentId": 0,
             "icon": "Document",
             "title": "地图",
             "name": "Map",
             "path": "/map",
             "component": "Layout",
             "hideInMenu": false,
             "children": [{
               "parentId": 4,
               "icon": "Document",
               "hideInMenu": false,
               "title": "地图",
               "name": "map1",
               "path": "/map/index1",
               "component": "/map/index1",
               "children": [],
               "id": 41
             }
             ],
           },
           {
             "id": 7,
             "parentId": 0,
             "icon": "Document",
             "title": "监听",
             "name": "Listen",
             "path": "/listen",
             "component": "Layout",
             "hideInMenu": false,
             "children": [{
               "parentId": 7,
               "icon": "Document",
               "hideInMenu": false,
               "title": "监听demo1",
               "name": "watch1",
               "path": "/listen/watch1",
               "component": "/listen/watch1",
               "children": [],
               "id": 71
             },
             {
               "parentId": 7,
               "icon": "Document",
               "hideInMenu": false,
               "title": "监听demo2",
               "name": "watch2",
               "path": "/listen/watch2",
               "component": "/listen/watch2",
               "children": [],
               "id": 72
             },
             {
               "parentId": 7,
               "icon": "Document",
               "hideInMenu": false,
               "title": "监听demo3",
               "name": "watch3",
               "path": "/listen/watch3",
               "component": "/listen/watch3",
               "children": [],
               "id": 73
             }
             ],
           },
           {
             "id": 8,
             "parentId": 0,
             "icon": "Document",
             "title": "动画",
             "name": "Animation",
             "path": "/animation",
             "component": "Layout",
             "hideInMenu": false,
             "children": [{
               "parentId": 8,
               "icon": "Document",
               "hideInMenu": false,
               "title": "动画1",
               "name": "animation1",
               "path": "/animation/animation1",
               "component": "/animation/animation1",
               "children": [],
               "id": 81
             },
             {
               "parentId": 8,
               "icon": "Document",
               "hideInMenu": false,
               "title": "动画2",
               "name": "animation2",
               "path": "/animation/animation2",
               "component": "/animation/animation2",
               "children": [],
               "id": 82
             }
             ],
           },
           {
             "id": 9,
             "parentId": 0,
             "icon": "Document",
             "title": "关于我们",
             "name": "AboutUs",
             "path": "/aboutUs",
             "component": "Layout",
             "hideInMenu": false,
             "children": [{
               "parentId": 9,
               "icon": "Document",
               "hideInMenu": false,
               "title": "公司简介",
               "name": "aboutUs",
               "path": "/aboutUs/companyProfile",
               "component": "/aboutUs/companyProfile",
               "children": [],
               "id": 91
             }
             ],
           },
           ]
 
            return {
                status: 200,
                message: 'success',
                data: routes
            }
        }
    }
]
 
export default mock