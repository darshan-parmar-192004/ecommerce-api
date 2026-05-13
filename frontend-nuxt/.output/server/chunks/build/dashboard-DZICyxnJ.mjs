import { _ as __nuxt_component_0 } from './nuxt-link-1RlcU5D7.mjs';
import { withAsyncContext, computed, mergeProps, unref, withCtx, createTextVNode, toDisplayString, createVNode, useSSRContext } from 'vue';
import { u as useApi } from './useApi-CR6WE1xi.mjs';
import { u as useAsyncData } from './asyncData--maOkXya.mjs';
import { a as useSeoMeta, s as script$3 } from './server.mjs';
import script$2 from './index-BGLbqIM8.mjs';
import script$1 from './index-Q5CgsuHz.mjs';
import script from './index-D2BzYgnQ.mjs';
import { ssrRenderAttrs, ssrInterpolate, ssrRenderList, ssrRenderComponent } from 'vue/server-renderer';
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
import './index-Bg3yUxEP.mjs';
import './index-BGjNF9Dd.mjs';
import './index-mL8pBhlS.mjs';
import './index-BdDxcS0p.mjs';
import './index-D2H2HXHf.mjs';
import './index-D4owPo0H.mjs';
import './index-D255zs2l.mjs';
import './index-83EVY8bO.mjs';
import './index-IyD_XvNl.mjs';
import './index-DDXXjm8b.mjs';
import './index-Dedjwyps.mjs';
import './index-DrzwqwtL.mjs';
import './index-rAVNvoJo.mjs';
import './index-C3Ej0YBE.mjs';
import './index-CMnz8U93.mjs';
import './index-B6Gs2AMF.mjs';
import './index-l8XG4jwb.mjs';
import './index-BsfLo0vt.mjs';
import './index-DrH3FigG.mjs';
import './index-Dik5Ht0d.mjs';
import './index-Ccf11gW7.mjs';
import './index-BkCm0XJl.mjs';
import './index-bmcwiX6i.mjs';
import './index-CvVdVfGD.mjs';
import './index-DzyF1dFx.mjs';
import './index-0q6EZZjm.mjs';

