// src/api/base.ts
import axios from "axios";

const API_BASE_URL = "/api/";

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
});

// 请求拦截器
apiClient.interceptors.request.use(
  (config) => {
    // 在这里可以添加认证 token 等
    // const token = localStorage.getItem("authToken");
    // if (token) {
    //   config.headers["Authorization"] = `Bearer ${token}`;
    // }
    return config;
  },
  (error) => {
    console.error("API 发送失败:", error);
    return Promise.reject(error);
  }
);

// 响应拦截器
apiClient.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    console.error("API 请求失败:", error);
    return Promise.reject(error);
  }
);

export default apiClient;
