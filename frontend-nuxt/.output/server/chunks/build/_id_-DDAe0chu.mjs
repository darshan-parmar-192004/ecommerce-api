import { _ as __nuxt_component_0 } from './nuxt-link-1RlcU5D7.mjs';
import { u as useRoute, b as useCartStore, a as useSeoMeta, s as script, f as fetchDefaults, e as useRequestFetch } from './server.mjs';
import { withAsyncContext, ref, computed, mergeProps, withCtx, createVNode, createTextVNode, unref, isRef, toValue, reactive, useSSRContext } from 'vue';
import { K as hash } from '../nitro/nitro.mjs';
import { u as useAsyncData } from './asyncData--maOkXya.mjs';
import script$4 from './index-zRb3CJ4f.mjs';
import script$3 from './index-CMnz8U93.mjs';
import script$2 from './index-CZDIRxVY.mjs';
import script$1 from './index-BGLbqIM8.mjs';
import { ssrRenderAttrs, ssrRenderComponent, ssrRenderStyle, ssrInterpolate, ssrRenderList } from 'vue/server-renderer';
import 'unhead';
import '@unhead/shared';
import 'vue-router';
import '@primeuix/utils/eventbus';
import '@primeuix/styled';
import '@primeuix/utils';
import '@primeuix/utils/object';
import '@primeuix/styles/base';
import '@primeuix/utils/dom';
import '@primeuix/utils/zindex';
import '@primeuix/styles/toast';
import '@primeuix/utils/uuid';
import '@primeuix/styles/ripple';
import '@primeuix/styles/badge';
import '@primeuix/styles/button';
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
import '@primeuix/styles/carousel';
import '@primeuix/styles/galleria';
import '@primeuix/styles/image';
import '@primeuix/styles/imagecompare';
import '@primeuix/styles/avatar';
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
import './index-B6Gs2AMF.mjs';
import './index-l8XG4jwb.mjs';
import './index-DDXXjm8b.mjs';
import './index-Dedjwyps.mjs';
import './index-DrzwqwtL.mjs';

