// 添加客户名称
import { request } from "@/utils/service"

export interface RowMeta extends GetLicenseData {
  /** vxe-table 自动添加上去的属性 */
  _VXE_ID?: string
}

export interface GetLicenseData {
  id: number
  customerName: string
  env: string
  expiredDate: string
  license: string
  createdAt: string
  active?: number
}

export type GetLicenseResponseData = ApiResponseData<{
  list: GetLicenseData[]
  total: number
}>
export interface reqLicense {
  customerID: number | string
  env: string
  compDisplayName: string
  expiredDate: string
  managerApis: number
  managerUsers: number
  managerProjects: number
  managerIntegrations: number
  managerOrganizations: number
  managerTests: number
  managerProducts: number
  studioUsers: number
  studioCompBlackList: string
  permissionMenu?: string
  checkList?: string
  active?: number
}

export interface LicenseRequestData {
  /** 当前页码 */
  currentPage?: number
  /** 查询条数 */
  size?: number
  /** 查询参数：用户名 */
  name?: string
  /** license状态 */
  active?: number
}

/** 新增 */
export function addCustomerEnvApi(data: reqLicense) {
  return request<ApiResponseData<null>>({
    url: "/license",
    method: "post",
    data
  })
}

/** 查 */
export function getLicenseApi(params: LicenseRequestData) {
  return request<GetLicenseResponseData>({
    url: "/license",
    method: "get",
    params
  })
}

/** 删 */
export function deleteLicenseApi(id: number) {
  return request({
    url: `license/${id}`,
    method: "delete"
  })
}
