<template>
    <el-card class="box-card">
        <template #header>
            <div class="card-header">
                <span>Request Records</span>
            </div>
        </template>

        <el-table show-header :data="requestRecords" stripe>
            <el-table-column prop="dateTime" label="Date + Time" />
            <el-table-column prop="IMEI" label="IMEI" />
            <el-table-column prop="type" label="Type" />
            <el-table-column prop="payload" label="Payload (hex)" />
            <el-table-column prop="base64" label="Payload (base64)" />
            <el-table-column prop="messageID" label="Message ID" />
            <el-table-column prop="flush" label="Flush" />
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