<script setup>
import otsHeader from './components/otsHeader.vue';
import otsFooter from './components/otsFooter.vue';
</script>

<template>
    <ots-header></ots-header>
    <main>
        <router-view></router-view>
    </main>
    <ots-footer></ots-footer>
</template>

<style scoped>

</style>
