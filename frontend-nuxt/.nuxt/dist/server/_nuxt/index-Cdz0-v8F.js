import { B as BaseStyle, t as script$1 } from "../server.mjs";
import { openBlock, createElementBlock, mergeProps, renderSlot, createCommentVNode } from "vue";
import "ofetch";
import "#internal/nuxt/paths";
import "hookable";
import "unctx";
import "h3";
import "ufo";
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
import "vue/server-renderer";
import "@primeuix/utils/zindex";
import "@primeuix/styles/toast";
import "@primeuix/utils/uuid";
import "@primeuix/styles/ripple";
import "@primeuix/styles/badge";
import "@primeuix/styles/button";
var DeferredContentStyle = BaseStyle.extend({
  name: "deferredcontent"
});
var script = {
  name: "DeferredContent",
  "extends": script$1,
  inheritAttrs: false,
  emits: ["load"],
  style: DeferredContentStyle,
  data: function data() {
    return {
      loaded: false
    };
  },
  mounted: function mounted() {
    if (!this.loaded) {
      if (this.shouldLoad()) this.load();
      else this.bindScrollListener();
    }
  },
  beforeUnmount: function beforeUnmount() {
    this.unbindScrollListener();
  },
  methods: {
    bindScrollListener: function bindScrollListener() {
      var _this = this;
      this.documentScrollListener = function() {
        if (_this.shouldLoad()) {
          _this.load();
          _this.unbindScrollListener();
        }
      };
      (void 0).addEventListener("scroll", this.documentScrollListener);
    },
    unbindScrollListener: function unbindScrollListener() {
      if (this.documentScrollListener) {
        (void 0).removeEventListener("scroll", this.documentScrollListener);
        this.documentScrollListener = null;
      }
    },
    shouldLoad: function shouldLoad() {
      if (this.loaded) {
        return false;
      } else {
        var rect = this.$refs.container.getBoundingClientRect();
        var docElement = (void 0).documentElement;
        var winHeight = docElement.clientHeight;
        return winHeight >= rect.top;
      }
    },
    load: function load(event) {
      this.loaded = true;
      this.$emit("load", event);
    }
  }
};
function render(_ctx, _cache, $props, $setup, $data, $options) {
  return openBlock(), createElementBlock("div", mergeProps({
    ref: "container"
  }, _ctx.ptmi("root")), [$data.loaded ? renderSlot(_ctx.$slots, "default", {
    key: 0
  }) : createCommentVNode("", true)], 16);
}
script.render = render;
export {
  script as default
};
//# sourceMappingURL=index-Cdz0-v8F.js.map
