<script setup>
import { ref } from 'vue'
const actual = ref(new Date(Date.now()).toLocaleDateString())
</script>

<template>
    <footer>
        <span class="copy">&COPY; from</span>
        <span class="author">Schmid Lothar</span>
        <span class="date">{{ actual }}</span>
    </footer>
</template>

<style scoped>

</style>