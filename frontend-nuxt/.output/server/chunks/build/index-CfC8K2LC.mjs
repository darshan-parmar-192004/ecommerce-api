import { u as useRoute, c as useRouter, a as useSeoMeta, s as script } from './server.mjs';
import { ref, watch, computed, mergeProps, withCtx, isRef, unref, createVNode, withKeys, useSSRContext } from 'vue';
import { u as useApi } from './useApi-CR6WE1xi.mjs';
import { u as useAsyncData } from './asyncData--maOkXya.mjs';
import script$6 from './index-BGjNF9Dd.mjs';
import script$5 from './index-CMnz8U93.mjs';
import script$4 from './index-BdDxcS0p.mjs';
import script$3 from './index-DDXXjm8b.mjs';
import script$2 from './index-IyD_XvNl.mjs';
import script$1 from './index-83EVY8bO.mjs';
import { ssrRenderAttrs, ssrRenderComponent, ssrRenderList, ssrInterpolate, ssrRenderStyle } from 'vue/server-renderer';
import { P as ProductCard } from './ProductCard-Di5Z6QG2.mjs';
import { _ as _export_sfc } from './_plugin-vue_export-helper-1tPrXgE0.mjs';
import '../nitro/nitro.mjs';
import 'node:http';
import 'node:https';
import 'node:events';
import 'node:buffer';
import 'node:fs';
import 'node:path';
import 'node:crypto';
import '@primevue/core/base/style';
import '@primevue/core/basecomponent/style';
import '@primeuix/styles/autocomplete';
import '@primeuix/utils/object';
import '@primeuix/styles/cascadeselect';
import '@primeuix/styles/checkbox';
import '@primeuix/styles/checkboxgroup';
import '@primeuix/styles/colorpicker';
import '@primeuix/styles/datepicker';
import '@primeuix/styles/floatlabel';
import '@primeuix/styles/iconfield';
import '@primeuix/styles/iftalabel';
import '@primeuix/styles/inputchips';
import '@primeuix/styles/inputgroup';
import '@primeuix/styles/inputnumber';
import '@primeuix/styles/inputotp';
import '@primeuix/styles/inputtext';
import '@primeuix/styles/knob';
import '@primeuix/styles/listbox';
import '@primeuix/styles/multiselect';
import '@primeuix/styles/password';
import '@primeuix/styles/radiobutton';
import '@primeuix/styles/radiobuttongroup';
import '@primeuix/styles/rating';
import '@primeuix/styles/select';
import '@primeuix/styles/selectbutton';
import '@primeuix/styles/slider';
import '@primeuix/styles/textarea';
import '@primeuix/styles/togglebutton';
import '@primeuix/styles/toggleswitch';
import '@primeuix/styles/treeselect';
import '@primeuix/styles/button';
import '@primeuix/styles/buttongroup';
import '@primeuix/styles/speeddial';
import '@primeuix/styles/splitbutton';
import '@primeuix/styles/datatable';
import '@primeuix/styles/dataview';
import '@primeuix/styles/orderlist';
import '@primeuix/styles/organizationchart';
import '@primeuix/styles/paginator';
import '@primeuix/styles/picklist';
import '@primeuix/styles/tree';
import '@primeuix/styles/treetable';
import '@primeuix/styles/timeline';
import '@primeuix/styles/virtualscroller';
import '@primeuix/styles/accordion';
import '@primeuix/styles/card';
import '@primeuix/styles/divider';
import '@primeuix/styles/fieldset';
import '@primeuix/styles/panel';
import '@primeuix/styles/scrollpanel';
import '@primeuix/styles/splitter';
import '@primeuix/styles/stepper';
import '@primeuix/styles/tabview';
import '@primeuix/styles/tabs';
import '@primeuix/styles/toolbar';
import '@primeuix/styles/confirmdialog';
import '@primeuix/styles/confirmpopup';
import '@primeuix/styles/dialog';
import '@primeuix/styles/drawer';
import '@primeuix/styles/popover';
import '@primeuix/styles/fileupload';
import '@primeuix/styles/breadcrumb';
import '@primeuix/styles/contextmenu';
import '@primeuix/styles/dock';
import '@primeuix/styles/menu';
import '@primeuix/styles/menubar';
import '@primeuix/styles/megamenu';
import '@primeuix/styles/panelmenu';
import '@primeuix/styles/steps';
import '@primeuix/styles/tabmenu';
import '@primeuix/styles/tieredmenu';
import '@primeuix/styles/message';
import '@primeuix/styles/inlinemessage';
import '@primeuix/styles/toast';
import '@primeuix/styles/carousel';
import '@primeuix/styles/galleria';
import '@primeuix/styles/image';
import '@primeuix/styles/imagecompare';
import '@primeuix/styles/avatar';
import '@primeuix/styles/badge';
import '@primeuix/styles/blockui';
import '@primeuix/styles/chip';
import '@primeuix/styles/inplace';
import '@primeuix/styles/metergroup';
import '@primeuix/styles/overlaybadge';
import '@primeuix/styles/scrolltop';
import '@primeuix/styles/skeleton';
import '@primeuix/styles/progressbar';
import '@primeuix/styles/progressspinner';
import '@primeuix/styles/tag';
import '@primeuix/styles/terminal';
import '@primevue/forms/form/style';
import '@primevue/forms/formfield/style';
import '@primeuix/styles/tooltip';
import '@primeuix/styles/ripple';
import '@primeuix/styled';
import 'unhead';
import '@unhead/shared';
import 'vue-router';
import '@primeuix/utils/eventbus';
import '@primeuix/utils';
import '@primeuix/styles/base';
import '@primeuix/utils/dom';
import '@primeuix/utils/zindex';
import '@primeuix/utils/uuid';
import './index-mL8pBhlS.mjs';
import './index-BsfLo0vt.mjs';
import './index-Bg3yUxEP.mjs';
import './index-D2H2HXHf.mjs';
import './index-D4owPo0H.mjs';
import './index-D255zs2l.mjs';
import './index-Dedjwyps.mjs';
import './index-DrzwqwtL.mjs';
import './index-rAVNvoJo.mjs';
import './index-C3Ej0YBE.mjs';
import './index-B6Gs2AMF.mjs';
import './index-l8XG4jwb.mjs';
import './nuxt-link-1RlcU5D7.mjs';

