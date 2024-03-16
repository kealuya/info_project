import axios from "@/utils/request";

/**
 * 获取标签列表
 * @param data
 * @returns {AxiosPromise}
 */
export function apiGetLabelList(params?: object) {
  // return axios.get("/label/list", params);


  var a = {
    " label": "String",
    "bgColor": "String",
    "createTime": {
      " type": "Date",
      "default": "Date.now"
    }
  }
  console.log(params)
  return  new Promise(() => {
    const data = a;
  });
}
