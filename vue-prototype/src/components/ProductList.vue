<script setup>
import { ref, computed } from "vue";
import ProductCard from "./ProductCard.vue";

const props = defineProps({
    products: {
        type: Array,
        required: true,
    },
});

const emit = defineEmits(["add-to-cart"]);

const searchQuery = ref("");
const minPrice = ref(0);
const maxPrice = ref(Infinity);

const filteredProducts = computed(() => {
    let result = props.products;

    if (searchQuery.value.trim()) {
        const query = searchQuery.value.toLowerCase().trim();
        result = result.filter(
            (product) =>
                product.name.toLowerCase().includes(query) ||
                product.category.toLowerCase().includes(query),
        );
    }

    if (minPrice.value > 0) {
        result = result.filter((product) => product.price >= minPrice.value);
    }

    if (maxPrice.value !== Infinity && maxPrice.value > 0) {
        result = result.filter((product) => product.price <= maxPrice.value);
    }

    return result;
});

const handleAddToCart = (product) => {
    emit("add-to-cart", product);
};
</script>

<template>
    <div class="product-list">
        <div class="filters">
            <div class="filter-group">
                <label for="search">Search:</label>
                <input
                    id="search"
                    v-model="searchQuery"
                    type="text"
                    placeholder="Search by name or category..."
                />
            </div>

            <div class="filter-group price-filters">
                <label for="minPrice">Min Price:</label>
                <input
                    id="minPrice"
                    v-model.number="minPrice"
                    type="number"
                    min="0"
                    placeholder="0"
                />

                <label for="maxPrice">Max Price:</label>
                <input
                    id="maxPrice"
                    v-model.number="maxPrice"
                    type="number"
                    min="0"
                    placeholder="∞"
                />
            </div>
        </div>

        <div v-if="filteredProducts.length > 0" class="products-grid">
            <ProductCard
                v-for="product in filteredProducts"
                :key="product.id"
                :product="product"
                @add-to-cart="handleAddToCart"
            />
        </div>

        <div v-else class="no-results">
            <p>No products match your filters</p>
            <button
                @click="
                    searchQuery = '';
                    minPrice = 0;
                    maxPrice = Infinity;
                "
            >
                Clear Filters
            </button>
        </div>
    </div>
</template>

<style scoped>
.product-list {
    width: 100%;
}

.filters {
    display: flex;
    flex-wrap: wrap;
    gap: 16px;
    margin-bottom: 24px;
    padding: 16px;
    background: #f8f9fa;
    border-radius: 8px;
}

.filter-group {
    display: flex;
    flex-direction: column;
    gap: 6px;
}

.filter-group label {
    font-size: 12px;
    font-weight: 500;
    color: #666;
    text-transform: uppercase;
}

.filter-group input {
    padding: 8px 12px;
    border: 1px solid #ddd;
    border-radius: 6px;
    font-size: 14px;
    min-width: 180px;
}

.price-filters {
    flex-direction: row;
    align-items: flex-end;
    gap: 12px;
}

.price-filters input {
    min-width: 100px;
}

.products-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 20px;
}

.no-results {
    text-align: center;
    padding: 48px;
    background: #f8f9fa;
    border-radius: 8px;
}

.no-results p {
    margin: 0 0 16px;
    font-size: 16px;
    color: #666;
}

.no-results button {
    padding: 10px 20px;
    background: #2c5282;
    color: white;
    border: none;
    border-radius: 6px;
    cursor: pointer;
    font-size: 14px;
}

.no-results button:hover {
    background: #1a365d;
}
</style>
