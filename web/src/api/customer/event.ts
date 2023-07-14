// 编辑客户事件
import { request } from "@/utils/service"
import { customerResponse } from "./index"

export type RowMeta = {
  id: number;
  customerID: number;
  name?: string;
  eventContent: string;
  eventTime: string;
  /** vxe-table 自动添加上去的属性 */
  _VXE_ID?: string;
};

export type EventResponseData = RowMeta & {
  updatedAt: string;
};


export type EventTableData = ApiResponseData<{
  list: EventResponseData[]
  total: number
}>

interface searchParams {
  /** 当前页码 */
  currentPage: number
  /** 查询条数 */
  size: number
  customerName?: string
  customerID?: number | string
  startTime?: number
  endTime?: number
  //* 排序字段 */
  sort?: string
}

// 添加客户事件
export function addEventApi(data: RowMeta) {
  return request<ApiResponseData<null>>({
    url: "/customer/event",
    method: "post",
    data
  })
}

// 编辑客户事件
export function updateEventApi(data: RowMeta) {
  return request<ApiResponseData<customerResponse>>({
    url: "/customer/event",
    method: "put",
    data
  })
}

// 获取客户事件
export function getCustomerEventApi(params: searchParams) {
  return request<EventTableData>({
    url: "/customer/event",
    method: "get",
    params
  })
}

/** 删 */
export function deleteCustomerEventApi(id: number) {
  return request({
    url: `customer/event/${id}`,
    method: "delete"
  })
}
