import { getWindowScrollTop } from "@primeuix/utils/dom";
import { ZIndex } from "@primeuix/utils/zindex";
import { s as script$3 } from "./index-BRkAJkA6.js";
import { B as BaseStyle, s as script$2, t as script$4 } from "../server.mjs";
import { style } from "@primeuix/styles/scrolltop";
import { resolveComponent, openBlock, createBlock, Transition, mergeProps, withCtx, renderSlot, normalizeClass, resolveDynamicComponent, createCommentVNode } from "vue";
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
import "cookie-es";
import "destr";
import "ohash";
import "vue/server-renderer";
import "@primeuix/styles/toast";
import "@primeuix/utils/uuid";
import "@primeuix/styles/ripple";
import "@primeuix/styles/badge";
import "@primeuix/styles/button";
var classes = {
  root: function root(_ref) {
    var props = _ref.props;
    return ["p-scrolltop", {
      "p-scrolltop-sticky": props.target !== "window"
    }];
  },
  icon: "p-scrolltop-icon"
};
var ScrollTopStyle = BaseStyle.extend({
  name: "scrolltop",
  style,
  classes
});
var script$1 = {
  name: "BaseScrollTop",
  "extends": script$4,
  props: {
    target: {
      type: String,
      "default": "window"
    },
    threshold: {
      type: Number,
      "default": 400
    },
    icon: {
      type: String,
      "default": void 0
    },
    behavior: {
      type: String,
      "default": "smooth"
    },
    buttonProps: {
      type: Object,
      "default": function _default() {
        return {
          rounded: true
        };
      }
    }
  },
  style: ScrollTopStyle,
  provide: function provide() {
    return {
      $pcScrollTop: this,
      $parentInstance: this
    };
  }
};
var script = {
  name: "ScrollTop",
  "extends": script$1,
  inheritAttrs: false,
  scrollListener: null,
  container: null,
  data: function data() {
    return {
      visible: false
    };
  },
  mounted: function mounted() {
    if (this.target === "window") this.bindDocumentScrollListener();
    else if (this.target === "parent") this.bindParentScrollListener();
  },
  beforeUnmount: function beforeUnmount() {
    if (this.target === "window") this.unbindDocumentScrollListener();
    else if (this.target === "parent") this.unbindParentScrollListener();
    if (this.container) {
      ZIndex.clear(this.container);
      this.overlay = null;
    }
  },
  methods: {
    onClick: function onClick() {
      var scrollElement = this.target === "window" ? void 0 : this.$el.parentElement;
      scrollElement.scroll({
        top: 0,
        behavior: this.behavior
      });
    },
    checkVisibility: function checkVisibility(scrollY) {
      if (scrollY > this.threshold) this.visible = true;
      else this.visible = false;
    },
    bindParentScrollListener: function bindParentScrollListener() {
      var _this = this;
      this.scrollListener = function() {
        _this.checkVisibility(_this.$el.parentElement.scrollTop);
      };
      this.$el.parentElement.addEventListener("scroll", this.scrollListener);
    },
    bindDocumentScrollListener: function bindDocumentScrollListener() {
      var _this2 = this;
      this.scrollListener = function() {
        _this2.checkVisibility(getWindowScrollTop());
      };
      (void 0).addEventListener("scroll", this.scrollListener);
    },
    unbindParentScrollListener: function unbindParentScrollListener() {
      if (this.scrollListener) {
        this.$el.parentElement.removeEventListener("scroll", this.scrollListener);
        this.scrollListener = null;
      }
    },
    unbindDocumentScrollListener: function unbindDocumentScrollListener() {
      if (this.scrollListener) {
        (void 0).removeEventListener("scroll", this.scrollListener);
        this.scrollListener = null;
      }
    },
    onEnter: function onEnter(el) {
      ZIndex.set("overlay", el, this.$primevue.config.zIndex.overlay);
    },
    onAfterLeave: function onAfterLeave(el) {
      ZIndex.clear(el);
    },
    containerRef: function containerRef(el) {
      this.container = el ? el.$el : void 0;
    }
  },
  computed: {
    scrollTopAriaLabel: function scrollTopAriaLabel() {
      return this.$primevue.config.locale.aria ? this.$primevue.config.locale.aria.scrollTop : void 0;
    }
  },
  components: {
    ChevronUpIcon: script$3,
    Button: script$2
  }
};
function render(_ctx, _cache, $props, $setup, $data, $options) {
  var _component_Button = resolveComponent("Button");
  return openBlock(), createBlock(Transition, mergeProps({
    name: "p-scrolltop",
    appear: "",
    onEnter: $options.onEnter,
    onAfterLeave: $options.onAfterLeave
  }, _ctx.ptm("transition")), {
    "default": withCtx(function() {
      return [$data.visible ? (openBlock(), createBlock(_component_Button, mergeProps({
        key: 0,
        ref: $options.containerRef,
        "class": _ctx.cx("root"),
        onClick: $options.onClick,
        "aria-label": $options.scrollTopAriaLabel,
        unstyled: _ctx.unstyled
      }, _ctx.buttonProps, {
        pt: _ctx.ptm("root")
      }), {
        icon: withCtx(function(slotProps) {
          return [renderSlot(_ctx.$slots, "icon", {
            "class": normalizeClass(_ctx.cx("icon"))
          }, function() {
            return [(openBlock(), createBlock(resolveDynamicComponent(_ctx.icon ? "span" : "ChevronUpIcon"), mergeProps({
              "class": [_ctx.cx("icon"), _ctx.icon, slotProps["class"]]
            }, _ctx.ptm("root")["icon"], {
              "data-pc-section": "icon"
            }), null, 16, ["class"]))];
          })];
        }),
        _: 3
      }, 16, ["class", "onClick", "aria-label", "unstyled", "pt"])) : createCommentVNode("", true)];
    }),
    _: 3
  }, 16, ["onEnter", "onAfterLeave"]);
}
script.render = render;
export {
  script as default
};
//# sourceMappingURL=index-68K7r8Mj.js.map
