import { _ as __nuxt_component_0 } from './nuxt-link-1RlcU5D7.mjs';
import { ref, mergeProps, withCtx, createVNode, unref, useSSRContext } from 'vue';
import { ssrRenderAttrs, ssrRenderComponent, ssrInterpolate } from 'vue/server-renderer';

const _sfc_main = {
  __name: "RedirectScreen",
  __ssrInlineRender: true,
  props: {
    message: { type: String, required: true },
    subtitle: { type: String, default: "" },
    redirectTo: { type: String, required: true }
  },
  setup(__props) {
    const countdown = ref(4);
    return (_ctx, _push, _parent, _attrs) => {
      const _component_NuxtLink = __nuxt_component_0;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "rounded-2xl p-12 bg-surface-container-lowest shadow-ambient text-center" }, _attrs))}>`);
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: "/",
        class: "inline-block mb-6"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`<span class="text-3xl font-bold font-display tracking-tight text-primary"${_scopeId}> The Curator </span>`);
          } else {
            return [
              createVNode("span", { class: "text-3xl font-bold font-display tracking-tight text-primary" }, " The Curator ")
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`<div class="flex justify-center mb-6"><i class="pi pi-spinner text-3xl text-primary animate-spin"></i></div><p class="text-lg font-medium text-on_surface mb-1">${ssrInterpolate(__props.message)}</p>`);
      if (__props.subtitle) {
        _push(`<p class="text-sm text-on_surface_variant mb-4">${ssrInterpolate(__props.subtitle)}</p>`);
      } else {
        _push(`<!---->`);
      }
      _push(`<p class="text-xs text-outline"> Redirecting in ${ssrInterpolate(unref(countdown))} second${ssrInterpolate(unref(countdown) !== 1 ? "s" : "")}... </p></div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("components/auth/RedirectScreen.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};

export { _sfc_main as _ };
//# sourceMappingURL=RedirectScreen-mve57wzr.mjs.map
