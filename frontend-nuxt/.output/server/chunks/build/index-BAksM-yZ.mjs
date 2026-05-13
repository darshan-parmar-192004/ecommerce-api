import { _ as __nuxt_component_0 } from './nuxt-link-1RlcU5D7.mjs';
import { withAsyncContext, computed, mergeProps, unref, withCtx, createVNode, toDisplayString, openBlock, createBlock, createCommentVNode, useSSRContext } from 'vue';
import { a as useSeoMeta, s as script } from './server.mjs';
import { u as useApi } from './useApi-CR6WE1xi.mjs';
import { u as useAsyncData } from './asyncData--maOkXya.mjs';
import script$1 from './index-BGLbqIM8.mjs';
import { ssrRenderAttrs, ssrRenderList, ssrInterpolate, ssrRenderComponent } from 'vue/server-renderer';
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
  __name: "index",
  __ssrInlineRender: true,
  async setup(__props) {
    let __temp, __restore;
    useSeoMeta({
      title: "My Orders - E-Commerce Store"
    });
    const { orders: ordersApi } = useApi();
    const { data: ordersData, pending, error, refresh } = ([__temp, __restore] = withAsyncContext(() => useAsyncData(
      "orders",
      () => ordersApi.list()
    )), __temp = await __temp, __restore(), __temp);
    const orders = computed(() => {
      var _a;
      return ((_a = ordersData.value) == null ? void 0 : _a.data) || [];
    });
    const formatDate = (dateStr) => {
      if (!dateStr) return "N/A";
      return new Date(dateStr).toLocaleDateString("en-US", {
        year: "numeric",
        month: "long",
        day: "numeric"
      });
    };
    return (_ctx, _push, _parent, _attrs) => {
      const _component_Button = script;
      const _component_NuxtLink = __nuxt_component_0;
      const _component_Tag = script$1;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8" }, _attrs))}><h1 class="text-3xl font-bold text-on_surface mb-8">My Orders</h1>`);
      if (unref(pending)) {
        _push(`<div class="space-y-4"><!--[-->`);
        ssrRenderList(3, (i) => {
          _push(`<div class="bg-surface-container-lowest rounded-lg p-6 animate-pulse"><div class="h-6 bg-surface-container rounded w-1/3 mb-4"></div><div class="h-4 bg-surface-container rounded w-1/2"></div></div>`);
        });
        _push(`<!--]--></div>`);
      } else if (unref(error)) {
        _push(`<div class="text-center py-12"><i class="pi pi-exclamation-triangle text-5xl text-error mb-4"></i><h2 class="text-xl font-semibold text-on_surface mb-2">Failed to load orders</h2><p class="text-on_surface_variant mb-4">${ssrInterpolate(unref(error).message)}</p>`);
        _push(ssrRenderComponent(_component_Button, {
          onClick: unref(refresh),
          label: "Try Again",
          icon: "pi pi-refresh"
        }, null, _parent));
        _push(`</div>`);
      } else if (unref(orders).length === 0) {
        _push(`<div class="text-center py-12"><i class="pi pi-box text-5xl text-outline mb-4"></i><h2 class="text-xl font-semibold text-on_surface mb-2">No orders yet</h2><p class="text-on_surface_variant mb-6">Start shopping to see your orders here.</p>`);
        _push(ssrRenderComponent(_component_NuxtLink, { to: "/products" }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(ssrRenderComponent(_component_Button, {
                label: "Browse Products",
                icon: "pi pi-shopping-cart"
              }, null, _parent2, _scopeId));
            } else {
              return [
                createVNode(_component_Button, {
                  label: "Browse Products",
                  icon: "pi pi-shopping-cart"
                })
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</div>`);
      } else {
        _push(`<div class="space-y-4"><!--[-->`);
        ssrRenderList(unref(orders), (order) => {
          _push(ssrRenderComponent(_component_NuxtLink, {
            key: order.order_id || order.id,
            to: `/orders/${order.order_id || order.id}`,
            class: "block bg-surface-container-lowest rounded-lg p-6 hover:shadow-lg transition-shadow"
          }, {
            default: withCtx((_, _push2, _parent2, _scopeId) => {
              if (_push2) {
                _push2(`<div class="flex items-center justify-between mb-3"${_scopeId}><span class="font-semibold text-on_surface"${_scopeId}> Order #${ssrInterpolate(order.order_id || order.id)}</span>`);
                _push2(ssrRenderComponent(_component_Tag, {
                  value: order.status || "Pending",
                  severity: order.status === "delivered" ? "success" : order.status === "cancelled" ? "danger" : "warn"
                }, null, _parent2, _scopeId));
                _push2(`</div><div class="text-sm text-on_surface_variant space-y-1"${_scopeId}><p${_scopeId}>${ssrInterpolate(formatDate(order.created_at || order.order_date))}</p><p class="font-medium text-primary"${_scopeId}> \u20B9 ${ssrInterpolate(Number(order.total_amount || order.total || 0).toFixed(2))}</p>`);
                if (order.items_count) {
                  _push2(`<p${_scopeId}>${ssrInterpolate(order.items_count)} item(s)</p>`);
                } else {
                  _push2(`<!---->`);
                }
                _push2(`</div>`);
              } else {
                return [
                  createVNode("div", { class: "flex items-center justify-between mb-3" }, [
                    createVNode("span", { class: "font-semibold text-on_surface" }, " Order #" + toDisplayString(order.order_id || order.id), 1),
                    createVNode(_component_Tag, {
                      value: order.status || "Pending",
                      severity: order.status === "delivered" ? "success" : order.status === "cancelled" ? "danger" : "warn"
                    }, null, 8, ["value", "severity"])
                  ]),
                  createVNode("div", { class: "text-sm text-on_surface_variant space-y-1" }, [
                    createVNode("p", null, toDisplayString(formatDate(order.created_at || order.order_date)), 1),
                    createVNode("p", { class: "font-medium text-primary" }, " \u20B9 " + toDisplayString(Number(order.total_amount || order.total || 0).toFixed(2)), 1),
                    order.items_count ? (openBlock(), createBlock("p", { key: 0 }, toDisplayString(order.items_count) + " item(s)", 1)) : createCommentVNode("", true)
                  ])
                ];
              }
            }),
            _: 2
          }, _parent));
        });
        _push(`<!--]--></div>`);
      }
      _push(`</div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/orders/index.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};

export { _sfc_main as default };
//# sourceMappingURL=index-BAksM-yZ.mjs.map
