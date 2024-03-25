import axios from 'axios';
import {Toast} from "vant";
const instance = axios.create({
    // 请求超时时间
    timeout: 5000,
});
instance.interceptors.request.use(
    function (config) {
        return config;
    },
    function (error) {
        // 对请求错误做些什么
        return Promise.reject(error);
    },
);
instance.interceptors.response.use(
    function (response) {
        if (response.status === 200) {
            return response.data
        } else {
            Toast.fail('网络错误,请联系管理员');
        }
    },
    function (error) {
        // 关闭 loading
        Toast.fail(error.message);

        return Promise.reject(error);
    },
);
//定位服务
export function fetchLocation(){
    return new Promise(function (resolve, reject) {
        instance({
            url:'',
            method: 'get'
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
//天气服务https://restapi.amap.com/v3/weather/weatherInfo?city=110101&key=<用户key>
export function fetchWeather(cityCode: string){
    return new Promise(function (resolve, reject) {
        instance({
            url:``,
            method: 'get'
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取账本和审批数据
export function getApprCountTeacher(query :object){
    return new Promise(function (resolve, reject) {
        instance({
            url:'/ApprovalEcl/v1/getApprCountTeacher',
            data: query,
            method: 'post'
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
