<template>
    <div class="hamburger-menu-container">
        <div class="hamburger-menu" :class="{ active: isOpen }" @click="toggleMenu">
            <div class="hamburger-line"></div>
            <div class="hamburger-line"></div>
            <div class="hamburger-line"></div>
        </div>
        <ul class="menu-items" :class="{ active: isOpen }">
            <li><img src="/src/assets/about.png" /></li>
        </ul>
    </div>
</template>

<script lang="ts">
export default {
    data() {
        return {
            isOpen: false
        }
    },
    methods: {
        toggleMenu() {
            this.isOpen = !this.isOpen;
        }
    }
}
</script>