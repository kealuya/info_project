import request from "./request";
// 帮助反馈
export function addFeedBack(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/feedback/addFeedBack',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}


