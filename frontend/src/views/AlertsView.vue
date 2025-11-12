<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { fetchAlerts, type Alert } from '@/services/alertService'
import AlertTable from '../components/AlertTable.vue'

const alerts = ref<Alert[]>([])
const loading = ref(true)

onMounted(async () => {
  try {
    setInterval(async () => {
        alerts.value = await fetchAlerts()
        console.log("fetch alerts")
    }, 3000);
  } catch (err) {
    console.error('Failed to fetch alerts:', err)
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold mb-4">Active Alerts</h1>

    <div v-if="loading" class="text-gray-500">Loading...</div>
    <div v-else>
      <AlertTable :alerts="alerts" :loading="loading" />
    </div>
  </div>
</template>
