import { type VxeColumnPropTypes } from "vxe-table/types/column"

const solts: VxeColumnPropTypes.Slots = {
  default: ({ row, column }) => {
    const cellValue = row[column.field]
    let value = "POC"
    if (cellValue === 2) {
      value = "正式"
    }
    // return [<span class={`el-tag el-tag--plain`}>{value}</span>]
    return value
  }
}

export default solts
