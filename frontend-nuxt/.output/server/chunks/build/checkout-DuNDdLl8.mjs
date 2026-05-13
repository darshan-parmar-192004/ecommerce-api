import { _ as __nuxt_component_0 } from './nuxt-link-1RlcU5D7.mjs';
import { ref, computed, reactive, mergeProps, unref, withCtx, createVNode, isRef, useSSRContext } from 'vue';
import { b as useCartStore, d as useAuthStore, c as useRouter, a as useSeoMeta, s as script } from './server.mjs';
import { u as useApi } from './useApi-CR6WE1xi.mjs';
import { u as useAppToast } from './useAppToast-CxoPk-aP.mjs';
import script$4 from './index-fZs5VxXm.mjs';
import script$3 from './index-DDXXjm8b.mjs';
import script$2 from './index-BAz9mnEE.mjs';
import script$1 from './index-otP1ftW0.mjs';
import { ssrRenderAttrs, ssrRenderComponent, ssrInterpolate, ssrRenderList, ssrRenderClass } from 'vue/server-renderer';
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
import './index-Dedjwyps.mjs';
import './index-DrzwqwtL.mjs';

const _sfc_main = {
  __name: "checkout",
  __ssrInlineRender: true,
  setup(__props) {
    const cartStore = useCartStore();
    const authStore = useAuthStore();
    const router = useRouter();
    const { orders: ordersApi } = useApi();
    const { success: showSuccess, error: showError } = useAppToast();
    useSeoMeta({
      title: "Checkout - E-Commerce Store"
    });
    const currentStep = ref(1);
    const loading = ref(false);
    const processingPayment = ref(false);
    const error = ref("");
    const success = ref(false);
    const selectedPaymentMethod = ref("razorpay");
    const paymentMethods = [
      { id: "razorpay", name: "Pay with Razorpay", icon: "credit-card", description: "Instant payment" },
      { id: "upi", name: "UPI", icon: "mobile", description: "Pay via UPI app" },
      { id: "wallet", name: "Wallet", icon: "wallet", description: "Paytm, PhonePe, etc." },
      { id: "cod", name: "Cash on Delivery", icon: "money", description: "Pay when you receive" }
    ];
    const promoCode = ref("");
    const promoApplied = ref(false);
    const promoDiscount = ref(0);
    const promoError = ref("");
    const applyPromoCode = () => {
      promoError.value = "";
      const code = promoCode.value.trim().toUpperCase();
      const promoCodes = { "SAVE10": 10, "FLAT20": 20, "FIRST50": 50 };
      if (promoCodes[code]) {
        promoDiscount.value = promoCodes[code];
        promoApplied.value = true;
      } else {
        promoError.value = "Invalid promo code";
      }
    };
    const removePromo = () => {
      promoCode.value = "";
      promoDiscount.value = 0;
      promoApplied.value = false;
    };
    const discountAmount = computed(() => cartStore.subtotal * promoDiscount.value / 100);
    const discountedTotal = computed(() => cartStore.total - discountAmount.value);
    const shippingForm = reactive({
      fullName: "",
      address: "",
      city: "",
      state: "",
      zipCode: "",
      phone: ""
    });
    const shippingErrors = reactive({});
    const validateShipping = () => {
      shippingErrors.fullName = !shippingForm.fullName.trim() ? "Full name is required" : "";
      shippingErrors.address = !shippingForm.address.trim() ? "Address is required" : "";
      shippingErrors.city = !shippingForm.city.trim() ? "City is required" : "";
      shippingErrors.state = !shippingForm.state.trim() ? "State is required" : "";
      shippingErrors.zipCode = !shippingForm.zipCode.trim() ? "ZIP code is required" : "";
      shippingErrors.phone = !shippingForm.phone.trim() ? "Phone number is required" : "";
      return !Object.values(shippingErrors).some((e) => e);
    };
    const goToShipping = () => {
      currentStep.value = 1;
    };
    const goToPayment = () => {
      if (validateShipping()) currentStep.value = 2;
    };
    const mockRazorpayPayment = async () => {
      processingPayment.value = true;
      try {
        await new Promise((resolve) => setTimeout(resolve, 2e3));
        await handleSubmit();
      } catch (err) {
        showError("Payment failed. Please try again.");
      } finally {
        processingPayment.value = false;
      }
    };
    const handleSubmit = async () => {
      var _a, _b;
      if (cartStore.isEmpty) {
        error.value = "Your cart is empty";
        return;
      }
      loading.value = true;
      error.value = "";
      try {
        const shippingAddress = shippingForm.fullName + ", " + shippingForm.address + ", " + shippingForm.city + ", " + shippingForm.state + " " + shippingForm.zipCode + ", " + shippingForm.phone;
        const orderData = {
          order: { customer_id: (_a = authStore.user) == null ? void 0 : _a.customer_id, shipping_address: shippingAddress, total_amount: discountedTotal.value },
          items: cartStore.items.map((item) => ({ product_id: item.product.product_id, quantity: item.quantity, unit_price: item.product.price }))
        };
        const result = await ordersApi.create(orderData);
        cartStore.clearCart();
        success.value = true;
        showSuccess("Order placed successfully!");
        const orderId = result.order_id || ((_b = result.data) == null ? void 0 : _b.order_id);
        setTimeout(() => {
          router.push("/orders/" + orderId);
        }, 1500);
      } catch (err) {
        error.value = err.message || "Failed to create order. Please try again.";
        showError(error.value);
      } finally {
        loading.value = false;
      }
    };
    return (_ctx, _push, _parent, _attrs) => {
      const _component_NuxtLink = __nuxt_component_0;
      const _component_Button = script;
      const _component_StepList = script$1;
      const _component_Message = script$2;
      const _component_InputText = script$3;
      const _component_Textarea = script$4;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8 bg-surface" }, _attrs))}><h1 class="text-3xl font-bold text-on_surface font-display mb-8">Checkout</h1>`);
      if (unref(cartStore).isEmpty) {
        _push(`<div class="text-center py-16 bg-surface-container-lowest rounded-xl shadow-ambient"><i class="pi pi-shopping-cart text-6xl text-outline mb-4"></i><h2 class="mt-4 text-xl font-medium text-on_surface">Your cart is empty</h2><p class="mt-2 text-on_surface_variant">Add some products before checking out.</p>`);
        _push(ssrRenderComponent(_component_NuxtLink, { to: "/products" }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(ssrRenderComponent(_component_Button, {
                label: "Browse Products",
                class: "mt-6",
                icon: "pi pi-shopping-bag"
              }, null, _parent2, _scopeId));
            } else {
              return [
                createVNode(_component_Button, {
                  label: "Browse Products",
                  class: "mt-6",
                  icon: "pi pi-shopping-bag"
                })
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</div>`);
      } else if (unref(success)) {
        _push(`<div class="text-center py-16 bg-surface-container-lowest rounded-xl shadow-ambient"><i class="pi pi-check-circle text-6xl text-green-600 mb-4"></i><h2 class="mt-4 text-2xl font-medium text-on_surface">Order Placed Successfully!</h2><p class="mt-2 text-on_surface_variant">Redirecting to your order...</p></div>`);
      } else {
        _push(`<div>`);
        _push(ssrRenderComponent(_component_StepList, {
          model: [{ label: "Shipping" }, { label: "Payment" }, { label: "Confirm" }],
          activeStep: unref(currentStep) - 1,
          class: "mb-8"
        }, null, _parent));
        _push(`<div class="grid grid-cols-1 lg:grid-cols-2 gap-8"><div>`);
        if (unref(currentStep) === 1) {
          _push(`<div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6"><h2 class="text-lg font-semibold text-on_surface font-display mb-4">Shipping Information</h2>`);
          if (unref(error)) {
            _push(ssrRenderComponent(_component_Message, {
              severity: "error",
              closable: false,
              class: "mb-4"
            }, null, _parent));
          } else {
            _push(`<!---->`);
          }
          _push(`<div class="space-y-4"><div><label class="block text-sm font-medium text-on_surface_variant mb-2">Full Name <span class="text-error">*</span></label>`);
          _push(ssrRenderComponent(_component_InputText, {
            modelValue: unref(shippingForm).fullName,
            "onUpdate:modelValue": ($event) => unref(shippingForm).fullName = $event,
            class: "w-full",
            placeholder: "John Doe"
          }, null, _parent));
          if (unref(shippingErrors).fullName) {
            _push(`<small class="text-error">${ssrInterpolate(unref(shippingErrors).fullName)}</small>`);
          } else {
            _push(`<!---->`);
          }
          _push(`</div><div><label class="block text-sm font-medium text-on_surface_variant mb-2">Address <span class="text-error">*</span></label>`);
          _push(ssrRenderComponent(_component_Textarea, {
            modelValue: unref(shippingForm).address,
            "onUpdate:modelValue": ($event) => unref(shippingForm).address = $event,
            rows: "2",
            class: "w-full",
            placeholder: "123 Main Street"
          }, null, _parent));
          if (unref(shippingErrors).address) {
            _push(`<small class="text-error">${ssrInterpolate(unref(shippingErrors).address)}</small>`);
          } else {
            _push(`<!---->`);
          }
          _push(`</div><div class="grid grid-cols-2 gap-4"><div><label class="block text-sm font-medium text-on_surface_variant mb-2">City <span class="text-error">*</span></label>`);
          _push(ssrRenderComponent(_component_InputText, {
            modelValue: unref(shippingForm).city,
            "onUpdate:modelValue": ($event) => unref(shippingForm).city = $event,
            class: "w-full",
            placeholder: "New York"
          }, null, _parent));
          if (unref(shippingErrors).city) {
            _push(`<small class="text-error">${ssrInterpolate(unref(shippingErrors).city)}</small>`);
          } else {
            _push(`<!---->`);
          }
          _push(`</div><div><label class="block text-sm font-medium text-on_surface_variant mb-2">State <span class="text-error">*</span></label>`);
          _push(ssrRenderComponent(_component_InputText, {
            modelValue: unref(shippingForm).state,
            "onUpdate:modelValue": ($event) => unref(shippingForm).state = $event,
            class: "w-full",
            placeholder: "NY"
          }, null, _parent));
          if (unref(shippingErrors).state) {
            _push(`<small class="text-error">${ssrInterpolate(unref(shippingErrors).state)}</small>`);
          } else {
            _push(`<!---->`);
          }
          _push(`</div></div><div class="grid grid-cols-2 gap-4"><div><label class="block text-sm font-medium text-on_surface_variant mb-2">ZIP Code <span class="text-error">*</span></label>`);
          _push(ssrRenderComponent(_component_InputText, {
            modelValue: unref(shippingForm).zipCode,
            "onUpdate:modelValue": ($event) => unref(shippingForm).zipCode = $event,
            class: "w-full",
            placeholder: "10001"
          }, null, _parent));
          if (unref(shippingErrors).zipCode) {
            _push(`<small class="text-error">${ssrInterpolate(unref(shippingErrors).zipCode)}</small>`);
          } else {
            _push(`<!---->`);
          }
          _push(`</div><div><label class="block text-sm font-medium text-on_surface_variant mb-2">Phone <span class="text-error">*</span></label>`);
          _push(ssrRenderComponent(_component_InputText, {
            modelValue: unref(shippingForm).phone,
            "onUpdate:modelValue": ($event) => unref(shippingForm).phone = $event,
            class: "w-full",
            placeholder: "+1 234 567 8900"
          }, null, _parent));
          if (unref(shippingErrors).phone) {
            _push(`<small class="text-error">${ssrInterpolate(unref(shippingErrors).phone)}</small>`);
          } else {
            _push(`<!---->`);
          }
          _push(`</div></div>`);
          _push(ssrRenderComponent(_component_Button, {
            onClick: goToPayment,
            label: "Continue to Payment",
            icon: "pi pi-arrow-right",
            iconPos: "right",
            class: "w-full !mt-4 !py-3"
          }, null, _parent));
          _push(`</div></div>`);
        } else {
          _push(`<!---->`);
        }
        if (unref(currentStep) === 2) {
          _push(`<div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6"><h2 class="text-lg font-semibold text-on_surface font-display mb-4">Payment Information</h2>`);
          if (unref(error)) {
            _push(ssrRenderComponent(_component_Message, {
              severity: "error",
              closable: false,
              class: "mb-4"
            }, null, _parent));
          } else {
            _push(`<!---->`);
          }
          _push(`<div class="mb-6 p-4 bg-surface-container rounded-lg"><label class="block text-sm font-medium text-on_surface_variant mb-2">Promo Code</label>`);
          if (unref(promoApplied)) {
            _push(`<div class="flex items-center justify-between"><div class="flex items-center gap-2"><i class="pi pi-check-circle text-green-600"></i><span class="text-green-600 font-medium">${ssrInterpolate(unref(promoCode))} applied (-${ssrInterpolate(unref(promoDiscount))}%)</span></div>`);
            _push(ssrRenderComponent(_component_Button, {
              onClick: removePromo,
              label: "Remove",
              text: "",
              severity: "danger",
              size: "small"
            }, null, _parent));
            _push(`</div>`);
          } else {
            _push(`<div class="flex gap-2">`);
            _push(ssrRenderComponent(_component_InputText, {
              modelValue: unref(promoCode),
              "onUpdate:modelValue": ($event) => isRef(promoCode) ? promoCode.value = $event : null,
              placeholder: "Enter promo code",
              class: "flex-1",
              onKeyup: applyPromoCode
            }, null, _parent));
            _push(ssrRenderComponent(_component_Button, {
              onClick: applyPromoCode,
              label: "Apply"
            }, null, _parent));
            _push(`</div>`);
          }
          if (unref(promoError)) {
            _push(`<small class="text-error">${ssrInterpolate(unref(promoError))}</small>`);
          } else {
            _push(`<!---->`);
          }
          _push(`<p class="text-xs text-outline mt-2">Try: SAVE10, FLAT20, FIRST50</p></div><div class="mb-6"><label class="block text-sm font-medium text-on_surface_variant mb-3">Payment Method</label><div class="space-y-3"><!--[-->`);
          ssrRenderList(paymentMethods, (method) => {
            _push(`<div class="${ssrRenderClass([unref(selectedPaymentMethod) === method.id ? "border-primary bg-primary/10" : "border-outline-variant/30 hover:border-primary/50 bg-surface-container", "flex items-center gap-3 p-3 rounded-lg border-2 cursor-pointer transition-all"])}"><i class="${ssrRenderClass(["pi pi-" + method.icon, "text-xl"])}"></i><div><p class="font-semibold text-on_surface text-sm">${ssrInterpolate(method.name)}</p><p class="text-xs text-on_surface_variant">${ssrInterpolate(method.description)}</p></div></div>`);
          });
          _push(`<!--]--></div></div><div class="bg-surface-container rounded-lg p-4 mb-6 text-left"><div class="flex justify-between items-center py-2 border-b border-outline-variant/20"><span class="text-on_surface_variant">Order Total</span><span class="text-xl font-bold text-primary">Rs ${ssrInterpolate(unref(discountedTotal).toFixed(2))}</span></div>`);
          if (unref(promoApplied)) {
            _push(`<div class="flex justify-between items-center py-2 border-b border-outline-variant/20 text-green-600"><span>Discount (${ssrInterpolate(unref(promoDiscount))}%)</span><span>-Rs ${ssrInterpolate(unref(discountAmount).toFixed(2))}</span></div>`);
          } else {
            _push(`<!---->`);
          }
          _push(`</div>`);
          _push(ssrRenderComponent(_component_Button, {
            onClick: mockRazorpayPayment,
            loading: unref(processingPayment),
            label: unref(processingPayment) ? "Processing Payment..." : "Pay Rs " + unref(discountedTotal).toFixed(2),
            icon: "pi pi-credit-card",
            class: "w-full !py-4"
          }, null, _parent));
          _push(ssrRenderComponent(_component_Button, {
            onClick: goToShipping,
            label: "Back",
            icon: "pi pi-arrow-left",
            text: "",
            class: "w-full mt-3"
          }, null, _parent));
          _push(`<p class="text-xs text-on_surface_variant mt-4 text-center"><i class="pi pi-lock mr-1"></i>This is a demo payment. No real money will be charged.</p></div>`);
        } else {
          _push(`<!---->`);
        }
        _push(`</div><div><div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6"><h2 class="text-lg font-semibold text-on_surface font-display mb-4">Order Summary</h2><div class="space-y-4 max-h-80 overflow-y-auto"><!--[-->`);
        ssrRenderList(unref(cartStore).items, (item) => {
          _push(`<div class="flex items-center gap-4 py-3 border-b border-outline-variant/20 last:border-0"><div class="w-14 h-14 bg-surface-container rounded-lg flex-shrink-0 flex items-center justify-center"><i class="pi pi-box text-2xl text-outline"></i></div><div class="flex-1 min-w-0"><p class="font-medium text-on_surface truncate">${ssrInterpolate(item.product.name)}</p><p class="text-sm text-on_surface_variant">Qty: ${ssrInterpolate(item.quantity)} x Rs ${ssrInterpolate(Number(item.product.price).toFixed(2))}</p></div><p class="font-semibold text-on_surface">Rs ${ssrInterpolate((Number(item.product.price) * item.quantity).toFixed(2))}</p></div>`);
        });
        _push(`<!--]--></div><div class="mt-6 pt-4 border-t border-outline-variant/20 space-y-2"><div class="flex justify-between text-sm text-on_surface_variant"><span>Subtotal</span><span class="font-medium text-on_surface">Rs ${ssrInterpolate(unref(cartStore).subtotal.toFixed(2))}</span></div>`);
        if (unref(promoApplied)) {
          _push(`<div class="flex justify-between text-sm text-green-600"><span>Discount (${ssrInterpolate(unref(promoDiscount))}%)</span><span>-Rs ${ssrInterpolate(unref(discountAmount).toFixed(2))}</span></div>`);
        } else {
          _push(`<!---->`);
        }
        _push(`<div class="flex justify-between text-sm text-on_surface_variant"><span>Tax (8%)</span><span class="font-medium text-on_surface">Rs ${ssrInterpolate(unref(cartStore).tax.toFixed(2))}</span></div><div class="flex justify-between text-base pt-2 border-t border-outline-variant/20"><span class="font-semibold text-on_surface">Total</span><span class="font-bold text-xl text-primary">Rs ${ssrInterpolate(unref(discountedTotal).toFixed(2))}</span></div></div></div></div></div></div>`);
      }
      _push(`</div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/checkout.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};

export { _sfc_main as default };
//# sourceMappingURL=checkout-DuNDdLl8.mjs.map
