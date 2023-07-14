// 添加客户名称
import { request } from "@/utils/service"
import * as Table from "@/api/table/types/table"

export interface reqCustomerEnv {
  id: string
  customerId: 0
  env: string
  vpn: string
  deploy_report: string
  description: string
}
export interface RowMeta {
  id: string
  name: string
  customerId: number
  env: string
  vpn: string
  deploy_report: string
  license: string
  description: string
  /** vxe-table 自动添加上去的属性 */
  _VXE_ID?: string
}

export type GetEnvResponseData = ApiResponseData<{
  list: RowMeta[]
  total: number
}>

/** 新增 */
export function addCustomerEnvApi(data: reqCustomerEnv) {
  return request<ApiResponseData<null>>({
    url: "/customer/env",
    method: "post",
    data
  })
}

/** 修改 */
export function updateCustomerEnvApi(data: reqCustomerEnv) {
  return request<ApiResponseData<null>>({
    url: "/customer/env",
    method: "put",
    data
  })
}

/** 查 */
export function getCustomerEnvApi(params: Table.GetTableRequestData) {
  return request<GetEnvResponseData>({
    url: "/customer/env",
    method: "get",
    params
  })
}

/** 删 */
export function deleteCustomerEnvApi(id: string) {
  return request({
    url: `customer/env/${id}`,
    method: "delete"
  })
}
