<template>
  <div class="filter-bar">
    <button @click="applyFilter('all')">All</button>
    <button @click="applyFilter('vegan')">Vegan</button>
    <button @click="applyFilter('gluten-free')">Gluten-Free</button>
    <!-- Add more filter options as needed -->
  </div>
</template>

<script>
export default {
  name: 'FilterBar',
  methods: {
    applyFilter(filter) {
      this.$emit('filter', filter);
    }
  }
}
</script>

<style scoped>
.filter-bar {
  display: flex;
  gap: 10px;
}
</style>
