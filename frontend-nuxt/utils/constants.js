export const STORAGE_KEYS = {
  AUTH_TOKEN: "auth_token",
  AUTH_USER: "auth_user",
  CART: "cart",
  SAVED_FOR_LATER: "saved_for_later",
};

export const PAGINATION = {
  DEFAULT_PAGE: 1,
  DEFAULT_LIMIT: 12,
  PAGE_SIZES: [12, 24, 36, 48],
};

export const TOAST = {
  DEFAULT_DURATION: 4000,
  MAX_TOASTS: 5,
};

export const FALLBACKS = {
  CUSTOMER_ID: 0,
  IMAGE_PLACEHOLDER: "/placeholder.jpg",
  IMAGE_PLACEHOLDER_SVG: "/placeholder.svg",
};

export const ANIMATION = {
  STAGGER_DELAY: 50,
  DRAWER_DURATION: 300,
  CART_ITEM_DURATION: 400,
};

export const ROUTES = {
  HOME: "/",
  LOGIN: "/auth/login",
  REGISTER: "/auth/register",
  UNAUTHORIZED: "/unauthorized",
  NOT_FOUND: "/404",
  ADMIN_DASHBOARD: "/admin",
};

export const ROLES = {
  CUSTOMER: "customer",
  ADMIN: "admin",
};

export const PASSWORD_POLICY = {
  MIN_LENGTH: 8,
};
