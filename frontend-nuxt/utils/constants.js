export const storageKeys = {
  authToken: "auth_token",
  authUser: "auth_user",
  cart: "cart",
  savedForLater: "saved_for_later",
};

export const pagination = {
  defaultPage: 1,
  defaultLimit: 12,
  pageSizes: [12, 24, 36, 48],
};

export const toast = {
  defaultDuration: 4000,
  maxToasts: 5,
};

export const fallbacks = {
  customerId: 0,
  imagePlaceholder: "/placeholder.jpg",
  imagePlaceholderSvg: "/placeholder.svg",
};

export const animation = {
  staggerDelay: 50,
  drawerDuration: 300,
  cartItemDuration: 400,
};

export const routes = {
  home: "/",
  login: "/auth/login",
  register: "/auth/register",
  unauthorized: "/unauthorized",
  notFound: "/404",
  adminDashboard: "/admin",
};

export const roles = {
  customer: "customer",
  admin: "admin",
};

export const passwordPolicy = {
  minLength: 8,
};
