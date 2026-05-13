import { _ as __nuxt_component_0 } from './nuxt-link-1RlcU5D7.mjs';
import { b as useCartStore, c as useRouter, a as useSeoMeta, s as script } from './server.mjs';
import { mergeProps, unref, withCtx, createVNode, createTextVNode, toDisplayString, useSSRContext } from 'vue';
import { ssrRenderAttrs, ssrRenderComponent, ssrRenderList, ssrInterpolate } from 'vue/server-renderer';
import { _ as _export_sfc } from './_plugin-vue_export-helper-1tPrXgE0.mjs';
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
  __name: "cart",
  __ssrInlineRender: true,
  setup(__props) {
    const cartStore = useCartStore();
    const router = useRouter();
    useSeoMeta({
      title: "Shopping Cart - E-Commerce Store"
    });
    const proceedToCheckout = () => {
      router.push("/checkout");
    };
    return (_ctx, _push, _parent, _attrs) => {
      const _component_NuxtLink = __nuxt_component_0;
      const _component_Button = script;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "min-h-screen bg-surface py-8" }, _attrs))} data-v-0883ee30><div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8" data-v-0883ee30><div class="mb-8" data-v-0883ee30><h1 class="text-3xl font-bold text-on_surface font-display" data-v-0883ee30>Shopping Cart</h1><p class="text-on_surface_variant mt-1" data-v-0883ee30>Manage your cart items</p></div>`);
      if (unref(cartStore).isEmpty) {
        _push(`<div class="bg-surface-container-lowest rounded-xl shadow-ambient p-12 text-center" data-v-0883ee30><i class="pi pi-shopping-cart text-6xl text-outline mb-4" data-v-0883ee30></i><h2 class="mt-4 text-xl font-semibold text-on_surface" data-v-0883ee30>Your cart is empty</h2><p class="mt-2 text-on_surface_variant" data-v-0883ee30>Add some products to get started!</p>`);
        _push(ssrRenderComponent(_component_NuxtLink, { to: "/products" }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(ssrRenderComponent(_component_Button, {
                label: "Browse Products",
                icon: "pi pi-shopping-bag",
                class: "mt-6"
              }, null, _parent2, _scopeId));
            } else {
              return [
                createVNode(_component_Button, {
                  label: "Browse Products",
                  icon: "pi pi-shopping-bag",
                  class: "mt-6"
                })
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</div>`);
      } else {
        _push(`<div class="grid grid-cols-1 lg:grid-cols-3 gap-8" data-v-0883ee30><div class="lg:col-span-2 space-y-4" data-v-0883ee30><!--[-->`);
        ssrRenderList(unref(cartStore).items, (item) => {
          _push(`<div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6 flex gap-6 hover:shadow-lg transition-all duration-300 group" data-v-0883ee30><div class="w-28 h-28 bg-gradient-to-br from-surface-container to-surface-container-high rounded-xl flex-shrink-0 flex items-center justify-center shadow-inner" data-v-0883ee30><i class="pi pi-box text-4xl text-outline group-hover:scale-110 transition-transform duration-300" data-v-0883ee30></i></div><div class="flex-1 min-w-0" data-v-0883ee30>`);
          _push(ssrRenderComponent(_component_NuxtLink, {
            to: `/products/${item.product.product_id}`,
            class: "font-semibold text-on_surface hover:text-primary line-clamp-1 text-lg transition-colors duration-200"
          }, {
            default: withCtx((_, _push2, _parent2, _scopeId) => {
              if (_push2) {
                _push2(`${ssrInterpolate(item.product.name)}`);
              } else {
                return [
                  createTextVNode(toDisplayString(item.product.name), 1)
                ];
              }
            }),
            _: 2
          }, _parent));
          _push(`<p class="text-primary font-bold text-lg mt-1" data-v-0883ee30> \u20B9 ${ssrInterpolate(Number(item.product.price).toFixed(2))}</p><div class="flex items-center gap-4 mt-4" data-v-0883ee30><div class="flex items-center border border-outline-variant/30 rounded-lg bg-surface-container overflow-hidden shadow-inner" data-v-0883ee30>`);
          _push(ssrRenderComponent(_component_Button, {
            onClick: ($event) => unref(cartStore).updateQuantity(item.product.product_id, item.quantity - 1),
            icon: "pi pi-minus",
            text: "",
            class: "!px-4 !py-2",
            size: "small"
          }, null, _parent));
          _push(`<span class="px-4 py-2 font-medium min-w-[3rem] text-center bg-surface-container-low" data-v-0883ee30>${ssrInterpolate(item.quantity)}</span>`);
          _push(ssrRenderComponent(_component_Button, {
            onClick: ($event) => unref(cartStore).updateQuantity(item.product.product_id, item.quantity + 1),
            icon: "pi pi-plus",
            text: "",
            class: "!px-4 !py-2",
            size: "small"
          }, null, _parent));
          _push(`</div>`);
          _push(ssrRenderComponent(_component_Button, {
            onClick: ($event) => unref(cartStore).removeItem(item.product.product_id),
            severity: "danger",
            text: "",
            size: "small"
          }, {
            default: withCtx((_, _push2, _parent2, _scopeId) => {
              if (_push2) {
                _push2(`<i class="pi pi-trash mr-1" data-v-0883ee30${_scopeId}></i> Remove `);
              } else {
                return [
                  createVNode("i", { class: "pi pi-trash mr-1" }),
                  createTextVNode(" Remove ")
                ];
              }
            }),
            _: 2
          }, _parent));
          _push(ssrRenderComponent(_component_Button, {
            onClick: ($event) => unref(cartStore).saveForLater(item.product.product_id),
            severity: "secondary",
            text: "",
            size: "small"
          }, {
            default: withCtx((_, _push2, _parent2, _scopeId) => {
              if (_push2) {
                _push2(`<i class="pi pi-bookmark mr-1" data-v-0883ee30${_scopeId}></i> Save for Later `);
              } else {
                return [
                  createVNode("i", { class: "pi pi-bookmark mr-1" }),
                  createTextVNode(" Save for Later ")
                ];
              }
            }),
            _: 2
          }, _parent));
          _push(`</div></div><div class="text-right" data-v-0883ee30><p class="text-xl font-bold text-on_surface" data-v-0883ee30> \u20B9 ${ssrInterpolate((Number(item.product.price) * item.quantity).toFixed(2))}</p></div></div>`);
        });
        _push(`<!--]-->`);
        if (unref(cartStore).savedCount > 0) {
          _push(`<div class="mt-8" data-v-0883ee30><h2 class="text-xl font-bold text-on_surface font-display mb-4" data-v-0883ee30>Saved for Later (${ssrInterpolate(unref(cartStore).savedCount)})</h2><div class="space-y-4" data-v-0883ee30><!--[-->`);
          ssrRenderList(unref(cartStore).savedForLater, (item) => {
            _push(`<div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6 flex gap-6 hover:shadow-lg transition-all duration-300 group" data-v-0883ee30><div class="w-24 h-24 bg-gradient-to-br from-surface-container to-surface-container-high rounded-xl flex-shrink-0 flex items-center justify-center shadow-inner" data-v-0883ee30><i class="pi pi-box text-3xl text-outline group-hover:scale-110 transition-transform duration-300" data-v-0883ee30></i></div><div class="flex-1 min-w-0" data-v-0883ee30>`);
            _push(ssrRenderComponent(_component_NuxtLink, {
              to: `/products/${item.product.product_id}`,
              class: "font-semibold text-on_surface hover:text-primary line-clamp-1 transition-colors duration-200"
            }, {
              default: withCtx((_, _push2, _parent2, _scopeId) => {
                if (_push2) {
                  _push2(`${ssrInterpolate(item.product.name)}`);
                } else {
                  return [
                    createTextVNode(toDisplayString(item.product.name), 1)
                  ];
                }
              }),
              _: 2
            }, _parent));
            _push(`<p class="text-primary font-bold mt-1" data-v-0883ee30> \u20B9 ${ssrInterpolate(Number(item.product.price).toFixed(2))}</p><div class="flex items-center gap-3 mt-3" data-v-0883ee30>`);
            _push(ssrRenderComponent(_component_Button, {
              onClick: ($event) => unref(cartStore).moveToCart(item.product.product_id),
              severity: "secondary",
              size: "small"
            }, {
              default: withCtx((_, _push2, _parent2, _scopeId) => {
                if (_push2) {
                  _push2(`<i class="pi pi-shopping-cart mr-1" data-v-0883ee30${_scopeId}></i> Move to Cart `);
                } else {
                  return [
                    createVNode("i", { class: "pi pi-shopping-cart mr-1" }),
                    createTextVNode(" Move to Cart ")
                  ];
                }
              }),
              _: 2
            }, _parent));
            _push(ssrRenderComponent(_component_Button, {
              onClick: ($event) => unref(cartStore).removeFromSaved(item.product.product_id),
              severity: "danger",
              text: "",
              size: "small"
            }, {
              default: withCtx((_, _push2, _parent2, _scopeId) => {
                if (_push2) {
                  _push2(`<i class="pi pi-trash mr-1" data-v-0883ee30${_scopeId}></i> Remove `);
                } else {
                  return [
                    createVNode("i", { class: "pi pi-trash mr-1" }),
                    createTextVNode(" Remove ")
                  ];
                }
              }),
              _: 2
            }, _parent));
            _push(`</div></div><div class="text-right" data-v-0883ee30><p class="font-bold text-on_surface" data-v-0883ee30> \u20B9 ${ssrInterpolate((Number(item.product.price) * item.quantity).toFixed(2))}</p></div></div>`);
          });
          _push(`<!--]--></div></div>`);
        } else {
          _push(`<!---->`);
        }
        _push(`</div><div class="lg:col-span-1" data-v-0883ee30><div class="bg-surface-container-lowest/80 backdrop-blur-md rounded-xl shadow-ambient p-6 sticky top-24 border border-white/10" data-v-0883ee30><h2 class="text-xl font-bold text-on_surface font-display mb-6" data-v-0883ee30>Order Summary</h2><div class="space-y-4 text-base" data-v-0883ee30><div class="flex justify-between text-on_surface_variant" data-v-0883ee30><span data-v-0883ee30>Subtotal (${ssrInterpolate(unref(cartStore).totalItems)} items)</span><span class="font-medium text-on_surface" data-v-0883ee30>\u20B9 ${ssrInterpolate(unref(cartStore).subtotal.toFixed(2))}</span></div><div class="flex justify-between text-on_surface_variant" data-v-0883ee30><span data-v-0883ee30>Tax (8%)</span><span class="font-medium text-on_surface" data-v-0883ee30>\u20B9 ${ssrInterpolate(unref(cartStore).tax.toFixed(2))}</span></div><div class="border-t border-outline-variant/20 pt-4 flex justify-between" data-v-0883ee30><span class="font-bold text-lg text-on_surface" data-v-0883ee30>Total</span><span class="font-bold text-2xl text-primary" data-v-0883ee30>\u20B9 ${ssrInterpolate(unref(cartStore).total.toFixed(2))}</span></div></div>`);
        _push(ssrRenderComponent(_component_Button, {
          onClick: proceedToCheckout,
          label: "Proceed to Checkout",
          icon: "pi pi-arrow-right",
          iconPos: "right",
          class: "w-full !mt-6 !py-4"
        }, null, _parent));
        _push(ssrRenderComponent(_component_NuxtLink, {
          to: "/products",
          class: "block text-center text-primary hover:text-primary/80 mt-4 font-medium"
        }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(` Continue Shopping `);
            } else {
              return [
                createTextVNode(" Continue Shopping ")
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</div></div></div>`);
      }
      _push(`</div></div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/cart.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const cart = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-0883ee30"]]);

export { cart as default };
//# sourceMappingURL=cart-6cHFHKNa.mjs.map
