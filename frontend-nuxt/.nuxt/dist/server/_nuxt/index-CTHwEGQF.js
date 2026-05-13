import { B as BaseStyle, t as script$2, G as script$3, J as script$4, w as script$5, K as script$6 } from "../server.mjs";
import { style } from "@primeuix/styles/inlinemessage";
import { openBlock, createElementBlock, mergeProps, renderSlot, createBlock, resolveDynamicComponent, createCommentVNode } from "vue";
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
    var props = _ref.props, instance = _ref.instance;
    return ["p-inlinemessage p-component p-inlinemessage-" + props.severity, {
      "p-inlinemessage-icon-only": !instance.$slots["default"]
    }];
  },
  icon: function icon(_ref2) {
    var props = _ref2.props;
    return ["p-inlinemessage-icon", props.icon];
  },
  text: "p-inlinemessage-text"
};
var InlineMessageStyle = BaseStyle.extend({
  name: "inlinemessage",
  style,
  classes
});
var script$1 = {
  name: "BaseInlineMessage",
  "extends": script$2,
  props: {
    severity: {
      type: String,
      "default": "error"
    },
    icon: {
      type: String,
      "default": void 0
    }
  },
  style: InlineMessageStyle,
  provide: function provide() {
    return {
      $pcInlineMessage: this,
      $parentInstance: this
    };
  }
};
var script = {
  name: "InlineMessage",
  "extends": script$1,
  inheritAttrs: false,
  timeout: null,
  data: function data() {
    return {
      visible: true
    };
  },
  mounted: function mounted() {
    var _this = this;
    if (!this.sticky) {
      setTimeout(function() {
        _this.visible = false;
      }, this.life);
    }
  },
  computed: {
    iconComponent: function iconComponent() {
      return {
        info: script$6,
        success: script$5,
        warn: script$4,
        error: script$3
      }[this.severity];
    }
  }
};
function render(_ctx, _cache, $props, $setup, $data, $options) {
  return openBlock(), createElementBlock("div", mergeProps({
    role: "alert",
    "aria-live": "assertive",
    "aria-atomic": "true",
    "class": _ctx.cx("root")
  }, _ctx.ptmi("root")), [renderSlot(_ctx.$slots, "icon", {}, function() {
    return [(openBlock(), createBlock(resolveDynamicComponent(_ctx.icon ? "span" : $options.iconComponent), mergeProps({
      "class": _ctx.cx("icon")
    }, _ctx.ptm("icon")), null, 16, ["class"]))];
  }), _ctx.$slots["default"] ? (openBlock(), createElementBlock("div", mergeProps({
    key: 0,
    "class": _ctx.cx("text")
  }, _ctx.ptm("text")), [renderSlot(_ctx.$slots, "default")], 16)) : createCommentVNode("", true)], 16);
}
script.render = render;
export {
  script as default
};
//# sourceMappingURL=index-CTHwEGQF.js.map
