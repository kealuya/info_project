const { Controller } = require('egg');

class HomeController extends Controller {
  // 新增的方法
  async fundAccountQuery() {
    const { ctx, service } = this;

    ctx.body = await service.alipey.fundAccountQuery();

  }

  async fundTransUniTransfer() {
    const { ctx, service } = this;

    ctx.body = await service.alipey.fundTransUniTransfer();

  }
}


module.exports = HomeController;
