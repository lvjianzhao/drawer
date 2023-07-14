import { request } from "@/utils/service"
import { GetTableRequestData } from "@/api/table/types/table"

export interface RowMeta {
  id: string
  name: string
  type: number
  businessName: string
  serviceStatus: number
  description: string
  serviceDate?: string[]
  serviceStartDate: string
  serviceEndDate: string
  maintenanceType: number
  /** vxe-table 自动添加上去的属性 */
  _VXE_ID?: string
}


export type CustomerResponseData = RowMeta & {
  updatedAt: string;
};


export type CustomerTableData = ApiResponseData<{
  list: CustomerResponseData[]
  total: number
}>

export interface customerResponse {
  createdAt: string
  ID: number
  name: string
}

export interface selectOption {
  label: string
  value: string
}


// 添加客户名称
export function addCustomerApi(data: RowMeta) {
  return request<ApiResponseData<null>>({
    url: "/customer",
    method: "post",
    data
  })
}


// 编辑客户名称
export function updateCustomerApi(data: RowMeta) {
  return request<ApiResponseData<customerResponse>>({
    url: "/customer",
    method: "put",
    data
  })
}

/** 删 */
export function deleteCustomerApi(id: string) {
  return request({
    url: `customer/${id}`,
    method: "delete"
  })
}

/** 查 */
export function getCustomerApi(params: GetTableRequestData) {
  return request<CustomerTableData>({
    url: "/customer",
    method: "get",
    params
  })
}

export function getAllCustomerApi() {
  return request<CustomerTableData>({
    url: "/customer/all",
    method: "get"
  })
}
