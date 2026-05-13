import { _ as __nuxt_component_0 } from "./nuxt-link-1RlcU5D7.js";
import { withAsyncContext, computed, mergeProps, unref, withCtx, createTextVNode, toDisplayString, createVNode, useSSRContext } from "vue";
import "hookable";
import { u as useApi } from "./useApi-CR6WE1xi.js";
import { u as useAsyncData } from "./asyncData--maOkXya.js";
import { a as useSeoMeta, s as script$3 } from "../server.mjs";
import script$2 from "./index-BGLbqIM8.js";
import script$1 from "./index-Q5CgsuHz.js";
import script from "./index-D2BzYgnQ.js";
import { ssrRenderAttrs, ssrInterpolate, ssrRenderList, ssrRenderComponent } from "vue/server-renderer";
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
import "@primeuix/styles/tag";
import "./index-Bg3yUxEP.js";
import "./index-BGjNF9Dd.js";
import "@primeuix/styles/paginator";
import "./index-mL8pBhlS.js";
import "./index-BdDxcS0p.js";
import "./index-D2H2HXHf.js";
import "./index-D4owPo0H.js";
import "./index-D255zs2l.js";
import "./index-83EVY8bO.js";
import "@primeuix/styles/iconfield";
import "./index-IyD_XvNl.js";
import "./index-DDXXjm8b.js";
import "./index-Dedjwyps.js";
import "./index-DrzwqwtL.js";
import "@primeuix/styles/inputtext";
import "./index-rAVNvoJo.js";
import "./index-C3Ej0YBE.js";
import "@primeuix/styles/virtualscroller";
import "@primeuix/styles/select";
import "./index-CMnz8U93.js";
import "./index-B6Gs2AMF.js";
import "./index-l8XG4jwb.js";
import "@primeuix/styles/inputnumber";
import "./index-BsfLo0vt.js";
import "@primeuix/styles/datatable";
import "./index-DrH3FigG.js";
import "./index-Dik5Ht0d.js";
import "./index-Ccf11gW7.js";
import "./index-BkCm0XJl.js";
import "@primeuix/styles/checkbox";
import "./index-bmcwiX6i.js";
import "@primeuix/styles/radiobutton";
import "./index-CvVdVfGD.js";
import "./index-DzyF1dFx.js";
import "./index-0q6EZZjm.js";
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
                    _push3(`₹ ${ssrInterpolate(Number(data.total_amount || data.total || 0).toFixed(2))}`);
                  } else {
                    return [
                      createTextVNode("₹ " + toDisplayString(Number(data.total_amount || data.total || 0).toFixed(2)), 1)
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
                    createTextVNode("₹ " + toDisplayString(Number(data.total_amount || data.total || 0).toFixed(2)), 1)
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
export {
  _sfc_main as default
};
//# sourceMappingURL=dashboard-DZICyxnJ.js.map