function useFetch(request, arg1, arg2) {
  const [opts = {}, autoKey] = [{}, arg1];
  const _request = computed(() => toValue(request));
  const _key = opts.key || hash([autoKey, typeof _request.value === "string" ? _request.value : "", ...generateOptionSegments(opts)]);
  if (!_key || typeof _key !== "string") {
    throw new TypeError("[nuxt] [useFetch] key must be a string: " + _key);
  }
  if (!request) {
    throw new Error("[nuxt] [useFetch] request is missing.");
  }
  const key = _key === autoKey ? "$f" + _key : _key;
  if (!opts.baseURL && typeof _request.value === "string" && (_request.value[0] === "/" && _request.value[1] === "/")) {
    throw new Error('[nuxt] [useFetch] the request URL must not start with "//".');
  }
  const {
    server,
    lazy,
    default: defaultFn,
    transform,
    pick,
    watch,
    immediate,
    getCachedData,
    deep,
    dedupe,
    ...fetchOptions
  } = opts;
  const _fetchOptions = reactive({
    ...fetchDefaults,
    ...fetchOptions,
    cache: typeof opts.cache === "boolean" ? void 0 : opts.cache
  });
  const _asyncDataOptions = {
    server,
    lazy,
    default: defaultFn,
    transform,
    pick,
    immediate,
    getCachedData,
    deep,
    dedupe,
    watch: watch === false ? [] : [_fetchOptions, _request, ...watch || []]
  };
  let controller;
  const asyncData = useAsyncData(key, () => {
    var _a;
    (_a = controller == null ? void 0 : controller.abort) == null ? void 0 : _a.call(controller);
    controller = typeof AbortController !== "undefined" ? new AbortController() : {};
    const timeoutLength = toValue(opts.timeout);
    let timeoutId;
    if (timeoutLength) {
      timeoutId = setTimeout(() => controller.abort(), timeoutLength);
      controller.signal.onabort = () => clearTimeout(timeoutId);
    }
    let _$fetch = opts.$fetch || globalThis.$fetch;
    if (!opts.$fetch) {
      const isLocalFetch = typeof _request.value === "string" && _request.value[0] === "/" && (!toValue(opts.baseURL) || toValue(opts.baseURL)[0] === "/");
      if (isLocalFetch) {
        _$fetch = useRequestFetch();
      }
    }
    return _$fetch(_request.value, { signal: controller.signal, ..._fetchOptions }).finally(() => {
      clearTimeout(timeoutId);
    });
  }, _asyncDataOptions);
  return asyncData;
}
function generateOptionSegments(opts) {
  var _a;
  const segments = [
    ((_a = toValue(opts.method)) == null ? void 0 : _a.toUpperCase()) || "GET",
    toValue(opts.baseURL)
  ];
  for (const _obj of [opts.params || opts.query]) {
    const obj = toValue(_obj);
    if (!obj) {
      continue;
    }
    const unwrapped = {};
    for (const [key, value] of Object.entries(obj)) {
      unwrapped[toValue(key)] = toValue(value);
    }
    segments.push(unwrapped);
  }
  return segments;
}
const _sfc_main = {
  __name: "[id]",
  __ssrInlineRender: true,
  async setup(__props) {
    let __temp, __restore;
    const route = useRoute();
    const cartStore = useCartStore();
    const { data: product, pending, error } = ([__temp, __restore] = withAsyncContext(() => useFetch(`/api/products/${route.params.id}`, "$H3OxYP9tuC")), __temp = await __temp, __restore(), __temp);
    useSeoMeta({
      title: () => product.value ? `${product.value.name} - E-Commerce Store` : "Product",
      description: () => {
        var _a;
        return ((_a = product.value) == null ? void 0 : _a.description) || "Product details";
      }
    });
    const quantity = ref(1);
    const added = ref(false);
    const isZoomed = ref(false);
    const mousePosition = ref({ x: 0, y: 0 });
    const reviews = ref([
      { id: 1, name: "John D.", rating: 5, comment: "Excellent product! Very satisfied with the quality.", date: "2024-01-15" },
      { id: 2, name: "Sarah M.", rating: 4, comment: "Good value for money. Fast delivery.", date: "2024-01-10" },
      { id: 3, name: "Alex K.", rating: 5, comment: "Highly recommend! Exceeded expectations.", date: "2024-01-05" }
    ]);
    const addToCart = () => {
      if (product.value) {
        for (let i = 0; i < quantity.value; i++) {
          cartStore.addItem(product.value);
        }
        added.value = true;
        cartStore.openCart();
        setTimeout(() => {
          added.value = false;
        }, 2e3);
      }
    };
    const averageRating = computed(() => {
      if (!reviews.value.length) return 0;
      return reviews.value.reduce((sum, r) => sum + r.rating, 0) / reviews.value.length;
    });
    return (_ctx, _push, _parent, _attrs) => {
      const _component_NuxtLink = __nuxt_component_0;
      const _component_Button = script;
      const _component_Tag = script$1;
      const _component_Rating = script$2;
      const _component_InputNumber = script$3;
      const _component_Avatar = script$4;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 bg-surface" }, _attrs))}>`);
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: "/products",
        class: "inline-flex items-center text-sm text-on_surface_variant hover:text-primary mb-6 transition-colors"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`<i class="pi pi-arrow-left mr-1"${_scopeId}></i> Back to Products `);
          } else {
            return [
              createVNode("i", { class: "pi pi-arrow-left mr-1" }),
              createTextVNode(" Back to Products ")
            ];
          }
        }),
        _: 1
      }, _parent));
      if (unref(pending)) {
        _push(`<div class="animate-pulse"><div class="grid grid-cols-1 md:grid-cols-2 gap-8"><div class="bg-surface-container aspect-square rounded-lg"></div><div><div class="h-8 bg-surface-container rounded w-3/4 mb-4"></div><div class="h-6 bg-surface-container rounded w-1/4 mb-6"></div><div class="h-4 bg-surface-container rounded w-full mb-2"></div><div class="h-4 bg-surface-container rounded w-2/3"></div></div></div></div>`);
      } else if (unref(error)) {
        _push(`<div class="bg-surface-container-lowest rounded-xl shadow-ambient p-8 text-center"><i class="pi pi-exclamation-triangle text-5xl text-error mb-4"></i><h1 class="mt-4 text-2xl font-bold text-on_surface mb-2">Product Not Found</h1><p class="text-on_surface_variant mb-6">The product you&#39;re looking for doesn&#39;t exist.</p>`);
        _push(ssrRenderComponent(_component_NuxtLink, { to: "/products" }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(ssrRenderComponent(_component_Button, {
                label: "Browse Products",
                icon: "pi pi-shopping-bag"
              }, null, _parent2, _scopeId));
            } else {
              return [
                createVNode(_component_Button, {
                  label: "Browse Products",
                  icon: "pi pi-shopping-bag"
                })
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</div>`);
      } else if (unref(product)) {
        _push(`<div class="grid grid-cols-1 md:grid-cols-2 gap-8"><div style="${ssrRenderStyle({ "position": "relative", "overflow": "hidden" })}"><div class="bg-surface-container aspect-square rounded-lg flex items-center justify-center overflow-hidden cursor-zoom-in" style="${ssrRenderStyle({ "position": "relative" })}"><div class="w-full h-full flex items-center justify-center transition-transform duration-200" style="${ssrRenderStyle({
          transform: unref(isZoomed) ? "scale(1.5)" : "scale(1)",
          transformOrigin: `${unref(mousePosition).x}% ${unref(mousePosition).y}%`
        })}"><i class="pi pi-box text-8xl text-outline"></i></div>`);
        if (unref(isZoomed)) {
          _push(`<div style="${ssrRenderStyle({ "position": "absolute", "top": "0", "left": "0", "width": "100%", "height": "100%", "pointer-events": "none", "border": "2px solid rgba(28, 49, 227, 0.5)", "border-radius": "0.5rem" })}"></div>`);
        } else {
          _push(`<!---->`);
        }
        _push(`</div>`);
        _push(ssrRenderComponent(_component_Button, {
          icon: "pi pi-heart",
          severity: "danger",
          text: "",
          rounded: "",
          class: "absolute top-4 right-4"
        }, null, _parent));
        if (unref(product).stock === 0) {
          _push(ssrRenderComponent(_component_Tag, {
            value: "Out of Stock",
            severity: "danger",
            class: "absolute top-4 left-4"
          }, null, _parent));
        } else {
          _push(`<!---->`);
        }
        _push(`</div><div><h1 class="text-3xl font-bold text-on_surface font-display mb-2">${ssrInterpolate(unref(product).name)}</h1><div class="flex items-center gap-4 mb-4"><div class="flex items-center gap-1">`);
        _push(ssrRenderComponent(_component_Rating, {
          modelValue: unref(averageRating),
          readonly: "",
          cancel: false
        }, null, _parent));
        _push(`<span class="ml-2 text-sm text-on_surface_variant">${ssrInterpolate(unref(averageRating).toFixed(1))} (${ssrInterpolate(unref(reviews).length)} reviews)</span></div></div><p class="text-3xl font-semibold text-primary mb-4"> \u20B9 ${ssrInterpolate(Number(unref(product).price).toFixed(2))}</p>`);
        if (unref(product).description) {
          _push(`<p class="text-on_surface_variant mb-6">${ssrInterpolate(unref(product).description)}</p>`);
        } else {
          _push(`<!---->`);
        }
        _push(`<div class="mb-6"><p class="text-sm text-on_surface_variant">`);
        if (unref(product).stock > 10) {
          _push(`<span class="text-green-600 flex items-center gap-1"><i class="pi pi-check-circle"></i> In Stock (${ssrInterpolate(unref(product).stock)} available) </span>`);
        } else if (unref(product).stock > 0) {
          _push(`<span class="text-yellow-600 flex items-center gap-1"><i class="pi pi-exclamation-circle"></i> Only ${ssrInterpolate(unref(product).stock)} left </span>`);
        } else {
          _push(`<!---->`);
        }
        _push(`</p></div><div class="mb-6"><label class="block text-sm font-semibold text-on_surface_variant mb-3">Quantity</label>`);
        _push(ssrRenderComponent(_component_InputNumber, {
          modelValue: unref(quantity),
          "onUpdate:modelValue": ($event) => isRef(quantity) ? quantity.value = $event : null,
          min: 1,
          max: unref(product).stock,
          showButtons: "",
          class: "w-auto"
        }, null, _parent));
        _push(`</div>`);
        _push(ssrRenderComponent(_component_Button, {
          onClick: addToCart,
          label: unref(added) ? "Added to Cart!" : "Add to Cart",
          icon: unref(added) ? "pi pi-check" : "pi pi-shopping-cart",
          severity: unref(added) ? "success" : "primary",
          class: "!py-3 !px-8"
        }, null, _parent));
        _push(`<div class="border-t border-outline-variant/20 pt-8 mt-8"><h3 class="text-xl font-bold text-on_surface font-display mb-6">Customer Reviews</h3><div class="space-y-4"><!--[-->`);
        ssrRenderList(unref(reviews), (review) => {
          _push(`<div class="bg-surface-container-low p-4 rounded-lg"><div class="flex items-center justify-between mb-2"><div class="flex items-center gap-2">`);
          _push(ssrRenderComponent(_component_Avatar, {
            label: review.name.charAt(0),
            shape: "circle",
            class: "bg-gradient-to-r from-primary to-primary-container text-white"
          }, null, _parent));
          _push(`<span class="font-semibold text-on_surface">${ssrInterpolate(review.name)}</span></div>`);
          _push(ssrRenderComponent(_component_Rating, {
            modelValue: review.rating,
            readonly: "",
            cancel: false
          }, null, _parent));
          _push(`</div><p class="text-on_surface_variant">${ssrInterpolate(review.comment)}</p><p class="text-xs text-outline mt-2">${ssrInterpolate(review.date)}</p></div>`);
        });
        _push(`<!--]--></div></div></div></div>`);
      } else {
        _push(`<!---->`);
      }
      _push(`</div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/products/[id].vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};

export { _sfc_main as default };
//# sourceMappingURL=_id_-DDAe0chu.mjs.map
