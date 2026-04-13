<script setup>
import { ref, computed } from "vue";
import ProductCard from "./ProductCard.vue";
import { useProducts } from "../composables/useProducts";

const { products, loading, error, refetch } = useProducts();

const searchQuery = ref("");
const minPrice = ref(0);
const maxPrice = ref(Infinity);

const filteredProducts = computed(() => {
    let result = products.value;

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

const clearFilters = () => {
    searchQuery.value = "";
    minPrice.value = 0;
    maxPrice.value = Infinity;
};

const hasActiveFilters = computed(() => {
    return searchQuery.value.trim() || minPrice.value > 0 || (maxPrice.value !== Infinity && maxPrice.value > 0);
});
</script>

<template>
    <div class="w-full">
        <div v-if="loading" class="flex flex-col items-center justify-center py-20">
            <div class="relative">
                <div class="w-12 h-12 border-4 border-gray-200 rounded-full"></div>
                <div class="absolute top-0 left-0 w-12 h-12 border-4 border-blue-600 border-t-transparent rounded-full animate-spin"></div>
            </div>
            <p class="mt-4 text-gray-500 font-medium">Loading products...</p>
        </div>

        <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-xl p-6 text-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10 mx-auto text-red-500 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
            </svg>
            <p class="text-red-600 mb-4">{{ error }}</p>
            <button
                @click="refetch"
                class="px-5 py-2 bg-red-600 hover:bg-red-700 text-white font-medium rounded-lg transition-colors inline-flex items-center gap-2"
            >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                </svg>
                Try Again
            </button>
        </div>

        <template v-else>
            <div class="bg-white border border-gray-200 rounded-xl p-4 mb-6">
                <div class="flex flex-col lg:flex-row lg:items-end gap-4">
                    <div class="flex-1">
                        <label for="search" class="block text-xs font-semibold text-gray-500 uppercase tracking-wider mb-1.5">Search Products</label>
                        <div class="relative">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                            </svg>
                            <input
                                id="search"
                                v-model="searchQuery"
                                type="text"
                                placeholder="Search by name or category..."
                                class="w-full pl-10 pr-4 py-2.5 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-shadow"
                            />
                        </div>
                    </div>

                    <div class="flex items-end gap-3">
                        <div class="flex flex-col">
                            <label for="minPrice" class="block text-xs font-semibold text-gray-500 uppercase tracking-wider mb-1.5">Min</label>
                            <input
                                id="minPrice"
                                v-model.number="minPrice"
                                type="number"
                                min="0"
                                placeholder="0"
                                class="w-24 px-3 py-2.5 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-shadow"
                            />
                        </div>

                        <span class="text-gray-400 pb-3">—</span>

                        <div class="flex flex-col">
                            <label for="maxPrice" class="block text-xs font-semibold text-gray-500 uppercase tracking-wider mb-1.5">Max</label>
                            <input
                                id="maxPrice"
                                v-model.number="maxPrice"
                                type="number"
                                min="0"
                                placeholder="Any"
                                class="w-24 px-3 py-2.5 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-shadow"
                            />
                        </div>

                        <button
                            v-if="hasActiveFilters"
                            @click="clearFilters"
                            class="px-3 py-2 text-sm text-gray-500 hover:text-gray-700 hover:bg-gray-100 rounded-lg transition-colors"
                        >
                            Clear
                        </button>
                    </div>
                </div>
            </div>

            <div class="mb-4 flex items-center justify-between">
                <p class="text-sm text-gray-500">
                    Showing <span class="font-semibold text-gray-900">{{ filteredProducts.length }}</span> of <span class="font-semibold text-gray-900">{{ products.length }}</span> products
                </p>
            </div>

            <div
                v-if="filteredProducts.length > 0"
                class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-5"
            >
                <ProductCard
                    v-for="product in filteredProducts"
                    :key="product.id"
                    :product="product"
                />
            </div>

            <div v-else class="text-center py-16 bg-white border border-gray-200 rounded-xl">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mx-auto text-gray-300 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <p class="text-gray-600 font-medium mb-2">No products found</p>
                <p class="text-gray-400 text-sm mb-4">Try adjusting your search or filters</p>
                <button
                    @click="clearFilters"
                    class="px-5 py-2 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg transition-colors"
                >
                    Clear All Filters
                </button>
            </div>
        </template>
    </div>
</template>
