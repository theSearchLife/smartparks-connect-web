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

<script>
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

<style>
.hamburger-menu-container {
    position: relative;
}

.hamburger-menu {
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    height: 24px;
    width: 24px;
    cursor: pointer;
}

.hamburger-line {
    width: 100%;
    height: 2px;
    background-color: #000;
    transition: all 0.2s ease-in-out;
}

.active .hamburger-line:nth-child(1) {
    transform: rotate(45deg) translate(5px, 5px);
}

.active .hamburger-line:nth-child(2) {
    opacity: 0;
}

.active .hamburger-line:nth-child(3) {
    transform: rotate(-45deg) translate(5px, -5px);
}

.menu-items {
    position: absolute;
    top: 100%;
    left: 0;
    width: 100%;
    padding: 0;
    margin: 0;
    list-style: none;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
    transition: all 0.2s ease-in-out;
    opacity: 0;
    pointer-events: none;
}

.active {
    opacity: 1;
    pointer-events: auto;
    z-index: 100;
}

.menu-items li {
    padding: 12px 16px;
    border-bottom: 1px solid #ddd;
}

.menu-items li:last-child {
    border-bottom: none;
}

.menu-items li a {
    color: #333;
    text-decoration: none;
}

.menu-items li a:hover {
    color: #000;
}
</style>
