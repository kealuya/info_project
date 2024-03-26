import request from "./request";

//提交借款申请
export function addLoanApply(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/loan/addLoanApply',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取报销列表
export function bxList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/bxlist',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取日常报销单详情
export function getDailyBxDetil(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/dailyBx/getDailyBxDetil',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取申请单成本中心冻结
export function getFreezingThawingHistory(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/finance/getFreezingThawingHistory',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

// 获取借款申请详情接口
export function getLoanDetail(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/loan/getLoanDetail',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

// 校验最大金额
export function checkMaxAmount(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/costControl/checkMaxAmount',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取默认收款人
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
// 获取借款单数组
export function getLoanList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/loan/getLoanList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 保存日常报销单
export function editSave(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/dailyBx/editSave',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 提交日常报销
export function sendDailyBx(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/dailyBx/sendDailyBx',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 删除未提交审核的报销单
export function deleteEdit(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/dailyBx/deleteEdit',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

// 差旅补助报销补助信息获取
export function getNeedTypeList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/travelBx/needTypeList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

// 差旅补助报销获取订单信息
export function getLeaderBxInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/travelBx/leaderBxInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

//差旅补助报销申请提交
export function submitApproval(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/travelBx/submitApproval',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}


//获取差旅补助报销详情
export function getTravelBxDetail(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/travelBx/travelBxDetail',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}



