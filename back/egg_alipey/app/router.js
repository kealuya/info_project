/**
 * @param {Egg.Application} app - egg application
 */
module.exports = app => {
  const { router, controller } = app;
  router.get('/', controller.home.fundAccountQuery);
  router.get('/pay', controller.home.fundTransUniTransfer);
};
