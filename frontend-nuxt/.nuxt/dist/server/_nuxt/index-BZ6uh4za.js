import { B as BaseStyle, t as script$2 } from "../server.mjs";
import { openBlock, createElementBlock, mergeProps, renderSlot } from "vue";
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
var classes = {
  root: function root(_ref) {
    var instance = _ref.instance;
    return ["p-splitterpanel", {
      "p-splitterpanel-nested": instance.isNested
    }];
  }
};
var SplitterPanelStyle = BaseStyle.extend({
  name: "splitterpanel",
  classes
});
var script$1 = {
  name: "BaseSplitterPanel",
  "extends": script$2,
  props: {
    size: {
      type: Number,
      "default": null
    },
    minSize: {
      type: Number,
      "default": null
    }
  },
  style: SplitterPanelStyle,
  provide: function provide() {
    return {
      $pcSplitterPanel: this,
      $parentInstance: this
    };
  }
};
var script = {
  name: "SplitterPanel",
  "extends": script$1,
  inheritAttrs: false,
  data: function data() {
    return {
      nestedState: null
    };
  },
  computed: {
    isNested: function isNested() {
      var _this = this;
      return this.$slots["default"]().some(function(child) {
        _this.nestedState = child.type.name === "Splitter" ? true : null;
        return _this.nestedState;
      });
    },
    getPTOptions: function getPTOptions() {
      return {
        context: {
          nested: this.isNested
        }
      };
    }
  }
};
function render(_ctx, _cache, $props, $setup, $data, $options) {
  return openBlock(), createElementBlock("div", mergeProps({
    ref: "container",
    "class": _ctx.cx("root")
  }, _ctx.ptmi("root", $options.getPTOptions)), [renderSlot(_ctx.$slots, "default")], 16);
}
script.render = render;
export {
  script as default
};
//# sourceMappingURL=index-BZ6uh4za.js.map
