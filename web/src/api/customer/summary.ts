import { request } from "@/utils/service"
type GetSummaryResponseData = ApiResponseData<{
  list: Object
  total: number
}>

export function getSummaryApi() {
  return request<GetSummaryResponseData>({
    url: "/customer/summary",
    method: "get"
  })
}
