import { api } from "../api.js";

const PRODUCTS_PER_PAGE = 12;

function LoadingSpinner() {
  return `
    <div class="flex justify-center items-center py-16">
      <div class="animate-spin rounded-full h-12 w-12 border-4 border-indigo-200 border-t-indigo-600" role="status" aria-label="Loading products">
        <span class="sr-only">Loading...</span>
      </div>
    </div>
  `;
}

function SkeletonCard() {
  return `
    <div class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden animate-pulse">
      <div class="bg-gray-100 w-full h-48"></div>
      <div class="p-4 space-y-3">
        <div class="h-4 bg-gray-200 rounded w-1/3"></div>
        <div class="h-6 bg-gray-200 rounded w-3/4"></div>
        <div class="h-8 bg-gray-200 rounded w-1/4"></div>
        <div class="h-4 bg-gray-200 rounded w-full"></div>
      </div>
    </div>
  `;
}

export async function renderHomePage(container, { navigate, showToast }) {
  let currentPage = 1;
  let totalPages = 1;
  let categories = [];
  let categoryMap = {};
  let currentFilters = {
    search: "",
    category: "",
    minPrice: "",
    maxPrice: "",
  };

  container.innerHTML = `
    <section aria-labelledby="products-heading" class="pt-20 md:pt-24 py-6 lg:py-10">
      <div class="flex flex-col md:flex-row md:items-center justify-between mb-8 gap-4">
        <p id="results-count" class="text-gray-500 text-sm" aria-live="polite"></p>
      </div>

      <aside class="bg-white rounded-2xl shadow-sm border border-gray-100 p-5 mb-8" aria-label="Product filters">
        <form id="filter-form" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-5 gap-4">
          <div class="space-y-1">
            <label for="search-input" class="block text-sm font-medium text-gray-700">Search</label>
            <input
              type="text"
              id="search-input"
              name="search"
              placeholder="Search products..."
              class="w-full px-4 py-2.5 border border-gray-300 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors"
            />
          </div>
          <div class="space-y-1">
            <label for="category-select" class="block text-sm font-medium text-gray-700">Category</label>
            <select
              id="category-select"
              name="category"
              class="w-full px-4 py-2.5 border border-gray-300 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors bg-white"
            >
              <option value="">All Categories</option>
            </select>
          </div>
          <div class="space-y-1">
            <label for="min-price" class="block text-sm font-medium text-gray-700">Min Price</label>
            <input
              type="number"
              id="min-price"
              name="minPrice"
              placeholder="0"
              min="0"
              step="0.01"
              class="w-full px-4 py-2.5 border border-gray-300 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors"
            />
          </div>
          <div class="space-y-1">
            <label for="max-price" class="block text-sm font-medium text-gray-700">Max Price</label>
            <input
              type="number"
              id="max-price"
              name="maxPrice"
              placeholder="9999"
              min="0"
              step="0.01"
              class="w-full px-4 py-2.5 border border-gray-300 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors"
            />
          </div>
          <div class="flex items-end gap-2 sm:col-span-2 lg:col-span-1">
            <button
              type="submit"
              class="flex-1 sm:flex-none px-5 py-2.5 bg-indigo-600 text-white font-medium rounded-xl hover:bg-indigo-700 active:scale-95 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
            >
              <span class="flex items-center justify-center gap-2">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                </svg>
                Filter
              </span>
            </button>
            <button
              type="button"
              id="clear-filters"
              class="px-4 py-2.5 border border-gray-300 text-gray-700 font-medium rounded-xl hover:bg-gray-50 active:scale-95 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-gray-300 focus:ring-offset-2"
            >
              Clear
            </button>
          </div>
        </form>
      </aside>

      <div id="loading" aria-live="polite">${LoadingSpinner()}</div>
      <div id="skeleton-grid" class="hidden grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6" aria-hidden="true">
        ${Array(PRODUCTS_PER_PAGE).fill(SkeletonCard()).join("")}
      </div>
      <div id="error" class="hidden bg-red-50 border border-red-200 text-red-700 px-5 py-5 rounded-2xl text-center" role="alert">
        <svg class="w-12 h-12 mx-auto mb-3 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        <p id="error-message"></p>
      </div>
      <div id="empty" class="hidden text-center py-16">
        <div class="inline-flex items-center justify-center w-32 h-32 rounded-3xl bg-gradient-to-br from-gray-50 to-indigo-50/50 border border-gray-100 mb-6">
          <svg class="w-16 h-16 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4"></path>
          </svg>
        </div>
        <p class="text-xl font-semibold text-gray-900 mb-2">No products found</p>
        <p class="text-gray-500">Try adjusting your filters or search terms</p>
      </div>
      <div id="products" class="hidden grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6" role="list" aria-label="Products"></div>

      <nav id="pagination" class="hidden flex items-center justify-center gap-3 mt-10" aria-label="Product pagination">
          <button
            id="prev-btn"
            class="flex items-center gap-2 px-4 py-2.5 border border-gray-300 text-gray-700 font-medium rounded-xl hover:bg-gray-50 active:scale-95 transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:bg-white focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
            aria-label="Previous page"
          >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
          </svg>
          Previous
        </button>
        <span id="page-info" class="px-4 py-2 text-gray-600 font-medium" aria-current="page"></span>
          <button
            id="next-btn"
            class="flex items-center gap-2 px-4 py-2.5 border border-gray-300 text-gray-700 font-medium rounded-xl hover:bg-gray-50 active:scale-95 transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:bg-white focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
            aria-label="Next page"
          >
          Next
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </button>
      </nav>
    </section>
  `;

  const searchInput = document.getElementById("search-input");
  const categorySelect = document.getElementById("category-select");
  const minPriceInput = document.getElementById("min-price");
  const maxPriceInput = document.getElementById("max-price");
  const filterForm = document.getElementById("filter-form");
  const clearFiltersBtn = document.getElementById("clear-filters");
  const productsContainer = document.getElementById("products");
  const skeletonGrid = document.getElementById("skeleton-grid");
  const loadingEl = document.getElementById("loading");
  const errorEl = document.getElementById("error");
  const errorMessage = document.getElementById("error-message");
  const emptyEl = document.getElementById("empty");
  const resultsCount = document.getElementById("results-count");
  const paginationEl = document.getElementById("pagination");
  const prevBtn = document.getElementById("prev-btn");
  const nextBtn = document.getElementById("next-btn");
  const pageInfo = document.getElementById("page-info");

  async function loadCategories() {
    try {
      const data = await api.categories.list();
      categories = data.data || [];
      categoryMap = {};
      categories.forEach((cat) => {
        categoryMap[cat.category_id] = cat.name;
      });
      categorySelect.innerHTML =
        '<option value="">All Categories</option>' +
        categories
          .map(
            (cat) =>
              `<option value="${cat.category_id}">${escapeHtml(cat.name)}</option>`,
          )
          .join("");
    } catch (err) {
      console.error("Failed to load categories:", err);
    }
  }

  async function loadProducts(page = 1) {
    loadingEl.classList.add("hidden");
    skeletonGrid.classList.remove("hidden");
    skeletonGrid.setAttribute("aria-hidden", "false");
    errorEl.classList.add("hidden");
    emptyEl.classList.add("hidden");
    productsContainer.classList.add("hidden");
    paginationEl.classList.add("hidden");
    productsContainer.innerHTML = "";

    try {
      const params = {
        page,
        limit: PRODUCTS_PER_PAGE,
      };

      if (currentFilters.search) params.search = currentFilters.search;
      if (currentFilters.category) params.category = currentFilters.category;
      if (currentFilters.minPrice) params.min_price = currentFilters.minPrice;
      if (currentFilters.maxPrice) params.max_price = currentFilters.maxPrice;

      const data = await api.products.list(params);
      const products = data.data || [];
      const totalItems = data.pagination?.total_items || 0;
      totalPages = data.pagination?.total_pages || 1;
      currentPage = page;

      skeletonGrid.classList.add("hidden");
      skeletonGrid.setAttribute("aria-hidden", "true");

      if (products.length === 0) {
        emptyEl.classList.remove("hidden");
        resultsCount.textContent = "No products found";
        return;
      }

      resultsCount.textContent = `Showing ${products.length} of ${totalItems} products`;

      productsContainer.innerHTML = products
        .map((product) => {
          const categoryName = categoryMap[product.category_id] || "";
          return `
        <article class="bg-white rounded-2xl shadow-sm border border-gray-100 hover:shadow-lg hover:-translate-y-1 hover:border-indigo-200 transition-all duration-300 overflow-hidden" role="listitem">
           <a href="/products/${product.product_id}" class="product-link block focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500 active:scale-95 transition-all duration-200" data-route>
            <div class="bg-gradient-to-br from-gray-100 to-gray-200 w-full h-48 flex items-center justify-center relative overflow-hidden group">
              <svg class="w-16 h-16 text-gray-400 transition-transform duration-300 group-hover:scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
            </div>
            <div class="p-5">
              ${categoryName ? `<span class="inline-block px-3 py-1 text-xs font-semibold text-indigo-700 bg-indigo-50 rounded-full uppercase tracking-wide">${escapeHtml(categoryName)}</span>` : ""}
               <h2 class="font-semibold text-lg text-gray-900 mt-3 line-clamp-2 group-hover:text-indigo-600 transition-all duration-200">${escapeHtml(product.name)}</h2>
              <p class="text-2xl font-bold text-gray-900 mt-2">$${product.price?.toFixed(2) || "0.00"}</p>
              <p class="text-gray-500 text-sm mt-3 line-clamp-2">${escapeHtml(product.description || "No description available.")}</p>
               <div class="mt-4 flex items-center text-indigo-600 font-medium text-sm group-hover:gap-2 transition-all duration-200 active:scale-95">
                 <span>View Details</span>
                <svg class="w-4 h-4 transition-transform duration-300 group-hover:translate-x-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                </svg>
              </div>
            </div>
          </a>
        </article>
      `;
        })
        .join("");

      productsContainer.classList.remove("hidden");
      updatePagination();
      paginationEl.classList.remove("hidden");
    } catch (err) {
      skeletonGrid.classList.add("hidden");
      skeletonGrid.setAttribute("aria-hidden", "true");
      errorMessage.textContent = err.message;
      errorEl.classList.remove("hidden");
    }
  }

  function updatePagination() {
    prevBtn.disabled = currentPage <= 1;
    nextBtn.disabled = currentPage >= totalPages;
    pageInfo.textContent = `Page ${currentPage} of ${totalPages}`;
  }

  filterForm.addEventListener("submit", (e) => {
    e.preventDefault();
    currentFilters = {
      search: searchInput.value.trim(),
      category: categorySelect.value,
      minPrice: minPriceInput.value,
      maxPrice: maxPriceInput.value,
    };
    loadProducts(1);
  });

  clearFiltersBtn.addEventListener("click", () => {
    searchInput.value = "";
    categorySelect.value = "";
    minPriceInput.value = "";
    maxPriceInput.value = "";
    currentFilters = { search: "", category: "", minPrice: "", maxPrice: "" };
    loadProducts(1);
  });

  prevBtn.addEventListener("click", () => {
    if (currentPage > 1) loadProducts(currentPage - 1);
  });

  nextBtn.addEventListener("click", () => {
    if (currentPage < totalPages) loadProducts(currentPage + 1);
  });

  await loadCategories();

  const url = new URL(window.location.href);
  currentFilters.search = url.searchParams.get("search") || "";
  searchInput.value = currentFilters.search;
  loadProducts(1);
}

function escapeHtml(text) {
  const div = document.createElement("div");
  div.textContent = text;
  return div.innerHTML;
}