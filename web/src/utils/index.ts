import dayjs from "dayjs"
import utc from "dayjs/plugin/utc"
import timezone from "dayjs/plugin/timezone"

// 引入插件
dayjs.extend(utc)
dayjs.extend(timezone)

/** 格式化日期 */
export const formatDate = (time: string | number | Date) => {
  return time ? dayjs.utc(time).local().format("YYYY-MM-DD") : "N/A"
}

/** 格式化日期时间 */
export const formatDateTime = (time: string | number | Date) => {
  return time ? dayjs.utc(time).local().format("YYYY-MM-DD HH:mm:ss") : "N/A"
}

/** 格式化时间戳 */
export const formatTimestamp = (time: string | number | Date) => {
  return time ? new Date(time).getTime() / 1000 : 0
}

/** 用 JS 获取全局 css 变量 */
export const getCssVariableValue = (cssVariableName: string) => {
  let cssVariableValue = ""
  try {
    // 没有拿到值时，会返回空串
    cssVariableValue = getComputedStyle(document.documentElement).getPropertyValue(cssVariableName)
  } catch (error) {
    console.error(error)
  }
  return cssVariableValue
}

/** 用 JS 设置全局 CSS 变量 */
export const setCssVariableValue = (cssVariableName: string, cssVariableValue: string) => {
  try {
    document.documentElement.style.setProperty(cssVariableName, cssVariableValue)
  } catch (error) {
    console.error(error)
  }
}
