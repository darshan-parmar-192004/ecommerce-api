import { ref, mergeProps, unref, withCtx, createTextVNode, toDisplayString, createVNode, isRef, useSSRContext } from 'vue';
import script$5 from './index-CMnz8U93.mjs';
import script$4 from './index-DDXXjm8b.mjs';
import script$3 from './index-D7DvHjf4.mjs';
import script$2 from './index-Q5CgsuHz.mjs';
import script$1 from './index-D2BzYgnQ.mjs';
import { s as script } from './server.mjs';
import { ssrRenderAttrs, ssrRenderComponent, ssrRenderList, ssrInterpolate } from 'vue/server-renderer';
import { u as useApi } from './useApi-CR6WE1xi.mjs';
import '@primeuix/utils';
import '@primeuix/utils/dom';
import '@primeuix/utils/object';
import './index-B6Gs2AMF.mjs';
import './index-l8XG4jwb.mjs';
import './index-Dedjwyps.mjs';
import './index-DrzwqwtL.mjs';
import '@primeuix/styles/inputnumber';
import '@primeuix/styles/inputtext';
import '@primeuix/utils/zindex';
import './index-DzyF1dFx.mjs';
import './index-Din928lO.mjs';
import '@primeuix/styled';
import '@primeuix/styles/dialog';
import './index-Bg3yUxEP.mjs';
import './index-BGjNF9Dd.mjs';
import '@primeuix/styles/paginator';
import './index-mL8pBhlS.mjs';
import './index-BdDxcS0p.mjs';
import './index-D2H2HXHf.mjs';
import './index-D4owPo0H.mjs';
import './index-D255zs2l.mjs';
import './index-83EVY8bO.mjs';
import '@primeuix/styles/iconfield';
import './index-IyD_XvNl.mjs';
import './index-rAVNvoJo.mjs';
import '@primeuix/utils/eventbus';
import './index-C3Ej0YBE.mjs';
import '@primeuix/styles/virtualscroller';
import '@primeuix/styles/select';
import './index-BsfLo0vt.mjs';
import '@primeuix/styles/datatable';
import './index-DrH3FigG.mjs';
import './index-Dik5Ht0d.mjs';
import './index-Ccf11gW7.mjs';
import './index-BkCm0XJl.mjs';
import '@primeuix/styles/checkbox';
import './index-bmcwiX6i.mjs';
import '@primeuix/styles/radiobutton';
import './index-CvVdVfGD.mjs';
import './index-0q6EZZjm.mjs';
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
import '@primeuix/styles/cascadeselect';
import '@primeuix/styles/checkboxgroup';
import '@primeuix/styles/colorpicker';
import '@primeuix/styles/datepicker';
import '@primeuix/styles/floatlabel';
import '@primeuix/styles/iftalabel';
import '@primeuix/styles/inputchips';
import '@primeuix/styles/inputgroup';
import '@primeuix/styles/inputotp';
import '@primeuix/styles/knob';
import '@primeuix/styles/listbox';
import '@primeuix/styles/multiselect';
import '@primeuix/styles/password';
import '@primeuix/styles/radiobuttongroup';
import '@primeuix/styles/rating';
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
import '@primeuix/styles/dataview';
import '@primeuix/styles/orderlist';
import '@primeuix/styles/organizationchart';
import '@primeuix/styles/picklist';
import '@primeuix/styles/tree';
import '@primeuix/styles/treetable';
import '@primeuix/styles/timeline';
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
import 'unhead';
import '@unhead/shared';
import 'vue-router';
import '@primeuix/styles/base';
import '@primeuix/utils/uuid';

