/**
 * 获取支付宝错误码描述信息
 * @param {string} code 错误码
 * @returns {{code: string, message: string, solution: string}} 错误信息对象
 */
function getAlipayErrorMessage(code) {
  const errorMap = {
    'SYSTEM_ERROR': {
      message: '系统繁忙',
      solution: '服务器异常 可能发生了网络或者系统异常，导致服务调用失败，商户可以用同样的请求发起重试'
    },
    'INVALID_PARAMETER': {
      message: '参数有误',
      solution: '请根据接口返回的参数非法的具体错误信息，修改参数后进行重试'
    },
    'ALIPAY_ACCOUNT_NOT_EXIST': {
      message: '对方没有支付宝账户',
      solution: '对方没有支付宝账户，可使用转账到卡服务给对方银行卡转账'
    },
    'AUTHOREE_IS_NOT_MATCH': {
      message: '请求的partner_id与支付授权协议的被授权方不一致',
      solution: '请检查授权协议的签约的partner_id是否与本次请求的partner_id一致'
    },
    'AUTH_CODE_ERROR': {
      message: '授权信息不存在',
      solution: '授权信息不存在，请检查协议号是否正确'
    },
    'AUTH_INFO_NOT_EXISTS': {
      message: '授权信息不存在',
      solution: '授权信息不存在，请检查协议号是否正确'
    },
    'AUTH_USERID_IS_NOT_MATCH': {
      message: '支付授权协议的授权方与查询的alipay_user_id不匹配',
      solution: '请检查支付授权协议授权方与alipay_user_id的一致性'
    },
    'MERCHANT_AGREEMENT_VERIFY_FAIL': {
      message: '支付授权协议对应的partner_id的商户协议校验失败',
      solution: '请检查请求的partner_id是否签约支付授权签约产品'
    },
    'NO_PERMISSION': {
      message: '没有接口使用权限',
      solution: '确认是否有查询余额权限'
    },
    'NO_PERMISSION_ACCOUNT': {
      message: '无权限操作当前付款账号',
      solution: '无权限操作当前付款账号'
    },
    'TRUSTEESHIP_ACCOUNT_NOT_EXIST': {
      message: '托管子户查询不存在',
      solution: '托管子户查询不存在'
    },
    'USER_AGREEMENT_IS_OUT_OF_DATE': {
      message: '用户协议已过期',
      solution: '请联系授权方重新进行支付授权协议的签约，并更新agreement_no'
    },
    'USER_AGREEMENT_STATUS_ABNORMAL': {
      message: '用户协议状态不正常',
      solution: '用户协议状态不正常，请检查协议有效性'
    },
    'USER_AGREEMENT_VERIFY_FAIL': {
      message: '用户协议校验失败',
      solution: '请检查用户协议号是否正确'
    }
  };

  // 如果找不到对应的错误码,返回未知错误
  if (!errorMap[code]) {
    return {
      message: '未知错误',
      solution: '请联系技术支持'
    };
  }

  // 返回错误信息对象
  return {
    message: errorMap[code].message,
    solution: errorMap[code].solution
  };
}


// 使用示例:
// const error = getAlipayErrorMessage('SYSTEM_ERROR');
// console.log(error);
// {
//   code: 'SYSTEM_ERROR',
//   message: '系统繁忙',
//   solution: '服务器异常 可能发生了网络或者系统异常，导致服务调用失败，商户可以用同样的请求发起重试'
// }
