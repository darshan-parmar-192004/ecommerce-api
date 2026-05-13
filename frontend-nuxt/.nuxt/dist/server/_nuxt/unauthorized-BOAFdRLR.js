import { _ as __nuxt_component_0 } from "./nuxt-link-1RlcU5D7.js";
import { g as useHead } from "../server.mjs";
import { mergeProps, withCtx, createTextVNode, useSSRContext } from "vue";
import { ssrRenderAttrs, ssrRenderComponent } from "vue/server-renderer";
import "ufo";
import "ofetch";
import "#internal/nuxt/paths";
import "hookable";
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
const _sfc_main = {
  __name: "unauthorized",
  __ssrInlineRender: true,
  setup(__props) {
    useHead({
      title: "Unauthorized - The Curator"
    });
    return (_ctx, _push, _parent, _attrs) => {
      const _component_NuxtLink = __nuxt_component_0;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "min-h-screen flex items-center justify-center bg-surface" }, _attrs))}><div class="text-center p-8"><div class="mb-8"><i class="pi pi-exclamation-triangle text-6xl text-error"></i></div><h1 class="text-4xl font-bold text-on_surface mb-4">Access Denied</h1><p class="text-on_surface_variant mb-8 max-w-md mx-auto"> You don&#39;t have permission to access this page. Please contact an administrator if you believe this is an error. </p>`);
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: "/",
        class: "inline-flex items-center px-6 py-3 rounded-md font-medium text-white bg-gradient-to-r from-primary to-primary-container shadow-glow hover:shadow-xl transition-all duration-300"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(` Go to Homepage `);
          } else {
            return [
              createTextVNode(" Go to Homepage ")
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`</div></div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/unauthorized.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
export {
  _sfc_main as default
};
//# sourceMappingURL=unauthorized-BOAFdRLR.js.map
