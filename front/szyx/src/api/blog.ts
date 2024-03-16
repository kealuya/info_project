import axios from "@/utils/request";
/**
 * 获取博客列表
 * @param data
 * @returns {AxiosPromise}
 */
export function apiGetBlogList(params?: object) {
  //return axios.get("/blog/list", params);

  var  a = [{
    type: Array,
    "title": "氨酚黄那敏胶囊",
    "desc": "氨酚咖那敏片，适应症为适用于缓解普通感冒及流行性感冒引起的发热、头痛、四肢酸痛、打喷嚏、流鼻涕、鼻塞、咽痛等症状。",
    "fileCoverImgUrl": "https://img0.baidu.com/it/u=2152382963,700737397&fm=253&fmt=auto&app=138&f=JPEG?w=688&h=500",
    "html": "String",
    "markdown": "String",
    "level": 1,
    "github": "String",
    "auth": "张文宏",
    "source": "Number",
    "isVisible": false,
    "releaseTime": "2024-03-05",
  //   "pv": {
  //   type: Number, default: 0
  // },
    "pv" :10,
  //   "likes": {
  //   type: Number, default: 0
  // },
    "likes":100,
  //   "comments": {
  //   type: Number, default: 0
  // }
    "comments":200
  },
    {
      type: Array,
      "title": "氨酚黄那敏胶囊",
      "desc": "氨酚咖那敏片，适应症为适用于缓解普通感冒及流行性感冒引起的发热、头痛、四肢酸痛、打喷嚏、流鼻涕、鼻塞、咽痛等症状。",
      "fileCoverImgUrl": "https://img0.baidu.com/it/u=2152382963,700737397&fm=253&fmt=auto&app=138&f=JPEG?w=688&h=500",
      "html": "String",
      "markdown": "String",
      "level": 1,
      "github": "String",
      "auth": "张文宏",
      "source": "Number",
      "isVisible": false,
      "releaseTime": "2024-03-05",
      //   "pv": {
      //   type: Number, default: 0
      // },
      "pv" :10,
      //   "likes": {
      //   type: Number, default: 0
      // },
      "likes":100,
      //   "comments": {
      //   type: Number, default: 0
      // }
      "comments":200
    },
    {
      type: Array,
      "title": "氨酚黄那敏胶囊",
      "desc": "氨酚咖那敏片，适应症为适用于缓解普通感冒及流行性感冒引起的发热、头痛、四肢酸痛、打喷嚏、流鼻涕、鼻塞、咽痛等症状。",
      "fileCoverImgUrl": "https://img0.baidu.com/it/u=2152382963,700737397&fm=253&fmt=auto&app=138&f=JPEG?w=688&h=500",
      "html": "String",
      "markdown": "String",
      "level": 1,
      "github": "String",
      "auth": "张文宏",
      "source": "Number",
      "isVisible": false,
      "releaseTime": "2024-03-05",
      //   "pv": {
      //   type: Number, default: 0
      // },
      "pv" :10,
      //   "likes": {
      //   type: Number, default: 0
      // },
      "likes":100,
      //   "comments": {
      //   type: Number, default: 0
      // }
      "comments":200
    },
    {
      type: Array,
      "title": "氨酚黄那敏胶囊",
      "desc": "氨酚咖那敏片，适应症为适用于缓解普通感冒及流行性感冒引起的发热、头痛、四肢酸痛、打喷嚏、流鼻涕、鼻塞、咽痛等症状。",
      "fileCoverImgUrl": "https://img0.baidu.com/it/u=2152382963,700737397&fm=253&fmt=auto&app=138&f=JPEG?w=688&h=500",
      "html": "String",
      "markdown": "String",
      "level": 1,
      "github": "String",
      "auth": "张文宏",
      "source": "Number",
      "isVisible": false,
      "releaseTime": "2024-03-05",
      //   "pv": {
      //   type: Number, default: 0
      // },
      "pv" :10,
      //   "likes": {
      //   type: Number, default: 0
      // },
      "likes":100,
      //   "comments": {
      //   type: Number, default: 0
      // }
      "comments":200
    },
    {
      type: Array,
      "title": "氨酚黄那敏胶囊",
      "desc": "氨酚咖那敏片，适应症为适用于缓解普通感冒及流行性感冒引起的发热、头痛、四肢酸痛、打喷嚏、流鼻涕、鼻塞、咽痛等症状。",
      "fileCoverImgUrl": "https://img0.baidu.com/it/u=2152382963,700737397&fm=253&fmt=auto&app=138&f=JPEG?w=688&h=500",
      "html": "String",
      "markdown": "String",
      "level": 1,
      "github": "String",
      "auth": "张文宏",
      "source": "Number",
      "isVisible": false,
      "releaseTime": "2024-03-05",
      //   "pv": {
      //   type: Number, default: 0
      // },
      "pv" :10,
      //   "likes": {
      //   type: Number, default: 0
      // },
      "likes":100,
      //   "comments": {
      //   type: Number, default: 0
      // }
      "comments":200
    },
    {
      type: Array,
      "title": "氨酚黄那敏胶囊",
      "desc": "氨酚咖那敏片，适应症为适用于缓解普通感冒及流行性感冒引起的发热、头痛、四肢酸痛、打喷嚏、流鼻涕、鼻塞、咽痛等症状。",
      "fileCoverImgUrl": "https://img0.baidu.com/it/u=2152382963,700737397&fm=253&fmt=auto&app=138&f=JPEG?w=688&h=500",
      "html": "String",
      "markdown": "String",
      "level": 1,
      "github": "String",
      "auth": "张文宏",
      "source": "Number",
      "isVisible": false,
      "releaseTime": "2024-03-05",
      //   "pv": {
      //   type: Number, default: 0
      // },
      "pv" :10,
      //   "likes": {
      //   type: Number, default: 0
      // },
      "likes":100,
      //   "comments": {
      //   type: Number, default: 0
      // }
      "comments":200
    }
 ]

  console.log(params)
  const response = {
    data :{
      list : a,
      total :6
    } ,
    success :true
  }


  return  new Promise((resolve) => {
    resolve(response)

  });

}
/**
 * 获取博客详情
 * @param data
 * @returns {AxiosPromise}
 */
