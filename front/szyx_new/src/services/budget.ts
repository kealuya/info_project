import request from "./request";
// 获取预算编制列表
export function getBudgetList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/budget/getBudgetList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取预算编制详情
export function getBudgetDetail(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/budget/getBudgetDetail',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
//获取预算科目列表
export function getSubjectList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/budget/getSubjectList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取部门成本中心列表
export function getBudgetDeptList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/budget/getBudgetDeptList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取项目成本中心列表
export function getCostCenterList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/costCenter/getCostCenterList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
