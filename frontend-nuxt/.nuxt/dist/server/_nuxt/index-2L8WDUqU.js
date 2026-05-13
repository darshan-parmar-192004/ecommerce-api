import { _ as __nuxt_component_0 } from "./nuxt-link-1RlcU5D7.js";
import { withAsyncContext, computed, ref, unref, withCtx, createVNode, createTextVNode, useSSRContext } from "vue";
import "hookable";
import { u as useApi } from "./useApi-CR6WE1xi.js";
import { u as useAsyncData } from "./asyncData--maOkXya.js";
import { a as useSeoMeta, s as script } from "../server.mjs";
import script$1 from "./index-DDXXjm8b.js";
import { ssrRenderAttrs, ssrRenderStyle, ssrRenderList, ssrInterpolate, ssrRenderComponent } from "vue/server-renderer";
import { P as ProductCard } from "./ProductCard-Di5Z6QG2.js";
import { _ as _export_sfc } from "./_plugin-vue_export-helper-1tPrXgE0.js";
import "ufo";
import "ofetch";
import "#internal/nuxt/paths";
import "unctx";
import "h3";
import "unhead";
import "@unhead/shared";
import "vue-router";
import "radix3";
import "defu";
import "klona";
import "@primeuix/utils/eventbus";
import "@primeuix/styled";
import "@primeuix/utils";
import "@primeuix/utils/object";
import "@primeuix/styles/base";
import "@primeuix/utils/dom";
import "cookie-es";
import "destr";
import "ohash";
import "@primeuix/utils/zindex";
import "@primeuix/styles/toast";
import "@primeuix/utils/uuid";
import "@primeuix/styles/ripple";
import "@primeuix/styles/badge";
import "@primeuix/styles/button";
import "./index-Dedjwyps.js";
import "./index-DrzwqwtL.js";
import "@primeuix/styles/inputtext";
const _sfc_main = {
  __name: "index",
  __ssrInlineRender: true,
  async setup(__props) {
    let __temp, __restore;
    const { products: productsApi } = useApi();
    const { data, pending, error, refresh } = ([__temp, __restore] = withAsyncContext(() => useAsyncData(
      "home-products",
      () => productsApi.list({ page: 1, limit: 8 })
    )), __temp = await __temp, __restore(), __temp);
    useSeoMeta({
      title: "Home - E-Commerce Store",
      description: "Shop the best products at our store"
    });
    const products = computed(() => {
      var _a;
      return ((_a = data.value) == null ? void 0 : _a.data) || [];
    });
    const parallaxOffset = ref(0);
    const floatingProducts = computed(() => products.value.slice(0, 3));
    return (_ctx, _push, _parent, _attrs) => {
      const _component_NuxtLink = __nuxt_component_0;
      const _component_Button = script;
      const _component_InputText = script$1;
      _push(`<div${ssrRenderAttrs(_attrs)} data-v-0cfd007b><section class="relative overflow-hidden bg-surface-container-low min-h-[85vh] flex items-center" data-v-0cfd007b><div class="absolute inset-0 overflow-hidden pointer-events-none" style="${ssrRenderStyle({ transform: `translateY(${unref(parallaxOffset)}px)` })}" data-v-0cfd007b><div class="absolute -top-20 -right-20 w-[600px] h-[600px] rounded-full blur-3xl opacity-20" style="${ssrRenderStyle({ "background": "radial-gradient(circle, #3e51fb, transparent)", "animation": "float 8s ease-in-out infinite" })}" data-v-0cfd007b></div><div class="absolute -bottom-40 -left-40 w-[500px] h-[500px] rounded-full blur-3xl opacity-15" style="${ssrRenderStyle({ "background": "radial-gradient(circle, #1c31e3, transparent)", "animation": "float 10s ease-in-out infinite reverse" })}" data-v-0cfd007b></div><div class="absolute top-1/2 left-1/2 -translate-x-1/2 w-[800px] h-[800px] rounded-full blur-3xl opacity-10" style="${ssrRenderStyle({ "background": "radial-gradient(circle, #6366f1, transparent)", "animation": "pulse-glow 6s ease-in-out infinite" })}" data-v-0cfd007b></div></div><div class="absolute inset-0 overflow-hidden pointer-events-none" data-v-0cfd007b><!--[-->`);
      ssrRenderList(unref(floatingProducts), (product, index2) => {
        _push(`<div class="absolute hidden lg:block w-40 h-48 bg-surface-container-lowest rounded-xl shadow-2xl border border-outline-variant/20 overflow-hidden animate-float-delayed" style="${ssrRenderStyle({
          top: `${15 + index2 * 25}%`,
          right: `${12 + index2 * 6}%`,
          animationDelay: `${index2 * 2}s`,
          transform: `translateY(${unref(parallaxOffset) * 0.5}px)`
        })}" data-v-0cfd007b><div class="aspect-[4/3] bg-surface-container flex items-center justify-center" data-v-0cfd007b><i class="pi pi-box text-4xl text-outline/30" data-v-0cfd007b></i></div><div class="p-2 text-center" data-v-0cfd007b><p class="text-xs font-medium text-on_surface truncate" data-v-0cfd007b>${ssrInterpolate(product.name)}</p><p class="text-xs text-primary font-bold" data-v-0cfd007b>₹${ssrInterpolate(Number(product.price).toFixed(0))}</p></div></div>`);
      });
      _push(`<!--]--></div><div class="relative w-full max-w-7xl mx-auto px-6 lg:px-8 py-28 md:py-40" data-v-0cfd007b><div class="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center" data-v-0cfd007b><div class="max-w-xl" data-v-0cfd007b><p class="text-xs font-semibold uppercase tracking-[0.2em] text-outline mb-6 animate-fade-in-up" data-v-0cfd007b> Curated Collection </p><h1 class="text-5xl md:text-6xl lg:text-7xl font-bold leading-[1.05] tracking-tight text-on_surface mb-8 animate-fade-in-up font-display" style="${ssrRenderStyle({ "animation-delay": "100ms", "letter-spacing": "-0.02em" })}" data-v-0cfd007b> Discover Products<br data-v-0cfd007b><span class="text-primary" data-v-0cfd007b>That Matter</span></h1><p class="text-lg md:text-xl text-on_surface_variant mb-10 max-w-xl animate-fade-in-up font-body leading-relaxed" style="${ssrRenderStyle({ "animation-delay": "200ms" })}" data-v-0cfd007b> Shop the latest trends with unbeatable prices and fast delivery </p><div class="flex flex-wrap gap-4 animate-fade-in-up" style="${ssrRenderStyle({ "animation-delay": "300ms" })}" data-v-0cfd007b>`);
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: "/products",
        class: "group relative inline-flex items-center gap-2 bg-gradient-to-r from-primary to-primary-container text-white px-8 py-4 rounded-md font-semibold transition-all duration-400 hover:shadow-xl hover:shadow-primary/30 hover:scale-105 active:scale-[0.98]"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`<span class="relative z-10" data-v-0cfd007b${_scopeId}>Shop Now</span><i class="pi pi-arrow-right relative z-10 transition-transform group-hover:translate-x-1" data-v-0cfd007b${_scopeId}></i><span class="absolute inset-0 rounded-md bg-white/20 scale-0 group-hover:scale-100 transition-transform duration-300" data-v-0cfd007b${_scopeId}></span>`);
          } else {
            return [
              createVNode("span", { class: "relative z-10" }, "Shop Now"),
              createVNode("i", { class: "pi pi-arrow-right relative z-10 transition-transform group-hover:translate-x-1" }),
              createVNode("span", { class: "absolute inset-0 rounded-md bg-white/20 scale-0 group-hover:scale-100 transition-transform duration-300" })
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: "/cart",
        class: "inline-flex items-center gap-2 px-8 py-4 rounded-md font-semibold transition-all duration-300 hover:bg-surface-container-high bg-surface-container-highest text-on_surface hover:scale-105 active:scale-[0.98]"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(` View Cart `);
          } else {
            return [
              createTextVNode(" View Cart ")
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`</div></div><div class="hidden lg:block" data-v-0cfd007b></div></div></div></section><section class="max-w-7xl mx-auto px-6 lg:px-8 py-20" data-v-0cfd007b><div class="flex items-end justify-between mb-12" data-v-0cfd007b><div data-v-0cfd007b><h2 class="text-3xl font-bold text-on_surface font-display tracking-tight" data-v-0cfd007b>Featured Products</h2><p class="text-on_surface_variant mt-2 font-body" data-v-0cfd007b>Handpicked for you</p></div>`);
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: "/products",
        class: "text-primary hover:text-primary-container font-semibold flex items-center gap-1 transition-colors duration-300 text-sm uppercase tracking-wider"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(` View All <i class="pi pi-arrow-right text-xs" data-v-0cfd007b${_scopeId}></i>`);
          } else {
            return [
              createTextVNode(" View All "),
              createVNode("i", { class: "pi pi-arrow-right text-xs" })
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`</div>`);
      if (unref(pending)) {
        _push(`<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8" data-v-0cfd007b><!--[-->`);
        ssrRenderList(4, (i) => {
          _push(`<div class="rounded-md overflow-hidden" style="${ssrRenderStyle({ "border": "1px solid rgba(197, 197, 217, 0.2)" })}" data-v-0cfd007b><div class="aspect-[4/3] bg-surface-container-low relative overflow-hidden" data-v-0cfd007b><div class="absolute inset-0 animate-shimmer" style="${ssrRenderStyle({ "background": "linear-gradient(90deg, transparent, rgba(255,255,255,0.4), transparent)", "background-size": "200% 100%" })}" data-v-0cfd007b></div></div><div class="p-5 bg-surface-container-lowest" data-v-0cfd007b><div class="h-5 bg-surface-container rounded w-3/4 mb-3 animate-shimmer" style="${ssrRenderStyle({ "background": "linear-gradient(90deg, #eeeef0, #e8e8ea, #eeeef0)", "background-size": "200% 100%", "animation": "shimmer 1.5s infinite" })}" data-v-0cfd007b></div><div class="h-4 bg-surface-container rounded w-1/2 mb-5 animate-shimmer" style="${ssrRenderStyle({ "background": "linear-gradient(90deg, #eeeef0, #e8e8ea, #eeeef0)", "background-size": "200% 100%", "animation": "shimmer 1.5s infinite" })}" data-v-0cfd007b></div><div class="flex justify-between items-center" data-v-0cfd007b><div class="h-6 bg-surface-container rounded w-20 animate-shimmer" style="${ssrRenderStyle({ "background": "linear-gradient(90deg, #eeeef0, #e8e8ea, #eeeef0)", "background-size": "200% 100%", "animation": "shimmer 1.5s infinite" })}" data-v-0cfd007b></div><div class="h-10 bg-surface-container rounded-md w-20 animate-shimmer" style="${ssrRenderStyle({ "background": "linear-gradient(90deg, #eeeef0, #e8e8ea, #eeeef0)", "background-size": "200% 100%", "animation": "shimmer 1.5s infinite" })}" data-v-0cfd007b></div></div></div></div>`);
        });
        _push(`<!--]--></div>`);
      } else if (unref(error)) {
        _push(`<div class="bg-surface-container-lowest rounded-md p-12 text-center shadow-ambient" style="${ssrRenderStyle({ "border": "1px solid rgba(197, 197, 217, 0.2)" })}" data-v-0cfd007b><i class="pi pi-exclamation-triangle text-5xl text-error mb-4" data-v-0cfd007b></i><h3 class="mt-4 text-lg font-semibold text-on_surface font-display" data-v-0cfd007b>Failed to load products</h3><p class="mt-2 text-on_surface_variant font-body" data-v-0cfd007b>Something went wrong. Please try again.</p>`);
        _push(ssrRenderComponent(_component_Button, {
          onClick: unref(refresh),
          label: "Try Again",
          class: "mt-6"
        }, null, _parent));
        _push(`</div>`);
      } else if (unref(products).length === 0) {
        _push(`<div class="bg-surface-container-lowest rounded-md p-16 text-center shadow-ambient" style="${ssrRenderStyle({ "border": "1px solid rgba(197, 197, 217, 0.2)" })}" data-v-0cfd007b><i class="pi pi-box text-5xl text-outline/30 mb-4" data-v-0cfd007b></i><h2 class="mt-4 text-xl font-semibold text-on_surface font-display" data-v-0cfd007b>No products available</h2><p class="mt-2 text-on_surface_variant font-body" data-v-0cfd007b>Check back soon for new products.</p></div>`);
      } else {
        _push(`<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8" data-v-0cfd007b><!--[-->`);
        ssrRenderList(unref(products), (product, index2) => {
          _push(ssrRenderComponent(ProductCard, {
            key: product.product_id,
            product,
            class: "animate-stagger-fade",
            style: { animationDelay: `${index2 * 80}ms` }
          }, null, _parent));
        });
        _push(`<!--]--></div>`);
      }
      _push(`</section><section class="bg-surface-container-low py-20" data-v-0cfd007b><div class="max-w-7xl mx-auto px-6 lg:px-8" data-v-0cfd007b><div class="grid grid-cols-1 md:grid-cols-3 gap-12" data-v-0cfd007b><div class="text-center group" data-v-0cfd007b><div class="w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-5 transition-transform duration-400 group-hover:scale-110 bg-gradient-to-r from-primary to-primary-container" data-v-0cfd007b><i class="pi pi-truck text-white text-2xl" data-v-0cfd007b></i></div><h3 class="text-lg font-semibold text-on_surface mb-2 font-display" data-v-0cfd007b>Free Shipping</h3><p class="text-on_surface_variant font-body" data-v-0cfd007b>On orders over ₹500</p></div><div class="text-center group" data-v-0cfd007b><div class="w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-5 transition-transform duration-400 group-hover:scale-110 bg-gradient-to-r from-primary to-primary-container" data-v-0cfd007b><i class="pi pi-shield text-white text-2xl" data-v-0cfd007b></i></div><h3 class="text-lg font-semibold text-on_surface mb-2 font-display" data-v-0cfd007b>Secure Payment</h3><p class="text-on_surface_variant font-body" data-v-0cfd007b>100% secure checkout</p></div><div class="text-center group" data-v-0cfd007b><div class="w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-5 transition-transform duration-400 group-hover:scale-110 bg-gradient-to-r from-primary to-primary-container" data-v-0cfd007b><i class="pi pi-refresh text-white text-2xl" data-v-0cfd007b></i></div><h3 class="text-lg font-semibold text-on_surface mb-2 font-display" data-v-0cfd007b>Easy Returns</h3><p class="text-on_surface_variant font-body" data-v-0cfd007b>30-day return policy</p></div></div></div></section><section class="py-20 bg-surface" data-v-0cfd007b><div class="max-w-7xl mx-auto px-6 lg:px-8" data-v-0cfd007b><div class="bg-gradient-to-r from-primary to-primary-container rounded-2xl p-12 text-center relative overflow-hidden" data-v-0cfd007b><div class="absolute inset-0 bg-[url(&#39;data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNjAiIGhlaWdodD0iNjAiIHZpZXdCb3g9IjAgMCA2MCA2MCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48ZyBmaWxsPSJub25lIiBmaWxsLXJ1bGU9ImV2ZW5vZGQiPjxwYXRoIGQ9Ik0zNiAxOGMtOS45NDEgMC0xOCA4LjA1OS0xOCAxOHM4LjA1OSAxOCAxOCAxOCAxOC04LjA1OSAxOC0xOC04LjA1OS0xOC0xOC0xOHptMCAzMmMtMy4zMTIgMC02LTIuNjg4LTYtNnMyLjY4OC02IDYtNiA2IDIuNjg4IDYgNi0yLjY4OCA2LTYgNnoiIGZpbGw9IiNmZmZmZmYiIGZpbGwtb3BhY2l0eT0iLjEiLz48L2c+PC9zdmc+&#39;)] opacity-20" data-v-0cfd007b></div><div class="relative z-10" data-v-0cfd007b><h2 class="text-3xl font-bold text-white font-display mb-4" data-v-0cfd007b>Stay Updated</h2><p class="text-white/80 mb-8 max-w-xl mx-auto font-body" data-v-0cfd007b>Subscribe to our newsletter for exclusive deals, new arrivals, and insider-only discounts.</p><div class="flex flex-col sm:flex-row gap-4 max-w-md mx-auto" data-v-0cfd007b>`);
      _push(ssrRenderComponent(_component_InputText, {
        type: "email",
        placeholder: "Enter your email",
        class: "flex-1 px-6 py-4 rounded-lg bg-white/20 backdrop-blur-sm text-white placeholder:text-white/60 border border-white/30 focus:outline-none focus:bg-white/30 focus:border-white"
      }, null, _parent));
      _push(ssrRenderComponent(_component_Button, {
        type: "submit",
        label: "Subscribe",
        class: "px-8 py-4 !bg-white !text-primary font-semibold rounded-lg transition-all duration-300 hover:bg-white/90 hover:scale-[1.02] active:scale-[0.98] white"
      }, null, _parent));
      _push(`</div></div></div></div></section></div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/index.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const index = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-0cfd007b"]]);
export {
  index as default
};
//# sourceMappingURL=index-2L8WDUqU.js.map
