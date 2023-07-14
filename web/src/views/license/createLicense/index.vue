<template>
  <div class="app-container">
    <el-form :model="formData" class="demo-form-inline" ref="formDom" :rules="rules" v-loading="loading">
      <el-space wrap fill style="width: 100%" size="large">
        <template v-for="(form_item, index) in form_config_item" :key="form_item.key">
          <el-card class="box-card">
            <warning-bar
              v-if="index === 0"
              title="需要先选择客户，才能加载出来环境列表。创建license前请先创建客户及其环境信息"
            />
            <div class="box-row">
              <el-row :gutter="50">
                <template v-for="item in form_item[form_item.key]" :key="item.prop">
                  <el-col :span="item.col || 24">
                    <!--                  select选择器-->
                    <el-form-item
                      v-if="item.type === 'select'"
                      :label="item.label"
                      :prop="item.prop"
                      :rules="item.rules"
                    >
                      <el-select
                        v-model="formData[item.prop]"
                        :style="`width: ${item.width}`"
                        placeholder="请选择"
                        clearable
                        filterable
                      >
                        <el-option
                          v-for="select in item.options"
                          :key="select.value"
                          :label="select.label"
                          :value="select.value"
                        ></el-option>
                      </el-select>
                    </el-form-item>

                    <!--                  时间日期选择器-->
                    <el-form-item v-if="item.type === 'date'" :label="item.label" :prop="item.prop" :rules="item.rules">
                      <el-date-picker
                        v-model="formData[item.prop]"
                        :style="`width: ${item.width}`"
                        :type="item.date_type || 'datetime'"
                        :format="item.date_format || 'YYYY-MM-DD HH:mm:ss'"
                        :value-format="item.date_value || 'YYYY-MM-DD HH:mm:ss'"
                        :placeholder="item.placeholder || '选择日期时间'"
                      ></el-date-picker>
                    </el-form-item>

                    <!--                  input输入框 -->
                    <el-form-item
                      v-if="item.type === 'input'"
                      :label="item.label"
                      :prop="item.prop"
                      :rules="item.rules"
                    >
                      <el-input
                        :maxlength="item.max_length"
                        :minlength="item.min_length"
                        :style="`width: ${item.width}`"
                        :placeholder="item.placeholder"
                        v-model="formData[item.prop]"
                        clearable
                      ></el-input>
                    </el-form-item>

                    <!-- inputNumber -->
                    <el-form-item
                      v-if="item.type === 'inputNumber'"
                      :label="item.label"
                      :prop="item.prop"
                      :rules="item.rules"
                    >
                      <el-input
                        :min="item.min || 0"
                        :max="item.max || 999999999"
                        v-model.number="formData[item.prop]"
                        clearable
                      ></el-input>
                    </el-form-item>
                  </el-col>
                </template>
              </el-row>
            </div>
          </el-card>
        </template>

        <!--        权限设置-->
        <Permission @submit-permission="handlerSubmitForm"/>
      </el-space>
    </el-form>
  </div>
</template>

<script lang="ts" setup>
import {reactive, ref, watch} from "vue"
import Permission from "../components/permissionSetting/index.vue"
import {selectOption} from "@/api/customer"
import {type reqCustomerEnv, addCustomerEnvApi, reqLicense} from "@/api/license/license"
import {formatDateTime} from "@/utils"
import {rulesHook} from "@/hooks/rulesHook"
import type {FormInstance} from "element-plus"
import {ElMessage} from "element-plus"
import {useRouter} from "vue-router"
import {getAllCustomer} from "@/utils/customer"
import {getCustomerEnvApi} from "@/api/customer/env"
import WarningBar from "@/components/WarningBar/warningBar.vue"

defineOptions({
  name: "License"
})

const {InitRules} = rulesHook()
const formDom = ref<FormInstance>()
const loading = ref(false)
const formData = reactive<reqLicense>({
  customerID: "",
  env: "",
  compDisplayName: "",
  expiredDate: "2999-12-31T00:00:00.000Z",
  managerApis: 9999999,
  managerUsers: 9999999,
  managerProjects: 9999999,
  managerIntegrations: 10,
  managerOrganizations: 100,
  managerTests: 100,
  managerProducts: 100,
  studioUsers: 100,
  studioCompBlackList: "",
  permissionMenu: "",
  checkList: ""
})

const version = ref<String>("")
const router = useRouter()
const customerOptions = ref<selectOption[]>([])
const envOptions = ref<selectOption[]>([])

getAllCustomer().then((Options) => {
  if (Options) {
    customerOptions.value = Options || []
  }
})

