import { _ as __nuxt_component_0 } from "./nuxt-link-1RlcU5D7.js";
import { ref, mergeProps, unref, withCtx, createVNode, Transition, withModifiers, openBlock, createBlock, createTextVNode, toDisplayString, createCommentVNode, useSSRContext } from "vue";
import { b as useCartStore, s as script } from "../server.mjs";
import { ssrRenderAttrs, ssrRenderComponent, ssrIncludeBooleanAttr, ssrInterpolate } from "vue/server-renderer";
import { _ as _export_sfc } from "./_plugin-vue_export-helper-1tPrXgE0.js";
const _sfc_main = {
  __name: "ProductCard",
  __ssrInlineRender: true,
  props: {
    product: {
      type: Object,
      required: true
    }
  },
  setup(__props) {
    const props = __props;
    const cartStore = useCartStore();
    const isAdding = ref(false);
    const cardRef = ref(null);
    const tiltStyle = ref({});
    const addToCart = async () => {
      isAdding.value = true;
      cartStore.addItem(props.product);
      cartStore.openCart();
      setTimeout(() => {
        isAdding.value = false;
      }, 500);
    };
    return (_ctx, _push, _parent, _attrs) => {
      const _component_NuxtLink = __nuxt_component_0;
      const _component_Button = script;
      _push(`<div${ssrRenderAttrs(mergeProps({
        ref_key: "cardRef",
        ref: cardRef,
        class: "group overflow-hidden rounded-lg bg-surface-container-lowest shadow-ambient transition-all duration-300 hover:shadow-xl",
        style: unref(tiltStyle)
      }, _attrs))} data-v-fd9e114d>`);
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: `/products/${__props.product.product_id}`,
        class: "block aspect-[4/3] bg-surface-container relative overflow-hidden"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`<div class="absolute inset-0 flex items-center justify-center transition-transform duration-700 ease-out group-hover:scale-110" data-v-fd9e114d${_scopeId}><i class="pi pi-box text-5xl text-outline/30" data-v-fd9e114d${_scopeId}></i></div><div class="absolute inset-0 bg-on_surface/40 opacity-0 group-hover:opacity-100 transition-opacity duration-300 flex items-center justify-center" data-v-fd9e114d${_scopeId}><button${ssrIncludeBooleanAttr(unref(isAdding) || __props.product.stock === 0) ? " disabled" : ""} class="transform scale-90 group-hover:scale-100 transition-transform duration-300 bg-white text-on_surface px-6 py-3 rounded-lg font-semibold flex items-center gap-2 shadow-lg hover:bg-primary hover:text-white" data-v-fd9e114d${_scopeId}>`);
            if (unref(isAdding)) {
              _push2(`<i class="pi pi-spin pi-spinner" data-v-fd9e114d${_scopeId}></i>`);
            } else {
              _push2(`<i class="pi pi-plus" data-v-fd9e114d${_scopeId}></i>`);
            }
            _push2(` ${ssrInterpolate(unref(isAdding) ? "Adding..." : "Quick Add")}</button></div>`);
            if (__props.product.stock === 0) {
              _push2(`<div class="absolute top-3 right-3 bg-error text-on_error text-[10px] font-bold uppercase tracking-wider px-3 py-1.5 rounded" data-v-fd9e114d${_scopeId}> Out of Stock </div>`);
            } else {
              _push2(`<!---->`);
            }
            if (__props.product.category_name) {
              _push2(`<div class="absolute top-3 left-3" data-v-fd9e114d${_scopeId}><span class="text-[10px] font-bold uppercase tracking-wider px-2 py-1 rounded bg-surface/80 backdrop-blur-sm text-on_surface" data-v-fd9e114d${_scopeId}>${ssrInterpolate(__props.product.category_name)}</span></div>`);
            } else {
              _push2(`<!---->`);
            }
          } else {
            return [
              createVNode("div", { class: "absolute inset-0 flex items-center justify-center transition-transform duration-700 ease-out group-hover:scale-110" }, [
                createVNode("i", { class: "pi pi-box text-5xl text-outline/30" })
              ]),
              createVNode(Transition, { name: "fade" }, {
                default: withCtx(() => [
                  createVNode("div", { class: "absolute inset-0 bg-on_surface/40 opacity-0 group-hover:opacity-100 transition-opacity duration-300 flex items-center justify-center" }, [
                    createVNode("button", {
                      onClick: withModifiers(addToCart, ["prevent"]),
                      disabled: unref(isAdding) || __props.product.stock === 0,
                      class: "transform scale-90 group-hover:scale-100 transition-transform duration-300 bg-white text-on_surface px-6 py-3 rounded-lg font-semibold flex items-center gap-2 shadow-lg hover:bg-primary hover:text-white"
                    }, [
                      unref(isAdding) ? (openBlock(), createBlock("i", {
                        key: 0,
                        class: "pi pi-spin pi-spinner"
                      })) : (openBlock(), createBlock("i", {
                        key: 1,
                        class: "pi pi-plus"
                      })),
                      createTextVNode(" " + toDisplayString(unref(isAdding) ? "Adding..." : "Quick Add"), 1)
                    ], 8, ["disabled"])
                  ])
                ]),
                _: 1
              }),
              createVNode(Transition, { name: "fade" }, {
                default: withCtx(() => [
                  __props.product.stock === 0 ? (openBlock(), createBlock("div", {
                    key: 0,
                    class: "absolute top-3 right-3 bg-error text-on_error text-[10px] font-bold uppercase tracking-wider px-3 py-1.5 rounded"
                  }, " Out of Stock ")) : createCommentVNode("", true)
                ]),
                _: 1
              }),
              __props.product.category_name ? (openBlock(), createBlock("div", {
                key: 0,
                class: "absolute top-3 left-3"
              }, [
                createVNode("span", { class: "text-[10px] font-bold uppercase tracking-wider px-2 py-1 rounded bg-surface/80 backdrop-blur-sm text-on_surface" }, toDisplayString(__props.product.category_name), 1)
              ])) : createCommentVNode("", true)
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`<div class="p-5" data-v-fd9e114d>`);
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: `/products/${__props.product.product_id}`,
        class: "block"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`<h3 class="font-display font-semibold text-on_surface transition-colors duration-300 group-hover:text-primary line-clamp-2 min-h-[2.5rem] text-base" data-v-fd9e114d${_scopeId}>${ssrInterpolate(__props.product.name)}</h3>`);
            if (__props.product.description) {
              _push2(`<p class="mt-1.5 text-sm text-outline line-clamp-2 font-body" data-v-fd9e114d${_scopeId}>${ssrInterpolate(__props.product.description)}</p>`);
            } else {
              _push2(`<!---->`);
            }
          } else {
            return [
              createVNode("h3", { class: "font-display font-semibold text-on_surface transition-colors duration-300 group-hover:text-primary line-clamp-2 min-h-[2.5rem] text-base" }, toDisplayString(__props.product.name), 1),
              __props.product.description ? (openBlock(), createBlock("p", {
                key: 0,
                class: "mt-1.5 text-sm text-outline line-clamp-2 font-body"
              }, toDisplayString(__props.product.description), 1)) : createCommentVNode("", true)
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`<div class="mt-5 flex items-center justify-between" data-v-fd9e114d><div data-v-fd9e114d><span class="text-xl font-bold text-primary font-display" data-v-fd9e114d> ₹${ssrInterpolate(Number(__props.product.price).toFixed(2))}</span>`);
      if (__props.product.original_price && __props.product.original_price > __props.product.price) {
        _push(`<p class="text-xs text-outline line-through mt-0.5" data-v-fd9e114d> ₹${ssrInterpolate(Number(__props.product.original_price).toFixed(2))}</p>`);
      } else {
        _push(`<!---->`);
      }
      _push(`</div>`);
      _push(ssrRenderComponent(_component_Button, {
        onClick: addToCart,
        disabled: unref(isAdding) || __props.product.stock === 0,
        class: "!px-5 !py-3"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            if (unref(isAdding)) {
              _push2(`<i class="pi pi-spin pi-spinner" data-v-fd9e114d${_scopeId}></i>`);
            } else {
              _push2(`<i class="pi pi-plus" data-v-fd9e114d${_scopeId}></i>`);
            }
            _push2(`<span class="font-medium text-xs uppercase tracking-wider ml-1" data-v-fd9e114d${_scopeId}>Add</span>`);
          } else {
            return [
              unref(isAdding) ? (openBlock(), createBlock("i", {
                key: 0,
                class: "pi pi-spin pi-spinner"
              })) : (openBlock(), createBlock("i", {
                key: 1,
                class: "pi pi-plus"
              })),
              createVNode("span", { class: "font-medium text-xs uppercase tracking-wider ml-1" }, "Add")
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`</div></div></div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("components/products/ProductCard.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const ProductCard = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-fd9e114d"]]);
export {
  ProductCard as P
};
//# sourceMappingURL=ProductCard-Di5Z6QG2.js.map
