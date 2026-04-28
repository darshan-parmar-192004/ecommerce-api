// In development, API calls go to backend directly (via Vite dev server)
// In production, this should be set to the API base URL (e.g., window.location.origin)
const API_BASE = import.meta.env.PROD ? '' : 'http://127.0.0.1:8080';

let _isAuthenticated = false;
let _authCheckDone = false;
let _userName = '';
let _userEmail = '';

export function isAuthenticated() {
  return _isAuthenticated;
}

export function getUserName() {
  return _userName;
}

export function getUserEmail() {
  return _userEmail;
}

export async function checkAuth() {
  try {
    const data = await fetchJSON("/auth/me", { method: "GET" });
    _isAuthenticated = true;
    _userName = data.name || '';
    _userEmail = data.email || '';
  } catch (err) {
    _isAuthenticated = false;
    _userName = '';
    _userEmail = '';
  }
  _authCheckDone = true;
  return _isAuthenticated;
}

export function resetAuthState() {
  _isAuthenticated = false;
  _authCheckDone = false;
  _userName = '';
  _userEmail = '';
}

export class ApiError extends Error {
  constructor(message, status, code = null) {
    super(message);
    this.name = "ApiError";
    this.status = status;
    this.code = code;
  }
}

function getErrorMessage(error, status) {
  const messages = {
    400: "Invalid request. Please check your input.",
    401: "Please log in to continue.",
    403: "You do not have permission to perform this action.",
    404: "The requested resource was not found.",
    409: "This item already exists.",
    422: "Validation failed. Please check your input.",
    429: "Too many requests. Please try again later.",
    500: "Server error. Please try again later.",
    503: "Service temporarily unavailable.",
  };

  if (messages[status]) {
    return messages[status];
  }

  const errorStr = String(error || "");

  if (errorStr.includes("network") || errorStr.includes("fetch")) {
    return "Network error. Please check your connection.";
  }

  return errorStr || "An unexpected error occurred.";
}

async function fetchJSON(url, options = {}) {
  const headers = {
    "Content-Type": "application/json",
    ...options.headers,
  };

  try {
    const response = await fetch(`${API_BASE}${url}`, {
      ...options,
      headers,
      credentials: "include",
    });

    const data = await response.json().catch(() => ({}));

    if (!response.ok) {
      let message = "Request failed";
      let code = null;

      if (data.error && typeof data.error === "object") {
        message = data.error.message || "Request failed";
        code = data.error.code || null;
      } else if (data.message) {
        message = data.message;
      } else if (data.error) {
        message = String(data.error);
      }

      if (response.status === 401) {
        window.dispatchEvent(new CustomEvent("auth:change"));
      }

      throw new ApiError(message, response.status, code);
    }

    return data;
  } catch (err) {
    if (err instanceof ApiError) {
      throw err;
    }
    throw new ApiError(getErrorMessage(err.message), 0);
  }
}

export const api = {
  products: {
    list: (params = {}) => {
      const query = new URLSearchParams(params).toString();
      return fetchJSON(`/products${query ? `?${query}` : ""}`);
    },
    get: (id) => fetchJSON(`/products/${id}`),
    create: (data) =>
      fetchJSON("/products", { method: "POST", body: JSON.stringify(data) }),
    update: (id, data) =>
      fetchJSON(`/products/${id}`, {
        method: "PUT",
        body: JSON.stringify(data),
      }),
    delete: (id) => fetchJSON(`/products/${id}`, { method: "DELETE" }),
  },
  categories: {
    list: () => fetchJSON("/categories"),
    get: (id) => fetchJSON(`/categories/${id}`),
    products: (id) => fetchJSON(`/categories/${id}/products`),
  },
  auth: {
    register: (data) =>
      fetchJSON("/auth/register", {
        method: "POST",
        body: JSON.stringify(data),
      }),
    login: (data) =>
      fetchJSON("/auth/login", {
        method: "POST",
        body: JSON.stringify(data),
      }),
    logout: async () => {
      try {
        await fetchJSON("/auth/logout", { method: "POST" });
      } finally {
        window.dispatchEvent(new CustomEvent("auth:change"));
      }
    },
    me: () => fetchJSON("/auth/me"),
  },
  orders: {
    create: (data) =>
      fetchJSON("/orders", { method: "POST", body: JSON.stringify(data) }),
    get: (id) => fetchJSON(`/orders/${id}`),
    list: () => fetchJSON("/orders"),
    cancel: (id) => fetchJSON(`/orders/${id}/cancel`, { method: "POST" }),
  },
};
