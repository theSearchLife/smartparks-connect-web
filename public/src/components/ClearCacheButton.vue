<template>
    <div>
        <el-button type="primary" @click="clearCache">Clear Cache</el-button>
    </div>
</template>
<script>
export default {
    methods: {
        clearCache() {
            if (confirm("Are you sure you want to clear all data?")) {
                localStorage.clear();
                window.location.reload(true);
            }
        }
    }
};
</script>
