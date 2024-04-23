import axios, { AxiosInstance, AxiosRequestConfig, Method } from "axios";

const service = {
  //   baseURL: import.meta.env.VITE_BASE_API,
  baseURL:
    // "https://www.fastmock.site/mock/d81efab1ab641aa5bd8561c24f6c8bb3/test",
   "http://localhost:5173",
  timeout: 5000,
  // 跨域时候允许携带凭证
  withCredentials: true,
};

class RequestHttp {
  constructor() {
    this.httpInterceptorsRequest();
    this.httpInterceptorsResponse();
  }

  // 实例化axios
  private static axiosInstance: AxiosInstance = axios.create(service);

  // 处理请求拦截器
  private httpInterceptorsRequest(): void {
    RequestHttp.axiosInstance.interceptors.request.use(
      (config) => {
        // const token = getToken();
        const token = sessionStorage.getItem("token");
        if (token) {
          config.headers.Authorization = token;
        }
        return config;
      },
      (error) => {
        return Promise.reject(error);
      }
    );
  }

  // 处理响应拦截器
  private httpInterceptorsResponse(): void {
    RequestHttp.axiosInstance.interceptors.response.use(
      (response) => {
        return response.data;
      },
      (e) => {
        return Promise.reject(e.response.data.message);
      }
    );
  }
}

export const http = new RequestHttp();
