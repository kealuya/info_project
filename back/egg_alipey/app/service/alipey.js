const Service = require('egg').Service;
const { AlipaySdk } = require('alipay-sdk');
const fs = require('fs');
const alipaySdk = new AlipaySdk({
  appId: '2021004191630933',
  privateKey: fs.readFileSync('cert/应用私钥RSA2048-敏感数据，请妥善保管.txt', 'ascii'),
  // 传入支付宝根证书、支付宝公钥证书和应用公钥证书。
  alipayRootCertPath: 'cert/alipayRootCert.crt',
  alipayPublicCertPath: 'cert/alipayCertPublicKey_RSA2.crt',
  appCertPath: 'cert/appCertPublicKey_2021004191630933.crt',
});

class AlipeyService extends Service {
  /**
   * 支付宝资金账户资产查询接口 alipay.fund.account.query
   * @param null
   *
   * @return
   * @property {string} code - 接口响应码，10000表示成功
   * @property {string} msg - 接口响应描述
   * @property {string} sub_code - 业务返回码
   * @property {string} sub_msg - 业务返回码描述
   * @property {string} available_amount - 可用余额
   * @property {string} freeze_amount - 冻结金额
   * @property {string} total_amount - 总金额
   * @property {Object} amount_detail - 金额明細
   * @property {string} amount_detail.acs - acs金额
   * @property {string} amount_detail.bank - 银行金额
   */
  async fundAccountQuery() {

    const result = await alipaySdk.exec('alipay.fund.account.query', {
      bizContent: {
        account_type: 'ACCTRANS_ACCOUNT',
      },
    });
    if ( result.code!=="10000"){
      this.logger.error(result)
    }else{
      // do something
      this.logger.info(result)
    }

    return result
  }

  async find(uid) {
    const user = await this.ctx.db.query(
      'select * from user where uid = ?',
      uid
    );
    return user;
  }
}

module.exports = AlipeyService;
