const Service = require('egg').Service;
const {AlipaySdk} = require('alipay-sdk');
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
     *
     * @return {Object} result
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
        if (result.code !== "10000") {
            this.logger.error(result)
        } else {
            // do something
            this.logger.info(result)
        }

        return result
    }

    /**
     * 支付宝转账接口 alipay.fund.trans.uni.transfer
     * @param {Object} params - 转账参数对象
     * @param {string} params.outBizNo - 商户转账唯一订单号，用于做幂等控制
     * @param {string} params.transAmount - 转账金额，单位：元，精确到小数点后两位
     * @param {string} params.orderTitle - 转账业务标题，会展示在支付宝账单中
     * @param {string} params.payeeInfoIdentity - 收款方支付宝账号
     * @param {string} params.payeeInfoName - 收款方真实姓名/企业名称
     * @param {string} params.remark - 转账备注，业务备注信息
     *
     * @return {Object} result
     * @property {string} code - 接口响应码，10000表示成功
     * @property {string} msg - 接口响应描述
     * @property {string} sub_code - 业务返回码
     * @property {string} sub_msg - 业务返回码描述
     * @property {string} out_biz_no - 商户转账唯一订单号
     * @property {string} order_id - 支付宝转账订单号
     * @property {string} pay_fund_order_id - 支付宝支付资金流水号
     * @property {string} status - 转账单据状态。SUCCESS：成功；FAIL：失败；DEALING：处理中
     * @property {string} trans_date - 订单支付时间，格式为yyyy-MM-dd HH:mm:ss
     *
     * @example
     * const result = await fundTransUniTransfer({
     *   outBizNo: "TRX202401010001",
     *   transAmount: "100.00",
     *   orderTitle: "员工报销款",
     *   payeeInfoIdentity: "user@example.com",
     *   payeeInfoName: "张三",
     *   remark: "1月份差旅费报销"
     * });
     */
    async fundTransUniTransfer(params) {
        /*
            bizContent: {
                out_biz_no: "JD0000000000002",
                trans_amount: "0.1",
                biz_scene: "DIRECT_TRANSFER",
                product_code: "TRANS_ACCOUNT_NO_PWD",
                // 【转账备注】转账业务的标题，用于在支付宝用户的账单里显示
                order_title: "测试支付宝转账",
                payee_info: {
                    identity: "1136700510@qq.com",
                    identity_type: "ALIPAY_LOGON_ID",
                    name:"苏州境步科技有限公司",
                    // bankcard_ext_info: {
                    //   inst_name: "招商银行",
                    //   account_type: "1",
                    // },
                },
                // 【理由】业务备注
                remark: "测试支付宝转账remark",
                business_params: {payer_show_name_use_alias: true}
            },
         */
        const result = await alipaySdk.exec('alipay.fund.trans.uni.transfer', {
            bizContent: {
                out_biz_no: params.outBizNo,
                trans_amount: params.transAmount,
                biz_scene: "DIRECT_TRANSFER",
                product_code: "TRANS_ACCOUNT_NO_PWD",
                // 【转账备注】转账业务的标题，用于在支付宝用户的账单里显示
                order_title: params.orderTitle,
                payee_info: {
                    identity: params.payeeInfoIdentity,
                    identity_type: "ALIPAY_LOGON_ID",
                    name: params.payeeInfoName,
                    // bankcard_ext_info: {
                    //   inst_name: "招商银行",
                    //   account_type: "1",
                    // },
                },
                // 【理由】业务备注
                remark: params.remark,
                business_params: {payer_show_name_use_alias: true}
            },
        });
        if (result.code !== "10000") {
            this.logger.error(result)
        } else {
            // do something
            this.logger.info(result)
        }

        return result
    }
}

module.exports = AlipeyService;
