import request from "./request";
//发票列表
export function getInvoiceWbList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/invoice/getInvoiceWbList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
//微信签名
export function getWxConfig(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/invoice/getInvoiceWbList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
//发票图片识别
export function getInvOcrResultWb(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/invoice/getInvOcrResultWb',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({'success':false,'data':{}})
        })
    })
}
//发票PDF、OFD识别
export function getInvOcrResultByPdfWb(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/invoice/getInvOcrResultByPdfWb',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
//发票保存invVerifyWb
export function invVerifyWb(query :object) {
    return new Promise(function (resolve, reject) {
        request({
            url:'/invoice/invVerifyWb',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
//获取 微信token
export function getWxTokenTicket(query :object) {
    return new Promise(function (resolve, reject) {
        request({
            url:'/invoice/getAccess_tokenTicket',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
// 获取微信票务
export function getWxSignToken(query :object) {
    return new Promise(function (resolve, reject) {
        request({
            url:'https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx5276498b4494a34a&secret=e8bf9e603d92ddc5076459e673d49014',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取刷新的token
export function getToken(query :object) {
    return new Promise(function (resolve, reject) {
        request({
            url:'/invoice/getToken',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({'success':false,'data':{}})
        })
    })
}
//获取发票图片
export function getImgDetail(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/invoice/getInvoiceImgUrl',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
//删除发票
export function deleteInv(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/invoice/deleteInvoiceToWb',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
