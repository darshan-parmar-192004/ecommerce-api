import { _ as __nuxt_component_0 } from './nuxt-link-1RlcU5D7.mjs';
import { withAsyncContext, mergeProps, withCtx, createVNode, createTextVNode, unref, useSSRContext } from 'vue';
import { u as useRoute, a as useSeoMeta, s as script } from './server.mjs';
import { u as useApi } from './useApi-CR6WE1xi.mjs';
import { u as useAsyncData } from './asyncData--maOkXya.mjs';
import script$1 from './index-BGLbqIM8.mjs';
import { ssrRenderAttrs, ssrRenderComponent, ssrInterpolate, ssrRenderList } from 'vue/server-renderer';
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

const _sfc_main = {
  __name: "[id]",
  __ssrInlineRender: true,
  async setup(__props) {
    let __temp, __restore;
    const route = useRoute();
    const { orders: ordersApi } = useApi();
    const { data: order, pending, error, refresh } = ([__temp, __restore] = withAsyncContext(() => useAsyncData(
      `order-${route.params.id}`,
      () => ordersApi.get(route.params.id)
    )), __temp = await __temp, __restore(), __temp);
    useSeoMeta({
      title: () => order.value ? `Order #${order.value.order_id || order.value.id} - E-Commerce Store` : "Order Details"
    });
    const formatDate = (dateStr) => {
      if (!dateStr) return "N/A";
      return new Date(dateStr).toLocaleDateString("en-US", { year: "numeric", month: "long", day: "numeric" });
    };
    return (_ctx, _push, _parent, _attrs) => {
      const _component_NuxtLink = __nuxt_component_0;
      const _component_Button = script;
      const _component_Tag = script$1;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8" }, _attrs))}>`);
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: "/orders",
        class: "inline-flex items-center text-sm text-on_surface_variant hover:text-primary mb-6"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`<i class="pi pi-arrow-left mr-1"${_scopeId}></i> Back to Orders `);
          } else {
            return [
              createVNode("i", { class: "pi pi-arrow-left mr-1" }),
              createTextVNode(" Back to Orders ")
            ];
          }
        }),
        _: 1
      }, _parent));
      if (unref(pending)) {
        _push(`<div class="animate-pulse space-y-4"><div class="h-8 bg-surface-container rounded w-1/3"></div><div class="h-4 bg-surface-container rounded w-1/2"></div></div>`);
      } else if (unref(error)) {
        _push(`<div class="text-center py-12"><i class="pi pi-exclamation-triangle text-5xl text-error mb-4"></i><h2 class="text-xl font-semibold text-on_surface mb-2">Failed to load order</h2>`);
        _push(ssrRenderComponent(_component_Button, {
          onClick: unref(refresh),
          label: "Try Again",
          icon: "pi pi-refresh"
        }, null, _parent));
        _push(`</div>`);
      } else if (unref(order)) {
        _push(`<div class="space-y-6"><div class="flex items-center justify-between"><div><h1 class="text-2xl font-bold text-on_surface font-display">Order #${ssrInterpolate(unref(order).order_id || unref(order).id)}</h1><p class="text-sm text-on_surface_variant mt-1">Placed on ${ssrInterpolate(formatDate(unref(order).created_at || unref(order).order_date))}</p></div>`);
        _push(ssrRenderComponent(_component_Tag, {
          value: unref(order).status || "Pending",
          severity: unref(order).status === "delivered" ? "success" : unref(order).status === "cancelled" ? "danger" : "warn"
        }, null, _parent));
        _push(`</div><div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6"><h2 class="text-lg font-semibold text-on_surface mb-4">Shipping Address</h2><p class="text-on_surface_variant">${ssrInterpolate(unref(order).shipping_address)}</p></div><div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6"><h2 class="text-lg font-semibold text-on_surface mb-4">Order Items</h2><div class="space-y-4"><!--[-->`);
        ssrRenderList(unref(order).items || unref(order).order_items || [], (item) => {
          _push(`<div class="flex items-center gap-4 py-3 border-b border-outline-variant/20 last:border-0"><div class="w-16 h-16 bg-surface-container rounded-lg flex items-center justify-center"><i class="pi pi-box text-2xl text-outline"></i></div><div class="flex-1"><p class="font-medium text-on_surface">${ssrInterpolate(item.product_name || item.name)}</p><p class="text-sm text-on_surface_variant">Qty: ${ssrInterpolate(item.quantity)} \xD7 \u20B9 ${ssrInterpolate(Number(item.unit_price || item.price).toFixed(2))}</p></div><p class="font-semibold text-on_surface">\u20B9 ${ssrInterpolate((Number(item.unit_price || item.price) * item.quantity).toFixed(2))}</p></div>`);
        });
        _push(`<!--]--></div><div class="mt-6 pt-4 border-t border-outline-variant/20"><div class="flex justify-between text-lg font-bold text-on_surface"><span>Total</span><span class="text-primary">\u20B9 ${ssrInterpolate(Number(unref(order).total_amount || unref(order).total || 0).toFixed(2))}</span></div></div></div></div>`);
      } else {
        _push(`<!---->`);
      }
      _push(`</div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/orders/[id].vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};

export { _sfc_main as default };
//# sourceMappingURL=_id_-RMYz5bGn.mjs.map
