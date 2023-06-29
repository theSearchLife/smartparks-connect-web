<template>
    <el-card class="box-card">
        <template #header>
            <div class="card-header">
                <span>Request Records</span>
            </div>
        </template>

        <el-table show-header :data="requestRecords" stripe>
            <el-table-column prop="dateTime" label="Date + Time" />
            <el-table-column prop="name" label="Name" />
            <el-table-column prop="devEUI" label="EUI" />
            <el-table-column prop="type" label="Type" />
            <el-table-column prop="port" label="Port" />
            <el-table-column prop="payload" label="Payload (hex)" />
            <el-table-column prop="base64" label="Payload (base64)" />
            <el-table-column prop="fCnt" label="FCnt" />
            <el-table-column prop="confirmed" label="Confirmed" />
        </el-table>
    </el-card>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
    name: "ChirpstackRequests",
    props: {
        requestRecords: {
            type: Array,
            required: true
        }
    },
});
</script>