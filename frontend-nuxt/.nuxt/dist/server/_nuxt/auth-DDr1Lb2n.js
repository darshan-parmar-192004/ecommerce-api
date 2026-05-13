import { mergeProps, useSSRContext } from "vue";
import { ssrRenderAttrs, ssrRenderClass, ssrRenderStyle, ssrRenderComponent, ssrRenderSlot } from "vue/server-renderer";
import { _ as _export_sfc } from "./_plugin-vue_export-helper-1tPrXgE0.js";
const _sfc_main$1 = {
  __name: "AnimatedGrid",
  __ssrInlineRender: true,
  props: {
    variant: {
      type: String,
      default: "full"
      // 'full' | 'hero' | 'minimal'
    }
  },
  setup(__props) {
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "fixed inset-0 -z-10 overflow-hidden pointer-events-none" }, _attrs))} data-v-c3b48cfd><div class="absolute inset-0 bg-gradient-to-br from-gray-50 via-white to-gray-100" data-v-c3b48cfd></div><div class="${ssrRenderClass([
        "absolute inset-0 opacity-[0.03]",
        __props.variant === "full" ? "bg-[linear-gradient(#1e293b_1px,transparent_1px),linear-gradient(90deg,#1e293b_1px,transparent_1px)] bg-[size:50px_50px] [animation:grid_20s_linear_infinite]" : "",
        __props.variant === "hero" ? "bg-[linear-gradient(#1e293b_1px,transparent_1px),linear-gradient(90deg,#1e293b_1px,transparent_1px)] bg-[size:30px_30px]" : "",
        __props.variant === "minimal" ? "bg-[linear-gradient(#1e293b_1px,transparent_1px),linear-gradient(90deg,#1e293b_1px,transparent_1px)] bg-[size:100px_100px]" : ""
      ])}" data-v-c3b48cfd></div><div class="absolute top-0 left-1/4 w-96 h-96 bg-primary-500/5 rounded-full blur-3xl animate-float" data-v-c3b48cfd></div><div class="absolute bottom-0 right-1/4 w-96 h-96 bg-accent-500/5 rounded-full blur-3xl animate-float" style="${ssrRenderStyle({ "animation-delay": "-3s" })}" data-v-c3b48cfd></div><div class="absolute top-0 left-0 right-0 h-px bg-gradient-to-r from-transparent via-primary-200/30 to-transparent" data-v-c3b48cfd></div></div>`);
    };
  }
};
const _sfc_setup$1 = _sfc_main$1.setup;
_sfc_main$1.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("components/common/AnimatedGrid.vue");
  return _sfc_setup$1 ? _sfc_setup$1(props, ctx) : void 0;
};
const AnimatedGrid = /* @__PURE__ */ _export_sfc(_sfc_main$1, [["__scopeId", "data-v-c3b48cfd"]]);
const _sfc_main = {
  __name: "auth",
  __ssrInlineRender: true,
  setup(__props) {
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "min-h-screen w-full flex items-center justify-center p-4 relative overflow-hidden bg-surface" }, _attrs))}>`);
      _push(ssrRenderComponent(AnimatedGrid, { variant: "full" }, null, _parent));
      _push(`<div class="relative z-10 w-full max-w-md">`);
      ssrRenderSlot(_ctx.$slots, "default", {}, null, _push, _parent);
      _push(`</div></div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("layouts/auth.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
export {
  _sfc_main as default
};
//# sourceMappingURL=auth-DDr1Lb2n.js.map