const _sfc_main = {
  __name: "products",
  __ssrInlineRender: true,
  setup(__props) {
    const products = ref([]);
    const loading = ref(true);
    const showModal = ref(false);
    const editingProduct = ref(null);
    const form = ref({ name: "", price: 0, category_id: "", description: "", stock: 0 });
    const { admin: adminApi } = useApi();
    const fetchProducts = async () => {
      loading.value = true;
      try {
        const response = await adminApi.products.list();
        products.value = response.data || response || [];
      } catch (err) {
        console.error("Failed to fetch products", err);
      } finally {
        loading.value = false;
      }
    };
    const editProduct = (product) => {
      editingProduct.value = product;
      form.value = { ...product };
      showModal.value = true;
    };
    const saveProduct = async () => {
      try {
        if (editingProduct.value) {
          await adminApi.products.update(editingProduct.value.product_id, form.value);
        } else {
          await adminApi.products.create(form.value);
        }
        showModal.value = false;
        await fetchProducts();
      } catch (err) {
        console.error("Failed to save product", err);
      }
    };
    const resetForm = () => {
      editingProduct.value = null;
      form.value = { name: "", price: 0, category_id: "", description: "", stock: 0 };
    };
    return (_ctx, _push, _parent, _attrs) => {
      const _component_Button = script;
      const _component_DataTable = script$1;
      const _component_Column = script$2;
      const _component_Dialog = script$3;
      const _component_InputText = script$4;
      const _component_InputNumber = script$5;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "p-6" }, _attrs))}><div class="flex items-center justify-between mb-8"><h1 class="text-3xl font-bold text-on_surface">Manage Products</h1>`);
      _push(ssrRenderComponent(_component_Button, {
        onClick: ($event) => {
          showModal.value = true;
          resetForm();
        },
        label: "Add Product",
        icon: "pi pi-plus"
      }, null, _parent));
      _push(`</div>`);
      if (unref(loading)) {
        _push(`<div class="space-y-3"><!--[-->`);
        ssrRenderList(5, (n) => {
          _push(`<div class="h-16 bg-surface-container rounded-lg animate-pulse"></div>`);
        });
        _push(`<!--]--></div>`);
      } else {
        _push(`<div class="bg-surface-container-lowest rounded-xl border border-outline-variant/20 overflow-hidden">`);
        _push(ssrRenderComponent(_component_DataTable, {
          value: unref(products),
          stripedRows: "",
          class: "w-full"
        }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(ssrRenderComponent(_component_Column, {
                field: "name",
                header: "Name"
              }, null, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_Column, { header: "Price" }, {
                body: withCtx(({ data }, _push3, _parent3, _scopeId2) => {
                  var _a, _b;
                  if (_push3) {
                    _push3(`$${ssrInterpolate((_a = data.price) == null ? void 0 : _a.toFixed(2))}`);
                  } else {
                    return [
                      createTextVNode("$" + toDisplayString((_b = data.price) == null ? void 0 : _b.toFixed(2)), 1)
                    ];
                  }
                }),
                _: 1
              }, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_Column, {
                field: "stock",
                header: "Stock"
              }, null, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_Column, { header: "Actions" }, {
                body: withCtx(({ data }, _push3, _parent3, _scopeId2) => {
                  if (_push3) {
                    _push3(ssrRenderComponent(_component_Button, {
                      onClick: ($event) => editProduct(data),
                      icon: "pi pi-pencil",
                      text: "",
                      rounded: "",
                      size: "small"
                    }, null, _parent3, _scopeId2));
                  } else {
                    return [
                      createVNode(_component_Button, {
                        onClick: ($event) => editProduct(data),
                        icon: "pi pi-pencil",
                        text: "",
                        rounded: "",
                        size: "small"
                      }, null, 8, ["onClick"])
                    ];
                  }
                }),
                _: 1
              }, _parent2, _scopeId));
            } else {
              return [
                createVNode(_component_Column, {
                  field: "name",
                  header: "Name"
                }),
                createVNode(_component_Column, { header: "Price" }, {
                  body: withCtx(({ data }) => {
                    var _a;
                    return [
                      createTextVNode("$" + toDisplayString((_a = data.price) == null ? void 0 : _a.toFixed(2)), 1)
                    ];
                  }),
                  _: 1
                }),
                createVNode(_component_Column, {
                  field: "stock",
                  header: "Stock"
                }),
                createVNode(_component_Column, { header: "Actions" }, {
                  body: withCtx(({ data }) => [
                    createVNode(_component_Button, {
                      onClick: ($event) => editProduct(data),
                      icon: "pi pi-pencil",
                      text: "",
                      rounded: "",
                      size: "small"
                    }, null, 8, ["onClick"])
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
      _push(ssrRenderComponent(_component_Dialog, {
        visible: unref(showModal),
        "onUpdate:visible": ($event) => isRef(showModal) ? showModal.value = $event : null,
        header: unref(editingProduct) ? "Edit Product" : "Add Product",
        modal: "",
        class: "w-full max-w-md"
      }, {
        footer: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(ssrRenderComponent(_component_Button, {
              onClick: ($event) => showModal.value = false,
              label: "Cancel",
              text: ""
            }, null, _parent2, _scopeId));
            _push2(ssrRenderComponent(_component_Button, {
              onClick: saveProduct,
              label: "Save",
              icon: "pi pi-check"
            }, null, _parent2, _scopeId));
          } else {
            return [
              createVNode(_component_Button, {
                onClick: ($event) => showModal.value = false,
                label: "Cancel",
                text: ""
              }, null, 8, ["onClick"]),
              createVNode(_component_Button, {
                onClick: saveProduct,
                label: "Save",
                icon: "pi pi-check"
              })
            ];
          }
        }),
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`<div class="space-y-4 p-4"${_scopeId}><div${_scopeId}><label class="block text-sm font-medium text-on_surface mb-1"${_scopeId}>Name</label>`);
            _push2(ssrRenderComponent(_component_InputText, {
              modelValue: unref(form).name,
              "onUpdate:modelValue": ($event) => unref(form).name = $event,
              class: "w-full",
              placeholder: "Product Name"
            }, null, _parent2, _scopeId));
            _push2(`</div><div${_scopeId}><label class="block text-sm font-medium text-on_surface mb-1"${_scopeId}>Price</label>`);
            _push2(ssrRenderComponent(_component_InputNumber, {
              modelValue: unref(form).price,
              "onUpdate:modelValue": ($event) => unref(form).price = $event,
              mode: "currency",
              currency: "USD",
              locale: "en-US",
              class: "w-full"
            }, null, _parent2, _scopeId));
            _push2(`</div><div${_scopeId}><label class="block text-sm font-medium text-on_surface mb-1"${_scopeId}>Stock</label>`);
            _push2(ssrRenderComponent(_component_InputNumber, {
              modelValue: unref(form).stock,
              "onUpdate:modelValue": ($event) => unref(form).stock = $event,
              class: "w-full"
            }, null, _parent2, _scopeId));
            _push2(`</div></div>`);
          } else {
            return [
              createVNode("div", { class: "space-y-4 p-4" }, [
                createVNode("div", null, [
                  createVNode("label", { class: "block text-sm font-medium text-on_surface mb-1" }, "Name"),
                  createVNode(_component_InputText, {
                    modelValue: unref(form).name,
                    "onUpdate:modelValue": ($event) => unref(form).name = $event,
                    class: "w-full",
                    placeholder: "Product Name"
                  }, null, 8, ["modelValue", "onUpdate:modelValue"])
                ]),
                createVNode("div", null, [
                  createVNode("label", { class: "block text-sm font-medium text-on_surface mb-1" }, "Price"),
                  createVNode(_component_InputNumber, {
                    modelValue: unref(form).price,
                    "onUpdate:modelValue": ($event) => unref(form).price = $event,
                    mode: "currency",
                    currency: "USD",
                    locale: "en-US",
                    class: "w-full"
                  }, null, 8, ["modelValue", "onUpdate:modelValue"])
                ]),
                createVNode("div", null, [
                  createVNode("label", { class: "block text-sm font-medium text-on_surface mb-1" }, "Stock"),
                  createVNode(_component_InputNumber, {
                    modelValue: unref(form).stock,
                    "onUpdate:modelValue": ($event) => unref(form).stock = $event,
                    class: "w-full"
                  }, null, 8, ["modelValue", "onUpdate:modelValue"])
                ])
              ])
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`</div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/admin/products.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};

export { _sfc_main as default };
//# sourceMappingURL=products-B-c8mJ4i.mjs.map
