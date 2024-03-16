import axios from "@/utils/request";

/**
 * 获取留言列表
 * @param data
 * @returns {AxiosPromise}
 */
export function apiGetMessageList(params?: object) {
  // return axios.get("/message/list", params);

  var a = {
    "content": "String",
    "headerColor": {
      "type": "String",
      "default": "#ff6c1a"
    },
    "nickname": {
      "type": "String",
      "default": "匿名网友"
    },
    "createTime": "String",
    "likes": {
      " type": "Number",
      "default": 0
    },
    "comments": {
      "type": "Number",
      "default": 0
    },
    "replyList": [
      {
        "replyHeaderColor": {
          "type": "String",
          "default": "#009688"
        },
        "replyContent": "String",
        "replyUser": {
          "type": "String",
          "default": "匿名网友"
        },
        "byReplyUser": "String",
        "replyTime": "String"
      }
    ]
  }

  console.log(params)
  return  new Promise(() => {
    const response = a;
  });

}

/**
 * 获取回复数量
 * @param data
 * @returns {AxiosPromise}
 */
export function apiGetReplyCount(params?: object) {
  // return axios.get("/message/replyCount", params);

  var a = {
    "content": "String",
    "headerColor": {
      "type": "String",
      "default": "#ff6c1a"
    },
    "nickname": {
      "type": "String",
      "default": "匿名网友"
    },
    "createTime": "String",
    "likes": {
      " type": "Number",
      "default": 0
    },
    "comments": {
      "type": "Number",
      "default": 0
    },
    "replyList": [
      {
        "replyHeaderColor": {
          "type": "String",
          "default": "#009688"
        },
        "replyContent": "String",
        "replyUser": {
          "type": "String",
          "default": "匿名网友"
        },
        "byReplyUser": "String",
        "replyTime": "String"
      }
    ]
  }

  console.log(params)
  return  new Promise(() => {
    const response = a;
  });

}

/**
 * 添加留言
 * @param data
 * @returns {AxiosPromise}
 */
export function apiAddMessage(params: object) {
  // return axios.post("/message/add", params);

  var a = {
    "content": "String",
    "headerColor": {
      "type": "String",
      "default": "#ff6c1a"
    },
    "nickname": {
      "type": "String",
      "default": "匿名网友"
    },
    "createTime": "String",
    "likes": {
      " type": "Number",
      "default": 0
    },
    "comments": {
      "type": "Number",
      "default": 0
    },
    "replyList": [
      {
        "replyHeaderColor": {
          "type": "String",
          "default": "#009688"
        },
        "replyContent": "String",
        "replyUser": {
          "type": "String",
          "default": "匿名网友"
        },
        "byReplyUser": "String",
        "replyTime": "String"
      }
    ]
  }

  console.log(params)
  return  new Promise(() => {
    const response = a;
  });

}

/**
 * 点赞
 * @param data
 * @returns {AxiosPromise}
 */
export function apiUpdateLikes(params: object) {
  // return axios.post("/message/updateLikes", params);

  var a = {
    "content": "String",
    "headerColor": {
      "type": "String",
      "default": "#ff6c1a"
    },
    "nickname": {
      "type": "String",
      "default": "匿名网友"
    },
    "createTime": "String",
    "likes": {
      " type": "Number",
      "default": 0
    },
    "comments": {
      "type": "Number",
      "default": 0
    },
    "replyList": [
      {
        "replyHeaderColor": {
          "type": "String",
          "default": "#009688"
        },
        "replyContent": "String",
        "replyUser": {
          "type": "String",
          "default": "匿名网友"
        },
        "byReplyUser": "String",
        "replyTime": "String"
      }
    ]
  }

  console.log(params)
  return  new Promise(() => {
    const response = a;
  });

}

/**
 * 回复
 * @param data
 * @returns {AxiosPromise}
 */
export function apiUpdateReplys(params: object) {
  // return axios.post("/message/updateReplys", params);

  var a = {
    "content": "String",
    "headerColor": {
      "type": "String",
      "default": "#ff6c1a"
    },
    "nickname": {
      "type": "String",
      "default": "匿名网友"
    },
    "createTime": "String",
    "likes": {
      " type": "Number",
      "default": 0
    },
    "comments": {
      "type": "Number",
      "default": 0
    },
    "replyList": [
      {
        "replyHeaderColor": {
          "type": "String",
          "default": "#009688"
        },
        "replyContent": "String",
        "replyUser": {
          "type": "String",
          "default": "匿名网友"
        },
        "byReplyUser": "String",
        "replyTime": "String"
      }
    ]
  }

  console.log(params)
  return  new Promise(() => {
    const response = a;
  });

}
