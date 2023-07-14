export interface CreateTableRequestData {
  username: string
  password: string
}

export interface UpdateTableRequestData {
  id: string
  username: string
  password?: string
}

export interface GetTableRequestData {
  /** 当前页码 */
  currentPage?: number
  /** 查询条数 */
  size?: number
  /** 查询参数：用户名 */
  username?: string
  /** 查询参数：维保状态 */
  serviceStatus?: number
  /** 部署报告中的IP */
  ip?: string
  /** license状态 */
  active?: number
  /** customerId*/
  customerId?: string | number
  /** 维保类型*/
  maintenanceType?: number
}

export interface GetTableData {
  createTime: string
  email: string
  id: string
  phone: string
  roles: string
  status: boolean
  username: string
  name: string
  createdAt: string
  ID: number
  env: string
}

export type GetTableResponseData = ApiResponseData<{
  list: GetTableData[]
  total: number
}>
