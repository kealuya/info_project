import request from "./request";

//获取账本列表
export function getAccountBookList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/accountBook/getAccountBookList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}


//删除账本
export function deleteBillInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/accountBook/deleteBillInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

//获取费用关联的发票信息
export function getInvoiceBillList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/accountBook/getInvoiceBillList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

//获取审批信息列表
export function getApprovalFlow(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/approval/getApprovalFlow',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

//获取发票列表
export function getInvoiceWbList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/accountBook/getInvoiceWbList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 新增费用
export function addBillInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/accountBook/addBillInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 账本列表
export function getBillList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/bxBill/getBillList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}



