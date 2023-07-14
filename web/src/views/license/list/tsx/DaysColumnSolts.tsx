import { type VxeColumnPropTypes } from "vxe-table/types/column"

const solts: VxeColumnPropTypes.Slots = {
  default: ({ row, column }) => {
    // 给定日期
    const givenDate: Date = new Date(row[column.field])
    // 当前日期
    const currentDate: Date = new Date()
    // 计算时间间隔（以天为单位）
    const timeDiff = Math.ceil((givenDate.getTime() - currentDate.getTime()) / (1000 * 60 * 60 * 24))
    let type = "danger"
    if (timeDiff > 31) {
      type = ""
    } else if (timeDiff > 0) {
      type = "warning"
    }
    return [<el-tag type={`${type}`}>{timeDiff}</el-tag>]
  }
}

export default solts
