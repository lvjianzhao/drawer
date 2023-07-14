import {request} from "@/utils/service"
import axios from 'axios';

export interface LoginRequestData {
  /** admin 或 editor */
  username: string
  /** 密码 */
  password: string
  /** 验证码 */
  captcha: string
  captchaId: string
}

export interface upload {
  image?: any
  file?: string
  type?: number
}

type LoginCodeResponseData = ApiResponseData<{ picPath: string; captchaId: string }>
type LoginResponseData = ApiResponseData<{ token: string }>

// 获取验证码
export const captcha = () => {
  return request<LoginCodeResponseData>({
    url: "/base/captcha",
    method: "post"
  })
}

/** 登录并返回 Token */
export function loginApi(data: LoginRequestData) {
  return request<LoginResponseData>({
    url: "/base/login",
    method: "post",
    data
  })
}

// 上传文件
export function uploadFile(data = {}) {
  return request<LoginCodeResponseData>({
    url: "/base/uploadFile",
    method: "post",
    data,
    headers: {
      "Content-Type": "multipart/form-data"
    }
  })
}
