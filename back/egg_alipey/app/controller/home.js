const { Controller } = require('egg');
const { AlipaySdk } = require('alipay-sdk');
const fs = require('fs');

class HomeController extends Controller {
  // 新增的方法
  async fundAccountQuery() {
    const { ctx, service } = this;

    ctx.body = await service.alipey.fundAccountQuery();

  }

}


module.exports = HomeController;
