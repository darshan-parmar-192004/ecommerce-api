import { mergeProps, openBlock, createBlock, resolveDynamicComponent, withCtx, renderSlot, normalizeClass } from "vue";
import { B as BaseStyle, t as script$2 } from "../server.mjs";
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
    var instance = _ref.instance, props = _ref.props;
    return ["p-accordionpanel", {
      "p-accordionpanel-active": instance.active,
      "p-disabled": props.disabled
    }];
  }
};
var AccordionPanelStyle = BaseStyle.extend({
  name: "accordionpanel",
  classes
});
var script$1 = {
  name: "BaseAccordionPanel",
  "extends": script$2,
  props: {
    value: {
      type: [String, Number],
      "default": void 0
    },
    disabled: {
      type: Boolean,
      "default": false
    },
    as: {
      type: [String, Object],
      "default": "DIV"
    },
    asChild: {
      type: Boolean,
      "default": false
    }
  },
  style: AccordionPanelStyle,
  provide: function provide() {
    return {
      $pcAccordionPanel: this,
      $parentInstance: this
    };
  }
};
var script = {
  name: "AccordionPanel",
  "extends": script$1,
  inheritAttrs: false,
  inject: ["$pcAccordion"],
  computed: {
    active: function active() {
      return this.$pcAccordion.isItemActive(this.value);
    },
    attrs: function attrs() {
      return mergeProps(this.a11yAttrs, this.ptmi("root", this.ptParams));
    },
    a11yAttrs: function a11yAttrs() {
      return {
        "data-pc-name": "accordionpanel",
        "data-p-disabled": this.disabled,
        "data-p-active": this.active
      };
    },
    ptParams: function ptParams() {
      return {
        context: {
          active: this.active
        }
      };
    }
  }
};
function render(_ctx, _cache, $props, $setup, $data, $options) {
  return !_ctx.asChild ? (openBlock(), createBlock(resolveDynamicComponent(_ctx.as), mergeProps({
    key: 0,
    "class": _ctx.cx("root")
  }, $options.attrs), {
    "default": withCtx(function() {
      return [renderSlot(_ctx.$slots, "default")];
    }),
    _: 3
  }, 16, ["class"])) : renderSlot(_ctx.$slots, "default", {
    key: 1,
    "class": normalizeClass(_ctx.cx("root")),
    active: $options.active,
    a11yAttrs: $options.a11yAttrs
  });
}
script.render = render;
export {
  script as default
};
//# sourceMappingURL=index-BsM5ypMj.js.map
