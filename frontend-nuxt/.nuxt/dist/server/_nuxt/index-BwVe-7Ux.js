import { ref, unref, useSSRContext } from "vue";
import "hookable";
import { ssrRenderAttrs, ssrInterpolate } from "vue/server-renderer";
const _sfc_main = {
  __name: "index",
  __ssrInlineRender: true,
  setup(__props) {
    const stats = ref({ totalOrders: 0, totalProducts: 0, totalCustomers: 0, revenue: 0 });
    const loading = ref(false);
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<div${ssrRenderAttrs(_attrs)}><h1 class="text-2xl font-bold text-on_surface mb-6">Dashboard</h1><div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6"><div class="bg-surface-container-lowest rounded-lg p-6"><p class="text-sm text-on_surface_variant mb-1">Total Orders</p><p class="text-3xl font-bold text-on_surface">${ssrInterpolate(unref(loading) ? "—" : unref(stats).totalOrders)}</p></div><div class="bg-surface-container-lowest rounded-lg p-6"><p class="text-sm text-on_surface_variant mb-1">Total Products</p><p class="text-3xl font-bold text-on_surface">${ssrInterpolate(unref(loading) ? "—" : unref(stats).totalProducts)}</p></div><div class="bg-surface-container-lowest rounded-lg p-6"><p class="text-sm text-on_surface_variant mb-1">Total Customers</p><p class="text-3xl font-bold text-on_surface">${ssrInterpolate(unref(loading) ? "—" : unref(stats).totalCustomers)}</p></div><div class="bg-surface-container-lowest rounded-lg p-6"><p class="text-sm text-on_surface_variant mb-1">Revenue</p><p class="text-3xl font-bold text-on_surface">$${ssrInterpolate(unref(loading) ? "—" : unref(stats).revenue)}</p></div></div></div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/admin/index.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
export {
  _sfc_main as default
};
//# sourceMappingURL=index-BwVe-7Ux.js.map
