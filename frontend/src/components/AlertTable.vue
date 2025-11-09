<script lang="ts" setup>
import { defineProps } from 'vue'
import type { Alert } from '@/services/alertService'

defineProps<{ alerts: Alert[]; loading: boolean }>()
</script>

<template>
  <table class="min-w-full divide-y divide-gray-400 border border-gray-400 rounded-lg">
    <thead class="bg-green-200">
      <tr>
        <th class="px-4 py-2 text-left text-sm font-semibold text-black">Instance</th>
        <th class="px-4 py-2 text-left text-sm font-semibold text-black">Alertname</th>
        <th class="px-4 py-2 text-left text-sm font-semibold text-black">Status</th>
      </tr>
    </thead>
    <tbody class="divide-y divide-gray-200">
      <tr v-for="alert in alerts" :key="alert.fingerprint">
        <td class="px-4 py-2">{{ alert.labels.instance }}</td>
        <td class="px-4 py-2">{{ alert.alertname }}</td>
        <td class="px-4 py-2">
          <span
            :class="{
              'text-red-600 font-semibold': alert.status === 'active',
              'text-green-600': alert.status === 'resolved',
              'text-black': !['active', 'resolved'].includes(alert.status),
            }"
          >
            {{ alert.status }}
          </span>
        </td>
      </tr>
    </tbody>
  </table>
</template>
