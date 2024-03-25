import request from "./request";
// 获取账户列表
export function payeeList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/cwxtPay/payeeList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 添加账户信息
export function addPayeeInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/cwxtPay/addPayeeInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 删除或取消绑定账户信息
export function deletePayeeInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/cwxtPay/deletePayeeInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

export function updatePayeeInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/cwxtPay/updatePayeeInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取默认个人账户
export function getDefaultPay(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/cwxtPay/getDefaultPay',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}




