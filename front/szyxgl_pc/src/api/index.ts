import api from '../utils/request'

/**
 * 测试接口
 */

// 测试用Hello World
export const login = () => api({ url: '/api/login', method: 'get' })

/**
 * 动态路由接口
 */
export const GetDynamicRoutes = () => api({ url: '/api/routes', method: 'get' })
// export const GetDynamicRoutes = () => api({ url: '/api/routes', method: 'post' })