// 监听formData.customerID的变化，以便获取对应的环境列表
watch(
  () => formData.customerID,
  (newValue) => {
    // 客户id变了，需要先清空当前选择的env
    formData.env = ""
    loading.value = true
    // 获取客户的所有环境
    getCustomerEnvApi({customerId: newValue})
      .then((res) => {
        envOptions.value = res.data.list.map((item) => {
          return {
            value: item.env,
            label: item.env
          }
        })
        loading.value = false
        if (envOptions.value.length === 0) {
          ElMessage({
            type: "error",
            message: "当前客户没有对应的环境信息，如需创建license，请先创建此客户的环境信息！"
          })
        }
      })
      .catch((error) => {
        loading.value = false
        console.log(error)
      })
  }
)

const form_config_item: { [key: string]: any } = reactive([
  {
    key: "customer",
    customer: [
      {
        type: "select",
        label: "客户名称",
        prop: "customerID",
        options: customerOptions,
        width: "100%",
        col: 12,
        required: true
      },
      {
        type: "select",
        label: "环境",
        prop: "env",
        options: envOptions,
        width: "100%",
        col: 12,
        required: true
      },
      {
        type: "date",
        label: "过期日期",
        prop: "expiredDate",
        date_type: "date",
        col: 12,
        date_format: "YYYY-MM-DD",
        date_value: "YYYY-MM-DD",
        width: "1000px",
        required: true
      },
      {
        type: "input",
        label: "公司显示名",
        prop: "compDisplayName",
        col: 12,
        placeholder: "公司显示用名，可用中文全名（可选）eg: 白山云科技有限公司",
        width: "100%",
        max_length: 20,
        min_length: 2
      }
    ]
  },
  {
    key: "manager",
    manager: [
      {
        type: "inputNumber",
        label: "管理平台API数",
        prop: "managerApis",
        col: 8,
        required: true
      },
      {
        type: "inputNumber",
        label: "管理平台用户数",
        prop: "managerUsers",
        col: 8,
        required: true
      },
      {
        type: "inputNumber",
        label: "管理平台项目数",
        prop: "managerProjects",
        col: 8,
        required: true
      },
      {
        type: "inputNumber",
        label: "管理平台集成平台限制",
        prop: "managerIntegrations",
        col: 8,
        required: true
      },
      {
        type: "inputNumber",
        label: "管理平台支持组织数",
        prop: "managerOrganizations",
        col: 8,
        required: true
      },
      {
        type: "inputNumber",
        label: "管理平台测试限制",
        prop: "managerTests",
        col: 8,
        required: true
      },
      {
        type: "inputNumber",
        label: "管理平台发布限制",
        prop: "managerProducts",
        col: 8,
        required: true
      },
      {
        type: "input",
        label: "版本",
        prop: "version",
        col: 8,
        placeholder: "请输入版本号"
      }
    ]
  },
  {
    key: "runtime",
    runtime: [
      {
        type: "inputNumber",
        label: "集成平台用户数",
        prop: "studioUsers",
        col: 8,
        required: true
      },
      {
        type: "input",
        label: "集成组件黑名单",
        prop: "studioCompBlackList",
        col: 16,
        placeholder: "集成平台组件黑名单（可选），通过英文逗号分隔. eg: InvokeSOAP, InvokeSAPRFC"
      }
    ]
  }
])


// 合并所有的value为一个列表
const mergedValues: any[] = []
form_config_item.forEach((formItem: any) => {
  Object.keys(formItem).forEach((key) => {
    if (formItem[key] !== "key" && Array.isArray(formItem[key])) {
      formItem[key].forEach((item: any) => {
        mergedValues.push(item)
      })
    }
  })
})

const rules = reactive(InitRules(mergedValues))

const handlerSubmitForm = async (Permission: Object, CheckList: object) => {
  Permission["version"] = version.value
  formData.expiredDate = formatDateTime(formData.expiredDate)
  formData.permissionMenu = JSON.stringify(Permission)
  formData.checkList = JSON.stringify(CheckList)

  await formDom.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        const res = await addCustomerEnvApi({...formData})
        ElMessage({type: "success", message: res.msg})
        router.push({name: "licenseList"})
      } finally {
        loading.value = false
      }
    }
  })
}
</script>

<style lang="scss" scoped>
.tooltip-base-box {
  width: 600px;
}

.tooltip-base-box .row {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.tooltip-base-box .center {
  justify-content: center;
}

.tooltip-base-box .box-item {
  width: 110px;
  margin-top: 10px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.box-card {
  width: 480px;
}

.box-row {
  margin-left: 5%;
  margin-right: 5%;
}
</style>
