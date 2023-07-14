<template>
  <div>
    <div ref="chartContainer" style="width: 100%; height: 190px"></div>
  </div>
</template>

<script lang="ts" setup>
import {ref, onMounted, PropType} from "vue"
import {init as echartsInit, ECharts, EChartsOption} from "echarts"

defineOptions({
  name: "MyEcharts"
})

// 设置全局的 textStyle


const chartContainer = ref<HTMLElement | null>(null)
let chart: ECharts | null = null

onMounted(() => {
  if (chartContainer.value) {
    chart = echartsInit(chartContainer.value)
    chart.setOption(getChartOption())
  }
})

const textStyle = {
  fontWeight: "normal",
  fontSize: 16,
  color: "#409EFF"
}

const getChartOption = (): EChartsOption => {
  return {
    tooltip: {},
    series: [
      {
        name: "客户类型",
        type: "pie",
        radius: "50%",
        center: ["10%", "50%"],
        encode: {
          itemName: 0,
          value: 1
        },
        label: {
          show: true,
          formatter: "{b}: {c}"
        },
        data: props.data.customerType
      },
      {
        name: "客户维保",
        type: "pie",
        radius: "50%",
        center: ["30%", "50%"],
        encode: {
          itemName: 0,
          value: 1
        },
        label: {
          show: true,
          formatter: "{b}: {c}"
        },
        data: props.data.technicalSupport
      },
      {
        name: "维保类型",
        type: "pie",
        radius: "50%",
        center: ["50%", "50%"],
        encode: {
          itemName: 0,
          value: 1
        },
        label: {
          show: true,
          formatter: "{b}: {c}"
        },
        data: props.data.maintenanceType
      },
      {
        name: "ENV",
        type: "pie",
        radius: "50%",
        center: ["70%", "50%"],
        encode: {
          itemName: 0,
          value: 1
        },
        label: {
          show: true,
          formatter: "{b}: {c}"
        },
        data: props.data.env
      },
      {
        name: "license",
        type: "pie",
        radius: "50%",
        center: ["90%", "50%"],
        encode: {
          itemName: 0,
          value: 1
        },
        label: {
          show: true,
          formatter: "{b}: {c}"
        },
        data: props.data.license
      }
    ],
    title: [
      {
        text: "客户类型分析",
        top: "1%",
        left: "%5",
        textStyle
      },
      {
        text: "维保数量分析",
        top: "1%",
        left: "20%",
        textStyle
      },
      {
        text: "维保类型分析",
        top: "1%",
        left: "40%",
        textStyle
      },
      {
        text: "环境数量分析",
        top: "1%",
        left: "60%",
        textStyle
      },
      {
        text: "license数量分析",
        top: "1%",
        left: "80%",
        textStyle
      }
    ]
  }
}

const props = defineProps({
  data: {
    type: Object as PropType<{
      customerType?: any[];
      technicalSupport?: any[];
      env?: any[];
      license?: any[];
    }>,
    default: () => ({
      customerType: [],
      technicalSupport: [],
      env: [],
      license: []
    })
  }
})
</script>

<style scoped></style>
