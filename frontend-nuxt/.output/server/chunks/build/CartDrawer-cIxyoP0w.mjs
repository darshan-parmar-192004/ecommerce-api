import { b as useCartStore, d as useAuthStore, c as useRouter, s as script } from './server.mjs';
import { unref, withCtx, createVNode, createTextVNode, useSSRContext } from 'vue';
import { u as useAppToast } from './useAppToast-CxoPk-aP.mjs';
import { ssrRenderTeleport, ssrInterpolate, ssrRenderComponent, ssrRenderList } from 'vue/server-renderer';
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
  __name: "CartDrawer",
  __ssrInlineRender: true,
  setup(__props) {
    const cartStore = useCartStore();
    const { isAuthenticated } = useAuthStore();
    const router = useRouter();
    useAppToast();
    const checkout = async () => {
      if (!isAuthenticated.value) {
        router.push("/auth/login?redirect=/checkout");
        return;
      }
      router.push("/checkout");
    };
    const formatPrice = (price) => {
      return `\u20B9${Number(price).toFixed(2)}`;
    };
    return (_ctx, _push, _parent, _attrs) => {
      const _component_Button = script;
      ssrRenderTeleport(_push, (_push2) => {
        if (unref(cartStore).isOpen) {
          _push2(`<div class="fixed inset-0 z-50 overflow-hidden" data-v-db7a4723><div class="absolute inset-0 bg-on_surface/30 backdrop-blur-md" data-v-db7a4723></div><div class="absolute inset-y-0 right-0 max-w-md w-full bg-surface-container-lowest/95 backdrop-blur-xl shadow-ambient flex flex-col border-l border-white/10" data-v-db7a4723><div class="flex items-center justify-between p-6 border-b border-outline-variant/20" data-v-db7a4723><div data-v-db7a4723><h2 class="text-xl font-bold text-on_surface font-display" data-v-db7a4723>Shopping Cart</h2><p class="text-xs text-outline mt-1" data-v-db7a4723>${ssrInterpolate(unref(cartStore).totalItems)} item${ssrInterpolate(unref(cartStore).totalItems !== 1 ? "s" : "")}</p></div>`);
          _push2(ssrRenderComponent(_component_Button, {
            onClick: unref(cartStore).toggleCart,
            icon: "pi pi-times",
            text: "",
            rounded: "",
            "aria-label": "Close cart"
          }, null, _parent));
          _push2(`</div><div class="flex-1 overflow-y-auto p-6" data-v-db7a4723>`);
          if (unref(cartStore).isEmpty) {
            _push2(`<div class="h-full flex flex-col items-center justify-center text-center" data-v-db7a4723><div class="w-24 h-24 rounded-full bg-surface-container flex items-center justify-center mb-6 animate-bounce-slow" data-v-db7a4723><i class="pi pi-shopping-cart text-4xl text-outline/40" data-v-db7a4723></i></div><h3 class="text-lg font-semibold text-on_surface font-display" data-v-db7a4723>Your cart is empty</h3><p class="text-on_surface_variant mt-2 font-body" data-v-db7a4723>Add products to get started</p>`);
            _push2(ssrRenderComponent(_component_Button, {
              onClick: unref(cartStore).toggleCart,
              class: "mt-6",
              label: "Continue Shopping",
              icon: "pi pi-shopping-bag"
            }, null, _parent));
            _push2(`</div>`);
          } else {
            _push2(`<div class="space-y-4" data-v-db7a4723><!--[-->`);
            ssrRenderList(unref(cartStore).items, (item) => {
              _push2(`<div class="flex gap-4 p-4 rounded-xl bg-surface-container-low/80 backdrop-blur-sm transition-all duration-300 hover:bg-surface-container hover:shadow-lg group" data-v-db7a4723><div class="w-24 h-24 bg-surface-container rounded-lg flex-shrink-0 overflow-hidden flex items-center justify-center shadow-inner" data-v-db7a4723><i class="pi pi-box text-4xl text-outline/30 group-hover:scale-110 transition-transform duration-300" data-v-db7a4723></i></div><div class="flex-1 min-w-0" data-v-db7a4723><h3 class="font-semibold text-on_surface font-display truncate group-hover:text-primary transition-colors duration-200" data-v-db7a4723>${ssrInterpolate(item.product.name)}</h3><p class="text-primary font-bold mt-1 font-display" data-v-db7a4723>${ssrInterpolate(formatPrice(item.product.price))}</p><div class="flex items-center gap-2 mt-3" data-v-db7a4723>`);
              _push2(ssrRenderComponent(_component_Button, {
                onClick: ($event) => unref(cartStore).updateQuantity(item.product.product_id, item.quantity - 1),
                icon: "pi pi-minus",
                text: "",
                rounded: "",
                class: "!w-8 !h-8",
                size: "small"
              }, null, _parent));
              _push2(`<span class="w-10 text-center font-medium text-on_surface bg-surface-container-low rounded" data-v-db7a4723>${ssrInterpolate(item.quantity)}</span>`);
              _push2(ssrRenderComponent(_component_Button, {
                onClick: ($event) => unref(cartStore).updateQuantity(item.product.product_id, item.quantity + 1),
                icon: "pi pi-plus",
                text: "",
                rounded: "",
                class: "!w-8 !h-8",
                size: "small"
              }, null, _parent));
              _push2(ssrRenderComponent(_component_Button, {
                onClick: ($event) => unref(cartStore).removeItem(item.product.product_id),
                severity: "danger",
                text: "",
                size: "small",
                class: "ml-auto"
              }, {
                default: withCtx((_, _push3, _parent2, _scopeId) => {
                  if (_push3) {
                    _push3(`<i class="pi pi-trash mr-1" data-v-db7a4723${_scopeId}></i> Remove `);
                  } else {
                    return [
                      createVNode("i", { class: "pi pi-trash mr-1" }),
                      createTextVNode(" Remove ")
                    ];
                  }
                }),
                _: 2
              }, _parent));
              _push2(`</div></div></div>`);
            });
            _push2(`<!--]--></div>`);
          }
          _push2(`</div>`);
          if (!unref(cartStore).isEmpty) {
            _push2(`<div class="border-t border-outline-variant/20 p-6 space-y-4 bg-surface-container-low/80 backdrop-blur-md" data-v-db7a4723><div class="space-y-2" data-v-db7a4723><div class="flex justify-between text-sm text-on_surface_variant" data-v-db7a4723><span data-v-db7a4723>Subtotal</span><span data-v-db7a4723>${ssrInterpolate(formatPrice(unref(cartStore).subtotal))}</span></div><div class="flex justify-between text-sm text-on_surface_variant" data-v-db7a4723><span data-v-db7a4723>Tax (8%)</span><span data-v-db7a4723>${ssrInterpolate(formatPrice(unref(cartStore).tax))}</span></div><div class="flex justify-between text-lg font-bold text-on_surface pt-2 border-t border-outline-variant/20" data-v-db7a4723><span data-v-db7a4723>Total</span><span class="text-primary" data-v-db7a4723>${ssrInterpolate(formatPrice(unref(cartStore).total))}</span></div></div>`);
            _push2(ssrRenderComponent(_component_Button, {
              onClick: checkout,
              label: "Checkout",
              icon: "pi pi-arrow-right",
              "icon-pos": "right",
              class: "w-full !py-4"
            }, null, _parent));
            _push2(`</div>`);
          } else {
            _push2(`<!---->`);
          }
          _push2(`</div></div>`);
        } else {
          _push2(`<!---->`);
        }
      }, "body", false, _parent);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("components/cart/CartDrawer.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const CartDrawer = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-db7a4723"]]);

export { CartDrawer as default };
//# sourceMappingURL=CartDrawer-cIxyoP0w.mjs.map
