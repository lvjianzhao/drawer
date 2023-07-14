import { getAllCustomerApi } from "@/api/customer"
import { getSummaryApi } from "@/api/customer/summary"

export const getAllCustomer = async () => {
  try {
    const res = await getAllCustomerApi()
    const customer = res.data.list.map((item) => {
      return {
        value: item.id,
        label: String(item.name)
      }
    })
    return customer
  } catch (error) {
    console.log(error)
  }
}