export function apiGetBlogDetail(params?: object) {
  // return axios.get("/blog/info", params);

  var  a = {
    "type": "Array",
    "title": "协和医院",
    "desc": "2",
    "fileCoverImgUrl": "https://bkimg.cdn.bcebos.com/pic/d01373f082025aafa40f428635babc64034f78f0ac7d?x-bce-process=image/resize,m_lfit,w_536,limit_1/quality,Q_70",
    "html": "String",
    "markdown": "String",
    "level": 1,
    "github": "String",
    "auth": "张文宏",
    "source": "Number",
    "isVisible": false,
    "releaseTime": "2024-03-05",
    "pv": {
      "type": "Number",
      "default": 0
    },
    "likes": {
      "type": "Number",
      "default": 0
    },
    "comments": {
      "type": "Number",
      "default": 0
    }
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
  // return axios.post("/blog/updateLikes", params);

  var  a = {
    "type": "Array",
    "title": "协和医院",
    "desc": "2",
    "fileCoverImgUrl": "https://bkimg.cdn.bcebos.com/pic/d01373f082025aafa40f428635babc64034f78f0ac7d?x-bce-process=image/resize,m_lfit,w_536,limit_1/quality,Q_70",
    "html": "String",
    "markdown": "String",
    "level": 1,
    "github": "String",
    "auth": "张文宏",
    "source": "Number",
    "isVisible": false,
    "releaseTime": "2024-03-05",
    "pv": {
      "type": "Number",
      "default": 0
    },
    "likes": {
      "type": "Number",
      "default": 0
    },
    "comments": {
      "type": "Number",
      "default": 0
    }
  }

  console.log(params)
  return  new Promise(() => {
    const response = a;
  });

}
/**
 * 浏览量
 * @param data
 * @returns {AxiosPromise}
 */
export function apiUpdatePV(params: object) {
  // return axios.post("/blog/updatePV", params);

  var  a = {
    "type": "Array",
    "title": "协和医院",
    "desc": "2",
    "fileCoverImgUrl": "https://bkimg.cdn.bcebos.com/pic/d01373f082025aafa40f428635babc64034f78f0ac7d?x-bce-process=image/resize,m_lfit,w_536,limit_1/quality,Q_70",
    "html": "String",
    "markdown": "String",
    "level": 1,
    "github": "String",
    "auth": "张文宏",
    "source": "Number",
    "isVisible": false,
    "releaseTime": "2024-03-05",
    "pv": {
      "type": "Number",
      "default": 0
    },
    "likes": {
      "type": "Number",
      "default": 0
    },
    "comments": {
      "type": "Number",
      "default": 0
    }
  }

  console.log(params)
  return  new Promise(() => {
    const response = a;
  });
}