const limit = 12;
const _sfc_main = {
  __name: "index",
  __ssrInlineRender: true,
  setup(__props) {
    const route = useRoute();
    const router = useRouter();
    const search = ref(route.query.search || "");
    const minPrice = ref(route.query.min_price || "");
    const maxPrice = ref(route.query.max_price || "");
    const selectedCategory = ref(route.query.category || "");
    const currentPage = ref(Number(route.query.page) || 1);
    const showSuggestions = ref(false);
    const searchSuggestions = ref([]);
    const recentSearches = ref([]);
    ref(0);
    const debouncedSearch = ref("");
    let searchTimeout = null;
    const saveRecentSearch = (searchTerm) => {
      if (!searchTerm || searchTerm.length < 2) return;
      const updated = [searchTerm, ...recentSearches.value.filter((s) => s !== searchTerm)].slice(0, 5);
      recentSearches.value = updated;
    };
    const clearRecentSearches = () => {
      recentSearches.value = [];
    };
    const { products: productsApi, categories: categoriesApi } = useApi();
    const fetchSuggestions = async (query) => {
      if (!query || query.length < 2) {
        searchSuggestions.value = [];
        return;
      }
      try {
        const data = await productsApi.list({ search: query, limit: 5 });
        searchSuggestions.value = data.data || [];
      } catch (e) {
        searchSuggestions.value = [];
      }
    };
    const handleSearchInput = (e) => {
      const value = e.target.value;
      search.value = value;
      if (searchTimeout) clearTimeout(searchTimeout);
      searchTimeout = setTimeout(() => {
        fetchSuggestions(value);
      }, 300);
    };
    watch(() => route.query, (newQuery) => {
      search.value = newQuery.search || "";
      minPrice.value = newQuery.min_price || "";
      maxPrice.value = newQuery.max_price || "";
      selectedCategory.value = newQuery.category || "";
      currentPage.value = Number(newQuery.page) || 1;
    }, { immediate: true, deep: true });
    const productsData = ref(null);
    const productsLoading = ref(false);
    const productsError = ref(null);
    const loadProducts = async () => {
      productsLoading.value = true;
      productsError.value = null;
      try {
        const response = await productsApi.list({
          page: currentPage.value > 1 ? currentPage.value : void 0,
          limit,
          category: selectedCategory.value,
          search: debouncedSearch.value,
          min_price: minPrice.value,
          max_price: maxPrice.value
        });
        productsData.value = response;
      } catch (e) {
        productsError.value = e;
      } finally {
        productsLoading.value = false;
      }
    };
    const refreshProducts = () => {
      loadProducts();
    };
    watch([selectedCategory, minPrice, maxPrice, debouncedSearch, currentPage], () => {
      loadProducts();
    });
    const { data: categoriesData } = useAsyncData("categories-all", () => categoriesApi.list());
    const products = computed(() => {
      var _a;
      return ((_a = productsData.value) == null ? void 0 : _a.data) || productsData.value || [];
    });
    const pagination = computed(() => {
      var _a;
      return ((_a = productsData.value) == null ? void 0 : _a.pagination) || { page: 1, total_pages: 1, total_items: 0 };
    });
    const categories = computed(() => {
      var _a;
      return ((_a = categoriesData.value) == null ? void 0 : _a.data) || categoriesData.value || [];
    });
    const applyFilters = () => {
      if (search.value) {
        saveRecentSearch(search.value);
      }
      router.push({
        query: {
          page: 1,
          ...search.value && { search: search.value },
          ...minPrice.value && { min_price: minPrice.value },
          ...maxPrice.value && { max_price: maxPrice.value },
          ...selectedCategory.value && { category: selectedCategory.value }
        }
      });
    };
    const clearFilters = () => {
      search.value = "";
      minPrice.value = "";
      maxPrice.value = "";
      selectedCategory.value = "";
      debouncedSearch.value = "";
      router.push({ query: { page: 1 } });
    };
    const goToPage = (newPage) => {
      if (newPage >= 1 && newPage <= pagination.value.total_pages) {
        router.push({
          query: {
            ...route.query,
            page: newPage
          }
        });
      }
    };
    watch(search, (newVal) => {
      if (searchTimeout) clearTimeout(searchTimeout);
      searchTimeout = setTimeout(() => {
        debouncedSearch.value = newVal;
        applyFilters();
      }, 500);
    });
    watch([minPrice, maxPrice], () => {
      applyFilters();
    });
    watch(selectedCategory, () => {
      applyFilters();
    });
    useSeoMeta({
      title: "Products - E-Commerce Store",
      description: "Browse our collection of products"
    });
    return (_ctx, _push, _parent, _attrs) => {
      const _component_Button = script;
      const _component_IconField = script$1;
      const _component_InputIcon = script$2;
      const _component_InputText = script$3;
      const _component_Select = script$4;
      const _component_InputNumber = script$5;
      const _component_Paginator = script$6;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "min-h-screen bg-surface py-8" }, _attrs))} data-v-912af6d4><div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8" data-v-912af6d4><div class="mb-8" data-v-912af6d4><h1 class="text-3xl font-bold text-on_surface font-display" data-v-912af6d4>Our Collection</h1><p class="text-on_surface_variant mt-2" data-v-912af6d4>Browse our curated selection of products</p></div><div class="flex flex-col lg:flex-row gap-8" data-v-912af6d4><aside class="w-full lg:w-72 flex-shrink-0" data-v-912af6d4><div class="bg-surface-container-lowest/80 backdrop-blur-sm rounded-xl p-6 sticky top-24 shadow-ambient transition-all duration-300 hover:shadow-lg border border-white/5" data-v-912af6d4><div class="flex items-center justify-between mb-6" data-v-912af6d4><h3 class="font-semibold text-on_surface" data-v-912af6d4>Filters</h3>`);
      _push(ssrRenderComponent(_component_Button, {
        onClick: clearFilters,
        label: "Clear all",
        text: "",
        size: "small",
        class: "!text-sm"
      }, null, _parent));
      _push(`</div><div class="flex flex-col gap-4" data-v-912af6d4><div data-v-912af6d4><label class="block text-sm font-medium text-on_surface_variant mb-2" data-v-912af6d4>Search</label><div class="relative" data-v-912af6d4>`);
      _push(ssrRenderComponent(_component_IconField, null, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(ssrRenderComponent(_component_InputIcon, { class: "pi pi-search" }, null, _parent2, _scopeId));
            _push2(ssrRenderComponent(_component_InputText, {
              modelValue: unref(search),
              "onUpdate:modelValue": ($event) => isRef(search) ? search.value = $event : null,
              placeholder: "Search products...",
              class: "w-full",
              onInput: handleSearchInput,
              onFocus: ($event) => showSuggestions.value = true,
              onBlur: ($event) => _ctx.setTimeout(() => showSuggestions.value = false, 200),
              onKeyup: applyFilters
            }, null, _parent2, _scopeId));
          } else {
            return [
              createVNode(_component_InputIcon, { class: "pi pi-search" }),
              createVNode(_component_InputText, {
                modelValue: unref(search),
                "onUpdate:modelValue": ($event) => isRef(search) ? search.value = $event : null,
                placeholder: "Search products...",
                class: "w-full",
                onInput: handleSearchInput,
                onFocus: ($event) => showSuggestions.value = true,
                onBlur: ($event) => _ctx.setTimeout(() => showSuggestions.value = false, 200),
                onKeyup: withKeys(applyFilters, ["enter"])
              }, null, 8, ["modelValue", "onUpdate:modelValue", "onFocus", "onBlur"])
            ];
          }
        }),
        _: 1
      }, _parent));
      if (unref(showSuggestions) && (unref(searchSuggestions).length > 0 || unref(recentSearches).length > 0)) {
        _push(`<div class="absolute z-50 w-full mt-2 bg-surface-container-lowest rounded-lg shadow-xl border border-outline-variant/20 overflow-hidden" data-v-912af6d4>`);
        if (unref(searchSuggestions).length > 0) {
          _push(`<div class="py-2" data-v-912af6d4><p class="px-3 py-1 text-xs font-semibold text-on_surface_variant uppercase" data-v-912af6d4>Products</p><!--[-->`);
          ssrRenderList(unref(searchSuggestions), (product) => {
            _push(`<div class="w-full px-3 py-2 text-left hover:bg-surface-container flex items-center gap-3 transition-colors cursor-pointer" data-v-912af6d4><i class="pi pi-box text-outline flex-shrink-0" data-v-912af6d4></i><div class="flex-1 min-w-0" data-v-912af6d4><p class="text-on_surface truncate" data-v-912af6d4>${ssrInterpolate(product.name)}</p><p class="text-xs text-primary font-medium" data-v-912af6d4>\u20B9${ssrInterpolate(Number(product.price).toFixed(2))}</p></div></div>`);
          });
          _push(`<!--]--></div>`);
        } else {
          _push(`<!---->`);
        }
        if (unref(recentSearches).length > 0) {
          _push(`<div class="border-t border-outline-variant/20 py-2" data-v-912af6d4><div class="flex items-center justify-between px-3 py-1" data-v-912af6d4><p class="text-xs font-semibold text-on_surface_variant uppercase" data-v-912af6d4>Recent Searches</p>`);
          _push(ssrRenderComponent(_component_Button, {
            onMousedown: clearRecentSearches,
            label: "Clear",
            text: "",
            size: "small",
            class: "!text-xs"
          }, null, _parent));
          _push(`</div><!--[-->`);
          ssrRenderList(unref(recentSearches), (term) => {
            _push(`<div class="w-full px-3 py-2 text-left hover:bg-surface-container flex items-center gap-3 transition-colors cursor-pointer" data-v-912af6d4><i class="pi pi-clock text-outline text-xs" data-v-912af6d4></i><span class="text-on_surface" data-v-912af6d4>${ssrInterpolate(term)}</span></div>`);
          });
          _push(`<!--]--></div>`);
        } else {
          _push(`<!---->`);
        }
        _push(`</div>`);
      } else {
        _push(`<!---->`);
      }
      _push(`</div></div><div data-v-912af6d4><label class="block text-sm font-medium text-on_surface_variant mb-2" data-v-912af6d4>Category</label>`);
      _push(ssrRenderComponent(_component_Select, {
        modelValue: unref(selectedCategory),
        "onUpdate:modelValue": ($event) => isRef(selectedCategory) ? selectedCategory.value = $event : null,
        options: unref(categories),
        optionLabel: "name",
        optionValue: "category_id",
        placeholder: "All Categories",
        class: "w-full",
        onChange: applyFilters
      }, null, _parent));
      _push(`</div><div data-v-912af6d4><label class="block text-sm font-medium text-on_surface_variant mb-2" data-v-912af6d4>Price Range</label><div style="${ssrRenderStyle({ "display": "grid", "grid-template-columns": "1fr 1fr", "gap": "0.5rem" })}" data-v-912af6d4>`);
      _push(ssrRenderComponent(_component_InputNumber, {
        modelValue: unref(minPrice),
        "onUpdate:modelValue": ($event) => isRef(minPrice) ? minPrice.value = $event : null,
        placeholder: "Min",
        style: { "width": "100%", "min-width": "0" },
        "input-style": "width: 100%; box-sizing: border-box;",
        min: 0
      }, null, _parent));
      _push(ssrRenderComponent(_component_InputNumber, {
        modelValue: unref(maxPrice),
        "onUpdate:modelValue": ($event) => isRef(maxPrice) ? maxPrice.value = $event : null,
        placeholder: "Max",
        style: { "width": "100%", "min-width": "0" },
        "input-style": "width: 100%; box-sizing: border-box;",
        min: 0
      }, null, _parent));
      _push(`</div></div>`);
      _push(ssrRenderComponent(_component_Button, {
        onClick: applyFilters,
        label: "Apply Filters",
        class: "w-full"
      }, null, _parent));
      _push(`</div></div></aside><main class="flex-1" data-v-912af6d4>`);
      if (unref(productsLoading)) {
        _push(`<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6" data-v-912af6d4><!--[-->`);
        ssrRenderList(6, (i) => {
          _push(`<div class="bg-surface-container-lowest rounded-xl p-4 animate-pulse" data-v-912af6d4><div class="aspect-[4/3] bg-surface-container rounded-lg mb-4" data-v-912af6d4></div><div class="h-4 bg-surface-container rounded w-3/4 mb-2" data-v-912af6d4></div><div class="h-4 bg-surface-container rounded w-1/2" data-v-912af6d4></div></div>`);
        });
        _push(`<!--]--></div>`);
      } else if (unref(productsError)) {
        _push(`<div class="bg-surface-container-lowest rounded-xl shadow-ambient p-8 text-center" data-v-912af6d4><i class="pi pi-exclamation-triangle text-5xl text-error mb-4" data-v-912af6d4></i><h3 class="mt-4 text-lg font-medium text-on_surface" data-v-912af6d4>Failed to load products</h3><p class="mt-2 text-on_surface_variant" data-v-912af6d4>Something went wrong. Please try again.</p>`);
        _push(ssrRenderComponent(_component_Button, {
          onClick: refreshProducts,
          label: "Try Again",
          class: "mt-4"
        }, null, _parent));
        _push(`</div>`);
      } else {
        _push(`<!--[-->`);
        if (unref(products).length === 0) {
          _push(`<div class="bg-surface-container-lowest rounded-xl shadow-ambient p-12 text-center" data-v-912af6d4><i class="pi pi-box text-6xl text-outline mb-4" data-v-912af6d4></i><h2 class="mt-4 text-xl font-medium text-on_surface" data-v-912af6d4>No products found</h2><p class="mt-2 text-on_surface_variant" data-v-912af6d4>Try adjusting your filters or search terms.</p>`);
          _push(ssrRenderComponent(_component_Button, {
            onClick: clearFilters,
            label: "Clear Filters",
            class: "mt-6"
          }, null, _parent));
          _push(`</div>`);
        } else {
          _push(`<div data-v-912af6d4><p class="text-on_surface_variant mb-4" data-v-912af6d4>${ssrInterpolate(unref(pagination).total_items)} products found</p><div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6" data-v-912af6d4><!--[-->`);
          ssrRenderList(unref(products), (product, index2) => {
            _push(ssrRenderComponent(ProductCard, {
              key: product.product_id,
              product,
              class: "animate-fade-in-up",
              style: { animationDelay: `${index2 * 50}ms` }
            }, null, _parent));
          });
          _push(`<!--]--></div>`);
          if (unref(pagination).total_pages > 1) {
            _push(`<div class="mt-12 flex justify-center gap-2" data-v-912af6d4>`);
            _push(ssrRenderComponent(_component_Paginator, {
              first: (unref(currentPage) - 1) * limit,
              rows: limit,
              totalRecords: unref(pagination).total_items,
              onPage: ($event) => goToPage($event.page + 1),
              class: "mt-6"
            }, null, _parent));
            _push(`</div>`);
          } else {
            _push(`<!---->`);
          }
          _push(`</div>`);
        }
        _push(`<!--]-->`);
      }
      _push(`</main></div></div></div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/products/index.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const index = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-912af6d4"]]);

export { index as default };
//# sourceMappingURL=index-CfC8K2LC.mjs.map
