import { ref, computed, mergeProps, unref, useSSRContext } from 'vue';
import { ssrRenderAttrs, ssrRenderList, ssrInterpolate } from 'vue/server-renderer';

const _sfc_main = {
  __name: "categories",
  __ssrInlineRender: true,
  setup(__props) {
    const categories = ref([]);
    const loading = ref(true);
    const categoryTree = computed(() => {
      const tree = [];
      const categoryMap = {};
      categories.value.forEach((cat) => {
        categoryMap[cat.category_id] = { ...cat, children: [] };
      });
      categories.value.forEach((cat) => {
        if (cat.parent_category_id && categoryMap[cat.parent_category_id]) {
          categoryMap[cat.parent_category_id].children.push(categoryMap[cat.category_id]);
        } else {
          tree.push(categoryMap[cat.category_id]);
        }
      });
      return tree;
    });
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "p-6" }, _attrs))}><h1 class="text-3xl font-bold text-on_surface mb-8">Category Hierarchy</h1>`);
      if (unref(loading)) {
        _push(`<div class="space-y-3"><!--[-->`);
        ssrRenderList(5, (n) => {
          _push(`<div class="h-16 bg-surface-container rounded-xl animate-pulse"></div>`);
        });
        _push(`<!--]--></div>`);
      } else {
        _push(`<div class="bg-surface-container-lowest rounded-xl border border-outline-variant/20 p-6">`);
        if (unref(categoryTree).length === 0) {
          _push(`<div class="text-center py-8"><i class="pi pi-tags text-5xl text-outline mb-4"></i><p class="text-on_surface_variant">No categories found</p></div>`);
        } else {
          _push(`<div class="space-y-4"><!--[-->`);
          ssrRenderList(unref(categoryTree), (category) => {
            var _a;
            _push(`<div class="mb-4"><div class="flex items-center gap-3 p-3 bg-surface-container rounded-lg"><i class="pi pi-tag text-primary"></i><span class="font-medium text-on_surface">${ssrInterpolate(category.name)}</span><span class="text-xs text-outline ml-auto">${ssrInterpolate(category.category_id)}</span></div>`);
            if ((_a = category.children) == null ? void 0 : _a.length) {
              _push(`<div class="ml-8 mt-2 space-y-2"><!--[-->`);
              ssrRenderList(category.children, (child) => {
                _push(`<div class="flex items-center gap-3 p-2 border border-outline-variant/20 rounded-lg"><i class="pi pi-tag text-outline text-xs"></i><span class="text-on_surface">${ssrInterpolate(child.name)}</span><span class="text-xs text-outline ml-auto">${ssrInterpolate(child.category_id)}</span></div>`);
              });
              _push(`<!--]--></div>`);
            } else {
              _push(`<!---->`);
            }
            _push(`</div>`);
          });
          _push(`<!--]--></div>`);
        }
        _push(`</div>`);
      }
      _push(`</div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/admin/categories.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};

export { _sfc_main as default };
//# sourceMappingURL=categories-8xlc1Gr-.mjs.map
