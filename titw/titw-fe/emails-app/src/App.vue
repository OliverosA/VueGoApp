<template>
  <!-- Manteniendo el estado del componente (el estado se pierde
  al recargar la pagina) -->
  <router-view v-slot="{ Component, route }">
    <keep-alive>
      <component :is="Component" :key="route.name" />
    </keep-alive>
  </router-view>
</template>

<style>
body {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  background: linear-gradient(
    0deg,
    rgba(183, 228, 255, 1) 7%,
    rgba(26, 110, 217, 1) 81%
  );
  width: 100%;
  min-height: 100vh;
}
</style>
