<template>
  <div class="app-container">
    <el-space wrap fill style="width: 100%" size="large">
      <el-row :gutter="20" v-if="customerData.total">
        <el-col :span="8">
          <el-card shadow="hover">
            <el-statistic title="累计客户数量" :value="customerData.total.customer"/>
          </el-card>
        </el-col>

        <el-col :span="8">
          <el-card shadow="hover">
            <el-statistic title="累计部署的环境数量" :value="customerData.total.env"/>
          </el-card>
        </el-col>

        <el-col :span="8">
          <el-card shadow="hover">
            <el-statistic title="License数量" :value="customerData.total.license"/>
          </el-card>
        </el-col>
      </el-row>
      <el-card v-else>
        <el-skeleton :rows="1" animated/>
      </el-card>

      <el-card>
        <customer-charts v-if="customerData.analyse" :data="customerData.analyse"></customer-charts>
        <!--        骨架屏-->
        <div v-else class="skeleton-container">
          <template v-for="index in 5" :key="index">
            <el-skeleton style="--el-skeleton-circle-size: 150px">
              <template #template>
                <el-skeleton-item variant="circle"/>
              </template>
            </el-skeleton>
          </template>
        </div>
      </el-card>
      <el-card header="客户历史重要事件">
        <!--        客户事件搜索表单-->
        <el-form :inline="true"
                 :model="eventSearch"
                 class="demo-form-inline"
        >
          <el-form-item label="选择客户：">
            <el-select v-model="eventSearch.customerID" placeholder="请选择" clearable filterable>
              <el-option
                v-for="select in customerOptions"
                :key="select.value"
                :label="select.label"
                :value="select.value"
              ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="选择日期：">
            <el-date-picker
              v-model="eventSearch.timeRange"
              type="datetimerange"
              unlink-panels
              range-separator="~"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
              :shortcuts="shortcuts"
              format="YYYY/MM/DD HH:mm:ss"
              value-format="YYYY-MM-DD HH:mm:ss"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSearch('init')">查询</el-button>
          </el-form-item>
        </el-form>
        <!--        事件线-->
        <el-empty
          v-if="events.length === 0"
          v-loading="eventSearch.loading"
          description="请选择客户，以便查看其相关事件"
        />
        <div v-else style="margin-top: 50px; margin-right: 50px">
          <el-timeline v-loading="eventSearch.loading">
            <el-timeline-item
              v-for="(activity, index) in events"
              :key="index"
              type="primary"
              hollow
              :timestamp="activity.eventTime"
              placement="top"
            >
              <el-card class="box-card">
                {{ activity.eventContent }}
              </el-card>
            </el-timeline-item>
          </el-timeline>

          <div class="button-center">
            <el-button type="text" @click="onSearch()" :loading="eventSearch.loading" v-if="eventSearch.loadingMore"
            >加载更多
            </el-button>
            <el-text v-else type="primary">没有更多了~</el-text>
          </div>
        </div>
      </el-card>
    </el-space>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from "vue"
import {shortcuts} from "@/config/dateDict"
import {selectOption} from "@/api/customer"
import {getAllCustomer} from "@/utils/customer"
import {getCustomerEventApi} from "@/api/customer/event"
import {ElMessage} from "element-plus"
import {formatDateTime, formatTimestamp} from "@/utils"
import CustomerCharts from "./components/echarts/index.vue"
import {getSummaryApi} from "@/api/customer/summary"

defineOptions({
  name: "Summary"
})

// onMounted(() => {
//   console.log(`the component is now mounted.`)
// })

interface EventItem {
  eventContent: string
  eventTime: string
}

const customerOptions = ref<selectOption[]>([])
const eventSearch = reactive({
  customerID: "",
  timeRange: "",
  pageSize: 30,
  currentPage: 1,
  loading: false,
  loadingMore: true
})

getAllCustomer().then((Options) => {
  if (Options) {
    customerOptions.value = Options || []
  }
})

const events = ref<EventItem[]>([])

const onSearch = (action?: string) => {
  if (eventSearch.customerID === "") {
    ElMessage({type: "error", message: "请选择客户名称！"})
    return
  }
  eventSearch.loading = true
  if (action === "init") {
    events.value = []
    eventSearch.currentPage = 1
    eventSearch.loadingMore = true
  }

  let startTimestamp
  let endTimestamp
  if (typeof eventSearch.timeRange === "object" && Object.keys(eventSearch.timeRange).length === 2) {
    let s = eventSearch.timeRange[0]
    let e = eventSearch.timeRange[1]
    startTimestamp = formatTimestamp(s)
    endTimestamp = formatTimestamp(e)
  }

  getCustomerEventApi({
    customerID: eventSearch.customerID,
    size: eventSearch.pageSize,
    currentPage: eventSearch.currentPage,
    startTime: startTimestamp || undefined,
    endTime: endTimestamp || undefined,
    sort: "eventTime"
  })
    .then((res) => {
      const formattedList = res.data.list.map((item) => {
        return {
          ...item,
          eventTime: formatDateTime(item.eventTime)
        }
      })

      events.value.push(...formattedList)
      // 如果当前事件列表元素数量大于查询结果总数，则不显示加载更多按钮
      if (events.value.length >= res.data.total) {
        eventSearch.loadingMore = false
      } else {
        eventSearch.currentPage += 1
      }
      eventSearch.loading = false
    })
    .catch((err) => {
      console.log(err)
      eventSearch.loading = false
    })
}

const customerData = ref<any>({
  total: null,
  analyse: null
})

async function fetchSummaryData() {
  try {
    const res = await getSummaryApi()
    if (res) {
      customerData.value = res.data.list
    }
  } catch (error) {
    console.log(error)
  }
}

fetchSummaryData()
</script>

<style scoped>
.box-card {
  width: 90%;
}

.skeleton-container {
  display: flex;
  justify-content: space-between;
}
</style>
