import Axios from "axios";
import {
    MessageBox,
    Message
} from "element-ui";
// 默认值
Axios.defaults.baseURL = "http://127.0.0.1:8808" // 环境地址
// Axios.defaults.baseURL = process.env.NODE_ENV === "production" ? "/" : "http://127.0.0.1:8808" // 环境地址

Axios.defaults.timeout = 3000 // 超时
Axios.defaults.withCredentials = true // 带cookies

/**
 * @description: 请求前拦截
 * @param {*} function
 * @return {*}
 */
Axios.interceptors.request.use((config) => {
    console.log(config.baseURL);
    // post 请求时，防止data = undefined
    if (config.method.toUpperCase() === 'POST' && !isEmptyObject(config.data)) {
        config.data = {}
    }
    return config
}, (err) => {
    return Promise.reject(err)
})

/**
 * @description: 相应拦截
 * @param {*}
 * @return {*}
 */
Axios.interceptors.response.use((res) => {
    const data = res.data;
    if (data.code !== 200) {
        // 用户未登录
        if (data.code === 1008) {
            // 确定未登录重新登录
            MessageBox.confirm(
                "登录已超时，请重新登录！",
                "提示", {
                    confirmButtonText: "重新登录",
                    type: "warning",
                }
            ).then(() => {
                // store.dispatch("user/resetToken").then(() => {
                //   location.reload();
                // });
            });
        } else {
            // 提示错误
            Message({
                message: data.message || "Error",
                type: "error",
                duration: 2 * 1000,
            });
        }
    }
    return data;
}, (err) => {
    return Promise.reject(err)
})

/**
 * @description: 导出post方法
 * @param {*} url 路径
 * @param {*} data 参数
 * @param {*} cb 回调函数
 * @return {*}
 */
export function POST(url, data, cb) {
    // 是回调函数才回调
    var b = typeof cb === "function"
    Axios.post(url, data).then(function (res) {
        if (b) {
            cb(res)
        }
    }).catch(function (err) {
        if (b) {
            cb(err)
        }
    })
}

/**
 * @description: 导出get方法
 * @param {*} url
 * @param {*} data
 * @param {*} cb
 * @return {*}
 */
export function GET(url, data, cb) {
    // 是回调函数才回调
    var b = typeof cb === "function"
    Axios.get(url, {
        params: data
    }).then((res) => {
        if (b) {
            cb(res)
        }
    }).catch((err) => {
        if (b) {
            cb(err)
        }
    })
}
/**
 * @description: 是否是对象
 * @param {*} params
 * @return {*}
 */
function isEmptyObject(params) {
    return Object.prototype.toString.call(params) === Object.prototype.toString.call({})
}