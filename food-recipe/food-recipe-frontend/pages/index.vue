<template>
  <div>
    <h1>Recipe Categories</h1>
    <div class="categories">
      <!-- Categories will be dynamically rendered here -->
    </div>

    <h2>Featured Recipes</h2>
    <div class="recipes">
      <!-- Recipes will be dynamically rendered here -->
      <div v-for="recipe in recipes" :key="recipe.id">
        <h3>{{ recipe.title }}</h3>
        <p>{{ recipe.description }}</p>
      </div>
    </div>

    <input type="text" placeholder="Search recipes by title" v-model="searchQuery" />
    
    <FilterBar @filter="applyFilter" />
  </div>
</template>

<script>
import FilterBar from 'components/FilterBar.vue';

export default {
  name: 'IndexPage',
  components: {
    FilterBar
  },
  data() {
    return {
      searchQuery: '',
      recipes: [] // Array to hold fetched recipes
    };
  },
  methods: {
    applyFilter(filters) {
      // Logic to apply filters
    },
    async fetchRecipes() {
      try {
        const response = await this.$axios.get(`${this.$config.apiUrl}/recipes`); // Fetch recipes from backend
        this.recipes = response.data; // Store fetched recipes
      } catch (error) {
        console.error('Error fetching recipes:', error);
      }
    }
  },
  mounted() {
    this.fetchRecipes(); // Fetch recipes when the component is mounted
  }
}
</script>

<style scoped>
/* Add styles for categories and recipes */
</style>
