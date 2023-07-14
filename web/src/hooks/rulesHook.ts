/**
 * @returns el-form组件创建校验规则
 */
export function rulesHook() {
  /**
   * @param {*} data 配置额外组件
   */
  const InitRules = (data: any = []) => {
    if (data.length === 0) {
      return false
    }
    let formRules: Record<string, any[]> = {}
    // 判断是否有required属性
    // @ts-ignore
    data.forEach((item) => {
      //初始化规则数组
      let rulesArr = []
      if (item.required) {
        let json = { required: true, message: item.message || messageType(item), trigger: "change" }
        rulesArr.push(json)
      }
      // 是否有其他校验规则
      const rule = item.rule
      if (rule && Array.isArray(rule) && rule.length > 0) {
        rulesArr = rulesArr.concat(rule)
      }
      formRules[item.prop] = rulesArr
    })
    return formRules
  }
  /**
   *
   * @description 提示文本
   */
  // @ts-ignore
  const messageType = (data) => {
    let msg = ""
    switch (data.type) {
      case "input":
        msg = "请输入"
        break
      case "upload":
        msg = "请上传"
        break
      case "radio":
      case "checkbox":
      case "select":
      case "date":
        msg = "请选择"
        break
      default:
        msg = "未定义"
    }
    return `${msg}${data.label}`
  }
  return { InitRules }
}
