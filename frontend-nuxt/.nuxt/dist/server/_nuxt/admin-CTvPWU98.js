import { _ as __nuxt_component_0 } from "./nuxt-link-1RlcU5D7.js";
import { u as useRoute, d as useAuthStore, s as script, _ as __nuxt_component_1, L as script$1 } from "../server.mjs";
import { mergeProps, withCtx, createVNode, createTextVNode, toDisplayString, useSSRContext } from "vue";
import { ssrRenderAttrs, ssrRenderList, ssrRenderComponent, ssrRenderClass, ssrInterpolate } from "vue/server-renderer";
import { useRouter } from "vue-router";
import { _ as _export_sfc } from "./_plugin-vue_export-helper-1tPrXgE0.js";
import "ufo";
import "ofetch";
import "#internal/nuxt/paths";
import "hookable";
import "unctx";
import "h3";
import "unhead";
import "@unhead/shared";
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
const _sfc_main$1 = {
  __name: "AdminSidebar",
  __ssrInlineRender: true,
  setup(__props) {
    const route = useRoute();
    const router = useRouter();
    const authStore = useAuthStore();
    const navigation = [
      { name: "Dashboard", href: "/admin", icon: "pi-th-large" },
      { name: "Products", href: "/admin/products", icon: "pi-box" },
      { name: "Categories", href: "/admin/categories", icon: "pi-tags" },
      { name: "Inventory", href: "/admin/inventory", icon: "pi-shopping-cart" },
      { name: "Orders", href: "/admin/orders", icon: "pi-clipboard" }
    ];
    const isActive = (href) => route.path === href;
    const logout = () => {
      authStore.clearAuth();
      router.push("/");
    };
    return (_ctx, _push, _parent, _attrs) => {
      const _component_NuxtLink = __nuxt_component_0;
      const _component_Button = script;
      _push(`<nav${ssrRenderAttrs(mergeProps({ class: "w-64 bg-surface border-r border-outline-variant/20 p-4 flex flex-col" }, _attrs))}><div class="flex-1 space-y-1"><!--[-->`);
      ssrRenderList(navigation, (item) => {
        _push(ssrRenderComponent(_component_NuxtLink, {
          key: item.name,
          to: item.href,
          class: ["flex items-center gap-3 px-3 py-2 rounded-md text-sm font-medium transition-colors", isActive(item.href) ? "bg-primary-container text-primary" : "text-on_surface_variant hover:bg-surface-container hover:text-on_surface"]
        }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(`<i class="${ssrRenderClass("pi " + item.icon)}"${_scopeId}></i> ${ssrInterpolate(item.name)}`);
            } else {
              return [
                createVNode("i", {
                  class: "pi " + item.icon
                }, null, 2),
                createTextVNode(" " + toDisplayString(item.name), 1)
              ];
            }
          }),
          _: 2
        }, _parent));
      });
      _push(`<!--]--></div><div class="mt-4 space-y-1">`);
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: "/",
        class: "flex items-center gap-3 px-3 py-2 rounded-md text-sm font-medium text-on_surface_variant hover:bg-surface-container hover:text-on_surface transition-colors"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`<i class="pi pi-home"${_scopeId}></i> Back to Website `);
          } else {
            return [
              createVNode("i", { class: "pi pi-home" }),
              createTextVNode(" Back to Website ")
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(ssrRenderComponent(_component_Button, {
        onClick: logout,
        severity: "danger",
        text: "",
        class: "w-full justify-start gap-3 px-3 py-2"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`<i class="pi pi-sign-out"${_scopeId}></i> Logout `);
          } else {
            return [
              createVNode("i", { class: "pi pi-sign-out" }),
              createTextVNode(" Logout ")
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`</div></nav>`);
    };
  }
};
const _sfc_setup$1 = _sfc_main$1.setup;
_sfc_main$1.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("components/admin/AdminSidebar.vue");
  return _sfc_setup$1 ? _sfc_setup$1(props, ctx) : void 0;
};
const _sfc_main = {};
function _sfc_ssrRender(_ctx, _push, _parent, _attrs) {
  const _component_AdminSidebar = _sfc_main$1;
  const _component_NuxtPage = __nuxt_component_1;
  const _component_Toast = script$1;
  _push(`<div${ssrRenderAttrs(mergeProps({ class: "min-h-screen bg-gray-50 dark:bg-brand-900 flex" }, _attrs))}>`);
  _push(ssrRenderComponent(_component_AdminSidebar, null, null, _parent));
  _push(`<main class="flex-1 p-8 overflow-auto dark:bg-brand-900">`);
  _push(ssrRenderComponent(_component_NuxtPage, null, null, _parent));
  _push(`</main>`);
  _push(ssrRenderComponent(_component_Toast, { position: "bottom-right" }, null, _parent));
  _push(`</div>`);
}
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("layouts/admin.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const admin = /* @__PURE__ */ _export_sfc(_sfc_main, [["ssrRender", _sfc_ssrRender]]);
export {
  admin as default
};
//# sourceMappingURL=admin-CTvPWU98.js.map