const _sfc_main = {
  __name: "dashboard",
  __ssrInlineRender: true,
  async setup(__props) {
    let __temp, __restore;
    const { orders: ordersApi } = useApi();
    const { data: ordersData, pending, error } = ([__temp, __restore] = withAsyncContext(() => useAsyncData("user-orders", () => ordersApi.list())), __temp = await __temp, __restore(), __temp);
    useSeoMeta({
      title: "Dashboard - E-Commerce Store"
    });
    const orders = computed(() => {
      var _a;
      return ((_a = ordersData.value) == null ? void 0 : _a.data) || ordersData.value || [];
    });
    const stats = computed(() => ({
      totalOrders: orders.value.length,
      pending: orders.value.filter((o) => o.status === "pending" || o.status === "processing").length,
      delivered: orders.value.filter((o) => o.status === "delivered").length,
      cancelled: orders.value.filter((o) => o.status === "cancelled").length
    }));
    const formatDate = (dateStr) => {
      if (!dateStr) return "N/A";
      return new Date(dateStr).toLocaleDateString("en-US", { year: "numeric", month: "long", day: "numeric" });
    };
    return (_ctx, _push, _parent, _attrs) => {
      const _component_DataTable = script;
      const _component_Column = script$1;
      const _component_Tag = script$2;
      const _component_NuxtLink = __nuxt_component_0;
      const _component_Button = script$3;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-8" }, _attrs))}><h1 class="text-3xl font-bold text-on_surface font-display mb-8">My Dashboard</h1><div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8"><div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6"><div class="flex items-center justify-between"><div><p class="text-sm text-on_surface_variant">Total Orders</p><p class="text-3xl font-bold text-on_surface">${ssrInterpolate(unref(stats).totalOrders)}</p></div><i class="pi pi-shopping-bag text-4xl text-primary/30"></i></div></div><div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6"><div class="flex items-center justify-between"><div><p class="text-sm text-on_surface_variant">Pending</p><p class="text-3xl font-bold text-yellow-600">${ssrInterpolate(unref(stats).pending)}</p></div><i class="pi pi-clock text-4xl text-yellow-600/30"></i></div></div><div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6"><div class="flex items-center justify-between"><div><p class="text-sm text-on_surface_variant">Delivered</p><p class="text-3xl font-bold text-green-600">${ssrInterpolate(unref(stats).delivered)}</p></div><i class="pi pi-check-circle text-4xl text-green-600/30"></i></div></div><div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6"><div class="flex items-center justify-between"><div><p class="text-sm text-on_surface_variant">Cancelled</p><p class="text-3xl font-bold text-red-600">${ssrInterpolate(unref(stats).cancelled)}</p></div><i class="pi pi-times-circle text-4xl text-red-600/30"></i></div></div></div><div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6"><h2 class="text-xl font-semibold text-on_surface mb-4">Recent Orders</h2>`);
      if (unref(pending)) {
        _push(`<div class="space-y-3"><!--[-->`);
        ssrRenderList(3, (i) => {
          _push(`<div class="h-16 bg-surface-container rounded-lg animate-pulse"></div>`);
        });
        _push(`<!--]--></div>`);
      } else if (unref(orders).length === 0) {
        _push(`<div class="text-center py-8"><i class="pi pi-file text-5xl text-outline mb-4"></i><p class="text-on_surface_variant">No orders yet</p></div>`);
      } else {
        _push(`<div class="overflow-x-auto">`);
        _push(ssrRenderComponent(_component_DataTable, {
          value: unref(orders),
          class: "w-full"
        }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(ssrRenderComponent(_component_Column, {
                field: "order_id",
                header: "Order ID"
              }, null, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_Column, { header: "Date" }, {
                body: withCtx(({ data }, _push3, _parent3, _scopeId2) => {
                  if (_push3) {
                    _push3(`${ssrInterpolate(formatDate(data.created_at || data.order_date))}`);
                  } else {
                    return [
                      createTextVNode(toDisplayString(formatDate(data.created_at || data.order_date)), 1)
                    ];
                  }
                }),
                _: 1
              }, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_Column, { header: "Status" }, {
                body: withCtx(({ data }, _push3, _parent3, _scopeId2) => {
                  if (_push3) {
                    _push3(ssrRenderComponent(_component_Tag, {
                      value: data.status || "Pending",
                      severity: data.status === "delivered" ? "success" : data.status === "cancelled" ? "danger" : "warn"
                    }, null, _parent3, _scopeId2));
                  } else {
                    return [
                      createVNode(_component_Tag, {
                        value: data.status || "Pending",
                        severity: data.status === "delivered" ? "success" : data.status === "cancelled" ? "danger" : "warn"
                      }, null, 8, ["value", "severity"])
                    ];
                  }
                }),
                _: 1
              }, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_Column, { header: "Total" }, {
                body: withCtx(({ data }, _push3, _parent3, _scopeId2) => {
                  if (_push3) {
                    _push3(`\u20B9 ${ssrInterpolate(Number(data.total_amount || data.total || 0).toFixed(2))}`);
                  } else {
                    return [
                      createTextVNode("\u20B9 " + toDisplayString(Number(data.total_amount || data.total || 0).toFixed(2)), 1)
                    ];
                  }
                }),
                _: 1
              }, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_Column, { header: "" }, {
                body: withCtx(({ data }, _push3, _parent3, _scopeId2) => {
                  if (_push3) {
                    _push3(ssrRenderComponent(_component_NuxtLink, {
                      to: `/orders/${data.order_id || data.id}`
                    }, {
                      default: withCtx((_2, _push4, _parent4, _scopeId3) => {
                        if (_push4) {
                          _push4(ssrRenderComponent(_component_Button, {
                            label: "View",
                            icon: "pi pi-eye",
                            text: "",
                            size: "small"
                          }, null, _parent4, _scopeId3));
                        } else {
                          return [
                            createVNode(_component_Button, {
                              label: "View",
                              icon: "pi pi-eye",
                              text: "",
                              size: "small"
                            })
                          ];
                        }
                      }),
                      _: 2
                    }, _parent3, _scopeId2));
                  } else {
                    return [
                      createVNode(_component_NuxtLink, {
                        to: `/orders/${data.order_id || data.id}`
                      }, {
                        default: withCtx(() => [
                          createVNode(_component_Button, {
                            label: "View",
                            icon: "pi pi-eye",
                            text: "",
                            size: "small"
                          })
                        ]),
                        _: 1
                      }, 8, ["to"])
                    ];
                  }
                }),
                _: 1
              }, _parent2, _scopeId));
            } else {
              return [
                createVNode(_component_Column, {
                  field: "order_id",
                  header: "Order ID"
                }),
                createVNode(_component_Column, { header: "Date" }, {
                  body: withCtx(({ data }) => [
                    createTextVNode(toDisplayString(formatDate(data.created_at || data.order_date)), 1)
                  ]),
                  _: 1
                }),
                createVNode(_component_Column, { header: "Status" }, {
                  body: withCtx(({ data }) => [
                    createVNode(_component_Tag, {
                      value: data.status || "Pending",
                      severity: data.status === "delivered" ? "success" : data.status === "cancelled" ? "danger" : "warn"
                    }, null, 8, ["value", "severity"])
                  ]),
                  _: 1
                }),
                createVNode(_component_Column, { header: "Total" }, {
                  body: withCtx(({ data }) => [
                    createTextVNode("\u20B9 " + toDisplayString(Number(data.total_amount || data.total || 0).toFixed(2)), 1)
                  ]),
                  _: 1
                }),
                createVNode(_component_Column, { header: "" }, {
                  body: withCtx(({ data }) => [
                    createVNode(_component_NuxtLink, {
                      to: `/orders/${data.order_id || data.id}`
                    }, {
                      default: withCtx(() => [
                        createVNode(_component_Button, {
                          label: "View",
                          icon: "pi pi-eye",
                          text: "",
                          size: "small"
                        })
                      ]),
                      _: 1
                    }, 8, ["to"])
                  ]),
                  _: 1
                })
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</div>`);
      }
      _push(`</div></div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/user/dashboard.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};

export { _sfc_main as default };
//# sourceMappingURL=dashboard-DZICyxnJ.mjs.map
