import axios from 'axios';
import { ElLoading } from 'element-plus'

var loading;
// Create an axios instance
const instance = axios.create({
  baseURL: '/api/',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
});

// Request interceptor
instance.interceptors.request.use(
  config => {
    // Do something before sending the request
    return config;
  },
  error => {
    // Handle the request error
    return Promise.reject(error);
  }
);

// Response interceptor
instance.interceptors.response.use(
    
  response => {
    if (loading != null){
        loading.close()
    }
    // Do something with the response data
    if (response.data.code != 200) {
        // Handle the response error
        return Promise.reject(response.data.err_msg);
    }
    return response.data.result
  },
  error => {
    // Handle the response error
    if (loading != null){
        loading.close()
    }
    return Promise.reject(error);
  }
);

/**
 * Common request method
 * @param {string} url Request URL
 * @param {string} method Request method, default is GET
 * @param {object} data Request data
 * @param {object} headers Request headers
 */
export function request(url, method = 'GET', data = {}, headers = {}) {
  loading = ElLoading.service({ fullscreen: true })
  return instance({
    url,
    method,
    data,
    headers
  });
}
