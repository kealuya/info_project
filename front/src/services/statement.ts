import {request} from "./request";
//获取个人考勤统计
export function countWorkOvertimeOrAskForLeave(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/statement/countWorkOvertimeOrAskForLeave',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取差旅费用统计和出行次数
export function countCorporateTravelSpending(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/statement/countCorporateTravelSpending',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取企业费用折线图
export function countFirmAccountingStatisticsData(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/statement/countFirmAccountingStatisticsData',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 根据时间范围、OA申请类别,统计个人OA办公提交的单数
export function countIndividualNumber(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/statement/countIndividualNumber',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取个人差旅消费统计
export function countTravelSpendingStatistics(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/statement/countMonthAccountingStatisticsData',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取个人出行次数消费统计
export function countTravelSpendingStatistic(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/statement/countTravelSpendingStatistics',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
